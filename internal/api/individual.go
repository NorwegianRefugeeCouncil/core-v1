package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
)

type Individual struct {
	// Address is the residence address of the individual
	Address string `json:"address" db:"address"`
	// Age is the age of the individual
	Age *int `json:"age" db:"age"`
	// BirthDate is the date of birth of the individual.
	BirthDate *time.Time `json:"birthDate" db:"birth_date"`
	// CognitiveDisabilityLevel is the cognitive disability level of the individual
	CognitiveDisabilityLevel DisabilityLevel `json:"cognitiveDisabilityLevel" db:"cognitive_disability_level"`
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
	CommunicationDisabilityLevel DisabilityLevel `json:"communicationDisabilityLevel" db:"communication_disability_level"`
	// CommunityID is the ID of the community the individual belongs to
	CommunityID string `json:"communityId" db:"community_id"`
	// CountryID is the ID of the country the individual belongs to
	CountryID string `json:"countryId" db:"country_id"`
	// CreatedAt is the time the individual record was created in the database
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	// DeletedAt is the time the individual record was deleted from the database
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
	// DisplacementStatus is the displacement status of the individual
	DisplacementStatus DisplacementStatus `json:"displacementStatus" db:"displacement_status"`
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
	Sex Sex `json:"sex" db:"sex"`
	// HasCognitiveDisability is true if the individual has a cognitive disability
	HasCognitiveDisability bool `json:"hasCognitiveDisability" db:"has_cognitive_disability"`
	// HasCommunicationDisability is true if the individual has a communication disability
	HasCommunicationDisability bool `json:"hasCommunicationDisability" db:"has_communication_disability"`
	// HasConsentedToRGPD is a flag indicating whether the individual has consented to the RGPD
	// (General Data Protection Regulation)
	HasConsentedToRGPD bool `json:"hasConsentedToRgpd" db:"has_consented_to_rgpd"`
	// HasConsentedToReferral is a flag indicating whether the individual has consented to be referred
	// to internal or external services
	HasConsentedToReferral bool `json:"hasConsentedToReferral" db:"has_consented_to_referral"`
	// HasHearingDisability is true if the individual has a hearing disability
	HasHearingDisability bool `json:"hasHearingDisability" db:"has_hearing_disability"`
	// HasMobilityDisability is true if the individual has a mobility disability
	HasMobilityDisability bool `json:"hasMobilityDisability" db:"has_mobility_disability"`
	// HasSelfCareDisability is true if the individual has a self care disability
	HasSelfCareDisability bool `json:"hasSelfCareDisability" db:"has_selfcare_disability"`
	// HasVisionDisability is true if the individual has a vision disability
	HasVisionDisability bool `json:"hasVisionDisability" db:"has_vision_disability"`
	// HearingDisabilityLevel is the hearing disability level of the individual
	HearingDisabilityLevel DisabilityLevel `json:"hearingDisabilityLevel" db:"hearing_disability_level"`
	// HouseholdID is the ID of the household the individual belongs to
	HouseholdID string `json:"householdId" db:"household_id"`
	// ID is the ID of the individual
	ID string `json:"id" db:"id"`
	// IdentificationType1 is the type of primary identification of the individual
	IdentificationType1 string `json:"identificationType1" db:"identification_type_1"`
	// IdentificationTypeExplanation1 is the explanation of the primary identification type of the individual
	// when the primary identification type is "other"
	IdentificationTypeExplanation1 string `json:"identificationTypeExplanation1" db:"identification_type_explanation_1"`
	// IdentificationNumber1 is the primary identification number of the individual
	IdentificationNumber1 string `json:"identificationNumber1" db:"identification_number_1"`
	// IdentificationType2 is the type of secondary identification of the individual
	IdentificationType2 string `json:"identificationType2" db:"identification_type_2"`
	// IdentificationTypeExplanation2 is the explanation of the secondary identification type of the individual
	// when the secondary identification type is "other"
	IdentificationTypeExplanation2 string `json:"identificationTypeExplanation2" db:"identification_type_explanation_2"`
	// IdentificationNumber2 is the secondary identification number of the individual
	IdentificationNumber2 string `json:"identificationNumber2" db:"identification_number_2"`
	// IdentificationType3 is the type of tertiary identification of the individual
	IdentificationType3 string `json:"identificationType3" db:"identification_type_3"`
	// IdentificationTypeExplanation3 is the explanation of the tertiary identification type of the individual
	// when the tertiary identification type is "other"
	IdentificationTypeExplanation3 string `json:"identificationTypeExplanation3" db:"identification_type_explanation_3"`
	// IdentificationNumber3 is the tertiary identification number of the individual
	IdentificationNumber3 string `json:"identificationNumber3" db:"identification_number_3"`
	// EngagementContext is the context of the identification of the individual
	EngagementContext string `json:"engagementContext" db:"engagement_context"`
	// InternalID is the internal ID of the individual. This is used to link individuals across different
	// systems.
	InternalID string `json:"internalId" db:"internal_id"`
	// IsHeadOfCommunity is a flag indicating whether the individual is the head of the community
	IsHeadOfCommunity bool `json:"isHeadOfCommunity" db:"is_head_of_community"`
	// IsHeadOfHousehold is a flag indicating whether the individual is the head of the household
	IsHeadOfHousehold bool `json:"isHeadOfHousehold" db:"is_head_of_household"`
	// IsMinor is a flag indicating whether the individual is a minor
	IsMinor bool `json:"isMinor" db:"is_minor"`
	// MobilityDisabilityLevel is the mobility disability level of the individual
	MobilityDisabilityLevel DisabilityLevel `json:"mobilityDisabilityLevel" db:"mobility_disability_level"`
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
	PreferredContactMethod string `json:"preferredContactMethod" db:"preferred_contact_method"`
	// PreferredContactMethodComments is the comments on the preferred contact method of the individual
	PreferredContactMethodComments string `json:"preferredContactMethodComments" db:"preferred_contact_method_comments"`
	// PreferredName is the preferred name of the individual
	PreferredName string `json:"preferredName" db:"preferred_name"`
	// PreferredCommunicationLanguage is the preferred communication language of the individual
	PreferredCommunicationLanguage string `json:"preferredCommunicationLanguage" db:"preferred_communication_language"`
	// PrefersToRemainAnonymous is a flag indicating whether the individual prefers to remain anonymous
	PrefersToRemainAnonymous bool `json:"prefersToRemainAnonymous" db:"prefers_to_remain_anonymous"`
	// PresentsProtectionConcerns is a flag indicating whether the individual presents protection concerns
	PresentsProtectionConcerns bool `json:"presentsProtectionConcerns" db:"presents_protection_concerns"`
	// SelfCareDisabilityLevel is the self-care disability level of the individual
	SelfCareDisabilityLevel DisabilityLevel `json:"selfCareDisabilityLevel" db:"selfcare_disability_level"`
	// SpokenLanguage1 is the primary spoken language of the individual
	SpokenLanguage1 string `json:"spokenLanguage1" db:"spoken_language_1"`
	// SpokenLanguage2 is the secondary spoken language of the individual
	SpokenLanguage2 string `json:"spokenLanguage2" db:"spoken_language_2"`
	// SpokenLanguage3 is the tertiary spoken language of the individual
	SpokenLanguage3 string `json:"spokenLanguage3" db:"spoken_language_3"`
	// UpdatedAt is the time the individual was last updated in the database
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	// VisionDisabilityLevel is the vision disability level of the individual
	VisionDisabilityLevel DisabilityLevel `json:"visionDisabilityLevel" db:"vision_disability_level"`
}

