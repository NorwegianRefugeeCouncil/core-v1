package forms

import (
	"bytes"
	"html/template"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/pkg/api/validation"
)

type FormSection struct {
	Title       string
	Collapsible bool
	Collapsed   bool
	Fields      []Field
}

type Form struct {
	Sections []*FormSection
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

		for _, field := range section.Fields {
			inputField, ok := field.(InputField)
			if !ok {
				continue
			}
			fieldName := inputField.GetName()
			if fieldName == "" {
				continue
			}
			if !v.Has(fieldName) {
				continue
			}
			urlValueForField := v.Get(fieldName)
			inputField.SetStringValue(urlValueForField)
		}
	}
}

func (f *Form) SetErrors(errors validation.ErrorList) {
	errsPerField := make(map[string][]string)
	for _, err := range errors {
		errsPerField[err.Field] = append(errsPerField[err.Field], err.ErrorBody())
	}
	for _, section := range f.Sections {
		for _, field := range section.Fields {
			inputField, ok := field.(InputField)
			if !ok {
				continue
			}
			fieldName := inputField.GetName()
			if fieldName == "" {
				continue
			}
			if errs, ok := errsPerField[fieldName]; ok {
				inputField.SetErrors(errs)
			}
		}
	}
}

func (f *Form) Into(i interface{}) error {

	reflectValue := reflect.ValueOf(i)
	structFieldsByName := computeStructFields(i)

	var allErrs validation.ErrorList

	for _, s := range f.Sections {
		for _, field := range s.Fields {
			inputField, ok := field.(InputField)
			if !ok {
				continue
			}
			name := inputField.GetName()
			if name == "" {
				continue
			}
			structField, ok := structFieldsByName[name]
			if !ok {
				allErrs = append(allErrs, validation.NotFound(nil, name))
				continue
			}
			propValue := reflectValue.Elem().FieldByName(structField.Name)
			propIsPointer := propValue.Kind() == reflect.Ptr

			v, err := inputField.GetValue()
			if err != nil {
				allErrs = append(allErrs, validation.Invalid(nil, name, err.Error()))
				continue
			}
			fieldValue := reflect.ValueOf(v)
			fieldValueIsPointer := fieldValue.Kind() == reflect.Ptr
			if propIsPointer && !fieldValueIsPointer {
				ptrValue := reflect.New(propValue.Type().Elem())
				ptrValue.Elem().Set(fieldValue)
				fieldValue = ptrValue
			} else if !propIsPointer && fieldValueIsPointer {
				if fieldValue.IsNil() {
					fieldValue = reflect.Zero(fieldValue.Type().Elem())
				} else {
					fieldValue = fieldValue.Elem()
				}

			}
			propValue.Set(fieldValue)
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
