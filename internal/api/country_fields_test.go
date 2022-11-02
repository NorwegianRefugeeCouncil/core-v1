package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Country_Field_String(t *testing.T) {
	tests := []struct {
		name     string
		field    Country_Field
		expected string
	}{
		{name: "id", field: Country_Field_ID, expected: "id"},
		{name: "code", field: Country_Field_Code, expected: "code"},
		{name: "name", field: Country_Field_Name, expected: "name"},
		{name: "jwtGroup", field: Country_Field_JwtGroup, expected: "jwtGroup"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.field.String())
		})
	}
}

func Test_Country_Field_MarshalJSON(t *testing.T) {
	type dummy struct{ Field Country_Field }
	tests := []struct {
		name     string
		field    Country_Field
		expected string
	}{
		{name: "id", field: Country_Field_ID, expected: "{\"Field\":\"id\"}"},
		{name: "code", field: Country_Field_Code, expected: "{\"Field\":\"code\"}"},
		{name: "name", field: Country_Field_Name, expected: "{\"Field\":\"name\"}"},
		{name: "jwtGroup", field: Country_Field_JwtGroup, expected: "{\"Field\":\"jwtGroup\"}"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := dummy{Field: test.field}
			bytes, err := json.Marshal(d)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, string(bytes))
		})
	}
}

func Test_Country_Field_UnmarshalJSON(t *testing.T) {
	type dummy struct{ Field Country_Field }
	tests := []struct {
		name     string
		json     string
		expected dummy
	}{
		{name: "id", json: "{\"Field\":\"id\"}", expected: dummy{Field: Country_Field_ID}},
		{name: "code", json: "{\"Field\":\"code\"}", expected: dummy{Field: Country_Field_Code}},
		{name: "name", json: "{\"Field\":\"name\"}", expected: dummy{Field: Country_Field_Name}},
		{name: "jwtGroup", json: "{\"Field\":\"jwtGroup\"}", expected: dummy{Field: Country_Field_JwtGroup}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var d dummy
			err := json.Unmarshal([]byte(test.json), &d)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, d)
		})
	}
}

func Test_Country_Field_MarshalText(t *testing.T) {
	tests := []struct {
		name     string
		field    Country_Field
		expected string
	}{
		{name: "id", field: Country_Field_ID, expected: "id"},
		{name: "code", field: Country_Field_Code, expected: "code"},
		{name: "name", field: Country_Field_Name, expected: "name"},
		{name: "jwtGroup", field: Country_Field_JwtGroup, expected: "jwtGroup"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes, err := test.field.MarshalText()
			assert.NoError(t, err)
			assert.Equal(t, test.expected, string(bytes))
		})
	}
}

func Test_Country_Field_UnmarshalText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected Country_Field
	}{
		{name: "id", text: "id", expected: Country_Field_ID},
		{name: "code", text: "code", expected: Country_Field_Code},
		{name: "name", text: "name", expected: Country_Field_Name},
		{name: "jwtGroup", text: "jwtGroup", expected: Country_Field_JwtGroup},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var field = new(Country_Field)
			err := field.UnmarshalText([]byte(test.text))
			assert.NoError(t, err)
			assert.Equal(t, test.expected, *field)
		})
	}
}

func Test_Country_AllCountry(t *testing.T) {
	assert.ElementsMatch(t, []Country_Field{
		Country_Field_ID,
		Country_Field_Code,
		Country_Field_Name,
		Country_Field_JwtGroup,
	}, All_Country_Fields())
}

func Test_Parse_Country_Field(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected Country_Field
	}{
		{name: "id", text: "id", expected: Country_Field_ID},
		{name: "code", text: "code", expected: Country_Field_Code},
		{name: "name", text: "name", expected: Country_Field_Name},
		{name: "jwtGroup", text: "jwtGroup", expected: Country_Field_JwtGroup},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parsed, err := Parse_Country_Field(test.text)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, parsed)
		})
	}
}

func Test_Get_Country_Field_Value(t *testing.T) {
	tests := []struct {
		name     string
		field    Country_Field
		object   *Country
		expected interface{}
	}{
		{"id", Country_Field_ID, &Country{ID: "mock"}, "mock"},
		{"code", Country_Field_Code, &Country{Code: "mock"}, "mock"},
		{"name", Country_Field_Name, &Country{Name: "mock"}, "mock"},
		{"jwtGroup", Country_Field_JwtGroup, &Country{JwtGroup: "mock"}, "mock"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := Get_Country_FieldValue(test.object, test.field)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, val)
		})
	}
}

func Test_Set_Country_Field_Value(t *testing.T) {
	tests := []struct {
		name     string
		field    Country_Field
		value    interface{}
		expected *Country
	}{
		{"id", Country_Field_ID, "mock", &Country{ID: "mock"}},
		{"code", Country_Field_Code, "mock", &Country{Code: "mock"}},
		{"name", Country_Field_Name, "mock", &Country{Name: "mock"}},
		{"jwtGroup", Country_Field_JwtGroup, "mock", &Country{JwtGroup: "mock"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			object := &Country{}
			err := Set_Country_FieldValue(object, test.field, test.value)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, object)
		})
	}
}

func Test_Country_Builder(t *testing.T) {
	tests := []struct {
		name        string
		builderFunc func(builder *Country_Builder)
		expect      *Country
	}{
		{"id", func(builder *Country_Builder) { builder.WithID("mock") }, &Country{ID: "mock"}},
		{"code", func(builder *Country_Builder) { builder.WithCode("mock") }, &Country{Code: "mock"}},
		{"name", func(builder *Country_Builder) { builder.WithName("mock") }, &Country{Name: "mock"}},
		{"jwtGroup", func(builder *Country_Builder) { builder.WithJwtGroup("mock") }, &Country{JwtGroup: "mock"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := New_Country_Builder()
			test.builderFunc(b)
			assert.Equal(t, test.expect, b.Build())
		})
	}
}
