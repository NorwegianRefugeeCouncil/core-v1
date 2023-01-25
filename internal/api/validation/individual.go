package validation

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/validation"
)

func ValidateIndividual(i *api.Individual) validation.ErrorList {
	return validateIndividual(i, nil)
}

func ValidateIndividualList(i *api.IndividualList) validation.ErrorList {
	allErrs := validation.ErrorList{}
	itemsPath := validation.NewPath("items")
	for i, individual := range i.Items {
		allErrs = append(allErrs, validateIndividual(individual, itemsPath.Index(i))...)
	}
	return allErrs
}

func validateIndividual(i *api.Individual, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	allErrs = append(allErrs, validateIndividualAddress(i.Address, p.Child("address"))...)
	allErrs = append(allErrs, validateIndividualAdministrativeArea(i.CollectionAdministrativeArea1, p.Child("collectionAdministrativeArea1"))...)
	allErrs = append(allErrs, validateIndividualAdministrativeArea(i.CollectionAdministrativeArea2, p.Child("collectionAdministrativeArea2"))...)
	allErrs = append(allErrs, validateIndividualAdministrativeArea(i.CollectionAdministrativeArea3, p.Child("collectionAdministrativeArea3"))...)
	allErrs = append(allErrs, validateIndividualAge(i.Age, p.Child("age"))...)
	allErrs = append(allErrs, validateIndividualBirthDate(i.BirthDate, p.Child("birthDate"))...)
	allErrs = append(allErrs, validateIndividualCollectionAgentName(i.CollectionAgentName, p.Child("collectionAgentName"))...)
	allErrs = append(allErrs, validateIndividualCollectionAgentTitle(i.CollectionAgentTitle, p.Child("collectionAgentTitle"))...)
	allErrs = append(allErrs, validateIndividualCollectionOffice(i.CollectionOffice, p.Child("collectionOffice"))...)
	allErrs = append(allErrs, validateIndividualCollectionTime(i.CollectionTime, p.Child("collectionTime"))...)
	allErrs = append(allErrs, validateIndividualComments(i.Comments, p.Child("comments"))...)
	allErrs = append(allErrs, validateIndividualCommunityID(i.CommunityID, p.Child("communityId"))...)
	allErrs = append(allErrs, validateIndividualCountryID(i.CountryID, p.Child("countryId"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.CognitiveDisabilityLevel, p.Child("cognitiveDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.CommunicationDisabilityLevel, p.Child("communicationDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.HearingDisabilityLevel, p.Child("hearingDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.MobilityDisabilityLevel, p.Child("mobilityDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.SelfCareDisabilityLevel, p.Child("selfCareDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.VisionDisabilityLevel, p.Child("visionDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisplacementStatus(i.DisplacementStatus, p.Child("displacementStatus"))...)
	allErrs = append(allErrs, validateIndividualDisplacementStatusComment(i.DisplacementStatusComment, p.Child("displacementStatusComment"))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email1, p.Child("email1"))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email2, p.Child("email2"))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email3, p.Child("email3"))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField1, p.Child("freeField1"))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField2, p.Child("freeField2"))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField3, p.Child("freeField3"))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField4, p.Child("freeField4"))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField5, p.Child("freeField5"))...)
	allErrs = append(allErrs, validateIndividualName(i.FullName, p.Child("fullName"))...)
	allErrs = append(allErrs, validateIndividualName(i.FirstName, p.Child("firstName"))...)
	allErrs = append(allErrs, validateIndividualName(i.MiddleName, p.Child("middleName"))...)
	allErrs = append(allErrs, validateIndividualName(i.LastName, p.Child("lastName"))...)
	allErrs = append(allErrs, validateIndividualSex(i.Sex, p.Child("sex"))...)
	allErrs = append(allErrs, validateIndividualHouseholdID(i.HouseholdID, p.Child("householdId"))...)
	allErrs = append(allErrs, validateIndividualEngagementContext(i.EngagementContext, p.Child("engagementContext"))...)
	allErrs = append(allErrs, validateIndividualIdentificationNumber(i.IdentificationNumber1, p.Child("identificationNumber1"))...)
	allErrs = append(allErrs, validateIndividualIdentificationNumber(i.IdentificationNumber2, p.Child("identificationNumber2"))...)
	allErrs = append(allErrs, validateIndividualIdentificationNumber(i.IdentificationNumber3, p.Child("identificationNumber3"))...)
	allErrs = append(allErrs, validateIndividualIdentificationType(i.IdentificationType1, p.Child("identificationType1"))...)
	allErrs = append(allErrs, validateIndividualIdentificationType(i.IdentificationType2, p.Child("identificationType2"))...)
	allErrs = append(allErrs, validateIndividualIdentificationType(i.IdentificationType3, p.Child("identificationType3"))...)
	allErrs = append(allErrs, validateIndividualIdentificationTypeExplanation(i.IdentificationTypeExplanation1, p.Child("identificationTypeExplanation1"))...)
	allErrs = append(allErrs, validateIndividualIdentificationTypeExplanation(i.IdentificationTypeExplanation2, p.Child("identificationTypeExplanation2"))...)
	allErrs = append(allErrs, validateIndividualIdentificationTypeExplanation(i.IdentificationTypeExplanation3, p.Child("identificationTypeExplanation3"))...)
	allErrs = append(allErrs, validateIndividualInternalID(i.InternalID, p.Child("internalId"))...)
	allErrs = append(allErrs, validateIndividualNationality(i.Nationality1, p.Child("nationality1"))...)
	allErrs = append(allErrs, validateIndividualNationality(i.Nationality2, p.Child("nationality2"))...)
	allErrs = append(allErrs, validateIndividualPhoneNumber(i.PhoneNumber1, p.Child("phoneNumber1"))...)
	allErrs = append(allErrs, validateIndividualPhoneNumber(i.PhoneNumber2, p.Child("phoneNumber2"))...)
	allErrs = append(allErrs, validateIndividualPhoneNumber(i.PhoneNumber3, p.Child("phoneNumber3"))...)
	allErrs = append(allErrs, validateIndividualPreferredCommunicationLanguage(i.PreferredCommunicationLanguage, p.Child("preferredCommunicationLanguage"))...)
	allErrs = append(allErrs, validateIndividualPreferredContactMethod(i.PreferredContactMethod, p.Child("preferredContactMethod"))...)
	allErrs = append(allErrs, validateIndividualPreferredContactMethodComments(i.PreferredContactMethodComments, p.Child("preferredContactMethodComments"))...)
	allErrs = append(allErrs, validateIndividualPreferredName(i.PreferredName, p.Child("preferredName"))...)
	allErrs = append(allErrs, validateIndividualSpokenLanguage(i.SpokenLanguage1, p.Child("spokenLanguage1"))...)
	allErrs = append(allErrs, validateIndividualSpokenLanguage(i.SpokenLanguage2, p.Child("spokenLanguage2"))...)
	allErrs = append(allErrs, validateIndividualSpokenLanguage(i.SpokenLanguage3, p.Child("spokenLanguage3"))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC1, p.Child("serviceCC1"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate1, p.Child("serviceRequestedDate1"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate1, p.Child("serviceDeliveredDate1"))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments1, p.Child("serviceComments1"))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC2, p.Child("serviceCC2"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate2, p.Child("serviceRequestedDate2"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate2, p.Child("serviceDeliveredDate2"))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments2, p.Child("serviceComments2"))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC3, p.Child("serviceCC3"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate3, p.Child("serviceRequestedDate3"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate3, p.Child("serviceDeliveredDate3"))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments3, p.Child("serviceComments3"))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC4, p.Child("serviceCC4"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate4, p.Child("serviceRequestedDate4"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate4, p.Child("serviceDeliveredDate4"))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments4, p.Child("serviceComments4"))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC5, p.Child("serviceCC5"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate5, p.Child("serviceRequestedDate5"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate5, p.Child("serviceDeliveredDate5"))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments5, p.Child("serviceComments5"))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC6, p.Child("serviceCC6"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate6, p.Child("serviceRequestedDate6"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate6, p.Child("serviceDeliveredDate6"))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments6, p.Child("serviceComments6"))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC7, p.Child("serviceCC7"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate7, p.Child("serviceRequestedDate7"))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate7, p.Child("serviceDeliveredDate7"))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments7, p.Child("serviceComments7"))...)
	return allErrs
}

