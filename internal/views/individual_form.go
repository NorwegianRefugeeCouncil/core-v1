package views

import (
	"fmt"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/pkg/views/forms"
)

type IndividualForm struct {
	*forms.Form
	individual            *api.Individual
	personalInfoSection   *forms.FormSection
	contactInfoSection    *forms.FormSection
	protectionSection     *forms.FormSection
	disabilitiesSection   *forms.FormSection
	dataCollectionSection *forms.FormSection
}

func NewIndividualForm(i *api.Individual) (*IndividualForm, error) {
	f := &IndividualForm{
		Form:       &forms.Form{},
		individual: i,
	}
	if err := f.build(); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *IndividualForm) build() error {
	type builderFuncs func() error

	runBuilderFunctions := func(builders ...builderFuncs) error {
		for _, builder := range builders {
			if err := builder(); err != nil {
				return err
			}
		}
		return nil
	}

	sectionBuilders := []builderFuncs{
		f.buildPersonalInfoSection,
		f.buildContactInfoSection,
		f.buildProtectionSection,
		f.buildDisabilitiesSection,
		f.buildDataColletionSection,
	}

	fieldBuilders := []builderFuncs{
		f.buildTitle,
		f.buildIdField,
		f.buildFullName,
		f.buildPreferredName,
		f.buildFirstName,
		f.buildMiddleName,
		f.buildLastName,
		f.buildPrefersToRemainAnonymous,
		f.buildSex,
		f.buildBirthDate,
		f.buildAge,
		f.buildIsMinor,
		f.buildNationality1,
		f.buildNationality2,
		f.buildIdentification1Type,
		f.buildIdentification1Other,
		f.buildIdentification1Number,
		f.buildIdentification2Type,
		f.buildIdentification2Other,
		f.buildIdentification2Number,
		f.buildIdentification3Type,
		f.buildIdentification3Other,
		f.buildIdentification3Number,
		f.buildInternalID,
		f.buildHouseholdID,
		f.buildIsHeadOfHousehold,
		f.buildIsFemaleHeadedHousehold,
		f.buildIsMinorHeadedHousehold,
		f.buildCommunityID,
		f.buildIsHeadOfCommunity,
		f.buildSpokenLanguage1,
		f.buildSpokenLanguage2,
		f.buildSpokenLanguage3,
		f.buildPreferredCommunicationLanguage,
		f.buildPhoneNumber1,
		f.buildPhoneNumber2,
		f.buildPhoneNumber3,
		f.buildEmailAddress1,
		f.buildEmailAddress2,
		f.buildEmailAddress3,
		f.buildAddress,
		f.buildPreferredMeansOfContact,
		f.buildContactInstructions,
		f.buildHasConsentedToRgpd,
		f.buildHasConsentedToReferral,
		f.buildPresentsProtectionConcerns,
		f.buildDisplacementStatus,
		f.buildDisplacementStatusComment,
		f.buildHasVisionDisability,
		f.buildVisionDisabilityLevel,
		f.buildHasHearingDisability,
		f.buildHearingDisabilityLevel,
		f.buildHasMobilityDisability,
		f.buildMobilityDisabilityLevel,
		f.buildHasCognitiveDisability,
		f.buildCognitiveDisabilityLevel,
		f.buildHasSelfCareDisability,
		f.buildSelfCareDisabilityLevel,
		f.buildHasCommunicationDisability,
		f.buildCommunicationDisabilityLevel,
		f.buildEngagementContext,
		f.buildComments,
		f.buildFreeField1,
		f.buildFreeField2,
		f.buildFreeField3,
		f.buildFreeField4,
		f.buildFreeField5,
		f.buildCollectionAgent,
		f.buildCollectionAgentTitle,
		f.buildCollectionDate,
		f.buildCollectionLocation1,
		f.buildCollectionLocation2,
		f.buildCollectionLocation3,
		f.buildCollectionOffice,
	}

	if err := runBuilderFunctions(sectionBuilders...); err != nil {
		return err
	}

	if err := runBuilderFunctions(fieldBuilders...); err != nil {
		return err
	}

	return nil
}

func (f *IndividualForm) buildTitle() error {
	if f.isNew() {
		f.Form.Title = "New Individual"
	} else if f.individual.FullName == "" {
		f.Form.Title = "Anonymous Individual"
	} else {
		f.Form.Title = f.individual.FullName
	}
	return nil
}

