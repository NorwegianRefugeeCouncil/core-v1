package alert

import (
	"html/template"
	"strings"

	"github.com/nrc-no/notcore/pkg/views/bootstrap"
)

type Alert struct {
	// Title is the optional title of the alert. If not set, the alert will not have a title.
	Title string
	// Content is the content of the alert. It is rendered as HTML
	Content template.HTML
	// Type is the type of the alert. Defaults to Success
	Type bootstrap.Style
	// Dismissible is whether the alert is dismissible.
	Dismissible bool
	// Icon is the optional icon to display in the alert. If not set, no icon will be displayed.
	Icon string
}

const alertTemplate = `

<div class="alert alert-{{ .Type }} {{ if .Dismissible }}alert-dismissible{{ end }}" role="alert">
{{ if .Title }}
	<div class="d-flex justify-content-between">
		<h4 class="alert-heading">
			{{ if .Icon }}<i class="bi bi-{{.Icon}}"></i>{{end}}
			{{.Title}}
		</h4>
		{{ if .Dismissible}}
				<i class="bi bi-x h4" data-bs-dismiss="alert" aria-label="Close"></i>
		{{end}}
	</div>
{{ end }}

<div>
{{ if not .Title}}
	{{ if .Icon}}
		<i class="bi bi-{{.Icon}}"></i>
	{{end}}
{{end}}
{{ .Content }}
</div class="{{if and (.Icon) (not .Title)}}d-flex align-items-center{{end}}">
</div>
`

var alertTemplateCompiled = template.Must(template.New("alert").Parse(alertTemplate))

func (a Alert) Render() template.HTML {
	var buf strings.Builder
	if err := alertTemplateCompiled.Execute(&buf, a); err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}
