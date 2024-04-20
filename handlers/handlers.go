package handlers

import (
	"fmt"
	"net/http"

	"github.com/mitdonga/go-web/config"
	"github.com/mitdonga/go-web/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	rp.App.SessionManager.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl")
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := rp.App.SessionManager.GetString(r.Context(), "remote_ip")
	fmt.Println(remoteIp)

	render.RenderTemplate(w, "about.page.tmpl")
}
