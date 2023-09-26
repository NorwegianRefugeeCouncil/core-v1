package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllFieldKinds(t *testing.T) {
	assert.ElementsMatch(t, []FieldKind{
		FieldKindUnknown,
		FieldKindID,
		FieldKindTextInput,
		FieldKindNumberInput,
		FieldKindDateInput,
		FieldKindSelect,
		FieldKindCheckboxInput,
		FieldKindOptionalBooleanInput,
		FieldKindTextarea,
		FieldKindHiddenInput,
	}, AllFieldKinds())
}

func TestKnownFieldKinds(t *testing.T) {
	assert.ElementsMatch(t, []FieldKind{
		FieldKindID,
		FieldKindTextInput,
		FieldKindNumberInput,
		FieldKindDateInput,
		FieldKindSelect,
		FieldKindCheckboxInput,
		FieldKindOptionalBooleanInput,
		FieldKindTextarea,
		FieldKindHiddenInput,
	}, KnownFieldKinds())
}
