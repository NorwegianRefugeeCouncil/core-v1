package constants

import "github.com/nrc-no/notcore/internal/containers"

const (
	FormParamIndividualInactive                        = "Inactive"
	FormParamIndividualAddress                         = "Address"
	FormParamIndividualAge                             = "Age"
	FormParamIndividualBirthDate                       = "BirthDate"
	FormParamIndividualCognitiveDisabilityLevel        = "CognitiveDisabilityLevel"
	FormParamIndividualCollectionAdministrativeArea1   = "CollectionAdministrativeArea1"
	FormParamIndividualCollectionAdministrativeArea2   = "CollectionAdministrativeArea2"
	FormParamIndividualCollectionAdministrativeArea3   = "CollectionAdministrativeArea3 "
	FormParamIndividualCollectionOffice                = "CollectionOffice "
	FormParamIndividualCollectionAgentName             = "CollectionAgentName "
	FormParamIndividualCollectionAgentTitle            = "CollectionAgentTitle "
	FormParamIndividualCollectionTime                  = "CollectionTime"
	FormParamIndividualComments                        = "Comments"
	FormParamIndividualCommunicationDisabilityLevel    = "CommunicationDisabilityLevel"
	FormParamIndividualCommunityID                     = "CommunityID"
	FormParamIndividualCommunitySize                   = "CommunitySize"
	FormParamIndividualCountryID                       = "Country"
	FormParamIndividualDisplacementStatus              = "DisplacementStatus"
	FormParamIndividualDisplacementStatusComment       = "DisplacementStatusComment"
	FormParamIndividualEmail1                          = "Email1"
	FormParamIndividualEmail2                          = "Email2"
	FormParamIndividualEmail3                          = "Email3"
	FormParamIndividualFreeField1                      = "FreeField1"
	FormParamIndividualFreeField2                      = "FreeField2"
	FormParamIndividualFreeField3                      = "FreeField3"
	FormParamIndividualFreeField4                      = "FreeField4"
	FormParamIndividualFreeField5                      = "FreeField5"
	FormParamIndividualFullName                        = "FullName"
	FormParamIndividualFirstName                       = "FirstName"
	FormParamIndividualMiddleName                      = "MiddleName"
	FormParamIndividualLastName                        = "LastName"
	FormParamIndividualNativeName                      = "NativeName"
	FormParamIndividualMothersName                     = "MothersName"
	FormParamIndividualSex                             = "Sex"
	FormParamIndividualHasCognitiveDisability          = "HasCognitiveDisability"
	FormParamIndividualHasCommunicationDisability      = "HasCommunicationDisability"
	FormParamIndividualHasConsentedToRGPD              = "HasConsentedToRGPD"
	FormParamIndividualHasConsentedToReferral          = "HasConsentedToReferral"
	FormParamIndividualHasDisability                   = "HasDisability"
	FormParamIndividualHasHearingDisability            = "HasHearingDisability"
	FormParamIndividualHasMobilityDisability           = "HasMobilityDisability"
	FormParamIndividualHasSelfCareDisability           = "HasSelfCareDisability"
	FormParamIndividualHasVisionDisability             = "HasVisionDisability"
	FormParamIndividualHearingDisabilityLevel          = "HearingDisabilityLevel"
	FormParamIndividualHouseholdID                     = "HouseholdID"
	FormParamIndividualHouseholdSize                   = "HouseholdSize"
	FormParamIndividualID                              = "ID"
	FormParamIndividualEngagementContext               = "EngagementContext"
	FormParamIndividualIdentificationNumber1           = "identificationNumber1"
	FormParamIndividualIdentificationNumber2           = "identificationNumber2"
	FormParamIndividualIdentificationNumber3           = "identificationNumber3"
	FormParamIndividualIdentificationType1             = "IdentificationType1"
	FormParamIndividualIdentificationType2             = "IdentificationType2"
	FormParamIndividualIdentificationType3             = "IdentificationType3"
	FormParamIndividualIdentificationTypeExplanation1  = "IdentificationTypeExplanation1"
	FormParamIndividualIdentificationTypeExplanation2  = "IdentificationTypeExplanation2"
	FormParamIndividualIdentificationTypeExplanation3  = "IdentificationTypeExplanation3"
	FormParamIndividualInternalID                      = "InternalID"
	FormParamIndividualIsHeadOfCommunity               = "IsHeadOfCommunity"
	FormParamIndividualIsHeadOfHousehold               = "IsHeadOfHousehold"
	FormParamIndividualIsFemaleHeadedHousehold         = "IsFemaleHeadedHousehold"
	FormParamIndividualIsMinorHeadedHousehold          = "IsMinorHeadedHousehold"
	FormParamIndividualIsMinor                         = "IsMinor"
	FormParamIndividualIsChildAtRisk                   = "IsChildAtRisk"
	FormParamIndividualIsElderAtRisk                   = "IsElderAtRisk"
	FormParamIndividualIsWomanAtRisk                   = "IsWomanAtRisk"
	FormParamIndividualIsPregnant                      = "IsPregnant"
	FormParamIndividualIsLactating                     = "IsLactating"
	FormParamIndividualIsSeparatedChild                = "IsSeparatedChild"
	FormParamIndividualIsSingleParent                  = "IsSingleParent"
	FormParamIndividualHasMedicalCondition             = "HasMedicalCondition"
	FormParamIndividualNeedsLegalAndPhysicalProtection = "NeedsLegalAndPhysicalProtection"
	FormParamIndividualMobilityDisabilityLevel         = "MobilityDisabilityLevel"
	FormParamIndividualNationality1                    = "Nationality1"
	FormParamIndividualNationality2                    = "Nationality2"
	FormParamIndividualPhoneNumber1                    = "PhoneNumber1"
	FormParamIndividualPhoneNumber2                    = "PhoneNumber2"
	FormParamIndividualPhoneNumber3                    = "PhoneNumber3"
	FormParamIndividualPreferredCommunicationLanguage  = "PreferredCommunicationLanguage"
	FormParamIndividualPreferredContactMethod          = "PreferredContactMethod"
	FormParamIndividualPreferredContactMethodComments  = "PreferredContactMethodComments"
	FormParamIndividualPreferredName                   = "PreferredName"
	FormParamIndividualPrefersToRemainAnonymous        = "PrefersToRemainAnonymous"
	FormParamIndividualPresentsProtectionConcerns      = "PresentsProtectionConcerns"
	FormParamIndividualPWDComments                     = "PWDComments"
	FormParamIndividualSelfCareDisabilityLevel         = "SelfCareDisabilityLevel"
	FormParamIndividualSpokenLanguage1                 = "SpokenLanguage1"
	FormParamIndividualSpokenLanguage2                 = "SpokenLanguage2"
	FormParamIndividualSpokenLanguage3                 = "SpokenLanguage3"
	FormParamIndividualVisionDisabilityLevel           = "VisionDisabilityLevel"
	FormParamIndividualServiceCC1                      = "ServiceCC1"
	FormParamIndividualServiceRequestedDate1           = "ServiceRequestedDate1"
	FormParamIndividualServiceDeliveredDate1           = "ServiceDeliveredDate1"
	FormParamIndividualServiceComments1                = "ServiceComments1"
	FormParamIndividualServiceCC2                      = "ServiceCC2"
	FormParamIndividualServiceRequestedDate2           = "ServiceRequestedDate2"
	FormParamIndividualServiceDeliveredDate2           = "ServiceDeliveredDate2"
	FormParamIndividualServiceComments2                = "ServiceComments2"
	FormParamIndividualServiceCC3                      = "ServiceCC3"
	FormParamIndividualServiceRequestedDate3           = "ServiceRequestedDate3"
	FormParamIndividualServiceDeliveredDate3           = "ServiceDeliveredDate3"
	FormParamIndividualServiceComments3                = "ServiceComments3"
	FormParamIndividualServiceCC4                      = "ServiceCC4"
	FormParamIndividualServiceRequestedDate4           = "ServiceRequestedDate4"
	FormParamIndividualServiceDeliveredDate4           = "ServiceDeliveredDate4"
	FormParamIndividualServiceComments4                = "ServiceComments4"
	FormParamIndividualServiceCC5                      = "ServiceCC5"
	FormParamIndividualServiceRequestedDate5           = "ServiceRequestedDate5"
	FormParamIndividualServiceDeliveredDate5           = "ServiceDeliveredDate5"
	FormParamIndividualServiceComments5                = "ServiceComments5"
	FormParamIndividualServiceCC6                      = "ServiceCC6"
	FormParamIndividualServiceRequestedDate6           = "ServiceRequestedDate6"
	FormParamIndividualServiceDeliveredDate6           = "ServiceDeliveredDate6"
	FormParamIndividualServiceComments6                = "ServiceComments6"
	FormParamIndividualServiceCC7                      = "ServiceCC7"
	FormParamIndividualServiceRequestedDate7           = "ServiceRequestedDate7"
	FormParamIndividualServiceDeliveredDate7           = "ServiceDeliveredDate7"
	FormParamIndividualServiceComments7                = "ServiceComments7"

	FormParamsGetIndividualCognitiveDisabilityLevel         = "cognitive_disability_level"
	FormParamsGetIndividualCollectionAdministrativeArea1    = "collection_administrative_area_1"
	FormParamsGetIndividualCollectionAdministrativeArea2    = "collection_administrative_area_2"
	FormParamsGetIndividualCollectionAdministrativeArea3    = "collection_administrative_area_3"
	FormParamsGetIndividualCollectionOffice                 = "collection_office"
	FormParamsGetIndividualCollectionAgentName              = "collection_agent_name"
	FormParamsGetIndividualCollectionAgentTitle             = "collection_agent_title"
	FormParamsGetIndividualsInactive                        = "inactive"
	FormParamsGetIndividualsAddress                         = "address"
	FormParamsGetIndividualsAgeFrom                         = "age_from"
	FormParamsGetIndividualsAgeTo                           = "age_to"
	FormParamsGetIndividualsBirthDateFrom                   = "birth_date_from"
	FormParamsGetIndividualsBirthDateTo                     = "birth_date_to"
	FormParamsGetIndividualsCognitiveDisabilityLevel        = "cognitive_disability_level"
	FormParamsGetIndividualsCollectionAdministrativeArea1   = "collection_administrative_area_1"
	FormParamsGetIndividualsCollectionAdministrativeArea2   = "collection_administrative_area_2"
	FormParamsGetIndividualsCollectionAdministrativeArea3   = "collection_administrative_area_3"
	FormParamsGetIndividualsCollectionOffice                = "collection_office"
	FormParamsGetIndividualsCollectionAgentName             = "collection_agent_name"
	FormParamsGetIndividualsCollectionAgentTitle            = "collection_agent_title"
	FormParamsGetIndividualsCollectionTimeFrom              = "collection_time_from"
	FormParamsGetIndividualsCollectionTimeTo                = "collection_time_to"
	FormParamsGetIndividualsCommunicationDisabilityLevel    = "communication_disability_level "
	FormParamsGetIndividualsCommunityID                     = "community_id"
	FormParamsGetIndividualsCommunitySize                   = "community_size"
	FormParamsGetIndividualsCountryID                       = "country_id"
	FormParamsGetIndividualsCreatedAtFrom                   = "created_at_from"
	FormParamsGetIndividualsCreatedAtTo                     = "created_at_to"
	FormParamsGetIndividualsDisplacementStatus              = "displacement_status"
	FormParamsGetIndividualsDisplacementStatusComment       = "displacement_status_comment"
	FormParamsGetIndividualsEmail                           = "email"
	FormParamsGetIndividualsFullName                        = "full_name"
	FormParamsGetIndividualsFirstName                       = "first_name"
	FormParamsGetIndividualsMiddleName                      = "middle_name"
	FormParamsGetIndividualsLastName                        = "last_name"
	FormParamsGetIndividualsNativeName                      = "native_name"
	FormParamsGetIndividualsMothersName                     = "mothers_name"
	FormParamsGetIndividualsFreeField1                      = "free_field_1"
	FormParamsGetIndividualsFreeField2                      = "free_field_2"
	FormParamsGetIndividualsFreeField3                      = "free_field_3"
	FormParamsGetIndividualsFreeField4                      = "free_field_4"
	FormParamsGetIndividualsFreeField5                      = "free_field_5"
	FormParamsGetIndividualsSex                             = "sex"
	FormParamsGetIndividualsHasCognitiveDisability          = "has_cognitive_disability"
	FormParamsGetIndividualsHasCommunicationDisability      = "has_communication_disability"
	FormParamsGetIndividualsHasConsentedToReferral          = "has_consented_to_referral"
	FormParamsGetIndividualsHasConsentedToRgpd              = "has_consented_to_rgpd"
	FormParamsGetIndividualsHasDisability                   = "has_disability"
	FormParamsGetIndividualsHasHearingDisability            = "has_hearing_disability"
	FormParamsGetIndividualsHasMobilityDisability           = "has_mobility_disability"
	FormParamsGetIndividualsHasSelfCareDisability           = "has_selfcare_disability"
	FormParamsGetIndividualsHasVisionDisability             = "has_vision_disability"
	FormParamsGetIndividualsHearingDisabilityLevel          = "hearing_disability_level"
	FormParamsGetIndividualsHouseholdID                     = "household_id"
	FormParamsGetIndividualsHouseholdSize                   = "household_size"
	FormParamsGetIndividualsID                              = "id"
	FormParamsGetIndividualsEngagementContext               = "engagement_context"
	FormParamsGetIndividualsIdentificationNumber            = "identification_number"
	FormParamsGetIndividualsInternalID                      = "internal_id"
	FormParamsGetIndividualsIsHeadOfCommunity               = "is_head_of_community"
	FormParamsGetIndividualsIsHeadOfHousehold               = "is_head_of_household"
	FormParamsGetIndividualsIsFemaleHeadedHousehold         = "is_female_headed_household"
	FormParamsGetIndividualsIsMinorHeadedHousehold          = "is_minor_headed_household"
	FormParamsGetIndividualsIsMinor                         = "is_minor"
	FormParamsGetIndividualsIsChildAtRisk                   = "is_child_at_risk"
	FormParamsGetIndividualsIsWomanAtRisk                   = "is_woman_at_risk"
	FormParamsGetIndividualsIsElderAtRisk                   = "is_elder_at_risk"
	FormParamsGetIndividualsIsPregnant                      = "is_pregnant"
	FormParamsGetIndividualsIsLactating                     = "is_lactating"
	FormParamsGetIndividualsIsSeparatedChild                = "is_separated_child"
	FormParamsGetIndividualsIsSingleParent                  = "is_single_parent"
	FormParamsGetIndividualsHasMedicalCondition             = "has_medical_condition"
	FormParamsGetIndividualsNeedsLegalAndPhysicalProtection = "needs_legal_and_physical_protection"
	FormParamsGetIndividualsMobilityDisabilityLevel         = "mobility_disability_level"
	FormParamsGetIndividualsNationality                     = "nationality"
	FormParamsGetIndividualsPhoneNumber                     = "phone_number"
	FormParamsGetIndividualsPreferredCommunicationLanguage  = "preferred_communication_language"
	FormParamsGetIndividualsPreferredContactMethod          = "preferred_contact_method"
	FormParamsGetIndividualsPrefersToRemainAnonymous        = "prefers_to_remain_anonymous"
	FormParamsGetIndividualsPresentsProtectionConcerns      = "presents_protection_concerns"
	FormParamsGetIndividualsPWDComments                     = "pwd_comments"
	FormParamsGetIndividualsSelfCareDisabilityLevel         = "selfcare_disability_level"
	FormParamsGetIndividualsSkip                            = "skip"
	FormParamsGetIndividualsSpokenLanguage                  = "spoken_language"
	FormParamsGetIndividualsTake                            = "take"
	FormParamsGetIndividualsUpdatedAtFrom                   = "updated_at_from"
	FormParamsGetIndividualsUpdatedAtTo                     = "updated_at_to"
	FormParamsGetIndividualsVisionDisabilityLevel           = "vision_disability_level"
	FormParamsGetIndividualsServiceCC                       = "service_cc"
	FormParamsGetIndividualsServiceRequestedDateFrom        = "service_requested_date_from"
	FormParamsGetIndividualsServiceRequestedDateTo          = "service_requested_date_to"
	FormParamsGetIndividualsServiceDeliveredDateFrom        = "service_delivered_date_from"
	FormParamsGetIndividualsServiceDeliveredDateTo          = "service_delivered_date_to"
	FormParamsGetIndividualsSort                            = "sort"

	DBColumnIndividualInactive                        = "inactive"
	DBColumnIndividualAddress                         = "address"
	DBColumnIndividualAge                             = "age"
	DBColumnIndividualBirthDate                       = "birth_date"
	DBColumnIndividualCognitiveDisabilityLevel        = "cognitive_disability_level"
	DBColumnIndividualCollectionAdministrativeArea1   = "collection_administrative_area_1"
	DBColumnIndividualCollectionAdministrativeArea2   = "collection_administrative_area_2"
	DBColumnIndividualCollectionAdministrativeArea3   = "collection_administrative_area_3"
	DBColumnIndividualCollectionOffice                = "collection_office"
	DBColumnIndividualCollectionAgentName             = "collection_agent_name"
	DBColumnIndividualCollectionAgentTitle            = "collection_agent_title"
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
	DBColumnIndividualFullName                        = "full_name"
	DBColumnIndividualFirstName                       = "first_name"
	DBColumnIndividualMiddleName                      = "middle_name"
	DBColumnIndividualLastName                        = "last_name"
	DBColumnIndividualNativeName                      = "native_name"
	DBColumnIndividualMothersName                     = "mothers_name"
	DBColumnIndividualFreeField1                      = "free_field_1"
	DBColumnIndividualFreeField2                      = "free_field_2"
	DBColumnIndividualFreeField3                      = "free_field_3"
	DBColumnIndividualFreeField4                      = "free_field_4"
	DBColumnIndividualFreeField5                      = "free_field_5"
	DBColumnIndividualSex                             = "sex"
	DBColumnIndividualHasCognitiveDisability          = "has_cognitive_disability"
	DBColumnIndividualHasCommunicationDisability      = "has_communication_disability"
	DBColumnIndividualHasConsentedToRGPD              = "has_consented_to_rgpd"
	DBColumnIndividualHasConsentedToReferral          = "has_consented_to_referral"
	DBColumnIndividualHasDisability                   = "has_disability"
	DBColumnIndividualHasHearingDisability            = "has_hearing_disability"
	DBColumnIndividualHasMobilityDisability           = "has_mobility_disability"
	DBColumnIndividualHasSelfCareDisability           = "has_selfcare_disability"
	DBColumnIndividualHasVisionDisability             = "has_vision_disability"
	DBColumnIndividualHearingDisabilityLevel          = "hearing_disability_level"
	DBColumnIndividualHouseholdID                     = "household_id"
	DBColumnIndividualHouseholdSize                   = "household_size"
	DBColumnIndividualID                              = "id"
	DBColumnIndividualEngagementContext               = "engagement_context"
	DBColumnIndividualIdentificationNumber1           = "identification_number_1"
	DBColumnIndividualIdentificationNumber2           = "identification_number_2"
	DBColumnIndividualIdentificationNumber3           = "identification_number_3"
	DBColumnIndividualIdentificationType1             = "identification_type_1"
	DBColumnIndividualIdentificationType2             = "identification_type_2"
	DBColumnIndividualIdentificationType3             = "identification_type_3"
	DBColumnIndividualIdentificationTypeExplanation1  = "identification_type_explanation_1"
	DBColumnIndividualIdentificationTypeExplanation2  = "identification_type_explanation_2"
	DBColumnIndividualIdentificationTypeExplanation3  = "identification_type_explanation_3"
	DBColumnIndividualInternalID                      = "internal_id"
	DBColumnIndividualIsHeadOfCommunity               = "is_head_of_community"
	DBColumnIndividualIsHeadOfHousehold               = "is_head_of_household"
	DBColumnIndividualIsFemaleHeadedHousehold         = "is_female_headed_household"
	DBColumnIndividualIsMinorHeadedHousehold          = "is_minor_headed_household"
	DBColumnIndividualIsMinor                         = "is_minor"
	DBColumnIndividualIsChildAtRisk                   = "is_child_at_risk"
	DBColumnIndividualIsElderAtRisk                   = "is_elder_at_risk"
	DBColumnIndividualIsSingleParent                  = "is_single_parent"
	DBColumnIndividualIsSeparatedChild                = "is_separated_child"
	DBColumnIndividualIsWomanAtRisk                   = "is_woman_at_risk"
	DBColumnIndividualIsPregnant                      = "is_pregnant"
	DBColumnIndividualIsLactating                     = "is_lactating"
	DBColumnIndividualHasMedicalCondition             = "has_medical_condition"
	DBColumnIndividualNeedsLegalAndPhysicalProtection = "needs_legal_and_physical_protection"
	DBColumnIndividualMobilityDisabilityLevel         = "mobility_disability_level"
	DBColumnIndividualNationality1                    = "nationality_1"
	DBColumnIndividualNationality2                    = "nationality_2"
	DBColumnIndividualNormalizedPhoneNumber1          = "normalized_phone_number_1"
	DBColumnIndividualNormalizedPhoneNumber2          = "normalized_phone_number_2"
	DBColumnIndividualNormalizedPhoneNumber3          = "normalized_phone_number_3"
	DBColumnIndividualPhoneNumber1                    = "phone_number_1"
	DBColumnIndividualPhoneNumber2                    = "phone_number_2"
	DBColumnIndividualPhoneNumber3                    = "phone_number_3"
	DBColumnIndividualPreferredCommunicationLanguage  = "preferred_communication_language"
	DBColumnIndividualPreferredContactMethod          = "preferred_contact_method"
	DBColumnIndividualPreferredContactMethodComments  = "preferred_contact_method_comments"
	DBColumnIndividualPreferredName                   = "preferred_name"
	DBColumnIndividualPrefersToRemainAnonymous        = "prefers_to_remain_anonymous"
	DBColumnIndividualPresentsProtectionConcerns      = "presents_protection_concerns"
	DBColumnIndividualPWDComments                     = "pwd_comments"
	DBColumnIndividualSelfCareDisabilityLevel         = "selfcare_disability_level"
	DBColumnIndividualSpokenLanguage1                 = "spoken_language_1"
	DBColumnIndividualSpokenLanguage2                 = "spoken_language_2"
	DBColumnIndividualSpokenLanguage3                 = "spoken_language_3"
	DBColumnIndividualUpdatedAt                       = "updated_at"
	DBColumnIndividualVisionDisabilityLevel           = "vision_disability_level"
	DBColumnIndividualServiceCC1                      = "service_cc_1"
	DBColumnIndividualServiceRequestedDate1           = "service_requested_date_1"
	DBColumnIndividualServiceDeliveredDate1           = "service_delivered_date_1"
	DBColumnIndividualServiceComments1                = "service_comments_1"
	DBColumnIndividualServiceCC2                      = "service_cc_2"
	DBColumnIndividualServiceRequestedDate2           = "service_requested_date_2"
	DBColumnIndividualServiceDeliveredDate2           = "service_delivered_date_2"
	DBColumnIndividualServiceComments2                = "service_comments_2"
	DBColumnIndividualServiceCC3                      = "service_cc_3"
	DBColumnIndividualServiceRequestedDate3           = "service_requested_date_3"
	DBColumnIndividualServiceDeliveredDate3           = "service_delivered_date_3"
	DBColumnIndividualServiceComments3                = "service_comments_3"
	DBColumnIndividualServiceCC4                      = "service_cc_4"
	DBColumnIndividualServiceRequestedDate4           = "service_requested_date_4"
	DBColumnIndividualServiceDeliveredDate4           = "service_delivered_date_4"
	DBColumnIndividualServiceComments4                = "service_comments_4"
	DBColumnIndividualServiceCC5                      = "service_cc_5"
	DBColumnIndividualServiceRequestedDate5           = "service_requested_date_5"
	DBColumnIndividualServiceDeliveredDate5           = "service_delivered_date_5"
	DBColumnIndividualServiceComments5                = "service_comments_5"
	DBColumnIndividualServiceCC6                      = "service_cc_6"
	DBColumnIndividualServiceRequestedDate6           = "service_requested_date_6"
	DBColumnIndividualServiceDeliveredDate6           = "service_delivered_date_6"
	DBColumnIndividualServiceComments6                = "service_comments_6"
	DBColumnIndividualServiceCC7                      = "service_cc_7"
	DBColumnIndividualServiceRequestedDate7           = "service_requested_date_7"
	DBColumnIndividualServiceDeliveredDate7           = "service_delivered_date_7"
	DBColumnIndividualServiceComments7                = "service_comments_7"

	FileColumnIndividualInactive                        = "inactive"
	FileColumnIndividualAddress                         = "address"
	FileColumnIndividualAge                             = "age"
	FileColumnIndividualBirthDate                       = "birth_date"
	FileColumnIndividualCognitiveDisabilityLevel        = "cognitive_disability_level"
	FileColumnIndividualCollectionAdministrativeArea1   = "collection_administrative_area_1"
	FileColumnIndividualCollectionAdministrativeArea2   = "collection_administrative_area_2"
	FileColumnIndividualCollectionAdministrativeArea3   = "collection_administrative_area_3"
	FileColumnIndividualCollectionOffice                = "collection_office"
	FileColumnIndividualCollectionAgentName             = "collection_agent_name"
	FileColumnIndividualCollectionAgentTitle            = "collection_agent_title"
	FileColumnIndividualCollectionTime                  = "collection_time"
	FileColumnIndividualComments                        = "comments"
	FileColumnIndividualCommunicationDisabilityLevel    = "communication_disability_level"
	FileColumnIndividualCommunityID                     = "community_id"
	FileColumnIndividualCommunitySize                   = "community_size"
	FileColumnIndividualCountryID                       = "country_id"
	FileColumnIndividualCreatedAt                       = "created_at"
	FileColumnIndividualDisplacementStatus              = "displacement_status"
	FileColumnIndividualDisplacementStatusComment       = "displacement_status_comment"
	FileColumnIndividualEmail1                          = "email_1"
	FileColumnIndividualEmail2                          = "email_2"
	FileColumnIndividualEmail3                          = "email_3"
	FileColumnIndividualFullName                        = "full_name"
	FileColumnIndividualFirstName                       = "first_name"
	FileColumnIndividualMiddleName                      = "middle_name"
	FileColumnIndividualLastName                        = "last_name"
	FileColumnIndividualNativeName                      = "native_name"
	FileColumnIndividualMothersName                     = "mothers_name"
	FileColumnIndividualFreeField1                      = "free_field_1"
	FileColumnIndividualFreeField2                      = "free_field_2"
	FileColumnIndividualFreeField3                      = "free_field_3"
	FileColumnIndividualFreeField4                      = "free_field_4"
	FileColumnIndividualFreeField5                      = "free_field_5"
	FileColumnIndividualSex                             = "sex"
	FileColumnIndividualHasMedicalCondition             = "has_medical_condition"
	FileColumnIndividualNeedsLegalAndPhysicalProtection = "needs_legal_and_physical_protection"
	FileColumnIndividualIsChildAtRisk                   = "is_child_at_risk"
	FileColumnIndividualIsWomanAtRisk                   = "is_woman_at_risk"
	FileColumnIndividualIsElderAtRisk                   = "is_elder_at_risk"
	FileColumnIndividualIsPregnant                      = "is_pregnant"
	FileColumnIndividualIsLactating                     = "is_lactating"
	FileColumnIndividualIsSingleParent                  = "is_single_parent"
	FileColumnIndividualIsSeparatedChild                = "is_separated_child"
	FileColumnIndividualHasCognitiveDisability          = "has_cognitive_disability"
	FileColumnIndividualHasCommunicationDisability      = "has_communication_disability"
	FileColumnIndividualHasConsentedToRGPD              = "has_consented_to_rgpd"
	FileColumnIndividualHasConsentedToReferral          = "has_consented_to_referral"
	FileColumnIndividualHasHearingDisability            = "has_hearing_disability"
	FileColumnIndividualHasDisability                   = "has_disability"
	FileColumnIndividualHasMobilityDisability           = "has_mobility_disability"
	FileColumnIndividualHasSelfCareDisability           = "has_selfcare_disability"
	FileColumnIndividualHasVisionDisability             = "has_vision_disability"
	FileColumnIndividualHearingDisabilityLevel          = "hearing_disability_level"
	FileColumnIndividualHouseholdID                     = "household_id"
	FileColumnIndividualHouseholdSize                   = "household_size"
	FileColumnIndividualID                              = "id"
	FileColumnIndividualEngagementContext               = "engagement_context"
	FileColumnIndividualIdentificationNumber1           = "identification_number_1"
	FileColumnIndividualIdentificationNumber2           = "identification_number_2"
	FileColumnIndividualIdentificationNumber3           = "identification_number_3"
	FileColumnIndividualIdentificationType1             = "identification_type_1"
	FileColumnIndividualIdentificationType2             = "identification_type_2"
	FileColumnIndividualIdentificationType3             = "identification_type_3"
	FileColumnIndividualIdentificationTypeExplanation1  = "identification_type_explanation_1"
	FileColumnIndividualIdentificationTypeExplanation2  = "identification_type_explanation_2"
	FileColumnIndividualIdentificationTypeExplanation3  = "identification_type_explanation_3"
	FileColumnIndividualInternalID                      = "internal_id"
	FileColumnIndividualIsHeadOfCommunity               = "is_head_of_community"
	FileColumnIndividualIsHeadOfHousehold               = "is_head_of_household"
	FileColumnIndividualIsFemaleHeadedHousehold         = "is_female_headed_household"
	FileColumnIndividualIsMinorHeadedHousehold          = "is_minor_headed_household"
	FileColumnIndividualIsMinor                         = "is_minor"
	FileColumnIndividualMobilityDisabilityLevel         = "mobility_disability_level"
	FileColumnIndividualNationality1                    = "nationality_1"
	FileColumnIndividualNationality2                    = "nationality_2"
	FileColumnIndividualPhoneNumber1                    = "phone_number_1"
	FileColumnIndividualPhoneNumber2                    = "phone_number_2"
	FileColumnIndividualPhoneNumber3                    = "phone_number_3"
	FileColumnIndividualPreferredCommunicationLanguage  = "preferred_communication_language"
	FileColumnIndividualPreferredContactMethod          = "preferred_contact_method"
	FileColumnIndividualPreferredContactMethodComments  = "preferred_contact_method_comments"
	FileColumnIndividualPreferredName                   = "preferred_name"
	FileColumnIndividualPrefersToRemainAnonymous        = "prefers_to_remain_anonymous"
	FileColumnIndividualPresentsProtectionConcerns      = "presents_protection_concerns"
	FileColumnIndividualPWDComments                     = "pwd_comments"
	FileColumnIndividualSelfCareDisabilityLevel         = "selfcare_disability_level"
	FileColumnIndividualSpokenLanguage1                 = "spoken_language_1"
	FileColumnIndividualSpokenLanguage2                 = "spoken_language_2"
	FileColumnIndividualSpokenLanguage3                 = "spoken_language_3"
	FileColumnIndividualUpdatedAt                       = "updated_at"
	FileColumnIndividualVisionDisabilityLevel           = "vision_disability_level"
	FileColumnIndividualServiceCC1                      = "service_cc_1"
	FileColumnIndividualServiceRequestedDate1           = "service_requested_date_1"
	FileColumnIndividualServiceDeliveredDate1           = "service_delivered_date_1"
	FileColumnIndividualServiceComments1                = "service_comments_1"
	FileColumnIndividualServiceCC2                      = "service_cc_2"
	FileColumnIndividualServiceRequestedDate2           = "service_requested_date_2"
	FileColumnIndividualServiceDeliveredDate2           = "service_delivered_date_2"
	FileColumnIndividualServiceComments2                = "service_comments_2"
	FileColumnIndividualServiceCC3                      = "service_cc_3"
	FileColumnIndividualServiceRequestedDate3           = "service_requested_date_3"
	FileColumnIndividualServiceDeliveredDate3           = "service_delivered_date_3"
	FileColumnIndividualServiceComments3                = "service_comments_3"
	FileColumnIndividualServiceCC4                      = "service_cc_4"
	FileColumnIndividualServiceRequestedDate4           = "service_requested_date_4"
	FileColumnIndividualServiceDeliveredDate4           = "service_delivered_date_4"
	FileColumnIndividualServiceComments4                = "service_comments_4"
	FileColumnIndividualServiceCC5                      = "service_cc_5"
	FileColumnIndividualServiceRequestedDate5           = "service_requested_date_5"
	FileColumnIndividualServiceDeliveredDate5           = "service_delivered_date_5"
	FileColumnIndividualServiceComments5                = "service_comments_5"
	FileColumnIndividualServiceCC6                      = "service_cc_6"
	FileColumnIndividualServiceRequestedDate6           = "service_requested_date_6"
	FileColumnIndividualServiceDeliveredDate6           = "service_delivered_date_6"
	FileColumnIndividualServiceComments6                = "service_comments_6"
	FileColumnIndividualServiceCC7                      = "service_cc_7"
	FileColumnIndividualServiceRequestedDate7           = "service_requested_date_7"
	FileColumnIndividualServiceDeliveredDate7           = "service_delivered_date_7"
	FileColumnIndividualServiceComments7                = "service_comments_7"
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
	DBColumnIndividualCommunitySize,
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
	DBColumnIndividualNativeName,
	DBColumnIndividualMothersName,
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
	DBColumnIndividualHasDisability,
	DBColumnIndividualHasMobilityDisability,
	DBColumnIndividualHasSelfCareDisability,
	DBColumnIndividualHasVisionDisability,
	DBColumnIndividualHearingDisabilityLevel,
	DBColumnIndividualHouseholdID,
	DBColumnIndividualHouseholdSize,
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
	DBColumnIndividualIsChildAtRisk,
	DBColumnIndividualIsWomanAtRisk,
	DBColumnIndividualIsElderAtRisk,
	DBColumnIndividualIsPregnant,
	DBColumnIndividualIsLactating,
	DBColumnIndividualIsSeparatedChild,
	DBColumnIndividualIsSingleParent,
	DBColumnIndividualHasMedicalCondition,
	DBColumnIndividualNeedsLegalAndPhysicalProtection,
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
	DBColumnIndividualPWDComments,
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
	FileColumnIndividualNativeName,
	FileColumnIndividualPrefersToRemainAnonymous,
	FileColumnIndividualSex,
	FileColumnIndividualBirthDate,
	FileColumnIndividualAge,
	FileColumnIndividualIsMinor,
	FileColumnIndividualIsChildAtRisk,
	FileColumnIndividualIsWomanAtRisk,
	FileColumnIndividualIsElderAtRisk,
	FileColumnIndividualIsSeparatedChild,
	FileColumnIndividualIsSingleParent,
	FileColumnIndividualIsPregnant,
	FileColumnIndividualIsLactating,
	FileColumnIndividualHasMedicalCondition,
	FileColumnIndividualNeedsLegalAndPhysicalProtection,
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
	DBColumnIndividualInactive:                        FileColumnIndividualInactive,
	DBColumnIndividualAddress:                         FileColumnIndividualAddress,
	DBColumnIndividualAge:                             FileColumnIndividualAge,
	DBColumnIndividualBirthDate:                       FileColumnIndividualBirthDate,
	DBColumnIndividualCognitiveDisabilityLevel:        FileColumnIndividualCognitiveDisabilityLevel,
	DBColumnIndividualCollectionAdministrativeArea1:   FileColumnIndividualCollectionAdministrativeArea1,
	DBColumnIndividualCollectionAdministrativeArea2:   FileColumnIndividualCollectionAdministrativeArea2,
	DBColumnIndividualCollectionAdministrativeArea3:   FileColumnIndividualCollectionAdministrativeArea3,
	DBColumnIndividualCollectionOffice:                FileColumnIndividualCollectionOffice,
	DBColumnIndividualCollectionAgentName:             FileColumnIndividualCollectionAgentName,
	DBColumnIndividualCollectionAgentTitle:            FileColumnIndividualCollectionAgentTitle,
	DBColumnIndividualCollectionTime:                  FileColumnIndividualCollectionTime,
	DBColumnIndividualComments:                        FileColumnIndividualComments,
	DBColumnIndividualCommunicationDisabilityLevel:    FileColumnIndividualCommunicationDisabilityLevel,
	DBColumnIndividualCommunityID:                     FileColumnIndividualCommunityID,
	DBColumnIndividualCommunitySize:                   FileColumnIndividualCommunitySize,
	DBColumnIndividualCountryID:                       FileColumnIndividualCountryID,
	DBColumnIndividualCreatedAt:                       FileColumnIndividualCreatedAt,
	DBColumnIndividualDisplacementStatus:              FileColumnIndividualDisplacementStatus,
	DBColumnIndividualDisplacementStatusComment:       FileColumnIndividualDisplacementStatusComment,
	DBColumnIndividualEmail1:                          FileColumnIndividualEmail1,
	DBColumnIndividualEmail2:                          FileColumnIndividualEmail2,
	DBColumnIndividualEmail3:                          FileColumnIndividualEmail3,
	DBColumnIndividualFullName:                        FileColumnIndividualFullName,
	DBColumnIndividualFirstName:                       FileColumnIndividualFirstName,
	DBColumnIndividualMiddleName:                      FileColumnIndividualMiddleName,
	DBColumnIndividualLastName:                        FileColumnIndividualLastName,
	DBColumnIndividualNativeName:                      FileColumnIndividualNativeName,
	DBColumnIndividualMothersName:                     FileColumnIndividualMothersName,
	DBColumnIndividualFreeField1:                      FileColumnIndividualFreeField1,
	DBColumnIndividualFreeField2:                      FileColumnIndividualFreeField2,
	DBColumnIndividualFreeField3:                      FileColumnIndividualFreeField3,
	DBColumnIndividualFreeField4:                      FileColumnIndividualFreeField4,
	DBColumnIndividualFreeField5:                      FileColumnIndividualFreeField5,
	DBColumnIndividualSex:                             FileColumnIndividualSex,
	DBColumnIndividualHasCognitiveDisability:          FileColumnIndividualHasCognitiveDisability,
	DBColumnIndividualHasCommunicationDisability:      FileColumnIndividualHasCommunicationDisability,
	DBColumnIndividualHasConsentedToRGPD:              FileColumnIndividualHasConsentedToRGPD,
	DBColumnIndividualHasConsentedToReferral:          FileColumnIndividualHasConsentedToReferral,
	DBColumnIndividualHasHearingDisability:            FileColumnIndividualHasHearingDisability,
	DBColumnIndividualHasDisability:                   FileColumnIndividualHasDisability,
	DBColumnIndividualHasMobilityDisability:           FileColumnIndividualHasMobilityDisability,
	DBColumnIndividualHasSelfCareDisability:           FileColumnIndividualHasSelfCareDisability,
	DBColumnIndividualHasVisionDisability:             FileColumnIndividualHasVisionDisability,
	DBColumnIndividualHearingDisabilityLevel:          FileColumnIndividualHearingDisabilityLevel,
	DBColumnIndividualHouseholdID:                     FileColumnIndividualHouseholdID,
	DBColumnIndividualHouseholdSize:                   FileColumnIndividualHouseholdSize,
	DBColumnIndividualID:                              FileColumnIndividualID,
	DBColumnIndividualEngagementContext:               FileColumnIndividualEngagementContext,
	DBColumnIndividualIdentificationNumber1:           FileColumnIndividualIdentificationNumber1,
	DBColumnIndividualIdentificationNumber2:           FileColumnIndividualIdentificationNumber2,
	DBColumnIndividualIdentificationNumber3:           FileColumnIndividualIdentificationNumber3,
	DBColumnIndividualIdentificationType1:             FileColumnIndividualIdentificationType1,
	DBColumnIndividualIdentificationType2:             FileColumnIndividualIdentificationType2,
	DBColumnIndividualIdentificationType3:             FileColumnIndividualIdentificationType3,
	DBColumnIndividualIdentificationTypeExplanation1:  FileColumnIndividualIdentificationTypeExplanation1,
	DBColumnIndividualIdentificationTypeExplanation2:  FileColumnIndividualIdentificationTypeExplanation2,
	DBColumnIndividualIdentificationTypeExplanation3:  FileColumnIndividualIdentificationTypeExplanation3,
	DBColumnIndividualInternalID:                      FileColumnIndividualInternalID,
	DBColumnIndividualIsHeadOfCommunity:               FileColumnIndividualIsHeadOfCommunity,
	DBColumnIndividualIsHeadOfHousehold:               FileColumnIndividualIsHeadOfHousehold,
	DBColumnIndividualIsFemaleHeadedHousehold:         FileColumnIndividualIsFemaleHeadedHousehold,
	DBColumnIndividualIsMinorHeadedHousehold:          FileColumnIndividualIsMinorHeadedHousehold,
	DBColumnIndividualIsMinor:                         FileColumnIndividualIsMinor,
	DBColumnIndividualIsPregnant:                      FileColumnIndividualIsPregnant,
	DBColumnIndividualIsLactating:                     FileColumnIndividualIsLactating,
	DBColumnIndividualIsSingleParent:                  FileColumnIndividualIsSingleParent,
	DBColumnIndividualIsSeparatedChild:                FileColumnIndividualIsSeparatedChild,
	DBColumnIndividualIsChildAtRisk:                   FileColumnIndividualIsChildAtRisk,
	DBColumnIndividualIsWomanAtRisk:                   FileColumnIndividualIsWomanAtRisk,
	DBColumnIndividualIsElderAtRisk:                   FileColumnIndividualIsElderAtRisk,
	DBColumnIndividualNeedsLegalAndPhysicalProtection: FileColumnIndividualNeedsLegalAndPhysicalProtection,
	DBColumnIndividualHasMedicalCondition:             FileColumnIndividualHasMedicalCondition,
	DBColumnIndividualMobilityDisabilityLevel:         FileColumnIndividualMobilityDisabilityLevel,
	DBColumnIndividualNationality1:                    FileColumnIndividualNationality1,
	DBColumnIndividualNationality2:                    FileColumnIndividualNationality2,
	DBColumnIndividualPhoneNumber1:                    FileColumnIndividualPhoneNumber1,
	DBColumnIndividualPhoneNumber2:                    FileColumnIndividualPhoneNumber2,
	DBColumnIndividualPhoneNumber3:                    FileColumnIndividualPhoneNumber3,
	DBColumnIndividualPreferredCommunicationLanguage:  FileColumnIndividualPreferredCommunicationLanguage,
	DBColumnIndividualPreferredContactMethod:          FileColumnIndividualPreferredContactMethod,
	DBColumnIndividualPreferredContactMethodComments:  FileColumnIndividualPreferredContactMethodComments,
	DBColumnIndividualPreferredName:                   FileColumnIndividualPreferredName,
	DBColumnIndividualPrefersToRemainAnonymous:        FileColumnIndividualPrefersToRemainAnonymous,
	DBColumnIndividualPresentsProtectionConcerns:      FileColumnIndividualPresentsProtectionConcerns,
	DBColumnIndividualPWDComments:                     FileColumnIndividualPWDComments,
	DBColumnIndividualSelfCareDisabilityLevel:         FileColumnIndividualSelfCareDisabilityLevel,
	DBColumnIndividualSpokenLanguage1:                 FileColumnIndividualSpokenLanguage1,
	DBColumnIndividualSpokenLanguage2:                 FileColumnIndividualSpokenLanguage2,
	DBColumnIndividualSpokenLanguage3:                 FileColumnIndividualSpokenLanguage3,
	DBColumnIndividualVisionDisabilityLevel:           FileColumnIndividualVisionDisabilityLevel,
	DBColumnIndividualUpdatedAt:                       FileColumnIndividualUpdatedAt,
	DBColumnIndividualServiceCC1:                      FileColumnIndividualServiceCC1,
	DBColumnIndividualServiceRequestedDate1:           FileColumnIndividualServiceRequestedDate1,
	DBColumnIndividualServiceDeliveredDate1:           FileColumnIndividualServiceDeliveredDate1,
	DBColumnIndividualServiceComments1:                FileColumnIndividualServiceComments1,
	DBColumnIndividualServiceCC2:                      FileColumnIndividualServiceCC2,
	DBColumnIndividualServiceRequestedDate2:           FileColumnIndividualServiceRequestedDate2,
	DBColumnIndividualServiceDeliveredDate2:           FileColumnIndividualServiceDeliveredDate2,
	DBColumnIndividualServiceComments2:                FileColumnIndividualServiceComments2,
	DBColumnIndividualServiceCC3:                      FileColumnIndividualServiceCC3,
	DBColumnIndividualServiceRequestedDate3:           FileColumnIndividualServiceRequestedDate3,
	DBColumnIndividualServiceDeliveredDate3:           FileColumnIndividualServiceDeliveredDate3,
	DBColumnIndividualServiceComments3:                FileColumnIndividualServiceComments3,
	DBColumnIndividualServiceCC4:                      FileColumnIndividualServiceCC4,
	DBColumnIndividualServiceRequestedDate4:           FileColumnIndividualServiceRequestedDate4,
	DBColumnIndividualServiceDeliveredDate4:           FileColumnIndividualServiceDeliveredDate4,
	DBColumnIndividualServiceComments4:                FileColumnIndividualServiceComments4,
	DBColumnIndividualServiceCC5:                      FileColumnIndividualServiceCC5,
	DBColumnIndividualServiceRequestedDate5:           FileColumnIndividualServiceRequestedDate5,
	DBColumnIndividualServiceDeliveredDate5:           FileColumnIndividualServiceDeliveredDate5,
	DBColumnIndividualServiceComments5:                FileColumnIndividualServiceComments5,
	DBColumnIndividualServiceCC6:                      FileColumnIndividualServiceCC6,
	DBColumnIndividualServiceRequestedDate6:           FileColumnIndividualServiceRequestedDate6,
	DBColumnIndividualServiceDeliveredDate6:           FileColumnIndividualServiceDeliveredDate6,
	DBColumnIndividualServiceComments6:                FileColumnIndividualServiceComments6,
	DBColumnIndividualServiceCC7:                      FileColumnIndividualServiceCC7,
	DBColumnIndividualServiceRequestedDate7:           FileColumnIndividualServiceRequestedDate7,
	DBColumnIndividualServiceDeliveredDate7:           FileColumnIndividualServiceDeliveredDate7,
	DBColumnIndividualServiceComments7:                FileColumnIndividualServiceComments7,
}

