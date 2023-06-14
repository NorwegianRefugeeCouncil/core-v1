package handlers

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/internal/validation"
	"go.uber.org/zap"
)

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
func (v ViewData) GetErrors() validation.ValidationErrors {
	if v == nil {
		return validation.ValidationErrors{}
	}
	if errs, ok := v[v.ErrorsKey()]; ok {
		if errsMap, ok := errs.(validation.ValidationErrors); ok {
			return errsMap
		}
	}
	return validation.ValidationErrors{}
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
	// Auth is the auth.Interface
	Auth auth.Interface
	// Countries is the list of all countries
	Countries []*api.Country
	// SelectedCountry is the selected country. May be nil
	SelectedCountry *api.Country
	// Session is the current user session
	Session auth.Session

	Locales locales.Interface
	// turn on dev features
	EnableBetaFeatures bool
}

func (r RequestContext) HasSelectedCountryWritePermission() bool {
	return r.Auth.HasCountryPermissionWrite(r.SelectedCountryID())
}

func (r RequestContext) HasSelectedCountryReadPermission() bool {
	return r.Auth.HasCountryPermissionRead(r.SelectedCountryID())
}

func (r RequestContext) SelectedCountryID() string {
	if r.SelectedCountry == nil {
		return ""
	}
	return r.SelectedCountry.ID
}

// viewParams is a map of key/value pairs that can be used to render a view.
type viewParams map[string]interface{}

type Renderer interface {
	RenderView(w http.ResponseWriter, r *http.Request, templateName string, data viewParams)
}

type renderer struct {
	templates          map[string]*template.Template
	enableBetaFeatures bool
}

func NewRenderer(templates map[string]*template.Template, enableBetaFeatures bool) Renderer {
	return &renderer{
		templates:          templates,
		enableBetaFeatures: enableBetaFeatures,
	}
}

func (r *renderer) RenderView(w http.ResponseWriter, req *http.Request, templateName string, data viewParams) {
	renderView(r.templates, templateName, w, req, data, r.enableBetaFeatures)
}

// renderView renders a view with the given name and data.
func renderView(
	templates map[string]*template.Template,
	tmpl string,
	w http.ResponseWriter,
	r *http.Request,
	data viewParams,
	enableBetaFeatures bool,
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

	authIntf, err := utils.GetAuthContext(ctx)
	if err != nil {
		l.Error("failed to get auth context", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	countries, err := utils.GetCountries(ctx)
	if err != nil {
		l.Error("failed to get countries", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	selectedCountryID, err := utils.GetSelectedCountryID(ctx)
	if err != nil {
		l.Error("failed to get selected country ID", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	session, ok := utils.GetSession(ctx)
	if !ok {
		l.Error("failed to get session")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var selectedCountry *api.Country
	if len(selectedCountryID) != 0 {
		for _, c := range countries {
			if c.ID == selectedCountryID {
				selectedCountry = c
				break
			}
		}
	}

	localesInterface := locales.GetLocales()

	rc := RequestContext{
		Request:            r,
		Auth:               authIntf,
		Countries:          countries,
		SelectedCountry:    selectedCountry,
		Session:            session,
		Locales:            localesInterface,
		EnableBetaFeatures: enableBetaFeatures,
	}
	vd[vd.RequestContextKey()] = rc
	vd["DeduplicationTypes"] = deduplication.DeduplicationTypes

	if err := templates[tmpl].ExecuteTemplate(w, "base", vd); err != nil {
		l.Error("failed to execute template", zap.Error(err))
	}
}
