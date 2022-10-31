package validation

import (
	"regexp"

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
	allErrs = append(allErrs, validateCountryJWTGroup(country.JwtGroup, path.Child("jwtGroup"))...)
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

var countryJwtGroupMaxLength = 255
var countryJwtGroupMinLength = 1
var countryJwtGroupPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+(?: [a-zA-Z0-9_-]+)*$`)

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

func validateCountryJWTGroup(jwtGroup string, path *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if jwtGroup == "" {
		allErrs = append(allErrs, validation.Required(path, "jwt group is required"))
	} else if len(jwtGroup) > countryJwtGroupMaxLength {
		allErrs = append(allErrs, validation.TooLongMaxLength(path, jwtGroup, countryJwtGroupMaxLength))
	} else if len(jwtGroup) < countryJwtGroupMinLength {
		allErrs = append(allErrs, validation.TooShortMinLength(path, jwtGroup, countryJwtGroupMinLength))
	} else if !countryJwtGroupPattern.MatchString(jwtGroup) {
		allErrs = append(allErrs, validation.Invalid(path, jwtGroup, "jwt group is invalid"))
	}
	return allErrs
}
