package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/handlers"
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
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/majestic-suite", handlers.Repo.Majestic)
	mux.Get("/comfortable-place", handlers.Repo.Comfortable)

	mux.Post("/make-reservation", handlers.Repo.PostMakeReservation)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	mux.Get("/booking", handlers.Repo.Booking)
	mux.Post("/booking", handlers.Repo.PostBooking)
	mux.Post("/booking-json", handlers.Repo.BookingJSON)

	mux.Get("/booking-summary", handlers.Repo.BookingSummary)

	// File server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
