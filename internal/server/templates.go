package server

import (
	"html/template"
	"strings"
	"time"

	"github.com/nrc-no/notcore/web"
)

type templates map[string]*template.Template

func parseTemplates() (templates, error) {
	t := make(templates)
	entries, err := web.Content.ReadDir("templates")
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if entry.Name() == "layout.gohtml" {
			continue
		}
		name := entry.Name()
		tpl := template.New("")
		tpl.Funcs(map[string]any{
			"add": func(a, b int) int {
				return a + b
			},
			"sub": func(a, b int) int {
				return a - b
			},
			"age": func(dateOfBirth time.Time) int {
				return int(time.Since(dateOfBirth).Hours() / 24 / 365)
			},
			"joinStrings": func(a []string, b string) string {
				return strings.Join(a, b)
			},
			"contains": func(arr []string, str string) bool {
				for _, v := range arr {
					if v == str {
						return true
					}
				}
				return false
			},
		})
		t[name], err = tpl.ParseFS(web.Content, "templates/layout.gohtml", "templates/"+name)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
