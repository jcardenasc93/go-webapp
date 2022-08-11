package main

import (
	"fmt"
	"github.com/jcardenasc93/go-webapp/pkg/handlers"
	"log"
	"net/http"
)

const serverPort = ":8000"

func main() {
	//NOTE: This is the way to handle requests with builtin go libs
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println(fmt.Sprintf("Starting server on %s", serverPort))

	//NOTE: Here starts the server
	http.ListenAndServe(serverPort, nil)
}