func (f *IndividualForm) buildPersonalInfoSection() error {
	f.personalInfoSection = &forms.FormSection{
		Title:       "Personal Information",
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   false,
	}
	f.Form.Sections = append(f.Form.Sections, f.personalInfoSection)
	return nil
}

func (f *IndividualForm) buildContactInfoSection() error {
	f.contactInfoSection = &forms.FormSection{
		Title:       "Contact Information",
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.contactInfoSection)
	return nil
}

func (f *IndividualForm) buildProtectionSection() error {
	f.protectionSection = &forms.FormSection{
		Title:       "Protection",
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.protectionSection)
	return nil
}

func (f *IndividualForm) buildDisabilitiesSection() error {
	f.disabilitiesSection = &forms.FormSection{
		Title:       "Disabilities",
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.disabilitiesSection)
	return nil
}

func (f *IndividualForm) buildDataColletionSection() error {
	f.dataCollectionSection = &forms.FormSection{
		Title:       "Data Collection",
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.dataCollectionSection)
	return nil
}

func (f *IndividualForm) buildIdField() error {
	if !f.isNew() {
		return buildField(&forms.IDField{
			Name:        "id",
			DisplayName: "ID",
			QRCodeURL:   fmt.Sprintf("/countries/%s/individuals/%s", f.individual.CountryID, f.individual.ID),
		}, f.personalInfoSection, f.individual.ID)
	}
	return nil
}

func (f *IndividualForm) isNew() bool {
	return len(f.individual.ID) == 0
}

func (f *IndividualForm) buildFullName() error {
	return buildField(&forms.TextInputField{
		Name:        "fullName",
		DisplayName: "Full Name",
	}, f.personalInfoSection, f.individual.FullName)
}

func (f *IndividualForm) buildPreferredName() error {
	return buildField(&forms.TextInputField{
		Name:        "preferredName",
		DisplayName: "Preferred Name",
		Value:       f.individual.PreferredName,
	}, f.personalInfoSection, f.individual.PreferredName)
}

func (f *IndividualForm) buildFirstName() error {
	return buildField(&forms.TextInputField{
		Name:        "firstName",
		DisplayName: "First Name",
	}, f.personalInfoSection, f.individual.FirstName)
}

func (f *IndividualForm) buildMiddleName() error {
	return buildField(&forms.TextInputField{
		Name:        "middleName",
		DisplayName: "Middle Name",
	}, f.personalInfoSection, f.individual.MiddleName)
}

func (f *IndividualForm) buildLastName() error {
	return buildField(&forms.TextInputField{
		Name:        "lastName",
		DisplayName: "Last Name",
	}, f.personalInfoSection, f.individual.LastName)
}

func (f *IndividualForm) buildPrefersToRemainAnonymous() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "prefersToRemainAnonymous",
		DisplayName: "Prefers to Remain Anonymous",
	}, f.personalInfoSection, f.individual.PrefersToRemainAnonymous)
}

func (f *IndividualForm) buildSex() error {
	sexOptions := getSexOptions()
	if f.isNew() {
		sexOptions = append([]forms.SelectInputFieldOption{
			{
				Value: "",
				Label: "",
			},
		}, sexOptions...)
	}
	return buildField(&forms.SelectInputField{
		Name:        "sex",
		DisplayName: "Sex",
		Options:     sexOptions,
		Codec:       &sexCodec{},
	}, f.personalInfoSection, f.individual.Sex)
}

func (f *IndividualForm) buildBirthDate() error {
	return buildField(&forms.DateInputField{
		Name:        "birthDate",
		DisplayName: "Birth Date",
	}, f.personalInfoSection, f.individual.BirthDate)
}

func (f *IndividualForm) buildAge() error {
	return buildField(&forms.NumberInputField{
		Name:        "age",
		DisplayName: "Age",
	}, f.personalInfoSection, f.individual.Age)
}

func (f *IndividualForm) buildIsMinor() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isMinor",
		DisplayName: "Is Minor",
	}, f.personalInfoSection, f.individual.IsMinor)
}

func (f *IndividualForm) buildNationality1() error {
	return buildField(&forms.SelectInputField{
		Name:        "nationality1",
		DisplayName: "Nationality 1",
		Options:     buildCountryOptions(),
	}, f.personalInfoSection, f.individual.Nationality1)
}

