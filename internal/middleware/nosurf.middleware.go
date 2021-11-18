package middleware

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf middleware sets (config.App).CSRFToken.
func (m *Middleware) NoSurf(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.appConfig.CSRFToken = nosurf.Token(r)

		if m.appConfig.InDevelopmentMode() {
			log.Printf("CSRFToken: %s", m.appConfig.CSRFToken)
		}

		next.ServeHTTP(w, r)
	})
}
