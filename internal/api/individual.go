package api

import (
	"encoding/json"
	"errors"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/locales"
	"net/mail"
	"strconv"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
)

type Individual struct {
	// Inactive is true if the individual is inactive
	Inactive bool `json:"inactive" db:"inactive"`
	// Address is the residence address of the individual
	Address string `json:"address" db:"address"`
	// Age is the age of the individual
	Age *int `json:"age" db:"age"`
	// BirthDate is the date of birth of the individual.
	BirthDate *time.Time `json:"birthDate" db:"birth_date"`
	// CognitiveDisabilityLevel is the cognitive disability level of the individual
	CognitiveDisabilityLevel enumTypes.DisabilityLevel `json:"cognitiveDisabilityLevel" db:"cognitive_disability_level"`
	// CollectionAdministrativeArea1 is the first administrative area of the collection
	// For example, in the case of a case in the Democratic Republic of Congo, this would be the province
	CollectionAdministrativeArea1 string `json:"collectionAdministrativeArea1" db:"collection_administrative_area_1"`
	// CollectionAdministrativeArea2 is the second administrative area of the collection
	// For example, in the case of a case in the Democratic Republic of Congo, this would be the territory
	CollectionAdministrativeArea2 string `json:"collectionAdministrativeArea2" db:"collection_administrative_area_2"`
	// CollectionAdministrativeArea3 is the third administrative area of the collection
	// For example, in the case of a case in the Democratic Republic of Congo, this would be the commune
	CollectionAdministrativeArea3 string `json:"collectionAdministrativeArea3" db:"collection_administrative_area_3"`
	// CollectionOffice is the office of the collection
	CollectionOffice string `json:"collectionOffice" db:"collection_office"`
	// CollectionAgentName is the ID of the agent who collected the data
	CollectionAgentName string `json:"collectionAgentName" db:"collection_agent_name"`
	// CollectionAgentTitle is the title of the agent who collected the data
	CollectionAgentTitle string `json:"collectionAgentTitle" db:"collection_agent_title"`
	// CollectionTime is the date/time the data was collected
	CollectionTime time.Time `json:"collectionTime" db:"collection_time"`
	// Comments is a free text field for comments
	Comments string `json:"comments" db:"comments"`
	// CommunicationDisabilityLevel is the communication disability level of the individual
	CommunicationDisabilityLevel enumTypes.DisabilityLevel `json:"communicationDisabilityLevel" db:"communication_disability_level"`
	// CommunityID is the ID of the community the individual belongs to
	CommunityID string `json:"communityId" db:"community_id"`
	// CommunitySize is the size of the community the individual belongs to
	CommunitySize *int `json:"communitySize" db:"community_size"`
	// CountryID is the ID of the country the individual belongs to
	CountryID string `json:"countryId" db:"country_id"`
	// CreatedAt is the time the individual record was created in the database
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	// DeletedAt is the time the individual record was deleted from the database
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
	// DisplacementStatus is the displacement status of the individual
	DisplacementStatus enumTypes.DisplacementStatus `json:"displacementStatus" db:"displacement_status"`
	// DisplacementStatusComment is a comment about the displacement status of the individual
	DisplacementStatusComment string `json:"displacementStatusComment" db:"displacement_status_comment"`
	// Email1 is the email address 1 of the individual
	Email1 string `json:"email1" db:"email_1"`
	// Email2 is the email address 2 of the individual
	Email2 string `json:"email2" db:"email_2"`
	// Email 3 is the email address 3 of the individual
	Email3 string `json:"email3" db:"email_3"`
	// FullName is the full name of the individual
	FullName string `json:"fullName" db:"full_name"`
	// FirstName is the first name of the individual
	FirstName string `json:"firstName" db:"first_name"`
	// MiddleName is the middle name of the individual
	MiddleName string `json:"middleName" db:"middle_name"`
	// MothersName is the name of the individuals mother
	MothersName string `json:"mothersName" db:"mothers_name"`
	// LastName is the last name of the individual
	LastName string `json:"lastName" db:"last_name"`
	// NativeName is the name of the individual in their native language
	NativeName string `json:"nativeName" db:"native_name"`
	// FreeField1 is a free field for the individual
	FreeField1 string `json:"freeField1" db:"free_field_1"`
	// FreeField2 is a free field for the individual
	FreeField2 string `json:"freeField2" db:"free_field_2"`
	// FreeField3 is a free field for the individual
	FreeField3 string `json:"freeField3" db:"free_field_3"`
	// FreeField4 is a free field for the individual
	FreeField4 string `json:"freeField4" db:"free_field_4"`
	// FreeField5 is a free field for the individual
	FreeField5 string `json:"freeField5" db:"free_field_5"`
	// Sex is the sex of the individual
	Sex enumTypes.Sex `json:"sex" db:"sex"`
	// HasCognitiveDisability is true if the individual has a cognitive disability
	HasCognitiveDisability *bool `json:"hasCognitiveDisability" db:"has_cognitive_disability"`
	// HasCommunicationDisability is true if the individual has a communication disability
	HasCommunicationDisability *bool `json:"hasCommunicationDisability" db:"has_communication_disability"`
	// HasConsentedToRGPD is a flag indicating whether the individual has consented to the RGPD
	// (General Data Protection Regulation)
	HasConsentedToRGPD *bool `json:"hasConsentedToRgpd" db:"has_consented_to_rgpd"`
	// HasConsentedToReferral is a flag indicating whether the individual has consented to be referred
	// to internal or external services
	HasConsentedToReferral *bool `json:"hasConsentedToReferral" db:"has_consented_to_referral"`
	// HasDisability is true if the individual has a disability
	HasDisability *bool `json:"hasDisability" db:"has_disability"`
	// HasHearingDisability is true if the individual has a hearing disability
	HasHearingDisability *bool `json:"hasHearingDisability" db:"has_hearing_disability"`
	// HasMobilityDisability is true if the individual has a mobility disability
	HasMobilityDisability *bool `json:"hasMobilityDisability" db:"has_mobility_disability"`
	// HasSelfCareDisability is true if the individual has a self care disability
	HasSelfCareDisability *bool `json:"hasSelfCareDisability" db:"has_selfcare_disability"`
	// HasVisionDisability is true if the individual has a vision disability
	HasVisionDisability *bool `json:"hasVisionDisability" db:"has_vision_disability"`
	// HearingDisabilityLevel is the hearing disability level of the individual
	HearingDisabilityLevel enumTypes.DisabilityLevel `json:"hearingDisabilityLevel" db:"hearing_disability_level"`
	// HouseholdID is the ID of the household the individual belongs to
	HouseholdID string `json:"householdId" db:"household_id"`
	// HouseholdSize is the size of the household the individual belongs to
	HouseholdSize *int `json:"householdSize" db:"household_size"`
	// ID is the ID of the individual
	ID string `json:"id" db:"id"`
	// IdentificationType1 is the type of primary identification of the individual
	IdentificationType1 enumTypes.IdentificationType `json:"identificationType1" db:"identification_type_1"`
	// IdentificationTypeExplanation1 is the explanation of the primary identification type of the individual
	// when the primary identification type is "other"
	IdentificationTypeExplanation1 string `json:"identificationTypeExplanation1" db:"identification_type_explanation_1"`
	// IdentificationNumber1 is the primary identification number of the individual
	IdentificationNumber1 string `json:"identificationNumber1" db:"identification_number_1"`
	// IdentificationType2 is the type of secondary identification of the individual
	IdentificationType2 enumTypes.IdentificationType `json:"identificationType2" db:"identification_type_2"`
	// IdentificationTypeExplanation2 is the explanation of the secondary identification type of the individual
	// when the secondary identification type is "other"
	IdentificationTypeExplanation2 string `json:"identificationTypeExplanation2" db:"identification_type_explanation_2"`
	// IdentificationNumber2 is the secondary identification number of the individual
	IdentificationNumber2 string `json:"identificationNumber2" db:"identification_number_2"`
	// IdentificationType3 is the type of tertiary identification of the individual
	IdentificationType3 enumTypes.IdentificationType `json:"identificationType3" db:"identification_type_3"`
	// IdentificationTypeExplanation3 is the explanation of the tertiary identification type of the individual
	// when the tertiary identification type is "other"
	IdentificationTypeExplanation3 string `json:"identificationTypeExplanation3" db:"identification_type_explanation_3"`
	// IdentificationNumber3 is the tertiary identification number of the individual
	IdentificationNumber3 string `json:"identificationNumber3" db:"identification_number_3"`
	// EngagementContext is the context of the identification of the individual
	EngagementContext enumTypes.EngagementContext `json:"engagementContext" db:"engagement_context"`
	// InternalID is the internal ID of the individual. This is used to link individuals across different
	// systems.
	InternalID string `json:"internalId" db:"internal_id"`
	// IsHeadOfCommunity is a flag indicating whether the individual is the head of the community
	IsHeadOfCommunity *bool `json:"isHeadOfCommunity" db:"is_head_of_community"`
	// IsHeadOfHousehold is a flag indicating whether the individual is the head of the household
	IsHeadOfHousehold *bool `json:"isHeadOfHousehold" db:"is_head_of_household"`
	// IsFemaleHeadedHousehold is a flag indicating whether the head of the household is female
	IsFemaleHeadedHousehold *bool `json:"isFemaleHeadedHousehold" db:"is_female_headed_household"`
	// IsMinorHeadedHousehold is a flag indicating whether the head of the household is a minor
	IsMinorHeadedHousehold *bool `json:"isMinorHeadedHousehold" db:"is_minor_headed_household"`
	// IsMinor is a flag indicating whether the individual is a minor
	IsMinor *bool `json:"isMinor" db:"is_minor"`
	// NeedsLegalAndPhysicalProtection is a flag indicating whether the individual needs legal and physical protection
	NeedsLegalAndPhysicalProtection *bool `json:"needsLegalAndPhysicalProtection" db:"needs_legal_and_physical_protection"`
	// HasMedicalCondition is a flag indicating whether the individual has a medical condition
	HasMedicalCondition *bool `json:"hasMedicalCondition" db:"has_medical_condition"`
	// IsPregnant is a flag indicating whether the individual is pregnant
	IsPregnant *bool `json:"isPregnant" db:"is_pregnant"`
	// IsLactating is a flag indicating whether the individual is lactating
	IsLactating *bool `json:"isLactating" db:"is_lactating"`
	// IsWomanAtRisk is a flag indicating whether the individual is a woman at risk
	IsWomanAtRisk *bool `json:"isWomanAtRisk" db:"is_woman_at_risk"`
	// IsChildAtRisk is a flag indicating whether the individual is a child at risk
	IsChildAtRisk *bool `json:"isChildAtRisk" db:"is_child_at_risk"`
	// IsElderAtRisk is a flag indicating whether the individual is an elder at risk
	IsElderAtRisk *bool `json:"isElderAtRisk" db:"is_elder_at_risk"`
	// IsSingleParent is a flag indicating whether the individual is a single parent
	IsSingleParent *bool `json:"isSingleParent" db:"is_single_parent"`
	// IsSeparatedChild is a flag indicating whether the individual is a separated child
	IsSeparatedChild *bool `json:"isSeparatedChild" db:"is_separated_child"`
	// MobilityDisabilityLevel is the mobility disability level of the individual
	MobilityDisabilityLevel enumTypes.DisabilityLevel `json:"mobilityDisabilityLevel" db:"mobility_disability_level"`
	// Nationality1 is the primary nationality of the individual
	Nationality1 string `json:"nationality1" db:"nationality_1"`
	// Nationality2 is the secondary nationality of the individual
	Nationality2 string `json:"nationality2" db:"nationality_2"`
	// NormalizedPhoneNumber1 is the normalized phone number of the individual
	// It is used for search purposes
	// TODO: do not expose this field on the api. This is a database concern only
	NormalizedPhoneNumber1 string `json:"-" db:"normalized_phone_number_1"`
	NormalizedPhoneNumber2 string `json:"-" db:"normalized_phone_number_2"`
	NormalizedPhoneNumber3 string `json:"-" db:"normalized_phone_number_3"`
	// PhoneNumber1 is the phone number 1 of the individual
	PhoneNumber1 string `json:"phoneNumber1" db:"phone_number_1"`
	// PhoneNumber2 is the phone number 2 of the individual
	PhoneNumber2 string `json:"phoneNumber2" db:"phone_number_2"`
	// PhoneNumber3 is the phone number 3 of the individual
	PhoneNumber3 string `json:"phoneNumber3" db:"phone_number_3"`
	// PreferredContactMethod is the preferred contact method of the individual
	PreferredContactMethod enumTypes.ContactMethod `json:"preferredContactMethod" db:"preferred_contact_method"`
	// PreferredContactMethodComments is the comments on the preferred contact method of the individual
	PreferredContactMethodComments string `json:"preferredContactMethodComments" db:"preferred_contact_method_comments"`
	// PreferredName is the preferred name of the individual
	PreferredName string `json:"preferredName" db:"preferred_name"`
	// PreferredCommunicationLanguage is the preferred communication language of the individual
	PreferredCommunicationLanguage string `json:"preferredCommunicationLanguage" db:"preferred_communication_language"`
	// PrefersToRemainAnonymous is a flag indicating whether the individual prefers to remain anonymous
	PrefersToRemainAnonymous *bool `json:"prefersToRemainAnonymous" db:"prefers_to_remain_anonymous"`
	// PresentsProtectionConcerns is a flag indicating whether the individual presents protection concerns
	PresentsProtectionConcerns *bool `json:"presentsProtectionConcerns" db:"presents_protection_concerns"`
	// PWDComments is the comments on the disability status of the individual
	PWDComments string `json:"pwdComments" db:"pwd_comments"`
	// SelfCareDisabilityLevel is the self-care disability level of the individual
	SelfCareDisabilityLevel enumTypes.DisabilityLevel `json:"selfCareDisabilityLevel" db:"selfcare_disability_level"`
	// SpokenLanguage1 is the primary spoken language of the individual
	SpokenLanguage1 string `json:"spokenLanguage1" db:"spoken_language_1"`
	// SpokenLanguage2 is the secondary spoken language of the individual
	SpokenLanguage2 string `json:"spokenLanguage2" db:"spoken_language_2"`
	// SpokenLanguage3 is the tertiary spoken language of the individual
	SpokenLanguage3 string `json:"spokenLanguage3" db:"spoken_language_3"`
	// UpdatedAt is the time the individual was last updated in the database
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	// VisionDisabilityLevel is the vision disability level of the individual
	VisionDisabilityLevel enumTypes.DisabilityLevel `json:"visionDisabilityLevel" db:"vision_disability_level"`
	// VulnerabilityComments is the comments on the vulnerabilities of the individual
	VulnerabilityComments string `json:"vulnerabilityComments" db:"vulnerability_comments"`

	ServiceCC1            enumTypes.ServiceCC `json:"serviceCC1" db:"service_cc_1"`
	ServiceRequestedDate1 *time.Time          `json:"serviceRequestedDate1" db:"service_requested_date_1"`
	ServiceDeliveredDate1 *time.Time          `json:"serviceDeliveredDate1" db:"service_delivered_date_1"`
	ServiceComments1      string              `json:"serviceComments1" db:"service_comments_1"`

	ServiceCC2            enumTypes.ServiceCC `json:"serviceCC2" db:"service_cc_2"`
	ServiceRequestedDate2 *time.Time          `json:"serviceRequestedDate2" db:"service_requested_date_2"`
	ServiceDeliveredDate2 *time.Time          `json:"serviceDeliveredDate2" db:"service_delivered_date_2"`
	ServiceComments2      string              `json:"serviceComments2" db:"service_comments_2"`

	ServiceCC3            enumTypes.ServiceCC `json:"serviceCC3" db:"service_cc_3"`
	ServiceRequestedDate3 *time.Time          `json:"serviceRequestedDate3" db:"service_requested_date_3"`
	ServiceDeliveredDate3 *time.Time          `json:"serviceDeliveredDate3" db:"service_delivered_date_3"`
	ServiceComments3      string              `json:"serviceComments3" db:"service_comments_3"`

	ServiceCC4            enumTypes.ServiceCC `json:"serviceCC4" db:"service_cc_4"`
	ServiceRequestedDate4 *time.Time          `json:"serviceRequestedDate4" db:"service_requested_date_4"`
	ServiceDeliveredDate4 *time.Time          `json:"serviceDeliveredDate4" db:"service_delivered_date_4"`
	ServiceComments4      string              `json:"serviceComments4" db:"service_comments_4"`

	ServiceCC5            enumTypes.ServiceCC `json:"serviceCC5" db:"service_cc_5"`
	ServiceRequestedDate5 *time.Time          `json:"serviceRequestedDate5" db:"service_requested_date_5"`
	ServiceDeliveredDate5 *time.Time          `json:"serviceDeliveredDate5" db:"service_delivered_date_5"`
	ServiceComments5      string              `json:"serviceComments5" db:"service_comments_5"`

	ServiceCC6            enumTypes.ServiceCC `json:"serviceCC6" db:"service_cc_6"`
	ServiceRequestedDate6 *time.Time          `json:"serviceRequestedDate6" db:"service_requested_date_6"`
	ServiceDeliveredDate6 *time.Time          `json:"serviceDeliveredDate6" db:"service_delivered_date_6"`
	ServiceComments6      string              `json:"serviceComments6" db:"service_comments_6"`

	ServiceCC7            enumTypes.ServiceCC `json:"serviceCC7" db:"service_cc_7"`
	ServiceRequestedDate7 *time.Time          `json:"serviceRequestedDate7" db:"service_requested_date_7"`
	ServiceDeliveredDate7 *time.Time          `json:"serviceDeliveredDate7" db:"service_delivered_date_7"`
	ServiceComments7      string              `json:"serviceComments7" db:"service_comments_7"`
}

