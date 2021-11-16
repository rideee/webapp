package config

import (
	"fmt"
)

// App holds the application config.
type App struct {
	Host            string
	Port            int
	DevelopmentMode bool
}

// GetAddress returns the host and port in the format host: port (127.0.0.1:8080).
func (a *App) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

// InDevMode() returns true when development mode is active.
func (a *App) InDevMode() bool {
	return a.DevelopmentMode
}

// NewApp returns a new App object with default values.
func NewApp() *App {
	return &App{
		Host:            "",
		Port:            8080,
		DevelopmentMode: false,
	}
}
