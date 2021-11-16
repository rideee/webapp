package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rideee/webapp/internal/config"
	"github.com/rideee/webapp/internal/tmpl"
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

	// TODO: For testing only...
	// Initialize new tmpl object.
	tpl := tmpl.New(app)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Render(w, r, "homepage.tmpl")
	})

	// Serve app.
	if app.InDevMode() {
		log.Println("Application is running in development mode.")
		log.Println("Server is running on", app.Address())
	}

	// When flag open is set, open app in web browser.
	if *openBrowser {
		url := app.Address()
		if app.Host == "" {
			url = "localhost" + url
		}
		url = app.Protocol + "://" + url
		browser.OpenURL(url)
	}

	log.Fatal(http.ListenAndServe(app.Address(), nil))
}
