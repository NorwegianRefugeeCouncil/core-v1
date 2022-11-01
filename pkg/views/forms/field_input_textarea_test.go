package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextAreaInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindTextarea, (&TextAreaInputField{}).GetKind())
}

func TestTextAreaInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&TextAreaInputField{Name: "name"}).GetName())
}

func TestTextAreaInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *TextAreaInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &TextAreaInputField{},
			want:  "",
		}, {
			name:  "string",
			field: &TextAreaInputField{Value: "foo"},
			want:  "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.GetValue()
			if tt.wantErr == nil {
				tt.wantErr = assert.NoError
			}
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTextAreaInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   TextAreaInputField
		value   interface{}
		want    TextAreaInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: TextAreaInputField{},
			value: (*string)(nil),
			want:  TextAreaInputField{Value: ""},
		}, {
			name:  "string",
			field: TextAreaInputField{},
			value: "abc",
			want:  TextAreaInputField{Value: "abc"},
		}, {
			name:    "invalid",
			field:   TextAreaInputField{},
			value:   123,
			want:    TextAreaInputField{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				tt.wantErr = assert.NoError
			}
			err := tt.field.SetValue(tt.value)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestTextAreaInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field TextAreaInputField
		value string
		want  TextAreaInputField
	}{
		{
			name:  "empty",
			field: TextAreaInputField{},
			value: "",
			want:  TextAreaInputField{Value: ""},
		}, {
			name:  "string",
			field: TextAreaInputField{},
			value: "abc",
			want:  TextAreaInputField{Value: "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestTextAreaInputField_GetStringValue(t *testing.T) {
	field := &TextAreaInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
