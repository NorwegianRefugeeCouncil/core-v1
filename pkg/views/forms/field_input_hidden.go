package forms

// HiddenInputField represents a hidden input field.
type HiddenInputField struct {
	// Name is the name of the field.
	Name string `json:"name"`
	// Value is the string value of the field.
	Value string `json:"value"`
}

// GetName implements FieldDefinition.GetName
func (f *HiddenInputField) GetName() string {
	return f.Name
}

// GetValue implements FieldDefinition.GetValue
func (f *HiddenInputField) GetValue() string {
	return f.Value
}

// SetValue implements FieldDefinition.SetValue
func (f *HiddenInputField) SetValue(value string) {
	f.Value = value
}

// SetErrors implements FieldDefinition.SetErrors
func (f *HiddenInputField) SetErrors(_ []string) {
	// do nothing. hidden fields cannot have errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *HiddenInputField) HasErrors() bool {
	return false
}

// GetKind implements FieldDefinition.GetKind
func (f *HiddenInputField) GetKind() FieldKind {
	return FieldKindHiddenInput
}
