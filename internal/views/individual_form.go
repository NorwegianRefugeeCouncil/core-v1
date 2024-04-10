package views

import (
	"fmt"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/pkg/views/forms"
)

type IndividualForm struct {
	*forms.Form
	individual             *api.Individual
	consentSection				 *forms.FormSection
	personalInfoSection    *forms.FormSection
	contactInfoSection     *forms.FormSection
	protectionSection      *forms.FormSection
	disabilitiesSection    *forms.FormSection
	vulnerabilitiesSection *forms.FormSection
	dataCollectionSection  *forms.FormSection
	serviceSection         *forms.FormSection
}

func NewIndividualForm(i *api.Individual) (*IndividualForm, error) {
	f := &IndividualForm{
		Form:       &forms.Form{},
		individual: i,
	}
	if err := f.build(locales.GetTranslator()); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *IndividualForm) build(t locales.Translator) error {
	type builderFuncs func(t locales.Translator) error

	runBuilderFunctions := func(builders ...builderFuncs) error {
		for _, builder := range builders {
			if err := builder(t); err != nil {
				return err
			}
		}
		return nil
	}

	sectionBuilders := []builderFuncs{
		f.buildConsentSection,
		f.buildPersonalInfoSection,
		f.buildContactInfoSection,
		f.buildProtectionSection,
		f.buildDisabilitiesSection,
		f.buildVulnerabilitiesSection,
		f.buildDataCollectionSection,
		f.buildServiceSection,
	}

	fieldBuilders := []builderFuncs{
		f.buildTitle,
		f.buildIdField,
		f.buildPrefersToRemainAnonymous,
		f.buildFullName,
		f.buildPreferredName,
		f.buildFirstName,
		f.buildMiddleName,
		f.buildLastName,
		f.buildNativeName,
		f.buildMothersName,
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
		f.buildSex,
		f.buildBirthDate,
		f.buildAge,
		f.buildIsMinor,
		f.buildIsChildAtRisk,
		f.buildIsElderAtRisk,
		f.buildIsWomanAtRisk,
		f.buildIsSingleParent,
		f.buildIsSeparatedChild,
		f.buildIsPregnant,
		f.buildIsLactating,
		f.buildHasMedicalCondition,
		f.buildNeedsLegalAndPhysicalProtection,
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
		f.buildVulnerabilityComments,
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
		f.buildServiceTextInputField("service_1", f.individual.Service1, "service_no", 1),
		f.buildServiceTextInputField("service_type_1", f.individual.ServiceType1, "service_type_no", 1),
		f.buildServiceTextInputField("service_sub_service_1", f.individual.ServiceSubService1, "service_sub_service_no", 1),
		f.buildServiceRequestedDate(1),
		f.buildServiceDeliveredDate(1),
		f.buildServiceComments(1),
		f.buildServiceTextInputField("service_location_1", f.individual.ServiceLocation1, "service_location_no", 1),
		f.buildServiceTextInputField("service_donor_1", f.individual.ServiceDonor1, "service_donor_no", 1),
		f.buildServiceTextInputField("service_project_name_1", f.individual.ServiceProjectName1, "service_project_name_no", 1),
		f.buildServiceTextInputField("service_agent_name_1", f.individual.ServiceAgentName1, "service_agent_name_no", 1),

		f.buildServiceCC(2),
		f.buildServiceTextInputField("service_2", f.individual.Service2, "service_no", 2),
		f.buildServiceTextInputField("service_type_2", f.individual.ServiceType2, "service_type_no", 2),
		f.buildServiceTextInputField("service_sub_service_2", f.individual.ServiceSubService2, "service_sub_service_no", 2),
		f.buildServiceRequestedDate(2),
		f.buildServiceDeliveredDate(2),
		f.buildServiceComments(2),
		f.buildServiceTextInputField("service_location_2", f.individual.ServiceLocation2, "service_location_no", 2),
		f.buildServiceTextInputField("service_donor_2", f.individual.ServiceDonor2, "service_donor_no", 2),
		f.buildServiceTextInputField("service_project_name_2", f.individual.ServiceProjectName2, "service_project_name_no", 2),
		f.buildServiceTextInputField("service_agent_name_2", f.individual.ServiceAgentName2, "service_agent_name_no", 2),

		f.buildServiceCC(3),
		f.buildServiceTextInputField("service_3", f.individual.Service3, "service_no", 3),
		f.buildServiceTextInputField("service_type_3", f.individual.ServiceType3, "service_type_no", 3),
		f.buildServiceTextInputField("service_sub_service_3", f.individual.ServiceSubService3, "service_sub_service_no", 3),
		f.buildServiceRequestedDate(3),
		f.buildServiceDeliveredDate(3),
		f.buildServiceComments(3),
		f.buildServiceTextInputField("service_location_3", f.individual.ServiceLocation3, "service_location_no", 3),
		f.buildServiceTextInputField("service_donor_3", f.individual.ServiceDonor3, "service_donor_no", 3),
		f.buildServiceTextInputField("service_project_name_3", f.individual.ServiceProjectName3, "service_project_name_no", 3),
		f.buildServiceTextInputField("service_agent_name_3", f.individual.ServiceAgentName3, "service_agent_name_no", 3),

		f.buildServiceCC(4),
		f.buildServiceTextInputField("service_4", f.individual.Service4, "service_no", 4),
		f.buildServiceTextInputField("service_type_4", f.individual.ServiceType4, "service_type_no", 4),
		f.buildServiceTextInputField("service_sub_service_4", f.individual.ServiceSubService4, "service_sub_service_no", 4),
		f.buildServiceRequestedDate(4),
		f.buildServiceDeliveredDate(4),
		f.buildServiceComments(4),
		f.buildServiceTextInputField("service_location_4", f.individual.ServiceLocation4, "service_location_no", 4),
		f.buildServiceTextInputField("service_donor_4", f.individual.ServiceDonor4, "service_donor_no", 4),
		f.buildServiceTextInputField("service_project_name_4", f.individual.ServiceProjectName4, "service_project_name_no", 4),
		f.buildServiceTextInputField("service_agent_name_4", f.individual.ServiceAgentName4, "service_agent_name_no", 4),

		f.buildServiceCC(5),
		f.buildServiceTextInputField("service_5", f.individual.Service5, "service_no", 5),
		f.buildServiceTextInputField("service_type_5", f.individual.ServiceType5, "service_type_no", 5),
		f.buildServiceTextInputField("service_sub_service_5", f.individual.ServiceSubService5, "service_sub_service_no", 5),
		f.buildServiceRequestedDate(5),
		f.buildServiceDeliveredDate(5),
		f.buildServiceComments(5),
		f.buildServiceTextInputField("service_location_5", f.individual.ServiceLocation5, "service_location_no", 5),
		f.buildServiceTextInputField("service_donor_5", f.individual.ServiceDonor5, "service_donor_no", 5),
		f.buildServiceTextInputField("service_project_name_5", f.individual.ServiceProjectName5, "service_project_name_no", 5),
		f.buildServiceTextInputField("service_agent_name_5", f.individual.ServiceAgentName5, "service_agent_name_no", 5),

		f.buildServiceCC(6),
		f.buildServiceTextInputField("service_6", f.individual.Service6, "service_no", 6),
		f.buildServiceTextInputField("service_type_6", f.individual.ServiceType6, "service_type_no", 6),
		f.buildServiceTextInputField("service_sub_service_6", f.individual.ServiceSubService6, "service_sub_service_no", 6),
		f.buildServiceRequestedDate(6),
		f.buildServiceDeliveredDate(6),
		f.buildServiceComments(6),
		f.buildServiceTextInputField("service_location_6", f.individual.ServiceLocation6, "service_location_no", 6),
		f.buildServiceTextInputField("service_donor_6", f.individual.ServiceDonor6, "service_donor_no", 6),
		f.buildServiceTextInputField("service_project_name_6", f.individual.ServiceProjectName6, "service_project_name_no", 6),
		f.buildServiceTextInputField("service_agent_name_6", f.individual.ServiceAgentName6, "service_agent_name_no", 6),

		f.buildServiceCC(7),
		f.buildServiceTextInputField("service_7", f.individual.Service7, "service_no", 7),
		f.buildServiceTextInputField("service_type_7", f.individual.ServiceType7, "service_type_no", 7),
		f.buildServiceTextInputField("service_sub_service_7", f.individual.ServiceSubService7, "service_sub_service_no", 7),
		f.buildServiceRequestedDate(7),
		f.buildServiceDeliveredDate(7),
		f.buildServiceComments(7),
		f.buildServiceTextInputField("service_location_7", f.individual.ServiceLocation7, "service_location_no", 7),
		f.buildServiceTextInputField("service_donor_7", f.individual.ServiceDonor7, "service_donor_no", 7),
		f.buildServiceTextInputField("service_project_name_7", f.individual.ServiceProjectName7, "service_project_name_no", 7),
		f.buildServiceTextInputField("service_agent_name_7", f.individual.ServiceAgentName7, "service_agent_name_no", 7),
	}

	if err := runBuilderFunctions(sectionBuilders...); err != nil {
		return err
	}

	if err := runBuilderFunctions(fieldBuilders...); err != nil {
		return err
	}

	return nil
}

