package handlers

import (
	"html/template"
	"net/http"
)

func HandleIndex(templates map[string]*template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates["index.gohtml"].ExecuteTemplate(w, "base", nil)
	})
}
