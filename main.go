package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs"
	"github.com/mitdonga/go-web/config"
	"github.com/mitdonga/go-web/handlers"
	"github.com/mitdonga/go-web/render"
)

const serverPort = ":3000"

var app config.AppConfig

func main() {

	ts, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println(err)
	}

	sessionManager := scs.New()
	sessionManager.Lifetime = 3 * time.Hour
	sessionManager.Cookie.Name = "go_session_id"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = app.IsProduction

	app.TemplateCache = ts
	app.UseCache = true
	app.IsProduction = false
	app.SessionManager = sessionManager

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
