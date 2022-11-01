# Fields

For each HTML form field supported, there is a corresponding field struct. 

For non-input fields, the struct should support the `Field` interface. 
This is for e.g. paragraphs, headings, and other non-input fields.

For input fields, the struct should implement the `InputField` interface.
This is for e.g. text inputs, checkboxes, and other input fields.

The currently supported input types are
- `Checkbox`
- `Date`
- `Number`
- `Select`
- `Text`
- `Textarea`
- `ID`
- `Hidden`

## Adding fields

To add a new field, you need to add a new struct that implements the `Field` or `InputField` interface.

The current file naming convention is as follows:
- `field_input_<type>.go` for input fields
- `field_<type>.go` for non-input fields

The struct name should be `Field<Type>` where `<type>` is the type of the field.