func (f *IndividualForm) buildNationality2() error {
	return buildField(&forms.SelectInputField{
		Name:        "nationality2",
		DisplayName: "Nationality 2",
		Options:     buildCountryOptions(),
	}, f.personalInfoSection, f.individual.Nationality2)
}

func (f *IndividualForm) buildIdentification1Type() error {
	return buildField(&forms.SelectInputField{
		Name:        "identificationType1",
		DisplayName: "Identification Type 1",
		Options:     getIdentificationTypeOptions(),
	}, f.personalInfoSection, f.individual.IdentificationType1)
}

func (f *IndividualForm) buildIdentification1Other() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "identificationTypeExplanation1",
		DisplayName: "If Other, please explain",
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation1)
}

func (f *IndividualForm) buildIdentification1Number() error {
	return buildField(&forms.TextInputField{
		Name:        "identificationNumber1",
		DisplayName: "Identification Number 1",
	}, f.personalInfoSection, f.individual.IdentificationNumber1)
}

func (f *IndividualForm) buildIdentification2Type() error {
	return buildField(&forms.SelectInputField{
		Name:        "identificationType2",
		DisplayName: "Identification Type 2",
		Options:     getIdentificationTypeOptions(),
	}, f.personalInfoSection, f.individual.IdentificationType2)
}

func (f *IndividualForm) buildIdentification2Other() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "identificationTypeExplanation2",
		DisplayName: "If Other, please explain",
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation2)
}

func (f *IndividualForm) buildIdentification2Number() error {
	return buildField(&forms.TextInputField{
		Name:        "identificationNumber2",
		DisplayName: "Identification Number 2",
	}, f.personalInfoSection, f.individual.IdentificationNumber2)
}
func (f *IndividualForm) buildIdentification3Type() error {
	return buildField(&forms.SelectInputField{
		Name:        "identificationType3",
		DisplayName: "Identification Type 3",
		Options:     getIdentificationTypeOptions(),
	}, f.personalInfoSection, f.individual.IdentificationType3)
}

func (f *IndividualForm) buildIdentification3Other() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "identificationTypeExplanation3",
		DisplayName: "If Other, please explain",
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation3)
}

func (f *IndividualForm) buildIdentification3Number() error {
	return buildField(&forms.TextInputField{
		Name:        "identificationNumber3",
		DisplayName: "Identification Number 3",
	}, f.personalInfoSection, f.individual.IdentificationNumber3)
}

func (f *IndividualForm) buildInternalID() error {
	return buildField(&forms.TextInputField{
		Name:        "internalId",
		DisplayName: "Internal ID",
	}, f.personalInfoSection, f.individual.InternalID)
}

func (f *IndividualForm) buildHouseholdID() error {
	return buildField(&forms.TextInputField{
		Name:        "householdId",
		DisplayName: "Household ID",
	}, f.personalInfoSection, f.individual.HouseholdID)
}

func (f *IndividualForm) buildIsHeadOfHousehold() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isHeadOfHousehold",
		DisplayName: "Is Head of Household",
	}, f.personalInfoSection, f.individual.IsHeadOfHousehold)
}

func (f *IndividualForm) buildIsFemaleHeadedHousehold() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isFemaleHeadedHousehold",
		DisplayName: "Is Female Headed Household",
	}, f.personalInfoSection, f.individual.IsFemaleHeadedHousehold)
}

func (f *IndividualForm) buildIsMinorHeadedHousehold() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isMinorHeadedHousehold",
		DisplayName: "Is Minor Headed Household",
	}, f.personalInfoSection, f.individual.IsMinorHeadedHousehold)
}

func (f *IndividualForm) buildCommunityID() error {
	return buildField(&forms.TextInputField{
		Name:        "communityId",
		DisplayName: "Community ID",
	}, f.personalInfoSection, f.individual.CommunityID)
}

func (f *IndividualForm) buildIsHeadOfCommunity() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isHeadOfCommunity",
		DisplayName: "Is Head of Community",
	}, f.personalInfoSection, f.individual.IsHeadOfCommunity)
}

func (f *IndividualForm) buildSpokenLanguage1() error {
	return buildField(&forms.SelectInputField{
		Name:        "spokenLanguage1",
		DisplayName: "Spoken Language 1",
		Options:     buildLanguageOptions(),
	}, f.personalInfoSection, f.individual.SpokenLanguage1)
}

