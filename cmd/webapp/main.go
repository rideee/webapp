package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rideee/webapp/internal/config"
)

func main() {
	// Parse flags.
	host := flag.String("host", "", "Set application host.")
	port := flag.Int("port", 8080, "Set application port.")
	devMode := flag.Bool("devMode", false, "Run application in development mode.")
	flag.Parse()

	// Initialize new application config.
	app := config.NewApp()
	app.SetHost(*host)
	app.SetPort(*port)
	app.SetDevelopmentMode(*devMode)

	// Serve app.
	if app.InDevMode() {
		log.Println("Application is running in development mode.")
	}
	log.Println("Server is running on", app.GetAddress())
	log.Fatal(http.ListenAndServe(app.GetAddress(), nil))
}
