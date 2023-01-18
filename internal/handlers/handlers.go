package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/config/helpers"
	"github.com/jcardenasc93/go-webapp/internal/models"
	"github.com/jcardenasc93/go-webapp/internal/models/forms"
	"github.com/jcardenasc93/go-webapp/internal/render"
)

// NOTE: Using repository pattern allows to swap componets around the app without huge changes in code base
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetupHandlers initialize Repository used by handlers
func SetupHandlers(r *Repository) {
	Repo = r
}

// NOTE: Using receivers grant func to access receiver data
// Home is the home handler
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Adds visitor remote address to session
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about handler
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Booking is the booking handler
func (rep *Repository) Booking(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{})
}

// Majestic is the majestic suite handler
func (rep *Repository) Majestic(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majestic-suite.page.tmpl", &models.TemplateData{})
}

// Comfortable is the comfortable place handler
func (rep *Repository) Comfortable(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "comfortable-place.page.tmpl", &models.TemplateData{})
}

// Contact is the comfortable place handler
func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// MakeReservation is the comfortable place handler
func (rep *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var initReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = initReservation

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostMakeReservation handles the post make reservation form
func (rep *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// reservation holds actual form state
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	// Makes form validation
	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLenght("first_name", 3)
	form.MinLenght("last_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		// Not valid form renders the form in its current state
		// with form errors
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// Put reservation in current session
	rep.App.Session.Put(r.Context(), "reservation", reservation)

	// Redirect
	http.Redirect(w, r, "/booking-summary", http.StatusSeeOther)
}

// PostBooking is handler to create a reservation
func (rep *Repository) PostBooking(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date: %s\nEnd date: %s", start, end)))
}

type availabilityResponse struct {
	OK  bool   `json:"ok"`
	Msg string `json:"message"`
}

// BookingJSON is handler to create a reservation
func (rep *Repository) BookingJSON(w http.ResponseWriter, r *http.Request) {
	resp := availabilityResponse{
		OK:  true,
		Msg: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rep.App.InfoLog.Println(out)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// BookingSummary is handler for display a reservation summary
func (rep *Repository) BookingSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := rep.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		rep.App.ErrorLog.Println("Cannot retrieve a valid reservation")
		rep.App.Session.Put(r.Context(), "error", "Cannot retrieve a valid reservation")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	rep.App.Session.Remove(r.Context(), "reservation") // Removes reservation after successfuly stored
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "booking-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
