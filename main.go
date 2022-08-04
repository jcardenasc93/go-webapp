package main

import (
	"fmt"
	"log"
	"net/http"
)

const serverPort = ":8000"

// Home is the home handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from GO server")
}

// About is the about handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func main() {
	//NOTE: This is the way to handle requests with builtin go libs
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println(fmt.Sprintf("Starting server on %s", serverPort))

	//NOTE: Here starts the server
	http.ListenAndServe(serverPort, nil)
}
