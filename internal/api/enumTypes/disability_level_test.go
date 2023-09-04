package enumTypes

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func disabilityLevelPtr(g DisabilityLevel) *DisabilityLevel {
	return &g
}

func TestDisabilityLevel_MarshalJSON(t *testing.T) {
	type dummy struct {
		DisabilityLevel DisabilityLevel `json:"d"`
	}
	type dummyPtr struct {
		DisabilityLevel *DisabilityLevel `json:"d"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{DisabilityLevel: DisabilityLevelMild})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"d":"mild"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{DisabilityLevel: disabilityLevelPtr(DisabilityLevelSevere)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"d":"severe"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"d":null}`), jsonBytes)
	}
}

func TestDisabilityLevel_UnmarshalJSON(t *testing.T) {

	type dummyPtr struct {
		DisabilityLevel *DisabilityLevel `json:"d"`
	}
	tests := []struct {
		name    string
		bytes   []byte
		want    DisabilityLevel
		wantErr bool
	}{
		{"severe", []byte(`{"d":"severe"}`), DisabilityLevelSevere, false},
		{"none", []byte(`{"d":"none"}`), DisabilityLevelNone, false},
		{"null", []byte(`{"d":"null"}`), "", true},
		{"invalid", []byte(`{"d":"invalid"}`), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dPtr dummyPtr
			if tt.wantErr {
				assert.Error(t, json.Unmarshal(tt.bytes, &dPtr))
			} else {
				assert.NoError(t, json.Unmarshal(tt.bytes, &dPtr))
				assert.Equal(t, disabilityLevelPtr(tt.want), dPtr.DisabilityLevel)
			}
		})
	}

}

func TestDisabilityLevel_String(t *testing.T) {
	tests := []struct {
		name string
		g    DisabilityLevel
		want string
	}{
		{"none", DisabilityLevelNone, "None"},
		{"mild", DisabilityLevelMild, "Mild"},
		{"moderate", DisabilityLevelModerate, "Moderate"},
		{"severe", DisabilityLevelSevere, "Severe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestDisabilityLevel_MarshalText(t *testing.T) {
	{
		got, err := DisabilityLevelNone.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "none", string(got))
	}
	{
		got, err := DisabilityLevelMild.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "mild", string(got))
	}
	{
		got, err := DisabilityLevelModerate.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "moderate", string(got))
	}
	{
		got, err := DisabilityLevelSevere.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "severe", string(got))
	}
}

func TestDisabilityLevel_UnmarshalText(t *testing.T) {
	{
		var g = new(DisabilityLevel)
		assert.NoError(t, g.UnmarshalText([]byte("none")))
		assert.Equal(t, DisabilityLevelNone, *g)
	}
	{
		var g = new(DisabilityLevel)
		assert.NoError(t, g.UnmarshalText([]byte("mild")))
		assert.Equal(t, DisabilityLevelMild, *g)
	}
	{
		var g = new(DisabilityLevel)
		assert.NoError(t, g.UnmarshalText([]byte("moderate")))
		assert.Equal(t, DisabilityLevelModerate, *g)
	}
	{
		var g = new(DisabilityLevel)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}
}

func TestParseDisabilityLevel(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    DisabilityLevel
		wantErr bool
	}{
		{"none", "none", DisabilityLevelNone, false},
		{"mild", "mild", DisabilityLevelMild, false},
		{"moderate", "moderate", DisabilityLevelModerate, false},
		{"severe", "severe", DisabilityLevelSevere, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDisabilityLevel(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllDisabilityLevels(t *testing.T) {
	assert.ElementsMatch(t, []DisabilityLevel{
		DisabilityLevelNone,
		DisabilityLevelMild,
		DisabilityLevelModerate,
		DisabilityLevelSevere,
	}, AllDisabilityLevels().Items())
}
