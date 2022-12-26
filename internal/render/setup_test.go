package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"

	"github.com/alexedwards/scs/v2"
	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/models"
)

var sessionMan *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// Includes complex types to store in session
	gob.Register(models.Reservation{})

	// Session management
	sessionMan = config.InitSession(sessionMan)
	// Initilize app
	testApp.InitApp("", sessionMan)
	// Setup templates config
	testApp.UseCacheTemplates = false
	// Setup templates
	SetupTemplates(&testApp)

	os.Exit(m.Run())
}

// testRespWriter is interface to mock http.ResponseWriter for testing
type testRespWriter struct{}

func (tw *testRespWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *testRespWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (tw *testRespWriter) WriteHeader(i int) {
}
