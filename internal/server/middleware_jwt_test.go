package server

import (
	"reflect"
	"testing"
)

func Test_unmarshapTokenClaims(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
		want    TokenClaims
		wantErr bool
	}{
		{
			name:    "empty",
			payload: []byte{},
			want:    TokenClaims{},
			wantErr: true,
		},
		{
			name:    "invalid",
			payload: []byte("invalid"),
			want:    TokenClaims{},
			wantErr: true,
		},
		{
			name:    "valid",
			payload: []byte(`{"iss":"issuer","sub":"subject","email":"email","groups":["group1","group2"]}`),
			want: TokenClaims{
				Iss:    "issuer",
				Sub:    "subject",
				Email:  "email",
				Groups: []string{"group1", "group2"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalTokenClaims(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalTokenClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalTokenClaims() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateTokenClaims(t *testing.T) {
	tests := []struct {
		name    string
		claims  TokenClaims
		wantErr bool
	}{
		{
			name: "valid",
			claims: TokenClaims{
				Iss:    "issuer",
				Sub:    "subject",
				Email:  "email",
				Groups: []string{"group1", "group2"},
			},
			wantErr: false,
		},
		{
			name: "missing issuer",
			claims: TokenClaims{
				Sub:    "subject",
				Email:  "email",
				Groups: []string{"group1", "group2"},
			},
			wantErr: true,
		},
		{
			name: "missing subject",
			claims: TokenClaims{
				Iss:    "issuer",
				Email:  "email",
				Groups: []string{"group1", "group2"},
			},
			wantErr: true,
		},
		{
			name: "missing email",
			claims: TokenClaims{
				Iss:    "issuer",
				Sub:    "subject",
				Groups: []string{"group1", "group2"},
			},
			wantErr: true,
		},
		{
			name: "missing groups",
			claims: TokenClaims{
				Iss:   "issuer",
				Sub:   "subject",
				Email: "email",
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
