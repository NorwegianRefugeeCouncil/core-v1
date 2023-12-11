package enumTypes

import (
	"encoding/json"
	"github.com/nrc-no/notcore/internal/locales"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sexPtr(g Sex) *Sex {
	return &g
}

func TestSex_MarshalJSON(t *testing.T) {
	type dummy struct {
		Sex Sex `json:"g"`
	}
	type dummyPtr struct {
		Sex *Sex `json:"g"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{Sex: SexMale})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"male"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{Sex: sexPtr(SexMale)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"male"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":null}`), jsonBytes)
	}
}

func TestSex_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		Sex Sex `json:"g"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"g":"male"}`), &d))
	assert.Equal(t, SexMale, d.Sex)

	type dummyPtr struct {
		Sex *Sex `json:"g"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":"male"}`), &dPtr))
		assert.Equal(t, sexPtr(SexMale), dPtr.Sex)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":null}`), &dPtr))
		assert.Equal(t, (*Sex)(nil), dPtr.Sex)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"g":"invalid"}`), &dPtr))
	}
}

func TestSex_String(t *testing.T) {
	tr := locales.GetTranslator()
	tests := []struct {
		name string
		g    Sex
		want string
	}{
		{"male", SexMale, tr("option_sex_male")},
		{"female", SexFemale, tr("option_sex_female")},
		{"other", SexOther, tr("option_other")},
		{"prefers_not_to_say", SexPreferNotToSay, tr("option_sex_prefers_not_to_say")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestSex_MarshalText(t *testing.T) {
	{
		got, err := SexMale.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "male", string(got))
	}
	{
		got, err := SexFemale.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "female", string(got))
	}
	{
		got, err := SexOther.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "other", string(got))
	}
	{
		got, err := SexPreferNotToSay.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "prefers_not_to_say", string(got))
	}
}

func TestSex_UnmarshalText(t *testing.T) {
	{
		var g = new(Sex)
		assert.NoError(t, g.UnmarshalText([]byte("male")))
		assert.Equal(t, SexMale, *g)
	}
	{
		var g = new(Sex)
		assert.NoError(t, g.UnmarshalText([]byte("female")))
		assert.Equal(t, SexFemale, *g)
	}
	{
		var g = new(Sex)
		assert.NoError(t, g.UnmarshalText([]byte("other")))
		assert.Equal(t, SexOther, *g)
	}
	{
		var g = new(Sex)
		assert.NoError(t, g.UnmarshalText([]byte("prefers_not_to_say")))
		assert.Equal(t, SexPreferNotToSay, *g)
	}
	{
		var g = new(Sex)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}
}

func TestParseSex(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    Sex
		wantErr bool
	}{
		{"male", "male", SexMale, false},
		{"female", "female", SexFemale, false},
		{"other", "other", SexOther, false},
		{"prefer_not_to_say", "prefers_not_to_say", SexPreferNotToSay, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSex(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllSexes(t *testing.T) {
	assert.ElementsMatch(t, []Sex{
		SexMale,
		SexFemale,
		SexOther,
		SexPreferNotToSay,
	}, AllSexes().Items())
}
