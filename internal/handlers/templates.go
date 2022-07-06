package handlers

import (
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

// ValidationErrors is a map of field names to error messages.
type ValidationErrors map[string]string

// ViewData is a map of key/value pairs that can be used to render a view.
// We can add methods to this type to make it more useful when rendering views.
type ViewData map[string]interface{}

// ErrorsKey returns the key for the errors map in the view data.
func (v ViewData) ErrorsKey() string {
	return "Errors"
}

// RequestContextKey returns the key for the request context in the view data.
func (v ViewData) RequestContextKey() string {
	return "__request_context"
}

// GetErrors returns a map of field names to error messages.
func (v ViewData) GetErrors() ValidationErrors {
	if v == nil {
		return ValidationErrors{}
	}
	if errs, ok := v[v.ErrorsKey()]; ok {
		if errsMap, ok := errs.(ValidationErrors); ok {
			return errsMap
		}
	}
	return ValidationErrors{}
}

// Error returns the error message for the given field.
func (v ViewData) Error(field string) string {
	if v == nil {
		return ""
	}
	if errs := v.GetErrors(); errs != nil {
		if err, ok := errs[field]; ok {
			return err
		}
	}
	return ""
}

// HasError returns true if the given field has an error.
func (v ViewData) HasError(field string) bool {
	if v == nil {
		return false
	}
	if errs := v.GetErrors(); errs != nil {
		if err, ok := errs[field]; ok {
			if len(err) > 0 {
				return true
			}
		}
	}
	return false
}

// HasErrors returns true if there are any errors
func (v ViewData) HasErrors() bool {
	if v == nil {
		return false
	}
	if errs := v.GetErrors(); errs != nil {
		return len(errs) > 0
	}
	return false
}

// RequestContext returns the request context for the current request.
func (v ViewData) RequestContext() RequestContext {
	if v == nil {
		return RequestContext{}
	}
	if rc, ok := v[v.RequestContextKey()]; ok {
		if rc, ok := rc.(RequestContext); ok {
			return rc
		}
	}
	return RequestContext{}
}

// RequestContext encapsulates useful information about the current request.
type RequestContext struct {
	// Request is the http request
	Request *http.Request
	// Permissions is the set of permissions for the current user
	Permissions api.UserPermissions
	// User is the current user
	User *api.User
}

// viewParams is a map of key/value pairs that can be used to render a view.
type viewParams map[string]interface{}

// renderView renders a view with the given name and data.
func renderView(
	templates map[string]*template.Template,
	tmpl string,
	w http.ResponseWriter,
	r *http.Request,
	data viewParams,
) {

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

	rc := RequestContext{
		User:        utils.GetRequestUser(ctx),
		Request:     r,
		Permissions: utils.GetRequestUserPermissions(ctx),
	}
	vd[vd.RequestContextKey()] = rc

	if err := templates[tmpl].ExecuteTemplate(w, "base", vd); err != nil {
		l.Error("failed to execute template", zap.Error(err))
	}
}
