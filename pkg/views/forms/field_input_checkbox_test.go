package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckboxInputField_HasErrors(t *testing.T) {
	assert.True(t, (&CheckboxInputField{Errors: []string{"a"}}).HasErrors())
	assert.False(t, (&CheckboxInputField{Errors: []string{}}).HasErrors())
}

func TestCheckboxInputField_GetErrors(t *testing.T) {
	var errs = []string{"a"}
	f := &CheckboxInputField{
		Errors: errs,
	}
	assert.Equalf(t, errs, f.GetErrors(), "GetErrors()")
}

func TestCheckboxInputField_SetErrors(t *testing.T) {
	field := &CheckboxInputField{}
	field.SetErrors([]string{"a"})
	assert.Equal(t, []string{"a"}, field.Errors)
}

func TestCheckboxInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindCheckboxInput, (&CheckboxInputField{}).GetKind())
}

func TestCheckboxInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&CheckboxInputField{Name: "name"}).GetName())
}

func TestCheckboxInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *CheckboxInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &CheckboxInputField{},
			want:  (*bool)(nil),
		}, {
			name:  "true",
			field: &CheckboxInputField{Value: "true"},
			want:  true,
		}, {
			name:  "false",
			field: &CheckboxInputField{Value: "false"},
			want:  false,
		}, {
			name:    "invalid",
			field:   &CheckboxInputField{Value: "invalid"},
			want:    (*bool)(nil),
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

func TestCheckboxInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   CheckboxInputField
		value   interface{}
		want    CheckboxInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: CheckboxInputField{},
			value: (*bool)(nil),
			want:  CheckboxInputField{Value: ""},
		}, {
			name:  "true",
			field: CheckboxInputField{},
			value: true,
			want:  CheckboxInputField{Value: "true"},
		}, {
			name:  "false",
			field: CheckboxInputField{},
			value: false,
			want:  CheckboxInputField{Value: "false"},
		}, {
			name:    "invalid",
			field:   CheckboxInputField{},
			value:   "invalid",
			want:    CheckboxInputField{},
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

func TestCheckboxInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field CheckboxInputField
		value string
		want  CheckboxInputField
	}{
		{
			name:  "empty",
			field: CheckboxInputField{},
			value: "",
			want:  CheckboxInputField{Value: ""},
		}, {
			name:  "true",
			field: CheckboxInputField{},
			value: "true",
			want:  CheckboxInputField{Value: "true"},
		}, {
			name:  "false",
			field: CheckboxInputField{},
			value: "false",
			want:  CheckboxInputField{Value: "false"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestCheckboxInputField_GetStringValue(t *testing.T) {
	field := &CheckboxInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
