package views

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/api/enumTypes"

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
	serviceSection        *forms.FormSection
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
		f.buildServiceSection,
	}

	fieldBuilders := []builderFuncs{
		f.buildTitle,
		f.buildIdField,
		f.buildFullName,
		f.buildPreferredName,
		f.buildFirstName,
		f.buildMiddleName,
		f.buildLastName,
		f.buildNativeName,
		f.buildMothersName,
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
		f.buildHouseholdSize,
		f.buildIsHeadOfHousehold,
		f.buildIsFemaleHeadedHousehold,
		f.buildIsMinorHeadedHousehold,
		f.buildCommunityID,
		f.buildCommunitySize,
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
		f.buildPreferredContactMethod,
		f.buildContactInstructions,
		f.buildHasConsentedToRgpd,
		f.buildHasConsentedToReferral,
		f.buildPresentsProtectionConcerns,
		f.buildDisplacementStatus,
		f.buildDisplacementStatusComment,
		f.buildHasDisability,
		f.buildPWDComments,
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
		f.buildServiceCC(1),
		f.buildServiceRequestedDate(1),
		f.buildServiceDeliveredDate(1),
		f.buildServiceComments(1),
		f.buildServiceCC(2),
		f.buildServiceRequestedDate(2),
		f.buildServiceDeliveredDate(2),
		f.buildServiceComments(2),
		f.buildServiceCC(3),
		f.buildServiceRequestedDate(3),
		f.buildServiceDeliveredDate(3),
		f.buildServiceComments(3),
		f.buildServiceCC(4),
		f.buildServiceRequestedDate(4),
		f.buildServiceDeliveredDate(4),
		f.buildServiceComments(4),
		f.buildServiceCC(5),
		f.buildServiceRequestedDate(5),
		f.buildServiceDeliveredDate(5),
		f.buildServiceComments(5),
		f.buildServiceCC(6),
		f.buildServiceRequestedDate(6),
		f.buildServiceDeliveredDate(6),
		f.buildServiceComments(6),
		f.buildServiceCC(7),
		f.buildServiceRequestedDate(7),
		f.buildServiceDeliveredDate(7),
		f.buildServiceComments(7),
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
		f.Form.Title = "New Participant"
	} else if f.individual.FullName == "" {
		f.Form.Title = "Anonymous Participant"
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

func (f *IndividualForm) buildServiceSection() error {
	f.serviceSection = &forms.FormSection{
		Title:       "Services",
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.serviceSection)
	return nil
}

func (f *IndividualForm) buildIdField() error {
	if !f.isNew() {
		return buildField(&forms.IDField{
			Name:        "id",
			DisplayName: "ID",
			QRCodeURL:   fmt.Sprintf("/countries/%s/participants/%s", f.individual.CountryID, f.individual.ID),
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
		DisplayName: "Surname",
	}, f.personalInfoSection, f.individual.LastName)
}

func (f *IndividualForm) buildMothersName() error {
	return buildField(&forms.TextInputField{
		Name:        "mothersName",
		DisplayName: "Mother's Name",
	}, f.personalInfoSection, f.individual.MothersName)
}

func (f *IndividualForm) buildNativeName() error {
	return buildField(&forms.TextInputField{
		Name:        "nativeName",
		DisplayName: "Native Name",
	}, f.personalInfoSection, f.individual.NativeName)
}

func (f *IndividualForm) buildPrefersToRemainAnonymous() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "prefersToRemainAnonymous",
		DisplayName: "The person prefers to remain anonymous",
	}, f.personalInfoSection, f.individual.PrefersToRemainAnonymous)
}

func (f *IndividualForm) buildSex() error {
	sexOptions := getSexOptions()
	sexOptions = append([]forms.SelectInputFieldOption{
		{Value: "", Label: "Select a value"},
	}, sexOptions...)
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
		DisplayName: "Date of Birth",
		MinValue:    "1900-01-01",
	}, f.personalInfoSection, f.individual.BirthDate)
}

func (f *IndividualForm) buildAge() error {
	return buildField(&forms.NumberInputField{
		Name:        "age",
		DisplayName: "Age of individual",
	}, f.personalInfoSection, f.individual.Age)
}

func (f *IndividualForm) buildIsMinor() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isMinor",
		DisplayName: "Is the person a minor",
	}, f.personalInfoSection, f.individual.IsMinor)
}

