package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const serverPort = ":8000"

// Home is the home handler
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from GO server")
	renderTemplate(w, "home.html")
}

// About is the about handler
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the about page")
	renderTemplate(w, "about.page.html")
}

func main() {
	//NOTE: This is the way to handle requests with builtin go libs
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println(fmt.Sprintf("Starting server on %s", serverPort))

	//NOTE: Here starts the server
	http.ListenAndServe(serverPort, nil)
}

// This method parse a given template in the ResponseWriter
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Error parsing template %s:\n", tmpl)
		log.Println(err)
	}
}
