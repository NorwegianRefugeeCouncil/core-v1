package validation

import (
	"strings"
	"testing"

	"github.com/nrc-no/notcore/internal/api"

	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

type CountryBuilder struct {
	country *api.Country
}

func ValidCountry() *CountryBuilder {
	return &CountryBuilder{
		country: &api.Country{
			ID:               "id",
			Code:             "code",
			Name:             "name",
			ReadGroup: 			  "read_group",
			WriteGroup: 		  "write_group",
		},
	}
}

func (b *CountryBuilder) Build() *api.Country {
	return b.country
}

func (b *CountryBuilder) WithName(name string) *CountryBuilder {
	b.country.Name = name
	return b
}

func (b *CountryBuilder) WithCode(code string) *CountryBuilder {
	b.country.Code = code
	return b
}

func (b *CountryBuilder) WithReadGroup(readGroup string) *CountryBuilder {
	b.country.ReadGroup = readGroup
	return b
}

func (b *CountryBuilder) WithWriteGroup(writeGroup string) *CountryBuilder {
	b.country.WriteGroup = writeGroup
	return b
}

func (b *CountryBuilder) WithID(id string) *CountryBuilder {
	b.country.ID = id
	return b
}

func TestValidateCountry(t *testing.T) {
	namePath := validation.NewPath("name")
	codePath := validation.NewPath("code")
	readGroupPath := validation.NewPath("readGroup")
	writeGroupPath := validation.NewPath("writeGroup")
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
			country: ValidCountry().WithName(bigstr(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(namePath, bigstr(256), 255)},
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
			country: ValidCountry().WithCode(bigstr(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(codePath, bigstr(256), 255)},
		}, {
			name: "missing read group",
			country: ValidCountry().
				WithReadGroup("").
				Build(),
			want: validation.ErrorList{validation.Required(readGroupPath, "group is required")},
		}, {
			name: "invalid read group",
			country: ValidCountry().
				WithReadGroup("!!").
				Build(),
			want: validation.ErrorList{validation.Invalid(readGroupPath, "!!", "group is invalid")},
		}, {
			name: "read group too long",
			country: ValidCountry().
				WithReadGroup(bigstr(256)).
				Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(readGroupPath, bigstr(256), 255)},
		}, {
			name: "missing write group",
			country: ValidCountry().
				WithWriteGroup("").
				Build(),
			want: validation.ErrorList{validation.Required(writeGroupPath, "group is required")},
		}, {
			name: "invalid write group",
			country: ValidCountry().
				WithWriteGroup("!!").
				Build(),
			want: validation.ErrorList{validation.Invalid(writeGroupPath, "!!", "group is invalid")},
		}, {
			name: "write group too long",
			country: ValidCountry().
				WithWriteGroup(bigstr(256)).
				Build(),
			want: validation.ErrorList{validation.TooLongMaxLength(writeGroupPath, bigstr(256), 255)},
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

func bigstr(n int) string {
	return strings.Repeat("a", n)
}
