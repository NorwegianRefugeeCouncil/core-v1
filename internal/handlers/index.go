package handlers

import (
	"html/template"
	"net/http"
)

func HandleIndex(templates map[string]*template.Template) http.Handler {
	const templateName = "index.gohtml"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderView(templates, templateName, w, r, viewParams{})
	})
}
