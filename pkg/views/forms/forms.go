package forms

import (
	"bytes"
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

type FormSection struct {
	Title       string
	Collapsible bool
	Collapsed   bool
	Fields      *FieldDefinitions
}

type Form struct {
	Action   string
	Method   string
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

		section.Fields.Each(func(fieldIndex int, field Field) {
			inputField, ok := field.(InputField)
			if !ok {
				return
			}
			fieldName := inputField.GetName()
			if fieldName == "" {
				return
			}
			urlValueForField := v.Get(fieldName)
			if inputField.GetKind() == FieldKindCheckboxInput {
				if urlValueForField == "true" {
					inputField.SetValue("true")
				} else {
					inputField.SetValue("false")
				}
			} else {
				inputField.SetValue(urlValueForField)
			}
		})
	}
}

func (f *Form) SetErrors(errors validation.ErrorList) {
	errsPerField := make(map[string][]string)
	for _, err := range errors {
		errsPerField[err.Field] = append(errsPerField[err.Field], err.ErrorBody())
	}
	for _, s := range f.Sections {
		s.Fields.Each(func(fieldIndex int, field Field) {
			inputField, ok := field.(InputField)
			if !ok {
				return
			}
			fieldName := inputField.GetName()
			if fieldName == "" {
				return
			}
			inputField.SetErrors(errsPerField[fieldName])
		})
	}
}

func (f *Form) Into(i interface{}) error {

	reflectValue := reflect.ValueOf(i)
	structFieldsByName := computeStructFields(i)

	var allErrs validation.ErrorList

	for _, s := range f.Sections {
		s.Fields.Each(func(fieldIndex int, field Field) {
			inputField, ok := field.(InputField)
			if !ok {
				return
			}
			name := inputField.GetName()
			if name == "" {
				return
			}
			structField, ok := structFieldsByName[name]
			if !ok {
				allErrs = append(allErrs, validation.NotFound(nil, name))
				return
			}
			reflectFieldValue := reflectValue.Elem().FieldByName(structField.Name)
			isPtrValue := reflectFieldValue.Kind() == reflect.Ptr
			actualValue, validationErrors := parseFieldValue(inputField)
			allErrs = append(allErrs, validationErrors...)
			if len(validationErrors) > 0 {
				return
			}
			if actualValue == nil {
				return
			}
			if isPtrValue {
				ptrValue := reflect.New(reflectFieldValue.Type().Elem())
				ptrValue.Elem().Set(*actualValue)
				actualValue = &ptrValue
			}
			reflectFieldValue.Set(*actualValue)
		})
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

func parseFieldValue(inputField InputField) (*reflect.Value, validation.ErrorList) {
	var actualValue reflect.Value
	fieldValue := inputField.GetValue()
	fieldName := inputField.GetName()

	switch inputField.GetKind() {
	case FieldKindCheckboxInput:
		boolValue := fieldValue == "true"
		actualValue = reflect.ValueOf(boolValue)
	case FieldKindDateInput:
		var dateValue time.Time
		if fieldValue == "" {
			return nil, nil
		}
		dateValue, err := time.Parse("2006-01-02", fieldValue)
		if err != nil {
			return nil, validation.ErrorList{
				validation.Invalid(
					validation.NewPath(fieldName),
					fieldValue,
					fmt.Sprintf("invalid date: %v", err)),
			}
		}
		actualValue = reflect.ValueOf(dateValue)
	case FieldKindTextInput | FieldKindTextarea | FieldKindID | FieldKindSelect:
		actualValue = reflect.ValueOf(fieldValue)
	case FieldKindNumberInput:
		intValue, err := strconv.Atoi(fieldValue)
		if err != nil {
			return nil, validation.ErrorList{
				validation.Invalid(
					validation.NewPath(fieldName),
					fieldValue,
					fmt.Sprintf("unable to convert %q to int", fieldValue)),
			}
		}
		actualValue = reflect.ValueOf(intValue)
	default:
		return nil, validation.ErrorList{
			validation.InternalError(
				validation.NewPath(fieldName),
				errors.New("unknown field type")),
		}
	}

	return &actualValue, nil
}
