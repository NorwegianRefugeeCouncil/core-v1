package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindNumberInput, (&NumberInputField{}).GetKind())
}

func TestNumberInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&NumberInputField{Name: "name"}).GetName())
}

func TestNumberInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *NumberInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &NumberInputField{},
			want:  (*int)(nil),
		}, {
			name:  "number",
			field: &NumberInputField{Value: "123"},
			want:  123,
		}, {
			name:    "invalid",
			field:   &NumberInputField{Value: "abc"},
			want:    (*int)(nil),
			wantErr: assert.Error,
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

func TestNumberInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   NumberInputField
		value   interface{}
		want    NumberInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: NumberInputField{},
			value: (*int)(nil),
			want:  NumberInputField{Value: ""},
		}, {
			name:  "number",
			field: NumberInputField{},
			value: 123,
			want:  NumberInputField{Value: "123"},
		}, {
			name:    "invalid",
			field:   NumberInputField{},
			value:   "abc",
			want:    NumberInputField{},
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

func TestNumberInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field NumberInputField
		value string
		want  NumberInputField
	}{
		{
			name:  "empty",
			field: NumberInputField{},
			value: "",
			want:  NumberInputField{Value: ""},
		}, {
			name:  "number",
			field: NumberInputField{},
			value: "123",
			want:  NumberInputField{Value: "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestNumberInputField_GetStringValue(t *testing.T) {
	field := &NumberInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
