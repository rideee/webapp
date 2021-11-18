package handler

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/rideee/webapp/internal/config"
)

// Handler struct stores application config and make it reachable for all handler methods related to object.
type Handler struct {
	appConfig *config.App
	views     *jet.Set
}

// New return new *Handler object.
func New(appConfig *config.App) *Handler {
	return &Handler{
		appConfig: appConfig,
		views:     appConfig.TemplateEngine,
	}
}
