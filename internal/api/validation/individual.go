package validation

import (
	"net/mail"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/validation"
)

func ValidateIndividual(i *api.Individual) validation.ErrorList {
	allErrs := validation.ErrorList{}
	allErrs = append(allErrs, validateIndividualCountryID(i.CountryID, validation.NewPath("countryId"))...)
	allErrs = append(allErrs, validateBirthDate(i.BirthDate, validation.NewPath("birthDate"))...)
	allErrs = append(allErrs, validateIndividualDisplacementStatus(i.DisplacementStatus, validation.NewPath("displacementStatus"))...)
	allErrs = append(allErrs, validateIndividualGender(i.Gender, validation.NewPath("gender"))...)
	allErrs = append(allErrs, validateIndividualEmail(i.Email, validation.NewPath("email"))...)
	return allErrs
}

func validateBirthDate(birthDate *time.Time, path *validation.Path) validation.ErrorList {
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

func validateIndividualDisplacementStatus(ds string, path *validation.Path) validation.ErrorList {
	switch ds {
	case "idp", "refugee", "host_community":
		return validation.ErrorList{}
	case "":
		return validation.ErrorList{validation.Required(path, "displacement status is required")}
	default:
		return validation.ErrorList{validation.NotSupported(path, ds, []string{
			"idp",
			"refugee",
			"host_community",
		})}
	}
}

func validateIndividualGender(gender string, path *validation.Path) validation.ErrorList {
	switch gender {
	case "male", "female", "other", "prefers_not_to_say":
		return validation.ErrorList{}
	case "":
		return validation.ErrorList{validation.Required(path, "gender is required")}
	default:
		return validation.ErrorList{validation.NotSupported(path, gender, []string{
			"male",
			"female",
			"other",
			"prefers_not_to_say",
		})}
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
