package api

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Individual_Field_String(t *testing.T) {
	tests := []struct {
		name     string
		field    Individual_Field
		expected string
	}{
		{name: "id", field: Individual_Field_ID, expected: "id"},
		{name: "countryId", field: Individual_Field_CountryID, expected: "countryId"},
		{name: "fullName", field: Individual_Field_FullName, expected: "fullName"},
		{name: "phoneNumber", field: Individual_Field_PhoneNumber, expected: "phoneNumber"},
		{name: "normalizedPhoneNumber", field: Individual_Field_NormalizedPhoneNumber, expected: "normalizedPhoneNumber"},
		{name: "email", field: Individual_Field_Email, expected: "email"},
		{name: "address", field: Individual_Field_Address, expected: "address"},
		{name: "birthDate", field: Individual_Field_BirthDate, expected: "birthDate"},
		{name: "gender", field: Individual_Field_Gender, expected: "gender"},
		{name: "displacementStatus", field: Individual_Field_DisplacementStatus, expected: "displacementStatus"},
		{name: "preferredName", field: Individual_Field_PreferredName, expected: "preferredName"},
		{name: "isMinor", field: Individual_Field_IsMinor, expected: "isMinor"},
		{name: "presentsProtectionConcerns", field: Individual_Field_PresentsProtectionConcerns, expected: "presentsProtectionConcerns"},
		{name: "physicalImpairment", field: Individual_Field_PhysicalImpairment, expected: "physicalImpairment"},
		{name: "sensoryImpairment", field: Individual_Field_SensoryImpairment, expected: "sensoryImpairment"},
		{name: "mentalImpairment", field: Individual_Field_MentalImpairment, expected: "mentalImpairment"},
		{name: "createdAt", field: Individual_Field_CreatedAt, expected: "createdAt"},
		{name: "updatedAt", field: Individual_Field_UpdatedAt, expected: "updatedAt"},
		{name: "deletedAt", field: Individual_Field_DeletedAt, expected: "deletedAt"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.field.String())
		})
	}
}

