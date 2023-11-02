package forms

import (
	_ "embed"
	"errors"
	"github.com/nrc-no/notcore/internal/locales"
	"html/template"
	"reflect"
)

var formTemplate *template.Template

//go:embed form.gohtml
var formHtmlTemplate string

func init() {
	t := template.New("form").Funcs(template.FuncMap{
		"isLast": func(i int, arr interface{}) bool {
			// ensure i is not a pointer to a slice
			if reflect.TypeOf(arr).Kind() == reflect.Ptr {
				arr = reflect.ValueOf(arr).Elem().Interface()
			}
			return i == reflect.ValueOf(arr).Len()-1
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
		"translate": func(id string, args ...interface{}) string {
			return locales.GetLocales().Translate(id, args...)
		},
	})
	formTemplate = template.Must(t.Parse(formHtmlTemplate))
}
