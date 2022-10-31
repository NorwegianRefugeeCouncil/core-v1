package forms

// FieldDefinition represents a field definition for a form.
// Only one of the fields should be set.
type FieldDefinition struct {
	// IDField represents an ID field.
	IDField *IDField `json:"id,omitempty"`
	// Text represents a text input field.
	Text *TextInputField `json:"text,omitempty"`
	// Number represents a number input field.
	Number *NumberInputField `json:"number,omitempty"`
	// Date represents a date input field.
	Date *DateInputField `json:"date,omitempty"`
	// Select represents a select input field.
	Select *SelectInputField `json:"select,omitempty"`
	// Checkbox represents a checkbox input field.
	Checkbox *CheckboxInputField `json:"checkbox,omitempty"`
	// TextArea represents a textarea input field.
	TextArea *TextAreaInputField `json:"textarea,omitempty"`
	// Hidden represents a hidden input field.
	Hidden *HiddenInputField `json:"hidden,omitempty"`
}

func NewFieldDefinition(field Field) *FieldDefinition {
	field = ensurePtr(field).(Field)
	switch typ := field.(type) {
	case *FieldDefinition:
		return typ
	case *IDField:
		return &FieldDefinition{IDField: typ}
	case *TextInputField:
		return &FieldDefinition{Text: typ}
	case *NumberInputField:
		return &FieldDefinition{Number: typ}
	case *DateInputField:
		return &FieldDefinition{Date: typ}
	case *SelectInputField:
		return &FieldDefinition{Select: typ}
	case *CheckboxInputField:
		return &FieldDefinition{Checkbox: typ}
	case *TextAreaInputField:
		return &FieldDefinition{TextArea: typ}
	case *HiddenInputField:
		return &FieldDefinition{Hidden: typ}
	default:
		panic("unknown field type")
	}
}

// GetKind implements InputField.GetKind
func (f *FieldDefinition) GetKind() FieldKind {
	field := f.getField()
	if inputField, ok := field.(InputField); ok {
		return inputField.GetKind()
	}
	return FieldKindUnknown
}

// GetName implements InputField.GetName
func (f *FieldDefinition) GetName() string {
	field := f.getField()
	if inputField, ok := field.(InputField); ok {
		return inputField.GetName()
	}
	return ""
}

// SetValue implements InputField.SetValue
func (f *FieldDefinition) SetValue(value string) {
	field := f.getField()
	if inputField, ok := field.(InputField); ok {
		inputField.SetValue(value)
	}
}

// GetValue implements InputField.GetValue
func (f *FieldDefinition) GetValue() string {
	field := f.getField()
	if inputField, ok := field.(InputField); ok {
		return inputField.GetValue()
	}
	return ""
}

// SetErrors implements InputField.SetErrors
func (f *FieldDefinition) SetErrors(errors []string) {
	field := f.getField()
	if inputField, ok := field.(InputField); ok {
		inputField.SetErrors(errors)
	}
}

// HasErrors implements InputField.HasErrors
func (f *FieldDefinition) HasErrors() bool {
	field := f.getField()
	if inputField, ok := field.(InputField); ok {
		return inputField.HasErrors()
	}
	return false
}

// GetErrors implements InputField.GetErrors
func (f *FieldDefinition) GetErrors() []string {
	field := f.getField()
	if inputField, ok := field.(InputField); ok {
		return inputField.GetErrors()
	}
	return nil
}
