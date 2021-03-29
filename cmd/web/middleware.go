package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF to all POSt requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, //TODO: Change to deployment mode
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//SessionLoad middleware loads and saves the session on each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
