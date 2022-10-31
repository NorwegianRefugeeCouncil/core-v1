package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFieldDefinitions_Add checks that all field kinds are handled
// by the FieldDefinitions.Add method.
func TestFieldDefinitions_Add(t *testing.T) {
	type test struct {
		name  string
		f     FieldDefinitions
		field Field
	}
	var tests []test
	for _, fieldKind := range KnownFieldKinds() {
		fieldForKind := genField(fieldKind)
		tests = append(tests, test{
			name:  fieldKind.String(),
			field: fieldForKind,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanicsf(t, func() {
				tt.f.Add(tt.field)
			}, `Error occurred while calling FieldDefinitions.Add(field with kind=%s).
This either means that a new field kind was added, and was not handled in FieldDefinitions.Add, or
that the file fieldkind_string.go was not re-generated. 
To fix
- Run make generate
- Make sure that FieldDefinitions.Add properly handles all field kinds`, tt.name)
		})
	}
}

func TestFieldDefinition_getField(t *testing.T) {
	type test struct {
		name  string
		field Field
		kind  FieldKind
	}
	var tests []test
	for _, fieldKind := range KnownFieldKinds() {
		fieldForKind := genField(fieldKind)
		tests = append(tests, test{
			name:  fieldKind.String(),
			field: fieldForKind,
			kind:  fieldKind,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := genField(tt.kind)
			assert.NotPanicsf(t, func() {
				fieldDef := NewFieldDefinition(tt.field)
				assert.Equal(t, want, fieldDef.getField())
			}, `Error occurred while calling FieldDefinition.getField() with field kind "%s".
This either means that a new field kind was added, and was not handled in FieldDefinition.getField(), or
that the file fieldkind_string.go was not re-generated. 
To fix
- Run make generate
- Make sure that FieldDefinition.getField() properly handles all field kinds`, tt.name)

		})
	}
}

func genField(fieldKind FieldKind) Field {
	var field Field
	switch fieldKind {
	case FieldKindUnknown:
		// Skip this field kind.
		return nil
	case FieldKindID:
		field = &IDField{}
	case FieldKindTextInput:
		field = &TextInputField{}
	case FieldKindNumberInput:
		field = &NumberInputField{}
	case FieldKindDateInput:
		field = &DateInputField{}
	case FieldKindSelect:
		field = &SelectInputField{}
	case FieldKindCheckboxInput:
		field = &CheckboxInputField{}
	case FieldKindTextarea:
		field = &TextAreaInputField{}
	case FieldKindHiddenInput:
		field = &HiddenInputField{}
	}
	return field
}
