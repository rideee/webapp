package main

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/rideee/webapp/internal/tmpl"
)

// routes keeps all application routes.
func routes(mux *mux.Router, tpl *tmpl.Tmpl) {

	// Homepage.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Render(w, r, "homepage.tmpl")
	})

	// Static files.
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join("web", "static")))))
}