func (f *IndividualForm) buildNationality1() error {
	return buildField(&forms.SelectInputField{
		Name:        "nationality1",
		DisplayName: "Nationality (main)",
		Options:     buildCountryOptions(),
	}, f.personalInfoSection, f.individual.Nationality1)
}

func (f *IndividualForm) buildNationality2() error {
	return buildField(&forms.SelectInputField{
		Name:        "nationality2",
		DisplayName: "Nationality (secondary)",
		Options:     buildCountryOptions(),
	}, f.personalInfoSection, f.individual.Nationality2)
}

func (f *IndividualForm) buildIdentification1Type() error {
	options := getIdentificationTypeOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "identificationType1",
		DisplayName: "Type of primary identification",
		Options:     options,
		Codec:       &identificationTypeCodec{},
	}, f.personalInfoSection, f.individual.IdentificationType1)
}

func (f *IndividualForm) buildIdentification1Other() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "identificationTypeExplanation1",
		DisplayName: "If other, please explain",
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation1)
}

func (f *IndividualForm) buildIdentification1Number() error {
	return buildField(&forms.TextInputField{
		Name:        "identificationNumber1",
		DisplayName: "Primary Identification number",
	}, f.personalInfoSection, f.individual.IdentificationNumber1)
}

func (f *IndividualForm) buildIdentification2Type() error {
	options := getIdentificationTypeOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "identificationType2",
		Options:     options,
		Codec:       &identificationTypeCodec{},
		DisplayName: "Type of secondary identification",
	}, f.personalInfoSection, f.individual.IdentificationType2)
}

func (f *IndividualForm) buildIdentification2Other() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "identificationTypeExplanation2",
		DisplayName: "If other, please explain",
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation2)
}

func (f *IndividualForm) buildIdentification2Number() error {
	return buildField(&forms.TextInputField{
		Name:        "identificationNumber2",
		DisplayName: "Secondary Identification number",
	}, f.personalInfoSection, f.individual.IdentificationNumber2)
}
func (f *IndividualForm) buildIdentification3Type() error {
	options := getIdentificationTypeOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "identificationType3",
		DisplayName: "Type of tertiary identification",
		Options:     options,
		Codec:       &identificationTypeCodec{},
	}, f.personalInfoSection, f.individual.IdentificationType3)
}

func (f *IndividualForm) buildIdentification3Other() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "identificationTypeExplanation3",
		DisplayName: "If other, please explain",
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation3)
}

func (f *IndividualForm) buildIdentification3Number() error {
	return buildField(&forms.TextInputField{
		Name:        "identificationNumber3",
		DisplayName: "Tertiary Identification number",
	}, f.personalInfoSection, f.individual.IdentificationNumber3)
}

func (f *IndividualForm) buildInternalID() error {
	return buildField(&forms.TextInputField{
		Name:        "internalId",
		DisplayName: "NRC ID number (if any)",
	}, f.personalInfoSection, f.individual.InternalID)
}

func (f *IndividualForm) buildHouseholdID() error {
	return buildField(&forms.TextInputField{
		Name:        "householdId",
		DisplayName: "Household ID Number",
	}, f.personalInfoSection, f.individual.HouseholdID)
}

func (f *IndividualForm) buildHouseholdSize() error {
	return buildField(&forms.NumberInputField{
		Name:        "householdSize",
		DisplayName: "Household Size",
	}, f.personalInfoSection, f.individual.HouseholdSize)
}

func (f *IndividualForm) buildIsHeadOfHousehold() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isHeadOfHousehold",
		DisplayName: "Is the person the head of the household",
	}, f.personalInfoSection, f.individual.IsHeadOfHousehold)
}

func (f *IndividualForm) buildIsFemaleHeadedHousehold() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isFemaleHeadedHousehold",
		DisplayName: "Is the household headed by a female",
	}, f.personalInfoSection, f.individual.IsFemaleHeadedHousehold)
}

func (f *IndividualForm) buildIsMinorHeadedHousehold() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isMinorHeadedHousehold",
		DisplayName: "Is the household headed by a minor",
	}, f.personalInfoSection, f.individual.IsMinorHeadedHousehold)
}

func (f *IndividualForm) buildCommunityID() error {
	return buildField(&forms.TextInputField{
		Name:        "communityId",
		DisplayName: "Community ID number",
	}, f.personalInfoSection, f.individual.CommunityID)
}

