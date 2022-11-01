package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDField_GetKind(t *testing.T) {
	assert.Equal(t, FieldKindID, (&IDField{}).GetKind())
}

func TestIDField_GetName(t *testing.T) {
	assert.Equal(t, "name", (&IDField{Name: "name"}).GetName())
}

func TestIDField_GetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   *IDField
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			field: &IDField{},
			want:  "",
		}, {
			name:  "string",
			field: &IDField{Value: "foo"},
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

func TestIDField_SetValue(t *testing.T) {
	tests := []struct {
		name    string
		field   IDField
		value   interface{}
		want    IDField
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:  "nil",
			field: IDField{},
			value: (*string)(nil),
			want:  IDField{Value: ""},
		}, {
			name:  "string",
			field: IDField{},
			value: "foo",
			want:  IDField{Value: "foo"},
		}, {
			name:    "invalid",
			field:   IDField{},
			value:   123,
			want:    IDField{},
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

func TestIDField_SetStringValue(t *testing.T) {
	tests := []struct {
		name  string
		field IDField
		value string
		want  IDField
	}{
		{
			name:  "empty",
			field: IDField{},
			value: "",
			want:  IDField{Value: ""},
		}, {
			name:  "string",
			field: IDField{},
			value: "foo",
			want:  IDField{Value: "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.SetStringValue(tt.value)
			assert.Equal(t, tt.want, tt.field)
		})
	}
}

func TestIDField_GetStringValue(t *testing.T) {
	field := &IDField{Value: "true"}
	assert.Equal(t, "true", field.GetStringValue())
}
