package middleware

import "github.com/rideee/webapp/internal/config"

type Middleware struct {
	appConfig *config.App
}

// New returns new *Middleware object.
func New(appConfig *config.App) *Middleware {
	return &Middleware{
		appConfig: appConfig,
	}
}
