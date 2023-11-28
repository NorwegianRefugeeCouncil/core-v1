package enumTypes

import (
	"encoding/json"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func optionalBooleanPtr(g OptionalBoolean) *OptionalBoolean {
	return &g
}

func TestOptionalBoolean_MarshalJSON(t *testing.T) {
	type dummy struct {
		OptionalBoolean OptionalBoolean `json:"d"`
	}
	type dummyPtr struct {
		OptionalBoolean *OptionalBoolean `json:"d"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{OptionalBoolean: OptionalBooleanYes})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"d":"yes"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{OptionalBoolean: optionalBooleanPtr(OptionalBooleanNo)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"d":"no"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"d":null}`), jsonBytes)
	}
}

func TestOptionalBoolean_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		OptionalBoolean OptionalBoolean `json:"d"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"d":""}`), &d))
	assert.Equal(t, OptionalBooleanUnknown, d.OptionalBoolean)

	type dummyPtr struct {
		OptionalBoolean *OptionalBoolean `json:"d"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"d":"no"}`), &dPtr))
		assert.Equal(t, optionalBooleanPtr(OptionalBooleanFalse), dPtr.OptionalBoolean)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"d":"YES"}`), &dPtr))
		assert.Equal(t, optionalBooleanPtr(OptionalBooleanTrue), dPtr.OptionalBoolean)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"d":null}`), &dPtr))
		assert.Equal(t, (*OptionalBoolean)(nil), dPtr.OptionalBoolean)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"d":"invalid"}`), &dPtr))
	}
}

func TestOptionalBoolean_String(t *testing.T) {
	tr := locales.GetTranslator()
	tests := []struct {
		name string
		g    OptionalBoolean
		want string
	}{
		{"no", OptionalBooleanNo, tr("no")},
		{"yes", OptionalBooleanYes, tr("yes")},
		{"", OptionalBooleanUnknown, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestOptionalBoolean_BoolPtr(t *testing.T) {
	tests := []struct {
		name string
		g    OptionalBoolean
		want *bool
	}{
		{"no", OptionalBooleanNo, pointers.Bool(false)},
		{"0", OptionalBoolean0, pointers.Bool(false)},
		{"false", OptionalBooleanFalse, pointers.Bool(false)},
		{"yes", OptionalBooleanYes, pointers.Bool(true)},
		{"1", OptionalBoolean1, pointers.Bool(true)},
		{"true", OptionalBooleanTrue, pointers.Bool(true)},
		{"", OptionalBooleanUnknown, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.BoolPtr())
		})
	}
}

func TestOptionalBoolean_MarshalText(t *testing.T) {
	{
		got, err := OptionalBooleanNo.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "no", string(got))
	}
	{
		got, err := OptionalBooleanYes.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "yes", string(got))
	}
	{
		got, err := OptionalBooleanUnknown.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "", string(got))
	}
}

func TestOptionalBoolean_UnmarshalText(t *testing.T) {
	{
		var g = new(OptionalBoolean)
		assert.NoError(t, g.UnmarshalText([]byte("no")))
		assert.Equal(t, OptionalBooleanFalse, *g)
	}
	{
		var g = new(OptionalBoolean)
		assert.NoError(t, g.UnmarshalText([]byte("yes")))
		assert.Equal(t, OptionalBooleanTrue, *g)
	}
	{
		var g = new(OptionalBoolean)
		assert.NoError(t, g.UnmarshalText([]byte("")))
		assert.Equal(t, OptionalBooleanUnknown, *g)
	}
}

func TestParseOptionalBoolean(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    OptionalBoolean
		wantErr bool
	}{
		{"no", "0", OptionalBooleanFalse, false},
		{"no", "NO", OptionalBooleanFalse, false},
		{"no", "No", OptionalBooleanFalse, false},
		{"no", "no", OptionalBooleanFalse, false},
		{"no", "FALSE", OptionalBooleanFalse, false},
		{"no", "faLSE", OptionalBooleanFalse, false},
		{"no", "False", OptionalBooleanFalse, false},
		{"no", "false", OptionalBooleanFalse, false},
		{"yes", "1", OptionalBooleanTrue, false},
		{"yes", "YES", OptionalBooleanTrue, false},
		{"yes", "YeS", OptionalBooleanTrue, false},
		{"yes", "Yes", OptionalBooleanTrue, false},
		{"yes", "yes", OptionalBooleanTrue, false},
		{"yes", "TRUE", OptionalBooleanTrue, false},
		{"yes", "True", OptionalBooleanTrue, false},
		{"yes", "true", OptionalBooleanTrue, false},
		{"", "", OptionalBooleanUnknown, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOptionalBoolean(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllOptionalBooleans(t *testing.T) {
	assert.ElementsMatch(t, []OptionalBoolean{
		OptionalBooleanNo,
		OptionalBooleanFalse,
		OptionalBoolean0,
		OptionalBooleanTrue,
		OptionalBoolean1,
		OptionalBooleanYes,
	}, AllOptionalBooleans().Items())
}
