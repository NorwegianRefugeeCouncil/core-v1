package constants

import (
	"github.com/nrc-no/notcore/internal/containers"
)

const (
	// Search Form
	FormParamsGetIndividualsAddress                         = "address"
	FormParamsGetIndividualsAgeFrom                         = "age_from"
	FormParamsGetIndividualsAgeTo                           = "age_to"
	FormParamsGetIndividualsBirthDateFrom                   = "birth_date_from"
	FormParamsGetIndividualsBirthDateTo                     = "birth_date_to"
	FormParamsGetIndividualsCognitiveDisabilityLevel        = "cognitive_disability_level"
	FormParamsGetIndividualsCollectionAdministrativeArea1   = "collection_administrative_area_1"
	FormParamsGetIndividualsCollectionAdministrativeArea2   = "collection_administrative_area_2"
	FormParamsGetIndividualsCollectionAdministrativeArea3   = "collection_administrative_area_3"
	FormParamsGetIndividualsCollectionAgentName             = "collection_agent_name"
	FormParamsGetIndividualsCollectionAgentTitle            = "collection_agent_title"
	FormParamsGetIndividualsCollectionOffice                = "collection_office"
	FormParamsGetIndividualsCollectionTimeFrom              = "collection_time_from"
	FormParamsGetIndividualsCollectionTimeTo                = "collection_time_to"
	FormParamsGetIndividualsCommunityID                     = "community_id"
	FormParamsGetIndividualsCountryID                       = "country_id"
	FormParamsGetIndividualsCreatedAtFrom                   = "created_at_from"
	FormParamsGetIndividualsCreatedAtTo                     = "created_at_to"
	FormParamsGetIndividualsDisplacementStatus              = "displacement_status"
	FormParamsGetIndividualsEmail                           = "email"
	FormParamsGetIndividualsEngagementContext               = "engagement_context"
	FormParamsGetIndividualsFreeField1                      = "free_field_1"
	FormParamsGetIndividualsFreeField2                      = "free_field_2"
	FormParamsGetIndividualsFreeField3                      = "free_field_3"
	FormParamsGetIndividualsFreeField4                      = "free_field_4"
	FormParamsGetIndividualsFreeField5                      = "free_field_5"
	FormParamsGetIndividualsFullName                        = "full_name"
	FormParamsGetIndividualsHasCognitiveDisability          = "has_cognitive_disability"
	FormParamsGetIndividualsHasCommunicationDisability      = "has_communication_disability"
	FormParamsGetIndividualsHasConsentedToReferral          = "has_consented_to_referral"
	FormParamsGetIndividualsHasConsentedToRgpd              = "has_consented_to_rgpd"
	FormParamsGetIndividualsHasDisability                   = "has_disability"
	FormParamsGetIndividualsHasHearingDisability            = "has_hearing_disability"
	FormParamsGetIndividualsHasMedicalCondition             = "has_medical_condition"
	FormParamsGetIndividualsHasMobilityDisability           = "has_mobility_disability"
	FormParamsGetIndividualsHasSelfCareDisability           = "has_selfcare_disability"
	FormParamsGetIndividualsHasVisionDisability             = "has_vision_disability"
	FormParamsGetIndividualsHearingDisabilityLevel          = "hearing_disability_level"
	FormParamsGetIndividualsHouseholdID                     = "household_id"
	FormParamsGetIndividualsID                              = "id"
	FormParamsGetIndividualsIdentificationNumber            = "identification_number"
	FormParamsGetIndividualsInactive                        = "inactive"
	FormParamsGetIndividualsInternalID                      = "internal_id"
	FormParamsGetIndividualsIsChildAtRisk                   = "is_child_at_risk"
	FormParamsGetIndividualsIsElderAtRisk                   = "is_elder_at_risk"
	FormParamsGetIndividualsIsFemaleHeadedHousehold         = "is_female_headed_household"
	FormParamsGetIndividualsIsHeadOfCommunity               = "is_head_of_community"
	FormParamsGetIndividualsIsHeadOfHousehold               = "is_head_of_household"
	FormParamsGetIndividualsIsLactating                     = "is_lactating"
	FormParamsGetIndividualsIsMinor                         = "is_minor"
	FormParamsGetIndividualsIsMinorHeadedHousehold          = "is_minor_headed_household"
	FormParamsGetIndividualsIsPregnant                      = "is_pregnant"
	FormParamsGetIndividualsIsSeparatedChild                = "is_separated_child"
	FormParamsGetIndividualsIsSingleParent                  = "is_single_parent"
	FormParamsGetIndividualsIsWomanAtRisk                   = "is_woman_at_risk"
	FormParamsGetIndividualsMobilityDisabilityLevel         = "mobility_disability_level"
	FormParamsGetIndividualsMothersName                     = "mothers_name"
	FormParamsGetIndividualsNationality                     = "nationality"
	FormParamsGetIndividualsNeedsLegalAndPhysicalProtection = "needs_legal_and_physical_protection"
	FormParamsGetIndividualsPhoneNumber                     = "phone_number"
	FormParamsGetIndividualsPreferredCommunicationLanguage  = "preferred_communication_language"
	FormParamsGetIndividualsPreferredContactMethod          = "preferred_contact_method"
	FormParamsGetIndividualsPrefersToRemainAnonymous        = "prefers_to_remain_anonymous"
	FormParamsGetIndividualsPresentsProtectionConcerns      = "presents_protection_concerns"
	FormParamsGetIndividualsSelfCareDisabilityLevel         = "selfcare_disability_level"
	FormParamsGetIndividualsServiceCC                       = "service_cc"
	FormParamsGetIndividualsServiceDeliveredDateFrom        = "service_delivered_date_from"
	FormParamsGetIndividualsServiceDeliveredDateTo          = "service_delivered_date_to"
	FormParamsGetIndividualsServiceRequestedDateFrom        = "service_requested_date_from"
	FormParamsGetIndividualsServiceRequestedDateTo          = "service_requested_date_to"
	FormParamsGetIndividualsSex                             = "sex"
	FormParamsGetIndividualsSkip                            = "skip"
	FormParamsGetIndividualsSort                            = "sort"
	FormParamsGetIndividualsSpokenLanguage                  = "spoken_language"
	FormParamsGetIndividualsTake                            = "take"
	FormParamsGetIndividualsUpdatedAtFrom                   = "updated_at_from"
	FormParamsGetIndividualsUpdatedAtTo                     = "updated_at_to"
	FormParamsGetIndividualsVisionDisabilityLevel           = "vision_disability_level"

	DBColumnIndividualAddress                         = "address"
	DBColumnIndividualAge                             = "age"
	DBColumnIndividualBirthDate                       = "birth_date"
	DBColumnIndividualCognitiveDisabilityLevel        = "cognitive_disability_level"
	DBColumnIndividualCollectionAdministrativeArea1   = "collection_administrative_area_1"
	DBColumnIndividualCollectionAdministrativeArea2   = "collection_administrative_area_2"
	DBColumnIndividualCollectionAdministrativeArea3   = "collection_administrative_area_3"
	DBColumnIndividualCollectionAgentName             = "collection_agent_name"
	DBColumnIndividualCollectionAgentTitle            = "collection_agent_title"
	DBColumnIndividualCollectionOffice                = "collection_office"
	DBColumnIndividualCollectionTime                  = "collection_time"
	DBColumnIndividualComments                        = "comments"
	DBColumnIndividualCommunicationDisabilityLevel    = "communication_disability_level"
	DBColumnIndividualCommunityID                     = "community_id"
	DBColumnIndividualCommunitySize                   = "community_size"
	DBColumnIndividualCountryID                       = "country_id"
	DBColumnIndividualCreatedAt                       = "created_at"
	DBColumnIndividualDeletedAt                       = "deleted_at"
	DBColumnIndividualDisplacementStatus              = "displacement_status"
	DBColumnIndividualDisplacementStatusComment       = "displacement_status_comment"
	DBColumnIndividualEmail1                          = "email_1"
	DBColumnIndividualEmail2                          = "email_2"
	DBColumnIndividualEmail3                          = "email_3"
	DBColumnIndividualEngagementContext               = "engagement_context"
	DBColumnIndividualFirstName                       = "first_name"
	DBColumnIndividualFreeField1                      = "free_field_1"
	DBColumnIndividualFreeField2                      = "free_field_2"
	DBColumnIndividualFreeField3                      = "free_field_3"
	DBColumnIndividualFreeField4                      = "free_field_4"
	DBColumnIndividualFreeField5                      = "free_field_5"
	DBColumnIndividualFullName                        = "full_name"
	DBColumnIndividualHasCognitiveDisability          = "has_cognitive_disability"
	DBColumnIndividualHasCommunicationDisability      = "has_communication_disability"
	DBColumnIndividualHasConsentedToRGPD              = "has_consented_to_rgpd"
	DBColumnIndividualHasConsentedToReferral          = "has_consented_to_referral"
	DBColumnIndividualHasDisability                   = "has_disability"
	DBColumnIndividualHasHearingDisability            = "has_hearing_disability"
	DBColumnIndividualHasMedicalCondition             = "has_medical_condition"
	DBColumnIndividualHasMobilityDisability           = "has_mobility_disability"
	DBColumnIndividualHasSelfCareDisability           = "has_selfcare_disability"
	DBColumnIndividualHasVisionDisability             = "has_vision_disability"
	DBColumnIndividualHearingDisabilityLevel          = "hearing_disability_level"
	DBColumnIndividualHouseholdID                     = "household_id"
	DBColumnIndividualHouseholdSize                   = "household_size"
	DBColumnIndividualID                              = "id"
	DBColumnIndividualIdentificationNumber1           = "identification_number_1"
	DBColumnIndividualIdentificationNumber2           = "identification_number_2"
	DBColumnIndividualIdentificationNumber3           = "identification_number_3"
	DBColumnIndividualIdentificationType1             = "identification_type_1"
	DBColumnIndividualIdentificationType2             = "identification_type_2"
	DBColumnIndividualIdentificationType3             = "identification_type_3"
	DBColumnIndividualIdentificationTypeExplanation1  = "identification_type_explanation_1"
	DBColumnIndividualIdentificationTypeExplanation2  = "identification_type_explanation_2"
	DBColumnIndividualIdentificationTypeExplanation3  = "identification_type_explanation_3"
	DBColumnIndividualInactive                        = "inactive"
	DBColumnIndividualInternalID                      = "internal_id"
	DBColumnIndividualIsChildAtRisk                   = "is_child_at_risk"
	DBColumnIndividualIsElderAtRisk                   = "is_elder_at_risk"
	DBColumnIndividualIsFemaleHeadedHousehold         = "is_female_headed_household"
	DBColumnIndividualIsHeadOfCommunity               = "is_head_of_community"
	DBColumnIndividualIsHeadOfHousehold               = "is_head_of_household"
	DBColumnIndividualIsLactating                     = "is_lactating"
	DBColumnIndividualIsMinor                         = "is_minor"
	DBColumnIndividualIsMinorHeadedHousehold          = "is_minor_headed_household"
	DBColumnIndividualIsPregnant                      = "is_pregnant"
	DBColumnIndividualIsSeparatedChild                = "is_separated_child"
	DBColumnIndividualIsSingleParent                  = "is_single_parent"
	DBColumnIndividualIsWomanAtRisk                   = "is_woman_at_risk"
	DBColumnIndividualLastName                        = "last_name"
	DBColumnIndividualMiddleName                      = "middle_name"
	DBColumnIndividualMobilityDisabilityLevel         = "mobility_disability_level"
	DBColumnIndividualMothersName                     = "mothers_name"
	DBColumnIndividualNationality1                    = "nationality_1"
	DBColumnIndividualNationality2                    = "nationality_2"
	DBColumnIndividualNativeName                      = "native_name"
	DBColumnIndividualNeedsLegalAndPhysicalProtection = "needs_legal_and_physical_protection"
	DBColumnIndividualNormalizedPhoneNumber1          = "normalized_phone_number_1"
	DBColumnIndividualNormalizedPhoneNumber2          = "normalized_phone_number_2"
	DBColumnIndividualNormalizedPhoneNumber3          = "normalized_phone_number_3"
	DBColumnIndividualPWDComments                     = "pwd_comments"
	DBColumnIndividualPhoneNumber1                    = "phone_number_1"
	DBColumnIndividualPhoneNumber2                    = "phone_number_2"
	DBColumnIndividualPhoneNumber3                    = "phone_number_3"
	DBColumnIndividualPreferredCommunicationLanguage  = "preferred_communication_language"
	DBColumnIndividualPreferredContactMethod          = "preferred_contact_method"
	DBColumnIndividualPreferredContactMethodComments  = "preferred_contact_method_comments"
	DBColumnIndividualPreferredName                   = "preferred_name"
	DBColumnIndividualPrefersToRemainAnonymous        = "prefers_to_remain_anonymous"
	DBColumnIndividualPresentsProtectionConcerns      = "presents_protection_concerns"
	DBColumnIndividualSelfCareDisabilityLevel         = "selfcare_disability_level"
	DBColumnIndividualServiceCC1                      = "service_cc_1"
	DBColumnIndividualServiceCC2                      = "service_cc_2"
	DBColumnIndividualServiceCC3                      = "service_cc_3"
	DBColumnIndividualServiceCC4                      = "service_cc_4"
	DBColumnIndividualServiceCC5                      = "service_cc_5"
	DBColumnIndividualServiceCC6                      = "service_cc_6"
	DBColumnIndividualServiceCC7                      = "service_cc_7"
	DBColumnIndividualServiceComments1                = "service_comments_1"
	DBColumnIndividualServiceComments2                = "service_comments_2"
	DBColumnIndividualServiceComments3                = "service_comments_3"
	DBColumnIndividualServiceComments4                = "service_comments_4"
	DBColumnIndividualServiceComments5                = "service_comments_5"
	DBColumnIndividualServiceComments6                = "service_comments_6"
	DBColumnIndividualServiceComments7                = "service_comments_7"
	DBColumnIndividualServiceDeliveredDate1           = "service_delivered_date_1"
	DBColumnIndividualServiceDeliveredDate2           = "service_delivered_date_2"
	DBColumnIndividualServiceDeliveredDate3           = "service_delivered_date_3"
	DBColumnIndividualServiceDeliveredDate4           = "service_delivered_date_4"
	DBColumnIndividualServiceDeliveredDate5           = "service_delivered_date_5"
	DBColumnIndividualServiceDeliveredDate6           = "service_delivered_date_6"
	DBColumnIndividualServiceDeliveredDate7           = "service_delivered_date_7"
	DBColumnIndividualServiceRequestedDate1           = "service_requested_date_1"
	DBColumnIndividualServiceRequestedDate2           = "service_requested_date_2"
	DBColumnIndividualServiceRequestedDate3           = "service_requested_date_3"
	DBColumnIndividualServiceRequestedDate4           = "service_requested_date_4"
	DBColumnIndividualServiceRequestedDate5           = "service_requested_date_5"
	DBColumnIndividualServiceRequestedDate6           = "service_requested_date_6"
	DBColumnIndividualServiceRequestedDate7           = "service_requested_date_7"
	DBColumnIndividualSex                             = "sex"
	DBColumnIndividualSpokenLanguage1                 = "spoken_language_1"
	DBColumnIndividualSpokenLanguage2                 = "spoken_language_2"
	DBColumnIndividualSpokenLanguage3                 = "spoken_language_3"
	DBColumnIndividualUpdatedAt                       = "updated_at"
	DBColumnIndividualVisionDisabilityLevel           = "vision_disability_level"
	DBColumnIndividualVulnerabilityComments           = "vulnerability_comments"

	FileColumnIndividualAddress                         = "file_address"
	FileColumnIndividualAge                             = "file_age"
	FileColumnIndividualBirthDate                       = "file_birth_date"
	FileColumnIndividualCognitiveDisabilityLevel        = "file_cognitive_disability_level"
	FileColumnIndividualCollectionAdministrativeArea1   = "file_collection_administrative_area_1"
	FileColumnIndividualCollectionAdministrativeArea2   = "file_collection_administrative_area_2"
	FileColumnIndividualCollectionAdministrativeArea3   = "file_collection_administrative_area_3"
	FileColumnIndividualCollectionAgentName             = "file_collection_agent_name"
	FileColumnIndividualCollectionAgentTitle            = "file_collection_agent_title"
	FileColumnIndividualCollectionOffice                = "file_collection_office"
	FileColumnIndividualCollectionTime                  = "file_collection_time"
	FileColumnIndividualComments                        = "file_comments"
	FileColumnIndividualCommunicationDisabilityLevel    = "file_communication_disability_level"
	FileColumnIndividualCommunityID                     = "file_community_id"
	FileColumnIndividualCommunitySize                   = "file_community_size"
	FileColumnIndividualCountryID                       = "file_country_id"
	FileColumnIndividualCreatedAt                       = "file_created_at"
	FileColumnIndividualDisplacementStatus              = "file_displacement_status"
	FileColumnIndividualDisplacementStatusComment       = "file_displacement_status_comment"
	FileColumnIndividualEmail1                          = "file_email_1"
	FileColumnIndividualEmail2                          = "file_email_2"
	FileColumnIndividualEmail3                          = "file_email_3"
	FileColumnIndividualEngagementContext               = "file_engagement_context"
	FileColumnIndividualFirstName                       = "file_first_name"
	FileColumnIndividualFreeField1                      = "file_free_field_1"
	FileColumnIndividualFreeField2                      = "file_free_field_2"
	FileColumnIndividualFreeField3                      = "file_free_field_3"
	FileColumnIndividualFreeField4                      = "file_free_field_4"
	FileColumnIndividualFreeField5                      = "file_free_field_5"
	FileColumnIndividualFullName                        = "file_full_name"
	FileColumnIndividualHasCognitiveDisability          = "file_has_cognitive_disability"
	FileColumnIndividualHasCommunicationDisability      = "file_has_communication_disability"
	FileColumnIndividualHasConsentedToRGPD              = "file_has_consented_to_rgpd"
	FileColumnIndividualHasConsentedToReferral          = "file_has_consented_to_referral"
	FileColumnIndividualHasDisability                   = "file_has_disability"
	FileColumnIndividualHasHearingDisability            = "file_has_hearing_disability"
	FileColumnIndividualHasMedicalCondition             = "file_has_medical_condition"
	FileColumnIndividualHasMobilityDisability           = "file_has_mobility_disability"
	FileColumnIndividualHasSelfCareDisability           = "file_has_selfcare_disability"
	FileColumnIndividualHasVisionDisability             = "file_has_vision_disability"
	FileColumnIndividualHearingDisabilityLevel          = "file_hearing_disability_level"
	FileColumnIndividualHouseholdID                     = "file_household_id"
	FileColumnIndividualHouseholdSize                   = "file_household_size"
	FileColumnIndividualID                              = "file_id"
	FileColumnIndividualIdentificationNumber1           = "file_identification_number_1"
	FileColumnIndividualIdentificationNumber2           = "file_identification_number_2"
	FileColumnIndividualIdentificationNumber3           = "file_identification_number_3"
	FileColumnIndividualIdentificationType1             = "file_identification_type_1"
	FileColumnIndividualIdentificationType2             = "file_identification_type_2"
	FileColumnIndividualIdentificationType3             = "file_identification_type_3"
	FileColumnIndividualIdentificationTypeExplanation1  = "file_identification_type_explanation_1"
	FileColumnIndividualIdentificationTypeExplanation2  = "file_identification_type_explanation_2"
	FileColumnIndividualIdentificationTypeExplanation3  = "file_identification_type_explanation_3"
	FileColumnIndividualInactive                        = "file_inactive"
	FileColumnIndividualInternalID                      = "file_internal_id"
	FileColumnIndividualIsChildAtRisk                   = "file_is_child_at_risk"
	FileColumnIndividualIsElderAtRisk                   = "file_is_elder_at_risk"
	FileColumnIndividualIsFemaleHeadedHousehold         = "file_is_female_headed_household"
	FileColumnIndividualIsHeadOfCommunity               = "file_is_head_of_community"
	FileColumnIndividualIsHeadOfHousehold               = "file_is_head_of_household"
	FileColumnIndividualIsLactating                     = "file_is_lactating"
	FileColumnIndividualIsMinor                         = "file_is_minor"
	FileColumnIndividualIsMinorHeadedHousehold          = "file_is_minor_headed_household"
	FileColumnIndividualIsPregnant                      = "file_is_pregnant"
	FileColumnIndividualIsSeparatedChild                = "file_is_separated_child"
	FileColumnIndividualIsSingleParent                  = "file_is_single_parent"
	FileColumnIndividualIsWomanAtRisk                   = "file_is_woman_at_risk"
	FileColumnIndividualLastName                        = "file_last_name"
	FileColumnIndividualMiddleName                      = "file_middle_name"
	FileColumnIndividualMobilityDisabilityLevel         = "file_mobility_disability_level"
	FileColumnIndividualMothersName                     = "file_mothers_name"
	FileColumnIndividualNationality1                    = "file_nationality_1"
	FileColumnIndividualNationality2                    = "file_nationality_2"
	FileColumnIndividualNativeName                      = "file_native_name"
	FileColumnIndividualNeedsLegalAndPhysicalProtection = "file_needs_legal_and_physical_protection"
	FileColumnIndividualPWDComments                     = "file_pwd_comments"
	FileColumnIndividualPhoneNumber1                    = "file_phone_number_1"
	FileColumnIndividualPhoneNumber2                    = "file_phone_number_2"
	FileColumnIndividualPhoneNumber3                    = "file_phone_number_3"
	FileColumnIndividualPreferredCommunicationLanguage  = "file_preferred_communication_language"
	FileColumnIndividualPreferredContactMethod          = "file_preferred_contact_method"
	FileColumnIndividualPreferredContactMethodComments  = "file_preferred_contact_method_comments"
	FileColumnIndividualPreferredName                   = "file_preferred_name"
	FileColumnIndividualPrefersToRemainAnonymous        = "file_prefers_to_remain_anonymous"
	FileColumnIndividualPresentsProtectionConcerns      = "file_presents_protection_concerns"
	FileColumnIndividualSelfCareDisabilityLevel         = "file_selfcare_disability_level"
	FileColumnIndividualServiceCC1                      = "file_service_cc_1"
	FileColumnIndividualServiceCC2                      = "file_service_cc_2"
	FileColumnIndividualServiceCC3                      = "file_service_cc_3"
	FileColumnIndividualServiceCC4                      = "file_service_cc_4"
	FileColumnIndividualServiceCC5                      = "file_service_cc_5"
	FileColumnIndividualServiceCC6                      = "file_service_cc_6"
	FileColumnIndividualServiceCC7                      = "file_service_cc_7"
	FileColumnIndividualServiceComments1                = "file_service_comments_1"
	FileColumnIndividualServiceComments2                = "file_service_comments_2"
	FileColumnIndividualServiceComments3                = "file_service_comments_3"
	FileColumnIndividualServiceComments4                = "file_service_comments_4"
	FileColumnIndividualServiceComments5                = "file_service_comments_5"
	FileColumnIndividualServiceComments6                = "file_service_comments_6"
	FileColumnIndividualServiceComments7                = "file_service_comments_7"
	FileColumnIndividualServiceDeliveredDate1           = "file_service_delivered_date_1"
	FileColumnIndividualServiceDeliveredDate2           = "file_service_delivered_date_2"
	FileColumnIndividualServiceDeliveredDate3           = "file_service_delivered_date_3"
	FileColumnIndividualServiceDeliveredDate4           = "file_service_delivered_date_4"
	FileColumnIndividualServiceDeliveredDate5           = "file_service_delivered_date_5"
	FileColumnIndividualServiceDeliveredDate6           = "file_service_delivered_date_6"
	FileColumnIndividualServiceDeliveredDate7           = "file_service_delivered_date_7"
	FileColumnIndividualServiceRequestedDate1           = "file_service_requested_date_1"
	FileColumnIndividualServiceRequestedDate2           = "file_service_requested_date_2"
	FileColumnIndividualServiceRequestedDate3           = "file_service_requested_date_3"
	FileColumnIndividualServiceRequestedDate4           = "file_service_requested_date_4"
	FileColumnIndividualServiceRequestedDate5           = "file_service_requested_date_5"
	FileColumnIndividualServiceRequestedDate6           = "file_service_requested_date_6"
	FileColumnIndividualServiceRequestedDate7           = "file_service_requested_date_7"
	FileColumnIndividualSex                             = "file_sex"
	FileColumnIndividualSpokenLanguage1                 = "file_spoken_language_1"
	FileColumnIndividualSpokenLanguage2                 = "file_spoken_language_2"
	FileColumnIndividualSpokenLanguage3                 = "file_spoken_language_3"
	FileColumnIndividualUpdatedAt                       = "file_updated_at"
	FileColumnIndividualVisionDisabilityLevel           = "file_vision_disability_level"
	FileColumnIndividualVulnerabilityComments           = "file_vulnerability_comments"
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
	DBColumnIndividualCollectionOffice,
	DBColumnIndividualCollectionTime,
	DBColumnIndividualComments,
	DBColumnIndividualCommunicationDisabilityLevel,
	DBColumnIndividualCommunityID,
	DBColumnIndividualCommunitySize,
	DBColumnIndividualCountryID,
	DBColumnIndividualCreatedAt,
	DBColumnIndividualDeletedAt,
	DBColumnIndividualDisplacementStatus,
	DBColumnIndividualDisplacementStatusComment,
	DBColumnIndividualEmail1,
	DBColumnIndividualEmail2,
	DBColumnIndividualEmail3,
	DBColumnIndividualEngagementContext,
	DBColumnIndividualFirstName,
	DBColumnIndividualFreeField1,
	DBColumnIndividualFreeField2,
	DBColumnIndividualFreeField3,
	DBColumnIndividualFreeField4,
	DBColumnIndividualFreeField5,
	DBColumnIndividualFullName,
	DBColumnIndividualHasCognitiveDisability,
	DBColumnIndividualHasCommunicationDisability,
	DBColumnIndividualHasConsentedToRGPD,
	DBColumnIndividualHasConsentedToReferral,
	DBColumnIndividualHasDisability,
	DBColumnIndividualHasHearingDisability,
	DBColumnIndividualHasMedicalCondition,
	DBColumnIndividualHasMobilityDisability,
	DBColumnIndividualHasSelfCareDisability,
	DBColumnIndividualHasVisionDisability,
	DBColumnIndividualHearingDisabilityLevel,
	DBColumnIndividualHouseholdID,
	DBColumnIndividualHouseholdSize,
	DBColumnIndividualID,
	DBColumnIndividualIdentificationNumber1,
	DBColumnIndividualIdentificationNumber2,
	DBColumnIndividualIdentificationNumber3,
	DBColumnIndividualIdentificationType1,
	DBColumnIndividualIdentificationType2,
	DBColumnIndividualIdentificationType3,
	DBColumnIndividualIdentificationTypeExplanation1,
	DBColumnIndividualIdentificationTypeExplanation2,
	DBColumnIndividualIdentificationTypeExplanation3,
	DBColumnIndividualInactive,
	DBColumnIndividualInternalID,
	DBColumnIndividualIsChildAtRisk,
	DBColumnIndividualIsElderAtRisk,
	DBColumnIndividualIsFemaleHeadedHousehold,
	DBColumnIndividualIsHeadOfCommunity,
	DBColumnIndividualIsHeadOfHousehold,
	DBColumnIndividualIsLactating,
	DBColumnIndividualIsMinor,
	DBColumnIndividualIsMinorHeadedHousehold,
	DBColumnIndividualIsPregnant,
	DBColumnIndividualIsSeparatedChild,
	DBColumnIndividualIsSingleParent,
	DBColumnIndividualIsWomanAtRisk,
	DBColumnIndividualLastName,
	DBColumnIndividualMiddleName,
	DBColumnIndividualMobilityDisabilityLevel,
	DBColumnIndividualMothersName,
	DBColumnIndividualNationality1,
	DBColumnIndividualNationality2,
	DBColumnIndividualNativeName,
	DBColumnIndividualNeedsLegalAndPhysicalProtection,
	DBColumnIndividualNormalizedPhoneNumber1,
	DBColumnIndividualNormalizedPhoneNumber2,
	DBColumnIndividualNormalizedPhoneNumber3,
	DBColumnIndividualPWDComments,
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
	DBColumnIndividualServiceCC1,
	DBColumnIndividualServiceCC2,
	DBColumnIndividualServiceCC3,
	DBColumnIndividualServiceCC4,
	DBColumnIndividualServiceCC5,
	DBColumnIndividualServiceCC6,
	DBColumnIndividualServiceCC7,
	DBColumnIndividualServiceComments1,
	DBColumnIndividualServiceComments2,
	DBColumnIndividualServiceComments3,
	DBColumnIndividualServiceComments4,
	DBColumnIndividualServiceComments5,
	DBColumnIndividualServiceComments6,
	DBColumnIndividualServiceComments7,
	DBColumnIndividualServiceDeliveredDate1,
	DBColumnIndividualServiceDeliveredDate2,
	DBColumnIndividualServiceDeliveredDate3,
	DBColumnIndividualServiceDeliveredDate4,
	DBColumnIndividualServiceDeliveredDate5,
	DBColumnIndividualServiceDeliveredDate6,
	DBColumnIndividualServiceDeliveredDate7,
	DBColumnIndividualServiceRequestedDate1,
	DBColumnIndividualServiceRequestedDate2,
	DBColumnIndividualServiceRequestedDate3,
	DBColumnIndividualServiceRequestedDate4,
	DBColumnIndividualServiceRequestedDate5,
	DBColumnIndividualServiceRequestedDate6,
	DBColumnIndividualServiceRequestedDate7,
	DBColumnIndividualSex,
	DBColumnIndividualSpokenLanguage1,
	DBColumnIndividualSpokenLanguage2,
	DBColumnIndividualSpokenLanguage3,
	DBColumnIndividualUpdatedAt,
	DBColumnIndividualVisionDisabilityLevel,
	DBColumnIndividualVulnerabilityComments,
)

