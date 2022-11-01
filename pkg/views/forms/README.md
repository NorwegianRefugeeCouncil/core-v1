# Fields

For each HTML form field supported, there is a corresponding field struct.

### non-input fields (`Field` interface)

For non-input fields, the struct should support the `Field` interface.
This is for e.g. paragraphs, headings, and other non-input fields.

### input fields (`InputField` interface)

For input fields, the struct should implement the `InputField` interface.
This is for e.g. text inputs, checkboxes, and other input fields that
contain data.

The currently supported fields are

- `Checkbox`
- `Date`
- `Number`
- `Select`
- `Text`
- `Textarea`
- `ID`
- `Hidden`

## Conversion

Each field type has a `Codec` that converts a string value to and from and underlying
value type. 

For example, the `Number` field has an `IntCodec` that converts a string to an `int` and
vice versa. The fields use a default codec, but this can be overridden by setting the
`Codec` field on the field struct.

This can be useful for custom fields, e.g. select fields where the string value should
be converted to an `iota` constant. Another example is for `Date` fields, where
the format of the date string can be changed.

## Adding fields

To add a new field, you need to add a new struct that implements the `Field` or `InputField` interface.

The current file naming convention is as follows:

- `field_input_<type>.go` for input fields
- `field_<type>.go` for non-input fields

The struct name should be `Field<Type>` where `<type>` is the type of the field.

Tasks

- [ ] Create the new struct in `field_<type>.go` for `Input` fields or `field_input_<type>.go` for `InputField` fields.
- [ ] Create the `FieldKind` constant for the new field in `field_kind.go`
- [ ] Register the new field kind in `pkg/views/forms/field_kind.go#init()`
- [ ] Handle the new field kind in `pkg/views/forms/form.gohtml`
- [ ] Add the new field to the list of supported fields in this file. 
- [ ] If necessary, define a new `Codec` for a field type if an existing codec is not suitable.