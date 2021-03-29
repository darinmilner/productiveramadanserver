package render

import (
	"github/darinmilner/productiveramadanserver/internal/config"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager

var testApp config.AppConfig

func TestMain(m *testing.M) {

	testApp.InProduction = true

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false //True in Production

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())

}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
