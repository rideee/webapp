package main

import (
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/rideee/webapp/internal/config"
)

// routes keeps all application routes.
func routes(mux *mux.Router, app *config.App) {
	views := app.TemplateEngine

	// Homepage.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view, err := views.GetTemplate(path.Join("home", "home.page.jet"))
		if err != nil {
			log.Println("Unexpected template err:", err.Error())
		}

		view.Execute(w, nil, nil)
	})

	// Static files.
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join("web", "static")))))
}
