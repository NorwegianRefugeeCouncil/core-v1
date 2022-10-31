package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFieldDefinition(t *testing.T) {
	type test struct {
		name string
		kind FieldKind
	}
	var tests []test
	for _, fieldKind := range KnownFieldKinds() {
		tests = append(tests, test{
			name: fieldKind.String(),
			kind: fieldKind,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				ret := NewFieldDefinition(genField(tt.kind))
				assert.NotNilf(t, ret, `Error occurred while creating field definition for kind %s.
This either means that a new field kind was added, and was not handled in NewFieldDefinition, or
that the file fieldkind_string.go was not re-generated. 
To fix
- Run make generate
- Make sure that NewFieldDefinition properly handles all field kinds`,
					tt.name)
			})
		})
	}
}
