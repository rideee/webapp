package middleware

import (
	"log"
	"net/http"
)

// Logger logs additional information when application running in development mode.
func (m *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.appConfig.InDevelopmentMode() {
			log.Printf("New handled %s '%s' request from %s", r.Method, r.RequestURI, r.RemoteAddr)
		}
		next.ServeHTTP(w, r)
	})
}
