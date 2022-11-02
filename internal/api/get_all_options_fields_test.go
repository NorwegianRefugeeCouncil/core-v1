package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	"time"
)

func Test_GetAllOptions_Field_String(t *testing.T) {
	tests := []struct {
		name     string
		field    GetAllOptions_Field
		expected string
	}{
		{name: "address", field: GetAllOptions_Field_Address, expected: "address"},
		{name: "ids", field: GetAllOptions_Field_IDs, expected: "ids"},
		{name: "birthDateFrom", field: GetAllOptions_Field_BirthDateFrom, expected: "birthDateFrom"},
		{name: "birthDateTo", field: GetAllOptions_Field_BirthDateTo, expected: "birthDateTo"},
		{name: "countryId", field: GetAllOptions_Field_CountryID, expected: "countryId"},
		{name: "displacementStatuses", field: GetAllOptions_Field_DisplacementStatuses, expected: "displacementStatuses"},
		{name: "email", field: GetAllOptions_Field_Email, expected: "email"},
		{name: "fullName", field: GetAllOptions_Field_FullName, expected: "fullName"},
		{name: "genders", field: GetAllOptions_Field_Genders, expected: "genders"},
		{name: "isMinor", field: GetAllOptions_Field_IsMinor, expected: "isMinor"},
		{name: "phoneNumber", field: GetAllOptions_Field_PhoneNumber, expected: "phoneNumber"},
		{name: "presentsProtectionConcerns", field: GetAllOptions_Field_PresentsProtectionConcerns, expected: "presentsProtectionConcerns"},
		{name: "skip", field: GetAllOptions_Field_Skip, expected: "skip"},
		{name: "take", field: GetAllOptions_Field_Take, expected: "take"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.field.String())
		})
	}
}

