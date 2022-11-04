package validation

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/validation"
)

func ValidateListIndividualsOptions(opts *api.ListIndividualsOptions) validation.ErrorList {
	allErrs := validation.ErrorList{}
	allErrs = append(allErrs, validateListIndividualsCountryID(opts.CountryID, validation.NewPath("countryId"))...)
	allErrs = append(allErrs, validateListIndividualsOptionsSkip(opts.Skip, validation.NewPath("skip"))...)
	allErrs = append(allErrs, validateListIndividualsOptionsTake(opts.Take, validation.NewPath("take"))...)
	allErrs = append(allErrs, validateListIndividualsOptionsGenders(opts.Genders, validation.NewPath("genders"))...)
	allErrs = append(allErrs, validateListIndividualsOptionsDisplacementStatuses(opts.DisplacementStatuses, validation.NewPath("displacementStatuses"))...)
	if opts.BirthDateFrom != nil && opts.BirthDateTo != nil && opts.BirthDateFrom.After(*opts.BirthDateTo) {
		allErrs = append(allErrs, validation.Invalid(validation.NewPath("birthDateFrom"), opts.BirthDateFrom, "birthDateFrom must be before birthDateTo"))
	}
	return allErrs
}

func validateListIndividualsCountryID(countryID string, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if countryID == "" {
		allErrs = append(allErrs, validation.Required(p, "country id is required"))
	}
	return allErrs
}

func validateListIndividualsOptionsSkip(skip int, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if skip < 0 {
		allErrs = append(allErrs, validation.Invalid(p, skip, "must be greater than or equal to 0"))
	}
	return allErrs
}

func validateListIndividualsOptionsTake(take int, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	if take < 0 {
		allErrs = append(allErrs, validation.Invalid(p, take, "must be greater than or equal to 0"))
	}
	return allErrs
}

func validateListIndividualsOptionsGenders(genders containers.Set[api.Gender], p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	for i, g := range genders.Items() {
		if !allowedGenders.Contains(g) {
			allErrs = append(allErrs, validation.NotSupported(p.Index(i), g, allowedGendersStr))
		}
	}
	return allErrs
}

func validateListIndividualsOptionsDisplacementStatuses(displacementStatuses containers.Set[api.DisplacementStatus], p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	for i, g := range displacementStatuses.Items() {
		if !allowedDisplacementStatuses.Contains(g) {
			allErrs = append(allErrs, validation.NotSupported(p.Index(i), g, allowedDisplacementStatusesStr))
		}
	}
	return allErrs
}