func (f *IndividualForm) buildCommunitySize() error {
	return buildField(&forms.NumberInputField{
		Name:        "communitySize",
		DisplayName: "Community Size",
	}, f.personalInfoSection, f.individual.CommunitySize)
}

func (f *IndividualForm) buildIsHeadOfCommunity() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "isHeadOfCommunity",
		DisplayName: "Is the person the representative of the community",
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
		DisplayName: "Spoken Preferred Language",
		Options:     buildLanguageOptions(),
	}, f.personalInfoSection, f.individual.PreferredCommunicationLanguage)
}

func (f *IndividualForm) buildPhoneNumber1() error {
	return buildField(&forms.TextInputField{
		Name:        "phoneNumber1",
		DisplayName: "Phone Number #1",
	}, f.contactInfoSection, f.individual.PhoneNumber1)
}

func (f *IndividualForm) buildPhoneNumber2() error {
	return buildField(&forms.TextInputField{
		Name:        "phoneNumber2",
		DisplayName: "Phone Number #2",
	}, f.contactInfoSection, f.individual.PhoneNumber2)
}

func (f *IndividualForm) buildPhoneNumber3() error {
	return buildField(&forms.TextInputField{
		Name:        "phoneNumber3",
		DisplayName: "Phone Number #3",
	}, f.contactInfoSection, f.individual.PhoneNumber3)
}

func (f *IndividualForm) buildEmailAddress1() error {
	return buildField(&forms.TextInputField{
		Name:        "email1",
		DisplayName: "Email Address #1",
	}, f.contactInfoSection, f.individual.Email1)
}

func (f *IndividualForm) buildEmailAddress2() error {
	return buildField(&forms.TextInputField{
		Name:        "email2",
		DisplayName: "Email Address #2",
	}, f.contactInfoSection, f.individual.Email2)
}

func (f *IndividualForm) buildEmailAddress3() error {
	return buildField(&forms.TextInputField{
		Name:        "email3",
		DisplayName: "Email Address #3",
	}, f.contactInfoSection, f.individual.Email3)
}

func (f *IndividualForm) buildAddress() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "address",
		DisplayName: "Residence Address",
	}, f.contactInfoSection, f.individual.Address)
}

func (f *IndividualForm) buildPreferredContactMethod() error {
	options := getPreferredContactMethodOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "preferredContactMethod",
		DisplayName: "Preferred means of contact ",
		Options:     options,
		Codec:       &preferredContactMethodCodec{},
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
		DisplayName: "The person presents protection concerns",
	}, f.protectionSection, f.individual.PresentsProtectionConcerns)
}

func (f *IndividualForm) buildDisplacementStatus() error {
	options := getDisplacementStatusOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
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
		DisplayName: "If other, please explain",
	}, f.protectionSection, f.individual.DisplacementStatusComment)
}

func (f *IndividualForm) buildHasVisionDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasVisionDisability",
		DisplayName: "Does the person have difficulty seeing, even if wearing glasses?",
	}, f.disabilitiesSection, f.individual.HasVisionDisability)
}

func (f *IndividualForm) buildVisionDisabilityLevel() error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "visionDisabilityLevel",
		DisplayName: "Would you say:",
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.VisionDisabilityLevel)
}

func (f *IndividualForm) buildHasDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasDisability",
		DisplayName: "Does the participant have any disability (PWD)?",
	}, f.disabilitiesSection, f.individual.HasDisability)
}

func (f *IndividualForm) buildPWDComments() error {
	return buildField(&forms.TextAreaInputField{
		Name:        "PWDComments",
		DisplayName: "PWD Comments",
	}, f.disabilitiesSection, f.individual.PWDComments)
}

func (f *IndividualForm) buildHasHearingDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasHearingDisability",
		DisplayName: "Does the person have difficulty hearing, even if using a hearing aid?",
	}, f.disabilitiesSection, f.individual.HasHearingDisability)
}

func (f *IndividualForm) buildHearingDisabilityLevel() error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "hearingDisabilityLevel",
		DisplayName: "Would you say:",
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.HearingDisabilityLevel)
}

func (f *IndividualForm) buildHasMobilityDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasMobilityDisability",
		DisplayName: "Do you have difficulty walking or climbing steps?",
	}, f.disabilitiesSection, f.individual.HasMobilityDisability)
}

