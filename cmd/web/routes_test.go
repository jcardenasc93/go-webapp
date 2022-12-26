package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/jcardenasc93/go-webapp/internal/config"
)

func TestRoutes(t *testing.T) {
	testApp := config.AppConfig{}

	mux := routing(&testApp)

	switch v := mux.(type) {
	case *chi.Mux:
		// Success
	default:
		t.Errorf("Expected *chi.Mux type but got %T instead", v)
	}
}
