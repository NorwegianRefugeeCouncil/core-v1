package forms

//go:generate stringer -type=FieldKind

// FieldKind represents the kind of field.
type FieldKind uint8

const (
	// FieldKindUnknown is the default value for a field kind.
	FieldKindUnknown FieldKind = iota
	// FieldKindID is the kind of field that represents an ID.
	FieldKindID
	// FieldKindTextInput is the kind of field that represents a text input.
	FieldKindTextInput
	// FieldKindNumberInput is the kind of field that represents a number input.
	FieldKindNumberInput
	// FieldKindDateInput is the kind of field that represents a date input.
	FieldKindDateInput
	// FieldKindSelect is the kind of field that represents a select input.
	FieldKindSelect
	// FieldKindCheckboxInput is the kind of field that represents a checkbox input.
	FieldKindCheckboxInput
	// FieldKindTextarea is the kind of field that represents a textarea input.
	FieldKindTextarea
	// FieldKindHiddenInput is the kind of field that represents a hidden input.
	FieldKindHiddenInput
)

func AllFieldKinds() []FieldKind {
	numFieldKinds := len(_FieldKind_index) - 1
	fieldKinds := make([]FieldKind, numFieldKinds)
	for i := 0; i < numFieldKinds; i++ {
		fieldKinds[i] = FieldKind(i)
	}
	return fieldKinds
}

func KnownFieldKinds() []FieldKind {
	numFieldKinds := len(_FieldKind_index) - 2
	fieldKinds := make([]FieldKind, numFieldKinds)
	for i := 1; i <= numFieldKinds; i++ {
		fieldKinds[i-1] = FieldKind(i)
	}
	return fieldKinds
}