type IndividualList struct {
	Items []*Individual `json:"items"`
}

func (i *Individual) GetFieldValue(field string) (interface{}, error) {
	switch field {
	case constants.DBColumnIndividualInactive:
		return i.Inactive, nil
	case constants.DBColumnIndividualAddress:
		return i.Address, nil
	case constants.DBColumnIndividualAge:
		return i.Age, nil
	case constants.DBColumnIndividualBirthDate:
		return i.BirthDate, nil
	case constants.DBColumnIndividualCollectionAdministrativeArea1:
		return i.CollectionAdministrativeArea1, nil
	case constants.DBColumnIndividualCollectionAdministrativeArea2:
		return i.CollectionAdministrativeArea2, nil
	case constants.DBColumnIndividualCollectionAdministrativeArea3:
		return i.CollectionAdministrativeArea3, nil
	case constants.DBColumnIndividualCollectionOffice:
		return i.CollectionOffice, nil
	case constants.DBColumnIndividualCollectionAgentName:
		return i.CollectionAgentName, nil
	case constants.DBColumnIndividualCollectionAgentTitle:
		return i.CollectionAgentTitle, nil
	case constants.DBColumnIndividualCollectionTime:
		return i.CollectionTime, nil
	case constants.DBColumnIndividualComments:
		return i.Comments, nil
	case constants.DBColumnIndividualCommunityID:
		return i.CommunityID, nil
	case constants.DBColumnIndividualCommunitySize:
		return i.CommunitySize, nil
	case constants.DBColumnIndividualCountryID:
		return i.CountryID, nil
	case constants.DBColumnIndividualCreatedAt:
		return i.CreatedAt, nil
	case constants.DBColumnIndividualCognitiveDisabilityLevel:
		return i.CognitiveDisabilityLevel, nil
	case constants.DBColumnIndividualCommunicationDisabilityLevel:
		return i.CommunicationDisabilityLevel, nil
	case constants.DBColumnIndividualHearingDisabilityLevel:
		return i.HearingDisabilityLevel, nil
	case constants.DBColumnIndividualMobilityDisabilityLevel:
		return i.MobilityDisabilityLevel, nil
	case constants.DBColumnIndividualSelfCareDisabilityLevel:
		return i.SelfCareDisabilityLevel, nil
	case constants.DBColumnIndividualVisionDisabilityLevel:
		return i.VisionDisabilityLevel, nil
	case constants.DBColumnIndividualDisplacementStatus:
		return i.DisplacementStatus, nil
	case constants.DBColumnIndividualDisplacementStatusComment:
		return i.DisplacementStatusComment, nil
	case constants.DBColumnIndividualEmail1:
		return i.Email1, nil
	case constants.DBColumnIndividualEmail2:
		return i.Email2, nil
	case constants.DBColumnIndividualEmail3:
		return i.Email3, nil
	case constants.DBColumnIndividualFullName:
		return i.FullName, nil
	case constants.DBColumnIndividualFirstName:
		return i.FirstName, nil
	case constants.DBColumnIndividualMiddleName:
		return i.MiddleName, nil
	case constants.DBColumnIndividualLastName:
		return i.LastName, nil
	case constants.DBColumnIndividualNativeName:
		return i.NativeName, nil
	case constants.DBColumnIndividualMothersName:
		return i.MothersName, nil
	case constants.DBColumnIndividualSex:
		return i.Sex, nil
	case constants.DBColumnIndividualHasCognitiveDisability:
		return i.HasCognitiveDisability, nil
	case constants.DBColumnIndividualHasCommunicationDisability:
		return i.HasCommunicationDisability, nil
	case constants.DBColumnIndividualHasConsentedToRGPD:
		return i.HasConsentedToRGPD, nil
	case constants.DBColumnIndividualHasDisability:
		return i.HasDisability, nil
	case constants.DBColumnIndividualHasHearingDisability:
		return i.HasHearingDisability, nil
	case constants.DBColumnIndividualHasMobilityDisability:
		return i.HasMobilityDisability, nil
	case constants.DBColumnIndividualHasSelfCareDisability:
		return i.HasSelfCareDisability, nil
	case constants.DBColumnIndividualHasVisionDisability:
		return i.HasVisionDisability, nil
	case constants.DBColumnIndividualHasConsentedToReferral:
		return i.HasConsentedToReferral, nil
	case constants.DBColumnIndividualHouseholdID:
		return i.HouseholdID, nil
	case constants.DBColumnIndividualHouseholdSize:
		return i.HouseholdSize, nil
	case constants.DBColumnIndividualID:
		return i.ID, nil
	case constants.DBColumnIndividualIdentificationType1:
		return i.IdentificationType1, nil
	case constants.DBColumnIndividualIdentificationTypeExplanation1:
		return i.IdentificationTypeExplanation1, nil
	case constants.DBColumnIndividualIdentificationNumber1:
		return i.IdentificationNumber1, nil
	case constants.DBColumnIndividualIdentificationType2:
		return i.IdentificationType2, nil
	case constants.DBColumnIndividualIdentificationTypeExplanation2:
		return i.IdentificationTypeExplanation2, nil
	case constants.DBColumnIndividualIdentificationNumber2:
		return i.IdentificationNumber2, nil
	case constants.DBColumnIndividualIdentificationType3:
		return i.IdentificationType3, nil
	case constants.DBColumnIndividualIdentificationTypeExplanation3:
		return i.IdentificationTypeExplanation3, nil
	case constants.DBColumnIndividualIdentificationNumber3:
		return i.IdentificationNumber3, nil
	case constants.DBColumnIndividualEngagementContext:
		return i.EngagementContext, nil
	case constants.DBColumnIndividualInternalID:
		return i.InternalID, nil
	case constants.DBColumnIndividualIsHeadOfCommunity:
		return i.IsHeadOfCommunity, nil
	case constants.DBColumnIndividualIsHeadOfHousehold:
		return i.IsHeadOfHousehold, nil
	case constants.DBColumnIndividualIsFemaleHeadedHousehold:
		return i.IsFemaleHeadedHousehold, nil
	case constants.DBColumnIndividualIsMinorHeadedHousehold:
		return i.IsMinorHeadedHousehold, nil
	case constants.DBColumnIndividualIsMinor:
		return i.IsMinor, nil
	case constants.DBColumnIndividualIsChildAtRisk:
		return i.IsChildAtRisk, nil
	case constants.DBColumnIndividualIsElderAtRisk:
		return i.IsElderAtRisk, nil
	case constants.DBColumnIndividualIsWomanAtRisk:
		return i.IsWomanAtRisk, nil
	case constants.DBColumnIndividualIsLactating:
		return i.IsLactating, nil
	case constants.DBColumnIndividualIsPregnant:
		return i.IsPregnant, nil
	case constants.DBColumnIndividualIsSeparatedChild:
		return i.IsSeparatedChild, nil
	case constants.DBColumnIndividualIsSingleParent:
		return i.IsSingleParent, nil
	case constants.DBColumnIndividualHasMedicalCondition:
		return i.HasMedicalCondition, nil
	case constants.DBColumnIndividualNeedsLegalAndPhysicalProtection:
		return i.NeedsLegalAndPhysicalProtection, nil
	case constants.DBColumnIndividualNationality1:
		return i.Nationality1, nil
	case constants.DBColumnIndividualNationality2:
		return i.Nationality2, nil
	case constants.DBColumnIndividualNormalizedPhoneNumber1:
		return i.NormalizedPhoneNumber1, nil
	case constants.DBColumnIndividualPhoneNumber1:
		return i.PhoneNumber1, nil
	case constants.DBColumnIndividualNormalizedPhoneNumber2:
		return i.NormalizedPhoneNumber2, nil
	case constants.DBColumnIndividualPhoneNumber2:
		return i.PhoneNumber2, nil
	case constants.DBColumnIndividualNormalizedPhoneNumber3:
		return i.NormalizedPhoneNumber3, nil
	case constants.DBColumnIndividualPhoneNumber3:
		return i.PhoneNumber3, nil
	case constants.DBColumnIndividualPreferredContactMethod:
		return i.PreferredContactMethod, nil
	case constants.DBColumnIndividualPreferredContactMethodComments:
		return i.PreferredContactMethodComments, nil
	case constants.DBColumnIndividualPreferredName:
		return i.PreferredName, nil
	case constants.DBColumnIndividualPreferredCommunicationLanguage:
		return i.PreferredCommunicationLanguage, nil
	case constants.DBColumnIndividualPrefersToRemainAnonymous:
		return i.PrefersToRemainAnonymous, nil
	case constants.DBColumnIndividualPresentsProtectionConcerns:
		return i.PresentsProtectionConcerns, nil
	case constants.DBColumnIndividualPWDComments:
		return i.PWDComments, nil
	case constants.DBColumnIndividualVulnerabilityComments:
		return i.VulnerabilityComments, nil
	case constants.DBColumnIndividualSpokenLanguage1:
		return i.SpokenLanguage1, nil
	case constants.DBColumnIndividualSpokenLanguage2:
		return i.SpokenLanguage2, nil
	case constants.DBColumnIndividualSpokenLanguage3:
		return i.SpokenLanguage3, nil
	case constants.DBColumnIndividualUpdatedAt:
		return i.UpdatedAt, nil
	case constants.DBColumnIndividualFreeField1:
		return i.FreeField1, nil
	case constants.DBColumnIndividualFreeField2:
		return i.FreeField2, nil
	case constants.DBColumnIndividualFreeField3:
		return i.FreeField3, nil
	case constants.DBColumnIndividualFreeField4:
		return i.FreeField4, nil
	case constants.DBColumnIndividualFreeField5:
		return i.FreeField5, nil
	case constants.DBColumnIndividualServiceCC1:
		return i.ServiceCC1, nil
	case constants.DBColumnIndividualServiceRequestedDate1:
		return i.ServiceRequestedDate1, nil
	case constants.DBColumnIndividualServiceDeliveredDate1:
		return i.ServiceDeliveredDate1, nil
	case constants.DBColumnIndividualServiceComments1:
		return i.ServiceComments1, nil
	case constants.DBColumnIndividualServiceCC2:
		return i.ServiceCC2, nil
	case constants.DBColumnIndividualServiceRequestedDate2:
		return i.ServiceRequestedDate2, nil
	case constants.DBColumnIndividualServiceDeliveredDate2:
		return i.ServiceDeliveredDate2, nil
	case constants.DBColumnIndividualServiceComments2:
		return i.ServiceComments2, nil
	case constants.DBColumnIndividualServiceCC3:
		return i.ServiceCC3, nil
	case constants.DBColumnIndividualServiceRequestedDate3:
		return i.ServiceRequestedDate3, nil
	case constants.DBColumnIndividualServiceDeliveredDate3:
		return i.ServiceDeliveredDate3, nil
	case constants.DBColumnIndividualServiceComments3:
		return i.ServiceComments3, nil
	case constants.DBColumnIndividualServiceCC4:
		return i.ServiceCC4, nil
	case constants.DBColumnIndividualServiceRequestedDate4:
		return i.ServiceRequestedDate4, nil
	case constants.DBColumnIndividualServiceDeliveredDate4:
		return i.ServiceDeliveredDate4, nil
	case constants.DBColumnIndividualServiceComments4:
		return i.ServiceComments4, nil
	case constants.DBColumnIndividualServiceCC5:
		return i.ServiceCC5, nil
	case constants.DBColumnIndividualServiceRequestedDate5:
		return i.ServiceRequestedDate5, nil
	case constants.DBColumnIndividualServiceDeliveredDate5:
		return i.ServiceDeliveredDate5, nil
	case constants.DBColumnIndividualServiceComments5:
		return i.ServiceComments5, nil
	case constants.DBColumnIndividualServiceCC6:
		return i.ServiceCC6, nil
	case constants.DBColumnIndividualServiceRequestedDate6:
		return i.ServiceRequestedDate6, nil
	case constants.DBColumnIndividualServiceDeliveredDate6:
		return i.ServiceDeliveredDate6, nil
	case constants.DBColumnIndividualServiceComments6:
		return i.ServiceComments6, nil
	case constants.DBColumnIndividualServiceCC7:
		return i.ServiceCC7, nil
	case constants.DBColumnIndividualServiceRequestedDate7:
		return i.ServiceRequestedDate7, nil
	case constants.DBColumnIndividualServiceDeliveredDate7:
		return i.ServiceDeliveredDate7, nil
	case constants.DBColumnIndividualServiceComments7:
		return i.ServiceComments7, nil
	default:
		return nil, errors.New(locales.GetTranslator()("error_unknown_field", field))
	}
}

