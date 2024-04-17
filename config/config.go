package config

import (
	"html/template"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
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