func Test_Individual_Field_MarshalJSON(t *testing.T) {
	type dummy struct{ Field Individual_Field }
	tests := []struct {
		name     string
		field    Individual_Field
		expected string
	}{
		{name: "id", field: Individual_Field_ID, expected: "{\"Field\":\"id\"}"},
		{name: "countryId", field: Individual_Field_CountryID, expected: "{\"Field\":\"countryId\"}"},
		{name: "fullName", field: Individual_Field_FullName, expected: "{\"Field\":\"fullName\"}"},
		{name: "phoneNumber", field: Individual_Field_PhoneNumber, expected: "{\"Field\":\"phoneNumber\"}"},
		{name: "normalizedPhoneNumber", field: Individual_Field_NormalizedPhoneNumber, expected: "{\"Field\":\"normalizedPhoneNumber\"}"},
		{name: "email", field: Individual_Field_Email, expected: "{\"Field\":\"email\"}"},
		{name: "address", field: Individual_Field_Address, expected: "{\"Field\":\"address\"}"},
		{name: "birthDate", field: Individual_Field_BirthDate, expected: "{\"Field\":\"birthDate\"}"},
		{name: "gender", field: Individual_Field_Gender, expected: "{\"Field\":\"gender\"}"},
		{name: "displacementStatus", field: Individual_Field_DisplacementStatus, expected: "{\"Field\":\"displacementStatus\"}"},
		{name: "preferredName", field: Individual_Field_PreferredName, expected: "{\"Field\":\"preferredName\"}"},
		{name: "isMinor", field: Individual_Field_IsMinor, expected: "{\"Field\":\"isMinor\"}"},
		{name: "presentsProtectionConcerns", field: Individual_Field_PresentsProtectionConcerns, expected: "{\"Field\":\"presentsProtectionConcerns\"}"},
		{name: "physicalImpairment", field: Individual_Field_PhysicalImpairment, expected: "{\"Field\":\"physicalImpairment\"}"},
		{name: "sensoryImpairment", field: Individual_Field_SensoryImpairment, expected: "{\"Field\":\"sensoryImpairment\"}"},
		{name: "mentalImpairment", field: Individual_Field_MentalImpairment, expected: "{\"Field\":\"mentalImpairment\"}"},
		{name: "createdAt", field: Individual_Field_CreatedAt, expected: "{\"Field\":\"createdAt\"}"},
		{name: "updatedAt", field: Individual_Field_UpdatedAt, expected: "{\"Field\":\"updatedAt\"}"},
		{name: "deletedAt", field: Individual_Field_DeletedAt, expected: "{\"Field\":\"deletedAt\"}"},
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

func Test_Individual_Field_UnmarshalJSON(t *testing.T) {
	type dummy struct{ Field Individual_Field }
	tests := []struct {
		name     string
		json     string
		expected dummy
	}{
		{name: "id", json: "{\"Field\":\"id\"}", expected: dummy{Field: Individual_Field_ID}},
		{name: "countryId", json: "{\"Field\":\"countryId\"}", expected: dummy{Field: Individual_Field_CountryID}},
		{name: "fullName", json: "{\"Field\":\"fullName\"}", expected: dummy{Field: Individual_Field_FullName}},
		{name: "phoneNumber", json: "{\"Field\":\"phoneNumber\"}", expected: dummy{Field: Individual_Field_PhoneNumber}},
		{name: "normalizedPhoneNumber", json: "{\"Field\":\"normalizedPhoneNumber\"}", expected: dummy{Field: Individual_Field_NormalizedPhoneNumber}},
		{name: "email", json: "{\"Field\":\"email\"}", expected: dummy{Field: Individual_Field_Email}},
		{name: "address", json: "{\"Field\":\"address\"}", expected: dummy{Field: Individual_Field_Address}},
		{name: "birthDate", json: "{\"Field\":\"birthDate\"}", expected: dummy{Field: Individual_Field_BirthDate}},
		{name: "gender", json: "{\"Field\":\"gender\"}", expected: dummy{Field: Individual_Field_Gender}},
		{name: "displacementStatus", json: "{\"Field\":\"displacementStatus\"}", expected: dummy{Field: Individual_Field_DisplacementStatus}},
		{name: "preferredName", json: "{\"Field\":\"preferredName\"}", expected: dummy{Field: Individual_Field_PreferredName}},
		{name: "isMinor", json: "{\"Field\":\"isMinor\"}", expected: dummy{Field: Individual_Field_IsMinor}},
		{name: "presentsProtectionConcerns", json: "{\"Field\":\"presentsProtectionConcerns\"}", expected: dummy{Field: Individual_Field_PresentsProtectionConcerns}},
		{name: "physicalImpairment", json: "{\"Field\":\"physicalImpairment\"}", expected: dummy{Field: Individual_Field_PhysicalImpairment}},
		{name: "sensoryImpairment", json: "{\"Field\":\"sensoryImpairment\"}", expected: dummy{Field: Individual_Field_SensoryImpairment}},
		{name: "mentalImpairment", json: "{\"Field\":\"mentalImpairment\"}", expected: dummy{Field: Individual_Field_MentalImpairment}},
		{name: "createdAt", json: "{\"Field\":\"createdAt\"}", expected: dummy{Field: Individual_Field_CreatedAt}},
		{name: "updatedAt", json: "{\"Field\":\"updatedAt\"}", expected: dummy{Field: Individual_Field_UpdatedAt}},
		{name: "deletedAt", json: "{\"Field\":\"deletedAt\"}", expected: dummy{Field: Individual_Field_DeletedAt}},
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

func Test_Individual_Field_MarshalText(t *testing.T) {
	tests := []struct {
		name     string
		field    Individual_Field
		expected string
	}{
		{name: "id", field: Individual_Field_ID, expected: "id"},
		{name: "countryId", field: Individual_Field_CountryID, expected: "countryId"},
		{name: "fullName", field: Individual_Field_FullName, expected: "fullName"},
		{name: "phoneNumber", field: Individual_Field_PhoneNumber, expected: "phoneNumber"},
		{name: "normalizedPhoneNumber", field: Individual_Field_NormalizedPhoneNumber, expected: "normalizedPhoneNumber"},
		{name: "email", field: Individual_Field_Email, expected: "email"},
		{name: "address", field: Individual_Field_Address, expected: "address"},
		{name: "birthDate", field: Individual_Field_BirthDate, expected: "birthDate"},
		{name: "gender", field: Individual_Field_Gender, expected: "gender"},
		{name: "displacementStatus", field: Individual_Field_DisplacementStatus, expected: "displacementStatus"},
		{name: "preferredName", field: Individual_Field_PreferredName, expected: "preferredName"},
		{name: "isMinor", field: Individual_Field_IsMinor, expected: "isMinor"},
		{name: "presentsProtectionConcerns", field: Individual_Field_PresentsProtectionConcerns, expected: "presentsProtectionConcerns"},
		{name: "physicalImpairment", field: Individual_Field_PhysicalImpairment, expected: "physicalImpairment"},
		{name: "sensoryImpairment", field: Individual_Field_SensoryImpairment, expected: "sensoryImpairment"},
		{name: "mentalImpairment", field: Individual_Field_MentalImpairment, expected: "mentalImpairment"},
		{name: "createdAt", field: Individual_Field_CreatedAt, expected: "createdAt"},
		{name: "updatedAt", field: Individual_Field_UpdatedAt, expected: "updatedAt"},
		{name: "deletedAt", field: Individual_Field_DeletedAt, expected: "deletedAt"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes, err := test.field.MarshalText()
			assert.NoError(t, err)
			assert.Equal(t, test.expected, string(bytes))
		})
	}
}

func Test_Individual_Field_UnmarshalText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected Individual_Field
	}{
		{name: "id", text: "id", expected: Individual_Field_ID},
		{name: "countryId", text: "countryId", expected: Individual_Field_CountryID},
		{name: "fullName", text: "fullName", expected: Individual_Field_FullName},
		{name: "phoneNumber", text: "phoneNumber", expected: Individual_Field_PhoneNumber},
		{name: "normalizedPhoneNumber", text: "normalizedPhoneNumber", expected: Individual_Field_NormalizedPhoneNumber},
		{name: "email", text: "email", expected: Individual_Field_Email},
		{name: "address", text: "address", expected: Individual_Field_Address},
		{name: "birthDate", text: "birthDate", expected: Individual_Field_BirthDate},
		{name: "gender", text: "gender", expected: Individual_Field_Gender},
		{name: "displacementStatus", text: "displacementStatus", expected: Individual_Field_DisplacementStatus},
		{name: "preferredName", text: "preferredName", expected: Individual_Field_PreferredName},
		{name: "isMinor", text: "isMinor", expected: Individual_Field_IsMinor},
		{name: "presentsProtectionConcerns", text: "presentsProtectionConcerns", expected: Individual_Field_PresentsProtectionConcerns},
		{name: "physicalImpairment", text: "physicalImpairment", expected: Individual_Field_PhysicalImpairment},
		{name: "sensoryImpairment", text: "sensoryImpairment", expected: Individual_Field_SensoryImpairment},
		{name: "mentalImpairment", text: "mentalImpairment", expected: Individual_Field_MentalImpairment},
		{name: "createdAt", text: "createdAt", expected: Individual_Field_CreatedAt},
		{name: "updatedAt", text: "updatedAt", expected: Individual_Field_UpdatedAt},
		{name: "deletedAt", text: "deletedAt", expected: Individual_Field_DeletedAt},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var field = new(Individual_Field)
			err := field.UnmarshalText([]byte(test.text))
			assert.NoError(t, err)
			assert.Equal(t, test.expected, *field)
		})
	}
}

