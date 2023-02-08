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
	FormParamIndividualCollectionOffice               = "CollectionOffice "
	FormParamIndividualCollectionAgentID              = "CollectionAgentName "
	FormParamIndividualCollectionTime                 = "CollectionTime"
	FormParamIndividualCommunicationDisabilityLevel   = "CommunicationDisabilityLevel"
	FormParamIndividualCommunityID                    = "CommunityID"
	FormParamIndividualCountryID                      = "Country"
	FormParamIndividualDisplacementStatus             = "DisplacementStatus"
	FormParamIndividualDisplacementStatusComment      = "DisplacementStatusComment"
	FormParamIndividualEmail1                         = "Email1"
	FormParamIndividualEmail2                         = "Email2"
	FormParamIndividualEmail3                         = "Email3"
	FormParamIndividualFullName                       = "FullName"
	FormParamIndividualFirstName                      = "FirstName"
	FormParamIndividualMiddleName                     = "MiddleName"
	FormParamIndividualLastName                       = "LastName"
	FormParamIndividualSex                            = "Sex"
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
	FormParamIndividualEngagementContext              = "EngagementContext"
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
	FormParamsIndividualsIsFemaleHeadedHousehold      = "IsFemaleHeadedHousehold"
	FormParamIndividualIsMinorHeadedHousehold         = "IsMinorHeadedHousehold"
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
	FormParamServiceCC1                               = "ServiceCC1"
	FormParamServiceRequestedDate1                    = "ServiceRequestedDate1"
	FormParamServiceDeliveredDate1                    = "ServiceDeliveredDate1"
	FormParamServiceComments1                         = "ServiceComments1"
	FormParamServiceCC2                               = "ServiceCC2"
	FormParamServiceRequestedDate2                    = "ServiceRequestedDate2"
	FormParamServiceDeliveredDate2                    = "ServiceDeliveredDate2"
	FormParamServiceComments2                         = "ServiceComments2"
	FormParamServiceCC3                               = "ServiceCC3"
	FormParamServiceRequestedDate3                    = "ServiceRequestedDate3"
	FormParamServiceDeliveredDate3                    = "ServiceDeliveredDate3"
	FormParamServiceComments3                         = "ServiceComments3"
	FormParamServiceCC4                               = "ServiceCC4"
	FormParamServiceRequestedDate4                    = "ServiceRequestedDate4"
	FormParamServiceDeliveredDate4                    = "ServiceDeliveredDate4"
	FormParamServiceComments4                         = "ServiceComments4"
	FormParamServiceCC5                               = "ServiceCC5"
	FormParamServiceRequestedDate5                    = "ServiceRequestedDate5"
	FormParamServiceDeliveredDate5                    = "ServiceDeliveredDate5"
	FormParamServiceComments5                         = "ServiceComments5"
	FormParamServiceCC6                               = "ServiceCC6"
	FormParamServiceRequestedDate6                    = "ServiceRequestedDate6"
	FormParamServiceDeliveredDate6                    = "ServiceDeliveredDate6"
	FormParamServiceComments6                         = "ServiceComments6"
	FormParamServiceCC7                               = "ServiceCC7"
	FormParamServiceRequestedDate7                    = "ServiceRequestedDate7"
	FormParamServiceDeliveredDate7                    = "ServiceDeliveredDate7"
	FormParamServiceComments7                         = "ServiceComments7"

	FormParamsGetIndividualCognitiveDisabilityLevel        = "cognitive_disability_level"
	FormParamsGetIndividualCollectionAdministrativeArea1   = "collection_administrative_area_1"
	FormParamsGetIndividualCollectionAdministrativeArea2   = "collection_administrative_area_2"
	FormParamsGetIndividualCollectionAdministrativeArea3   = "collection_administrative_area_3"
	FormParamsGetIndividualCollectionOffice                = "collection_office"
	FormParamsGetIndividualCollectionAgentName             = "collection_agent_name"
	FormParamsGetIndividualCollectionAgentTitle            = "collection_agent_title"
	FormParamsGetIndividualsActive                         = "inactive"
	FormParamsGetIndividualsAddress                        = "address"
	FormParamsGetIndividualsAgeFrom                        = "age_from"
	FormParamsGetIndividualsAgeTo                          = "age_to"
	FormParamsGetIndividualsBirthDateFrom                  = "birth_date_from"
	FormParamsGetIndividualsBirthDateTo                    = "birth_date_to"
	FormParamsGetIndividualsCognitiveDisabilityLevel       = "cognitive_disability_level"
	FormParamsGetIndividualsCollectionAdministrativeArea1  = "collection_administrative_area_1"
	FormParamsGetIndividualsCollectionAdministrativeArea2  = "collection_administrative_area_2"
	FormParamsGetIndividualsCollectionAdministrativeArea3  = "collection_administrative_area_3"
	FormParamsGetIndividualsCollectionOffice               = "collection_office"
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
	FormParamsGetIndividualsDisplacementStatusComment      = "displacement_status_comment"
	FormParamsGetIndividualsEmail                          = "email"
	FormParamsGetIndividualsFullName                       = "full_name"
	FormParamsGetIndividualsFirstName                      = "first_name"
	FormParamsGetIndividualsMiddleName                     = "middle_name"
	FormParamsGetIndividualsLastName                       = "last_name"
	FormParamsGetIndividualsFreeField1                     = "free_field_1"
	FormParamsGetIndividualsFreeField2                     = "free_field_2"
	FormParamsGetIndividualsFreeField3                     = "free_field_3"
	FormParamsGetIndividualsFreeField4                     = "free_field_4"
	FormParamsGetIndividualsFreeField5                     = "free_field_5"
	FormParamsGetIndividualsSex                            = "sex"
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
	FormParamsGetIndividualsEngagementContext              = "engagement_context"
	FormParamsGetIndividualsIdentificationNumber           = "identification_number"
	FormParamsGetIndividualsInternalID                     = "internal_id"
	FormParamsGetIndividualsIsHeadOfCommunity              = "is_head_of_community"
	FormParamsGetIndividualsIsHeadOfHousehold              = "is_head_of_household"
	FormParamsGetIndividualsIsFemaleHeadedHousehold        = "is_female_headed_household"
	FormParamsGetIndividualsIsMinorHeadedHousehold         = "is_minor_headed_household"
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
	FormParamsGetIndividualsServiceCC1                     = "service_cc_1"
	FormParamsGetIndividualsServiceRequestedDate1From      = "service_requested_date_1_from"
	FormParamsGetIndividualsServiceRequestedDate1To        = "service_requested_date_1_to"
	FormParamsGetIndividualsServiceDeliveredDate1From      = "service_delivered_date_1_from"
	FormParamsGetIndividualsServiceDeliveredDate1To        = "service_delivered_date_1_to"
	FormParamsGetIndividualsServiceCC2                     = "service_cc_2"
	FormParamsGetIndividualsServiceRequestedDate2From      = "service_requested_date_2_from"
	FormParamsGetIndividualsServiceRequestedDate2To        = "service_requested_date_2_to"
	FormParamsGetIndividualsServiceDeliveredDate2From      = "service_delivered_date_2_from"
	FormParamsGetIndividualsServiceDeliveredDate2To        = "service_delivered_date_2_to"
	FormParamsGetIndividualsServiceCC3                     = "service_cc_3"
	FormParamsGetIndividualsServiceRequestedDate3From      = "service_requested_date_3_from"
	FormParamsGetIndividualsServiceRequestedDate3To        = "service_requested_date_3_to"
	FormParamsGetIndividualsServiceDeliveredDate3From      = "service_delivered_date_3_from"
	FormParamsGetIndividualsServiceDeliveredDate3To        = "service_delivered_date_3_to"
	FormParamsGetIndividualsServiceCC4                     = "service_cc_4"
	FormParamsGetIndividualsServiceRequestedDate4From      = "service_requested_date_4_from"
	FormParamsGetIndividualsServiceRequestedDate4To        = "service_requested_date_4_to"
	FormParamsGetIndividualsServiceDeliveredDate4From      = "service_delivered_date_4_from"
	FormParamsGetIndividualsServiceDeliveredDate4To        = "service_delivered_date_4_to"
	FormParamsGetIndividualsServiceCC5                     = "service_cc_5"
	FormParamsGetIndividualsServiceRequestedDate5From      = "service_requested_date_5_from"
	FormParamsGetIndividualsServiceRequestedDate5To        = "service_requested_date_5_to"
	FormParamsGetIndividualsServiceDeliveredDate5From      = "service_delivered_date_5_from"
	FormParamsGetIndividualsServiceDeliveredDate5To        = "service_delivered_date_5_to"
	FormParamsGetIndividualsServiceCC6                     = "service_cc_6"
	FormParamsGetIndividualsServiceRequestedDate6From      = "service_requested_date_6_from"
	FormParamsGetIndividualsServiceRequestedDate6To        = "service_requested_date_6_to"
	FormParamsGetIndividualsServiceDeliveredDate6From      = "service_delivered_date_6_from"
	FormParamsGetIndividualsServiceDeliveredDate6To        = "service_delivered_date_6_to"
	FormParamsGetIndividualsServiceCC7                     = "service_cc_7"
	FormParamsGetIndividualsServiceRequestedDate7From      = "service_requested_date_7_from"
	FormParamsGetIndividualsServiceRequestedDate7To        = "service_requested_date_7_to"
	FormParamsGetIndividualsServiceDeliveredDate7From      = "service_delivered_date_7_from"
	FormParamsGetIndividualsServiceDeliveredDate7To        = "service_delivered_date_7_to"
	FormParamsGetIndividualsSort                           = "sort"

	DBColumnIndividualInactive                       = "inactive"
	DBColumnIndividualAddress                        = "address"
	DBColumnIndividualAge                            = "age"
	DBColumnIndividualBirthDate                      = "birth_date"
	DBColumnIndividualCognitiveDisabilityLevel       = "cognitive_disability_level"
	DBColumnIndividualCollectionAdministrativeArea1  = "collection_administrative_area_1"
	DBColumnIndividualCollectionAdministrativeArea2  = "collection_administrative_area_2"
	DBColumnIndividualCollectionAdministrativeArea3  = "collection_administrative_area_3"
	DBColumnIndividualCollectionOffice               = "collection_office"
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
	DBColumnIndividualDisplacementStatusComment      = "displacement_status_comment"
	DBColumnIndividualEmail1                         = "email_1"
	DBColumnIndividualEmail2                         = "email_2"
	DBColumnIndividualEmail3                         = "email_3"
	DBColumnIndividualFullName                       = "full_name"
	DBColumnIndividualFirstName                      = "first_name"
	DBColumnIndividualMiddleName                     = "middle_name"
	DBColumnIndividualLastName                       = "last_name"
	DBColumnIndividualFreeField1                     = "free_field_1"
	DBColumnIndividualFreeField2                     = "free_field_2"
	DBColumnIndividualFreeField3                     = "free_field_3"
	DBColumnIndividualFreeField4                     = "free_field_4"
	DBColumnIndividualFreeField5                     = "free_field_5"
	DBColumnIndividualSex                            = "sex"
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
	DBColumnIndividualEngagementContext              = "engagement_context"
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
	DBColumnIndividualIsFemaleHeadedHousehold        = "is_female_headed_household"
	DBColumnIndividualIsMinorHeadedHousehold         = "is_minor_headed_household"
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
	DBColumnIndividualServiceCC1                     = "service_cc_1"
	DBColumnIndividualServiceRequestedDate1          = "service_requested_date_1"
	DBColumnIndividualServiceDeliveredDate1          = "service_delivered_date_1"
	DBColumnIndividualServiceComments1               = "service_comments_1"
	DBColumnIndividualServiceCC2                     = "service_cc_2"
	DBColumnIndividualServiceRequestedDate2          = "service_requested_date_2"
	DBColumnIndividualServiceDeliveredDate2          = "service_delivered_date_2"
	DBColumnIndividualServiceComments2               = "service_comments_2"
	DBColumnIndividualServiceCC3                     = "service_cc_3"
	DBColumnIndividualServiceRequestedDate3          = "service_requested_date_3"
	DBColumnIndividualServiceDeliveredDate3          = "service_delivered_date_3"
	DBColumnIndividualServiceComments3               = "service_comments_3"
	DBColumnIndividualServiceCC4                     = "service_cc_4"
	DBColumnIndividualServiceRequestedDate4          = "service_requested_date_4"
	DBColumnIndividualServiceDeliveredDate4          = "service_delivered_date_4"
	DBColumnIndividualServiceComments4               = "service_comments_4"
	DBColumnIndividualServiceCC5                     = "service_cc_5"
	DBColumnIndividualServiceRequestedDate5          = "service_requested_date_5"
	DBColumnIndividualServiceDeliveredDate5          = "service_delivered_date_5"
	DBColumnIndividualServiceComments5               = "service_comments_5"
	DBColumnIndividualServiceCC6                     = "service_cc_6"
	DBColumnIndividualServiceRequestedDate6          = "service_requested_date_6"
	DBColumnIndividualServiceDeliveredDate6          = "service_delivered_date_6"
	DBColumnIndividualServiceComments6               = "service_comments_6"
	DBColumnIndividualServiceCC7                     = "service_cc_7"
	DBColumnIndividualServiceRequestedDate7          = "service_requested_date_7"
	DBColumnIndividualServiceDeliveredDate7          = "service_delivered_date_7"
	DBColumnIndividualServiceComments7               = "service_comments_7"

	FileColumnIndividualInactive                       = "inactive"
	FileColumnIndividualAddress                        = "address"
	FileColumnIndividualAge                            = "age"
	FileColumnIndividualBirthDate                      = "birth_date"
	FileColumnIndividualCognitiveDisabilityLevel       = "cognitive_disability_level"
	FileColumnIndividualCollectionAdministrativeArea1  = "collection_administrative_area_1"
	FileColumnIndividualCollectionAdministrativeArea2  = "collection_administrative_area_2"
	FileColumnIndividualCollectionAdministrativeArea3  = "collection_administrative_area_3"
	FileColumnIndividualCollectionOffice               = "collection_office"
	FileColumnIndividualCollectionAgentName            = "collection_agent_name"
	FileColumnIndividualCollectionAgentTitle           = "collection_agent_title"
	FileColumnIndividualCollectionTime                 = "collection_time"
	FileColumnIndividualComments                       = "comments"
	FileColumnIndividualCommunicationDisabilityLevel   = "communication_disability_level"
	FileColumnIndividualCommunityID                    = "community_id"
	FileColumnIndividualCountryID                      = "country_id"
	FileColumnIndividualCreatedAt                      = "created_at"
	FileColumnIndividualDisplacementStatus             = "displacement_status"
	FileColumnIndividualDisplacementStatusComment      = "displacement_status_comment"
	FileColumnIndividualEmail1                         = "email_1"
	FileColumnIndividualEmail2                         = "email_2"
	FileColumnIndividualEmail3                         = "email_3"
	FileColumnIndividualFullName                       = "full_name"
	FileColumnIndividualFirstName                      = "first_name"
	FileColumnIndividualMiddleName                     = "middle_name"
	FileColumnIndividualLastName                       = "last_name"
	FileColumnIndividualFreeField1                     = "free_field_1"
	FileColumnIndividualFreeField2                     = "free_field_2"
	FileColumnIndividualFreeField3                     = "free_field_3"
	FileColumnIndividualFreeField4                     = "free_field_4"
	FileColumnIndividualFreeField5                     = "free_field_5"
	FileColumnIndividualSex                            = "sex"
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
	FileColumnIndividualEngagementContext              = "engagement_context"
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
	FileColumnIndividualIsFemaleHeadedHousehold        = "is_female_headed_household"
	FileColumnIndividualIsMinorHeadedHousehold         = "is_minor_headed_household"
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
	FileColumnIndividualServiceCC1                     = "service_cc_1"
	FileColumnIndividualServiceRequestedDate1          = "service_requested_date_1"
	FileColumnIndividualServiceDeliveredDate1          = "service_delivered_date_1"
	FileColumnIndividualServiceComments1               = "service_comments_1"
	FileColumnIndividualServiceCC2                     = "service_cc_2"
	FileColumnIndividualServiceRequestedDate2          = "service_requested_date_2"
	FileColumnIndividualServiceDeliveredDate2          = "service_delivered_date_2"
	FileColumnIndividualServiceComments2               = "service_comments_2"
	FileColumnIndividualServiceCC3                     = "service_cc_3"
	FileColumnIndividualServiceRequestedDate3          = "service_requested_date_3"
	FileColumnIndividualServiceDeliveredDate3          = "service_delivered_date_3"
	FileColumnIndividualServiceComments3               = "service_comments_3"
	FileColumnIndividualServiceCC4                     = "service_cc_4"
	FileColumnIndividualServiceRequestedDate4          = "service_requested_date_4"
	FileColumnIndividualServiceDeliveredDate4          = "service_delivered_date_4"
	FileColumnIndividualServiceComments4               = "service_comments_4"
	FileColumnIndividualServiceCC5                     = "service_cc_5"
	FileColumnIndividualServiceRequestedDate5          = "service_requested_date_5"
	FileColumnIndividualServiceDeliveredDate5          = "service_delivered_date_5"
	FileColumnIndividualServiceComments5               = "service_comments_5"
	FileColumnIndividualServiceCC6                     = "service_cc_6"
	FileColumnIndividualServiceRequestedDate6          = "service_requested_date_6"
	FileColumnIndividualServiceDeliveredDate6          = "service_delivered_date_6"
	FileColumnIndividualServiceComments6               = "service_comments_6"
	FileColumnIndividualServiceCC7                     = "service_cc_7"
	FileColumnIndividualServiceRequestedDate7          = "service_requested_date_7"
	FileColumnIndividualServiceDeliveredDate7          = "service_delivered_date_7"
	FileColumnIndividualServiceComments7               = "service_comments_7"
)

