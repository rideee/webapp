package main

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/rideee/webapp/internal/config"
	"github.com/rideee/webapp/internal/handler"
)

// routes keeps all application routes.
func routes(app *config.App, mux *mux.Router, handler *handler.Handler) {

	// Homepage route.
	mux.HandleFunc("/", handler.Home)

	if app.InDevelopmentMode() {
		// CSRF Token Test routes.
		mux.HandleFunc("/test/CSRFToken", handler.CSRFTokenTest)
	}

	// Static files.
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join("web", "static")))))
}
