package constants

import "github.com/nrc-no/notcore/internal/containers"

const (
	FormParamIndividualAddress                        = "Address"
	FormParamIndividualAge                            = "Age"
	FormParamIndividualBirthDate                      = "BirthDate"
	FormParamIndividualCognitiveDisabilityLevel       = "CognitiveDisabilityLevel"
	FormParamIndividualCollectionAdministrativeArea1  = "CollectionAdministrativeArea1"
	FormParamIndividualCollectionAdministrativeArea2  = "CollectionAdministrativeArea2"
	FormParamIndividualCollectionAdministrativeArea3  = "CollectionAdministrativeArea3 "
	FormParamIndividualCollectionAgentID              = "CollectionAgentName "
	FormParamIndividualCollectionTime                 = "CollectionTime"
	FormParamIndividualCommunicationDisabilityLevel   = "CommunicationDisabilityLevel"
	FormParamIndividualCommunityID                    = "CommunityID"
	FormParamIndividualCountryID                      = "Country"
	FormParamIndividualDisplacementStatus             = "DisplacementStatus"
	FormParamIndividualEmail1                         = "Email1"
	FormParamIndividualEmail2                         = "Email2"
	FormParamIndividualEmail3                         = "Email3"
	FormParamIndividualFullName                       = "FullName"
	FormParamIndividualGender                         = "Gender"
	FormParamIndividualHasCognitiveDisability         = "HasCognitiveDisability"
	FormParamIndividualHasCommunicationDisability     = "HasCommunicationDisability"
	FormParamIndividualHasConsentedToRGPD             = "HasConsentedToRGPD"
	FormParamIndividualHasConsentedToReferral         = "HasConsentedToReferral"
	FormParamIndividualHasHearingDisability           = "HasHearingDisability"
	FormParamIndividualHasMobilityDisability          = "HasMobilityDisability"
	FormParamIndividualHasSelfCareDisability          = "HasSelfCareDisability"
	FormParamIndividualHasVisionDisability            = "HasVisionDisability"
	FormParamIndividualHearingDisabilityLevel         = "HearingDisabilityLevel"
	FormParamIndividualHouseholdID                    = "HouseholdID"
	FormParamIndividualID                             = "ID"
	FormParamIndividualIdentificationContext          = "IdentificationContext"
	FormParamIndividualIdentificationNumber1          = "IdentificationNumber1"
	FormParamIndividualIdentificationNumber2          = "IdentificationNumber2"
	FormParamIndividualIdentificationNumber3          = "IdentificationNumber3"
	FormParamIndividualIdentificationType1            = "IdentificationType1"
	FormParamIndividualIdentificationType2            = "IdentificationType2"
	FormParamIndividualIdentificationType3            = "IdentificationType3"
	FormParamIndividualIdentificationTypeExplanation1 = "IdentificationTypeExplanation1"
	FormParamIndividualIdentificationTypeExplanation2 = "IdentificationTypeExplanation2"
	FormParamIndividualIdentificationTypeExplanation3 = "IdentificationTypeExplanation3"
	FormParamIndividualInternalID                     = "InternalID"
	FormParamIndividualIsActive                       = "IsActive"
	FormParamIndividualIsHeadOfCommunity              = "IsHeadOfCommunity"
	FormParamIndividualIsHeadOfHousehold              = "IsHeadOfHousehold"
	FormParamIndividualIsMinor                        = "IsMinor"
	FormParamIndividualMobilityDisabilityLevel        = "MobilityDisabilityLevel"
	FormParamIndividualNationality1                   = "Nationality1"
	FormParamIndividualNationality2                   = "Nationality2"
	FormParamIndividualPhoneNumber                    = "PhoneNumber1"
	FormParamIndividualPreferredCommunicationLanguage = "PreferredCommunicationLanguage"
	FormParamIndividualPreferredContactMethod         = "PreferredContactMethod"
	FormParamIndividualPreferredContactMethodComments = "PreferredContactMethodComments"
	FormParamIndividualPreferredName                  = "PreferredName"
	FormParamIndividualPrefersToRemainAnonymous       = "PrefersToRemainAnonymous"
	FormParamIndividualPresentsProtectionConcerns     = "PresentsProtectionConcerns"
	FormParamIndividualSelfCareDisabilityLevel        = "SelfCareDisabilityLevel"
	FormParamIndividualSpokenLanguage1                = "SpokenLanguage1"
	FormParamIndividualSpokenLanguage2                = "SpokenLanguage2"
	FormParamIndividualSpokenLanguage3                = "SpokenLanguage3"
	FormParamIndividualVisionDisabilityLevel          = "VisionDisabilityLevel"

	FormParamsGetIndividualCognitiveDisabilityLevel        = "cognitive_disability_level"
	FormParamsGetIndividualCollectionAdministrativeArea1   = "collection_administrative_area_1"
	FormParamsGetIndividualCollectionAdministrativeArea2   = "collection_administrative_area_2"
	FormParamsGetIndividualCollectionAdministrativeArea3   = "collection_administrative_area_3"
	FormParamsGetIndividualCollectionAgentName             = "collection_agent_name"
	FormParamsGetIndividualCollectionAgentTitle            = "collection_agent_title"
	FormParamsGetIndividualsAddress                        = "address"
	FormParamsGetIndividualsAgeFrom                        = "age_from"
	FormParamsGetIndividualsAgeTo                          = "age_to"
	FormParamsGetIndividualsBirthDateFrom                  = "birth_date_from"
	FormParamsGetIndividualsBirthDateTo                    = "birth_date_to"
	FormParamsGetIndividualsCognitiveDisabilityLevel       = "cognitive_disability_level"
	FormParamsGetIndividualsCollectionAdministrativeArea1  = "collection_administrative_area_1"
	FormParamsGetIndividualsCollectionAdministrativeArea2  = "collection_administrative_area_2"
	FormParamsGetIndividualsCollectionAdministrativeArea3  = "collection_administrative_area_3"
	FormParamsGetIndividualsCollectionAgentName            = "collection_agent_name"
	FormParamsGetIndividualsCollectionAgentTitle           = "collection_agent_title"
	FormParamsGetIndividualsCollectionTimeFrom             = "collection_time_from"
	FormParamsGetIndividualsCollectionTimeTo               = "collection_time_to"
	FormParamsGetIndividualsCommunicationDisabilityLevel   = "communication_disability_level "
	FormParamsGetIndividualsCommunityID                    = "community_id"
	FormParamsGetIndividualsCountryID                      = "country_id"
	FormParamsGetIndividualsCreatedAtFrom                  = "created_at_from"
	FormParamsGetIndividualsCreatedAtTo                    = "created_at_to"
	FormParamsGetIndividualsDisplacementStatus             = "displacement_status"
	FormParamsGetIndividualsEmail                          = "email"
	FormParamsGetIndividualsFullName                       = "full_name"
	FormParamsGetIndividualsFreeField1                     = "free_field_1"
	FormParamsGetIndividualsFreeField2                     = "free_field_2"
	FormParamsGetIndividualsFreeField3                     = "free_field_3"
	FormParamsGetIndividualsFreeField4                     = "free_field_4"
	FormParamsGetIndividualsFreeField5                     = "free_field_5"
	FormParamsGetIndividualsGender                         = "gender"
	FormParamsGetIndividualsHasCognitiveDisability         = "has_cognitive_disability"
	FormParamsGetIndividualsHasCommunicationDisability     = "has_communication_disability"
	FormParamsGetIndividualsHasConsentedToReferral         = "has_consented_to_referral"
	FormParamsGetIndividualsHasConsentedToRgpd             = "has_consented_to_rgpd"
	FormParamsGetIndividualsHasHearingDisability           = "has_hearing_disability"
	FormParamsGetIndividualsHasMobilityDisability          = "has_mobility_disability"
	FormParamsGetIndividualsHasSelfCareDisability          = "has_selfcare_disability"
	FormParamsGetIndividualsHasVisionDisability            = "has_vision_disability"
	FormParamsGetIndividualsHearingDisabilityLevel         = "hearing_disability_level"
	FormParamsGetIndividualsHouseholdID                    = "household_id"
	FormParamsGetIndividualsID                             = "id"
	FormParamsGetIndividualsIdentificationContext          = "identification_context"
	FormParamsGetIndividualsIdentificationNumber           = "identification_number"
	FormParamsGetIndividualsInternalID                     = "internal_id"
	FormParamsGetIndividualsIsHeadOfCommunity              = "is_head_of_community"
	FormParamsGetIndividualsIsHeadOfHousehold              = "is_head_of_household"
	FormParamsGetIndividualsIsMinor                        = "is_minor"
	FormParamsGetIndividualsMobilityDisabilityLevel        = "mobility_disability_level"
	FormParamsGetIndividualsNationality                    = "nationality"
	FormParamsGetIndividualsPhoneNumber                    = "phone_number"
	FormParamsGetIndividualsPreferredCommunicationLanguage = "preferred_communication_language"
	FormParamsGetIndividualsPreferredContactMethod         = "preferred_contact_method"
	FormParamsGetIndividualsPrefersToRemainAnonymous       = "prefers_to_remain_anonymous"
	FormParamsGetIndividualsPresentsProtectionConcerns     = "presents_protection_concerns"
	FormParamsGetIndividualsSelfCareDisabilityLevel        = "selfcare_disability_level"
	FormParamsGetIndividualsSkip                           = "skip"
	FormParamsGetIndividualsSpokenLanguage                 = "spoken_language"
	FormParamsGetIndividualsTake                           = "take"
	FormParamsGetIndividualsUpdatedAtFrom                  = "updated_at_from"
	FormParamsGetIndividualsUpdatedAtTo                    = "updated_at_to"
	FormParamsGetIndividualsVisionDisabilityLevel          = "vision_disability_level"
	FormParamsGetIndividualsSort                           = "sort"

	DBColumnIndividualAddress                        = "address"
	DBColumnIndividualAge                            = "age"
	DBColumnIndividualBirthDate                      = "birth_date"
	DBColumnIndividualCognitiveDisabilityLevel       = "cognitive_disability_level"
	DBColumnIndividualCollectionAdministrativeArea1  = "collection_administrative_area_1"
	DBColumnIndividualCollectionAdministrativeArea2  = "collection_administrative_area_2"
	DBColumnIndividualCollectionAdministrativeArea3  = "collection_administrative_area_3"
	DBColumnIndividualCollectionAgentName            = "collection_agent_name"
	DBColumnIndividualCollectionAgentTitle           = "collection_agent_title"
	DBColumnIndividualCollectionTime                 = "collection_time"
	DBColumnIndividualComments                       = "comments"
	DBColumnIndividualCommunicationDisabilityLevel   = "communication_disability_level"
	DBColumnIndividualCommunityID                    = "community_id"
	DBColumnIndividualCountryID                      = "country_id"
	DBColumnIndividualCreatedAt                      = "created_at"
	DBColumnIndividualDeletedAt                      = "deleted_at"
	DBColumnIndividualDisplacementStatus             = "displacement_status"
	DBColumnIndividualEmail1                         = "email_1"
	DBColumnIndividualEmail2                         = "email_2"
	DBColumnIndividualEmail3                         = "email_3"
	DBColumnIndividualFullName                       = "full_name"
	DBColumnIndividualFreeField1                     = "free_field_1"
	DBColumnIndividualFreeField2                     = "free_field_2"
	DBColumnIndividualFreeField3                     = "free_field_3"
	DBColumnIndividualFreeField4                     = "free_field_4"
	DBColumnIndividualFreeField5                     = "free_field_5"
	DBColumnIndividualGender                         = "gender"
	DBColumnIndividualHasCognitiveDisability         = "has_cognitive_disability"
	DBColumnIndividualHasCommunicationDisability     = "has_communication_disability"
	DBColumnIndividualHasConsentedToRGPD             = "has_consented_to_rgpd"
	DBColumnIndividualHasConsentedToReferral         = "has_consented_to_referral"
	DBColumnIndividualHasHearingDisability           = "has_hearing_disability"
	DBColumnIndividualHasMobilityDisability          = "has_mobility_disability"
	DBColumnIndividualHasSelfCareDisability          = "has_selfcare_disability"
	DBColumnIndividualHasVisionDisability            = "has_vision_disability"
	DBColumnIndividualHearingDisabilityLevel         = "hearing_disability_level"
	DBColumnIndividualHouseholdID                    = "household_id"
	DBColumnIndividualID                             = "id"
	DBColumnIndividualIdentificationContext          = "identification_context"
	DBColumnIndividualIdentificationNumber1          = "identification_number_1"
	DBColumnIndividualIdentificationNumber2          = "identification_number_2"
	DBColumnIndividualIdentificationNumber3          = "identification_number_3"
	DBColumnIndividualIdentificationType1            = "identification_type_1"
	DBColumnIndividualIdentificationType2            = "identification_type_2"
	DBColumnIndividualIdentificationType3            = "identification_type_3"
	DBColumnIndividualIdentificationTypeExplanation1 = "identification_type_explanation_1"
	DBColumnIndividualIdentificationTypeExplanation2 = "identification_type_explanation_2"
	DBColumnIndividualIdentificationTypeExplanation3 = "identification_type_explanation_3"
	DBColumnIndividualInternalID                     = "internal_id"
	DBColumnIndividualIsHeadOfCommunity              = "is_head_of_community"
	DBColumnIndividualIsHeadOfHousehold              = "is_head_of_household"
	DBColumnIndividualIsMinor                        = "is_minor"
	DBColumnIndividualMobilityDisabilityLevel        = "mobility_disability_level"
	DBColumnIndividualNationality1                   = "nationality_1"
	DBColumnIndividualNationality2                   = "nationality_2"
	DBColumnIndividualNormalizedPhoneNumber1         = "normalized_phone_number_1"
	DBColumnIndividualNormalizedPhoneNumber2         = "normalized_phone_number_2"
	DBColumnIndividualNormalizedPhoneNumber3         = "normalized_phone_number_3"
	DBColumnIndividualPhoneNumber1                   = "phone_number_1"
	DBColumnIndividualPhoneNumber2                   = "phone_number_2"
	DBColumnIndividualPhoneNumber3                   = "phone_number_3"
	DBColumnIndividualPreferredCommunicationLanguage = "preferred_communication_language"
	DBColumnIndividualPreferredContactMethod         = "preferred_contact_method"
	DBColumnIndividualPreferredContactMethodComments = "preferred_contact_method_comments"
	DBColumnIndividualPreferredName                  = "preferred_name"
	DBColumnIndividualPrefersToRemainAnonymous       = "prefers_to_remain_anonymous"
	DBColumnIndividualPresentsProtectionConcerns     = "presents_protection_concerns"
	DBColumnIndividualSelfCareDisabilityLevel        = "selfcare_disability_level"
	DBColumnIndividualSpokenLanguage1                = "spoken_language_1"
	DBColumnIndividualSpokenLanguage2                = "spoken_language_2"
	DBColumnIndividualSpokenLanguage3                = "spoken_language_3"
	DBColumnIndividualUpdatedAt                      = "updated_at"
	DBColumnIndividualVisionDisabilityLevel          = "vision_disability_level"

	FileColumnIndividualAddress                        = "address"
	FileColumnIndividualAge                            = "age"
	FileColumnIndividualBirthDate                      = "birth_date"
	FileColumnIndividualCognitiveDisabilityLevel       = "cognitive_disability_level"
	FileColumnIndividualCollectionAdministrativeArea1  = "collection_administrative_area_1"
	FileColumnIndividualCollectionAdministrativeArea2  = "collection_administrative_area_2"
	FileColumnIndividualCollectionAdministrativeArea3  = "collection_administrative_area_3"
	FileColumnIndividualCollectionAgentName            = "collection_agent_name"
	FileColumnIndividualCollectionAgentTitle           = "collection_agent_title"
	FileColumnIndividualCollectionTime                 = "collection_time"
	FileColumnIndividualComments                       = "comments"
	FileColumnIndividualCommunicationDisabilityLevel   = "communication_disability_level"
	FileColumnIndividualCommunityID                    = "community_id"
	FileColumnIndividualCountryID                      = "country_id"
	FileColumnIndividualCreatedAt                      = "created_at"
	FileColumnIndividualDisplacementStatus             = "displacement_status"
	FileColumnIndividualEmail1                         = "email_1"
	FileColumnIndividualEmail2                         = "email_2"
	FileColumnIndividualEmail3                         = "email_3"
	FileColumnIndividualFullName                       = "full_name"
	FileColumnIndividualFreeField1                     = "free_field_1"
	FileColumnIndividualFreeField2                     = "free_field_2"
	FileColumnIndividualFreeField3                     = "free_field_3"
	FileColumnIndividualFreeField4                     = "free_field_4"
	FileColumnIndividualFreeField5                     = "free_field_5"
	FileColumnIndividualGender                         = "gender"
	FileColumnIndividualHasCognitiveDisability         = "has_cognitive_disability"
	FileColumnIndividualHasCommunicationDisability     = "has_communication_disability"
	FileColumnIndividualHasConsentedToRGPD             = "has_consented_to_rgpd"
	FileColumnIndividualHasConsentedToReferral         = "has_consented_to_referral"
	FileColumnIndividualHasHearingDisability           = "has_hearing_disability"
	FileColumnIndividualHasMobilityDisability          = "has_mobility_disability"
	FileColumnIndividualHasSelfCareDisability          = "has_selfcare_disability"
	FileColumnIndividualHasVisionDisability            = "has_vision_disability"
	FileColumnIndividualHearingDisabilityLevel         = "hearing_disability_level"
	FileColumnIndividualHouseholdID                    = "household_id"
	FileColumnIndividualID                             = "id"
	FileColumnIndividualIdentificationContext          = "identification_context"
	FileColumnIndividualIdentificationNumber1          = "identification_number_1"
	FileColumnIndividualIdentificationNumber2          = "identification_number_2"
	FileColumnIndividualIdentificationNumber3          = "identification_number_3"
	FileColumnIndividualIdentificationType1            = "identification_type_1"
	FileColumnIndividualIdentificationType2            = "identification_type_2"
	FileColumnIndividualIdentificationType3            = "identification_type_3"
	FileColumnIndividualIdentificationTypeExplanation1 = "identification_type_explanation_1"
	FileColumnIndividualIdentificationTypeExplanation2 = "identification_type_explanation_2"
	FileColumnIndividualIdentificationTypeExplanation3 = "identification_type_explanation_3"
	FileColumnIndividualInternalID                     = "internal_id"
	FileColumnIndividualIsHeadOfCommunity              = "is_head_of_community"
	FileColumnIndividualIsHeadOfHousehold              = "is_head_of_household"
	FileColumnIndividualIsMinor                        = "is_minor"
	FileColumnIndividualMobilityDisabilityLevel        = "mobility_disability_level"
	FileColumnIndividualNationality1                   = "nationality_1"
	FileColumnIndividualNationality2                   = "nationality_2"
	FileColumnIndividualPhoneNumber1                   = "phone_number_1"
	FileColumnIndividualPhoneNumber2                   = "phone_number_2"
	FileColumnIndividualPhoneNumber3                   = "phone_number_3"
	FileColumnIndividualPreferredCommunicationLanguage = "preferred_communication_language"
	FileColumnIndividualPreferredContactMethod         = "preferred_contact_method"
	FileColumnIndividualPreferredContactMethodComments = "preferred_contact_method_comments"
	FileColumnIndividualPreferredName                  = "preferred_name"
	FileColumnIndividualPrefersToRemainAnonymous       = "prefers_to_remain_anonymous"
	FileColumnIndividualPresentsProtectionConcerns     = "presents_protection_concerns"
	FileColumnIndividualSelfCareDisabilityLevel        = "selfcare_disability_level"
	FileColumnIndividualSpokenLanguage1                = "spoken_language_1"
	FileColumnIndividualSpokenLanguage2                = "spoken_language_2"
	FileColumnIndividualSpokenLanguage3                = "spoken_language_3"
	FileColumnIndividualUpdatedAt                      = "updated_at"
	FileColumnIndividualVisionDisabilityLevel          = "vision_disability_level"
)

