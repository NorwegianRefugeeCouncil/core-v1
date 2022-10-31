package forms

// Field represents a form field.
type Field interface {
	// GetKind returns the kind of the field.
	GetKind() FieldKind
}

// InputField represents an input field.
// It is a field that can be used to input data.
type InputField interface {
	Field
	// GetName returns the name of the field.
	GetName() string
	// GetValue returns the value of the field.
	GetValue() string
	// SetValue sets the value of the field.
	SetValue(value string)
	// SetErrors sets the errors of the field.
	SetErrors(errors []string)
	// HasErrors returns true if the field has errors.
	HasErrors() bool
}