var IndividualDBColumns = containers.NewStringSet(
	DBColumnIndividualInactive,
	DBColumnIndividualAddress,
	DBColumnIndividualAge,
	DBColumnIndividualBirthDate,
	DBColumnIndividualCognitiveDisabilityLevel,
	DBColumnIndividualCollectionAdministrativeArea1,
	DBColumnIndividualCollectionAdministrativeArea2,
	DBColumnIndividualCollectionAdministrativeArea3,
	DBColumnIndividualCollectionOffice,
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
	DBColumnIndividualDisplacementStatusComment,
	DBColumnIndividualEmail1,
	DBColumnIndividualEmail2,
	DBColumnIndividualEmail3,
	DBColumnIndividualFullName,
	DBColumnIndividualFirstName,
	DBColumnIndividualMiddleName,
	DBColumnIndividualLastName,
	DBColumnIndividualFreeField1,
	DBColumnIndividualFreeField2,
	DBColumnIndividualFreeField3,
	DBColumnIndividualFreeField4,
	DBColumnIndividualFreeField5,
	DBColumnIndividualSex,
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
	DBColumnIndividualEngagementContext,
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
	DBColumnIndividualIsFemaleHeadedHousehold,
	DBColumnIndividualIsMinorHeadedHousehold,
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
	DBColumnIndividualServiceCC1,
	DBColumnIndividualServiceRequestedDate1,
	DBColumnIndividualServiceDeliveredDate1,
	DBColumnIndividualServiceComments1,
	DBColumnIndividualServiceCC2,
	DBColumnIndividualServiceRequestedDate2,
	DBColumnIndividualServiceDeliveredDate2,
	DBColumnIndividualServiceComments2,
	DBColumnIndividualServiceCC3,
	DBColumnIndividualServiceRequestedDate3,
	DBColumnIndividualServiceDeliveredDate3,
	DBColumnIndividualServiceComments3,
	DBColumnIndividualServiceCC4,
	DBColumnIndividualServiceRequestedDate4,
	DBColumnIndividualServiceDeliveredDate4,
	DBColumnIndividualServiceComments4,
	DBColumnIndividualServiceCC5,
	DBColumnIndividualServiceRequestedDate5,
	DBColumnIndividualServiceDeliveredDate5,
	DBColumnIndividualServiceComments5,
	DBColumnIndividualServiceCC6,
	DBColumnIndividualServiceRequestedDate6,
	DBColumnIndividualServiceDeliveredDate6,
	DBColumnIndividualServiceComments6,
	DBColumnIndividualServiceCC7,
	DBColumnIndividualServiceRequestedDate7,
	DBColumnIndividualServiceDeliveredDate7,
	DBColumnIndividualServiceComments7,
)

