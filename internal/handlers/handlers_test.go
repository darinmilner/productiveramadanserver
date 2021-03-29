package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"Home", "/", "GET", []postData{}, http.StatusOK},
	{"About", "/about", "GET", []postData{}, http.StatusOK},
	{"Get All hadith", "/hadiths", "GET", []postData{}, http.StatusOK},
	{"Get All ayahs", "/ayahs", "GET", []postData{}, http.StatusOK},
	{"Get One hadith", "/hadiths/1", "GET", []postData{}, http.StatusOK},
	{"Get One ayah", "/ayahs/1", "GET", []postData{}, http.StatusOK},
	{"Get One hadith returns 404 with number out of range", "/hadiths/100", "GET", []postData{}, http.StatusNotFound},
	{"Get One ayah returns 404 with number out of range", "/ayahs/100", "GET", []postData{}, http.StatusNotFound},
}

func TestHandler(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)

	defer testServer.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := testServer.Client().Get(testServer.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal("Could not get URL", err)
			}

			if e.expectedStatusCode != resp.StatusCode {
				t.Errorf("Test %s expected %d but got code %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {

		}
	}
}
