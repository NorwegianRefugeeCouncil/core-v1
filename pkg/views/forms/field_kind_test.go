package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllFieldKinds(t *testing.T) {
	assert.Equal(t, []FieldKind{
		FieldKindUnknown,
		FieldKindID,
		FieldKindTextInput,
		FieldKindNumberInput,
		FieldKindDateInput,
		FieldKindSelect,
		FieldKindCheckboxInput,
		FieldKindTextarea,
		FieldKindHiddenInput,
	}, AllFieldKinds())
}

func TestKnownFieldKinds(t *testing.T) {
	assert.Equal(t, []FieldKind{
		FieldKindID,
		FieldKindTextInput,
		FieldKindNumberInput,
		FieldKindDateInput,
		FieldKindSelect,
		FieldKindCheckboxInput,
		FieldKindTextarea,
		FieldKindHiddenInput,
	}, KnownFieldKinds())
}
