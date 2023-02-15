package validation

import (
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateListIndividualsOptions(t *testing.T) {
	fromDate := time.Now().AddDate(-10, 0, 0)
	toDate := time.Now().AddDate(-5, 0, 0)
	tests := []struct {
		name string
		opts api.ListIndividualsOptions
		want validation.ErrorList
	}{
		{
			name: "valid",
			opts: api.ListIndividualsOptions{
				Take:                 10,
				Skip:                 10,
				Sexes:                containers.NewSet[enumTypes.Sex](enumTypes.SexFemale, enumTypes.SexMale, enumTypes.SexOther, enumTypes.SexPreferNotToSay),
				DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus](enumTypes.DisplacementStatusIDP, enumTypes.DisplacementStatusRefugee, enumTypes.DisplacementStatusHostCommunity),
				BirthDateFrom:        &fromDate,
				BirthDateTo:          &toDate,
				CountryID:            "countryId",
			},
			want: validation.ErrorList{},
		}, {
			name: "invalid skip",
			opts: api.ListIndividualsOptions{
				CountryID: "countryId",
				Skip:      -1,
			},
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("skip"), -1, "must be greater than or equal to 0"),
			},
		}, {
			name: "invalid take",
			opts: api.ListIndividualsOptions{
				CountryID: "countryId",
				Take:      -1,
			},
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("take"), -1, "must be greater than or equal to 0"),
			},
		}, {
			name: "invalid sex",
			opts: api.ListIndividualsOptions{
				CountryID: "countryId",
				Sexes:     containers.NewSet[enumTypes.Sex]("invalid"),
			},
			want: validation.ErrorList{
				validation.NotSupported(validation.NewPath("sexes").Index(0), enumTypes.Sex("invalid"), allowedSexesStr),
			},
		}, {
			name: "invalid displacement status",
			opts: api.ListIndividualsOptions{
				CountryID:            "countryId",
				DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus]("invalid"),
			},
			want: validation.ErrorList{
				validation.NotSupported(validation.NewPath("displacementStatuses").Index(0), enumTypes.DisplacementStatus("invalid"), allowedDisplacementStatusesStr),
			},
		}, {
			name: "from birthdate after to birthdate",
			opts: api.ListIndividualsOptions{
				CountryID:     "countryId",
				BirthDateFrom: &toDate,
				BirthDateTo:   &fromDate,
			},
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("birthDateFrom"), &toDate, "birthDateFrom must be before birthDateTo"),
			},
		}, {
			name: "missing country ID",
			opts: api.ListIndividualsOptions{
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