var IndividualFileToDBMap = map[string]string{
	FileColumnIndividualInactive:                        DBColumnIndividualInactive,
	FileColumnIndividualAddress:                         DBColumnIndividualAddress,
	FileColumnIndividualAge:                             DBColumnIndividualAge,
	FileColumnIndividualBirthDate:                       DBColumnIndividualBirthDate,
	FileColumnIndividualCognitiveDisabilityLevel:        DBColumnIndividualCognitiveDisabilityLevel,
	FileColumnIndividualCollectionAdministrativeArea1:   DBColumnIndividualCollectionAdministrativeArea1,
	FileColumnIndividualCollectionAdministrativeArea2:   DBColumnIndividualCollectionAdministrativeArea2,
	FileColumnIndividualCollectionAdministrativeArea3:   DBColumnIndividualCollectionAdministrativeArea3,
	FileColumnIndividualCollectionOffice:                DBColumnIndividualCollectionOffice,
	FileColumnIndividualCollectionAgentName:             DBColumnIndividualCollectionAgentName,
	FileColumnIndividualCollectionAgentTitle:            DBColumnIndividualCollectionAgentTitle,
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
	FileColumnIndividualFullName:                        DBColumnIndividualFullName,
	FileColumnIndividualFirstName:                       DBColumnIndividualFirstName,
	FileColumnIndividualMiddleName:                      DBColumnIndividualMiddleName,
	FileColumnIndividualLastName:                        DBColumnIndividualLastName,
	FileColumnIndividualNativeName:                      DBColumnIndividualNativeName,
	FileColumnIndividualMothersName:                     DBColumnIndividualMothersName,
	FileColumnIndividualFreeField1:                      DBColumnIndividualFreeField1,
	FileColumnIndividualFreeField2:                      DBColumnIndividualFreeField2,
	FileColumnIndividualFreeField3:                      DBColumnIndividualFreeField3,
	FileColumnIndividualFreeField4:                      DBColumnIndividualFreeField4,
	FileColumnIndividualFreeField5:                      DBColumnIndividualFreeField5,
	FileColumnIndividualSex:                             DBColumnIndividualSex,
	FileColumnIndividualHasCognitiveDisability:          DBColumnIndividualHasCognitiveDisability,
	FileColumnIndividualHasCommunicationDisability:      DBColumnIndividualHasCommunicationDisability,
	FileColumnIndividualHasConsentedToRGPD:              DBColumnIndividualHasConsentedToRGPD,
	FileColumnIndividualHasConsentedToReferral:          DBColumnIndividualHasConsentedToReferral,
	FileColumnIndividualHasDisability:                   DBColumnIndividualHasDisability,
	FileColumnIndividualHasHearingDisability:            DBColumnIndividualHasHearingDisability,
	FileColumnIndividualHasMobilityDisability:           DBColumnIndividualHasMobilityDisability,
	FileColumnIndividualHasSelfCareDisability:           DBColumnIndividualHasSelfCareDisability,
	FileColumnIndividualHasVisionDisability:             DBColumnIndividualHasVisionDisability,
	FileColumnIndividualHearingDisabilityLevel:          DBColumnIndividualHearingDisabilityLevel,
	FileColumnIndividualHouseholdID:                     DBColumnIndividualHouseholdID,
	FileColumnIndividualHouseholdSize:                   DBColumnIndividualHouseholdSize,
	FileColumnIndividualID:                              DBColumnIndividualID,
	FileColumnIndividualEngagementContext:               DBColumnIndividualEngagementContext,
	FileColumnIndividualIdentificationNumber1:           DBColumnIndividualIdentificationNumber1,
	FileColumnIndividualIdentificationNumber2:           DBColumnIndividualIdentificationNumber2,
	FileColumnIndividualIdentificationNumber3:           DBColumnIndividualIdentificationNumber3,
	FileColumnIndividualIdentificationType1:             DBColumnIndividualIdentificationType1,
	FileColumnIndividualIdentificationType2:             DBColumnIndividualIdentificationType2,
	FileColumnIndividualIdentificationType3:             DBColumnIndividualIdentificationType3,
	FileColumnIndividualIdentificationTypeExplanation1:  DBColumnIndividualIdentificationTypeExplanation1,
	FileColumnIndividualIdentificationTypeExplanation2:  DBColumnIndividualIdentificationTypeExplanation2,
	FileColumnIndividualIdentificationTypeExplanation3:  DBColumnIndividualIdentificationTypeExplanation3,
	FileColumnIndividualInternalID:                      DBColumnIndividualInternalID,
	FileColumnIndividualIsHeadOfCommunity:               DBColumnIndividualIsHeadOfCommunity,
	FileColumnIndividualIsHeadOfHousehold:               DBColumnIndividualIsHeadOfHousehold,
	FileColumnIndividualIsFemaleHeadedHousehold:         DBColumnIndividualIsFemaleHeadedHousehold,
	FileColumnIndividualIsMinorHeadedHousehold:          DBColumnIndividualIsMinorHeadedHousehold,
	FileColumnIndividualIsMinor:                         DBColumnIndividualIsMinor,
	FileColumnIndividualIsChildAtRisk:                   DBColumnIndividualIsChildAtRisk,
	FileColumnIndividualIsWomanAtRisk:                   DBColumnIndividualIsWomanAtRisk,
	FileColumnIndividualIsElderAtRisk:                   DBColumnIndividualIsElderAtRisk,
	FileColumnIndividualIsSeparatedChild:                DBColumnIndividualIsSeparatedChild,
	FileColumnIndividualIsSingleParent:                  DBColumnIndividualIsSingleParent,
	FileColumnIndividualIsPregnant:                      DBColumnIndividualIsPregnant,
	FileColumnIndividualIsLactating:                     DBColumnIndividualIsLactating,
	FileColumnIndividualHasMedicalCondition:             DBColumnIndividualHasMedicalCondition,
	FileColumnIndividualNeedsLegalAndPhysicalProtection: DBColumnIndividualNeedsLegalAndPhysicalProtection,
	FileColumnIndividualMobilityDisabilityLevel:         DBColumnIndividualMobilityDisabilityLevel,
	FileColumnIndividualNationality1:                    DBColumnIndividualNationality1,
	FileColumnIndividualNationality2:                    DBColumnIndividualNationality2,
	FileColumnIndividualPhoneNumber1:                    DBColumnIndividualPhoneNumber1,
	FileColumnIndividualPhoneNumber2:                    DBColumnIndividualPhoneNumber2,
	FileColumnIndividualPhoneNumber3:                    DBColumnIndividualPhoneNumber3,
	FileColumnIndividualPreferredCommunicationLanguage:  DBColumnIndividualPreferredCommunicationLanguage,
	FileColumnIndividualPreferredContactMethod:          DBColumnIndividualPreferredContactMethod,
	FileColumnIndividualPreferredContactMethodComments:  DBColumnIndividualPreferredContactMethodComments,
	FileColumnIndividualPreferredName:                   DBColumnIndividualPreferredName,
	FileColumnIndividualPrefersToRemainAnonymous:        DBColumnIndividualPrefersToRemainAnonymous,
	FileColumnIndividualPresentsProtectionConcerns:      DBColumnIndividualPresentsProtectionConcerns,
	FileColumnIndividualPWDComments:                     DBColumnIndividualPWDComments,
	FileColumnIndividualSelfCareDisabilityLevel:         DBColumnIndividualSelfCareDisabilityLevel,
	FileColumnIndividualSpokenLanguage1:                 DBColumnIndividualSpokenLanguage1,
	FileColumnIndividualSpokenLanguage2:                 DBColumnIndividualSpokenLanguage2,
	FileColumnIndividualSpokenLanguage3:                 DBColumnIndividualSpokenLanguage3,
	FileColumnIndividualVisionDisabilityLevel:           DBColumnIndividualVisionDisabilityLevel,
	FileColumnIndividualServiceCC1:                      DBColumnIndividualServiceCC1,
	FileColumnIndividualServiceRequestedDate1:           DBColumnIndividualServiceRequestedDate1,
	FileColumnIndividualServiceDeliveredDate1:           DBColumnIndividualServiceDeliveredDate1,
	FileColumnIndividualServiceComments1:                DBColumnIndividualServiceComments1,
	FileColumnIndividualServiceCC2:                      DBColumnIndividualServiceCC2,
	FileColumnIndividualServiceRequestedDate2:           DBColumnIndividualServiceRequestedDate2,
	FileColumnIndividualServiceDeliveredDate2:           DBColumnIndividualServiceDeliveredDate2,
	FileColumnIndividualServiceComments2:                DBColumnIndividualServiceComments2,
	FileColumnIndividualServiceCC3:                      DBColumnIndividualServiceCC3,
	FileColumnIndividualServiceRequestedDate3:           DBColumnIndividualServiceRequestedDate3,
	FileColumnIndividualServiceDeliveredDate3:           DBColumnIndividualServiceDeliveredDate3,
	FileColumnIndividualServiceComments3:                DBColumnIndividualServiceComments3,
	FileColumnIndividualServiceCC4:                      DBColumnIndividualServiceCC4,
	FileColumnIndividualServiceRequestedDate4:           DBColumnIndividualServiceRequestedDate4,
	FileColumnIndividualServiceDeliveredDate4:           DBColumnIndividualServiceDeliveredDate4,
	FileColumnIndividualServiceComments4:                DBColumnIndividualServiceComments4,
	FileColumnIndividualServiceCC5:                      DBColumnIndividualServiceCC5,
	FileColumnIndividualServiceRequestedDate5:           DBColumnIndividualServiceRequestedDate5,
	FileColumnIndividualServiceDeliveredDate5:           DBColumnIndividualServiceDeliveredDate5,
	FileColumnIndividualServiceComments5:                DBColumnIndividualServiceComments5,
	FileColumnIndividualServiceCC6:                      DBColumnIndividualServiceCC6,
	FileColumnIndividualServiceRequestedDate6:           DBColumnIndividualServiceRequestedDate6,
	FileColumnIndividualServiceDeliveredDate6:           DBColumnIndividualServiceDeliveredDate6,
	FileColumnIndividualServiceComments6:                DBColumnIndividualServiceComments6,
	FileColumnIndividualServiceCC7:                      DBColumnIndividualServiceCC7,
	FileColumnIndividualServiceRequestedDate7:           DBColumnIndividualServiceRequestedDate7,
	FileColumnIndividualServiceDeliveredDate7:           DBColumnIndividualServiceDeliveredDate7,
	FileColumnIndividualServiceComments7:                DBColumnIndividualServiceComments7,
}