var IndividualDBColumns = containers.NewStringSet(
	DBColumnIndividualAddress,
	DBColumnIndividualAge,
	DBColumnIndividualBirthDate,
	DBColumnIndividualCognitiveDisabilityLevel,
	DBColumnIndividualCollectionAdministrativeArea1,
	DBColumnIndividualCollectionAdministrativeArea2,
	DBColumnIndividualCollectionAdministrativeArea3,
	DBColumnIndividualCollectionAgentName,
	DBColumnIndividualCollectionAgentTitle,
	DBColumnIndividualCollectionTime,
	DBColumnIndividualComments,
	DBColumnIndividualCommunicationDisabilityLevel,
	DBColumnIndividualCommunityID,
	DBColumnIndividualCountryID,
	DBColumnIndividualCreatedAt,
	DBColumnIndividualDeletedAt,
	DBColumnIndividualDisplacementStatus,
	DBColumnIndividualEmail1,
	DBColumnIndividualEmail2,
	DBColumnIndividualEmail3,
	DBColumnIndividualFullName,
	DBColumnIndividualFreeField1,
	DBColumnIndividualFreeField2,
	DBColumnIndividualFreeField3,
	DBColumnIndividualFreeField4,
	DBColumnIndividualFreeField5,
	DBColumnIndividualGender,
	DBColumnIndividualHasCognitiveDisability,
	DBColumnIndividualHasCommunicationDisability,
	DBColumnIndividualHasConsentedToRGPD,
	DBColumnIndividualHasConsentedToReferral,
	DBColumnIndividualHasHearingDisability,
	DBColumnIndividualHasMobilityDisability,
	DBColumnIndividualHasSelfCareDisability,
	DBColumnIndividualHasVisionDisability,
	DBColumnIndividualHearingDisabilityLevel,
	DBColumnIndividualHouseholdID,
	DBColumnIndividualID,
	DBColumnIndividualIdentificationContext,
	DBColumnIndividualIdentificationNumber1,
	DBColumnIndividualIdentificationNumber2,
	DBColumnIndividualIdentificationNumber3,
	DBColumnIndividualIdentificationType1,
	DBColumnIndividualIdentificationType2,
	DBColumnIndividualIdentificationType3,
	DBColumnIndividualIdentificationTypeExplanation1,
	DBColumnIndividualIdentificationTypeExplanation2,
	DBColumnIndividualIdentificationTypeExplanation3,
	DBColumnIndividualInternalID,
	DBColumnIndividualIsHeadOfCommunity,
	DBColumnIndividualIsHeadOfHousehold,
	DBColumnIndividualIsMinor,
	DBColumnIndividualMobilityDisabilityLevel,
	DBColumnIndividualNationality1,
	DBColumnIndividualNationality2,
	DBColumnIndividualNormalizedPhoneNumber1,
	DBColumnIndividualNormalizedPhoneNumber2,
	DBColumnIndividualNormalizedPhoneNumber3,
	DBColumnIndividualPhoneNumber1,
	DBColumnIndividualPhoneNumber2,
	DBColumnIndividualPhoneNumber3,
	DBColumnIndividualPreferredCommunicationLanguage,
	DBColumnIndividualPreferredContactMethod,
	DBColumnIndividualPreferredContactMethodComments,
	DBColumnIndividualPreferredName,
	DBColumnIndividualPrefersToRemainAnonymous,
	DBColumnIndividualPresentsProtectionConcerns,
	DBColumnIndividualSelfCareDisabilityLevel,
	DBColumnIndividualSpokenLanguage1,
	DBColumnIndividualSpokenLanguage2,
	DBColumnIndividualSpokenLanguage3,
	DBColumnIndividualUpdatedAt,
	DBColumnIndividualVisionDisabilityLevel,
)