func Test_GetAllOptions_Field_MarshalJSON(t *testing.T) {
	type dummy struct{ Field GetAllOptions_Field }
	tests := []struct {
		name     string
		field    GetAllOptions_Field
		expected string
	}{
		{name: "address", field: GetAllOptions_Field_Address, expected: "{\"Field\":\"address\"}"},
		{name: "ids", field: GetAllOptions_Field_IDs, expected: "{\"Field\":\"ids\"}"},
		{name: "birthDateFrom", field: GetAllOptions_Field_BirthDateFrom, expected: "{\"Field\":\"birthDateFrom\"}"},
		{name: "birthDateTo", field: GetAllOptions_Field_BirthDateTo, expected: "{\"Field\":\"birthDateTo\"}"},
		{name: "countryId", field: GetAllOptions_Field_CountryID, expected: "{\"Field\":\"countryId\"}"},
		{name: "displacementStatuses", field: GetAllOptions_Field_DisplacementStatuses, expected: "{\"Field\":\"displacementStatuses\"}"},
		{name: "email", field: GetAllOptions_Field_Email, expected: "{\"Field\":\"email\"}"},
		{name: "fullName", field: GetAllOptions_Field_FullName, expected: "{\"Field\":\"fullName\"}"},
		{name: "genders", field: GetAllOptions_Field_Genders, expected: "{\"Field\":\"genders\"}"},
		{name: "isMinor", field: GetAllOptions_Field_IsMinor, expected: "{\"Field\":\"isMinor\"}"},
		{name: "phoneNumber", field: GetAllOptions_Field_PhoneNumber, expected: "{\"Field\":\"phoneNumber\"}"},
		{name: "presentsProtectionConcerns", field: GetAllOptions_Field_PresentsProtectionConcerns, expected: "{\"Field\":\"presentsProtectionConcerns\"}"},
		{name: "skip", field: GetAllOptions_Field_Skip, expected: "{\"Field\":\"skip\"}"},
		{name: "take", field: GetAllOptions_Field_Take, expected: "{\"Field\":\"take\"}"},
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

func Test_GetAllOptions_Field_UnmarshalJSON(t *testing.T) {
	type dummy struct{ Field GetAllOptions_Field }
	tests := []struct {
		name     string
		json     string
		expected dummy
	}{
		{name: "address", json: "{\"Field\":\"address\"}", expected: dummy{Field: GetAllOptions_Field_Address}},
		{name: "ids", json: "{\"Field\":\"ids\"}", expected: dummy{Field: GetAllOptions_Field_IDs}},
		{name: "birthDateFrom", json: "{\"Field\":\"birthDateFrom\"}", expected: dummy{Field: GetAllOptions_Field_BirthDateFrom}},
		{name: "birthDateTo", json: "{\"Field\":\"birthDateTo\"}", expected: dummy{Field: GetAllOptions_Field_BirthDateTo}},
		{name: "countryId", json: "{\"Field\":\"countryId\"}", expected: dummy{Field: GetAllOptions_Field_CountryID}},
		{name: "displacementStatuses", json: "{\"Field\":\"displacementStatuses\"}", expected: dummy{Field: GetAllOptions_Field_DisplacementStatuses}},
		{name: "email", json: "{\"Field\":\"email\"}", expected: dummy{Field: GetAllOptions_Field_Email}},
		{name: "fullName", json: "{\"Field\":\"fullName\"}", expected: dummy{Field: GetAllOptions_Field_FullName}},
		{name: "genders", json: "{\"Field\":\"genders\"}", expected: dummy{Field: GetAllOptions_Field_Genders}},
		{name: "isMinor", json: "{\"Field\":\"isMinor\"}", expected: dummy{Field: GetAllOptions_Field_IsMinor}},
		{name: "phoneNumber", json: "{\"Field\":\"phoneNumber\"}", expected: dummy{Field: GetAllOptions_Field_PhoneNumber}},
		{name: "presentsProtectionConcerns", json: "{\"Field\":\"presentsProtectionConcerns\"}", expected: dummy{Field: GetAllOptions_Field_PresentsProtectionConcerns}},
		{name: "skip", json: "{\"Field\":\"skip\"}", expected: dummy{Field: GetAllOptions_Field_Skip}},
		{name: "take", json: "{\"Field\":\"take\"}", expected: dummy{Field: GetAllOptions_Field_Take}},
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

func Test_GetAllOptions_Field_MarshalText(t *testing.T) {
	tests := []struct {
		name     string
		field    GetAllOptions_Field
		expected string
	}{
		{name: "address", field: GetAllOptions_Field_Address, expected: "address"},
		{name: "ids", field: GetAllOptions_Field_IDs, expected: "ids"},
		{name: "birthDateFrom", field: GetAllOptions_Field_BirthDateFrom, expected: "birthDateFrom"},
		{name: "birthDateTo", field: GetAllOptions_Field_BirthDateTo, expected: "birthDateTo"},
		{name: "countryId", field: GetAllOptions_Field_CountryID, expected: "countryId"},
		{name: "displacementStatuses", field: GetAllOptions_Field_DisplacementStatuses, expected: "displacementStatuses"},
		{name: "email", field: GetAllOptions_Field_Email, expected: "email"},
		{name: "fullName", field: GetAllOptions_Field_FullName, expected: "fullName"},
		{name: "genders", field: GetAllOptions_Field_Genders, expected: "genders"},
		{name: "isMinor", field: GetAllOptions_Field_IsMinor, expected: "isMinor"},
		{name: "phoneNumber", field: GetAllOptions_Field_PhoneNumber, expected: "phoneNumber"},
		{name: "presentsProtectionConcerns", field: GetAllOptions_Field_PresentsProtectionConcerns, expected: "presentsProtectionConcerns"},
		{name: "skip", field: GetAllOptions_Field_Skip, expected: "skip"},
		{name: "take", field: GetAllOptions_Field_Take, expected: "take"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes, err := test.field.MarshalText()
			assert.NoError(t, err)
			assert.Equal(t, test.expected, string(bytes))
		})
	}
}

func Test_GetAllOptions_Field_UnmarshalText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected GetAllOptions_Field
	}{
		{name: "address", text: "address", expected: GetAllOptions_Field_Address},
		{name: "ids", text: "ids", expected: GetAllOptions_Field_IDs},
		{name: "birthDateFrom", text: "birthDateFrom", expected: GetAllOptions_Field_BirthDateFrom},
		{name: "birthDateTo", text: "birthDateTo", expected: GetAllOptions_Field_BirthDateTo},
		{name: "countryId", text: "countryId", expected: GetAllOptions_Field_CountryID},
		{name: "displacementStatuses", text: "displacementStatuses", expected: GetAllOptions_Field_DisplacementStatuses},
		{name: "email", text: "email", expected: GetAllOptions_Field_Email},
		{name: "fullName", text: "fullName", expected: GetAllOptions_Field_FullName},
		{name: "genders", text: "genders", expected: GetAllOptions_Field_Genders},
		{name: "isMinor", text: "isMinor", expected: GetAllOptions_Field_IsMinor},
		{name: "phoneNumber", text: "phoneNumber", expected: GetAllOptions_Field_PhoneNumber},
		{name: "presentsProtectionConcerns", text: "presentsProtectionConcerns", expected: GetAllOptions_Field_PresentsProtectionConcerns},
		{name: "skip", text: "skip", expected: GetAllOptions_Field_Skip},
		{name: "take", text: "take", expected: GetAllOptions_Field_Take},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var field = new(GetAllOptions_Field)
			err := field.UnmarshalText([]byte(test.text))
			assert.NoError(t, err)
			assert.Equal(t, test.expected, *field)
		})
	}
}

