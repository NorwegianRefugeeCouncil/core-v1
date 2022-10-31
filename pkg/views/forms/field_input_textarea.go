package forms

// TextAreaInputField represents a text area input field.
type TextAreaInputField struct {
	// Name is the name of the field.
	Name string `json:"name"`
	// DisplayName is the display name of the field.
	DisplayName string `json:"displayName"`
	// Required is true if the field is required.
	Required bool `json:"required"`
	// Value is the string value of the field.
	Value string `json:"value"`
	// Help is the help text of the field.
	Help string `json:"help"`
	// Errors are the errors of the field.
	Errors []string `json:"errors"`
}

// GetName implements FieldDefinition.GetName
func (f *TextAreaInputField) GetName() string {
	return f.Name
}

// GetValue implements FieldDefinition.GetValue
func (f *TextAreaInputField) GetValue() string {
	return f.Value
}

// SetValue implements FieldDefinition.SetValue
func (f *TextAreaInputField) SetValue(value string) {
	f.Value = value
}

// SetErrors implements FieldDefinition.SetErrors
func (f *TextAreaInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *TextAreaInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetKind implements FieldDefinition.GetKind
func (f *TextAreaInputField) GetKind() FieldKind {
	return FieldKindTextarea
}