// Ordering is important
var IndividualFileColumns = []string{
	FileColumnIndividualID,
	FileColumnIndividualFullName,
	FileColumnIndividualPreferredName,
	FileColumnIndividualFirstName,
	FileColumnIndividualMiddleName,
	FileColumnIndividualLastName,
	FileColumnIndividualNativeName,
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
	FileColumnIndividualHouseholdSize,
	FileColumnIndividualIsHeadOfHousehold,
	FileColumnIndividualIsFemaleHeadedHousehold,
	FileColumnIndividualIsMinorHeadedHousehold,
	FileColumnIndividualCommunityID,
	FileColumnIndividualCommunitySize,
	FileColumnIndividualIsHeadOfCommunity,
	FileColumnIndividualMothersName,
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
	FileColumnIndividualHasDisability,
	FileColumnIndividualPWDComments,
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
	FileColumnIndividualIsChildAtRisk,
	FileColumnIndividualIsWomanAtRisk,
	FileColumnIndividualIsElderAtRisk,
	FileColumnIndividualIsSeparatedChild,
	FileColumnIndividualIsSingleParent,
	FileColumnIndividualIsPregnant,
	FileColumnIndividualIsLactating,
	FileColumnIndividualHasMedicalCondition,
	FileColumnIndividualNeedsLegalAndPhysicalProtection,
	FileColumnIndividualVulnerabilityComments,
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

var IndividualFileToDBMap = map[string]string{
	FileColumnIndividualAddress:                         DBColumnIndividualAddress,
	FileColumnIndividualAge:                             DBColumnIndividualAge,
	FileColumnIndividualBirthDate:                       DBColumnIndividualBirthDate,
	FileColumnIndividualCognitiveDisabilityLevel:        DBColumnIndividualCognitiveDisabilityLevel,
	FileColumnIndividualCollectionAdministrativeArea1:   DBColumnIndividualCollectionAdministrativeArea1,
	FileColumnIndividualCollectionAdministrativeArea2:   DBColumnIndividualCollectionAdministrativeArea2,
	FileColumnIndividualCollectionAdministrativeArea3:   DBColumnIndividualCollectionAdministrativeArea3,
	FileColumnIndividualCollectionAgentName:             DBColumnIndividualCollectionAgentName,
	FileColumnIndividualCollectionAgentTitle:            DBColumnIndividualCollectionAgentTitle,
	FileColumnIndividualCollectionOffice:                DBColumnIndividualCollectionOffice,
	FileColumnIndividualCollectionTime:                  DBColumnIndividualCollectionTime,
	FileColumnIndividualComments:                        DBColumnIndividualComments,
	FileColumnIndividualCommunicationDisabilityLevel:    DBColumnIndividualCommunicationDisabilityLevel,
	FileColumnIndividualCommunityID:                     DBColumnIndividualCommunityID,
	FileColumnIndividualCommunitySize:                   DBColumnIndividualCommunitySize,
	FileColumnIndividualCountryID:                       DBColumnIndividualCountryID,
	FileColumnIndividualDisplacementStatus:              DBColumnIndividualDisplacementStatus,
	FileColumnIndividualDisplacementStatusComment:       DBColumnIndividualDisplacementStatusComment,
	FileColumnIndividualEmail1:                          DBColumnIndividualEmail1,
	FileColumnIndividualEmail2:                          DBColumnIndividualEmail2,
	FileColumnIndividualEmail3:                          DBColumnIndividualEmail3,
	FileColumnIndividualEngagementContext:               DBColumnIndividualEngagementContext,
	FileColumnIndividualFirstName:                       DBColumnIndividualFirstName,
	FileColumnIndividualFreeField1:                      DBColumnIndividualFreeField1,
	FileColumnIndividualFreeField2:                      DBColumnIndividualFreeField2,
	FileColumnIndividualFreeField3:                      DBColumnIndividualFreeField3,
	FileColumnIndividualFreeField4:                      DBColumnIndividualFreeField4,
	FileColumnIndividualFreeField5:                      DBColumnIndividualFreeField5,
	FileColumnIndividualFullName:                        DBColumnIndividualFullName,
	FileColumnIndividualHasCognitiveDisability:          DBColumnIndividualHasCognitiveDisability,
	FileColumnIndividualHasCommunicationDisability:      DBColumnIndividualHasCommunicationDisability,
	FileColumnIndividualHasConsentedToRGPD:              DBColumnIndividualHasConsentedToRGPD,
	FileColumnIndividualHasConsentedToReferral:          DBColumnIndividualHasConsentedToReferral,
	FileColumnIndividualHasDisability:                   DBColumnIndividualHasDisability,
	FileColumnIndividualHasHearingDisability:            DBColumnIndividualHasHearingDisability,
	FileColumnIndividualHasMedicalCondition:             DBColumnIndividualHasMedicalCondition,
	FileColumnIndividualHasMobilityDisability:           DBColumnIndividualHasMobilityDisability,
	FileColumnIndividualHasSelfCareDisability:           DBColumnIndividualHasSelfCareDisability,
	FileColumnIndividualHasVisionDisability:             DBColumnIndividualHasVisionDisability,
	FileColumnIndividualHearingDisabilityLevel:          DBColumnIndividualHearingDisabilityLevel,
	FileColumnIndividualHouseholdID:                     DBColumnIndividualHouseholdID,
	FileColumnIndividualHouseholdSize:                   DBColumnIndividualHouseholdSize,
	FileColumnIndividualID:                              DBColumnIndividualID,
	FileColumnIndividualIdentificationNumber1:           DBColumnIndividualIdentificationNumber1,
	FileColumnIndividualIdentificationNumber2:           DBColumnIndividualIdentificationNumber2,
	FileColumnIndividualIdentificationNumber3:           DBColumnIndividualIdentificationNumber3,
	FileColumnIndividualIdentificationType1:             DBColumnIndividualIdentificationType1,
	FileColumnIndividualIdentificationType2:             DBColumnIndividualIdentificationType2,
	FileColumnIndividualIdentificationType3:             DBColumnIndividualIdentificationType3,
	FileColumnIndividualIdentificationTypeExplanation1:  DBColumnIndividualIdentificationTypeExplanation1,
	FileColumnIndividualIdentificationTypeExplanation2:  DBColumnIndividualIdentificationTypeExplanation2,
	FileColumnIndividualIdentificationTypeExplanation3:  DBColumnIndividualIdentificationTypeExplanation3,
	FileColumnIndividualInactive:                        DBColumnIndividualInactive,
	FileColumnIndividualInternalID:                      DBColumnIndividualInternalID,
	FileColumnIndividualIsChildAtRisk:                   DBColumnIndividualIsChildAtRisk,
	FileColumnIndividualIsElderAtRisk:                   DBColumnIndividualIsElderAtRisk,
	FileColumnIndividualIsFemaleHeadedHousehold:         DBColumnIndividualIsFemaleHeadedHousehold,
	FileColumnIndividualIsHeadOfCommunity:               DBColumnIndividualIsHeadOfCommunity,
	FileColumnIndividualIsHeadOfHousehold:               DBColumnIndividualIsHeadOfHousehold,
	FileColumnIndividualIsLactating:                     DBColumnIndividualIsLactating,
	FileColumnIndividualIsMinor:                         DBColumnIndividualIsMinor,
	FileColumnIndividualIsMinorHeadedHousehold:          DBColumnIndividualIsMinorHeadedHousehold,
	FileColumnIndividualIsPregnant:                      DBColumnIndividualIsPregnant,
	FileColumnIndividualIsSeparatedChild:                DBColumnIndividualIsSeparatedChild,
	FileColumnIndividualIsSingleParent:                  DBColumnIndividualIsSingleParent,
	FileColumnIndividualIsWomanAtRisk:                   DBColumnIndividualIsWomanAtRisk,
	FileColumnIndividualLastName:                        DBColumnIndividualLastName,
	FileColumnIndividualMiddleName:                      DBColumnIndividualMiddleName,
	FileColumnIndividualMobilityDisabilityLevel:         DBColumnIndividualMobilityDisabilityLevel,
	FileColumnIndividualMothersName:                     DBColumnIndividualMothersName,
	FileColumnIndividualNationality1:                    DBColumnIndividualNationality1,
	FileColumnIndividualNationality2:                    DBColumnIndividualNationality2,
	FileColumnIndividualNativeName:                      DBColumnIndividualNativeName,
	FileColumnIndividualNeedsLegalAndPhysicalProtection: DBColumnIndividualNeedsLegalAndPhysicalProtection,
	FileColumnIndividualPWDComments:                     DBColumnIndividualPWDComments,
	FileColumnIndividualPhoneNumber1:                    DBColumnIndividualPhoneNumber1,
	FileColumnIndividualPhoneNumber2:                    DBColumnIndividualPhoneNumber2,
	FileColumnIndividualPhoneNumber3:                    DBColumnIndividualPhoneNumber3,
	FileColumnIndividualPreferredCommunicationLanguage:  DBColumnIndividualPreferredCommunicationLanguage,
	FileColumnIndividualPreferredContactMethod:          DBColumnIndividualPreferredContactMethod,
	FileColumnIndividualPreferredContactMethodComments:  DBColumnIndividualPreferredContactMethodComments,
	FileColumnIndividualPreferredName:                   DBColumnIndividualPreferredName,
	FileColumnIndividualPrefersToRemainAnonymous:        DBColumnIndividualPrefersToRemainAnonymous,
	FileColumnIndividualPresentsProtectionConcerns:      DBColumnIndividualPresentsProtectionConcerns,
	FileColumnIndividualSelfCareDisabilityLevel:         DBColumnIndividualSelfCareDisabilityLevel,
	FileColumnIndividualServiceCC1:                      DBColumnIndividualServiceCC1,
	FileColumnIndividualServiceCC2:                      DBColumnIndividualServiceCC2,
	FileColumnIndividualServiceCC3:                      DBColumnIndividualServiceCC3,
	FileColumnIndividualServiceCC4:                      DBColumnIndividualServiceCC4,
	FileColumnIndividualServiceCC5:                      DBColumnIndividualServiceCC5,
	FileColumnIndividualServiceCC6:                      DBColumnIndividualServiceCC6,
	FileColumnIndividualServiceCC7:                      DBColumnIndividualServiceCC7,
	FileColumnIndividualServiceComments1:                DBColumnIndividualServiceComments1,
	FileColumnIndividualServiceComments2:                DBColumnIndividualServiceComments2,
	FileColumnIndividualServiceComments3:                DBColumnIndividualServiceComments3,
	FileColumnIndividualServiceComments4:                DBColumnIndividualServiceComments4,
	FileColumnIndividualServiceComments5:                DBColumnIndividualServiceComments5,
	FileColumnIndividualServiceComments6:                DBColumnIndividualServiceComments6,
	FileColumnIndividualServiceComments7:                DBColumnIndividualServiceComments7,
	FileColumnIndividualServiceDeliveredDate1:           DBColumnIndividualServiceDeliveredDate1,
	FileColumnIndividualServiceDeliveredDate2:           DBColumnIndividualServiceDeliveredDate2,
	FileColumnIndividualServiceDeliveredDate3:           DBColumnIndividualServiceDeliveredDate3,
	FileColumnIndividualServiceDeliveredDate4:           DBColumnIndividualServiceDeliveredDate4,
	FileColumnIndividualServiceDeliveredDate5:           DBColumnIndividualServiceDeliveredDate5,
	FileColumnIndividualServiceDeliveredDate6:           DBColumnIndividualServiceDeliveredDate6,
	FileColumnIndividualServiceDeliveredDate7:           DBColumnIndividualServiceDeliveredDate7,
	FileColumnIndividualServiceRequestedDate1:           DBColumnIndividualServiceRequestedDate1,
	FileColumnIndividualServiceRequestedDate2:           DBColumnIndividualServiceRequestedDate2,
	FileColumnIndividualServiceRequestedDate3:           DBColumnIndividualServiceRequestedDate3,
	FileColumnIndividualServiceRequestedDate4:           DBColumnIndividualServiceRequestedDate4,
	FileColumnIndividualServiceRequestedDate5:           DBColumnIndividualServiceRequestedDate5,
	FileColumnIndividualServiceRequestedDate6:           DBColumnIndividualServiceRequestedDate6,
	FileColumnIndividualServiceRequestedDate7:           DBColumnIndividualServiceRequestedDate7,
	FileColumnIndividualSex:                             DBColumnIndividualSex,
	FileColumnIndividualSpokenLanguage1:                 DBColumnIndividualSpokenLanguage1,
	FileColumnIndividualSpokenLanguage2:                 DBColumnIndividualSpokenLanguage2,
	FileColumnIndividualSpokenLanguage3:                 DBColumnIndividualSpokenLanguage3,
	FileColumnIndividualVisionDisabilityLevel:           DBColumnIndividualVisionDisabilityLevel,
	FileColumnIndividualVulnerabilityComments:           DBColumnIndividualVulnerabilityComments,
	FileColumnIndividualCreatedAt:                       DBColumnIndividualCreatedAt,
	FileColumnIndividualUpdatedAt:                       DBColumnIndividualUpdatedAt,
}

var IndividualSystemFileColumns = []string{
	DBColumnIndividualCreatedAt,
	DBColumnIndividualUpdatedAt,
}
