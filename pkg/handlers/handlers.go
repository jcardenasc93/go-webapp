package handlers

import (
	"net/http"

	"github.com/jcardenasc93/go-webapp/pkg/render"
)

// Home is the home handler
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from GO server")
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about handler
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the about page")
	render.RenderTemplate(w, "about.page.tmpl")
}
