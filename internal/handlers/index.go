package handlers

import (
	"github.com/nrc-no/notcore/internal/clients"
	"html/template"
	"net/http"
)

func HandleIndex(templates map[string]*template.Template, client zanzibar.Client) http.Handler {
	const templateName = "index.gohtml"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderView(templates, templateName, w, r, viewParams{})
	})
}
