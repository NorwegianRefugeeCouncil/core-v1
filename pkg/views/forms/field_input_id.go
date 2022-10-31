package forms

// IDField represents an ID field.
type IDField struct {
	// QRCodeURL is the URL of the QR code.
	QRCodeURL string `json:"qrCodeURL"`
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

// Ensure IDField implements InputField
var _ InputField = (*IDField)(nil)

// GetName implements FieldDefinition.GetName
func (f *IDField) GetName() string {
	return f.Name
}

// GetValue implements FieldDefinition.GetValue
func (f *IDField) GetValue() string {
	return f.Value
}

// SetValue implements FieldDefinition.SetValue
func (f *IDField) SetValue(value string) {
	f.Value = value
}

// SetErrors implements FieldDefinition.SetErrors
func (f *IDField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *IDField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetKind implements FieldDefinition.GetKind
func (f *IDField) GetKind() FieldKind {
	return FieldKindID
}