func (f *IndividualForm) buildMobilityDisabilityLevel() error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "mobilityDisabilityLevel",
		DisplayName: "Would you say:",
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.MobilityDisabilityLevel)
}

func (f *IndividualForm) buildHasCognitiveDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasCognitiveDisability",
		DisplayName: "Does the person have difficulty remembering or concentrating?",
	}, f.disabilitiesSection, f.individual.HasCognitiveDisability)
}

func (f *IndividualForm) buildCognitiveDisabilityLevel() error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "cognitiveDisabilityLevel",
		DisplayName: "Would you say:",
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.CognitiveDisabilityLevel)
}

func (f *IndividualForm) buildHasSelfCareDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasSelfCareDisability",
		DisplayName: "Does the person have difficulty with self-care such as washing all over or dressing?",
	}, f.disabilitiesSection, f.individual.HasSelfCareDisability)
}

func (f *IndividualForm) buildSelfCareDisabilityLevel() error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "selfCareDisabilityLevel",
		DisplayName: "Would you say:",
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.SelfCareDisabilityLevel)
}

func (f *IndividualForm) buildHasCommunicationDisability() error {
	return buildField(&forms.CheckboxInputField{
		Name:        "hasCommunicationDisability",
		DisplayName: "Using their usual (customary) language, do they have difficulty communicating, for example understanding or being understood?",
	}, f.disabilitiesSection, f.individual.HasCommunicationDisability)
}

func (f *IndividualForm) buildCommunicationDisabilityLevel() error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "communicationDisabilityLevel",
		DisplayName: "Would you say:",
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.CommunicationDisabilityLevel)
}

func (f *IndividualForm) buildEngagementContext() error {
	options := getEngagementContextOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        "engagementContext",
		DisplayName: "Context of Engagement",
		Options:     options,
		Codec:       &engagementContextCodec{},
	}, f.dataCollectionSection, f.individual.EngagementContext)
}

func (f *IndividualForm) buildCollectionAgent() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAgentName",
		DisplayName: "NRC agent responsible for data collection",
	}, f.dataCollectionSection, f.individual.CollectionAgentName)
}

func (f *IndividualForm) buildCollectionAgentTitle() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAgentTitle",
		DisplayName: "Title of NRC agent responsible for data collection",
	}, f.dataCollectionSection, f.individual.CollectionAgentTitle)
}

func (f *IndividualForm) buildCollectionDate() error {
	return buildField(&forms.DateInputField{
		Name:        "collectionTime",
		DisplayName: "Date of registration",
	}, f.dataCollectionSection, f.individual.CollectionTime)
}

func (f *IndividualForm) buildCollectionLocation1() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAdministrativeArea1",
		DisplayName: "Location of registration (admin1)",
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea1)
}

func (f *IndividualForm) buildCollectionLocation2() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAdministrativeArea2",
		DisplayName: "Location of registration (admin2)",
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea2)
}

func (f *IndividualForm) buildCollectionLocation3() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionAdministrativeArea3",
		DisplayName: "Location of registration (admin3)",
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea3)
}