// Ordering is important
var IndividualFileColumns = []string{
	FileColumnIndividualID,
	FileColumnIndividualFullName,
	FileColumnIndividualPreferredName,
	FileColumnIndividualFirstName,
	FileColumnIndividualMiddleName,
	FileColumnIndividualLastName,
	FileColumnIndividualPrefersToRemainAnonymous,
	FileColumnIndividualSex,
	FileColumnIndividualBirthDate,
	FileColumnIndividualAge,
	FileColumnIndividualIsMinor,
	FileColumnIndividualNationality1,
	FileColumnIndividualNationality2,
	FileColumnIndividualIdentificationNumber1,
	FileColumnIndividualIdentificationType1,
	FileColumnIndividualIdentificationTypeExplanation1,
	FileColumnIndividualIdentificationNumber2,
	FileColumnIndividualIdentificationType2,
	FileColumnIndividualIdentificationTypeExplanation2,
	FileColumnIndividualIdentificationNumber3,
	FileColumnIndividualIdentificationType3,
	FileColumnIndividualIdentificationTypeExplanation3,
	FileColumnIndividualInternalID,
	FileColumnIndividualHouseholdID,
	FileColumnIndividualIsHeadOfHousehold,
	FileColumnIndividualIsFemaleHeadedHousehold,
	FileColumnIndividualIsMinorHeadedHousehold,
	FileColumnIndividualCommunityID,
	FileColumnIndividualIsHeadOfCommunity,
	FileColumnIndividualSpokenLanguage1,
	FileColumnIndividualSpokenLanguage2,
	FileColumnIndividualSpokenLanguage3,
	FileColumnIndividualPreferredCommunicationLanguage,
	FileColumnIndividualPhoneNumber1,
	FileColumnIndividualPhoneNumber2,
	FileColumnIndividualPhoneNumber3,
	FileColumnIndividualEmail1,
	FileColumnIndividualEmail2,
	FileColumnIndividualEmail3,
	FileColumnIndividualAddress,
	FileColumnIndividualPreferredContactMethod,
	FileColumnIndividualPreferredContactMethodComments,
	FileColumnIndividualHasConsentedToRGPD,
	FileColumnIndividualHasConsentedToReferral,
	FileColumnIndividualPresentsProtectionConcerns,
	FileColumnIndividualDisplacementStatus,
	FileColumnIndividualDisplacementStatusComment,
	FileColumnIndividualHasVisionDisability,
	FileColumnIndividualVisionDisabilityLevel,
	FileColumnIndividualHasHearingDisability,
	FileColumnIndividualHearingDisabilityLevel,
	FileColumnIndividualHasMobilityDisability,
	FileColumnIndividualMobilityDisabilityLevel,
	FileColumnIndividualHasCognitiveDisability,
	FileColumnIndividualCognitiveDisabilityLevel,
	FileColumnIndividualHasSelfCareDisability,
	FileColumnIndividualSelfCareDisabilityLevel,
	FileColumnIndividualHasCommunicationDisability,
	FileColumnIndividualCommunicationDisabilityLevel,
	FileColumnIndividualEngagementContext,
	FileColumnIndividualComments,
	FileColumnIndividualFreeField1,
	FileColumnIndividualFreeField2,
	FileColumnIndividualFreeField3,
	FileColumnIndividualFreeField4,
	FileColumnIndividualFreeField5,
	FileColumnIndividualCollectionAgentName,
	FileColumnIndividualCollectionAgentTitle,
	FileColumnIndividualCollectionTime,
	FileColumnIndividualCollectionAdministrativeArea1,
	FileColumnIndividualCollectionAdministrativeArea2,
	FileColumnIndividualCollectionAdministrativeArea3,
	FileColumnIndividualCollectionOffice,
	FileColumnIndividualServiceCC1,
	FileColumnIndividualServiceRequestedDate1,
	FileColumnIndividualServiceDeliveredDate1,
	FileColumnIndividualServiceComments1,
	FileColumnIndividualServiceCC2,
	FileColumnIndividualServiceRequestedDate2,
	FileColumnIndividualServiceDeliveredDate2,
	FileColumnIndividualServiceComments2,
	FileColumnIndividualServiceCC3,
	FileColumnIndividualServiceRequestedDate3,
	FileColumnIndividualServiceDeliveredDate3,
	FileColumnIndividualServiceComments3,
	FileColumnIndividualServiceCC4,
	FileColumnIndividualServiceRequestedDate4,
	FileColumnIndividualServiceDeliveredDate4,
	FileColumnIndividualServiceComments4,
	FileColumnIndividualServiceCC5,
	FileColumnIndividualServiceRequestedDate5,
	FileColumnIndividualServiceDeliveredDate5,
	FileColumnIndividualServiceComments5,
	FileColumnIndividualServiceCC6,
	FileColumnIndividualServiceRequestedDate6,
	FileColumnIndividualServiceDeliveredDate6,
	FileColumnIndividualServiceComments6,
	FileColumnIndividualServiceCC7,
	FileColumnIndividualServiceRequestedDate7,
	FileColumnIndividualServiceDeliveredDate7,
	FileColumnIndividualServiceComments7,
	FileColumnIndividualInactive,
	FileColumnIndividualCreatedAt,
	FileColumnIndividualUpdatedAt,
}

