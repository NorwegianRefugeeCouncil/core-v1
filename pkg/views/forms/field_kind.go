package forms

import "fmt"

// FieldKind represents the kind of field.
type FieldKind string

var knownFieldKinds = map[FieldKind]struct{}{}

func RegisterFieldKind(fieldKind FieldKind) {
	if _, ok := knownFieldKinds[fieldKind]; ok {
		panic(fmt.Sprintf("field kind %s is already registered", fieldKind))
	}
	knownFieldKinds[fieldKind] = struct{}{}
}

func init() {
	RegisterFieldKind(FieldKindCheckboxInput)
	RegisterFieldKind(FieldKindOptionalBooleanInput)
	RegisterFieldKind(FieldKindDateInput)
	RegisterFieldKind(FieldKindHiddenInput)
	RegisterFieldKind(FieldKindID)
	RegisterFieldKind(FieldKindNumberInput)
	RegisterFieldKind(FieldKindSelect)
	RegisterFieldKind(FieldKindTextInput)
	RegisterFieldKind(FieldKindTextarea)
	RegisterFieldKind(FieldKindUnknown)
}

const (
	// FieldKindUnknown is the default value for a field kind.
	FieldKindUnknown FieldKind = "unknown"

	// FieldKindOptionalBooleanInput is the kind of field that represents an optional boolean input.
	FieldKindOptionalBooleanInput FieldKind = "optionalBoolean"
	// FieldKindCheckboxInput is the kind of field that represents a checkbox input.
	FieldKindCheckboxInput FieldKind = "checkbox"
	// FieldKindDateInput is the kind of field that represents a date input.
	FieldKindDateInput FieldKind = "date"
	// FieldKindHiddenInput is the kind of field that represents a hidden input.
	FieldKindHiddenInput FieldKind = "hidden"
	// FieldKindID is the kind of field that represents an ID.
	FieldKindID FieldKind = "id"
	// FieldKindNumberInput is the kind of field that represents a number input.
	FieldKindNumberInput FieldKind = "number"
	// FieldKindSelect is the kind of field that represents a select input.
	FieldKindSelect FieldKind = "select"
	// FieldKindTextInput is the kind of field that represents a text input.
	FieldKindTextInput FieldKind = "text"
	// FieldKindTextarea is the kind of field that represents a textarea input.
	FieldKindTextarea FieldKind = "textarea"
)

func AllFieldKinds() []FieldKind {
	ret := make([]FieldKind, 0, len(knownFieldKinds)-1)
	for k := range knownFieldKinds {
		ret = append(ret, k)
	}
	return ret
}

func KnownFieldKinds() []FieldKind {
	ret := make([]FieldKind, 0, len(knownFieldKinds))
	for k := range knownFieldKinds {
		if k == FieldKindUnknown {
			continue
		}
		ret = append(ret, k)
	}
	return ret
}
