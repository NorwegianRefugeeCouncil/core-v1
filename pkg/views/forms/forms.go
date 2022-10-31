package forms

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/nrc-no/notcore/pkg/api/validation"
)

var formTemplate *template.Template

//go:embed form.gohtml
var formHtmlTemplate string

func init() {
	t := template.New("form").Funcs(template.FuncMap{
		"isLast": func(i int, arr interface{}) bool {
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
	})
	formTemplate = template.Must(t.Parse(formHtmlTemplate))
}

type FormSection struct {
	Title       string
	Collapsible bool
	Collapsed   bool
	Fields      []*Field
}

type FormInteraction struct {
	ButtonIcon            string
	ButtonLabel           string
	ButtonStyle           string
	ButtonTitle           string
	FormAction            string
	FormMethod            string
	ModalIcon             string
	ModalIconStyle        string
	ModalContent          template.HTML
	ShowConfirmationModal bool
}

type Form struct {
	ColClasses   string
	Action       string
	Method       string
	WasValidated bool
	Sections     []*FormSection
	Interactions []*FormInteraction
}

func (f *Form) HTML() (template.HTML, error) {
	var b bytes.Buffer
	err := formTemplate.ExecuteTemplate(&b, "form", f)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}

func (f *Form) ParseURLValues(v url.Values) {
	for sectionIndex, section := range f.Sections {

		collapsedValue := v.Get("__form__section-" + strconv.Itoa(sectionIndex) + "-collapsed")
		if collapsedValue == "true" {
			section.Collapsed = true
		} else if collapsedValue == "false" {
			section.Collapsed = false
		}

		for _, f := range section.Fields {
			fieldName := f.GetName()
			if fieldName == "" {
				continue
			}
			if f.Checkbox != nil {
				if v.Get(fieldName) == "true" {
					f.Checkbox.Value = "true"
				} else {
					f.Checkbox.Value = "false"
				}
			} else if v.Has(fieldName) {
				f.SetValue(v.Get(fieldName))
			}
		}
	}
}

func (f *Form) SetErrors(errors validation.ErrorList) {
	if len(errors) > 0 {
		f.WasValidated = true
	}
	errsPerField := make(map[string][]*validation.Error)
	for _, err := range errors {
		errsPerField[err.Field] = append(errsPerField[err.Field], err)
	}
	for _, s := range f.Sections {
		for _, f := range s.Fields {
			var errStrings []string
			if errs, ok := errsPerField[f.GetName()]; ok {
				for _, err := range errs {
					s.Collapsed = false
					errStrings = append(errStrings, err.ErrorBody())
				}
			}
			f.SetErrors(errStrings)
		}
	}
}

func (f *Form) Into(i interface{}) error {

	reflectValue := reflect.ValueOf(i)
	structFieldsByName := computeStructFields(i)

	var allErrs validation.ErrorList

	for _, s := range f.Sections {
		for _, f := range s.Fields {
			name := f.GetName()
			if name == "" {
				continue
			}
			structField, ok := structFieldsByName[name]
			if !ok {
				return fmt.Errorf("form field %q not found in struct %T", name, i)
			}
			reflectFieldValue := reflectValue.Elem().FieldByName(structField.Name)
			isPtrValue := reflectFieldValue.Kind() == reflect.Ptr
			actualValue, validationErrors := parseFieldValue(f)
			allErrs = append(allErrs, validationErrors...)
			if len(validationErrors) > 0 {
				continue
			}
			if actualValue == nil {
				continue
			}
			if isPtrValue {
				ptrValue := reflect.New(reflectFieldValue.Type().Elem())
				ptrValue.Elem().Set(*actualValue)
				actualValue = &ptrValue
			}
			reflectFieldValue.Set(*actualValue)
		}
	}
	if len(allErrs) > 0 {
		return allErrs.ToAggregate()
	}
	return nil
}

func computeStructFields(i interface{}) map[string]reflect.StructField {
	reflectValue := reflect.ValueOf(i)
	fieldCount := reflectValue.Elem().NumField()
	reflectType := reflect.TypeOf(i)
	structFieldsByName := map[string]reflect.StructField{}
	for i := 0; i < fieldCount; i++ {
		field := reflectType.Elem().Field(i)
		fieldName := field.Name
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			jsonTag = strings.Split(jsonTag, ",")[0]
			fieldName = jsonTag
		}
		structFieldsByName[fieldName] = field
	}
	return structFieldsByName
}

func parseFieldValue(f *Field) (*reflect.Value, validation.ErrorList) {
	var actualValue reflect.Value
	switch {
	case f.Checkbox != nil:
		boolValue := f.Checkbox.Value == "true"
		actualValue = reflect.ValueOf(boolValue)
	case f.Date != nil:
		var dateValue time.Time
		if f.Date.Value == "" {
			return nil, nil
		}
		dateValue, err := time.Parse("2006-01-02", f.Date.Value)
		if err != nil {
			return nil, validation.ErrorList{
				validation.Invalid(
					validation.NewPath(f.Date.Name),
					f.Date.Value,
					fmt.Sprintf("invalid date: %v", err)),
			}
		}
		actualValue = reflect.ValueOf(dateValue)
	case f.Text != nil ||
		f.MultilineText != nil ||
		f.IDField != nil ||
		f.Select != nil:
		actualValue = reflect.ValueOf(f.GetValue())
	case f.Number != nil:
		intValue, err := strconv.Atoi(f.Number.Value)
		if err != nil {
			return nil, validation.ErrorList{
				validation.Invalid(
					validation.NewPath(f.Number.Name),
					f.Number.Value,
					fmt.Sprintf("unable to convert %q to int", f.Number.Value)),
			}
		}
		actualValue = reflect.ValueOf(intValue)
	default:
		return nil, validation.ErrorList{
			validation.InternalError(
				validation.NewPath(f.GetName()),
				errors.New("unknown field type")),
		}
	}
	return &actualValue, nil
}
