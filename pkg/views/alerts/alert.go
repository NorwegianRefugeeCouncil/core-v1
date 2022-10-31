package alerts

import (
	"html/template"
	"strings"
)

type AlertStyle int

const (
	AlertStylePrimary AlertStyle = iota
	AlertStyleSecondary
	AlertStyleSuccess
	AlertStyleDanger
	AlertStyleWarning
	AlertStyleInfo
)

func (a AlertStyle) Class() string {
	switch a {
	case AlertStylePrimary:
		return "alert-primary"
	case AlertStyleSecondary:
		return "alert-secondary"
	case AlertStyleSuccess:
		return "alert-success"
	case AlertStyleDanger:
		return "alert-danger"
	case AlertStyleWarning:
		return "alert-warning"
	case AlertStyleInfo:
		return "alert-info"
	default:
		return ""
	}
}

// Alert represents a UI alert
type Alert struct {
	// Style is the style of the alert
	Style AlertStyle
	// Title is the title of the alert
	Title string
	// Message is the message of the alert
	Body string
	// Dismissible determines if the alert can be dismissed
	Dismissible bool
	// Icon is the icon to display in the alert
	Icon string
}

type Alerts []*Alert

func NewAlerts() Alerts {
	return []*Alert{}
}

// Render renders the alert
func (a Alert) Render() (template.HTML, error) {
	var tplBytes = &strings.Builder{}
	err := tpl.Execute(tplBytes, a)
	if err != nil {
		return "", err
	}
	return template.HTML(tplBytes.String()), nil
}

var tpl = template.Must(template.New("alert").Parse(`
<div class="alert {{.Style.Class}} {{if .Dismissible}}alert-dismissible{{end}}" role="alert">
	{{if .Title}}
		<h4 class="alert-heading">{{ if .Icon}}<i class="bi bi-{{.Icon}} me-2"></i>{{end}}{{.Title}}</h4>
	{{end}}
	{{if .Body}}
		{{ if not .Title }}
			{{ if .Icon }}<p class="mb-0"><i class="bi bi-{{.Icon}} me-2"></i>{{.Body}}</p>{{end}}
		{{else}}
			<p class="mb-0">{{.Body}}</p>
		{{end}}
	{{end}}
	{{if .Dismissible}}
		<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
	{{end}}
</div>
`))