func (i *Individual) String() string {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func (i *Individual) Normalize() {
	i.Address = trimString(i.Address)
	i.CollectionAdministrativeArea1 = trimString(i.CollectionAdministrativeArea1)
	i.CollectionAdministrativeArea2 = trimString(i.CollectionAdministrativeArea2)
	i.CollectionAdministrativeArea3 = trimString(i.CollectionAdministrativeArea3)
	i.CollectionOffice = trimString(i.CollectionOffice)
	i.CollectionAgentName = trimString(i.CollectionAgentName)
	i.Comments = trimString(i.Comments)
	i.CommunityID = trimString(i.CommunityID)
	i.CountryID = trimString(i.CountryID)
	i.DisplacementStatus = enumTypes.DisplacementStatus(trimString(string(i.DisplacementStatus)))
	i.DisplacementStatusComment = trimString(i.DisplacementStatusComment)
	i.Email1 = normalizeEmail(i.Email1)
	i.Email2 = normalizeEmail(i.Email2)
	i.Email3 = normalizeEmail(i.Email3)
	i.FreeField1 = trimString(i.FreeField1)
	i.FreeField2 = trimString(i.FreeField2)
	i.FreeField3 = trimString(i.FreeField3)
	i.FreeField4 = trimString(i.FreeField4)
	i.FreeField5 = trimString(i.FreeField5)
	i.FullName = trimString(i.FullName)
	i.FirstName = trimString(i.FirstName)
	i.MiddleName = trimString(i.MiddleName)
	i.LastName = trimString(i.LastName)
	i.NativeName = trimString(i.NativeName)
	i.MothersName = trimString(i.MothersName)
	i.HouseholdID = trimString(i.HouseholdID)
	i.ID = trimString(i.ID)
	i.EngagementContext = enumTypes.EngagementContext(trimString(string(i.EngagementContext)))
	i.IdentificationNumber1 = trimString(i.IdentificationNumber1)
	i.IdentificationNumber2 = trimString(i.IdentificationNumber2)
	i.IdentificationNumber3 = trimString(i.IdentificationNumber3)
	i.IdentificationType1 = enumTypes.IdentificationType(trimString(string(i.IdentificationType1)))
	i.IdentificationType2 = enumTypes.IdentificationType(trimString(string(i.IdentificationType2)))
	i.IdentificationType3 = enumTypes.IdentificationType(trimString(string(i.IdentificationType3)))
	i.IdentificationTypeExplanation1 = trimString(i.IdentificationTypeExplanation1)
	i.IdentificationTypeExplanation2 = trimString(i.IdentificationTypeExplanation2)
	i.IdentificationTypeExplanation3 = trimString(i.IdentificationTypeExplanation3)
	i.InternalID = trimString(i.InternalID)
	i.Nationality1 = trimString(i.Nationality1)
	i.Nationality2 = trimString(i.Nationality2)

	// not alphabetically sorted because phone numbers need to be trimmed
	// before storing NormalizedPhoneNumbers
	i.PhoneNumber1 = trimString(i.PhoneNumber1)
	i.PhoneNumber2 = trimString(i.PhoneNumber2)
	i.PhoneNumber3 = trimString(i.PhoneNumber3)

	i.NormalizedPhoneNumber1 = NormalizePhoneNumber(i.PhoneNumber1)
	i.NormalizedPhoneNumber2 = NormalizePhoneNumber(i.PhoneNumber2)
	i.NormalizedPhoneNumber3 = NormalizePhoneNumber(i.PhoneNumber3)
	i.PreferredContactMethod = enumTypes.ContactMethod(trimString(string(i.PreferredContactMethod)))
	i.PreferredContactMethodComments = trimString(i.PreferredContactMethodComments)
	i.PreferredName = trimString(i.PreferredName)
	if i.PreferredName == "" {
		i.PreferredName = i.FullName
	}
	i.PreferredCommunicationLanguage = trimString(i.PreferredCommunicationLanguage)
	i.PWDComments = trimString(i.PWDComments)
	i.VulnerabilityComments = trimString(i.VulnerabilityComments)
	i.SpokenLanguage1 = trimString(i.SpokenLanguage1)
	i.SpokenLanguage2 = trimString(i.SpokenLanguage2)
	i.SpokenLanguage3 = trimString(i.SpokenLanguage3)

	i.ServiceCC1 = enumTypes.ServiceCC(trimString(string(i.ServiceCC1)))
	i.ServiceComments1 = trimString(i.ServiceComments1)
	i.ServiceCC2 = enumTypes.ServiceCC(trimString(string(i.ServiceCC2)))
	i.ServiceComments2 = trimString(i.ServiceComments2)
	i.ServiceCC3 = enumTypes.ServiceCC(trimString(string(i.ServiceCC3)))
	i.ServiceComments3 = trimString(i.ServiceComments3)
	i.ServiceCC4 = enumTypes.ServiceCC(trimString(string(i.ServiceCC4)))
	i.ServiceComments4 = trimString(i.ServiceComments4)
	i.ServiceCC5 = enumTypes.ServiceCC(trimString(string(i.ServiceCC5)))
	i.ServiceComments5 = trimString(i.ServiceComments5)
	i.ServiceCC6 = enumTypes.ServiceCC(trimString(string(i.ServiceCC6)))
	i.ServiceComments6 = trimString(i.ServiceComments6)
	i.ServiceCC7 = enumTypes.ServiceCC(trimString(string(i.ServiceCC7)))
	i.ServiceComments7 = trimString(i.ServiceComments7)

	if i.PrefersToRemainAnonymous != nil && *i.PrefersToRemainAnonymous == true {
		i.FullName = ""
		i.PreferredName = ""
		i.FirstName = ""
		i.MiddleName = ""
		i.LastName = ""
		i.NativeName = ""
	}
}

func (i *Individual) MarshalTabularData() ([]string, error) {
	row := make([]string, len(constants.IndividualFileColumns))
	for j, col := range constants.IndividualFileColumns {
		field, ok := constants.IndividualDBToFileMap[col]
		if !ok {
			return nil, fmt.Errorf("unknown column %s", col) // should not happen but we never know.
		}
		value, err := i.GetFieldValue(field)
		if err != nil {
			return nil, err
		}

		switch v := value.(type) {
		case bool:
			row[j] = strconv.FormatBool(v)
		case *bool:
			if v != nil {
				row[j] = strconv.FormatBool(*v)
			}
		case enumTypes.OptionalBoolean:
			row[j] = string(v)
		case int:
			row[j] = strconv.Itoa(v)
		case *int:
			if v != nil {
				row[j] = strconv.Itoa(*value.(*int))
			}
		case string:
			if (field == constants.DBColumnIndividualNationality1 || field == constants.DBColumnIndividualNationality2) && v != "" {
				row[j] = constants.CountriesByCode[v].Name
				break
			}
			if (field == constants.DBColumnIndividualSpokenLanguage1 || field == constants.DBColumnIndividualSpokenLanguage2 || field == constants.DBColumnIndividualSpokenLanguage3 || field == constants.DBColumnIndividualPreferredCommunicationLanguage) && v != "" {
				row[j] = constants.LanguagesByCode[v].Name
				break
			}
			row[j] = v
		case time.Time:
			row[j] = v.Format(getTimeFormatForField(field))
		case *time.Time:
			if v != nil {
				row[j] = v.Format(getTimeFormatForField(field))
			}
		case enumTypes.DisabilityLevel:
			row[j] = string(v)
		case enumTypes.DisplacementStatus:
			row[j] = string(v)
		case enumTypes.ServiceCC:
			row[j] = string(v)
		case enumTypes.Sex:
			row[j] = string(v)
		default:
			row[j] = fmt.Sprintf("%v", v)
		}
	}
	return row, nil
}

func (i *Individual) UnmarshalTabularData(colMapping map[string]int, cols []string) []error {
	var errors []error
	if len(cols) <= len(colMapping) {
		filler := make([]string, len(colMapping)-len(cols))
		cols = append(cols, filler...)
	}
	for field, idx := range colMapping {
		switch field {
		case constants.FileColumnIndividualID:
			i.ID = cols[idx]
		case constants.FileColumnIndividualInactive:
			i.Inactive = isExplicitlyTrue(cols[idx])
		case constants.FileColumnIndividualAddress:
			i.Address = cols[idx]
		case constants.FileColumnIndividualAge:
			age, err := ParseAge(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.Age = age
		case constants.FileColumnIndividualBirthDate:
			var birthDate *time.Time
			birthDate, err := ParseBirthdate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.BirthDate = birthDate
		case constants.FileColumnIndividualCognitiveDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualCognitiveDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.CognitiveDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualCollectionAdministrativeArea1:
			i.CollectionAdministrativeArea1 = cols[idx]
		case constants.FileColumnIndividualCollectionAdministrativeArea2:
			i.CollectionAdministrativeArea2 = cols[idx]
		case constants.FileColumnIndividualCollectionAdministrativeArea3:
			i.CollectionAdministrativeArea3 = cols[idx]
		case constants.FileColumnIndividualCollectionOffice:
			i.CollectionOffice = cols[idx]
		case constants.FileColumnIndividualCollectionAgentName:
			i.CollectionAgentName = cols[idx]
		case constants.FileColumnIndividualCollectionAgentTitle:
			i.CollectionAgentTitle = cols[idx]
		case constants.FileColumnIndividualComments:
			i.Comments = cols[idx]
		case constants.FileColumnIndividualCollectionTime:
			var collectionTime *time.Time
			collectionTime, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualCollectionTime, err))
				break
			}
			if collectionTime != nil {
				i.CollectionTime = *collectionTime
			}
		case constants.FileColumnIndividualCommunicationDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualCommunicationDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.CommunicationDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualCommunityID:
			i.CommunityID = cols[idx]
		case constants.FileColumnIndividualCommunitySize:
			var communitySizeStr = cols[idx]
			if communitySizeStr == "" {
				continue
			}
			communitySize, err := strconv.Atoi(communitySizeStr)
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.CommunitySize = &communitySize
		case constants.FileColumnIndividualDisplacementStatus:
			displacementStatus, err := enumTypes.ParseDisplacementStatus(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualDisplacementStatus, err, enumTypes.AllDisplacementStatuses().String()))
				break
			}
			i.DisplacementStatus = displacementStatus
		case constants.FileColumnIndividualDisplacementStatusComment:
			i.DisplacementStatusComment = cols[idx]
		case constants.FileColumnIndividualEmail1:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail1, err))
					break
				}
				i.Email1 = email.Address
			}
		case constants.FileColumnIndividualEmail2:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail2, err))
					break
				}
				i.Email2 = email.Address
			}
		case constants.FileColumnIndividualEmail3:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail3, err))
					break
				}
				i.Email3 = email.Address
			}
		case constants.FileColumnIndividualFullName:
			i.FullName = cols[idx]
		case constants.FileColumnIndividualFirstName:
			i.FirstName = cols[idx]
		case constants.FileColumnIndividualMiddleName:
			i.MiddleName = cols[idx]
		case constants.FileColumnIndividualLastName:
			i.LastName = cols[idx]
		case constants.FileColumnIndividualNativeName:
			i.NativeName = cols[idx]
		case constants.FileColumnIndividualMothersName:
			i.MothersName = cols[idx]
		case constants.FileColumnIndividualFreeField1:
			i.FreeField1 = cols[idx]
		case constants.FileColumnIndividualFreeField2:
			i.FreeField2 = cols[idx]
		case constants.FileColumnIndividualFreeField3:
			i.FreeField3 = cols[idx]
		case constants.FileColumnIndividualFreeField4:
			i.FreeField4 = cols[idx]
		case constants.FileColumnIndividualFreeField5:
			i.FreeField5 = cols[idx]
		case constants.FileColumnIndividualSex:
			sex, err := enumTypes.ParseSex(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualSex, err, enumTypes.AllSexes().String()))
				break
			}
			i.Sex = sex
		case constants.FileColumnIndividualHasMedicalCondition:
			hasMedicalCondition, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasMedicalCondition = hasMedicalCondition.BoolPtr()
		case constants.FileColumnIndividualNeedsLegalAndPhysicalProtection:
			needsLegalAndPhysicalProtection, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.NeedsLegalAndPhysicalProtection = needsLegalAndPhysicalProtection.BoolPtr()
		case constants.FileColumnIndividualIsChildAtRisk:
			isChildAtRisk, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsChildAtRisk = isChildAtRisk.BoolPtr()
		case constants.FileColumnIndividualIsWomanAtRisk:
			isWomanAtRisk, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsWomanAtRisk = isWomanAtRisk.BoolPtr()
		case constants.FileColumnIndividualIsElderAtRisk:
			isElderAtRisk, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsElderAtRisk = isElderAtRisk.BoolPtr()
		case constants.FileColumnIndividualIsLactating:
			isLactating, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsLactating = isLactating.BoolPtr()
		case constants.FileColumnIndividualIsPregnant:
			isPregnant, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsPregnant = isPregnant.BoolPtr()
		case constants.FileColumnIndividualIsSingleParent:
			isSingleParent, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsSingleParent = isSingleParent.BoolPtr()
		case constants.FileColumnIndividualIsSeparatedChild:
			isSeparatedChild, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsSeparatedChild = isSeparatedChild.BoolPtr()
		case constants.FileColumnIndividualHasCognitiveDisability:
			hasCognitiveDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasCognitiveDisability = hasCognitiveDisability.BoolPtr()
		case constants.FileColumnIndividualHasCommunicationDisability:
			hasCommunicationDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasCommunicationDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasCommunicationDisability = hasCommunicationDisability.BoolPtr()
		case constants.FileColumnIndividualHasConsentedToRGPD:
			hasConsentedToRGPD, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasConsentedToRGPD, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasConsentedToRGPD = hasConsentedToRGPD.BoolPtr()
		case constants.FileColumnIndividualHasConsentedToReferral:
			hasConsentedToReferral, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasConsentedToReferral, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasConsentedToReferral = hasConsentedToReferral.BoolPtr()
		case constants.FileColumnIndividualHasDisability:
			hasDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasDisability = hasDisability.BoolPtr()
		case constants.FileColumnIndividualHasHearingDisability:
			hasHearingDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasHearingDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasHearingDisability = hasHearingDisability.BoolPtr()
		case constants.FileColumnIndividualHasMobilityDisability:
			hasMobilityDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasMobilityDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasMobilityDisability = hasMobilityDisability.BoolPtr()
		case constants.FileColumnIndividualHasSelfCareDisability:
			hasSelfCareDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasSelfCareDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasSelfCareDisability = hasSelfCareDisability.BoolPtr()
		case constants.FileColumnIndividualHasVisionDisability:
			hasVisionDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHasVisionDisability, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.HasVisionDisability = hasVisionDisability.BoolPtr()
		case constants.FileColumnIndividualHearingDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHearingDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.HearingDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualHouseholdID:
			i.HouseholdID = cols[idx]
		case constants.FileColumnIndividualHouseholdSize:
			var householdSizeStr = cols[idx]
			if householdSizeStr == "" {
				continue
			}
			householdSize, err := strconv.Atoi(householdSizeStr)
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.HouseholdSize = &householdSize
		case constants.FileColumnIndividualIdentificationType1:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIdentificationType1, err, enumTypes.AllIdentificationTypes().String()))
				break
			}
			i.IdentificationType1 = identificationType
		case constants.FileColumnIndividualIdentificationTypeExplanation1:
			i.IdentificationTypeExplanation1 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber1:
			i.IdentificationNumber1 = cols[idx]
		case constants.FileColumnIndividualIdentificationType2:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIdentificationType2, err, enumTypes.AllIdentificationTypes().String()))
				break
			}
			i.IdentificationType2 = identificationType
		case constants.FileColumnIndividualIdentificationTypeExplanation2:
			i.IdentificationTypeExplanation2 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber2:
			i.IdentificationNumber2 = cols[idx]
		case constants.FileColumnIndividualIdentificationType3:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIdentificationType3, err, enumTypes.AllIdentificationTypes().String()))
				break
			}
			i.IdentificationType3 = identificationType
		case constants.FileColumnIndividualIdentificationTypeExplanation3:
			i.IdentificationTypeExplanation3 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber3:
			i.IdentificationNumber3 = cols[idx]
		case constants.FileColumnIndividualEngagementContext:
			engagementContext, err := enumTypes.ParseEngagementContext(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualEngagementContext, err, enumTypes.AllEngagementContexts().String()))
				break
			}
			i.EngagementContext = engagementContext
		case constants.FileColumnIndividualInternalID:
			i.InternalID = cols[idx]
		case constants.FileColumnIndividualIsHeadOfCommunity:
			isHeadOfCommunity, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIsHeadOfCommunity, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsHeadOfCommunity = isHeadOfCommunity.BoolPtr()
		case constants.FileColumnIndividualIsHeadOfHousehold:
			isHeadOfHousehold, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIsHeadOfHousehold, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsHeadOfHousehold = isHeadOfHousehold.BoolPtr()
		case constants.FileColumnIndividualIsFemaleHeadedHousehold:
			isFemaleHeadedHousehold, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIsFemaleHeadedHousehold, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsFemaleHeadedHousehold = isFemaleHeadedHousehold.BoolPtr()
		case constants.FileColumnIndividualIsMinorHeadedHousehold:
			isMinorHeadedHousehold, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIsMinorHeadedHousehold, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsMinorHeadedHousehold = isMinorHeadedHousehold.BoolPtr()
		case constants.FileColumnIndividualIsMinor:
			isMinor, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIsMinor, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.IsMinor = isMinor.BoolPtr()
		case constants.FileColumnIndividualMobilityDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualMobilityDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.MobilityDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualNationality1:
			if cols[idx] != "" {
				if c := constants.CountriesByCode[cols[idx]].Name; c != "" {
					i.Nationality1 = cols[idx]
				} else if c := constants.CountriesByName[cols[idx]].Name; c != "" {
					i.Nationality1 = constants.CountriesByName[cols[idx]].ISO3166Alpha3
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\". valid values adhere to the ISO3166Alpha3 norm", constants.FileColumnIndividualNationality1, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualNationality2:
			if cols[idx] != "" {
				if c := constants.CountriesByCode[cols[idx]].Name; c != "" {
					i.Nationality2 = cols[idx]
				} else if c := constants.CountriesByName[cols[idx]].Name; c != "" {
					i.Nationality2 = constants.CountriesByName[cols[idx]].ISO3166Alpha3
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\". valid values adhere to the ISO3166Alpha3 norm", constants.FileColumnIndividualNationality2, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualPhoneNumber1:
			i.PhoneNumber1 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber2:
			i.PhoneNumber2 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber3:
			i.PhoneNumber3 = cols[idx]
		case constants.FileColumnIndividualPreferredContactMethod:
			preferredContactMethod, err := enumTypes.ParseContactMethod(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualPreferredContactMethod, err, enumTypes.AllContactMethods().String()))
				break
			}
			i.PreferredContactMethod = preferredContactMethod
		case constants.FileColumnIndividualPreferredContactMethodComments:
			i.PreferredContactMethodComments = cols[idx]
		case constants.FileColumnIndividualPreferredName:
			i.PreferredName = cols[idx]
		case constants.FileColumnIndividualPreferredCommunicationLanguage:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.PreferredCommunicationLanguage = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.PreferredCommunicationLanguage = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualPreferredCommunicationLanguage, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualPrefersToRemainAnonymous:
			prefersToRemainAnonymous, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualPrefersToRemainAnonymous, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.PrefersToRemainAnonymous = prefersToRemainAnonymous.BoolPtr()
		case constants.FileColumnIndividualPresentsProtectionConcerns:
			presentsProtectionConcerns, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualPresentsProtectionConcerns, err, enumTypes.AllOptionalBooleans().String()))
			}
			i.PresentsProtectionConcerns = presentsProtectionConcerns.BoolPtr()
		case constants.FileColumnIndividualPWDComments:
			i.PWDComments = cols[idx]
		case constants.FileColumnIndividualVulnerabilityComments:
			i.VulnerabilityComments = cols[idx]
		case constants.FileColumnIndividualSelfCareDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualSelfCareDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.SelfCareDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualSpokenLanguage1:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage1 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage1 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualSpokenLanguage1, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualSpokenLanguage2:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage2 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage2 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualSpokenLanguage2, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualSpokenLanguage3:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage3 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage3 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualSpokenLanguage3, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualVisionDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualVisionDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.VisionDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualServiceCC1:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC1, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC1 = cc
		case constants.FileColumnIndividualServiceRequestedDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate1, err))
				break
			}
			i.ServiceRequestedDate1 = date
		case constants.FileColumnIndividualServiceDeliveredDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate1, err))
				break
			}
			i.ServiceDeliveredDate1 = date
		case constants.FileColumnIndividualServiceComments1:
			i.ServiceComments1 = cols[idx]
		case constants.FileColumnIndividualServiceCC2:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC2, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC2 = cc
		case constants.FileColumnIndividualServiceRequestedDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate2, err))
				break
			}
			i.ServiceRequestedDate2 = date
		case constants.FileColumnIndividualServiceDeliveredDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate2, err))
				break
			}
			i.ServiceDeliveredDate2 = date
		case constants.FileColumnIndividualServiceComments2:
			i.ServiceComments2 = cols[idx]
		case constants.FileColumnIndividualServiceCC3:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC3, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC3 = cc
		case constants.FileColumnIndividualServiceRequestedDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate3, err))
				break
			}
			i.ServiceRequestedDate3 = date
		case constants.FileColumnIndividualServiceDeliveredDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate3, err))
				break
			}
			i.ServiceDeliveredDate3 = date
		case constants.FileColumnIndividualServiceComments3:
			i.ServiceComments3 = cols[idx]
		case constants.FileColumnIndividualServiceCC4:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC4, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC4 = cc
		case constants.FileColumnIndividualServiceRequestedDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate4, err))
				break
			}
			i.ServiceRequestedDate4 = date
		case constants.FileColumnIndividualServiceDeliveredDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate4, err))
				break
			}
			i.ServiceDeliveredDate4 = date
		case constants.FileColumnIndividualServiceComments4:
			i.ServiceComments4 = cols[idx]
		case constants.FileColumnIndividualServiceCC5:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC5, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC5 = cc
		case constants.FileColumnIndividualServiceRequestedDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate5, err))
				break
			}
			i.ServiceRequestedDate5 = date
		case constants.FileColumnIndividualServiceDeliveredDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate5, err))
				break
			}
			i.ServiceDeliveredDate5 = date
		case constants.FileColumnIndividualServiceComments5:
			i.ServiceComments5 = cols[idx]
		case constants.FileColumnIndividualServiceCC6:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC6, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC6 = cc
		case constants.FileColumnIndividualServiceRequestedDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate6, err))
				break
			}
			i.ServiceRequestedDate6 = date
		case constants.FileColumnIndividualServiceDeliveredDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate6, err))
				break
			}
			i.ServiceDeliveredDate6 = date
		case constants.FileColumnIndividualServiceComments6:
			i.ServiceComments6 = cols[idx]
		case constants.FileColumnIndividualServiceCC7:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC7, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC7 = cc
		case constants.FileColumnIndividualServiceRequestedDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate7, err))
				break
			}
			i.ServiceRequestedDate7 = date
		case constants.FileColumnIndividualServiceDeliveredDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate7, err))
				break
			}
			i.ServiceDeliveredDate7 = date
		case constants.FileColumnIndividualServiceComments7:
			i.ServiceComments7 = cols[idx]
		}
	}
	if len(errors) > 0 {
		return errors
	}
	i.Normalize()
	return nil
}
