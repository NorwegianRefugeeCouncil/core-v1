package handlers

import (
	"html/template"
	"net/http"
)

func HandleHome(templates map[string]*template.Template) http.Handler {
	const templateName = "individuals.gohtml"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderView(templates, templateName, w, r, viewParams{})
	})
}