func (f *IndividualForm) buildTitle(t locales.Translator) error {
	if f.isNew() {
		f.Form.Title = t("new_participant")
	} else if f.individual.FullName == "" && f.individual.FirstName == "" && f.individual.MiddleName == "" && f.individual.LastName == "" && f.individual.NativeName == "" {
		f.Form.Title = t("anonymous_participant")
	} else {
		if f.individual.FullName != "" {
			f.Form.Title = f.individual.FullName
		} else if f.individual.PreferredName != "" {
			f.Form.Title = f.individual.PreferredName
		} else if f.individual.FirstName != "" || f.individual.MiddleName != "" || f.individual.LastName != "" {
			f.Form.Title = f.individual.FirstName + " " + f.individual.MiddleName + " " + f.individual.LastName
		} else if f.individual.NativeName != "" {
			f.Form.Title = f.individual.NativeName
		}
	}
	return nil
}

func (f *IndividualForm) buildConsentSection(t locales.Translator) error {
	f.consentSection = &forms.FormSection{
		Title:       t("consent"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   false,
	}
	f.Form.Sections = append(f.Form.Sections, f.consentSection)
	return nil
}

func (f *IndividualForm) buildPersonalInfoSection(t locales.Translator) error {
	f.personalInfoSection = &forms.FormSection{
		Title:       t("personal_info"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   false,
	}
	f.Form.Sections = append(f.Form.Sections, f.personalInfoSection)
	return nil
}

func (f *IndividualForm) buildContactInfoSection(t locales.Translator) error {
	f.contactInfoSection = &forms.FormSection{
		Title:       t("contact"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.contactInfoSection)
	return nil
}

func (f *IndividualForm) buildProtectionSection(t locales.Translator) error {
	f.protectionSection = &forms.FormSection{
		Title:       t("protection"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.protectionSection)
	return nil
}

func (f *IndividualForm) buildDisabilitiesSection(t locales.Translator) error {
	f.disabilitiesSection = &forms.FormSection{
		Title:       t("disabilities"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.disabilitiesSection)
	return nil
}

func (f *IndividualForm) buildVulnerabilitiesSection(t locales.Translator) error {
	f.vulnerabilitiesSection = &forms.FormSection{
		Title:       t("vulnerabilities"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.vulnerabilitiesSection)
	return nil
}

func (f *IndividualForm) buildDataCollectionSection(t locales.Translator) error {
	f.dataCollectionSection = &forms.FormSection{
		Title:       t("data_collection"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.dataCollectionSection)
	return nil
}

func (f *IndividualForm) buildServiceSection(t locales.Translator) error {
	f.serviceSection = &forms.FormSection{
		Title:       t("services"),
		Fields:      []forms.Field{},
		Collapsible: true,
		Collapsed:   !f.isNew(),
	}
	f.Form.Sections = append(f.Form.Sections, f.serviceSection)
	return nil
}

func (f *IndividualForm) buildIdField(t locales.Translator) error {
	if !f.isNew() {
		return buildField(&forms.IDField{
			Name:        constants.DBColumnIndividualID,
			DisplayName: t("individual_id"),
			QRCodeURL:   fmt.Sprintf("/countries/%s/participants/%s", f.individual.CountryID, f.individual.ID),
		}, f.personalInfoSection, f.individual.ID)
	}
	return nil
}

func (f *IndividualForm) isNew() bool {
	return len(f.individual.ID) == 0
}

func (f *IndividualForm) buildFullName(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualFullName,
		DisplayName: t("full_name"),
	}, f.personalInfoSection, f.individual.FullName)
}

func (f *IndividualForm) buildPreferredName(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualPreferredName,
		DisplayName: t("preferred_name"),
		Value:       f.individual.PreferredName,
	}, f.personalInfoSection, f.individual.PreferredName)
}

func (f *IndividualForm) buildFirstName(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualFirstName,
		DisplayName: t("first_name"),
	}, f.personalInfoSection, f.individual.FirstName)
}

func (f *IndividualForm) buildMiddleName(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualMiddleName,
		DisplayName: t("middle_name"),
	}, f.personalInfoSection, f.individual.MiddleName)
}

