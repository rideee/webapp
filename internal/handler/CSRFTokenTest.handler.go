package handler

import (
	"log"
	"net/http"
	"path"

	"github.com/CloudyKit/jet/v6"
)

// CSRFTokenTest handler.
func (h *Handler) CSRFTokenTest(w http.ResponseWriter, r *http.Request) {

	view, err := h.views.GetTemplate(path.Join("CSRFTokenTest", "CSRFTokenTest.page.jet"))
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
	}

	vars := make(jet.VarMap)
	vars.Set("pageTitle", "CSRF Token Test Page")
	vars.Set("CSRFToken", h.appConfig.CSRFToken)

	if r.Method == "POST" {
		vars.Set("HiddenMessage", r.FormValue("message"))
	}

	view.Execute(w, vars, nil)

}
