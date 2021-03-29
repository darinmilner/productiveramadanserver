package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/darinmilner/productiveapp/internal/config"
	"github.com/darinmilner/productiveapp/internal/render"
	"github.com/go-chi/chi"
)

const pathToTemplates = "./../../templates"

var app config.AppConfig
var functions = template.FuncMap{}

func getRoutes() http.Handler {

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}

	//Change to true when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)

	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	//mux.Use(WriteToConsole) //Testing middleware

	hadithHandler := NewHadithHandlers()
	ayahHandler := NewAyahHandlers()
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/hadiths", hadithHandler.GetHadith)
	mux.Get("/ayahs", ayahHandler.GetAyahs)
	mux.Get("/hadiths/{id}", hadithHandler.GetHadith)
	mux.Get("/ayahs/{id}", ayahHandler.GetAyahs)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

//CreateTemplateCache creates a map of templateCache
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}
	return myCache, nil
}