func (f *IndividualForm) buildLastName(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualLastName,
		DisplayName: t("last_name"),
	}, f.personalInfoSection, f.individual.LastName)
}

func (f *IndividualForm) buildMothersName(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualMothersName,
		DisplayName: t("mother_name"),
	}, f.personalInfoSection, f.individual.MothersName)
}

func (f *IndividualForm) buildNativeName(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualNativeName,
		DisplayName: t("native_name"),
	}, f.personalInfoSection, f.individual.NativeName)
}

func (f *IndividualForm) buildPrefersToRemainAnonymous(t locales.Translator) error {
	return buildField(&forms.CheckboxInputField{
		Name:        constants.DBColumnIndividualPrefersToRemainAnonymous,
		DisplayName: t("prefers_to_stay_anonymous"),
	}, f.personalInfoSection, f.individual.PrefersToRemainAnonymous)
}

func (f *IndividualForm) buildSex(t locales.Translator) error {
	sexOptions := getSexOptions()
	sexOptions = append([]forms.SelectInputFieldOption{
		{Value: "", Label: t("select_a_value")},
	}, sexOptions...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualSex,
		DisplayName: t("sex"),
		Options:     sexOptions,
		Codec:       &sexCodec{},
	}, f.personalInfoSection, f.individual.Sex)
}