func (f *IndividualForm) buildCollectionOffice() error {
	return buildField(&forms.TextInputField{
		Name:        "collectionOffice",
		DisplayName: "Main office where the individual is linked to",
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

func (f *IndividualForm) buildServiceCC(idx int) func() error {
	return func() error {
		var value interface{}
		switch idx {
		case 1:
			value = f.individual.ServiceCC1
		case 2:
			value = f.individual.ServiceCC2
		case 3:
			value = f.individual.ServiceCC3
		case 4:
			value = f.individual.ServiceCC4
		case 5:
			value = f.individual.ServiceCC5
		case 6:
			value = f.individual.ServiceCC6
		case 7:
			value = f.individual.ServiceCC7
		default:
			return fmt.Errorf("invalid service CC index: %d", idx)
		}

		options := getServiceCCOptions()
		options = append([]forms.SelectInputFieldOption{{Value: "", Label: "Select a value"}}, options...)
		return buildField(&forms.SelectInputField{
			Name:        fmt.Sprintf("serviceCC%d", idx),
			DisplayName: fmt.Sprintf("Type of service #%d", idx),
			Options:     options,
			Codec:       &serviceCCCodec{},
		}, f.serviceSection, value)
	}
}

func (f *IndividualForm) buildServiceRequestedDate(idx int) func() error {
	return func() error {
		var value interface{}
		switch idx {
		case 1:
			value = f.individual.ServiceRequestedDate1
		case 2:
			value = f.individual.ServiceRequestedDate2
		case 3:
			value = f.individual.ServiceRequestedDate3
		case 4:
			value = f.individual.ServiceRequestedDate4
		case 5:
			value = f.individual.ServiceRequestedDate5
		case 6:
			value = f.individual.ServiceRequestedDate6
		case 7:
			value = f.individual.ServiceRequestedDate7
		default:
			return fmt.Errorf("invalid service requested date index: %d", idx)
		}

		return buildField(&forms.DateInputField{
			Name:        fmt.Sprintf("serviceRequestedDate%d", idx),
			DisplayName: fmt.Sprintf("Date of service #%d requested", idx),
		}, f.serviceSection, value)
	}
}

func (f *IndividualForm) buildServiceDeliveredDate(idx int) func() error {
	return func() error {
		var value interface{}
		switch idx {
		case 1:
			value = f.individual.ServiceDeliveredDate1
		case 2:
			value = f.individual.ServiceDeliveredDate2
		case 3:
			value = f.individual.ServiceDeliveredDate3
		case 4:
			value = f.individual.ServiceDeliveredDate4
		case 5:
			value = f.individual.ServiceDeliveredDate5
		case 6:
			value = f.individual.ServiceDeliveredDate6
		case 7:
			value = f.individual.ServiceDeliveredDate7
		default:
			return fmt.Errorf("invalid service delivered date index: %d", idx)
		}

		return buildField(&forms.DateInputField{
			Name:        fmt.Sprintf("serviceDeliveredDate%d", idx),
			DisplayName: fmt.Sprintf("Date of service #%d delivered", idx),
		}, f.serviceSection, value)
	}
}

func (f *IndividualForm) buildServiceComments(idx int) func() error {
	return func() error {
		var value interface{}
		switch idx {
		case 1:
			value = f.individual.ServiceComments1
		case 2:
			value = f.individual.ServiceComments2
		case 3:
			value = f.individual.ServiceComments3
		case 4:
			value = f.individual.ServiceComments4
		case 5:
			value = f.individual.ServiceComments5
		case 6:
			value = f.individual.ServiceComments6
		case 7:
			value = f.individual.ServiceComments7
		default:
			return fmt.Errorf("invalid service comments index: %d", idx)
		}

		return buildField(&forms.TextAreaInputField{
			Name:        fmt.Sprintf("serviceComments%d", idx),
			DisplayName: fmt.Sprintf("Service %d Comments", idx),
		}, f.serviceSection, value)
	}
}

func buildField(field forms.InputField, section *forms.FormSection, value interface{}) error {
	if err := field.SetValue(value); err != nil {
		return err
	}
	section.Fields = append(section.Fields, field)
	return nil
}

func getDisabilityLevels() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, g := range enumTypes.AllDisabilityLevels().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: g.String(),
			Value: string(g),
		})
	}
	return ret
}

func getIdentificationTypeOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, g := range enumTypes.AllIdentificationTypes().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: g.String(),
			Value: string(g),
		})
	}
	return ret
}

func getEngagementContextOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, g := range enumTypes.AllEngagementContexts().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: g.String(),
			Value: string(g),
		})
	}
	return ret
}

func getSexOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, g := range enumTypes.AllSexes().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: g.String(),
			Value: string(g),
		})
	}
	return ret
}

func getDisplacementStatusOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, s := range enumTypes.AllDisplacementStatuses().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: s.String(),
			Value: string(s),
		})
	}
	return ret
}

func getPreferredContactMethodOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, s := range enumTypes.AllContactMethods().Items() {
		ret = append(ret, forms.SelectInputFieldOption{
			Label: s.String(),
			Value: string(s),
		})
	}
	return ret
}