var IndividualFileToStructMap = map[string]string{
	FileColumnIndividualInactive:                        FormParamIndividualInactive,
	FileColumnIndividualAddress:                         FormParamIndividualAddress,
	FileColumnIndividualAge:                             FormParamIndividualAge,
	FileColumnIndividualBirthDate:                       FormParamIndividualBirthDate,
	FileColumnIndividualCognitiveDisabilityLevel:        FormParamIndividualCognitiveDisabilityLevel,
	FileColumnIndividualCollectionAdministrativeArea1:   FormParamIndividualCollectionAdministrativeArea1,
	FileColumnIndividualCollectionAdministrativeArea2:   FormParamIndividualCollectionAdministrativeArea2,
	FileColumnIndividualCollectionAdministrativeArea3:   FormParamIndividualCollectionAdministrativeArea3,
	FileColumnIndividualCollectionOffice:                FormParamIndividualCollectionOffice,
	FileColumnIndividualCollectionAgentName:             FormParamIndividualCollectionAgentName,
	FileColumnIndividualCollectionAgentTitle:            FormParamIndividualCollectionAgentTitle,
	FileColumnIndividualCollectionTime:                  FormParamIndividualCollectionTime,
	FileColumnIndividualComments:                        FormParamIndividualComments,
	FileColumnIndividualCommunicationDisabilityLevel:    FormParamIndividualCommunicationDisabilityLevel,
	FileColumnIndividualCommunityID:                     FormParamIndividualCommunityID,
	FileColumnIndividualCountryID:                       FormParamIndividualCountryID,
	FileColumnIndividualDisplacementStatus:              FormParamIndividualDisplacementStatus,
	FileColumnIndividualDisplacementStatusComment:       FormParamIndividualDisplacementStatusComment,
	FileColumnIndividualEmail1:                          FormParamIndividualEmail1,
	FileColumnIndividualEmail2:                          FormParamIndividualEmail2,
	FileColumnIndividualEmail3:                          FormParamIndividualEmail3,
	FileColumnIndividualFullName:                        FormParamIndividualFullName,
	FileColumnIndividualFirstName:                       FormParamIndividualFirstName,
	FileColumnIndividualMiddleName:                      FormParamIndividualMiddleName,
	FileColumnIndividualLastName:                        FormParamIndividualLastName,
	FileColumnIndividualFreeField1:                      FormParamIndividualFreeField1,
	FileColumnIndividualFreeField2:                      FormParamIndividualFreeField2,
	FileColumnIndividualFreeField3:                      FormParamIndividualFreeField3,
	FileColumnIndividualFreeField4:                      FormParamIndividualFreeField4,
	FileColumnIndividualFreeField5:                      FormParamIndividualFreeField5,
	FileColumnIndividualSex:                             FormParamIndividualSex,
	FileColumnIndividualHasCognitiveDisability:          FormParamIndividualHasCognitiveDisability,
	FileColumnIndividualHasCommunicationDisability:      FormParamIndividualHasCommunicationDisability,
	FileColumnIndividualHasConsentedToRGPD:              FormParamIndividualHasConsentedToRGPD,
	FileColumnIndividualHasConsentedToReferral:          FormParamIndividualHasConsentedToReferral,
	FileColumnIndividualHasHearingDisability:            FormParamIndividualHasHearingDisability,
	FileColumnIndividualHasMobilityDisability:           FormParamIndividualHasMobilityDisability,
	FileColumnIndividualHasSelfCareDisability:           FormParamIndividualHasSelfCareDisability,
	FileColumnIndividualHasVisionDisability:             FormParamIndividualHasVisionDisability,
	FileColumnIndividualHearingDisabilityLevel:          FormParamIndividualHearingDisabilityLevel,
	FileColumnIndividualHouseholdID:                     FormParamIndividualHouseholdID,
	FileColumnIndividualID:                              FormParamIndividualID,
	FileColumnIndividualEngagementContext:               FormParamIndividualEngagementContext,
	FileColumnIndividualIdentificationNumber1:           FormParamIndividualIdentificationNumber1,
	FileColumnIndividualIdentificationNumber2:           FormParamIndividualIdentificationNumber2,
	FileColumnIndividualIdentificationNumber3:           FormParamIndividualIdentificationNumber3,
	FileColumnIndividualIdentificationType1:             FormParamIndividualIdentificationType1,
	FileColumnIndividualIdentificationType2:             FormParamIndividualIdentificationType2,
	FileColumnIndividualIdentificationType3:             FormParamIndividualIdentificationType3,
	FileColumnIndividualIdentificationTypeExplanation1:  FormParamIndividualIdentificationTypeExplanation1,
	FileColumnIndividualIdentificationTypeExplanation2:  FormParamIndividualIdentificationTypeExplanation2,
	FileColumnIndividualIdentificationTypeExplanation3:  FormParamIndividualIdentificationTypeExplanation3,
	FileColumnIndividualInternalID:                      FormParamIndividualInternalID,
	FileColumnIndividualIsHeadOfCommunity:               FormParamIndividualIsHeadOfCommunity,
	FileColumnIndividualIsHeadOfHousehold:               FormParamIndividualIsHeadOfHousehold,
	FileColumnIndividualIsFemaleHeadedHousehold:         FormParamIndividualIsFemaleHeadedHousehold,
	FileColumnIndividualIsMinorHeadedHousehold:          FormParamIndividualIsMinorHeadedHousehold,
	FileColumnIndividualIsMinor:                         FormParamIndividualIsMinor,
	FileColumnIndividualIsChildAtRisk:                   FormParamIndividualIsChildAtRisk,
	FileColumnIndividualIsWomanAtRisk:                   FormParamIndividualIsWomanAtRisk,
	FileColumnIndividualIsElderAtRisk:                   FormParamIndividualIsElderAtRisk,
	FileColumnIndividualIsSeparatedChild:                FormParamIndividualIsSeparatedChild,
	FileColumnIndividualIsSingleParent:                  FormParamIndividualIsSingleParent,
	FileColumnIndividualIsPregnant:                      FormParamIndividualIsPregnant,
	FileColumnIndividualIsLactating:                     FormParamIndividualIsLactating,
	FileColumnIndividualHasMedicalCondition:             FormParamIndividualHasMedicalCondition,
	FileColumnIndividualNeedsLegalAndPhysicalProtection: FormParamIndividualNeedsLegalAndPhysicalProtection,
	FileColumnIndividualMobilityDisabilityLevel:         FormParamIndividualMobilityDisabilityLevel,
	FileColumnIndividualNationality1:                    FormParamIndividualNationality1,
	FileColumnIndividualNationality2:                    FormParamIndividualNationality2,
	FileColumnIndividualPhoneNumber1:                    FormParamIndividualPhoneNumber1,
	FileColumnIndividualPhoneNumber2:                    FormParamIndividualPhoneNumber2,
	FileColumnIndividualPhoneNumber3:                    FormParamIndividualPhoneNumber3,
	FileColumnIndividualPreferredCommunicationLanguage:  FormParamIndividualPreferredCommunicationLanguage,
	FileColumnIndividualPreferredContactMethod:          FormParamIndividualPreferredContactMethod,
	FileColumnIndividualPreferredContactMethodComments:  FormParamIndividualPreferredContactMethodComments,
	FileColumnIndividualPreferredName:                   FormParamIndividualPreferredName,
	FileColumnIndividualPrefersToRemainAnonymous:        FormParamIndividualPrefersToRemainAnonymous,
	FileColumnIndividualPresentsProtectionConcerns:      FormParamIndividualPresentsProtectionConcerns,
	FileColumnIndividualSelfCareDisabilityLevel:         FormParamIndividualSelfCareDisabilityLevel,
	FileColumnIndividualSpokenLanguage1:                 FormParamIndividualSpokenLanguage1,
	FileColumnIndividualSpokenLanguage2:                 FormParamIndividualSpokenLanguage2,
	FileColumnIndividualSpokenLanguage3:                 FormParamIndividualSpokenLanguage3,
	FileColumnIndividualVisionDisabilityLevel:           FormParamIndividualVisionDisabilityLevel,
	FileColumnIndividualServiceCC1:                      FormParamIndividualServiceCC1,
	FileColumnIndividualServiceRequestedDate1:           FormParamIndividualServiceRequestedDate1,
	FileColumnIndividualServiceDeliveredDate1:           FormParamIndividualServiceDeliveredDate1,
	FileColumnIndividualServiceComments1:                FormParamIndividualServiceComments1,
	FileColumnIndividualServiceCC2:                      FormParamIndividualServiceCC2,
	FileColumnIndividualServiceRequestedDate2:           FormParamIndividualServiceRequestedDate2,
	FileColumnIndividualServiceDeliveredDate2:           FormParamIndividualServiceDeliveredDate2,
	FileColumnIndividualServiceComments2:                FormParamIndividualServiceComments2,
	FileColumnIndividualServiceCC3:                      FormParamIndividualServiceCC3,
	FileColumnIndividualServiceRequestedDate3:           FormParamIndividualServiceRequestedDate3,
	FileColumnIndividualServiceDeliveredDate3:           FormParamIndividualServiceDeliveredDate3,
	FileColumnIndividualServiceComments3:                FormParamIndividualServiceComments3,
	FileColumnIndividualServiceCC4:                      FormParamIndividualServiceCC4,
	FileColumnIndividualServiceRequestedDate4:           FormParamIndividualServiceRequestedDate4,
	FileColumnIndividualServiceDeliveredDate4:           FormParamIndividualServiceDeliveredDate4,
	FileColumnIndividualServiceComments4:                FormParamIndividualServiceComments4,
	FileColumnIndividualServiceCC5:                      FormParamIndividualServiceCC5,
	FileColumnIndividualServiceRequestedDate5:           FormParamIndividualServiceRequestedDate5,
	FileColumnIndividualServiceDeliveredDate5:           FormParamIndividualServiceDeliveredDate5,
	FileColumnIndividualServiceComments5:                FormParamIndividualServiceComments5,
	FileColumnIndividualServiceCC6:                      FormParamIndividualServiceCC6,
	FileColumnIndividualServiceRequestedDate6:           FormParamIndividualServiceRequestedDate6,
	FileColumnIndividualServiceDeliveredDate6:           FormParamIndividualServiceDeliveredDate6,
	FileColumnIndividualServiceComments6:                FormParamIndividualServiceComments6,
	FileColumnIndividualServiceCC7:                      FormParamIndividualServiceCC7,
	FileColumnIndividualServiceRequestedDate7:           FormParamIndividualServiceRequestedDate7,
	FileColumnIndividualServiceDeliveredDate7:           FormParamIndividualServiceDeliveredDate7,
	FileColumnIndividualServiceComments7:                FormParamIndividualServiceComments7,
}

var IndividualSystemFileColumns = containers.NewStringSet(
	FileColumnIndividualCreatedAt,
	FileColumnIndividualUpdatedAt,
)
