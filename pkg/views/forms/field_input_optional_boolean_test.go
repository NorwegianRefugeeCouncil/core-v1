package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionalBooleanInputField_HasErrors(t *testing.T) {
	assert.True(t, (&OptionalBooleanInputField{Errors: []string{"a"}}).HasErrors())
	assert.False(t, (&OptionalBooleanInputField{Errors: []string{}}).HasErrors())
}

func TestOptionalBooleanInputField_GetErrors(t *testing.T) {
	var errs = []string{"a"}
	f := &OptionalBooleanInputField{
		Errors: errs,
	}
	assert.Equalf(t, errs, f.GetErrors(), "GetErrors()")
}

func TestOptionalBooleanInputField_SetErrors(t *testing.T) {
	field := &OptionalBooleanInputField{}
	field.SetErrors([]string{"a"})
	assert.Equal(t, []string{"a"}, field.Errors)
}

func TestOptionalBooleanInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindOptionalBooleanInput, (&OptionalBooleanInputField{}).GetKind())
}

func TestOptionalBooleanInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&OptionalBooleanInputField{Name: "name"}).GetName())
}

func TestOptionalBooleanInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *OptionalBooleanInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &OptionalBooleanInputField{},
			want:  (*bool)(nil),
		}, {
			name:  "true",
			field: &OptionalBooleanInputField{Value: "true"},
			want:  true,
		}, {
			name:  "false",
			field: &OptionalBooleanInputField{Value: "false"},
			want:  false,
		}, {
			name:    "invalid",
			field:   &OptionalBooleanInputField{Value: "invalid"},
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

func TestOptionalBooleanInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   OptionalBooleanInputField
		value   interface{}
		want    OptionalBooleanInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: OptionalBooleanInputField{},
			value: (*bool)(nil),
			want:  OptionalBooleanInputField{Value: ""},
		}, {
			name:  "true",
			field: OptionalBooleanInputField{},
			value: true,
			want:  OptionalBooleanInputField{Value: "true"},
		}, {
			name:  "false",
			field: OptionalBooleanInputField{},
			value: false,
			want:  OptionalBooleanInputField{Value: "false"},
		}, {
			name:    "invalid",
			field:   OptionalBooleanInputField{},
			value:   "invalid",
			want:    OptionalBooleanInputField{},
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

func TestOptionalBooleanInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field OptionalBooleanInputField
		value string
		want  OptionalBooleanInputField
	}{
		{
			name:  "empty",
			field: OptionalBooleanInputField{},
			value: "",
			want:  OptionalBooleanInputField{Value: ""},
		}, {
			name:  "true",
			field: OptionalBooleanInputField{},
			value: "true",
			want:  OptionalBooleanInputField{Value: "true"},
		}, {
			name:  "false",
			field: OptionalBooleanInputField{},
			value: "false",
			want:  OptionalBooleanInputField{Value: "false"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestOptionalBooleanInputField_GetStringValue(t *testing.T) {
	field := &OptionalBooleanInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
