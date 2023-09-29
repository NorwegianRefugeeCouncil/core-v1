package middleware

import (
	"reflect"
	"testing"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils"
)

var (
	country1 = api.Country{
		ID:           "1",
		ReadGroup:    "nrc-country-1-read",
		WriteGroup:   "nrc-country-1-write",
	}
	country2 = api.Country{
		ID:           "2",
		ReadGroup:    "nrc-country-2-read",
		WriteGroup:   "nrc-country-2-write",
	}
	country3 = api.Country{
		ID:           "3",
		ReadGroup:    "nrc-country-1-read",
		WriteGroup:   "nrc-country-1-write",
	}
)

func Test_parsePermissions(t *testing.T) {
	jwtGroups := utils.JwtGroupOptions{
		GlobalAdmin: "global-admin",
	}
	type args struct {
		allCountries    []*api.Country
		jwtGroups       utils.JwtGroupOptions
		userGroups      []string
	}
	tests := []struct {
		name string
		args args
		want *ParsedPermissions
	}{
		{
			name: "global admin, no countries defined",
			args: args{
				allCountries: []*api.Country{},
				jwtGroups: jwtGroups,
				userGroups: []string{
					jwtGroups.GlobalAdmin,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: true,
				CountryPermissions: auth.CountryPermissions{},
			},
		},
		{
			name: "global admin, with countries defined",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					jwtGroups.GlobalAdmin,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: true,
				CountryPermissions: auth.CountryPermissions{},
			},
		},
		{
			name: "global admin, with countries defined and read access to one",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					jwtGroups.GlobalAdmin,
					country1.ReadGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: true,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
					),
				},
			},
		},
		{
			name: "single country, no access",
			args: args{
				allCountries: []*api.Country{
					&country1,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{},
			},
		},
		{
			name: "single country, read access",
			args: args{
				allCountries: []*api.Country{
					&country1,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
					),
				},
			},
		},
		{
			name: "single country, write access",
			args: args{
				allCountries: []*api.Country{
					&country1,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionWrite,
					),
				},
			},
		},
		{
			name: "single country, read and write access",
			args: args{
				allCountries: []*api.Country{
					&country1,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
					country1.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
						auth.PermissionWrite,
					),
				},
			},
		},
		{
			name: "multiple countries, no access",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{},
			},
		},
		{
			name: "multiple countries, read access to one",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
					),
				},
			},
		},
		{
			name: "multiple countries, write access to one",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionWrite,
					),
				},
			},
		},
		{
			name: "multiple countries, read and write access to one",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
					country1.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
						auth.PermissionWrite,
					),
				},
			},
		},
		{
			name: "multiple countries, read access to multiple",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
					country2.ReadGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
					),
					country2.ID: containers.NewSet(
						auth.PermissionRead,
					),
				},
			},
		},
		{
			name: "multiple countries, write access to multiple",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.WriteGroup,
					country2.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionWrite,
					),
					country2.ID: containers.NewSet(
						auth.PermissionWrite,
					),
				},
			},
		},
		{
			name: "multiple countries, read and write access to multiple",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
					country1.WriteGroup,
					country2.ReadGroup,
					country2.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
						auth.PermissionWrite,
					),
					country2.ID: containers.NewSet(
						auth.PermissionRead,
						auth.PermissionWrite,
					),
				},
			},
		},
		{
			name: "multiple countries, read access to one and write access to another",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
					country2.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
					),
					country2.ID: containers.NewSet(
						auth.PermissionWrite,
					),
				},
			},
		},
		{
			name: "multiple countries, read and write access to one and write access to another",
			args: args{
				allCountries: []*api.Country{
					&country1,
					&country2,
				},
				jwtGroups: jwtGroups,
				userGroups: []string{
					country1.ReadGroup,
					country1.WriteGroup,
					country2.WriteGroup,
				},
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CountryPermissions: auth.CountryPermissions{
					country1.ID: containers.NewSet(
						auth.PermissionRead,
						auth.PermissionWrite,
					),
					country2.ID: containers.NewSet(
						auth.PermissionWrite,
					),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePermissions(tt.args.allCountries, tt.args.jwtGroups, tt.args.userGroups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
