package validation

import (
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateListIndividualsOptions(t *testing.T) {
	fromDate := time.Now().AddDate(-10, 0, 0)
	toDate := time.Now().AddDate(-5, 0, 0)
	tests := []struct {
		name string
		opts api.GetAllOptions
		want validation.ErrorList
	}{
		{
			name: "valid",
			opts: api.GetAllOptions{
				Take:                 10,
				Skip:                 10,
				Genders:              []string{"female", "male", "other", "prefers_not_to_say"},
				DisplacementStatuses: []string{"idp", "refugee", "host_community"},
				BirthDateFrom:        &fromDate,
				BirthDateTo:          &toDate,
				CountryID:            "countryId",
			},
			want: validation.ErrorList{},
		}, {
			name: "invalid skip",
			opts: api.GetAllOptions{
				CountryID: "countryId",
				Skip:      -1,
			},
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("skip"), -1, "must be greater than or equal to 0"),
			},
		}, {
			name: "invalid take",
			opts: api.GetAllOptions{
				CountryID: "countryId",
				Take:      -1,
			},
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("take"), -1, "must be greater than or equal to 0"),
			},
		}, {
			name: "invalid gender",
			opts: api.GetAllOptions{
				CountryID: "countryId",
				Genders:   []string{"invalid"},
			},
			want: validation.ErrorList{
				validation.NotSupported(validation.NewPath("genders").Index(0), "invalid", allowedGenders.Items()),
			},
		}, {
			name: "duplicate gender",
			opts: api.GetAllOptions{
				CountryID: "countryId",
				Genders:   []string{"male", "male"},
			},
			want: validation.ErrorList{
				validation.Duplicate(validation.NewPath("genders").Index(1), "male", "gender specified multiple times in options"),
			},
		}, {
			name: "invalid displacement status",
			opts: api.GetAllOptions{
				CountryID:            "countryId",
				DisplacementStatuses: []string{"invalid"},
			},
			want: validation.ErrorList{
				validation.NotSupported(validation.NewPath("displacementStatuses").Index(0), "invalid", allowedDisplacementStatuses.Items()),
			},
		}, {
			name: "duplicate displacement status",
			opts: api.GetAllOptions{
				CountryID:            "countryId",
				DisplacementStatuses: []string{"refugee", "refugee"},
			},
			want: validation.ErrorList{
				validation.Duplicate(validation.NewPath("displacementStatuses").Index(1), "refugee", "displacement status specified multiple times in options"),
			},
		}, {
			name: "from birthdate after to birthdate",
			opts: api.GetAllOptions{
				CountryID:     "countryId",
				BirthDateFrom: &toDate,
				BirthDateTo:   &fromDate,
			},
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("birthDateFrom"), &toDate, "birthDateFrom must be before birthDateTo"),
			},
		}, {
			name: "missing country ID",
			opts: api.GetAllOptions{
				CountryID: "",
			},
			want: validation.ErrorList{
				validation.Required(validation.NewPath("countryId"), "country id is required"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateListIndividualsOptions(&tt.opts)
			assert.Equal(t, tt.want, got)
		})
	}
}
