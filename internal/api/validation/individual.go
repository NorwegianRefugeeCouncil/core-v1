package validation

import (
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/constants"
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
	allErrs = append(allErrs, validateIndividualAddress(i.Address, p.Child(constants.DBColumnIndividualAddress))...)
	allErrs = append(allErrs, validateIndividualAdministrativeArea(i.CollectionAdministrativeArea1, p.Child(constants.DBColumnIndividualCollectionAdministrativeArea1))...)
	allErrs = append(allErrs, validateIndividualAdministrativeArea(i.CollectionAdministrativeArea2, p.Child(constants.DBColumnIndividualCollectionAdministrativeArea2))...)
	allErrs = append(allErrs, validateIndividualAdministrativeArea(i.CollectionAdministrativeArea3, p.Child(constants.DBColumnIndividualCollectionAdministrativeArea3))...)
	allErrs = append(allErrs, validateIndividualAge(i.Age, p.Child(constants.DBColumnIndividualAge))...)
	allErrs = append(allErrs, validateIndividualBirthDate(i.BirthDate, p.Child(constants.DBColumnIndividualBirthDate))...)
	allErrs = append(allErrs, validateIndividualCollectionAgentName(i.CollectionAgentName, p.Child(constants.DBColumnIndividualCollectionAgentName))...)
	allErrs = append(allErrs, validateIndividualCollectionAgentTitle(i.CollectionAgentTitle, p.Child(constants.DBColumnIndividualCollectionAgentTitle))...)
	allErrs = append(allErrs, validateIndividualCollectionOffice(i.CollectionOffice, p.Child(constants.DBColumnIndividualCollectionOffice))...)
	allErrs = append(allErrs, validateIndividualCollectionTime(i.CollectionTime, p.Child(constants.DBColumnIndividualCollectionTime))...)
	allErrs = append(allErrs, validateIndividualComments(i.Comments, p.Child(constants.DBColumnIndividualComments))...)
	allErrs = append(allErrs, validateIndividualCommunityID(i.CommunityID, p.Child(constants.DBColumnIndividualCommunityID))...)
	allErrs = append(allErrs, validateIndividualCommunitySize(i.CommunitySize, p.Child(constants.DBColumnIndividualCommunitySize))...)
	allErrs = append(allErrs, validateIndividualCountryID(i.CountryID, p.Child(constants.DBColumnIndividualCountryID))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.CognitiveDisabilityLevel, p.Child(constants.DBColumnIndividualCognitiveDisabilityLevel))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.CommunicationDisabilityLevel, p.Child(constants.DBColumnIndividualCommunicationDisabilityLevel))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.HearingDisabilityLevel, p.Child(constants.DBColumnIndividualHearingDisabilityLevel))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.MobilityDisabilityLevel, p.Child(constants.DBColumnIndividualMobilityDisabilityLevel))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.SelfCareDisabilityLevel, p.Child(constants.DBColumnIndividualSelfCareDisabilityLevel))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.VisionDisabilityLevel, p.Child(constants.DBColumnIndividualVisionDisabilityLevel))...)
	allErrs = append(allErrs, validateIndividualDisplacementStatus(i.DisplacementStatus, p.Child(constants.DBColumnIndividualDisplacementStatus))...)
	allErrs = append(allErrs, validateIndividualDisplacementStatusComment(i.DisplacementStatusComment, p.Child(constants.DBColumnIndividualDisplacementStatusComment))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email1, p.Child(constants.DBColumnIndividualEmail1))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email2, p.Child(constants.DBColumnIndividualEmail2))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email3, p.Child(constants.DBColumnIndividualEmail3))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField1, p.Child(constants.DBColumnIndividualFreeField1))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField2, p.Child(constants.DBColumnIndividualFreeField2))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField3, p.Child(constants.DBColumnIndividualFreeField3))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField4, p.Child(constants.DBColumnIndividualFreeField4))...)
	allErrs = append(allErrs, validateIndividualFreeField(i.FreeField5, p.Child(constants.DBColumnIndividualFreeField5))...)
	allErrs = append(allErrs, validateIndividualName(i.FullName, p.Child(constants.DBColumnIndividualFullName))...)
	allErrs = append(allErrs, validateIndividualName(i.FirstName, p.Child(constants.DBColumnIndividualFirstName))...)
	allErrs = append(allErrs, validateIndividualName(i.MiddleName, p.Child(constants.DBColumnIndividualMiddleName))...)
	allErrs = append(allErrs, validateIndividualName(i.LastName, p.Child(constants.DBColumnIndividualLastName))...)
	allErrs = append(allErrs, validateIndividualName(i.NativeName, p.Child(constants.DBColumnIndividualNativeName))...)
	allErrs = append(allErrs, validateIndividualName(i.MothersName, p.Child(constants.DBColumnIndividualMothersName))...)
	allErrs = append(allErrs, validateIndividualSex(i.Sex, p.Child(constants.DBColumnIndividualSex))...)
	allErrs = append(allErrs, validateIndividualHouseholdID(i.HouseholdID, p.Child(constants.DBColumnIndividualHouseholdID))...)
	allErrs = append(allErrs, validateIndividualHouseholdSize(i.HouseholdSize, p.Child(constants.DBColumnIndividualHouseholdSize))...)
	allErrs = append(allErrs, validateIndividualEngagementContext(i.EngagementContext, p.Child(constants.DBColumnIndividualEngagementContext))...)
	allErrs = append(allErrs, validateIndividualIdentificationNumber(i.IdentificationNumber1, p.Child(constants.DBColumnIndividualIdentificationNumber1))...)
	allErrs = append(allErrs, validateIndividualIdentificationNumber(i.IdentificationNumber2, p.Child(constants.DBColumnIndividualIdentificationNumber2))...)
	allErrs = append(allErrs, validateIndividualIdentificationNumber(i.IdentificationNumber3, p.Child(constants.DBColumnIndividualIdentificationNumber3))...)
	allErrs = append(allErrs, validateIndividualIdentificationType(i.IdentificationType1, p.Child(constants.DBColumnIndividualIdentificationType1))...)
	allErrs = append(allErrs, validateIndividualIdentificationType(i.IdentificationType2, p.Child(constants.DBColumnIndividualIdentificationType2))...)
	allErrs = append(allErrs, validateIndividualIdentificationType(i.IdentificationType3, p.Child(constants.DBColumnIndividualIdentificationType3))...)
	allErrs = append(allErrs, validateIndividualIdentificationTypeExplanation(i.IdentificationTypeExplanation1, p.Child(constants.DBColumnIndividualIdentificationTypeExplanation1))...)
	allErrs = append(allErrs, validateIndividualIdentificationTypeExplanation(i.IdentificationTypeExplanation2, p.Child(constants.DBColumnIndividualIdentificationTypeExplanation2))...)
	allErrs = append(allErrs, validateIndividualIdentificationTypeExplanation(i.IdentificationTypeExplanation3, p.Child(constants.DBColumnIndividualIdentificationTypeExplanation3))...)
	allErrs = append(allErrs, validateIndividualInternalID(i.InternalID, p.Child(constants.DBColumnIndividualInternalID))...)
	allErrs = append(allErrs, validateIndividualNationality(i.Nationality1, p.Child(constants.DBColumnIndividualNationality1))...)
	allErrs = append(allErrs, validateIndividualNationality(i.Nationality2, p.Child(constants.DBColumnIndividualNationality2))...)
	allErrs = append(allErrs, validateIndividualPhoneNumber(i.PhoneNumber1, p.Child(constants.DBColumnIndividualPhoneNumber1))...)
	allErrs = append(allErrs, validateIndividualPhoneNumber(i.PhoneNumber2, p.Child(constants.DBColumnIndividualPhoneNumber2))...)
	allErrs = append(allErrs, validateIndividualPhoneNumber(i.PhoneNumber3, p.Child(constants.DBColumnIndividualPhoneNumber3))...)
	allErrs = append(allErrs, validateIndividualPreferredCommunicationLanguage(i.PreferredCommunicationLanguage, p.Child(constants.DBColumnIndividualPreferredCommunicationLanguage))...)
	allErrs = append(allErrs, validateIndividualPreferredContactMethod(i.PreferredContactMethod, p.Child(constants.DBColumnIndividualPreferredContactMethod))...)
	allErrs = append(allErrs, validateIndividualPreferredContactMethodComments(i.PreferredContactMethodComments, p.Child(constants.DBColumnIndividualPreferredContactMethodComments))...)
	allErrs = append(allErrs, validateIndividualPreferredName(i.PreferredName, p.Child(constants.DBColumnIndividualPreferredName))...)
	allErrs = append(allErrs, validateIndividualSpokenLanguage(i.SpokenLanguage1, p.Child(constants.DBColumnIndividualSpokenLanguage1))...)
	allErrs = append(allErrs, validateIndividualSpokenLanguage(i.SpokenLanguage2, p.Child(constants.DBColumnIndividualSpokenLanguage2))...)
	allErrs = append(allErrs, validateIndividualSpokenLanguage(i.SpokenLanguage3, p.Child(constants.DBColumnIndividualSpokenLanguage3))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC1, p.Child(constants.DBColumnIndividualServiceCC1))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate1, p.Child(constants.DBColumnIndividualServiceRequestedDate1))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate1, p.Child(constants.DBColumnIndividualServiceDeliveredDate1))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments1, p.Child(constants.DBColumnIndividualServiceComments1))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC2, p.Child(constants.DBColumnIndividualServiceCC2))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate2, p.Child(constants.DBColumnIndividualServiceRequestedDate2))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate2, p.Child(constants.DBColumnIndividualServiceDeliveredDate2))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments2, p.Child(constants.DBColumnIndividualServiceComments2))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC3, p.Child(constants.DBColumnIndividualServiceCC3))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate3, p.Child(constants.DBColumnIndividualServiceRequestedDate3))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate3, p.Child(constants.DBColumnIndividualServiceDeliveredDate3))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments3, p.Child(constants.DBColumnIndividualServiceComments3))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC4, p.Child(constants.DBColumnIndividualServiceCC4))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate4, p.Child(constants.DBColumnIndividualServiceRequestedDate4))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate4, p.Child(constants.DBColumnIndividualServiceDeliveredDate4))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments4, p.Child(constants.DBColumnIndividualServiceComments4))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC5, p.Child(constants.DBColumnIndividualServiceCC5))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate5, p.Child(constants.DBColumnIndividualServiceRequestedDate5))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate5, p.Child(constants.DBColumnIndividualServiceDeliveredDate5))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments5, p.Child(constants.DBColumnIndividualServiceComments5))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC6, p.Child(constants.DBColumnIndividualServiceCC6))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate6, p.Child(constants.DBColumnIndividualServiceRequestedDate6))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate6, p.Child(constants.DBColumnIndividualServiceDeliveredDate6))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments6, p.Child(constants.DBColumnIndividualServiceComments6))...)
	allErrs = append(allErrs, validateIndividualServiceCC(i.ServiceCC7, p.Child(constants.DBColumnIndividualServiceCC7))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceRequestedDate7, p.Child(constants.DBColumnIndividualServiceRequestedDate7))...)
	allErrs = append(allErrs, validateIndividualServiceDate(i.ServiceDeliveredDate7, p.Child(constants.DBColumnIndividualServiceDeliveredDate7))...)
	allErrs = append(allErrs, validateIndividualServiceComment(i.ServiceComments7, p.Child(constants.DBColumnIndividualServiceComments7))...)
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

