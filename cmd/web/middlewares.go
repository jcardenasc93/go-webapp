package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// GenCSRFToken creates middleware to generate CSRF token to secure POST requests
func GenCSRFToken(next http.Handler) http.Handler {
	tokenHandler := nosurf.New(next)
	tokenHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction, // dev env not uses https
		SameSite: app.SameSite,
	})
	return tokenHandler
}

// LoadSession load and save session for each request
func LoadSession(next http.Handler) http.Handler {
	return sessionMan.LoadAndSave(next)
}
