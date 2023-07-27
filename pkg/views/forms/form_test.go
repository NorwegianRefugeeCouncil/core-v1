package forms

import (
	"net/url"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

type DummyStruct struct {
	TextField        string     `db:"textField"`
	TextFieldPtr     *string    `db:"textFieldPtr"`
	NumberField      int        `db:"numberField"`
	NumberFieldPtr   *int       `db:"numberFieldPtr"`
	TextAreaField    string     `db:"textAreaField"`
	TextAreaFieldPtr *string    `db:"textAreaFieldPtr"`
	CheckboxField    bool       `db:"checkboxField"`
	CheckboxFieldPtr *bool      `db:"checkboxFieldPtr"`
	DateField        time.Time  `db:"dateField"`
	DateFieldPtr     *time.Time `db:"dateFieldPtr"`
	SelectField      string     `db:"selectField"`
	SelectFieldPtr   *string    `db:"selectFieldPtr"`
}

func formWithFields(fields ...Field) *Form {
	return &Form{
		Sections: []*FormSection{
			{
				Fields: fields,
			},
		},
	}
}

func textField(name, value string) *TextInputField {
	return &TextInputField{
		Name:  name,
		Value: value,
	}
}
func numberField(name, value string) *NumberInputField {
	return &NumberInputField{
		Name:  name,
		Value: value,
	}
}
func textAreaField(name, value string) *TextAreaInputField {
	return &TextAreaInputField{
		Name:  name,
		Value: value,
	}
}
func checkboxField(name, value string) *CheckboxInputField {
	return &CheckboxInputField{
		Name:  name,
		Value: value,
	}
}
func dateField(name, value string) *DateInputField {
	return &DateInputField{
		Name:  name,
		Value: value,
	}
}
func selectField(name, value string, options []SelectInputFieldOption) *SelectInputField {
	return &SelectInputField{
		Name:    name,
		Value:   value,
		Options: options,
	}
}

func TestFormInto(t *testing.T) {

	selectFieldOptions := []SelectInputFieldOption{
		{Value: "foo", Label: "Foo"},
		{Value: "bar", Label: "Bar"},
	}

	tests := []struct {
		name    string
		form    *Form
		want    DummyStruct
		wantErr bool
	}{
		{
			name: "textField",
			form: formWithFields(textField("textField", "foo")),
			want: DummyStruct{TextField: "foo"},
		}, {
			name: "textField_pointer",
			form: formWithFields(textField("textFieldPtr", "foo")),
			want: DummyStruct{TextFieldPtr: pointers.String("foo")},
		}, {
			name: "textField_pointer_empty",
			form: formWithFields(textField("textFieldPtr", "")),
			want: DummyStruct{TextFieldPtr: pointers.String("")},
		}, {
			name: "numberField",
			form: formWithFields(numberField("numberField", "123")),
			want: DummyStruct{NumberField: 123},
		}, {
			name: "numberField_pointer",
			form: formWithFields(numberField("numberFieldPtr", "123")),
			want: DummyStruct{NumberFieldPtr: pointers.Int(123)},
		}, {
			name:    "numberField_invalid",
			form:    formWithFields(numberField("numberField", "abc")),
			wantErr: true,
		}, {
			name:    "numberField_pointer_invalid",
			form:    formWithFields(numberField("numberField", "abc")),
			wantErr: true,
		}, {
			name: "numberField_pointer_empty",
			form: formWithFields(numberField("numberFieldPtr", "")),
			want: DummyStruct{NumberFieldPtr: nil},
		}, {
			name: "textAreaField",
			form: formWithFields(textAreaField("textAreaField", "foo\nbar")),
			want: DummyStruct{TextAreaField: "foo\nbar"},
		}, {
			name: "textAreaField_pointer",
			form: formWithFields(textAreaField("textAreaFieldPtr", "buzz\nbar")),
			want: DummyStruct{TextAreaFieldPtr: pointers.String("buzz\nbar")},
		}, {
			name: "textAreaField_pointer_empty",
			form: formWithFields(textAreaField("textAreaFieldPtr", "")),
			want: DummyStruct{TextAreaFieldPtr: pointers.String("")},
		}, {
			name: "checkboxField",
			form: formWithFields(checkboxField("checkboxField", "true")),
			want: DummyStruct{CheckboxField: true},
		}, {
			name: "checkboxField_pointer",
			form: formWithFields(checkboxField("checkboxFieldPtr", "true")),
			want: DummyStruct{CheckboxFieldPtr: pointers.Bool(true)},
		}, {
			name: "checkboxField_pointer_empty",
			form: formWithFields(checkboxField("checkboxFieldPtr", "")),
			want: DummyStruct{CheckboxFieldPtr: (*bool)(nil)},
		}, {
			name: "dateField",
			form: formWithFields(dateField("dateField", "2020-01-01")),
			want: DummyStruct{DateField: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		}, {
			name: "dateField_pointer",
			form: formWithFields(dateField("dateFieldPtr", "2020-01-01")),
			want: DummyStruct{DateFieldPtr: pointers.Time(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))},
		}, {
			name: "dateField_pointer_empty",
			form: formWithFields(dateField("dateFieldPtr", "")),
			want: DummyStruct{DateFieldPtr: nil},
		}, {
			name:    "dateField_invalid",
			form:    formWithFields(dateField("dateField", "invalid")),
			wantErr: true,
		}, {
			name:    "dateField_pointer_invalid",
			form:    formWithFields(dateField("dateFieldPtr", "invalid")),
			wantErr: true,
		}, {
			name: "selectField",
			form: formWithFields(selectField("selectField", "foo", selectFieldOptions)),
			want: DummyStruct{SelectField: "foo"},
		}, {
			name: "selectField_pointer",
			form: formWithFields(selectField("selectFieldPtr", "foo", selectFieldOptions)),
			want: DummyStruct{SelectFieldPtr: pointers.String("foo")},
		}, {
			name: "selectField_pointer_empty",
			form: formWithFields(selectField("selectFieldPtr", "", selectFieldOptions)),
			want: DummyStruct{SelectFieldPtr: pointers.String("")},
		}, {
			name:    "unknownField",
			form:    formWithFields(selectField("invalid", "", selectFieldOptions)),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ds DummyStruct
			err := tt.form.Into(&ds)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				if !assert.NoError(t, err) {
					return
				}
				assert.Equal(t, tt.want, ds)
			}
		})
	}
}

func TestFormParseURLValues(t *testing.T) {
	tests := []struct {
		name   string
		form   *Form
		values url.Values
		want   *Form
	}{
		{
			name:   "empty",
			form:   &Form{},
			values: url.Values{},
			want:   &Form{},
		}, {
			name:   "single",
			form:   formWithFields(textField("foo", "bar")),
			values: url.Values{"foo": []string{"bar"}},
			want:   formWithFields(textField("foo", "bar")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.form.ParseURLValues(tt.values)
			assert.Equal(t, tt.want, tt.form)
		})
	}
}
