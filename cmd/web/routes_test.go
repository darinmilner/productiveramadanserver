package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/darinmilner/productiveapp/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		log.Print("Test passed")
	default:
		t.Error(fmt.Sprintf("router is not a mux type it is type %t", v))
	}

}
