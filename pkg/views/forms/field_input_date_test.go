package forms

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateInputField_HasErrors(t *testing.T) {
	assert.True(t, (&DateInputField{Errors: []string{"a"}}).HasErrors())
	assert.False(t, (&DateInputField{Errors: []string{}}).HasErrors())
}

func TestDateInputField_GetErrors(t *testing.T) {
	var errs = []string{"a"}
	f := &DateInputField{
		Errors: errs,
	}
	assert.Equal(t, errs, f.GetErrors())
}

func TestDateInputField_SetErrors(t *testing.T) {
	field := &DateInputField{}
	field.SetErrors([]string{"a"})
	assert.Equal(t, []string{"a"}, field.Errors)
}

func TestDateInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindDateInput, (&DateInputField{}).GetKind())
}

func TestDateInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&DateInputField{Name: "name"}).GetName())
}

func TestDateInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *DateInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &DateInputField{},
			want:  (*time.Time)(nil),
		}, {
			name:  "date",
			field: &DateInputField{Value: "2020-01-01"},
			want:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}, {
			name:    "invalid",
			field:   &DateInputField{Value: "invalid"},
			want:    (*time.Time)(nil),
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

func TestDateInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   DateInputField
		value   interface{}
		want    DateInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: DateInputField{},
			value: (*time.Time)(nil),
			want:  DateInputField{Value: ""},
		}, {
			name:  "date",
			field: DateInputField{},
			value: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  DateInputField{Value: "2020-01-01"},
		}, {
			name:    "invalid",
			field:   DateInputField{},
			value:   "invalid",
			want:    DateInputField{},
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

func TestDateInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field DateInputField
		value string
		want  DateInputField
	}{
		{
			name:  "empty",
			field: DateInputField{},
			value: "",
			want:  DateInputField{Value: ""},
		}, {
			name:  "true",
			field: DateInputField{},
			value: "2020-01-01",
			want:  DateInputField{Value: "2020-01-01"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestDateInputField_GetStringValue(t *testing.T) {
	field := &DateInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