func validateIndividualHouseholdSize(householdSize *int, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if householdSize == nil {
		return allErrs
	}
	if *householdSize < 0 {
		allErrs = append(allErrs, validation.Invalid(path, householdSize, "must be greater than or equal to 0"))
	}
	return allErrs
}

func validateIndividualCommunitySize(communitySize *int, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if communitySize == nil {
		return allErrs
	}
	if *communitySize < 0 {
		allErrs = append(allErrs, validation.Invalid(path, communitySize, "must be greater than or equal to 0"))
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
	if birthDate.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		allErrs = append(allErrs, validation.Invalid(path, birthDate, "birthdate cannot be before 1900-01-01"))
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

func validateIndividualDisplacementStatus(ds enumTypes.DisplacementStatus, path *validation.Path) validation.ErrorList {
	switch {
	case allowedDisplacementStatuses.Contains(ds):
		return validation.ErrorList{}
	case ds == enumTypes.DisplacementStatusUnspecified:
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

func validateIndividualSex(sex enumTypes.Sex, path *validation.Path) validation.ErrorList {
	switch {
	case allowedSexes.Contains(sex):
		return validation.ErrorList{}
	case sex == enumTypes.SexUnspecified:
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

func validateIndividualDisabilityLevel(disabilityLevel enumTypes.DisabilityLevel, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// disability level is optional
	if disabilityLevel == enumTypes.DisabilityLevelUnspecified {
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

func validateIndividualIdentificationType(identificationType enumTypes.IdentificationType, path *validation.Path) validation.ErrorList {
	if identificationType == enumTypes.IdentificationTypeUnspecified {
		return nil
	}
	if allowedIdentificationTypes.Contains(identificationType) {
		return validation.ErrorList{}
	}
	return validation.ErrorList{validation.NotSupported(path, identificationType, allowedIdentificationTypesStr)}
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

func validateIndividualEngagementContext(context enumTypes.EngagementContext, path *validation.Path) validation.ErrorList {
	if context == enumTypes.EngagementContextUnspecified {
		return nil
	}
	if allowedEngagementContexts.Contains(context) {
		return validation.ErrorList{}
	}
	return validation.ErrorList{validation.NotSupported(path, context, allowedEngagementContextsStr)}
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

func validateIndividualPreferredContactMethod(preferredContactMethod enumTypes.ContactMethod, path *validation.Path) validation.ErrorList {
	if preferredContactMethod == enumTypes.ContactMethodUnspecified {
		return nil
	}
	if allowedContactMethods.Contains(preferredContactMethod) {
		return validation.ErrorList{}
	}
	return validation.ErrorList{validation.NotSupported(path, preferredContactMethod, allowedContactMethodsStr)}
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

func validateIndividualServiceCC(cc enumTypes.ServiceCC, path *validation.Path) validation.ErrorList {
	if cc == enumTypes.ServiceCCNone {
		return nil
	}
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
