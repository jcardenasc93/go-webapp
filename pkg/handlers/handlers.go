package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// This method parse a given template in the ResponseWriter
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Error parsing template %s:\n", tmpl)
		log.Println(err)
	}
}

// Home is the home handler
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from GO server")
	renderTemplate(w, "home.page.tmpl")
}

// About is the about handler
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the about page")
	renderTemplate(w, "about.page.tmpl")
}
