package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jcardenasc93/go-webapp/pkg/config"
	"github.com/jcardenasc93/go-webapp/pkg/handlers"
)

func routing(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	// Adds middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(GenCSRFToken)
	mux.Use(LoadSession)
	// Adds routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
