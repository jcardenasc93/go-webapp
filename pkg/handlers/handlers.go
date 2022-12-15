package handlers

import (
	"net/http"

	"github.com/jcardenasc93/go-webapp/pkg/config"
	"github.com/jcardenasc93/go-webapp/pkg/models"
	"github.com/jcardenasc93/go-webapp/pkg/render"
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about handler
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}

// Booking is the booking handler
func (rep *Repository) Booking(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "booking.page.tmpl", &models.TemplateData{})
}

// Majestic is the majestic suite handler
func (rep *Repository) Majestic(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majestic-suite.page.tmpl", &models.TemplateData{})
}

// Comfortable is the comfortable place handler
func (rep *Repository) Comfortable(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "comfortable-place.page.tmpl", &models.TemplateData{})
}

// Contact is the comfortable place handler
func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// MakeReservation is the comfortable place handler
func (rep *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// PostBooking is handler to create a reservation
func (rep *Repository) PostBooking(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
