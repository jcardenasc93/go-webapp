package main

import (
	"net/http"
	"testing"
)

func TestGenCSRFToken(t *testing.T) {
	mh := mockHandler{}
	tokenHandler := GenCSRFToken(&mh)
	switch th := tokenHandler.(type) {
	case http.Handler:
		// Success
	default:
		t.Errorf("Expected http.Handler type. But got %T", th)
	}
}

func TestLoadSession(t *testing.T) {
	mh := mockHandler{}
	session := LoadSession(&mh)
	switch ts := session.(type) {
	case http.Handler:
		// Success
	default:
		t.Errorf("Expected http.Handler type. But got %T", ts)
	}
}