func Test_Individual_AllIndividual(t *testing.T) {
	assert.ElementsMatch(t, []Individual_Field{
		Individual_Field_ID,
		Individual_Field_CountryID,
		Individual_Field_FullName,
		Individual_Field_PhoneNumber,
		Individual_Field_NormalizedPhoneNumber,
		Individual_Field_Email,
		Individual_Field_Address,
		Individual_Field_BirthDate,
		Individual_Field_Gender,
		Individual_Field_DisplacementStatus,
		Individual_Field_PreferredName,
		Individual_Field_IsMinor,
		Individual_Field_PresentsProtectionConcerns,
		Individual_Field_PhysicalImpairment,
		Individual_Field_SensoryImpairment,
		Individual_Field_MentalImpairment,
		Individual_Field_CreatedAt,
		Individual_Field_UpdatedAt,
		Individual_Field_DeletedAt,
	}, All_Individual_Fields())
}

func Test_Parse_Individual_Field(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected Individual_Field
	}{
		{name: "id", text: "id", expected: Individual_Field_ID},
		{name: "countryId", text: "countryId", expected: Individual_Field_CountryID},
		{name: "fullName", text: "fullName", expected: Individual_Field_FullName},
		{name: "phoneNumber", text: "phoneNumber", expected: Individual_Field_PhoneNumber},
		{name: "normalizedPhoneNumber", text: "normalizedPhoneNumber", expected: Individual_Field_NormalizedPhoneNumber},
		{name: "email", text: "email", expected: Individual_Field_Email},
		{name: "address", text: "address", expected: Individual_Field_Address},
		{name: "birthDate", text: "birthDate", expected: Individual_Field_BirthDate},
		{name: "gender", text: "gender", expected: Individual_Field_Gender},
		{name: "displacementStatus", text: "displacementStatus", expected: Individual_Field_DisplacementStatus},
		{name: "preferredName", text: "preferredName", expected: Individual_Field_PreferredName},
		{name: "isMinor", text: "isMinor", expected: Individual_Field_IsMinor},
		{name: "presentsProtectionConcerns", text: "presentsProtectionConcerns", expected: Individual_Field_PresentsProtectionConcerns},
		{name: "physicalImpairment", text: "physicalImpairment", expected: Individual_Field_PhysicalImpairment},
		{name: "sensoryImpairment", text: "sensoryImpairment", expected: Individual_Field_SensoryImpairment},
		{name: "mentalImpairment", text: "mentalImpairment", expected: Individual_Field_MentalImpairment},
		{name: "createdAt", text: "createdAt", expected: Individual_Field_CreatedAt},
		{name: "updatedAt", text: "updatedAt", expected: Individual_Field_UpdatedAt},
		{name: "deletedAt", text: "deletedAt", expected: Individual_Field_DeletedAt},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parsed, err := Parse_Individual_Field(test.text)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, parsed)
		})
	}
}