const (
	individualAddressMaxLength                        = 512
	individualAdministrativeAreaMaxLength             = 128
	individualCollectionOfficeMaxLength               = 128
	individualCollectionAgentNameMaxLength            = 64
	individualCollectionAgentTitleMaxLength           = 64
	individualCommunityIDMaxLength                    = 64
	individualEmailMaxLength                          = 255
	individualNameMaxLength                           = 255
	individualFreeFieldMaxLength                      = 255
	individualHouseholdIDMaxLength                    = 64
	individualIdentificationTypeMaxLength             = 64
	individualIdentificationNumberMaxLength           = 64
	individualEngagementContextMaxLength              = 64
	individualInternalIDMaxLength                     = 64
	individualNationalityMaxLength                    = 64
	individualPhoneNumberMaxLength                    = 64
	individualPreferredContactMethodMaxLength         = 64
	individualPreferredNameMaxLength                  = 255
	individualPreferredCommunicationLanguageMaxLength = 64
	individualSpokenLanguageMaxLength                 = 64
	maxTextLength                                     = 65535
)

func validateIndividualAddress(address string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(address) > individualAddressMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, address, individualAddressMaxLength))
	}
	return allErrs
}

func validateIndividualAge(age *int, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if age == nil {
		return allErrs
	}
	if *age < 0 {
		allErrs = append(allErrs, validation.Invalid(path, age, "must be greater than or equal to 0"))
	}
	return allErrs
}

