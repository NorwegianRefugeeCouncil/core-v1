package enumTypes

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func contactMethodPtr(g ContactMethod) *ContactMethod {
	return &g
}

func TestContactMethod_MarshalJSON(t *testing.T) {
	type dummy struct {
		ContactMethod ContactMethod `json:"g"`
	}
	type dummyPtr struct {
		ContactMethod *ContactMethod `json:"g"`
	}
	{
		jsonBytes, err := json.Marshal(dummy{ContactMethod: ContactMethodEmail})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"email"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{ContactMethod: contactMethodPtr(ContactMethodEmail)})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":"email"}`), jsonBytes)
	}
	{
		jsonBytes, err := json.Marshal(dummyPtr{})
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"g":null}`), jsonBytes)
	}
}

func TestContactMethod_UnmarshalJSON(t *testing.T) {
	type dummy struct {
		ContactMethod ContactMethod `json:"g"`
	}
	var d dummy
	assert.NoError(t, json.Unmarshal([]byte(`{"g":"email"}`), &d))
	assert.Equal(t, ContactMethodEmail, d.ContactMethod)

	type dummyPtr struct {
		ContactMethod *ContactMethod `json:"g"`
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":"whatsapp"}`), &dPtr))
		assert.Equal(t, contactMethodPtr(ContactMethodWhatsapp), dPtr.ContactMethod)
	}
	{
		var dPtr dummyPtr
		assert.NoError(t, json.Unmarshal([]byte(`{"g":null}`), &dPtr))
		assert.Equal(t, (*ContactMethod)(nil), dPtr.ContactMethod)
	}
	{
		var dPtr dummyPtr
		assert.Error(t, json.Unmarshal([]byte(`{"g":"invalid"}`), &dPtr))
	}
}

func TestContactMethod_String(t *testing.T) {
	tests := []struct {
		name string
		g    ContactMethod
		want string
	}{
		{"email", ContactMethodEmail, "option_contact_method_email"},
		{"whatsapp", ContactMethodWhatsapp, "option_contact_method_whatsapp"},
		{"phone", ContactMethodPhone, "option_contact_method_phone"},
		{"visit", ContactMethodVisit, "option_contact_method_visit"},
		{"other", ContactMethodOther, "option_other"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.String())
		})
	}
}

func TestContactMethod_MarshalText(t *testing.T) {
	{
		got, err := ContactMethodEmail.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "email", string(got))
	}
	{
		got, err := ContactMethodWhatsapp.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "whatsapp", string(got))
	}
	{
		got, err := ContactMethodPhone.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "phone", string(got))
	}
	{
		got, err := ContactMethodVisit.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "visit", string(got))
	}
	{
		got, err := ContactMethodOther.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, "other", string(got))
	}
}

func TestContactMethod_UnmarshalText(t *testing.T) {
	{
		var g = new(ContactMethod)
		assert.NoError(t, g.UnmarshalText([]byte("email")))
		assert.Equal(t, ContactMethodEmail, *g)
	}
	{
		var g = new(ContactMethod)
		assert.NoError(t, g.UnmarshalText([]byte("whatsapp")))
		assert.Equal(t, ContactMethodWhatsapp, *g)
	}
	{
		var g = new(ContactMethod)
		assert.NoError(t, g.UnmarshalText([]byte("phone")))
		assert.Equal(t, ContactMethodPhone, *g)
	}
	{
		var g = new(ContactMethod)
		assert.NoError(t, g.UnmarshalText([]byte("visit")))
		assert.Equal(t, ContactMethodVisit, *g)
	}
	{
		var g = new(ContactMethod)
		assert.NoError(t, g.UnmarshalText([]byte("other")))
		assert.Equal(t, ContactMethodOther, *g)
	}
	{
		var g = new(ContactMethod)
		assert.Error(t, g.UnmarshalText([]byte("invalid")))
	}

}

func TestParseContactMethod(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    ContactMethod
		wantErr bool
	}{
		{"email", "email", ContactMethodEmail, false},
		{"whatsapp", "whatsapp", ContactMethodWhatsapp, false},
		{"phone", "phone", ContactMethodPhone, false},
		{"visit", "visit", ContactMethodVisit, false},
		{"invalid", "invalid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseContactMethod(tt.str)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllContactMethods(t *testing.T) {
	assert.ElementsMatch(t, []ContactMethod{
		ContactMethodWhatsapp,
		ContactMethodPhone,
		ContactMethodEmail,
		ContactMethodVisit,
		ContactMethodOther,
	}, AllContactMethods().Items())
}
