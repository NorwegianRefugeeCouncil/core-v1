package enumTypes

import (
	"encoding/json"
	"github.com/nrc-no/notcore/internal/locales"
	"testing"

	"github.com/stretchr/testify/assert"
)

func displacementStatusPtr(g DisplacementStatus) *DisplacementStatus {
	return &g
}

func TestDisplacementStatus_MarshalJSON(t *testing.T) {
	type dummy struct {
		DisplacementStatus DisplacementStatus `json:"g"`
	}
	type dummyPtr struct {
		DisplacementStatus *DisplacementStatus `json:"g"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{DisplacementStatus: DisplacementStatusRefugee})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"refugee"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{DisplacementStatus: displacementStatusPtr(DisplacementStatusRefugee)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"refugee"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":null}`), jsonBytes)
	}
}

func TestDisplacementStatus_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		DisplacementStatus DisplacementStatus `json:"g"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"g":"refugee"}`), &d))
	assert.Equal(t, DisplacementStatusRefugee, d.DisplacementStatus)

	type dummyPtr struct {
		DisplacementStatus *DisplacementStatus `json:"g"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":"idp"}`), &dPtr))
		assert.Equal(t, displacementStatusPtr(DisplacementStatusIDP), dPtr.DisplacementStatus)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":null}`), &dPtr))
		assert.Equal(t, (*DisplacementStatus)(nil), dPtr.DisplacementStatus)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"g":"invalid"}`), &dPtr))
	}
}

func TestDisplacementStatus_String(t *testing.T) {
	tr := locales.GetTranslator()
	tests := []struct {
		name string
		g    DisplacementStatus
		want string
	}{
		{"refugee", DisplacementStatusRefugee, tr("option_displacement_status_refugee")},
		{"idp", DisplacementStatusIDP, tr("option_displacement_status_idp")},
		{"host_community", DisplacementStatusHostCommunity, tr("option_displacement_status_host_community")},
		{"returnee", DisplacementStatusReturnee, tr("option_displacement_status_returnee")},
		{"asylum_seeker", DisplacementStatusAsylumSeeker, tr("option_displacement_status_asylum_seeker")},
		{"non_displaced", DisplacementStatusNonDisplaced, tr("option_displacement_status_non_displaced")},
		{"other", DisplacementStatusOther, tr("option_other")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestDisplacementStatus_MarshalText(t *testing.T) {
	{
		got, err := DisplacementStatusRefugee.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "refugee", string(got))
	}
	{
		got, err := DisplacementStatusIDP.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "idp", string(got))
	}
	{
		got, err := DisplacementStatusHostCommunity.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "host_community", string(got))
	}
	{
		got, err := DisplacementStatusReturnee.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "returnee", string(got))
	}
	{
		got, err := DisplacementStatusAsylumSeeker.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "asylum_seeker", string(got))
	}
	{
		got, err := DisplacementStatusNonDisplaced.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "non_displaced", string(got))
	}
	{
		got, err := DisplacementStatusOther.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "other", string(got))
	}
}

func TestDisplacementStatus_UnmarshalText(t *testing.T) {
	{
		var g = new(DisplacementStatus)
		assert.NoError(t, g.UnmarshalText([]byte("refugee")))
		assert.Equal(t, DisplacementStatusRefugee, *g)
	}
	{
		var g = new(DisplacementStatus)
		assert.NoError(t, g.UnmarshalText([]byte("idp")))
		assert.Equal(t, DisplacementStatusIDP, *g)
	}
	{
		var g = new(DisplacementStatus)
		assert.NoError(t, g.UnmarshalText([]byte("host_community")))
		assert.Equal(t, DisplacementStatusHostCommunity, *g)
	}
	{
		var g = new(DisplacementStatus)
		assert.NoError(t, g.UnmarshalText([]byte("asylum_seeker")))
		assert.Equal(t, DisplacementStatusAsylumSeeker, *g)
	}
	{
		var g = new(DisplacementStatus)
		assert.NoError(t, g.UnmarshalText([]byte("non_displaced")))
		assert.Equal(t, DisplacementStatusNonDisplaced, *g)
	}
	{
		var g = new(DisplacementStatus)
		assert.NoError(t, g.UnmarshalText([]byte("returnee")))
		assert.Equal(t, DisplacementStatusReturnee, *g)
	}
	{
		var g = new(DisplacementStatus)
		assert.NoError(t, g.UnmarshalText([]byte("other")))
		assert.Equal(t, DisplacementStatusOther, *g)
	}
	{
		var g = new(DisplacementStatus)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}

}

func TestParseDisplacementStatus(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    DisplacementStatus
		wantErr bool
	}{
		{"refugee", "refugee", DisplacementStatusRefugee, false},
		{"idp", "idp", DisplacementStatusIDP, false},
		{"host_community", "host_community", DisplacementStatusHostCommunity, false},
		{"asylum_seeker", "asylum_seeker", DisplacementStatusAsylumSeeker, false},
		{"non_displaced", "non_displaced", DisplacementStatusNonDisplaced, false},
		{"returnee", "returnee", DisplacementStatusReturnee, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDisplacementStatus(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllDisplacementStatuss(t *testing.T) {
	assert.ElementsMatch(t, []DisplacementStatus{
		DisplacementStatusIDP,
		DisplacementStatusHostCommunity,
		DisplacementStatusRefugee,
		DisplacementStatusReturnee,
		DisplacementStatusAsylumSeeker,
		DisplacementStatusNonDisplaced,
		DisplacementStatusOther,
	}, AllDisplacementStatuses().Items())
}