func (f *IndividualForm) buildBirthDate(t locales.Translator) error {
	return buildField(&forms.DateInputField{
		Name:        constants.DBColumnIndividualBirthDate,
		DisplayName: t("birth_date"),
		MinValue:    "1900-01-01",
	}, f.personalInfoSection, f.individual.BirthDate)
}

func (f *IndividualForm) buildAge(t locales.Translator) error {
	return buildField(&forms.NumberInputField{
		Name:        constants.DBColumnIndividualAge,
		DisplayName: t("age"),
	}, f.personalInfoSection, f.individual.Age)
}

func (f *IndividualForm) buildIsMinor(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsMinor,
		DisplayName: t("is_minor"),
	}, f.personalInfoSection, f.individual.IsMinor)
}

func (f *IndividualForm) buildIsChildAtRisk(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsChildAtRisk,
		DisplayName: t("is_child_at_risk"),
	}, f.vulnerabilitiesSection, f.individual.IsChildAtRisk)
}

func (f *IndividualForm) buildIsWomanAtRisk(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsWomanAtRisk,
		DisplayName: t("is_woman_at_risk"),
	}, f.vulnerabilitiesSection, f.individual.IsWomanAtRisk)
}

func (f *IndividualForm) buildIsElderAtRisk(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsElderAtRisk,
		DisplayName: t("is_elder_at_risk"),
	}, f.vulnerabilitiesSection, f.individual.IsElderAtRisk)
}

func (f *IndividualForm) buildIsSingleParent(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsSingleParent,
		DisplayName: t("is_single_parent"),
	}, f.vulnerabilitiesSection, f.individual.IsSingleParent)
}

func (f *IndividualForm) buildIsSeparatedChild(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsSeparatedChild,
		DisplayName: t("is_separated_child"),
	}, f.vulnerabilitiesSection, f.individual.IsSeparatedChild)
}

func (f *IndividualForm) buildIsPregnant(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsPregnant,
		DisplayName: t("is_pregnant"),
	}, f.vulnerabilitiesSection, f.individual.IsPregnant)
}

func (f *IndividualForm) buildIsLactating(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsLactating,
		DisplayName: t("is_lactating"),
	}, f.vulnerabilitiesSection, f.individual.IsLactating)
}

func (f *IndividualForm) buildHasMedicalCondition(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasMedicalCondition,
		DisplayName: t("has_medical_condition"),
	}, f.vulnerabilitiesSection, f.individual.HasMedicalCondition)
}

func (f *IndividualForm) buildNeedsLegalAndPhysicalProtection(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualNeedsLegalAndPhysicalProtection,
		DisplayName: t("needs_legal_and_physical_protection"),
	}, f.vulnerabilitiesSection, f.individual.NeedsLegalAndPhysicalProtection)
}

func (f *IndividualForm) buildNationality1(t locales.Translator) error {
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualNationality1,
		DisplayName: t("nationality_1"),
		Options:     buildCountryOptions(t),
	}, f.personalInfoSection, f.individual.Nationality1)
}

func (f *IndividualForm) buildNationality2(t locales.Translator) error {
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualNationality2,
		DisplayName: t("nationality_2"),
		Options:     buildCountryOptions(t),
	}, f.personalInfoSection, f.individual.Nationality2)
}

