package main

import (
	"log"
	"net/http"

	"github.com/darinmilner/productiveapp/internal/config"
	"github.com/darinmilner/productiveapp/internal/handlers"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(WriteToConsole) //Testing middleware

	//Session middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	hadithHandler := handlers.NewHadithHandlers()
	ayahHandler := handlers.NewAyahHandlers()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/hadiths", hadithHandler.GetHadith)
	mux.Get("/ayahs", ayahHandler.GetAyahs)
	mux.Get("/hadiths/{id}", hadithHandler.GetHadith)
	mux.Get("/ayahs/{id}", ayahHandler.GetAyahs)
	mux.Get("/signup", handlers.Repo.Signup)
	mux.Post("/signup", handlers.Repo.PostSignUp)
	mux.Get("/signup-success", handlers.Repo.SignupSuccess)
	mux.Get("/*", handlers.Repo.DoesNotExistPage)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

//WriteToConsole middleware--USELESS-Just For Testing
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}
