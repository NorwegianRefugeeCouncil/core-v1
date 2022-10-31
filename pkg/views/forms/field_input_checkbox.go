package forms

// CheckboxInputField represents a checkbox input field.
type CheckboxInputField struct {
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
func (f *CheckboxInputField) GetName() string {
	return f.Name
}

// GetValue implements FieldDefinition.GetValue
func (f *CheckboxInputField) GetValue() string {
	return f.Value
}

// SetValue implements FieldDefinition.SetValue
func (f *CheckboxInputField) SetValue(value string) {
	f.Value = value
}

// SetErrors implements FieldDefinition.SetErrors
func (f *CheckboxInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *CheckboxInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetKind implements FieldDefinition.GetKind
func (f *CheckboxInputField) GetKind() FieldKind {
	return FieldKindCheckboxInput
}
