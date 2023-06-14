package enumTypes

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func engagementContextPtr(g EngagementContext) *EngagementContext {
	return &g
}

func TestEngagementContext_MarshalJSON(t *testing.T) {
	type dummy struct {
		EngagementContext EngagementContext `json:"g"`
	}
	type dummyPtr struct {
		EngagementContext *EngagementContext `json:"g"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{EngagementContext: EngagementContextHouseVisit})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"houseVisit"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{EngagementContext: engagementContextPtr(EngagementContextHouseVisit)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"houseVisit"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":null}`), jsonBytes)
	}
}

func TestEngagementContext_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		EngagementContext EngagementContext `json:"g"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"g":"houseVisit"}`), &d))
	assert.Equal(t, EngagementContextHouseVisit, d.EngagementContext)

	type dummyPtr struct {
		EngagementContext *EngagementContext `json:"g"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":"fieldActivity"}`), &dPtr))
		assert.Equal(t, engagementContextPtr(EngagementContextFieldActivity), dPtr.EngagementContext)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":null}`), &dPtr))
		assert.Equal(t, (*EngagementContext)(nil), dPtr.EngagementContext)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"g":"invalid"}`), &dPtr))
	}
}

func TestEngagementContext_String(t *testing.T) {
	tests := []struct {
		name string
		g    EngagementContext
		want string
	}{
		{"houseVisit", EngagementContextHouseVisit, "option_engagement_context_house_visit"},
		{"fieldActivity", EngagementContextFieldActivity, "option_engagement_context_field_activity"},
		{"inOffice", EngagementContextInOffice, "option_engagement_context_in_office"},
		{"referred", EngagementContextReferred, "option_engagement_context_referred"},
		{"remoteChannels", EngagementContextRemoteChannels, "option_engagement_context_remote_channels"},
		{"other", EngagementContextOther, "option_other"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestEngagementContext_MarshalText(t *testing.T) {
	{
		got, err := EngagementContextHouseVisit.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "houseVisit", string(got))
	}
	{
		got, err := EngagementContextFieldActivity.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "fieldActivity", string(got))
	}
	{
		got, err := EngagementContextInOffice.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "inOffice", string(got))
	}
	{
		got, err := EngagementContextReferred.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "referred", string(got))
	}
	{
		got, err := EngagementContextRemoteChannels.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "remoteChannels", string(got))
	}
	{
		got, err := EngagementContextOther.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "other", string(got))
	}
}

func TestEngagementContext_UnmarshalText(t *testing.T) {
	{
		var g = new(EngagementContext)
		assert.NoError(t, g.UnmarshalText([]byte("houseVisit")))
		assert.Equal(t, EngagementContextHouseVisit, *g)
	}
	{
		var g = new(EngagementContext)
		assert.NoError(t, g.UnmarshalText([]byte("fieldActivity")))
		assert.Equal(t, EngagementContextFieldActivity, *g)
	}
	{
		var g = new(EngagementContext)
		assert.NoError(t, g.UnmarshalText([]byte("inOffice")))
		assert.Equal(t, EngagementContextInOffice, *g)
	}
	{
		var g = new(EngagementContext)
		assert.NoError(t, g.UnmarshalText([]byte("remoteChannels")))
		assert.Equal(t, EngagementContextRemoteChannels, *g)
	}
	{
		var g = new(EngagementContext)
		assert.NoError(t, g.UnmarshalText([]byte("referred")))
		assert.Equal(t, EngagementContextReferred, *g)
	}
	{
		var g = new(EngagementContext)
		assert.NoError(t, g.UnmarshalText([]byte("other")))
		assert.Equal(t, EngagementContextOther, *g)
	}
	{
		var g = new(EngagementContext)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}

}

func TestParseEngagementContext(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    EngagementContext
		wantErr bool
	}{
		{"houseVisit", "houseVisit", EngagementContextHouseVisit, false},
		{"fieldActivity", "fieldActivity", EngagementContextFieldActivity, false},
		{"inOffice", "inOffice", EngagementContextInOffice, false},
		{"remoteChannels", "remoteChannels", EngagementContextRemoteChannels, false},
		{"referred", "referred", EngagementContextReferred, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseEngagementContext(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllEngagementContexts(t *testing.T) {
	assert.ElementsMatch(t, []EngagementContext{
		EngagementContextFieldActivity,
		EngagementContextInOffice,
		EngagementContextHouseVisit,
		EngagementContextReferred,
		EngagementContextRemoteChannels,
		EngagementContextOther,
	}, AllEngagementContexts().Items())
}