func (f *IndividualForm) buildIdentification1Type(t locales.Translator) error {
	options := getIdentificationTypeOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualIdentificationType1,
		DisplayName: t("identification_type_1"),
		Options:     options,
		Codec:       &identificationTypeCodec{},
	}, f.personalInfoSection, f.individual.IdentificationType1)
}

func (f *IndividualForm) buildIdentification1Other(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualIdentificationTypeExplanation1,
		DisplayName: t("if_other_explain"),
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation1)
}

func (f *IndividualForm) buildIdentification1Number(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualIdentificationNumber1,
		DisplayName: t("identification_number_1"),
	}, f.personalInfoSection, f.individual.IdentificationNumber1)
}

func (f *IndividualForm) buildIdentification2Type(t locales.Translator) error {
	options := getIdentificationTypeOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualIdentificationType2,
		Options:     options,
		Codec:       &identificationTypeCodec{},
		DisplayName: t("identification_type_2"),
	}, f.personalInfoSection, f.individual.IdentificationType2)
}

func (f *IndividualForm) buildIdentification2Other(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualIdentificationTypeExplanation2,
		DisplayName: t("if_other_explain"),
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation2)
}

func (f *IndividualForm) buildIdentification2Number(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualIdentificationNumber2,
		DisplayName: t("identification_number_2"),
	}, f.personalInfoSection, f.individual.IdentificationNumber2)
}
func (f *IndividualForm) buildIdentification3Type(t locales.Translator) error {
	options := getIdentificationTypeOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualIdentificationType3,
		DisplayName: t("identification_type_3"),
		Options:     options,
		Codec:       &identificationTypeCodec{},
	}, f.personalInfoSection, f.individual.IdentificationType3)
}

func (f *IndividualForm) buildIdentification3Other(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualIdentificationTypeExplanation3,
		DisplayName: t("if_other_explain"),
	}, f.personalInfoSection, f.individual.IdentificationTypeExplanation3)
}

func (f *IndividualForm) buildIdentification3Number(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualIdentificationNumber3,
		DisplayName: t("identification_number_3"),
	}, f.personalInfoSection, f.individual.IdentificationNumber3)
}

func (f *IndividualForm) buildInternalID(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualInternalID,
		DisplayName: t("internal_id"),
	}, f.personalInfoSection, f.individual.InternalID)
}

func (f *IndividualForm) buildHouseholdID(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualHouseholdID,
		DisplayName: t("household_id"),
	}, f.personalInfoSection, f.individual.HouseholdID)
}

func (f *IndividualForm) buildHouseholdSize(t locales.Translator) error {
	return buildField(&forms.NumberInputField{
		Name:        constants.DBColumnIndividualHouseholdSize,
		DisplayName: t("household_size"),
	}, f.personalInfoSection, f.individual.HouseholdSize)
}

func (f *IndividualForm) buildIsHeadOfHousehold(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsHeadOfHousehold,
		DisplayName: t("is_head_of_household"),
	}, f.personalInfoSection, f.individual.IsHeadOfHousehold)
}

func (f *IndividualForm) buildIsFemaleHeadedHousehold(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsFemaleHeadedHousehold,
		DisplayName: t("is_female_headed_household"),
	}, f.personalInfoSection, f.individual.IsFemaleHeadedHousehold)
}

func (f *IndividualForm) buildIsMinorHeadedHousehold(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsMinorHeadedHousehold,
		DisplayName: t("is_minor_headed_household"),
	}, f.personalInfoSection, f.individual.IsMinorHeadedHousehold)
}

func (f *IndividualForm) buildCommunityID(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualCommunityID,
		DisplayName: t("community_id"),
	}, f.personalInfoSection, f.individual.CommunityID)
}

func (f *IndividualForm) buildCommunitySize(t locales.Translator) error {
	return buildField(&forms.NumberInputField{
		Name:        constants.DBColumnIndividualCommunitySize,
		DisplayName: t("community_size"),
	}, f.personalInfoSection, f.individual.CommunitySize)
}

func (f *IndividualForm) buildIsHeadOfCommunity(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualIsHeadOfCommunity,
		DisplayName: t("is_head_of_community"),
	}, f.personalInfoSection, f.individual.IsHeadOfCommunity)
}

func (f *IndividualForm) buildSpokenLanguage1(t locales.Translator) error {
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualSpokenLanguage1,
		DisplayName: t("spoken_language_1"),
		Options:     buildLanguageOptions(t),
	}, f.personalInfoSection, f.individual.SpokenLanguage1)
}

