package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func genderPtr(g Gender) *Gender {
	return &g
}

func TestGender_MarshalJSON(t *testing.T) {
	type dummy struct {
		Gender Gender `json:"g"`
	}
	type dummyPtr struct {
		Gender *Gender `json:"g"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{Gender: GenderMale})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"male"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{Gender: genderPtr(GenderMale)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"male"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":null}`), jsonBytes)
	}
}

func TestGender_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		Gender Gender `json:"g"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"g":"male"}`), &d))
	assert.Equal(t, GenderMale, d.Gender)

	type dummyPtr struct {
		Gender *Gender `json:"g"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":"male"}`), &dPtr))
		assert.Equal(t, genderPtr(GenderMale), dPtr.Gender)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":null}`), &dPtr))
		assert.Equal(t, (*Gender)(nil), dPtr.Gender)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"g":"invalid"}`), &dPtr))
	}
}

func TestGender_String(t *testing.T) {
	tests := []struct {
		name string
		g    Gender
		want string
	}{
		{"male", GenderMale, "Male"},
		{"female", GenderFemale, "Female"},
		{"other", GenderOther, "Other"},
		{"prefers_not_to_say", GenderPreferNotToSay, "Prefer not to say"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestGender_MarshalText(t *testing.T) {
	{
		got, err := GenderMale.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "male", string(got))
	}
	{
		got, err := GenderFemale.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "female", string(got))
	}
	{
		got, err := GenderOther.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "other", string(got))
	}
	{
		got, err := GenderPreferNotToSay.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "prefers_not_to_say", string(got))
	}
}

func TestGender_UnmarshalText(t *testing.T) {
	{
		var g = new(Gender)
		assert.NoError(t, g.UnmarshalText([]byte("male")))
		assert.Equal(t, GenderMale, *g)
	}
	{
		var g = new(Gender)
		assert.NoError(t, g.UnmarshalText([]byte("female")))
		assert.Equal(t, GenderFemale, *g)
	}
	{
		var g = new(Gender)
		assert.NoError(t, g.UnmarshalText([]byte("other")))
		assert.Equal(t, GenderOther, *g)
	}
	{
		var g = new(Gender)
		assert.NoError(t, g.UnmarshalText([]byte("prefers_not_to_say")))
		assert.Equal(t, GenderPreferNotToSay, *g)
	}
	{
		var g = new(Gender)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}
}

func TestParseGender(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    Gender
		wantErr bool
	}{
		{"male", "male", GenderMale, false},
		{"female", "female", GenderFemale, false},
		{"other", "other", GenderOther, false},
		{"prefer_not_to_say", "prefers_not_to_say", GenderPreferNotToSay, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGender(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllGenders(t *testing.T) {
	assert.ElementsMatch(t, []Gender{
		GenderMale,
		GenderFemale,
		GenderOther,
		GenderPreferNotToSay,
	}, AllGenders().Items())
}