var IndividualFileColumns = []string{
	FileColumnIndividualAddress,
	FileColumnIndividualAge,
	FileColumnIndividualBirthDate,
	FileColumnIndividualCognitiveDisabilityLevel,
	FileColumnIndividualCollectionAdministrativeArea1,
	FileColumnIndividualCollectionAdministrativeArea2,
	FileColumnIndividualCollectionAdministrativeArea3,
	FileColumnIndividualCollectionAgentName,
	FileColumnIndividualCollectionAgentTitle,
	FileColumnIndividualCollectionTime,
	FileColumnIndividualComments,
	FileColumnIndividualCommunicationDisabilityLevel,
	FileColumnIndividualCommunityID,
	FileColumnIndividualCountryID,
	FileColumnIndividualCreatedAt,
	FileColumnIndividualDisplacementStatus,
	FileColumnIndividualEmail1,
	FileColumnIndividualEmail2,
	FileColumnIndividualEmail3,
	FileColumnIndividualFullName,
	FileColumnIndividualFreeField1,
	FileColumnIndividualFreeField2,
	FileColumnIndividualFreeField3,
	FileColumnIndividualFreeField4,
	FileColumnIndividualFreeField5,
	FileColumnIndividualGender,
	FileColumnIndividualHasCognitiveDisability,
	FileColumnIndividualHasCommunicationDisability,
	FileColumnIndividualHasConsentedToRGPD,
	FileColumnIndividualHasConsentedToReferral,
	FileColumnIndividualHasHearingDisability,
	FileColumnIndividualHasMobilityDisability,
	FileColumnIndividualHasSelfCareDisability,
	FileColumnIndividualHasVisionDisability,
	FileColumnIndividualHearingDisabilityLevel,
	FileColumnIndividualHouseholdID,
	FileColumnIndividualID,
	FileColumnIndividualIdentificationContext,
	FileColumnIndividualIdentificationNumber1,
	FileColumnIndividualIdentificationNumber2,
	FileColumnIndividualIdentificationNumber3,
	FileColumnIndividualIdentificationType1,
	FileColumnIndividualIdentificationType2,
	FileColumnIndividualIdentificationType3,
	FileColumnIndividualIdentificationTypeExplanation1,
	FileColumnIndividualIdentificationTypeExplanation2,
	FileColumnIndividualIdentificationTypeExplanation3,
	FileColumnIndividualInternalID,
	FileColumnIndividualIsHeadOfCommunity,
	FileColumnIndividualIsHeadOfHousehold,
	FileColumnIndividualIsMinor,
	FileColumnIndividualMobilityDisabilityLevel,
	FileColumnIndividualNationality1,
	FileColumnIndividualNationality2,
	FileColumnIndividualPhoneNumber1,
	FileColumnIndividualPhoneNumber2,
	FileColumnIndividualPhoneNumber3,
	FileColumnIndividualPreferredCommunicationLanguage,
	FileColumnIndividualPreferredContactMethod,
	FileColumnIndividualPreferredContactMethodComments,
	FileColumnIndividualPreferredName,
	FileColumnIndividualPrefersToRemainAnonymous,
	FileColumnIndividualPresentsProtectionConcerns,
	FileColumnIndividualSelfCareDisabilityLevel,
	FileColumnIndividualSpokenLanguage1,
	FileColumnIndividualSpokenLanguage2,
	FileColumnIndividualSpokenLanguage3,
	FileColumnIndividualUpdatedAt,
	FileColumnIndividualVisionDisabilityLevel,
}