func (f *IndividualForm) buildSpokenLanguage2(t locales.Translator) error {
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualSpokenLanguage2,
		DisplayName: t("spoken_language_2"),
		Options:     buildLanguageOptions(t),
	}, f.personalInfoSection, f.individual.SpokenLanguage2)
}

func (f *IndividualForm) buildSpokenLanguage3(t locales.Translator) error {
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualSpokenLanguage3,
		DisplayName: t("spoken_language_3"),
		Options:     buildLanguageOptions(t),
	}, f.personalInfoSection, f.individual.SpokenLanguage3)
}

func (f *IndividualForm) buildPreferredCommunicationLanguage(t locales.Translator) error {
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualPreferredCommunicationLanguage,
		DisplayName: t("spoken_language_preferred"),
		Options:     buildLanguageOptions(t),
	}, f.personalInfoSection, f.individual.PreferredCommunicationLanguage)
}

func (f *IndividualForm) buildPhoneNumber1(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualPhoneNumber1,
		DisplayName: t("phone_number_1"),
	}, f.contactInfoSection, f.individual.PhoneNumber1)
}

func (f *IndividualForm) buildPhoneNumber2(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualPhoneNumber2,
		DisplayName: t("phone_number_2"),
	}, f.contactInfoSection, f.individual.PhoneNumber2)
}

func (f *IndividualForm) buildPhoneNumber3(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualPhoneNumber3,
		DisplayName: t("phone_number_3"),
	}, f.contactInfoSection, f.individual.PhoneNumber3)
}

func (f *IndividualForm) buildEmailAddress1(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualEmail1,
		DisplayName: t("email_1"),
	}, f.contactInfoSection, f.individual.Email1)
}

func (f *IndividualForm) buildEmailAddress2(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualEmail2,
		DisplayName: t("email_2"),
	}, f.contactInfoSection, f.individual.Email2)
}

func (f *IndividualForm) buildEmailAddress3(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualEmail3,
		DisplayName: t("email_3"),
	}, f.contactInfoSection, f.individual.Email3)
}

func (f *IndividualForm) buildAddress(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualAddress,
		DisplayName: t("address"),
	}, f.contactInfoSection, f.individual.Address)
}

func (f *IndividualForm) buildPreferredContactMethod(t locales.Translator) error {
	options := getPreferredContactMethodOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualPreferredContactMethod,
		DisplayName: t("preferred_contact_method"),
		Options:     options,
		Codec:       &preferredContactMethodCodec{},
	}, f.contactInfoSection, f.individual.PreferredContactMethod)
}

func (f *IndividualForm) buildContactInstructions(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualPreferredContactMethodComments,
		DisplayName: t("preferred_contact_method_comments"),
	}, f.contactInfoSection, f.individual.PreferredContactMethodComments)
}

func (f *IndividualForm) buildHasConsentedToRgpd(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasConsentedToRGPD,
		DisplayName: t("has_consented_to_rgpd"),
	}, f.consentSection, f.individual.HasConsentedToRGPD)
}

func (f *IndividualForm) buildHasConsentedToReferral(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasConsentedToReferral,
		DisplayName: t("has_consented_to_referral"),
	}, f.consentSection, f.individual.HasConsentedToReferral)
}

func (f *IndividualForm) buildPresentsProtectionConcerns(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualPresentsProtectionConcerns,
		DisplayName: t("presents_protection_concerns"),
	}, f.protectionSection, f.individual.PresentsProtectionConcerns)
}

func (f *IndividualForm) buildDisplacementStatus(t locales.Translator) error {
	options := getDisplacementStatusOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualDisplacementStatus,
		DisplayName: t("displacement_status"),
		Options:     options,
		Codec:       &displacementStatusCodec{},
	}, f.protectionSection, f.individual.DisplacementStatus)
}

func (f *IndividualForm) buildDisplacementStatusComment(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualDisplacementStatusComment,
		DisplayName: t("if_other_explain"),
	}, f.protectionSection, f.individual.DisplacementStatusComment)
}

func (f *IndividualForm) buildHasVisionDisability(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasVisionDisability,
		DisplayName: t("has_vision_disability"),
	}, f.disabilitiesSection, f.individual.HasVisionDisability)
}

func (f *IndividualForm) buildVisionDisabilityLevel(t locales.Translator) error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualVisionDisabilityLevel,
		DisplayName: t("disability_level"),
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.VisionDisabilityLevel)
}

