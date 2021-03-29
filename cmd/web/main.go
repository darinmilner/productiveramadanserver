package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/darinmilner/productiveapp/internal/config"
	"github.com/darinmilner/productiveapp/internal/handlers"
	"github.com/darinmilner/productiveapp/internal/models"
	"github.com/darinmilner/productiveapp/internal/render"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//const portNumber = ":8001"

var session *scs.SessionManager
var app config.AppConfig

var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	//os.Setenv("PORT", "8001")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	portNumber := os.Getenv("PORT")
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

	//put into the session
	gob.Register(models.Signup{})

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

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	dbConnectionString := os.Getenv("DbConnectionString")
	clientOptions := options.Client().ApplyURI(dbConnectionString)
	config.Client, _ = mongo.Connect(ctx, clientOptions)

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	return nil
}
