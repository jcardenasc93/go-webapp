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
	remoteIP := r.RemoteAddr
	// Adds visitor remote address to session
	rep.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about handler
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rep.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["greetings"] = "Hello from Golang"
	stringMap["remoteIP"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
