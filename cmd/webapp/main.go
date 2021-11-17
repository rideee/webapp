package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/mux"
	"github.com/rideee/webapp/internal/config"
	"github.com/rideee/webapp/pkg/browser"
)

func main() {
	// Parse flags.
	host := flag.String("host", "", "Set application host.")
	port := flag.Int("port", 8080, "Set application port.")
	devMode := flag.Bool("dev", false, "Run application in development mode.")
	openBrowser := flag.Bool("open", false, "Open application in default web browser when starting the server.")
	flag.Parse()

	// Initialize new application config.
	app := config.NewApp()
	app.Host = *host
	app.Port = *port
	app.DevelopmentMode = *devMode

	if app.InDevelopmentMode() {
		log.Println("Application is running in development mode.")
		log.Println("Server is running on", app.Address())

		// Initialize Jet template engine.
		app.TemplateEngine = jet.NewSet(
			jet.NewOSFileSystemLoader(app.TemplateRootDir),
			jet.InDevelopmentMode(),
		)
	} else {
		app.TemplateEngine = jet.NewSet(
			jet.NewOSFileSystemLoader(app.TemplateRootDir),
		)
	}

	// When flag 'open' is set, open app in web browser.
	if *openBrowser {
		url := app.Address()
		if app.Host == "" {
			url = "localhost" + url
		}
		url = app.Protocol + "://" + url
		browser.OpenURL(url)
	}

	// Initialize gorilla/mux router.
	router := mux.NewRouter()
	routes(router, app)

	// Serv the application.
	srv := &http.Server{
		Handler: router,
		Addr:    app.Address(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
