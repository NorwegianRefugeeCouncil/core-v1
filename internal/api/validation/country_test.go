package validation

import (
	"strings"
	"testing"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

func ValidCountry() *api.Country_Builder {
	return api.New_Country_Builder().
		WithID("id").
		WithCode("code").
		WithJwtGroup("jwt_group").
		WithName("name")
}

func TestValidateCountry(t *testing.T) {
	namePath := validation.NewPath("name")
	codePath := validation.NewPath("code")
	jwtGroupPath := validation.NewPath("jwtGroup")
	weirdString := string([]byte{0x7f, 0x7f})
	tests := []struct {
		name    string
		country *api.Country
		want    validation.ErrorList
	}{
		{
			name:    "valid",
			country: ValidCountry().Build(),
			want:    validation.ErrorList{},
		}, {
			name:    "invalid name",
			country: ValidCountry().WithName(weirdString).Build(),
			want:    validation.ErrorList{validation.Invalid(namePath, weirdString, "country name can only contain letters or spaces")},
		}, {
			name:    "missing name",
			country: ValidCountry().WithName("").Build(),
			want:    validation.ErrorList{validation.Required(namePath, "country name is required")},
		}, {
			name:    "name too long",
			country: ValidCountry().WithName(longString(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(namePath, longString(256), 255)},
		}, {
			name:    "name too short",
			country: ValidCountry().WithName("a").Build(),
			want:    validation.ErrorList{validation.TooShortMinLength(namePath, "a", 2)},
		}, {
			name:    "missing code",
			country: ValidCountry().WithCode("").Build(),
			want:    validation.ErrorList{validation.Required(codePath, "country code is required")},
		}, {
			name:    "invalid code",
			country: ValidCountry().WithCode("!!").Build(),
			want:    validation.ErrorList{validation.Invalid(codePath, "!!", "country code can only contain letters, numbers, underscores and hyphens")},
		}, {
			name:    "code too short",
			country: ValidCountry().WithCode("a").Build(),
			want:    validation.ErrorList{validation.TooShortMinLength(codePath, "a", 2)},
		}, {
			name:    "code too long",
			country: ValidCountry().WithCode(longString(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(codePath, longString(256), 255)},
		}, {
			name:    "missing jwt group",
			country: ValidCountry().WithJwtGroup("").Build(),
			want:    validation.ErrorList{validation.Required(jwtGroupPath, "jwt group is required")},
		}, {
			name:    "invalid jwt group",
			country: ValidCountry().WithJwtGroup("!!").Build(),
			want:    validation.ErrorList{validation.Invalid(jwtGroupPath, "!!", "jwt group is invalid")},
		}, {
			name:    "jwt group too long",
			country: ValidCountry().WithJwtGroup(longString(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(jwtGroupPath, longString(256), 255)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateCountry(tt.country)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValidateCountryList(t *testing.T) {
	tests := []struct {
		name        string
		countryList *api.CountryList
		want        validation.ErrorList
	}{
		{
			name:        "empty",
			countryList: &api.CountryList{},
			want:        validation.ErrorList{},
		}, {
			name: "valid",
			countryList: &api.CountryList{
				Items: []*api.Country{
					ValidCountry().Build(),
				},
			},
			want: validation.ErrorList{},
		}, {
			name: "invalid",
			countryList: &api.CountryList{
				Items: []*api.Country{
					ValidCountry().WithName("").Build(),
				},
			},
			want: validation.ErrorList{
				validation.Required(validation.NewPath("items[0].name"), "country name is required"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateCountryList(tt.countryList)
			assert.Equal(t, tt.want, got)
		})
	}
}

func longString(n int) string {
	return strings.Repeat("a", n)
}