var IndividualDBToFileMap = map[string]string{
	DBColumnIndividualInactive:                       FileColumnIndividualInactive,
	DBColumnIndividualAddress:                        FileColumnIndividualAddress,
	DBColumnIndividualAge:                            FileColumnIndividualAge,
	DBColumnIndividualBirthDate:                      FileColumnIndividualBirthDate,
	DBColumnIndividualCognitiveDisabilityLevel:       FileColumnIndividualCognitiveDisabilityLevel,
	DBColumnIndividualCollectionAdministrativeArea1:  FileColumnIndividualCollectionAdministrativeArea1,
	DBColumnIndividualCollectionAdministrativeArea2:  FileColumnIndividualCollectionAdministrativeArea2,
	DBColumnIndividualCollectionAdministrativeArea3:  FileColumnIndividualCollectionAdministrativeArea3,
	DBColumnIndividualCollectionOffice:               FileColumnIndividualCollectionOffice,
	DBColumnIndividualCollectionAgentName:            FileColumnIndividualCollectionAgentName,
	DBColumnIndividualCollectionAgentTitle:           FileColumnIndividualCollectionAgentTitle,
	DBColumnIndividualCollectionTime:                 FileColumnIndividualCollectionTime,
	DBColumnIndividualComments:                       FileColumnIndividualComments,
	DBColumnIndividualCommunicationDisabilityLevel:   FileColumnIndividualCommunicationDisabilityLevel,
	DBColumnIndividualCommunityID:                    FileColumnIndividualCommunityID,
	DBColumnIndividualCountryID:                      FileColumnIndividualCountryID,
	DBColumnIndividualCreatedAt:                      FileColumnIndividualCreatedAt,
	DBColumnIndividualDisplacementStatus:             FileColumnIndividualDisplacementStatus,
	DBColumnIndividualDisplacementStatusComment:      FileColumnIndividualDisplacementStatusComment,
	DBColumnIndividualEmail1:                         FileColumnIndividualEmail1,
	DBColumnIndividualEmail2:                         FileColumnIndividualEmail2,
	DBColumnIndividualEmail3:                         FileColumnIndividualEmail3,
	DBColumnIndividualFullName:                       FileColumnIndividualFullName,
	DBColumnIndividualFirstName:                      FileColumnIndividualFirstName,
	DBColumnIndividualMiddleName:                     FileColumnIndividualMiddleName,
	DBColumnIndividualLastName:                       FileColumnIndividualLastName,
	DBColumnIndividualFreeField1:                     FileColumnIndividualFreeField1,
	DBColumnIndividualFreeField2:                     FileColumnIndividualFreeField2,
	DBColumnIndividualFreeField3:                     FileColumnIndividualFreeField3,
	DBColumnIndividualFreeField4:                     FileColumnIndividualFreeField4,
	DBColumnIndividualFreeField5:                     FileColumnIndividualFreeField5,
	DBColumnIndividualSex:                            FileColumnIndividualSex,
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
	DBColumnIndividualEngagementContext:              FileColumnIndividualEngagementContext,
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
	DBColumnIndividualIsFemaleHeadedHousehold:        FileColumnIndividualIsFemaleHeadedHousehold,
	DBColumnIndividualIsMinorHeadedHousehold:         FileColumnIndividualIsMinorHeadedHousehold,
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
	DBColumnIndividualServiceCC1:                     FileColumnIndividualServiceCC1,
	DBColumnIndividualServiceRequestedDate1:          FileColumnIndividualServiceRequestedDate1,
	DBColumnIndividualServiceDeliveredDate1:          FileColumnIndividualServiceDeliveredDate1,
	DBColumnIndividualServiceComments1:               FileColumnIndividualServiceComments1,
	DBColumnIndividualServiceCC2:                     FileColumnIndividualServiceCC2,
	DBColumnIndividualServiceRequestedDate2:          FileColumnIndividualServiceRequestedDate2,
	DBColumnIndividualServiceDeliveredDate2:          FileColumnIndividualServiceDeliveredDate2,
	DBColumnIndividualServiceComments2:               FileColumnIndividualServiceComments2,
	DBColumnIndividualServiceCC3:                     FileColumnIndividualServiceCC3,
	DBColumnIndividualServiceRequestedDate3:          FileColumnIndividualServiceRequestedDate3,
	DBColumnIndividualServiceDeliveredDate3:          FileColumnIndividualServiceDeliveredDate3,
	DBColumnIndividualServiceComments3:               FileColumnIndividualServiceComments3,
	DBColumnIndividualServiceCC4:                     FileColumnIndividualServiceCC4,
	DBColumnIndividualServiceRequestedDate4:          FileColumnIndividualServiceRequestedDate4,
	DBColumnIndividualServiceDeliveredDate4:          FileColumnIndividualServiceDeliveredDate4,
	DBColumnIndividualServiceComments4:               FileColumnIndividualServiceComments4,
	DBColumnIndividualServiceCC5:                     FileColumnIndividualServiceCC5,
	DBColumnIndividualServiceRequestedDate5:          FileColumnIndividualServiceRequestedDate5,
	DBColumnIndividualServiceDeliveredDate5:          FileColumnIndividualServiceDeliveredDate5,
	DBColumnIndividualServiceComments5:               FileColumnIndividualServiceComments5,
	DBColumnIndividualServiceCC6:                     FileColumnIndividualServiceCC6,
	DBColumnIndividualServiceRequestedDate6:          FileColumnIndividualServiceRequestedDate6,
	DBColumnIndividualServiceDeliveredDate6:          FileColumnIndividualServiceDeliveredDate6,
	DBColumnIndividualServiceComments6:               FileColumnIndividualServiceComments6,
	DBColumnIndividualServiceCC7:                     FileColumnIndividualServiceCC7,
	DBColumnIndividualServiceRequestedDate7:          FileColumnIndividualServiceRequestedDate7,
	DBColumnIndividualServiceDeliveredDate7:          FileColumnIndividualServiceDeliveredDate7,
	DBColumnIndividualServiceComments7:               FileColumnIndividualServiceComments7,
}