func getServiceCCOptions() []forms.SelectInputFieldOption {
	var ret []forms.SelectInputFieldOption
	for _, s := range enumTypes.AllServiceCCs().Items() {
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
		Label: "Select a value",
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
		Label: "Select a value",
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
	case enumTypes.DisplacementStatus:
		switch v.(enumTypes.DisplacementStatus) {
		case enumTypes.DisplacementStatusIDP:
			return string(enumTypes.DisplacementStatusIDP), nil
		case enumTypes.DisplacementStatusRefugee:
			return string(enumTypes.DisplacementStatusRefugee), nil
		case enumTypes.DisplacementStatusHostCommunity:
			return string(enumTypes.DisplacementStatusHostCommunity), nil
		case enumTypes.DisplacementStatusReturnee:
			return string(enumTypes.DisplacementStatusReturnee), nil
		case enumTypes.DisplacementStatusAsylumSeeker:
			return string(enumTypes.DisplacementStatusAsylumSeeker), nil
		case enumTypes.DisplacementStatusNonDisplaced:
			return string(enumTypes.DisplacementStatusNonDisplaced), nil
		case enumTypes.DisplacementStatusOther:
			return string(enumTypes.DisplacementStatusOther), nil
		case enumTypes.DisplacementStatusUnspecified:
			return string(enumTypes.DisplacementStatusUnspecified), nil
		default:
			return "", fmt.Errorf("invalid displacement status: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid displacement status type: %T", v)
	}
}

func (d *displacementStatusCodec) Decode(v string) (interface{}, error) {
	switch v {
	case string(enumTypes.DisplacementStatusIDP):
		return enumTypes.DisplacementStatusIDP, nil
	case string(enumTypes.DisplacementStatusRefugee):
		return enumTypes.DisplacementStatusRefugee, nil
	case string(enumTypes.DisplacementStatusHostCommunity):
		return enumTypes.DisplacementStatusHostCommunity, nil
	case string(enumTypes.DisplacementStatusReturnee):
		return enumTypes.DisplacementStatusReturnee, nil
	case string(enumTypes.DisplacementStatusAsylumSeeker):
		return enumTypes.DisplacementStatusAsylumSeeker, nil
	case string(enumTypes.DisplacementStatusNonDisplaced):
		return enumTypes.DisplacementStatusNonDisplaced, nil
	case string(enumTypes.DisplacementStatusOther):
		return enumTypes.DisplacementStatusOther, nil
	case string(enumTypes.DisplacementStatusUnspecified):
		return enumTypes.DisplacementStatusUnspecified, nil
	default:
		return nil, fmt.Errorf("invalid displacement status: %v", v)
	}
}

type identificationTypeCodec struct{}

func (d *identificationTypeCodec) Encode(v interface{}) (string, error) {
	switch v.(type) {
	case enumTypes.IdentificationType:
		switch v.(enumTypes.IdentificationType) {
		case enumTypes.IdentificationTypeNational:
			return string(enumTypes.IdentificationTypeNational), nil
		case enumTypes.IdentificationTypePassport:
			return string(enumTypes.IdentificationTypePassport), nil
		case enumTypes.IdentificationTypeUNHCR:
			return string(enumTypes.IdentificationTypeUNHCR), nil
		case enumTypes.IdentificationTypeOther:
			return string(enumTypes.IdentificationTypeOther), nil
		case enumTypes.IdentificationTypeUnspecified:
			return string(enumTypes.IdentificationTypeUnspecified), nil
		default:
			return "", fmt.Errorf("invalid identificationType: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid identificationType type: %T", v)
	}
}

func (d *identificationTypeCodec) Decode(v string) (interface{}, error) {
	switch v {
	case string(enumTypes.IdentificationTypeNational):
		return enumTypes.IdentificationTypeNational, nil
	case string(enumTypes.IdentificationTypeUNHCR):
		return enumTypes.IdentificationTypeUNHCR, nil
	case string(enumTypes.IdentificationTypePassport):
		return enumTypes.IdentificationTypePassport, nil
	case string(enumTypes.IdentificationTypeOther):
		return enumTypes.IdentificationTypeOther, nil
	case string(enumTypes.IdentificationTypeUnspecified):
		return enumTypes.IdentificationTypeUnspecified, nil
	default:
		return nil, fmt.Errorf("invalid identificationType: %v", v)
	}
}

type serviceCCCodec struct{}

func (d *serviceCCCodec) Encode(v interface{}) (string, error) {
	switch v.(type) {
	case enumTypes.ServiceCC:
		switch v.(enumTypes.ServiceCC) {
		case enumTypes.ServiceCCNone:
			return string(enumTypes.ServiceCCNone), nil
		case enumTypes.ServiceCCShelter:
			return string(enumTypes.ServiceCCShelter), nil
		case enumTypes.ServiceCCWash:
			return string(enumTypes.ServiceCCWash), nil
		case enumTypes.ServiceCCProtection:
			return string(enumTypes.ServiceCCProtection), nil
		case enumTypes.ServiceCCEducation:
			return string(enumTypes.ServiceCCEducation), nil
		case enumTypes.ServiceCCICLA:
			return string(enumTypes.ServiceCCICLA), nil
		case enumTypes.ServiceCCLFS:
			return string(enumTypes.ServiceCCLFS), nil
		case enumTypes.ServiceCCCVA:
			return string(enumTypes.ServiceCCCVA), nil
		case enumTypes.ServiceCCOther:
			return string(enumTypes.ServiceCCOther), nil
		default:
			return "", fmt.Errorf("invalid service CC: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid service CC type: %T", v)
	}
}

func (d *serviceCCCodec) Decode(v string) (interface{}, error) {
	switch v {
	case string(enumTypes.ServiceCCNone):
		return enumTypes.ServiceCCNone, nil
	case string(enumTypes.ServiceCCShelter):
		return enumTypes.ServiceCCShelter, nil
	case string(enumTypes.ServiceCCWash):
		return enumTypes.ServiceCCWash, nil
	case string(enumTypes.ServiceCCProtection):
		return enumTypes.ServiceCCProtection, nil
	case string(enumTypes.ServiceCCEducation):
		return enumTypes.ServiceCCEducation, nil
	case string(enumTypes.ServiceCCICLA):
		return enumTypes.ServiceCCICLA, nil
	case string(enumTypes.ServiceCCLFS):
		return enumTypes.ServiceCCLFS, nil
	case string(enumTypes.ServiceCCCVA):
		return enumTypes.ServiceCCCVA, nil
	case string(enumTypes.ServiceCCOther):
		return enumTypes.ServiceCCOther, nil
	default:
		return nil, fmt.Errorf("invalid service CC: %v", v)
	}
}

type disabilityLevelCodec struct{}

func (d disabilityLevelCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case enumTypes.DisabilityLevel:
		switch v {
		case enumTypes.DisabilityLevelNone:
			return string(enumTypes.DisabilityLevelNone), nil
		case enumTypes.DisabilityLevelMild:
			return string(enumTypes.DisabilityLevelMild), nil
		case enumTypes.DisabilityLevelModerate:
			return string(enumTypes.DisabilityLevelModerate), nil
		case enumTypes.DisabilityLevelSevere:
			return string(enumTypes.DisabilityLevelSevere), nil
		case enumTypes.DisabilityLevelUnspecified:
			return string(enumTypes.DisabilityLevelUnspecified), nil
		default:
			return "", fmt.Errorf("unknown disability level: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid type %T", value)
	}
}

func (d disabilityLevelCodec) Decode(value string) (interface{}, error) {
	switch value {
	case string(enumTypes.DisabilityLevelNone):
		return enumTypes.DisabilityLevelNone, nil
	case string(enumTypes.DisabilityLevelMild):
		return enumTypes.DisabilityLevelMild, nil
	case string(enumTypes.DisabilityLevelModerate):
		return enumTypes.DisabilityLevelModerate, nil
	case string(enumTypes.DisabilityLevelSevere):
		return enumTypes.DisabilityLevelSevere, nil
	case string(enumTypes.DisabilityLevelUnspecified):
		return enumTypes.DisabilityLevelUnspecified, nil
	default:
		return nil, fmt.Errorf("unknown disability level: %v", value)
	}
}

type engagementContextCodec struct{}

func (d engagementContextCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case enumTypes.EngagementContext:
		switch v {
		case enumTypes.EngagementContextFieldActivity:
			return string(enumTypes.EngagementContextFieldActivity), nil
		case enumTypes.EngagementContextInOffice:
			return string(enumTypes.EngagementContextInOffice), nil
		case enumTypes.EngagementContextHouseVisit:
			return string(enumTypes.EngagementContextHouseVisit), nil
		case enumTypes.EngagementContextReferred:
			return string(enumTypes.EngagementContextReferred), nil
		case enumTypes.EngagementContextRemoteChannels:
			return string(enumTypes.EngagementContextRemoteChannels), nil
		case enumTypes.EngagementContextOther:
			return string(enumTypes.EngagementContextOther), nil
		case enumTypes.EngagementContextUnspecified:
			return string(enumTypes.EngagementContextUnspecified), nil
		default:
			return "", fmt.Errorf("unknown engagement context: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid type %T", value)
	}
}

func (d engagementContextCodec) Decode(value string) (interface{}, error) {
	switch value {
	case string(enumTypes.EngagementContextFieldActivity):
		return enumTypes.EngagementContextFieldActivity, nil
	case string(enumTypes.EngagementContextInOffice):
		return enumTypes.EngagementContextInOffice, nil
	case string(enumTypes.EngagementContextHouseVisit):
		return enumTypes.EngagementContextHouseVisit, nil
	case string(enumTypes.EngagementContextReferred):
		return enumTypes.EngagementContextReferred, nil
	case string(enumTypes.EngagementContextRemoteChannels):
		return enumTypes.EngagementContextRemoteChannels, nil
	case string(enumTypes.EngagementContextOther):
		return enumTypes.EngagementContextOther, nil
	case string(enumTypes.EngagementContextUnspecified):
		return enumTypes.EngagementContextUnspecified, nil
	default:
		return nil, fmt.Errorf("unknown engagement context: %v", value)
	}
}

var _ forms.Codec = &disabilityLevelCodec{}

type preferredContactMethodCodec struct{}

func (d preferredContactMethodCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case enumTypes.ContactMethod:
		switch v {
		case enumTypes.ContactMethodPhone:
			return string(enumTypes.ContactMethodPhone), nil
		case enumTypes.ContactMethodWhatsapp:
			return string(enumTypes.ContactMethodWhatsapp), nil
		case enumTypes.ContactMethodEmail:
			return string(enumTypes.ContactMethodEmail), nil
		case enumTypes.ContactMethodVisit:
			return string(enumTypes.ContactMethodVisit), nil
		case enumTypes.ContactMethodOther:
			return string(enumTypes.ContactMethodOther), nil
		case enumTypes.ContactMethodUnspecified:
			return string(enumTypes.ContactMethodUnspecified), nil
		default:
			return "", fmt.Errorf("unknown contact method: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid type %T", value)
	}
}

func (d preferredContactMethodCodec) Decode(value string) (interface{}, error) {
	switch value {
	case string(enumTypes.ContactMethodPhone):
		return enumTypes.ContactMethodPhone, nil
	case string(enumTypes.ContactMethodWhatsapp):
		return enumTypes.ContactMethodWhatsapp, nil
	case string(enumTypes.ContactMethodEmail):
		return enumTypes.ContactMethodEmail, nil
	case string(enumTypes.ContactMethodVisit):
		return enumTypes.ContactMethodVisit, nil
	case string(enumTypes.ContactMethodOther):
		return enumTypes.ContactMethodOther, nil
	case string(enumTypes.ContactMethodUnspecified):
		return enumTypes.ContactMethodUnspecified, nil
	default:
		return nil, fmt.Errorf("unknown contact method: %v", value)
	}
}

var _ forms.Codec = &preferredContactMethodCodec{}

type sexCodec struct{}

func (g sexCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case enumTypes.Sex:
		switch v {
		case enumTypes.SexMale:
			return string(enumTypes.SexMale), nil
		case enumTypes.SexFemale:
			return string(enumTypes.SexFemale), nil
		case enumTypes.SexOther:
			return string(enumTypes.SexOther), nil
		case enumTypes.SexPreferNotToSay:
			return string(enumTypes.SexPreferNotToSay), nil
		case enumTypes.SexUnspecified:
			return string(enumTypes.SexUnspecified), nil
		default:
			return "", fmt.Errorf("unknown sex: %v", v)
		}
	default:
		return "", fmt.Errorf("invalid type %T", value)
	}
}

func (g sexCodec) Decode(value string) (interface{}, error) {
	switch value {
	case string(enumTypes.SexMale):
		return enumTypes.SexMale, nil
	case string(enumTypes.SexFemale):
		return enumTypes.SexFemale, nil
	case string(enumTypes.SexOther):
		return enumTypes.SexOther, nil
	case string(enumTypes.SexPreferNotToSay):
		return enumTypes.SexPreferNotToSay, nil
	case string(enumTypes.SexUnspecified):
		return enumTypes.SexUnspecified, nil
	default:
		return nil, fmt.Errorf("unknown sex: %v", value)
	}
}