func Test_Get_Individual_Field_Value(t *testing.T) {
	tests := []struct {
		name     string
		field    Individual_Field
		object   *Individual
		expected interface{}
	}{
		{"id", Individual_Field_ID, &Individual{ID: "mock"}, "mock"},
		{"countryId", Individual_Field_CountryID, &Individual{CountryID: "mock"}, "mock"},
		{"fullName", Individual_Field_FullName, &Individual{FullName: "mock"}, "mock"},
		{"phoneNumber", Individual_Field_PhoneNumber, &Individual{PhoneNumber: "mock"}, "mock"},
		{"normalizedPhoneNumber", Individual_Field_NormalizedPhoneNumber, &Individual{NormalizedPhoneNumber: "mock"}, "mock"},
		{"email", Individual_Field_Email, &Individual{Email: "mock"}, "mock"},
		{"address", Individual_Field_Address, &Individual{Address: "mock"}, "mock"},
		{"birthDate", Individual_Field_BirthDate, &Individual{BirthDate: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()},
		{"gender", Individual_Field_Gender, &Individual{Gender: "mock"}, "mock"},
		{"displacementStatus", Individual_Field_DisplacementStatus, &Individual{DisplacementStatus: "mock"}, "mock"},
		{"preferredName", Individual_Field_PreferredName, &Individual{PreferredName: "mock"}, "mock"},
		{"isMinor", Individual_Field_IsMinor, &Individual{IsMinor: true}, true},
		{"presentsProtectionConcerns", Individual_Field_PresentsProtectionConcerns, &Individual{PresentsProtectionConcerns: true}, true},
		{"physicalImpairment", Individual_Field_PhysicalImpairment, &Individual{PhysicalImpairment: "mock"}, "mock"},
		{"sensoryImpairment", Individual_Field_SensoryImpairment, &Individual{SensoryImpairment: "mock"}, "mock"},
		{"mentalImpairment", Individual_Field_MentalImpairment, &Individual{MentalImpairment: "mock"}, "mock"},
		{"createdAt", Individual_Field_CreatedAt, &Individual{CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)}, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"updatedAt", Individual_Field_UpdatedAt, &Individual{UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)}, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"deletedAt", Individual_Field_DeletedAt, &Individual{DeletedAt: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := Get_Individual_FieldValue(test.object, test.field)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, val)
		})
	}
}

