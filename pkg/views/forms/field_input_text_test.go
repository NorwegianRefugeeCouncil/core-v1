package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindTextInput, (&TextInputField{}).GetKind())
}

func TestTextInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&TextInputField{Name: "name"}).GetName())
}

func TestTextInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *TextInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &TextInputField{},
			want:  "",
		}, {
			name:  "string",
			field: &TextInputField{Value: "foo"},
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

func TestTextInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   TextInputField
		value   interface{}
		want    TextInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: TextInputField{},
			value: (*string)(nil),
			want:  TextInputField{Value: ""},
		}, {
			name:  "string",
			field: TextInputField{},
			value: "abc",
			want:  TextInputField{Value: "abc"},
		}, {
			name:    "invalid",
			field:   TextInputField{},
			value:   123,
			want:    TextInputField{},
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

func TestTextInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field TextInputField
		value string
		want  TextInputField
	}{
		{
			name:  "empty",
			field: TextInputField{},
			value: "",
			want:  TextInputField{Value: ""},
		}, {
			name:  "string",
			field: TextInputField{},
			value: "abc",
			want:  TextInputField{Value: "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestTextInputField_GetStringValue(t *testing.T) {
	field := &TextInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
