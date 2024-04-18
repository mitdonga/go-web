package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mitdonga/go-web/config"
	"github.com/mitdonga/go-web/handlers"
)

func routes(a *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(WriteToLog)
	mux.Use(SessionHandler)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
