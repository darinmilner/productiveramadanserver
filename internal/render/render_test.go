package render

import (
	"fmt"
	"github/darinmilner/productiveramadanserver/internal/models"
	"net/http"
	"testing"
)

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/testurl", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
func TestAddDefaultData(t *testing.T) {
	var tData models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&tData, r)

	if result.FlashMsg != "123" {
		t.Error("Flash value 123 not found in session")
	}
}

func TestRenderTemplates(t *testing.T) {

	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	var testW myWriter

	err = RenderTemplates(&testW, r, "home.page.html", &models.TemplateData{})

	if err != nil {
		fmt.Print(err)
		t.Error("Error writing template data", err)
	}

	err = RenderTemplates(&testW, r, "not-here.page.html", &models.TemplateData{})

	if err == nil {
		t.Error("Returns nonexisting template page")
	}
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"

	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}
