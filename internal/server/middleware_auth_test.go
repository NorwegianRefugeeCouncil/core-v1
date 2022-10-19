package server

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Test_unmarshalJsonClaims(t *testing.T) {
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
			got, err := unmarshalJsonClaims(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalJsonClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalJsonClaims() got = %v, want %v", got, tt.want)
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

func Test_extractJwtTokenClaims(t *testing.T) {

	var createJwt = func() string {
		t := jwt.New(jwt.SigningMethodHS256)
		t.Claims = jwt.MapClaims{
			"iss":    "issuer",
			"sub":    "subject",
			"email":  "email",
			"groups": []string{"group1", "group2"},
			"aud":    "audience",
			"exp":    time.Now().Add(time.Hour).Unix(),
			"iat":    time.Now().Unix(),
		}
		token, _ := t.SignedString([]byte("secret"))
		return token
	}

	tests := []struct {
		name    string
		payload string
		want    TokenClaims
		wantErr bool
	}{
		{
			name:    "empty",
			payload: "",
			want:    TokenClaims{},
			wantErr: true,
		}, {
			name:    "invalid",
			payload: "invalid",
			want:    TokenClaims{},
			wantErr: true,
		}, {
			name:    "valid",
			payload: createJwt(),
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
			got, err := extractJwtTokenClaims(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractJwtTokenClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractJwtTokenClaims() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractBearerTokenClaims(t *testing.T) {

	var createJwt = func() string {
		t := jwt.New(jwt.SigningMethodHS256)
		t.Claims = jwt.MapClaims{
			"iss":    "issuer",
			"sub":    "subject",
			"email":  "email",
			"groups": []string{"group1", "group2"},
			"aud":    "audience",
			"exp":    time.Now().Add(time.Hour).Unix(),
			"iat":    time.Now().Unix(),
		}
		token, _ := t.SignedString([]byte("secret"))
		return token
	}

	tests := []struct {
		name    string
		payload string
		want    TokenClaims
		wantErr bool
	}{
		{
			name:    "empty",
			payload: "",
			want:    TokenClaims{},
			wantErr: true,
		}, {
			name:    "invalid",
			payload: "invalid",
			want:    TokenClaims{},
			wantErr: true,
		}, {
			name:    "valid",
			payload: fmt.Sprintf("Bearer %s", createJwt()),
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
			got, err := extractBearerTokenClaims(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractBearerTokenClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractBearerTokenClaims() got = %v, want %v", got, tt.want)
			}
		})
	}
}
