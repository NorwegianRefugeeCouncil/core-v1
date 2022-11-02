package validation

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/validation"
)

func ValidateListIndividualsOptions(opts *api.GetAllOptions) validation.ErrorList {
	allErrs := validation.ErrorList{}
	allErrs = append(allErrs, validateListIndividualsCountryID(opts.CountryID, validation.NewPath(api.GetAllOptions_Field_CountryID.String()))...)
	allErrs = append(allErrs, validateListIndividualsOptionsSkip(opts.Skip, validation.NewPath(api.GetAllOptions_Field_Skip.String()))...)
	allErrs = append(allErrs, validateListIndividualsOptionsTake(opts.Take, validation.NewPath(api.GetAllOptions_Field_Take.String()))...)
	allErrs = append(allErrs, validateListIndividualsOptionsGenders(opts.Genders, validation.NewPath(api.GetAllOptions_Field_Genders.String()))...)
	allErrs = append(allErrs, validateListIndividualsOptionsDisplacementStatuses(opts.DisplacementStatuses, validation.NewPath(api.GetAllOptions_Field_DisplacementStatuses.String()))...)
	if opts.BirthDateFrom != nil && opts.BirthDateTo != nil && opts.BirthDateFrom.After(*opts.BirthDateTo) {
		allErrs = append(allErrs, validation.Invalid(validation.NewPath(api.GetAllOptions_Field_BirthDateFrom.String()), opts.BirthDateFrom, "birthDateFrom must be before birthDateTo"))
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

func validateListIndividualsOptionsGenders(genders []string, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	seenGenders := containers.NewStringSet()
	for i, g := range genders {
		if !allowedGenders.Contains(g) {
			allErrs = append(allErrs, validation.NotSupported(p.Index(i), g, allowedGenders.Items()))
		} else {
			if seenGenders.Contains(g) {
				allErrs = append(allErrs, validation.Duplicate(p.Index(i), g, "gender specified multiple times in options"))
			}
			seenGenders.Add(g)
		}
	}
	return allErrs
}

func validateListIndividualsOptionsDisplacementStatuses(displacementStatuses []string, p *validation.Path) validation.ErrorList {
	allErrs := validation.ErrorList{}
	seenDisplacementStatuses := containers.NewStringSet()
	for i, g := range displacementStatuses {
		if !allowedDisplacementStatuses.Contains(g) {
			allErrs = append(allErrs, validation.NotSupported(p.Index(i), g, allowedDisplacementStatuses.Items()))
		} else {
			if seenDisplacementStatuses.Contains(g) {
				allErrs = append(allErrs, validation.Duplicate(p.Index(i), g, "displacement status specified multiple times in options"))
			}
			seenDisplacementStatuses.Add(g)
		}
	}
	return allErrs
}
