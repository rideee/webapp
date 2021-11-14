package config

import (
	"fmt"
)

// App holds the application config.
type App struct {
	host            string
	port            int
	developmentMode bool
}

// SetHost sets application host.
func (a *App) SetHost(host string) {
	a.host = host
}

// GetHost returns application host.
func (a *App) GetHost() string {
	return a.host
}

// SetPort sets application port.
func (a *App) SetPort(port int) {
	a.port = port
}

// GetPort returns application port.
func (a *App) GetPort() int {
	return a.port
}

// GetAddress returns the host and port in the format host: port (127.0.0.1:8080).
func (a *App) GetAddress() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}

// SetDevelopmentMode sets development mode.
func (a *App) SetDevelopmentMode(mode bool) {
	a.developmentMode = mode
}

// NewApp returns a new App object with default values.
func NewApp() *App {
	return &App{
		host:            "",
		port:            8080,
		developmentMode: false,
	}
}