func (f *IndividualForm) buildSpokenLanguage2() error {
	return buildField(&forms.SelectInputField{
		Name:        "spokenLanguage2",
		DisplayName: "Spoken Language 2",
		Options:     buildLanguageOptions(),
	}, f.personalInfoSection, f.individual.SpokenLanguage2)
}

func (f *IndividualForm) buildSpokenLanguage3() error {
	return buildField(&forms.SelectInputField{
		Name:        "spokenLanguage3",
		DisplayName: "Spoken Language 3",
		Options:     buildLanguageOptions(),
	}, f.personalInfoSection, f.individual.SpokenLanguage3)
}

func (f *IndividualForm) buildPreferredCommunicationLanguage() error {
	return buildField(&forms.SelectInputField{
		Name:        "preferredCommunicationLanguage",
		DisplayName: "Preferred Communication Language",
		Options:     buildLanguageOptions(),
	}, f.personalInfoSection, f.individual.PreferredCommunicationLanguage)
}

func (f *IndividualForm) buildPhoneNumber1() error {
	return buildField(&forms.TextInputField{
		Name:        "phoneNumber1",
		DisplayName: "Phone Number 1",
	}, f.contactInfoSection, f.individual.PhoneNumber1)
}

func (f *IndividualForm) buildPhoneNumber2() error {
	return buildField(&forms.TextInputField{
		Name:        "phoneNumber2",
		DisplayName: "Phone Number 2",
	}, f.contactInfoSection, f.individual.PhoneNumber2)
}

func (f *IndividualForm) buildPhoneNumber3() error {
	return buildField(&forms.TextInputField{
		Name:        "phoneNumber3",
		DisplayName: "Phone Number 3",
	}, f.contactInfoSection, f.individual.PhoneNumber3)
}

func (f *IndividualForm) buildEmailAddress1() error {
	return buildField(&forms.TextInputField{
		Name:        "email1",
		DisplayName: "Email Address 1",
	}, f.contactInfoSection, f.individual.Email1)
}

func (f *IndividualForm) buildEmailAddress2() error {
	return buildField(&forms.TextInputField{
		Name:        "email2",
		DisplayName: "Email Address 2",
	}, f.contactInfoSection, f.individual.Email2)
}

func (f *IndividualForm) buildEmailAddress3() error {
	return buildField(&forms.TextInputField{
		Name:        "email3",
		DisplayName: "Email Address 3",
	}, f.contactInfoSection, f.individual.Email3)
}

func (f *IndividualForm) buildAddress() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "address",
		DisplayName: "Residence Address",
	}, f.contactInfoSection, f.individual.Address)
}

func (f *IndividualForm) buildPreferredMeansOfContact() error {
	options := []forms.SelectInputFieldOption{
		{Label: "Phone", Value: "phone"},
		{Label: "WhatsApp", Value: "whatsapp"},
		{Label: "Email", Value: "email"},
		{Label: "Visit", Value: "visit"},
		{Label: "Other", Value: "other"},
	}
	if f.isNew() {
		options = append([]forms.SelectInputFieldOption{
			{Label: "Select Preferred Means of Contact", Value: ""},
		}, options...)
	}
	return buildField(&forms.SelectInputField{
		Name:        "preferredContactMethod",
		DisplayName: "Preferred contact method",
		Options:     options,
	}, f.contactInfoSection, f.individual.PreferredContactMethod)
}

func (f *IndividualForm) buildContactInstructions() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "preferredContactMethodComments",
		DisplayName: "Instructions for contact or other comments",
	}, f.contactInfoSection, f.individual.PreferredContactMethodComments)
}

func (f *IndividualForm) buildHasConsentedToRgpd() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasConsentedToRgpd",
		DisplayName: "Has the person consented to NRC using their data?",
	}, f.protectionSection, f.individual.HasConsentedToRGPD)
}

func (f *IndividualForm) buildHasConsentedToReferral() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasConsentedToReferral",
		DisplayName: "Has the person consented to NRC referring them to other service providers within or outside of NRC",
	}, f.protectionSection, f.individual.HasConsentedToRGPD)
}

func (f *IndividualForm) buildPresentsProtectionConcerns() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "presentsProtectionConcerns",
		DisplayName: "Presents protection concerns",
	}, f.protectionSection, f.individual.PresentsProtectionConcerns)
}

