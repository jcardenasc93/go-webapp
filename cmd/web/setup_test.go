package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run()) // Run tests
}

// mockHandler mocks http.Handler for testing purposes
type mockHandler struct{}

func (mh *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
