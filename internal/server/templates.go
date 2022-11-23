package server

import (
	"errors"
	"html/template"
	"strings"
	"time"

	"github.com/nrc-no/notcore/web"
)

type templates map[string]*template.Template

func parseTemplates(
	loginURL string,
	logoutURL string,
	refreshURL string,
	tokenRefreshInterval time.Duration,
) (templates, error) {
	t := make(templates)
	entries, err := web.Content.ReadDir("templates")
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if entry.Name() == "nav.gohtml" {
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
			"logoutURL": func() string {
				return logoutURL
			},
			"loginURL": func() string {
				return loginURL
			},
			"tokenRefreshURL": func() string {
				return refreshURL
			},
			"tokenRefreshInterval": func() time.Duration {
				return tokenRefreshInterval
			},
			"time": func() TimeFunctions {
				return TimeFunctions{}
			},
			"concat": func(strs ...string) string {
				return strings.Join(strs, "")
			},
			"attr": func(val string) template.HTMLAttr {
				return template.HTMLAttr(val)
			},
			"safe": func(val string) template.HTML {
				return template.HTML(val)
			},
			"dict": func(values ...interface{}) (map[string]interface{}, error) {
				if len(values)%2 != 0 {
					return nil, errors.New("dict must have an even number of arguments")
				}
				dict := make(map[string]interface{})
				for i := 0; i < len(values); i += 2 {
					key, ok := values[i].(string)
					if !ok {
						return nil, errors.New("dict keys must be strings")
					}
					dict[key] = values[i+1]
				}
				return dict, nil
			},
		})
		t[name], err = tpl.ParseFS(web.Content,
			"templates/nav.gohtml",
			"templates/base.gohtml",
			"templates/searchForm.gohtml",
			"templates/"+name)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
