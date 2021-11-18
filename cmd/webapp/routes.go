package main

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/rideee/webapp/internal/config"
	"github.com/rideee/webapp/internal/handler"
)

// routes keeps all application routes.
func routes(mux *mux.Router, app *config.App, handler *handler.Handler) {

	// Homepage route.
	mux.HandleFunc("/", handler.Home)

	// Static files.
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join("web", "static")))))
}
