package render

import (
	"net/http"
	"testing"

	"github.com/jcardenasc93/go-webapp/internal/models"
)

func TestAddDefaultTempData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Fatal("Cannot retrieve a valid session")
	}

	flashMsg := "flash test"
	sessionMan.Put(r.Context(), "flash", flashMsg)
	result := AddDefaultTempData(&td, r)
	if result.FlashMsg != flashMsg {
		t.Errorf("Flash msg set wrong. Expected %s but got %s", flashMsg, result.FlashMsg)
	}
}

func TestRenderTemplate(t *testing.T) {
	templsDir = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Fatal(err)
	}
	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww testRespWriter

	tmp := "home.page.tmpl"
	err = RenderTemplate(&ww, r, tmp, &models.TemplateData{})
	if err != nil {
		t.Errorf("Cannot render template %s", tmp)
	}

	noValidTmp := "novalid.page.tmpl"
	err = RenderTemplate(&ww, r, noValidTmp, &models.TemplateData{})
	if err == nil {
		t.Errorf("Unexistent template rendered %s", noValidTmp)
	}
}

func TestSetupTemplates(t *testing.T) {
	SetupTemplates(&testApp)
	if app == nil {
		t.Error("Error setting up app templates")
	}
}

func TestCreateTemplateCache(t *testing.T) {
	templsDir = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error("Create template cache FAILS")
	}
	if len(tc) <= 0 {
		t.Error("Cache templates from templates dir FAILS")
	}

	// Force error
	templsDir = "./fakedir"
	tc, _ = CreateTemplateCache()
	if len(tc) != 0 {
		t.Error("Templates cache created from not existing dir")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = sessionMan.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
