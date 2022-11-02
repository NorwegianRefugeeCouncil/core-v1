package views

import (
	"path"
	"strconv"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/views/forms"
)

type IndividualForm struct {
	*forms.Form
}

func NewIndividualForm(i *api.Individual) *IndividualForm {

	isNew := i.ID == ""

	impairmentOptions := []forms.SelectInputFieldOption{
		{
			Value: "",
			Label: "None",
		}, {
			Value: "mild",
			Label: "Mild",
		}, {
			Value: "moderate",
			Label: "Moderate",
		}, {
			Value: "severe",
			Label: "Severe",
		},
	}

	var birthDate string
	if i.BirthDate != nil {
		birthDate = i.BirthDate.Format("2006-01-02")
	}

	idField := &forms.IDField{
		Name:        "id",
		DisplayName: "ID",
		Value:       i.ID,
		QRCodeURL:   path.Join("/countries", i.CountryID, "individuals", i.ID),
	}

	fullNameField := &forms.TextInputField{
		Name:        "fullName",
		DisplayName: "Full Name",
		Value:       i.FullName,
	}

	preferredNameField := &forms.TextInputField{
		Name:        "preferredName",
		DisplayName: "Preferred Name",
		Value:       i.PreferredName,
	}

	genderOptions := []forms.SelectInputFieldOption{
		{
			Value: "male",
			Label: "Male",
		}, {
			Value: "female",
			Label: "Female",
		}, {
			Value: "other",
			Label: "Other",
		}, {
			Value: "prefers_not_to_say",
			Label: "Prefer not to say",
		},
	}

	if isNew {
		genderOptions = append([]forms.SelectInputFieldOption{
			{
				Value: "",
				Label: "",
			},
		}, genderOptions...)
	}

	genderField := &forms.SelectInputField{
		Name:        "gender",
		DisplayName: "Gender",
		Value:       i.Gender,
		Required:    true,
		Options:     genderOptions,
	}

	birthDateField := &forms.DateInputField{
		Name:        "birthDate",
		DisplayName: "Birth Date",
		Value:       birthDate,
	}

	isMinorField := &forms.CheckboxInputField{
		Name:        "isMinor",
		DisplayName: "Is Minor",
		Value:       strconv.FormatBool(i.IsMinor),
	}

	displacementStatusOptions := []forms.SelectInputFieldOption{
		{
			Value: "refugee",
			Label: "Refugee",
		}, {
			Value: "idp",
			Label: "Internally Displaced Person",
		}, {
			Value: "hostCommunity",
			Label: "Host Community",
		},
	}

	if isNew {
		displacementStatusOptions = append([]forms.SelectInputFieldOption{
			{
				Value: "",
				Label: "",
			},
		}, displacementStatusOptions...)
	}

	displacementStatusField := &forms.SelectInputField{
		Options:     displacementStatusOptions,
		Name:        "displacementStatus",
		DisplayName: "Displacement Status",
		Value:       i.DisplacementStatus,
	}

	emailField := &forms.TextInputField{
		Name:        "email",
		DisplayName: "Email",
		Value:       i.Email,
		Required:    true,
	}

	phoneNumberField := &forms.TextInputField{
		Name:        "phoneNumber",
		DisplayName: "Phone",
		Value:       i.PhoneNumber,
	}

	addressField := &forms.TextAreaInputField{
		Name:        "address",
		DisplayName: "Address",
		Value:       i.Address,
	}

	presentsProtectionConcernsField := &forms.CheckboxInputField{
		Name:        "presentsProtectionConcerns",
		DisplayName: "Presents Protection Concerns",
		Value:       strconv.FormatBool(i.PresentsProtectionConcerns),
	}

	physicalImpairmentField := &forms.SelectInputField{
		Name:        "physicalImpairment",
		DisplayName: "Physical Impairment",
		Value:       i.PhysicalImpairment,
		Options:     impairmentOptions,
	}

	mentalImpairmentField := &forms.SelectInputField{
		Name:        "mentalImpairment",
		DisplayName: "Mental Impairment",
		Value:       i.MentalImpairment,
		Required:    false,
		Options:     impairmentOptions,
	}

	sensoryImpairmentField := &forms.SelectInputField{
		Name:        "sensoryImpairment",
		DisplayName: "Sensory Impairment",
		Value:       i.SensoryImpairment,
		Required:    false,
		Options:     impairmentOptions,
	}

	personalInfoFields := []forms.Field{
		fullNameField,
		preferredNameField,
		genderField,
		birthDateField,
		isMinorField,
		displacementStatusField,
	}
	if !isNew {
		personalInfoFields = append([]forms.Field{idField}, personalInfoFields...)
	}

	contactInfoFields := []forms.Field{
		emailField,
		phoneNumberField,
		addressField,
	}

	protectionFields := []forms.Field{
		presentsProtectionConcernsField,
	}

	disabilityFields := []forms.Field{
		physicalImpairmentField,
		mentalImpairmentField,
		sensoryImpairmentField,
	}

	personalInfoSection := &forms.FormSection{
		Title:       "Personal Information",
		Fields:      personalInfoFields,
		Collapsible: true,
	}

	contactInfoSection := &forms.FormSection{
		Title:       "Contact Information",
		Fields:      contactInfoFields,
		Collapsible: true,
	}

	protectionSection := &forms.FormSection{
		Title:       "Protection Concerns",
		Fields:      protectionFields,
		Collapsible: true,
	}

	disabilitySection := &forms.FormSection{
		Title:       "Disability",
		Fields:      disabilityFields,
		Collapsible: true,
	}

	formSections := []*forms.FormSection{
		personalInfoSection,
		contactInfoSection,
		protectionSection,
		disabilitySection,
	}

	action := "/countries/" + i.CountryID + "/individuals/"
	if len(i.ID) != 0 {
		action += i.ID
	} else {
		action += "new"
	}
	f := forms.Form{
		Sections: formSections,
	}

	return &IndividualForm{
		Form: &f,
	}
}
