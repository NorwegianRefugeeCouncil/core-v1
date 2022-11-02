package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHiddenInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindHiddenInput, (&HiddenInputField{}).GetKind())
}

func TestHiddenInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&HiddenInputField{Name: "name"}).GetName())
}

func TestHiddenInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *HiddenInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &HiddenInputField{},
			want:  "",
		}, {
			name:  "string",
			field: &HiddenInputField{Value: "foo"},
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

func TestHiddenInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   HiddenInputField
		value   interface{}
		want    HiddenInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: HiddenInputField{},
			value: (*string)(nil),
			want:  HiddenInputField{Value: ""},
		}, {
			name:  "string",
			field: HiddenInputField{},
			value: "foo",
			want:  HiddenInputField{Value: "foo"},
		}, {
			name:    "invalid",
			field:   HiddenInputField{},
			value:   123,
			want:    HiddenInputField{},
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

func TestHiddenInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field HiddenInputField
		value string
		want  HiddenInputField
	}{
		{
			name:  "empty",
			field: HiddenInputField{},
			value: "",
			want:  HiddenInputField{Value: ""},
		}, {
			name:  "true",
			field: HiddenInputField{},
			value: "foo",
			want:  HiddenInputField{Value: "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestHiddenInputField_GetStringValue(t *testing.T) {
	field := &HiddenInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
