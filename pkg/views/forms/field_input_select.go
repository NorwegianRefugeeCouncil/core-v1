package forms

// SelectInputField represents a select input field.
type SelectInputField struct {
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
	// Options are the options of the field.
	Errors []string `json:"errors"`
	// Options are the options of the field.
	Options []SelectInputFieldOption `json:"options"`
}

// GetName implements FieldDefinition.GetName
func (f *SelectInputField) GetName() string {
	return f.Name
}

// GetValue implements FieldDefinition.GetValue
func (f *SelectInputField) GetValue() string {
	return f.Value
}

// SetValue implements FieldDefinition.SetValue
func (f *SelectInputField) SetValue(value string) {
	f.Value = value
}

// SetErrors implements FieldDefinition.SetErrors
func (f *SelectInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *SelectInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetKind implements FieldDefinition.GetKind
func (f *SelectInputField) GetKind() FieldKind {
	return FieldKindSelect
}

// SelectInputFieldOption represents an option of a select input field.
type SelectInputFieldOption struct {
	// Value is the value of the option.
	Value string
	// Label is the label of the option.
	Label string
}
