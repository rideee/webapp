package config

import (
	"fmt"
	"path"

	"github.com/CloudyKit/jet/v6"
)

// App holds the application config.
type App struct {
	Protocol        string
	Host            string
	Port            int
	DevelopmentMode bool
	TemplateEngine  *jet.Set
	TemplateRootDir string
}

// GetAddress returns the host and port in the format host: port (127.0.0.1:8080).
func (a *App) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

// InDevMode() returns true when development mode is active.
func (a *App) InDevelopmentMode() bool {
	return a.DevelopmentMode
}

// NewApp returns a new App object with default values.
func NewApp() *App {
	return &App{
		Protocol:        "http",
		Host:            "",
		Port:            8080,
		DevelopmentMode: false,
		TemplateRootDir: path.Join("web", "template"),
	}
}