func validateIndividualAdministrativeArea(administrativeArea string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(administrativeArea) > individualAdministrativeAreaMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, administrativeArea, individualAdministrativeAreaMaxLength))
	}
	return allErrs
}

func validateIndividualCollectionOffice(collectionOffice string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(collectionOffice) > individualCollectionOfficeMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, collectionOffice, individualCollectionOfficeMaxLength))
	}
	return allErrs
}

func validateIndividualCollectionAgentName(collectionAgentName string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(collectionAgentName) > individualCollectionAgentNameMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, collectionAgentName, individualCollectionAgentNameMaxLength))
	}
	return allErrs
}

func validateIndividualCollectionAgentTitle(collectionAgentTitle string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(collectionAgentTitle) > individualCollectionAgentTitleMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, collectionAgentTitle, individualCollectionAgentTitleMaxLength))
	}
	return allErrs
}

func validateIndividualCollectionTime(collectionTime time.Time, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	now := time.Now()
	if collectionTime.After(now) {
		allErrs = append(allErrs, validation.Invalid(path, collectionTime, "collection time cannot be in the future"))
	}
	return allErrs
}

func validateIndividualComments(comments string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(comments) == 0 {
		return allErrs
	}
	if len(comments) > maxTextLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, comments, maxTextLength))
	}
	return allErrs
}

func validateIndividualCommunityID(communityID string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(communityID) > individualCommunityIDMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, communityID, individualCommunityIDMaxLength))
	}
	return allErrs
}

func validateIndividualBirthDate(birthDate *time.Time, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// birthDate is optional
	if birthDate == nil {
		return allErrs
	}
	if birthDate.After(time.Now()) {
		allErrs = append(allErrs, validation.Invalid(path, birthDate, "birthdate cannot be in the future"))
	}
	if *birthDate == (time.Time{}) {
		allErrs = append(allErrs, validation.Invalid(path, birthDate, "must be a valid date"))
	}
	return allErrs
}

func validateIndividualCountryID(countryID string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if countryID == "" {
		allErrs = append(allErrs, validation.Required(path, "country id is required"))
	} else if _, err := uuid.Parse(countryID); err != nil {
		allErrs = append(allErrs, validation.Invalid(path, countryID, "must be a valid UUID"))
	}
	return allErrs
}

func validateIndividualDisplacementStatus(ds api.DisplacementStatus, path *validation.Path) validation.ErrorList {
	switch {
	case allowedDisplacementStatuses.Contains(ds):
		return validation.ErrorList{}
	case ds == api.DisplacementStatusUnspecified:
		return validation.ErrorList{}
	default:
		return validation.ErrorList{validation.NotSupported(path, ds, allowedDisplacementStatusesStr)}
	}
}

func validateIndividualDisplacementStatusComment(dsc string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(dsc) > maxTextLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, dsc, maxTextLength))
	}
	return allErrs
}

func validateIndividualSex(sex api.Sex, path *validation.Path) validation.ErrorList {
	switch {
	case allowedSexes.Contains(sex):
		return validation.ErrorList{}
	case sex == api.SexUnspecified:
		return validation.ErrorList{}
	default:
		return validation.ErrorList{validation.NotSupported(path, sex, allowedSexesStr)}
	}
}

