package forms

type FieldNamer interface {
	GetName() string
}

type Field struct {
	IDField       *IDField
	Text          *TextInputField
	Number        *NumberInputField
	Date          *DateInputField
	Select        *SelectInputField
	Checkbox      *CheckboxInputField
	MultilineText *MultilineTextInputField
	Hidden        *HiddenInputField
}

func (f Field) GetName() string {
	switch {
	case f.IDField != nil:
		return f.IDField.GetName()
	case f.Text != nil:
		return f.Text.GetName()
	case f.Number != nil:
		return f.Number.GetName()
	case f.Date != nil:
		return f.Date.GetName()
	case f.Select != nil:
		return f.Select.GetName()
	case f.Checkbox != nil:
		return f.Checkbox.GetName()
	case f.MultilineText != nil:
		return f.MultilineText.GetName()
	case f.Hidden != nil:
		return f.Hidden.GetName()
	default:
		return ""
	}
}

func (f Field) SetValue(value string) {
	switch {
	case f.IDField != nil:
		f.IDField.Value = value
	case f.Text != nil:
		f.Text.Value = value
	case f.Number != nil:
		f.Number.Value = value
	case f.Date != nil:
		f.Date.Value = value
	case f.Select != nil:
		f.Select.Value = value
	case f.Checkbox != nil:
		f.Checkbox.Value = value
	case f.MultilineText != nil:
		f.MultilineText.Value = value
	case f.Hidden != nil:
		f.Hidden.Value = value
	}
}

func (f Field) GetValue() string {
	switch {
	case f.IDField != nil:
		return f.IDField.Value
	case f.Text != nil:
		return f.Text.Value
	case f.Number != nil:
		return f.Number.Value
	case f.Date != nil:
		return f.Date.Value
	case f.Select != nil:
		return f.Select.Value
	case f.Checkbox != nil:
		return f.Checkbox.Value
	case f.MultilineText != nil:
		return f.MultilineText.Value
	case f.Hidden != nil:
		return f.Hidden.Value
	default:
		return ""
	}
}

func (f Field) SetErrors(errors []string) {
	switch {
	case f.IDField != nil:
		f.IDField.Errors = errors
	case f.Text != nil:
		f.Text.Errors = errors
	case f.Number != nil:
		f.Number.Errors = errors
	case f.Date != nil:
		f.Date.Errors = errors
	case f.Select != nil:
		f.Select.Errors = errors
	case f.Checkbox != nil:
		f.Checkbox.Errors = errors
	case f.MultilineText != nil:
		f.MultilineText.Errors = errors
	}
}

type IDField struct {
	QRCodeURL   string
	Name        string
	DisplayName string
	Required    bool
	Value       string
	Help        string
	Errors      []string
}

func (f *IDField) GetName() string {
	return f.Name
}

type TextInputField struct {
	Name        string
	DisplayName string
	Required    bool
	Value       string
	Help        string
	Errors      []string
}

func (f *TextInputField) GetName() string {
	return f.Name
}

type NumberInputField struct {
	Name        string
	DisplayName string
	Required    bool
	Value       string
	Help        string
	Errors      []string
}

func (f *NumberInputField) GetName() string {
	return f.Name
}

type DateInputField struct {
	Name        string
	DisplayName string
	Required    bool
	Value       string
	Help        string
	Errors      []string
}

func (f *DateInputField) GetName() string {
	return f.Name
}

type SelectInputFieldOption struct {
	Value string
	Label string
}

type SelectInputField struct {
	Name        string
	DisplayName string
	Required    bool
	Value       string
	Help        string
	Errors      []string
	Options     []SelectInputFieldOption
}

func (f *SelectInputField) GetName() string {
	return f.Name
}

type RadioInputFieldOption struct {
	Value string
	Label string
}

type CheckboxInputField struct {
	Name        string
	DisplayName string
	Required    bool
	Value       string
	Help        string
	Errors      []string
}

func (f *CheckboxInputField) GetName() string {
	return f.Name
}

type MultilineTextInputField struct {
	Name        string
	DisplayName string
	Required    bool
	Value       string
	Help        string
	Errors      []string
}

func (f *MultilineTextInputField) GetName() string {
	return f.Name
}

type HiddenInputField struct {
	Name  string
	Value string
}

func (f *HiddenInputField) GetName() string {
	return f.Name
}
