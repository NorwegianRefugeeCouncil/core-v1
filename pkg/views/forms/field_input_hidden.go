package forms

// HiddenInputField represents a hidden input field.
type HiddenInputField struct {
	// Name is the name of the field.
	Name string
	// Value is the string value of the field.
	Value string
	// Codec is the codec of the field.
	Codec Codec
}

// Ensure HiddenInputField implements InputField
var _ InputField = (*HiddenInputField)(nil)

// GetName implements FieldDefinition.GetName
func (f *HiddenInputField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *HiddenInputField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *HiddenInputField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *HiddenInputField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *HiddenInputField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *HiddenInputField) SetErrors(_ []string) {
	// do nothing. hidden fields cannot have errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *HiddenInputField) HasErrors() bool {
	return false
}

// GetErrors implements FieldDefinition.GetErrors
func (f *HiddenInputField) GetErrors() []string {
	return []string{}
}

// GetKind implements FieldDefinition.GetKind
func (f *HiddenInputField) GetKind() FieldKind {
	return FieldKindHiddenInput
}

func (f *HiddenInputField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &StringCodec{}
	}
	return codec
}
