package main

import (
	"log"
	"net/http"

	"github.com/rideee/webapp/internal/config"
)

func main() {
	app := config.NewApp()
	app.SetPort(8080)

	log.Println("Server running on ", app.GetAddress())
	log.Fatal(http.ListenAndServe(app.GetAddress(), nil))
}
