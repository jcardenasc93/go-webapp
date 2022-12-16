package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/models"
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
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
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
		log.Println(err)
	}

	log.Println(out)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