func Test_GetAllOptions_AllGetAllOptions(t *testing.T) {
	assert.ElementsMatch(t, []GetAllOptions_Field{
		GetAllOptions_Field_Address,
		GetAllOptions_Field_IDs,
		GetAllOptions_Field_BirthDateFrom,
		GetAllOptions_Field_BirthDateTo,
		GetAllOptions_Field_CountryID,
		GetAllOptions_Field_DisplacementStatuses,
		GetAllOptions_Field_Email,
		GetAllOptions_Field_FullName,
		GetAllOptions_Field_Genders,
		GetAllOptions_Field_IsMinor,
		GetAllOptions_Field_PhoneNumber,
		GetAllOptions_Field_PresentsProtectionConcerns,
		GetAllOptions_Field_Skip,
		GetAllOptions_Field_Take,
	}, All_GetAllOptions_Fields())
}

func Test_Parse_GetAllOptions_Field(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected GetAllOptions_Field
	}{
		{name: "address", text: "address", expected: GetAllOptions_Field_Address},
		{name: "ids", text: "ids", expected: GetAllOptions_Field_IDs},
		{name: "birthDateFrom", text: "birthDateFrom", expected: GetAllOptions_Field_BirthDateFrom},
		{name: "birthDateTo", text: "birthDateTo", expected: GetAllOptions_Field_BirthDateTo},
		{name: "countryId", text: "countryId", expected: GetAllOptions_Field_CountryID},
		{name: "displacementStatuses", text: "displacementStatuses", expected: GetAllOptions_Field_DisplacementStatuses},
		{name: "email", text: "email", expected: GetAllOptions_Field_Email},
		{name: "fullName", text: "fullName", expected: GetAllOptions_Field_FullName},
		{name: "genders", text: "genders", expected: GetAllOptions_Field_Genders},
		{name: "isMinor", text: "isMinor", expected: GetAllOptions_Field_IsMinor},
		{name: "phoneNumber", text: "phoneNumber", expected: GetAllOptions_Field_PhoneNumber},
		{name: "presentsProtectionConcerns", text: "presentsProtectionConcerns", expected: GetAllOptions_Field_PresentsProtectionConcerns},
		{name: "skip", text: "skip", expected: GetAllOptions_Field_Skip},
		{name: "take", text: "take", expected: GetAllOptions_Field_Take},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parsed, err := Parse_GetAllOptions_Field(test.text)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, parsed)
		})
	}
}

func Test_Get_GetAllOptions_Field_Value(t *testing.T) {
	tests := []struct {
		name     string
		field    GetAllOptions_Field
		object   *GetAllOptions
		expected interface{}
	}{
		{"address", GetAllOptions_Field_Address, &GetAllOptions{Address: "mock"}, "mock"},
		{"ids", GetAllOptions_Field_IDs, &GetAllOptions{IDs: []string{"mock"}}, []string{"mock"}},
		{"birthDateFrom", GetAllOptions_Field_BirthDateFrom, &GetAllOptions{BirthDateFrom: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()},
		{"birthDateTo", GetAllOptions_Field_BirthDateTo, &GetAllOptions{BirthDateTo: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()},
		{"countryId", GetAllOptions_Field_CountryID, &GetAllOptions{CountryID: "mock"}, "mock"},
		{"displacementStatuses", GetAllOptions_Field_DisplacementStatuses, &GetAllOptions{DisplacementStatuses: []string{"mock"}}, []string{"mock"}},
		{"email", GetAllOptions_Field_Email, &GetAllOptions{Email: "mock"}, "mock"},
		{"fullName", GetAllOptions_Field_FullName, &GetAllOptions{FullName: "mock"}, "mock"},
		{"genders", GetAllOptions_Field_Genders, &GetAllOptions{Genders: []string{"mock"}}, []string{"mock"}},
		{"isMinor", GetAllOptions_Field_IsMinor, &GetAllOptions{IsMinor: func() *bool { b := true; return &b }()}, func() *bool { b := true; return &b }()},
		{"phoneNumber", GetAllOptions_Field_PhoneNumber, &GetAllOptions{PhoneNumber: "mock"}, "mock"},
		{"presentsProtectionConcerns", GetAllOptions_Field_PresentsProtectionConcerns, &GetAllOptions{PresentsProtectionConcerns: func() *bool { b := true; return &b }()}, func() *bool { b := true; return &b }()},
		{"skip", GetAllOptions_Field_Skip, &GetAllOptions{Skip: 1}, 1},
		{"take", GetAllOptions_Field_Take, &GetAllOptions{Take: 1}, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := Get_GetAllOptions_FieldValue(test.object, test.field)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, val)
		})
	}
}