func (f *IndividualForm) buildDisplacementStatus() error {
	options := getDisplacementStatusOptions()
	if f.isNew() {
		options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	}
	return buildField(&forms.SelectInputField{
		Name:        "displacementStatus",
		DisplayName: "Displacement Status",
		Options:     options,
		Codec:       &displacementStatusCodec{},
	}, f.protectionSection, f.individual.DisplacementStatus)
}

func (f *IndividualForm) buildDisplacementStatusComment() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "displacementStatusComment",
		DisplayName: "If Other, please explain",
	}, f.protectionSection, f.individual.DisplacementStatusComment)
}

func (f *IndividualForm) buildHasVisionDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasVisionDisability",
		DisplayName: "Has Vision disability",
	}, f.disabilitiesSection, f.individual.HasVisionDisability)
}

func (f *IndividualForm) buildVisionDisabilityLevel() error {
	return buildField(&forms.SelectInputField{
		Name:        "visionDisabilityLevel",
		DisplayName: "Vision disability",
		Options:     getDisabilityLevels(),
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.VisionDisabilityLevel)
}

func (f *IndividualForm) buildHasHearingDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasHearingDisability",
		DisplayName: "Has Hearing Disability",
	}, f.disabilitiesSection, f.individual.HasHearingDisability)
}

func (f *IndividualForm) buildHearingDisabilityLevel() error {
	return buildField(&forms.SelectInputField{
		Name:        "hearingDisabilityLevel",
		DisplayName: "Hearing disability level",
		Options:     getDisabilityLevels(),
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.HearingDisabilityLevel)
}

func (f *IndividualForm) buildHasMobilityDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasMobilityDisability",
		DisplayName: "Has Mobility Disability",
	}, f.disabilitiesSection, f.individual.HasMobilityDisability)
}

func (f *IndividualForm) buildMobilityDisabilityLevel() error {
	return buildField(&forms.SelectInputField{
		Name:        "mobilityDisabilityLevel",
		DisplayName: "Mobility disability level",
		Options:     getDisabilityLevels(),
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.MobilityDisabilityLevel)
}

func (f *IndividualForm) buildHasCognitiveDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasCognitiveDisability",
		DisplayName: "Has Cognitive Disability",
	}, f.disabilitiesSection, f.individual.HasCognitiveDisability)
}

func (f *IndividualForm) buildCognitiveDisabilityLevel() error {
	return buildField(&forms.SelectInputField{
		Name:        "cognitiveDisabilityLevel",
		DisplayName: "Cognitive disability level",
		Options:     getDisabilityLevels(),
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.CognitiveDisabilityLevel)
}

func (f *IndividualForm) buildHasSelfCareDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasSelfCareDisability",
		DisplayName: "Has SelfCare Disability",
	}, f.disabilitiesSection, f.individual.HasSelfCareDisability)
}

func (f *IndividualForm) buildSelfCareDisabilityLevel() error {
	return buildField(&forms.SelectInputField{
		Name:        "selfCareDisabilityLevel",
		DisplayName: "SelfCare disability level",
		Options:     getDisabilityLevels(),
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.SelfCareDisabilityLevel)
}

func (f *IndividualForm) buildHasCommunicationDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasCommunicationDisability",
		DisplayName: "Has Communication Disability",
	}, f.disabilitiesSection, f.individual.HasCommunicationDisability)
}

func (f *IndividualForm) buildCommunicationDisabilityLevel() error {
	return buildField(&forms.SelectInputField{
		Name:        "communicationDisabilityLevel",
		DisplayName: "Communication disability level",
		Options:     getDisabilityLevels(),
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.CommunicationDisabilityLevel)
}

func (f *IndividualForm) buildEngagementContext() error {
	return buildField(&forms.SelectInputField{
		Name:        "engagementContext",
		DisplayName: "Context of Engagement",
		Options: []forms.SelectInputFieldOption{
			{Label: "", Value: ""},
			{Label: "House Visit", Value: "houseVisit"},
			{Label: "Field Activity", Value: "fieldActivity"},
			{Label: "In-Office", Value: "inOffice"},
			{Label: "Remote Channels", Value: "remoteChannels"},
			{Label: "Referred", Value: "referred"},
			{Label: "Other", Value: "other"},
		},
	}, f.dataCollectionSection, f.individual.EngagementContext)
}

func (f *IndividualForm) buildCollectionAgent() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAgentName",
		DisplayName: "Collection Agent Name",
	}, f.dataCollectionSection, f.individual.CollectionAgentName)
}

