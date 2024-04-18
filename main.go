package main

import (
	"fmt"
	"net/http"

	"github.com/mitdonga/go-web/config"
	"github.com/mitdonga/go-web/handlers"
	"github.com/mitdonga/go-web/render"
)

const serverPort = ":3000"

func main() {
	var app config.AppConfig

	ts, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println(err)
	}

	app.TemplateCache = ts
	app.UseCache = true

	render.SetAppConfig(&app)
	Repo := handlers.NewRepo(&app)
	handlers.NewHandler(Repo)

	srv := &http.Server{
		Addr:    serverPort,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
