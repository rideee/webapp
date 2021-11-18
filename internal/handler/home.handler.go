package handler

import (
	"log"
	"net/http"
	"path"

	"github.com/CloudyKit/jet/v6"
)

// Home handler.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	view, err := h.views.GetTemplate(path.Join("home", "home.page.jet"))
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
	}

	vars := make(jet.VarMap)
	vars.Set("pageTitle", "Home")

	view.Execute(w, vars, nil)
}
