package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func GenCSRFToken(next http.Handler) http.Handler {
	tokenHandler := nosurf.New(next)
	tokenHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false, // dev env not uses https
		SameSite: http.SameSiteLaxMode,
	})
	return tokenHandler
}
