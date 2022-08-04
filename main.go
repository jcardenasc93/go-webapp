package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//NOTE: This is the way to handle requests with builtin go libs
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		n, err := fmt.Fprintf(w, "Hello from GO server")
		log.Println(fmt.Sprintf("Number of written bytes: %d", n))
		if err != nil {
			log.Panic(err)
		}
	})

	//NOTE: Here starts the server
	http.ListenAndServe(":8000", nil)
}
