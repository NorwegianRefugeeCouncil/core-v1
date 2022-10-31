package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFieldDefinitions_Add checks that all field kinds are handled
// by the FieldDefinitions.Add method.
func TestFieldDefinitions_Add(t *testing.T) {
	for _, fieldKind := range KnownFieldKinds() {
		fieldForKind := genField(fieldKind)
		t.Run(fieldKind.String(), func(t *testing.T) {
			fieldDefs := NewFieldDefinitions()
			assert.NotPanicsf(t, func() {
				fieldDefs.Add(fieldForKind)
			}, `Error occurred while calling FieldDefinitions.Add(field with kind=%s).
This either means that a new field kind was added, and was not handled in FieldDefinitions.Add, or
that the file fieldkind_string.go was not re-generated. 
To fix
- Run make generate
- Make sure that FieldDefinitions.Add properly handles all field kinds`, fieldKind.String())
		})
	}
}

func TestFieldDefinition_getField(t *testing.T) {
	for _, fieldKind := range KnownFieldKinds() {
		t.Run(fieldKind.String(), func(t *testing.T) {
			want := genField(fieldKind)
			assert.NotPanicsf(t, func() {
				fieldDef := NewFieldDefinition(want)
				assert.Equal(t, want, fieldDef.getField())
			}, `Error occurred while calling FieldDefinition.getField() with field kind "%s".
This either means that a new field kind was added, and was not handled in FieldDefinition.getField(), or
that the file fieldkind_string.go was not re-generated. 
To fix
- Run make generate
- Make sure that FieldDefinition.getField() properly handles all field kinds`, fieldKind.String())

		})
	}
}

func TestGenFields(t *testing.T) {
	for _, fieldKind := range KnownFieldKinds() {
		t.Run(fieldKind.String(), func(t *testing.T) {
			assert.NotPanicsf(t, func() {
				genField(fieldKind)
			}, `Error occurred while calling genField with field kind "%s".
This either means that a new field kind was added, and was not handled in genField, or
that the file fieldkind_string.go was not re-generated. 
To fix
- Run make generate
- Make sure that genField properly handles all field kinds`, fieldKind.String())

		})
	}
}

func genField(fieldKind FieldKind) Field {
	var field Field
	switch fieldKind {
	case FieldKindUnknown:
		panic("cannot generate field for unknown field kind")
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
