package enumTypes

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func identificationTypePtr(g IdentificationType) *IdentificationType {
	return &g
}

func TestIdentificationType_MarshalJSON(t *testing.T) {
	type dummy struct {
		IdentificationType IdentificationType `json:"g"`
	}
	type dummyPtr struct {
		IdentificationType *IdentificationType `json:"g"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{IdentificationType: IdentificationTypePassport})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"passport"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{IdentificationType: identificationTypePtr(IdentificationTypePassport)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"passport"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":null}`), jsonBytes)
	}
}

func TestIdentificationType_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		IdentificationType IdentificationType `json:"g"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"g":"passport"}`), &d))
	assert.Equal(t, IdentificationTypePassport, d.IdentificationType)

	type dummyPtr struct {
		IdentificationType *IdentificationType `json:"g"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":"unhcr_id"}`), &dPtr))
		assert.Equal(t, identificationTypePtr(IdentificationTypeUNHCR), dPtr.IdentificationType)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":null}`), &dPtr))
		assert.Equal(t, (*IdentificationType)(nil), dPtr.IdentificationType)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"g":"invalid"}`), &dPtr))
	}
}

func TestIdentificationType_String(t *testing.T) {
	tests := []struct {
		name string
		g    IdentificationType
		want string
	}{
		{"passport", IdentificationTypePassport, "option_identification_type_passport"},
		{"unhcr_id", IdentificationTypeUNHCR, "option_identification_type_unhcr"},
		{"national_id", IdentificationTypeNational, "option_identification_type_national_id"},
		{"other", IdentificationTypeOther, "option_other"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestIdentificationType_MarshalText(t *testing.T) {
	{
		got, err := IdentificationTypePassport.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "passport", string(got))
	}
	{
		got, err := IdentificationTypeUNHCR.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "unhcr_id", string(got))
	}
	{
		got, err := IdentificationTypeNational.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "national_id", string(got))
	}
	{
		got, err := IdentificationTypeOther.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "other", string(got))
	}
}

func TestIdentificationType_UnmarshalText(t *testing.T) {
	{
		var g = new(IdentificationType)
		assert.NoError(t, g.UnmarshalText([]byte("passport")))
		assert.Equal(t, IdentificationTypePassport, *g)
	}
	{
		var g = new(IdentificationType)
		assert.NoError(t, g.UnmarshalText([]byte("unhcr_id")))
		assert.Equal(t, IdentificationTypeUNHCR, *g)
	}
	{
		var g = new(IdentificationType)
		assert.NoError(t, g.UnmarshalText([]byte("national_id")))
		assert.Equal(t, IdentificationTypeNational, *g)
	}
	{
		var g = new(IdentificationType)
		assert.NoError(t, g.UnmarshalText([]byte("other")))
		assert.Equal(t, IdentificationTypeOther, *g)
	}
	{
		var g = new(IdentificationType)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}

}

func TestParseIdentificationType(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    IdentificationType
		wantErr bool
	}{
		{"passport", "passport", IdentificationTypePassport, false},
		{"unhcr_id", "unhcr_id", IdentificationTypeUNHCR, false},
		{"national_id", "national_id", IdentificationTypeNational, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIdentificationType(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllIdentificationTypes(t *testing.T) {
	assert.ElementsMatch(t, []IdentificationType{
		IdentificationTypeUNHCR,
		IdentificationTypeNational,
		IdentificationTypePassport,
		IdentificationTypeOther,
	}, AllIdentificationTypes().Items())
}
