package enumTypes

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func serviceCCPtr(g ServiceCC) *ServiceCC {
	return &g
}

func TestServiceCC_MarshalJSON(t *testing.T) {
	type dummy struct {
		ServiceCC ServiceCC `json:"g"`
	}
	type dummyPtr struct {
		ServiceCC *ServiceCC `json:"g"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{ServiceCC: ServiceCCCVA})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"cva"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{ServiceCC: serviceCCPtr(ServiceCCCVA)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"cva"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":null}`), jsonBytes)
	}
}

func TestServiceCC_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		ServiceCC ServiceCC `json:"g"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"g":"cva"}`), &d))
	assert.Equal(t, ServiceCCCVA, d.ServiceCC)

	type dummyPtr struct {
		ServiceCC *ServiceCC `json:"g"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":"icla"}`), &dPtr))
		assert.Equal(t, serviceCCPtr(ServiceCCICLA), dPtr.ServiceCC)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":null}`), &dPtr))
		assert.Equal(t, (*ServiceCC)(nil), dPtr.ServiceCC)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"g":"invalid"}`), &dPtr))
	}
}

func TestServiceCC_String(t *testing.T) {
	tests := []struct {
		name string
		g    ServiceCC
		want string
	}{
		{"cva", ServiceCCCVA, "option_service_cva"},
		{"icla", ServiceCCICLA, "option_service_icla"},
		{"education", ServiceCCEducation, "option_service_education"},
		{"lfs", ServiceCCLFS, "option_service_lfs"},
		{"protection", ServiceCCProtection, "option_service_protection"},
		{"shelter_and_settlements", ServiceCCShelter, "option_service_shelter"},
		{"wash", ServiceCCWash, "option_service_wash"},
		{"other", ServiceCCOther, "option_other"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestServiceCC_MarshalText(t *testing.T) {
	{
		got, err := ServiceCCCVA.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "cva", string(got))
	}
	{
		got, err := ServiceCCICLA.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "icla", string(got))
	}
	{
		got, err := ServiceCCEducation.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "education", string(got))
	}
	{
		got, err := ServiceCCLFS.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "lfs", string(got))
	}
	{
		got, err := ServiceCCProtection.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "protection", string(got))
	}
	{
		got, err := ServiceCCShelter.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "shelter_and_settlements", string(got))
	}
	{
		got, err := ServiceCCWash.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "wash", string(got))
	}
	{
		got, err := ServiceCCOther.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "other", string(got))
	}
}

func TestServiceCC_UnmarshalText(t *testing.T) {
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("cva")))
		assert.Equal(t, ServiceCCCVA, *g)
	}
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("icla")))
		assert.Equal(t, ServiceCCICLA, *g)
	}
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("education")))
		assert.Equal(t, ServiceCCEducation, *g)
	}
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("protection")))
		assert.Equal(t, ServiceCCProtection, *g)
	}
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("shelter_and_settlements")))
		assert.Equal(t, ServiceCCShelter, *g)
	}
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("wash")))
		assert.Equal(t, ServiceCCWash, *g)
	}
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("lfs")))
		assert.Equal(t, ServiceCCLFS, *g)
	}
	{
		var g = new(ServiceCC)
		assert.NoError(t, g.UnmarshalText([]byte("other")))
		assert.Equal(t, ServiceCCOther, *g)
	}
	{
		var g = new(ServiceCC)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}

}

func TestParseServiceCC(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    ServiceCC
		wantErr bool
	}{
		{"cva", "cva", ServiceCCCVA, false},
		{"icla", "icla", ServiceCCICLA, false},
		{"education", "education", ServiceCCEducation, false},
		{"protection", "protection", ServiceCCProtection, false},
		{"shelter_and_settlements", "shelter_and_settlements", ServiceCCShelter, false},
		{"wash", "wash", ServiceCCWash, false},
		{"lfs", "lfs", ServiceCCLFS, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseServiceCC(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllServiceCCs(t *testing.T) {
	assert.ElementsMatch(t, []ServiceCC{
		ServiceCCICLA,
		ServiceCCEducation,
		ServiceCCCVA,
		ServiceCCLFS,
		ServiceCCProtection,
		ServiceCCShelter,
		ServiceCCWash,
		ServiceCCOther,
	}, AllServiceCCs().Items())
}
