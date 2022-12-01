package middleware

import (
	"testing"
)

func Test_validateTokenClaims(t *testing.T) {
	tests := []struct {
		name    string
		claims  TokenClaims
		wantErr bool
	}{
		{
			name: "valid",
			claims: TokenClaims{
				Iss:             "issuer",
				Sub:             "subject",
				Email:           "email",
				Groups:          []string{"group1", "group2"},
				NrcOrganisation: "NRC Uganda",
				Exp:             1,
				Iat:             2,
			},
			wantErr: false,
		},
		{
			name: "missing issuer",
			claims: TokenClaims{
				Sub:             "subject",
				Email:           "email",
				Groups:          []string{"group1", "group2"},
				NrcOrganisation: "NRC Uganda",
				Exp:             1,
				Iat:             2,
			},
			wantErr: true,
		},
		{
			name: "missing subject",
			claims: TokenClaims{
				Iss:             "issuer",
				Email:           "email",
				Groups:          []string{"group1", "group2"},
				NrcOrganisation: "NRC Uganda",
				Exp:             1,
				Iat:             2,
			},
			wantErr: true,
		},
		{
			name: "missing email",
			claims: TokenClaims{
				Iss:             "issuer",
				Sub:             "subject",
				Groups:          []string{"group1", "group2"},
				NrcOrganisation: "NRC Uganda",
				Exp:             1,
				Iat:             2,
			},
			wantErr: true,
		},
		{
			name: "missing groups",
			claims: TokenClaims{
				Iss:             "issuer",
				Sub:             "subject",
				Email:           "email",
				NrcOrganisation: "NRC Uganda",
				Exp:             1,
				Iat:             2,
			},
			wantErr: true,
		},
		{
			name: "missing nrcOrganisation",
			claims: TokenClaims{
				Iss:   "issuer",
				Sub:   "subject",
				Email: "email",
				Exp:   1,
				Iat:   2,
			},
			wantErr: true,
		},
		{
			name: "missing expiration",
			claims: TokenClaims{
				Iss:             "issuer",
				Sub:             "subject",
				Email:           "email",
				Groups:          []string{"group1", "group2"},
				NrcOrganisation: "NRC Uganda",
				Iat:             2,
			},
			wantErr: true,
		}, {
			name: "missing issuedAt",
			claims: TokenClaims{
				Iss:             "issuer",
				Sub:             "subject",
				Email:           "email",
				Groups:          []string{"group1", "group2"},
				NrcOrganisation: "NRC Uganda",
				Exp:             2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTokenClaims(tt.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateTokenClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
