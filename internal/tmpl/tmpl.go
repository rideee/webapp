package tmpl

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/rideee/webapp/internal/config"
)

// tmpl represents template structure.
type Tmpl struct {
	AppConfig    *config.App
	TemplateData *struct {
		Data      map[string]interface{}
		CSRFToken string
		Flash     string
		Warning   string
		Error     string
	}
}

// New creates new tmpl object.
func New(cfg *config.App) *Tmpl {
	return &Tmpl{
		AppConfig: cfg,
	}
}

// Render renders a template.
func (t *Tmpl) Render(w http.ResponseWriter, r *http.Request, tmpl string) {
	parsedTemplate, err := template.ParseFiles(path.Join("web", "template", tmpl))
	if err != nil {
		log.Printf("Render: %v\n", err)
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Render->Execute: %v\n", err)
	}

}