func (f *IndividualForm) buildCollectionAgentTitle() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAgentTitle",
		DisplayName: "Collection Agent Title",
	}, f.dataCollectionSection, f.individual.CollectionAgentTitle)
}

func (f *IndividualForm) buildCollectionDate() error {
	return buildField(&forms.DateInputField{
		Name:        "collectionTime",
		DisplayName: "Collection Date",
	}, f.dataCollectionSection, f.individual.CollectionTime)
}

func (f *IndividualForm) buildCollectionLocation1() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAdministrativeArea1",
		DisplayName: "Collection Location 1",
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea1)
}

func (f *IndividualForm) buildCollectionLocation2() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAdministrativeArea2",
		DisplayName: "Collection Location 2",
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea2)
}

func (f *IndividualForm) buildCollectionLocation3() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAdministrativeArea3",
		DisplayName: "Collection Location 3",
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea3)
}

func (f *IndividualForm) buildCollectionOffice() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionOffice",
		DisplayName: "Collection Office",
	}, f.dataCollectionSection, f.individual.CollectionOffice)
}

func (f *IndividualForm) buildComments() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "comments",
		DisplayName: "Comments",
	}, f.dataCollectionSection, f.individual.Comments)
}

func (f *IndividualForm) buildFreeField1() error {
	return buildField(&forms.TextInputField{
		Name:        "freeField1",
		DisplayName: "Free Field 1",
	}, f.dataCollectionSection, f.individual.FreeField1)
}

func (f *IndividualForm) buildFreeField2() error {
	return buildField(&forms.TextInputField{
		Name:        "freeField2",
		DisplayName: "Free Field 2",
	}, f.dataCollectionSection, f.individual.FreeField2)
}

func (f *IndividualForm) buildFreeField3() error {
	return buildField(&forms.TextInputField{
		Name:        "freeField3",
		DisplayName: "Free Field 3",
	}, f.dataCollectionSection, f.individual.FreeField3)
}

func (f *IndividualForm) buildFreeField4() error {
	return buildField(&forms.TextInputField{
		Name:        "freeField4",
		DisplayName: "Free Field 4",
	}, f.dataCollectionSection, f.individual.FreeField4)
}

func (f *IndividualForm) buildFreeField5() error {
	return buildField(&forms.TextInputField{
		Name:        "freeField5",
		DisplayName: "Free Field 5",
	}, f.dataCollectionSection, f.individual.FreeField5)
}

func buildField(field forms.InputField, section *forms.FormSection, value interface{}) error {
	if err := field.SetValue(value); err != nil {
		return err
	}
	section.Fields = append(section.Fields, field)
	return nil
}

func getDisabilityLevels() []forms.SelectInputFieldOption {
	return []forms.SelectInputFieldOption{
		{Value: "0", Label: "No disability"},
		{Value: "1", Label: "Mild"},
		{Value: "2", Label: "Moderate"},
		{Value: "3", Label: "Severe"},
	}
}

func getIdentificationTypeOptions() []forms.SelectInputFieldOption {
	return []forms.SelectInputFieldOption{
		{Label: "", Value: ""},
		{Label: "Passport", Value: "passport"},
		{Label: "UNHCR ID", Value: "unhcr_id"},
		{Label: "National ID", Value: "national_id"},
		{Label: "Other", Value: "other"},
	}
}

func getSexOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, g := range api.AllSexes().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: g.String(),
			Value: string(g),
		})
	}
	return ret
}

func getDisplacementStatusOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, s := range api.AllDisplacementStatuses().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: s.String(),
			Value: string(s),
		})
	}
	return ret
}

func buildCountryOptions() []forms.SelectInputFieldOption {
	var opts = make([]forms.SelectInputFieldOption, 0, len(constants.Countries))
	opts = append(opts, forms.SelectInputFieldOption{
		Value: "",
		Label: "",
	})
	for _, country := range constants.Countries {
		opts = append(opts, forms.SelectInputFieldOption{
			Value: country.ISO3166Alpha3,
			Label: country.Name,
		})
	}
	return opts
}

