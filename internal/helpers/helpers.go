package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/darinmilner/productiveapp/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client Error with a status of", status)

	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.ErrorLog.Println(trace)
}
