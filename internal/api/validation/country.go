package validation

import (
	"regexp"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/validation"
)

func ValidateCountry(country *api.Country) validation.ErrorList {
	return validateCountry(nil, country)
}

func ValidateCountryList(countryList *api.CountryList) validation.ErrorList {
	allErrs := validation.ErrorList{}
	itemsPath := validation.NewPath("items")
	for i, country := range countryList.Items {
		allErrs = append(allErrs, validateCountry(itemsPath.Index(i), country)...)
	}
	return allErrs
}

func validateCountry(path *validation.Path, country *api.Country) validation.ErrorList {
	allErrs := validation.ErrorList{}
	allErrs = append(allErrs, validateCountryName(country.Name, path.Child("name"))...)
	allErrs = append(allErrs, validateCountryCode(country.Code, path.Child("code"))...)
	allErrs = append(allErrs, validateCountryNrcOrganisations(country.NrcOrganisations, path.Child("nrcOrganisations"))...)
	return allErrs
}

var countryNameMaxLength = 255
var countryNameMinLength = 2
var allowedCountryNameChars = map[rune]bool{}

func init() {
	for _, r := range `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789èéêëàâäôöûüçÈÉÊËÀÂÄÔÖÛÜÇ_-'"!@#$%^&*()[]{}|;:,./<>? ` {
		allowedCountryNameChars[r] = true
	}
}

var countryCodeMaxLength = 255
var countryCodeMinLength = 2
var countryCodePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

var countryNrcOrganisationMaxLength = 255
var countryNrcOrganisationMinLength = 1
var countryNrcOrganisationPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+(?: [a-zA-Z0-9_-]+)*$`)

func validateCountryName(name string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if name == "" {
		allErrs = append(allErrs, validation.Required(path, "country name is required"))
	} else if len(name) > countryNameMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, name, countryNameMaxLength))
	} else if len(name) < countryNameMinLength {
		allErrs = append(allErrs, validation.TooShortMinLength(path, name, countryNameMinLength))
	} else if !countryNameContainsOnlyValidChars(name) {
		allErrs = append(allErrs, validation.Invalid(path, name, "country name can only contain letters or spaces"))
	}
	return allErrs
}

func countryNameContainsOnlyValidChars(name string) bool {
	for _, r := range name {
		if !allowedCountryNameChars[r] {
			return false
		}
	}
	return true
}

func validateCountryCode(code string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if code == "" {
		allErrs = append(allErrs, validation.Required(path, "country code is required"))
	} else if len(code) > countryCodeMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, code, countryCodeMaxLength))
	} else if len(code) < countryCodeMinLength {
		allErrs = append(allErrs, validation.TooShortMinLength(path, code, countryCodeMinLength))
	} else if !countryCodePattern.MatchString(code) {
		allErrs = append(allErrs, validation.Invalid(path, code, "country code can only contain letters, numbers, underscores and hyphens"))
	}
	return allErrs
}

func validateCountryNrcOrganisations(nrcOrganisations string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if nrcOrganisations == "" {
		allErrs = append(allErrs, validation.Required(path, "nrc organisation is required"))
	}
	for _, org := range strings.Split(nrcOrganisations, ",") {
		if len(org) > countryNrcOrganisationMaxLength {
			allErrs = append(allErrs, validation.TooLongMaxLength(path, org, countryNrcOrganisationMaxLength))
		} else if len(org) < countryNrcOrganisationMinLength {
			allErrs = append(allErrs, validation.TooShortMinLength(path, org, countryNrcOrganisationMinLength))
		} else if !countryNrcOrganisationPattern.MatchString(org) {
			allErrs = append(allErrs, validation.Invalid(path, org, "nrc organisation is invalid"))
		}
	}
	return allErrs
}
