package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfig has the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
