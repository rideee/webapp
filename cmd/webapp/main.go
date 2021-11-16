package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rideee/webapp/internal/config"
	"github.com/rideee/webapp/internal/tmpl"
)

func main() {
	// Parse flags.
	host := flag.String("host", "", "Set application host.")
	port := flag.Int("port", 8080, "Set application port.")
	devMode := flag.Bool("devMode", false, "Run application in development mode.")
	flag.Parse()

	// Initialize new application config.
	app := config.NewApp()
	app.Host = *host
	app.Port = *port
	app.DevelopmentMode = *devMode

	// TODO: For testing only...
	t := tmpl.New(app)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Render(w, r, "homepage.tmpl")
	})

	// Serve app.
	if app.InDevMode() {
		log.Println("Application is running in development mode.")
	}
	log.Println("Server is running on", app.Address())
	log.Fatal(http.ListenAndServe(app.Address(), nil))
}
