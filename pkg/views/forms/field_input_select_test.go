package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectInputField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindSelect, (&SelectInputField{}).GetKind())
}

func TestSelectInputField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&SelectInputField{Name: "name"}).GetName())
}

func TestSelectInputField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *SelectInputField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &SelectInputField{},
			want:  "",
		}, {
			name:  "string",
			field: &SelectInputField{Value: "foo"},
			want:  "foo",
		}, {
			name:  "[multiple] empty",
			field: &SelectInputField{AllowMultiple: true, Value: ""},
			want:  []string{},
		}, {
			name:  "[multiple] string",
			field: &SelectInputField{AllowMultiple: true, Value: "foo"},
			want:  []string{"foo"},
		}, {
			name:  "[multiple] multiple strings",
			field: &SelectInputField{AllowMultiple: true, Value: "foo,bar"},
			want:  []string{"foo", "bar"},
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

func TestSelectInputField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   SelectInputField
		value   interface{}
		want    SelectInputField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: SelectInputField{},
			value: (*string)(nil),
			want:  SelectInputField{Value: ""},
		}, {
			name:  "string",
			field: SelectInputField{},
			value: "abc",
			want:  SelectInputField{Value: "abc"},
		}, {
			name:    "invalid",
			field:   SelectInputField{},
			value:   123,
			want:    SelectInputField{},
			wantErr: assert.Error,
		}, {
			name:  "[multiple] nil",
			field: SelectInputField{AllowMultiple: true},
			value: ([]string)(nil),
			want:  SelectInputField{AllowMultiple: true, Value: ""},
		}, {
			name:  "[multiple] empty",
			field: SelectInputField{AllowMultiple: true},
			value: []string{},
			want:  SelectInputField{AllowMultiple: true, Value: ""},
		}, {
			name:  "[multiple] string",
			field: SelectInputField{AllowMultiple: true},
			value: []string{"abc"},
			want:  SelectInputField{AllowMultiple: true, Value: "abc"},
		}, {
			name:  "[multiple] multiple strings",
			field: SelectInputField{AllowMultiple: true},
			value: []string{"abc", "def"},
			want:  SelectInputField{AllowMultiple: true, Value: "abc,def"},
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

func TestSelectInputField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field SelectInputField
		value string
		want  SelectInputField
	}{
		{
			name:  "empty",
			field: SelectInputField{},
			value: "",
			want:  SelectInputField{Value: ""},
		}, {
			name:  "string",
			field: SelectInputField{},
			value: "abc",
			want:  SelectInputField{Value: "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestSelectInputField_GetStringValue(t *testing.T) {
	field := &SelectInputField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