var IndividualDBToFileMap = map[string]string{
	DBColumnIndividualAddress:                        FileColumnIndividualAddress,
	DBColumnIndividualAge:                            FileColumnIndividualAge,
	DBColumnIndividualBirthDate:                      FileColumnIndividualBirthDate,
	DBColumnIndividualCognitiveDisabilityLevel:       FileColumnIndividualCognitiveDisabilityLevel,
	DBColumnIndividualCollectionAdministrativeArea1:  FileColumnIndividualCollectionAdministrativeArea1,
	DBColumnIndividualCollectionAdministrativeArea2:  FileColumnIndividualCollectionAdministrativeArea2,
	DBColumnIndividualCollectionAdministrativeArea3:  FileColumnIndividualCollectionAdministrativeArea3,
	DBColumnIndividualCollectionAgentName:            FileColumnIndividualCollectionAgentName,
	DBColumnIndividualCollectionAgentTitle:           FileColumnIndividualCollectionAgentTitle,
	DBColumnIndividualCollectionTime:                 FileColumnIndividualCollectionTime,
	DBColumnIndividualComments:                       FileColumnIndividualComments,
	DBColumnIndividualCommunicationDisabilityLevel:   FileColumnIndividualCommunicationDisabilityLevel,
	DBColumnIndividualCommunityID:                    FileColumnIndividualCommunityID,
	DBColumnIndividualCountryID:                      FileColumnIndividualCountryID,
	DBColumnIndividualCreatedAt:                      FileColumnIndividualCreatedAt,
	DBColumnIndividualDisplacementStatus:             FileColumnIndividualDisplacementStatus,
	DBColumnIndividualEmail1:                         FileColumnIndividualEmail1,
	DBColumnIndividualEmail2:                         FileColumnIndividualEmail2,
	DBColumnIndividualEmail3:                         FileColumnIndividualEmail3,
	DBColumnIndividualFullName:                       FileColumnIndividualFullName,
	DBColumnIndividualFreeField1:                     FileColumnIndividualFreeField1,
	DBColumnIndividualFreeField2:                     FileColumnIndividualFreeField2,
	DBColumnIndividualFreeField3:                     FileColumnIndividualFreeField3,
	DBColumnIndividualFreeField4:                     FileColumnIndividualFreeField4,
	DBColumnIndividualFreeField5:                     FileColumnIndividualFreeField5,
	DBColumnIndividualGender:                         FileColumnIndividualGender,
	DBColumnIndividualHasCognitiveDisability:         FileColumnIndividualHasCognitiveDisability,
	DBColumnIndividualHasCommunicationDisability:     FileColumnIndividualHasCommunicationDisability,
	DBColumnIndividualHasConsentedToRGPD:             FileColumnIndividualHasConsentedToRGPD,
	DBColumnIndividualHasConsentedToReferral:         FileColumnIndividualHasConsentedToReferral,
	DBColumnIndividualHasHearingDisability:           FileColumnIndividualHasHearingDisability,
	DBColumnIndividualHasMobilityDisability:          FileColumnIndividualHasMobilityDisability,
	DBColumnIndividualHasSelfCareDisability:          FileColumnIndividualHasSelfCareDisability,
	DBColumnIndividualHasVisionDisability:            FileColumnIndividualHasVisionDisability,
	DBColumnIndividualHearingDisabilityLevel:         FileColumnIndividualHearingDisabilityLevel,
	DBColumnIndividualHouseholdID:                    FileColumnIndividualHouseholdID,
	DBColumnIndividualID:                             FileColumnIndividualID,
	DBColumnIndividualIdentificationContext:          FileColumnIndividualIdentificationContext,
	DBColumnIndividualIdentificationNumber1:          FileColumnIndividualIdentificationNumber1,
	DBColumnIndividualIdentificationNumber2:          FileColumnIndividualIdentificationNumber2,
	DBColumnIndividualIdentificationNumber3:          FileColumnIndividualIdentificationNumber3,
	DBColumnIndividualIdentificationType1:            FileColumnIndividualIdentificationType1,
	DBColumnIndividualIdentificationType2:            FileColumnIndividualIdentificationType2,
	DBColumnIndividualIdentificationType3:            FileColumnIndividualIdentificationType3,
	DBColumnIndividualIdentificationTypeExplanation1: FileColumnIndividualIdentificationTypeExplanation1,
	DBColumnIndividualIdentificationTypeExplanation2: FileColumnIndividualIdentificationTypeExplanation2,
	DBColumnIndividualIdentificationTypeExplanation3: FileColumnIndividualIdentificationTypeExplanation3,
	DBColumnIndividualInternalID:                     FileColumnIndividualInternalID,
	DBColumnIndividualIsHeadOfCommunity:              FileColumnIndividualIsHeadOfCommunity,
	DBColumnIndividualIsHeadOfHousehold:              FileColumnIndividualIsHeadOfHousehold,
	DBColumnIndividualIsMinor:                        FileColumnIndividualIsMinor,
	DBColumnIndividualMobilityDisabilityLevel:        FileColumnIndividualMobilityDisabilityLevel,
	DBColumnIndividualNationality1:                   FileColumnIndividualNationality1,
	DBColumnIndividualNationality2:                   FileColumnIndividualNationality2,
	DBColumnIndividualPhoneNumber1:                   FileColumnIndividualPhoneNumber1,
	DBColumnIndividualPhoneNumber2:                   FileColumnIndividualPhoneNumber2,
	DBColumnIndividualPhoneNumber3:                   FileColumnIndividualPhoneNumber3,
	DBColumnIndividualPreferredCommunicationLanguage: FileColumnIndividualPreferredCommunicationLanguage,
	DBColumnIndividualPreferredContactMethod:         FileColumnIndividualPreferredContactMethod,
	DBColumnIndividualPreferredContactMethodComments: FileColumnIndividualPreferredContactMethodComments,
	DBColumnIndividualPreferredName:                  FileColumnIndividualPreferredName,
	DBColumnIndividualPrefersToRemainAnonymous:       FileColumnIndividualPrefersToRemainAnonymous,
	DBColumnIndividualPresentsProtectionConcerns:     FileColumnIndividualPresentsProtectionConcerns,
	DBColumnIndividualSelfCareDisabilityLevel:        FileColumnIndividualSelfCareDisabilityLevel,
	DBColumnIndividualSpokenLanguage1:                FileColumnIndividualSpokenLanguage1,
	DBColumnIndividualSpokenLanguage2:                FileColumnIndividualSpokenLanguage2,
	DBColumnIndividualSpokenLanguage3:                FileColumnIndividualSpokenLanguage3,
	DBColumnIndividualVisionDisabilityLevel:          FileColumnIndividualVisionDisabilityLevel,
	DBColumnIndividualUpdatedAt:                      FileColumnIndividualUpdatedAt,
}

