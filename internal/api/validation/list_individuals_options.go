package validation

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/validation"
)

func ValidateListIndividualsOptions(opts *api.GetAllOptions) validation.ErrorList {
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

func validateListIndividualsOptionsGenders(genders []api.Gender, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	seenGenders := containers.NewSet[api.Gender]()
	for i, g := range genders {
		if !allowedGenders.Contains(g) {
			genderStrings := make([]string, len(allowedGenders))
			for i, g := range allowedGenders.Items() {
				genderStrings[i] = string(g)
			}
			allErrs = append(allErrs, validation.NotSupported(p.Index(i), g, genderStrings))
		} else {
			if seenGenders.Contains(g) {
				allErrs = append(allErrs, validation.Duplicate(p.Index(i), g, "gender specified multiple times in options"))
			}
			seenGenders.Add(g)
		}
	}
	return allErrs
}

func validateListIndividualsOptionsDisplacementStatuses(displacementStatuses []api.DisplacementStatus, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	seenDisplacementStatuses := containers.NewSet[api.DisplacementStatus]()
	for i, g := range displacementStatuses {
		if !allowedDisplacementStatuses.Contains(g) {
			displacementStatusStrings := make([]string, len(allowedDisplacementStatuses))
			for i, g := range allowedDisplacementStatuses.Items() {
				displacementStatusStrings[i] = string(g)
			}
			allErrs = append(allErrs, validation.NotSupported(p.Index(i), g, displacementStatusStrings))
		} else {
			if seenDisplacementStatuses.Contains(g) {
				allErrs = append(allErrs, validation.Duplicate(p.Index(i), g, "displacement status specified multiple times in options"))
			}
			seenDisplacementStatuses.Add(g)
		}
	}
	return allErrs
}