func (f *IndividualForm) buildHasDisability(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasDisability,
		DisplayName: t("has_disability"),
	}, f.disabilitiesSection, f.individual.HasDisability)
}

func (f *IndividualForm) buildPWDComments(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualPWDComments,
		DisplayName: t("pwd_comments"),
	}, f.disabilitiesSection, f.individual.PWDComments)
}

func (f *IndividualForm) buildVulnerabilityComments(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualVulnerabilityComments,
		DisplayName: t("vulnerability_comments"),
	}, f.vulnerabilitiesSection, f.individual.VulnerabilityComments)
}

func (f *IndividualForm) buildHasHearingDisability(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasHearingDisability,
		DisplayName: t("has_hearing_disability"),
	}, f.disabilitiesSection, f.individual.HasHearingDisability)
}

func (f *IndividualForm) buildHearingDisabilityLevel(t locales.Translator) error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualHearingDisabilityLevel,
		DisplayName: t("disability_level"),
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.HearingDisabilityLevel)
}

func (f *IndividualForm) buildHasMobilityDisability(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasMobilityDisability,
		DisplayName: t("has_mobility_disability"),
	}, f.disabilitiesSection, f.individual.HasMobilityDisability)
}

func (f *IndividualForm) buildMobilityDisabilityLevel(t locales.Translator) error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualMobilityDisabilityLevel,
		DisplayName: t("disability_level"),
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.MobilityDisabilityLevel)
}

func (f *IndividualForm) buildHasCognitiveDisability(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasCognitiveDisability,
		DisplayName: t("has_cognitive_disability"),
	}, f.disabilitiesSection, f.individual.HasCognitiveDisability)
}

func (f *IndividualForm) buildCognitiveDisabilityLevel(t locales.Translator) error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualCognitiveDisabilityLevel,
		DisplayName: t("disability_level"),
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.CognitiveDisabilityLevel)
}

func (f *IndividualForm) buildHasSelfCareDisability(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasSelfCareDisability,
		DisplayName: t("has_self_care_disability"),
	}, f.disabilitiesSection, f.individual.HasSelfCareDisability)
}

func (f *IndividualForm) buildSelfCareDisabilityLevel(t locales.Translator) error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualSelfCareDisabilityLevel,
		DisplayName: t("disability_level"),
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.SelfCareDisabilityLevel)
}

func (f *IndividualForm) buildHasCommunicationDisability(t locales.Translator) error {
	return buildField(&forms.OptionalBooleanInputField{
		Name:        constants.DBColumnIndividualHasCommunicationDisability,
		DisplayName: t("has_communication_disability"),
	}, f.disabilitiesSection, f.individual.HasCommunicationDisability)
}

func (f *IndividualForm) buildCommunicationDisabilityLevel(t locales.Translator) error {
	options := getDisabilityLevels()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualCommunicationDisabilityLevel,
		DisplayName: t("disability_level"),
		Options:     options,
		Codec:       &disabilityLevelCodec{},
	}, f.disabilitiesSection, f.individual.CommunicationDisabilityLevel)
}

func (f *IndividualForm) buildEngagementContext(t locales.Translator) error {
	options := getEngagementContextOptions()
	options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
	return buildField(&forms.SelectInputField{
		Name:        constants.DBColumnIndividualEngagementContext,
		DisplayName: t("engagement_context"),
		Options:     options,
		Codec:       &engagementContextCodec{},
	}, f.dataCollectionSection, f.individual.EngagementContext)
}

func (f *IndividualForm) buildCollectionAgent(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualCollectionAgentName,
		DisplayName: t("collection_agent_name"),
	}, f.dataCollectionSection, f.individual.CollectionAgentName)
}

func (f *IndividualForm) buildCollectionAgentTitle(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualCollectionAgentTitle,
		DisplayName: t("collection_agent_title"),
	}, f.dataCollectionSection, f.individual.CollectionAgentTitle)
}

func (f *IndividualForm) buildCollectionDate(t locales.Translator) error {
	return buildField(&forms.DateInputField{
		Name:        constants.DBColumnIndividualCollectionTime,
		DisplayName: t("collection_time"),
	}, f.dataCollectionSection, f.individual.CollectionTime)
}

func (f *IndividualForm) buildCollectionLocation1(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualCollectionAdministrativeArea1,
		DisplayName: t("collection_area_1"),
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea1)
}