type IndividualList struct {
	Items []*Individual `json:"items"`
}

func (i *Individual) GetFieldValue(field string) (interface{}, error) {
	switch field {
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
	case constants.DBColumnIndividualSex:
		return i.Sex, nil
	case constants.DBColumnIndividualHasCognitiveDisability:
		return i.HasCognitiveDisability, nil
	case constants.DBColumnIndividualHasCommunicationDisability:
		return i.HasCommunicationDisability, nil
	case constants.DBColumnIndividualHasConsentedToRGPD:
		return i.HasConsentedToRGPD, nil
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
	case constants.DBColumnIndividualIsMinor:
		return i.IsMinor, nil
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
	default:
		return nil, fmt.Errorf("unknown field: %s", field)
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
	i.DisplacementStatus = DisplacementStatus(trimString(string(i.DisplacementStatus)))
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
	i.HouseholdID = trimString(i.HouseholdID)
	i.ID = trimString(i.ID)
	i.EngagementContext = trimString(i.EngagementContext)
	i.IdentificationNumber1 = trimString(i.IdentificationNumber1)
	i.IdentificationNumber2 = trimString(i.IdentificationNumber2)
	i.IdentificationNumber3 = trimString(i.IdentificationNumber3)
	i.IdentificationType1 = trimString(i.IdentificationType1)
	i.IdentificationType2 = trimString(i.IdentificationType2)
	i.IdentificationType3 = trimString(i.IdentificationType3)
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
	i.PreferredContactMethod = trimString(i.PreferredContactMethod)
	i.PreferredContactMethodComments = trimString(i.PreferredContactMethodComments)
	i.PreferredName = trimString(i.PreferredName)
	if i.PreferredName == "" {
		i.PreferredName = i.FullName
	}
	i.PreferredCommunicationLanguage = trimString(i.PreferredCommunicationLanguage)
	i.SpokenLanguage1 = trimString(i.SpokenLanguage1)
	i.SpokenLanguage2 = trimString(i.SpokenLanguage2)
	i.SpokenLanguage3 = trimString(i.SpokenLanguage3)

	if i.PrefersToRemainAnonymous {
		i.FullName = ""
		i.PreferredName = ""
	}
}