var IndividualFileToDBMap = map[string]string{
	FileColumnIndividualInactive:                       DBColumnIndividualInactive,
	FileColumnIndividualAddress:                        DBColumnIndividualAddress,
	FileColumnIndividualAge:                            DBColumnIndividualAge,
	FileColumnIndividualBirthDate:                      DBColumnIndividualBirthDate,
	FileColumnIndividualCognitiveDisabilityLevel:       DBColumnIndividualCognitiveDisabilityLevel,
	FileColumnIndividualCollectionAdministrativeArea1:  DBColumnIndividualCollectionAdministrativeArea1,
	FileColumnIndividualCollectionAdministrativeArea2:  DBColumnIndividualCollectionAdministrativeArea2,
	FileColumnIndividualCollectionAdministrativeArea3:  DBColumnIndividualCollectionAdministrativeArea3,
	FileColumnIndividualCollectionOffice:               DBColumnIndividualCollectionOffice,
	FileColumnIndividualCollectionAgentName:            DBColumnIndividualCollectionAgentName,
	FileColumnIndividualCollectionAgentTitle:           DBColumnIndividualCollectionAgentTitle,
	FileColumnIndividualCollectionTime:                 DBColumnIndividualCollectionTime,
	FileColumnIndividualComments:                       DBColumnIndividualComments,
	FileColumnIndividualCommunicationDisabilityLevel:   DBColumnIndividualCommunicationDisabilityLevel,
	FileColumnIndividualCommunityID:                    DBColumnIndividualCommunityID,
	FileColumnIndividualCountryID:                      DBColumnIndividualCountryID,
	FileColumnIndividualDisplacementStatus:             DBColumnIndividualDisplacementStatus,
	FileColumnIndividualDisplacementStatusComment:      DBColumnIndividualDisplacementStatusComment,
	FileColumnIndividualEmail1:                         DBColumnIndividualEmail1,
	FileColumnIndividualEmail2:                         DBColumnIndividualEmail2,
	FileColumnIndividualEmail3:                         DBColumnIndividualEmail3,
	FileColumnIndividualFullName:                       DBColumnIndividualFullName,
	FileColumnIndividualFirstName:                      DBColumnIndividualFirstName,
	FileColumnIndividualMiddleName:                     DBColumnIndividualMiddleName,
	FileColumnIndividualLastName:                       DBColumnIndividualLastName,
	FileColumnIndividualFreeField1:                     DBColumnIndividualFreeField1,
	FileColumnIndividualFreeField2:                     DBColumnIndividualFreeField2,
	FileColumnIndividualFreeField3:                     DBColumnIndividualFreeField3,
	FileColumnIndividualFreeField4:                     DBColumnIndividualFreeField4,
	FileColumnIndividualFreeField5:                     DBColumnIndividualFreeField5,
	FileColumnIndividualSex:                            DBColumnIndividualSex,
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
	FileColumnIndividualEngagementContext:              DBColumnIndividualEngagementContext,
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
	FileColumnIndividualIsFemaleHeadedHousehold:        DBColumnIndividualIsFemaleHeadedHousehold,
	FileColumnIndividualIsMinorHeadedHousehold:         DBColumnIndividualIsMinorHeadedHousehold,
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
	FileColumnIndividualServiceCC1:                     DBColumnIndividualServiceCC1,
	FileColumnIndividualServiceRequestedDate1:          DBColumnIndividualServiceRequestedDate1,
	FileColumnIndividualServiceDeliveredDate1:          DBColumnIndividualServiceDeliveredDate1,
	FileColumnIndividualServiceComments1:               DBColumnIndividualServiceComments1,
	FileColumnIndividualServiceCC2:                     DBColumnIndividualServiceCC2,
	FileColumnIndividualServiceRequestedDate2:          DBColumnIndividualServiceRequestedDate2,
	FileColumnIndividualServiceDeliveredDate2:          DBColumnIndividualServiceDeliveredDate2,
	FileColumnIndividualServiceComments2:               DBColumnIndividualServiceComments2,
	FileColumnIndividualServiceCC3:                     DBColumnIndividualServiceCC3,
	FileColumnIndividualServiceRequestedDate3:          DBColumnIndividualServiceRequestedDate3,
	FileColumnIndividualServiceDeliveredDate3:          DBColumnIndividualServiceDeliveredDate3,
	FileColumnIndividualServiceComments3:               DBColumnIndividualServiceComments3,
	FileColumnIndividualServiceCC4:                     DBColumnIndividualServiceCC4,
	FileColumnIndividualServiceRequestedDate4:          DBColumnIndividualServiceRequestedDate4,
	FileColumnIndividualServiceDeliveredDate4:          DBColumnIndividualServiceDeliveredDate4,
	FileColumnIndividualServiceComments4:               DBColumnIndividualServiceComments4,
	FileColumnIndividualServiceCC5:                     DBColumnIndividualServiceCC5,
	FileColumnIndividualServiceRequestedDate5:          DBColumnIndividualServiceRequestedDate5,
	FileColumnIndividualServiceDeliveredDate5:          DBColumnIndividualServiceDeliveredDate5,
	FileColumnIndividualServiceComments5:               DBColumnIndividualServiceComments5,
	FileColumnIndividualServiceCC6:                     DBColumnIndividualServiceCC6,
	FileColumnIndividualServiceRequestedDate6:          DBColumnIndividualServiceRequestedDate6,
	FileColumnIndividualServiceDeliveredDate6:          DBColumnIndividualServiceDeliveredDate6,
	FileColumnIndividualServiceComments6:               DBColumnIndividualServiceComments6,
	FileColumnIndividualServiceCC7:                     DBColumnIndividualServiceCC7,
	FileColumnIndividualServiceRequestedDate7:          DBColumnIndividualServiceRequestedDate7,
	FileColumnIndividualServiceDeliveredDate7:          DBColumnIndividualServiceDeliveredDate7,
	FileColumnIndividualServiceComments7:               DBColumnIndividualServiceComments7,
}

var IndividualSystemFileColumns = containers.NewStringSet(
	FileColumnIndividualCreatedAt,
	FileColumnIndividualUpdatedAt,
)