func validateIndividualEmail(email string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(email) == 0 {
		return allErrs
	}

	if len(email) > individualEmailMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, email, individualEmailMaxLength))
	} else if _, err := mail.ParseAddress(email); err != nil {
		allErrs = append(allErrs, validation.Invalid(path, email, "invalid email address"))
	}

	return allErrs
}

func validateIndividualName(name string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// email is optional
	if len(name) > individualNameMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, name, individualNameMaxLength))
	}
	return allErrs
}

func validateIndividualFreeField(freeField string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// email is optional
	if len(freeField) > individualFreeFieldMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, freeField, individualFreeFieldMaxLength))
	}
	return allErrs
}

func validateIndividualDisabilityLevel(disabilityLevel api.DisabilityLevel, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// disability level is optional
	if disabilityLevel == api.DisabilityLevelUnspecified {
		return nil
	}
	if !allowedDisabilityLevels.Contains(disabilityLevel) {
		allErrs = append(allErrs, validation.NotSupported(path, disabilityLevel, allowedDisabilityLevelsStr))
	}
	return allErrs
}

func validateIndividualHouseholdID(householdID string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(householdID) > individualHouseholdIDMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, householdID, individualHouseholdIDMaxLength))
	}
	return allErrs
}

func validateIndividualIdentificationType(identificationType string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(identificationType) > individualIdentificationTypeMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, identificationType, individualIdentificationTypeMaxLength))
	}
	return allErrs
}

func validateIndividualIdentificationTypeExplanation(explanation string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(explanation) > maxTextLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, explanation, maxTextLength))
	}
	return allErrs
}
func validateIndividualIdentificationNumber(number string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(number) > individualIdentificationNumberMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, number, individualIdentificationNumberMaxLength))
	}
	return allErrs
}

func validateIndividualEngagementContext(context string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(context) > individualEngagementContextMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, context, individualEngagementContextMaxLength))
	}
	return allErrs
}

func validateIndividualInternalID(internalID string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(internalID) > individualInternalIDMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, internalID, individualInternalIDMaxLength))
	}
	return allErrs
}

func validateIndividualNationality(nationality string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(nationality) > individualNationalityMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, nationality, individualNationalityMaxLength))
	}
	return allErrs
}

func validateIndividualPhoneNumber(phoneNumber string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(phoneNumber) > individualPhoneNumberMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, phoneNumber, individualPhoneNumberMaxLength))
	}
	return allErrs
}

func validateIndividualPreferredContactMethod(preferredContactMethod string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(preferredContactMethod) > individualPreferredContactMethodMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, preferredContactMethod, individualPreferredContactMethodMaxLength))
	}
	return allErrs
}

func validateIndividualPreferredContactMethodComments(comments string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(comments) > maxTextLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, comments, maxTextLength))
	}
	return allErrs
}

func validateIndividualPreferredName(preferredName string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(preferredName) > individualPreferredNameMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, preferredName, individualPreferredNameMaxLength))
	}
	return allErrs
}

func validateIndividualPreferredCommunicationLanguage(language string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(language) > individualPreferredCommunicationLanguageMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, language, individualPreferredCommunicationLanguageMaxLength))
	}
	return allErrs
}

func validateIndividualSpokenLanguage(language string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(language) > individualSpokenLanguageMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, language, individualSpokenLanguageMaxLength))
	}
	return allErrs
}

func validateIndividualServiceCC(cc api.ServiceCC, path *validation.Path) validation.ErrorList {
	if allowedServiceCCs.Contains(cc) {
		return validation.ErrorList{}
	}
	return validation.ErrorList{validation.NotSupported(path, cc, allowedServiceCCsStr)}
}

func validateIndividualServiceDate(d *time.Time, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if d == nil {
		return allErrs
	}
	if d.After(time.Now()) {
		allErrs = append(allErrs, validation.Invalid(path, d, "service date cannot be in the future"))
	}
	if *d == (time.Time{}) {
		allErrs = append(allErrs, validation.Invalid(path, d, "must be a valid date"))
	}
	return allErrs
}

func validateIndividualServiceComment(comment string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(comment) > maxTextLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, comment, maxTextLength))
	}
	return allErrs
}
