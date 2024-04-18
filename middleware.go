package main

import (
	"fmt"
	"net/http"
)

func WriteToLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hitting the middleware...")
		next.ServeHTTP(w, r)
	})
}

func SessionHandler(next http.Handler) http.Handler {
	return app.SessionManager.LoadAndSave(next)
}
