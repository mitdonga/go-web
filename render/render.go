package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mitdonga/go-web/config"
)

// Defined functions to use inside template
var functions = template.FuncMap{}
var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, fileName string) {
	// appConfig := config.AppConfig()
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	ts, ok := templateCache[fileName]
	if !ok {
		log.Fatalln("Template cache not found for file ", fileName)
	}

	buffBytes := new(bytes.Buffer)
	err := ts.Execute(buffBytes, nil)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = buffBytes.WriteTo(w)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		fmt.Println("Found page", page)
		file_name := filepath.Base(page)

		// create template set
		ts, err := template.New(file_name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return nil, err
			}
		}
		myCache[file_name] = ts
	}
	return myCache, nil
}