func (f *IndividualForm) buildCollectionLocation2(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualCollectionAdministrativeArea2,
		DisplayName: t("collection_area_2"),
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea2)
}

func (f *IndividualForm) buildCollectionLocation3(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualCollectionAdministrativeArea3,
		DisplayName: t("collection_area_3"),
	}, f.dataCollectionSection, f.individual.CollectionAdministrativeArea3)
}

func (f *IndividualForm) buildCollectionOffice(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualCollectionOffice,
		DisplayName: t("collection_office"),
	}, f.dataCollectionSection, f.individual.CollectionOffice)
}

func (f *IndividualForm) buildComments(t locales.Translator) error {
	return buildField(&forms.TextAreaInputField{
		Name:        constants.DBColumnIndividualComments,
		DisplayName: t("comments"),
	}, f.dataCollectionSection, f.individual.Comments)
}

func (f *IndividualForm) buildFreeField1(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualFreeField1,
		DisplayName: t("free_field_1"),
	}, f.dataCollectionSection, f.individual.FreeField1)
}

func (f *IndividualForm) buildFreeField2(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualFreeField2,
		DisplayName: t("free_field_2"),
	}, f.dataCollectionSection, f.individual.FreeField2)
}

func (f *IndividualForm) buildFreeField3(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualFreeField3,
		DisplayName: t("free_field_3"),
	}, f.dataCollectionSection, f.individual.FreeField3)
}

func (f *IndividualForm) buildFreeField4(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualFreeField4,
		DisplayName: t("free_field_4"),
	}, f.dataCollectionSection, f.individual.FreeField4)
}

func (f *IndividualForm) buildFreeField5(t locales.Translator) error {
	return buildField(&forms.TextInputField{
		Name:        constants.DBColumnIndividualFreeField5,
		DisplayName: t("free_field_5"),
	}, f.dataCollectionSection, f.individual.FreeField5)
}

func (f *IndividualForm) buildServiceCC(idx int) func(t locales.Translator) error {
	return func(t locales.Translator) error {
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
		options = append([]forms.SelectInputFieldOption{{Value: "", Label: t("select_a_value")}}, options...)
		return buildField(&forms.SelectInputField{
			Name:        fmt.Sprintf("service_cc_%d", idx),
			DisplayName: t("service_cc_abrv_no", idx),
			Options:     options,
			Codec:       &serviceCCCodec{},
		}, f.serviceSection, value)
	}
}

func (f *IndividualForm) buildServiceRequestedDate(idx int) func(t locales.Translator) error {
	return func(t locales.Translator) error {
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
			Name:        fmt.Sprintf("service_requested_date_%d", idx),
			DisplayName: t("service_requested_date_no", idx),
		}, f.serviceSection, value)
	}
}

func (f *IndividualForm) buildServiceDeliveredDate(idx int) func(t locales.Translator) error {
	return func(t locales.Translator) error {
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
			Name:        fmt.Sprintf("service_delivered_date_%d", idx),
			DisplayName: t("service_delivery_date_no", idx),
		}, f.serviceSection, value)
	}
}

func (f *IndividualForm) buildServiceComments(idx int) func(t locales.Translator) error {
	return func(t locales.Translator) error {
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
			Name:        fmt.Sprintf("service_comments_%d", idx),
			DisplayName: t("service_delivery_date_comments_no", idx),
		}, f.serviceSection, value)
	}
}

func (f *IndividualForm) buildServiceTextInputField(fieldName string, fieldValue interface{}, translationKey string, idx int) func(t locales.Translator) error {
	return func(t locales.Translator) error {
		return buildField(&forms.TextInputField{
			Name:        fieldName,
			DisplayName: t(translationKey, idx),
		}, f.serviceSection, fieldValue)
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

func buildCountryOptions(t locales.Translator) []forms.SelectInputFieldOption {
	var opts = make([]forms.SelectInputFieldOption, 0, len(constants.Countries))
	opts = append(opts, forms.SelectInputFieldOption{
		Value: "",
		Label: t("select_a_value"),
	})
	for _, country := range constants.Countries {
		opts = append(opts, forms.SelectInputFieldOption{
			Value: country.ISO3166Alpha3,
			Label: country.Name,
		})
	}
	return opts
}

func buildLanguageOptions(t locales.Translator) []forms.SelectInputFieldOption {
	var opts = make([]forms.SelectInputFieldOption, 0, len(constants.Languages))
	opts = append(opts, forms.SelectInputFieldOption{
		Value: "",
		Label: t("select_a_value"),
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