var IndividualFileToDBMap = map[string]string{
	FileColumnIndividualAddress:                        DBColumnIndividualAddress,
	FileColumnIndividualAge:                            DBColumnIndividualAge,
	FileColumnIndividualBirthDate:                      DBColumnIndividualBirthDate,
	FileColumnIndividualCognitiveDisabilityLevel:       DBColumnIndividualCognitiveDisabilityLevel,
	FileColumnIndividualCollectionAdministrativeArea1:  DBColumnIndividualCollectionAdministrativeArea1,
	FileColumnIndividualCollectionAdministrativeArea2:  DBColumnIndividualCollectionAdministrativeArea2,
	FileColumnIndividualCollectionAdministrativeArea3:  DBColumnIndividualCollectionAdministrativeArea3,
	FileColumnIndividualCollectionAgentName:            DBColumnIndividualCollectionAgentName,
	FileColumnIndividualCollectionAgentTitle:           DBColumnIndividualCollectionAgentTitle,
	FileColumnIndividualCollectionTime:                 DBColumnIndividualCollectionTime,
	FileColumnIndividualComments:                       DBColumnIndividualComments,
	FileColumnIndividualCommunicationDisabilityLevel:   DBColumnIndividualCommunicationDisabilityLevel,
	FileColumnIndividualCommunityID:                    DBColumnIndividualCommunityID,
	FileColumnIndividualCountryID:                      DBColumnIndividualCountryID,
	FileColumnIndividualDisplacementStatus:             DBColumnIndividualDisplacementStatus,
	FileColumnIndividualEmail1:                         DBColumnIndividualEmail1,
	FileColumnIndividualEmail2:                         DBColumnIndividualEmail2,
	FileColumnIndividualEmail3:                         DBColumnIndividualEmail3,
	FileColumnIndividualFullName:                       DBColumnIndividualFullName,
	FileColumnIndividualFreeField1:                     DBColumnIndividualFreeField1,
	FileColumnIndividualFreeField2:                     DBColumnIndividualFreeField2,
	FileColumnIndividualFreeField3:                     DBColumnIndividualFreeField3,
	FileColumnIndividualFreeField4:                     DBColumnIndividualFreeField4,
	FileColumnIndividualFreeField5:                     DBColumnIndividualFreeField5,
	FileColumnIndividualGender:                         DBColumnIndividualGender,
	FileColumnIndividualHasCognitiveDisability:         DBColumnIndividualHasCognitiveDisability,
	FileColumnIndividualHasCommunicationDisability:     DBColumnIndividualHasCommunicationDisability,
	FileColumnIndividualHasConsentedToRGPD:             DBColumnIndividualHasConsentedToRGPD,
	FileColumnIndividualHasConsentedToReferral:         DBColumnIndividualHasConsentedToReferral,
	FileColumnIndividualHasHearingDisability:           DBColumnIndividualHasHearingDisability,
	FileColumnIndividualHasMobilityDisability:          DBColumnIndividualHasMobilityDisability,
	FileColumnIndividualHasSelfCareDisability:          DBColumnIndividualHasSelfCareDisability,
	FileColumnIndividualHasVisionDisability:            DBColumnIndividualHasVisionDisability,
	FileColumnIndividualHearingDisabilityLevel:         DBColumnIndividualHearingDisabilityLevel,
	FileColumnIndividualHouseholdID:                    DBColumnIndividualHouseholdID,
	FileColumnIndividualID:                             DBColumnIndividualID,
	FileColumnIndividualIdentificationContext:          DBColumnIndividualIdentificationContext,
	FileColumnIndividualIdentificationNumber1:          DBColumnIndividualIdentificationNumber1,
	FileColumnIndividualIdentificationNumber2:          DBColumnIndividualIdentificationNumber2,
	FileColumnIndividualIdentificationNumber3:          DBColumnIndividualIdentificationNumber3,
	FileColumnIndividualIdentificationType1:            DBColumnIndividualIdentificationType1,
	FileColumnIndividualIdentificationType2:            DBColumnIndividualIdentificationType2,
	FileColumnIndividualIdentificationType3:            DBColumnIndividualIdentificationType3,
	FileColumnIndividualIdentificationTypeExplanation1: DBColumnIndividualIdentificationTypeExplanation1,
	FileColumnIndividualIdentificationTypeExplanation2: DBColumnIndividualIdentificationTypeExplanation2,
	FileColumnIndividualIdentificationTypeExplanation3: DBColumnIndividualIdentificationTypeExplanation3,
	FileColumnIndividualInternalID:                     DBColumnIndividualInternalID,
	FileColumnIndividualIsHeadOfCommunity:              DBColumnIndividualIsHeadOfCommunity,
	FileColumnIndividualIsHeadOfHousehold:              DBColumnIndividualIsHeadOfHousehold,
	FileColumnIndividualIsMinor:                        DBColumnIndividualIsMinor,
	FileColumnIndividualMobilityDisabilityLevel:        DBColumnIndividualMobilityDisabilityLevel,
	FileColumnIndividualNationality1:                   DBColumnIndividualNationality1,
	FileColumnIndividualNationality2:                   DBColumnIndividualNationality2,
	FileColumnIndividualPhoneNumber1:                   DBColumnIndividualPhoneNumber1,
	FileColumnIndividualPhoneNumber2:                   DBColumnIndividualPhoneNumber2,
	FileColumnIndividualPhoneNumber3:                   DBColumnIndividualPhoneNumber3,
	FileColumnIndividualPreferredCommunicationLanguage: DBColumnIndividualPreferredCommunicationLanguage,
	FileColumnIndividualPreferredContactMethod:         DBColumnIndividualPreferredContactMethod,
	FileColumnIndividualPreferredContactMethodComments: DBColumnIndividualPreferredContactMethodComments,
	FileColumnIndividualPreferredName:                  DBColumnIndividualPreferredName,
	FileColumnIndividualPrefersToRemainAnonymous:       DBColumnIndividualPrefersToRemainAnonymous,
	FileColumnIndividualPresentsProtectionConcerns:     DBColumnIndividualPresentsProtectionConcerns,
	FileColumnIndividualSelfCareDisabilityLevel:        DBColumnIndividualSelfCareDisabilityLevel,
	FileColumnIndividualSpokenLanguage1:                DBColumnIndividualSpokenLanguage1,
	FileColumnIndividualSpokenLanguage2:                DBColumnIndividualSpokenLanguage2,
	FileColumnIndividualSpokenLanguage3:                DBColumnIndividualSpokenLanguage3,
	FileColumnIndividualVisionDisabilityLevel:          DBColumnIndividualVisionDisabilityLevel,
}
