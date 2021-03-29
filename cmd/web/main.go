package main

import (
	"fmt"
	"github/darinmilner/productiveramadanserver/internal/config"
	"github/darinmilner/productiveramadanserver/internal/handlers"
	"github/darinmilner/productiveramadanserver/internal/render"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/joho/godotenv"
)

//const portNumber = ":8001"

var session *scs.SessionManager
var app config.AppConfig

var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	portNumber := os.Getenv("PORT")

	if portNumber == "" {
		portNumber = "8000"
	}

	fmt.Print(portNumber)
	err = run()
	log.Println("Server running on port: ", portNumber)
	srv := &http.Server{
		Addr:    ":" + portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {

	app.InProduction = true

	//Info log
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //True in Production

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	return nil
}