func buildLanguageOptions() []forms.SelectInputFieldOption {
	var opts = make([]forms.SelectInputFieldOption, 0, len(constants.Languages))
	opts = append(opts, forms.SelectInputFieldOption{
		Value: "",
		Label: "",
	})
	for _, lang := range constants.Languages {
		opts = append(opts, forms.SelectInputFieldOption{
			Value: lang.ID,
			Label: lang.Name,
		})
	}
	return opts
}

type displacementStatusCodec struct{}

func (d *displacementStatusCodec) Encode(v interface{}) (string, error) {
	switch v.(type) {
	case api.DisplacementStatus:
		switch v.(api.DisplacementStatus) {
		case api.DisplacementStatusIDP:
			return string(api.DisplacementStatusIDP), nil
		case api.DisplacementStatusRefugee:
			return string(api.DisplacementStatusRefugee), nil
		case api.DisplacementStatusHostCommunity:
			return string(api.DisplacementStatusHostCommunity), nil
		case api.DisplacementStatusReturnee:
			return string(api.DisplacementStatusReturnee), nil
		case api.DisplacementStatusNonDisplaced:
			return string(api.DisplacementStatusNonDisplaced), nil
		case api.DisplacementStatusOther:
			return string(api.DisplacementStatusOther), nil
		case api.DisplacementStatusUnspecified:
			return string(api.DisplacementStatusUnspecified), nil
		default:
			return "", fmt.Errorf("invalid displacement status: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid displacement status type: %T", v)
	}
}

func (d *displacementStatusCodec) Decode(v string) (interface{}, error) {
	switch v {
	case string(api.DisplacementStatusIDP):
		return api.DisplacementStatusIDP, nil
	case string(api.DisplacementStatusRefugee):
		return api.DisplacementStatusRefugee, nil
	case string(api.DisplacementStatusHostCommunity):
		return api.DisplacementStatusHostCommunity, nil
	case string(api.DisplacementStatusReturnee):
		return api.DisplacementStatusReturnee, nil
	case string(api.DisplacementStatusNonDisplaced):
		return api.DisplacementStatusNonDisplaced, nil
	case string(api.DisplacementStatusOther):
		return api.DisplacementStatusOther, nil
	case string(api.DisplacementStatusUnspecified):
		return api.DisplacementStatusUnspecified, nil
	default:
		return nil, fmt.Errorf("invalid displacement status: %v", v)
	}
}

type disabilityLevelCodec struct{}

func (d disabilityLevelCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case api.DisabilityLevel:
		switch v {
		case api.DisabilityLevelNone:
			return "0", nil
		case api.DisabilityLevelMild:
			return "1", nil
		case api.DisabilityLevelModerate:
			return "2", nil
		case api.DisabilityLevelSevere:
			return "3", nil
		case api.DisabilityLevelUnspecified:
			return "", nil
		default:
			return "", fmt.Errorf("unknown disability level: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid type %T", value)
	}
}

func (d disabilityLevelCodec) Decode(value string) (interface{}, error) {
	switch value {
	case "0":
		return api.DisabilityLevelNone, nil
	case "1":
		return api.DisabilityLevelMild, nil
	case "2":
		return api.DisabilityLevelModerate, nil
	case "3":
		return api.DisabilityLevelSevere, nil
	case "":
		return api.DisabilityLevelUnspecified, nil
	default:
		return nil, fmt.Errorf("unknown disability level: %v", value)
	}
}

var _ forms.Codec = &disabilityLevelCodec{}

type sexCodec struct{}

func (g sexCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case api.Sex:
		switch v {
		case api.SexMale:
			return string(api.SexMale), nil
		case api.SexFemale:
			return string(api.SexFemale), nil
		case api.SexOther:
			return string(api.SexOther), nil
		case api.SexPreferNotToSay:
			return string(api.SexPreferNotToSay), nil
		case api.SexUnspecified:
			return string(api.SexUnspecified), nil
		default:
			return "", fmt.Errorf("unknown sex: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid type %T", value)
	}
}

func (g sexCodec) Decode(value string) (interface{}, error) {
	switch value {
	case string(api.SexMale):
		return api.SexMale, nil
	case string(api.SexFemale):
		return api.SexFemale, nil
	case string(api.SexOther):
		return api.SexOther, nil
	case string(api.SexPreferNotToSay):
		return api.SexPreferNotToSay, nil
	case string(api.SexUnspecified):
		return api.SexUnspecified, nil
	default:
		return nil, fmt.Errorf("unknown sex: %v", value)
	}
}
