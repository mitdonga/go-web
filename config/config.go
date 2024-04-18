package config

import (
	"html/template"

	"github.com/alexedwards/scs"
)

type AppConfig struct {
	TemplateCache  map[string]*template.Template
	UseCache       bool
	IsProduction   bool
	SessionManager *scs.SessionManager
}

// func AppConfig() config {
// 	var app config

// 	templateCache, err := createTemplateCache()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	app.templateCache = templateCache
// 	return app
// }
