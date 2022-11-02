package forms

// DateInputField represents a date input field.
type DateInputField struct {
	// Name is the name of the field.
	Name string
	// DisplayName is the display name of the field.
	DisplayName string
	// Required is true if the field is required.
	Required bool
	// Value is the string value of the field.
	Value string
	// Help is the help text of the field.
	Help string
	// Errors are the errors of the field.
	Errors []string
	// Codec is the codec of the field.
	Codec Codec
}

// Ensure DateInputField implements InputField
var _ InputField = (*DateInputField)(nil)

// GetName implements FieldDefinition.GetName
func (f *DateInputField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *DateInputField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *DateInputField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *DateInputField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *DateInputField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *DateInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *DateInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetErrors implements FieldDefinition.GetErrors
func (f *DateInputField) GetErrors() []string {
	copied := make([]string, len(f.Errors))
	copy(copied, f.Errors)
	return copied
}

// GetKind implements FieldDefinition.GetKind
func (f *DateInputField) GetKind() FieldKind {
	return FieldKindDateInput
}

func (f *DateInputField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &TimeCodec{
			format: "2006-01-02",
		}
	}
	return codec
}
