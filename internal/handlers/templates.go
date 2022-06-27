package handlers

import (
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

type ViewData map[string]interface{}

func (v ViewData) GetErrors() ValidationErrors {
	if errs, ok := v["Errors"]; ok {
		if errsMap, ok := errs.(ValidationErrors); ok {
			return errsMap
		}
	}
	return nil
}

func (v ViewData) Error(field string) string {
	if errs := v.GetErrors(); errs != nil {
		if err, ok := errs[field]; ok {
			return err
		}
	}
	return ""
}

func (v ViewData) HasError(field string) bool {
	if errs := v.GetErrors(); errs != nil {
		if err, ok := errs[field]; ok {
			if len(err) > 0 {
				return true
			}
		}
	}
	return false
}

func (v ViewData) HasErrors() bool {
	if errs := v.GetErrors(); errs != nil {
		return len(errs) > 0
	}
	return false
}

func renderView(
	templates map[string]*template.Template,
	tmpl string,
	w http.ResponseWriter,
	r *http.Request,
	data map[string]interface{}) {

	ctx := r.Context()
	l := logging.NewLogger(ctx)
	l.Debug("rendering view", zap.String("template", tmpl))

	if data == nil {
		data = make(map[string]interface{})
	}

	vd := ViewData{}
	for k, v := range data {
		vd[k] = v
	}

	if err := templates[tmpl].ExecuteTemplate(w, "base", vd); err != nil {
		l.Error("failed to execute template", zap.Error(err))
	}
}
