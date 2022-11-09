package validation

import (
	"net/mail"
	"time"

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
	allErrs = append(allErrs, validateIndividualCountryID(i.CountryID, p.Child("countryId"))...)
	allErrs = append(allErrs, validateIndividualBirthDate(i.BirthDate, p.Child("birthDate"))...)
	allErrs = append(allErrs, validateIndividualDisplacementStatus(i.DisplacementStatus, p.Child("displacementStatus"))...)
	allErrs = append(allErrs, validateIndividualGender(i.Gender, p.Child("gender"))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email1, p.Child("email1"))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email2, p.Child("email2"))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email3, p.Child("email3"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.CognitiveDisabilityLevel, p.Child("cognitiveDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.CommunicationDisabilityLevel, p.Child("communicationDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.HearingDisabilityLevel, p.Child("hearingDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.MobilityDisabilityLevel, p.Child("mobilityDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.SelfCareDisabilityLevel, p.Child("selfCareDisabilityLevel"))...)
	allErrs = append(allErrs, validateIndividualDisabilityLevel(i.VisionDisabilityLevel, p.Child("visionDisabilityLevel"))...)

	return allErrs
}

func validateIndividualBirthDate(birthDate *time.Time, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// birthDate is optional
	if birthDate != nil {
		if birthDate.After(time.Now()) {
			allErrs = append(allErrs, validation.Invalid(path, birthDate, "birthdate cannot be in the future"))
		}
		if *birthDate == (time.Time{}) {
			allErrs = append(allErrs, validation.Invalid(path, birthDate, "must be a valid date"))
		}
	}
	return allErrs
}

func validateIndividualCountryID(countryID string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if countryID == "" {
		allErrs = append(allErrs, validation.Required(path, "country id is required"))
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

func validateIndividualGender(gender api.Gender, path *validation.Path) validation.ErrorList {
	switch {
	case allowedGenders.Contains(gender):
		return validation.ErrorList{}
	case gender == api.GenderUnspecified:
		return validation.ErrorList{}
	default:
		return validation.ErrorList{validation.NotSupported(path, gender, allowedGendersStr)}
	}
}

func validateIndividualEmail(email string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// email is optional
	if len(email) != 0 {
		if _, err := mail.ParseAddress(email); err != nil {
			allErrs = append(allErrs, validation.Invalid(path, email, "invalid email address"))
		}
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
