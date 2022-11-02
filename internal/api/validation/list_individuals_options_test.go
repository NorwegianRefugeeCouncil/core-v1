package validation

import (
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

func ValidIndividualListOptions() *api.GetAllOptions_Builder {
	fromDate := time.Now().AddDate(-10, 0, 0)
	toDate := time.Now().AddDate(-5, 0, 0)
	return api.New_GetAllOptions_Builder().
		WithTake(10).
		WithSkip(10).
		WithGenders([]string{"female", "male", "other", "prefers_not_to_say"}).
		WithDisplacementStatuses([]string{"idp", "refugee", "host_community"}).
		WithBirthDateFrom(&fromDate).
		WithBirthDateTo(&toDate).
		WithCountryID("countryId")
}

func TestValidateListIndividualsOptions(t *testing.T) {
	fromDate := time.Now().AddDate(-10, 0, 0)
	toDate := time.Now().AddDate(-5, 0, 0)
	tests := []struct {
		name string
		opts *api.GetAllOptions
		want validation.ErrorList
	}{
		{
			name: "valid",
			opts: ValidIndividualListOptions().Build(),
			want: validation.ErrorList{},
		}, {
			name: "invalid skip",
			opts: ValidIndividualListOptions().WithSkip(-1).Build(),
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("skip"), -1, "must be greater than or equal to 0"),
			},
		}, {
			name: "invalid take",
			opts: ValidIndividualListOptions().WithTake(-1).Build(),
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("take"), -1, "must be greater than or equal to 0"),
			},
		}, {
			name: "invalid gender",
			opts: ValidIndividualListOptions().WithGenders([]string{"invalid"}).Build(),
			want: validation.ErrorList{
				validation.NotSupported(validation.NewPath("genders").Index(0), "invalid", allowedGenders.Items()),
			},
		}, {
			name: "duplicate gender",
			opts: ValidIndividualListOptions().WithGenders([]string{"male", "male"}).Build(),
			want: validation.ErrorList{
				validation.Duplicate(validation.NewPath("genders").Index(1), "male", "gender specified multiple times in options"),
			},
		}, {
			name: "invalid displacement status",
			opts: ValidIndividualListOptions().WithDisplacementStatuses([]string{"invalid"}).Build(),
			want: validation.ErrorList{
				validation.NotSupported(validation.NewPath("displacementStatuses").Index(0), "invalid", allowedDisplacementStatuses.Items()),
			},
		}, {
			name: "duplicate displacement status",
			opts: ValidIndividualListOptions().WithDisplacementStatuses([]string{"refugee", "refugee"}).Build(),
			want: validation.ErrorList{
				validation.Duplicate(validation.NewPath("displacementStatuses").Index(1), "refugee", "displacement status specified multiple times in options"),
			},
		}, {
			name: "from birthdate after to birthdate",
			opts: ValidIndividualListOptions().WithBirthDateFrom(&toDate).WithBirthDateTo(&fromDate).Build(),
			want: validation.ErrorList{
				validation.Invalid(validation.NewPath("birthDateFrom"), &toDate, "birthDateFrom must be before birthDateTo"),
			},
		}, {
			name: "missing country ID",
			opts: ValidIndividualListOptions().WithCountryID("").Build(),
			want: validation.ErrorList{
				validation.Required(validation.NewPath("countryId"), "country id is required"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateListIndividualsOptions(tt.opts)
			assert.Equal(t, tt.want, got)
		})
	}
}
