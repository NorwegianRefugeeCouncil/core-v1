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
	allErrs = append(allErrs, validateIndividualEmail(i.Email, p.Child("email"))...)
	allErrs = append(allErrs, validateIndividualPreferredContactMethod(i.PreferredContactMethod, p.Child("preferredContactMethod"))...)
	allErrs = append(allErrs, validateIndividualCollectionAgentName(i.CollectionAgentName, p.Child("collectionAgentName"))...)
	allErrs = append(allErrs, validateIndividualCollectionAgentTitle(i.CollectionAgentTitle, p.Child("collectionAgentTitle"))...)
	allErrs = append(allErrs, validateIndividualDateOfRegistration(i.CollectionTime, p.Child("collectionTime"))...)
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
	case len(ds) == 0:
		return validation.ErrorList{validation.Required(path, "displacement status is required")}
	default:
		return validation.ErrorList{validation.NotSupported(path, ds, allowedDisplacementStatusesStr)}
	}
}

func validateIndividualGender(gender api.Gender, path *validation.Path) validation.ErrorList {
	switch {
	case allowedGenders.Contains(gender):
		return validation.ErrorList{}
	case len(gender) == 0:
		return validation.ErrorList{validation.Required(path, "gender is required")}
	default:
		return validation.ErrorList{validation.NotSupported(path, gender, allowedGendersStr)}
	}
}

func validateIndividualPreferredContactMethod(pcm string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(pcm) == 0 {
		allErrs = append(allErrs, validation.Required(path, "preferred contact method is required"))
	}
	return allErrs
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

func validateIndividualCollectionAgentName(name string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(name) == 0 {
		allErrs = append(allErrs, validation.Required(path, "collection agent name is required"))
	}
	return allErrs
}

func validateIndividualCollectionAgentTitle(name string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if len(name) == 0 {
		allErrs = append(allErrs, validation.Required(path, "collection agent title is required"))
	}
	return allErrs
}

func validateIndividualDateOfRegistration(dateOfRegistration time.Time, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if dateOfRegistration.IsZero() {
		allErrs = append(allErrs, validation.Required(path, "date of registration is required"))
	}
	return allErrs
}