func Test_Set_Individual_Field_Value(t *testing.T) {
	tests := []struct {
		name     string
		field    Individual_Field
		value    interface{}
		expected *Individual
	}{
		{"id", Individual_Field_ID, "mock", &Individual{ID: "mock"}},
		{"countryId", Individual_Field_CountryID, "mock", &Individual{CountryID: "mock"}},
		{"fullName", Individual_Field_FullName, "mock", &Individual{FullName: "mock"}},
		{"phoneNumber", Individual_Field_PhoneNumber, "mock", &Individual{PhoneNumber: "mock"}},
		{"normalizedPhoneNumber", Individual_Field_NormalizedPhoneNumber, "mock", &Individual{NormalizedPhoneNumber: "mock"}},
		{"email", Individual_Field_Email, "mock", &Individual{Email: "mock"}},
		{"address", Individual_Field_Address, "mock", &Individual{Address: "mock"}},
		{"birthDate", Individual_Field_BirthDate, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(), &Individual{BirthDate: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
		{"gender", Individual_Field_Gender, "mock", &Individual{Gender: "mock"}},
		{"displacementStatus", Individual_Field_DisplacementStatus, "mock", &Individual{DisplacementStatus: "mock"}},
		{"preferredName", Individual_Field_PreferredName, "mock", &Individual{PreferredName: "mock"}},
		{"isMinor", Individual_Field_IsMinor, true, &Individual{IsMinor: true}},
		{"presentsProtectionConcerns", Individual_Field_PresentsProtectionConcerns, true, &Individual{PresentsProtectionConcerns: true}},
		{"physicalImpairment", Individual_Field_PhysicalImpairment, "mock", &Individual{PhysicalImpairment: "mock"}},
		{"sensoryImpairment", Individual_Field_SensoryImpairment, "mock", &Individual{SensoryImpairment: "mock"}},
		{"mentalImpairment", Individual_Field_MentalImpairment, "mock", &Individual{MentalImpairment: "mock"}},
		{"createdAt", Individual_Field_CreatedAt, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), &Individual{CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)}},
		{"updatedAt", Individual_Field_UpdatedAt, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), &Individual{UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)}},
		{"deletedAt", Individual_Field_DeletedAt, func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(), &Individual{DeletedAt: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			object := &Individual{}
			err := Set_Individual_FieldValue(object, test.field, test.value)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, object)
		})
	}
}

func Test_Individual_Builder(t *testing.T) {
	tests := []struct {
		name        string
		builderFunc func(builder *Individual_Builder)
		expect      *Individual
	}{
		{"id", func(builder *Individual_Builder) { builder.WithID("mock") }, &Individual{ID: "mock"}},
		{"countryId", func(builder *Individual_Builder) { builder.WithCountryID("mock") }, &Individual{CountryID: "mock"}},
		{"fullName", func(builder *Individual_Builder) { builder.WithFullName("mock") }, &Individual{FullName: "mock"}},
		{"phoneNumber", func(builder *Individual_Builder) { builder.WithPhoneNumber("mock") }, &Individual{PhoneNumber: "mock"}},
		{"normalizedPhoneNumber", func(builder *Individual_Builder) { builder.WithNormalizedPhoneNumber("mock") }, &Individual{NormalizedPhoneNumber: "mock"}},
		{"email", func(builder *Individual_Builder) { builder.WithEmail("mock") }, &Individual{Email: "mock"}},
		{"address", func(builder *Individual_Builder) { builder.WithAddress("mock") }, &Individual{Address: "mock"}},
		{"birthDate", func(builder *Individual_Builder) {
			builder.WithBirthDate(func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }())
		}, &Individual{BirthDate: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
		{"gender", func(builder *Individual_Builder) { builder.WithGender("mock") }, &Individual{Gender: "mock"}},
		{"displacementStatus", func(builder *Individual_Builder) { builder.WithDisplacementStatus("mock") }, &Individual{DisplacementStatus: "mock"}},
		{"preferredName", func(builder *Individual_Builder) { builder.WithPreferredName("mock") }, &Individual{PreferredName: "mock"}},
		{"isMinor", func(builder *Individual_Builder) { builder.WithIsMinor(true) }, &Individual{IsMinor: true}},
		{"presentsProtectionConcerns", func(builder *Individual_Builder) { builder.WithPresentsProtectionConcerns(true) }, &Individual{PresentsProtectionConcerns: true}},
		{"physicalImpairment", func(builder *Individual_Builder) { builder.WithPhysicalImpairment("mock") }, &Individual{PhysicalImpairment: "mock"}},
		{"sensoryImpairment", func(builder *Individual_Builder) { builder.WithSensoryImpairment("mock") }, &Individual{SensoryImpairment: "mock"}},
		{"mentalImpairment", func(builder *Individual_Builder) { builder.WithMentalImpairment("mock") }, &Individual{MentalImpairment: "mock"}},
		{"createdAt", func(builder *Individual_Builder) { builder.WithCreatedAt(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)) }, &Individual{CreatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)}},
		{"updatedAt", func(builder *Individual_Builder) { builder.WithUpdatedAt(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)) }, &Individual{UpdatedAt: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)}},
		{"deletedAt", func(builder *Individual_Builder) {
			builder.WithDeletedAt(func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }())
		}, &Individual{DeletedAt: func() *time.Time { t := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC); return &t }()}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := New_Individual_Builder()
			test.builderFunc(b)
			assert.Equal(t, test.expect, b.Build())
		})
	}
}
