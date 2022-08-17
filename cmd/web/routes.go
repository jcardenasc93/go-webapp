package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/jcardenasc93/go-webapp/pkg/config"
	"github.com/jcardenasc93/go-webapp/pkg/handlers"
)

func routing(a *config.AppConfig) http.Handler {
	mux := pat.New()
	// Adds routes
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