func Test_Set_GetAllOptions_Field_Value(t *testing.T) {
	tests := []struct {
		name     string
		field    GetAllOptions_Field
		value    interface{}
		expected *GetAllOptions
	}{
		{"address", GetAllOptions_Field_Address, "mock", &GetAllOptions{Address: "mock"}},
		{"ids", GetAllOptions_Field_IDs, []string{"mock"}, &GetAllOptions{IDs: []string{"mock"}}},
		{"birthDateFrom", GetAllOptions_Field_BirthDateFrom, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(), &GetAllOptions{BirthDateFrom: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
		{"birthDateTo", GetAllOptions_Field_BirthDateTo, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(), &GetAllOptions{BirthDateTo: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
		{"countryId", GetAllOptions_Field_CountryID, "mock", &GetAllOptions{CountryID: "mock"}},
		{"displacementStatuses", GetAllOptions_Field_DisplacementStatuses, []string{"mock"}, &GetAllOptions{DisplacementStatuses: []string{"mock"}}},
		{"email", GetAllOptions_Field_Email, "mock", &GetAllOptions{Email: "mock"}},
		{"fullName", GetAllOptions_Field_FullName, "mock", &GetAllOptions{FullName: "mock"}},
		{"genders", GetAllOptions_Field_Genders, []string{"mock"}, &GetAllOptions{Genders: []string{"mock"}}},
		{"isMinor", GetAllOptions_Field_IsMinor, func() *bool { b := true; return &b }(), &GetAllOptions{IsMinor: func() *bool { b := true; return &b }()}},
		{"phoneNumber", GetAllOptions_Field_PhoneNumber, "mock", &GetAllOptions{PhoneNumber: "mock"}},
		{"presentsProtectionConcerns", GetAllOptions_Field_PresentsProtectionConcerns, func() *bool { b := true; return &b }(), &GetAllOptions{PresentsProtectionConcerns: func() *bool { b := true; return &b }()}},
		{"skip", GetAllOptions_Field_Skip, 1, &GetAllOptions{Skip: 1}},
		{"take", GetAllOptions_Field_Take, 1, &GetAllOptions{Take: 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			object := &GetAllOptions{}
			err := Set_GetAllOptions_FieldValue(object, test.field, test.value)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, object)
		})
	}
}

func Test_GetAllOptions_Builder(t *testing.T) {
	tests := []struct {
		name        string
		builderFunc func(builder *GetAllOptions_Builder)
		expect      *GetAllOptions
	}{
		{"address", func(builder *GetAllOptions_Builder) { builder.WithAddress("mock") }, &GetAllOptions{Address: "mock"}},
		{"ids", func(builder *GetAllOptions_Builder) { builder.WithIDs([]string{"mock"}) }, &GetAllOptions{IDs: []string{"mock"}}},
		{"birthDateFrom", func(builder *GetAllOptions_Builder) {
			builder.WithBirthDateFrom(func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }())
		}, &GetAllOptions{BirthDateFrom: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
		{"birthDateTo", func(builder *GetAllOptions_Builder) {
			builder.WithBirthDateTo(func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }())
		}, &GetAllOptions{BirthDateTo: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
		{"countryId", func(builder *GetAllOptions_Builder) { builder.WithCountryID("mock") }, &GetAllOptions{CountryID: "mock"}},
		{"displacementStatuses", func(builder *GetAllOptions_Builder) { builder.WithDisplacementStatuses([]string{"mock"}) }, &GetAllOptions{DisplacementStatuses: []string{"mock"}}},
		{"email", func(builder *GetAllOptions_Builder) { builder.WithEmail("mock") }, &GetAllOptions{Email: "mock"}},
		{"fullName", func(builder *GetAllOptions_Builder) { builder.WithFullName("mock") }, &GetAllOptions{FullName: "mock"}},
		{"genders", func(builder *GetAllOptions_Builder) { builder.WithGenders([]string{"mock"}) }, &GetAllOptions{Genders: []string{"mock"}}},
		{"isMinor", func(builder *GetAllOptions_Builder) { builder.WithIsMinor(func() *bool { b := true; return &b }()) }, &GetAllOptions{IsMinor: func() *bool { b := true; return &b }()}},
		{"phoneNumber", func(builder *GetAllOptions_Builder) { builder.WithPhoneNumber("mock") }, &GetAllOptions{PhoneNumber: "mock"}},
		{"presentsProtectionConcerns", func(builder *GetAllOptions_Builder) {
			builder.WithPresentsProtectionConcerns(func() *bool { b := true; return &b }())
		}, &GetAllOptions{PresentsProtectionConcerns: func() *bool { b := true; return &b }()}},
		{"skip", func(builder *GetAllOptions_Builder) { builder.WithSkip(1) }, &GetAllOptions{Skip: 1}},
		{"take", func(builder *GetAllOptions_Builder) { builder.WithTake(1) }, &GetAllOptions{Take: 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := New_GetAllOptions_Builder()
			test.builderFunc(b)
			assert.Equal(t, test.expect, b.Build())
		})
	}
}
