package db

import (
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func Test_newGetAllIndividualsSQLQuery(t *testing.T) {
	someDate, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	zeroTime := time.Time{}
	const defaultQuery = `SELECT * FROM individual_registrations WHERE deleted_at IS NULL ORDER BY created_at`
	tests := []struct {
		name     string
		args     api.ListIndividualsOptions
		wantSql  string
		wantArgs []interface{}
	}{
		{
			name:    "empty",
			args:    api.ListIndividualsOptions{},
			wantSql: defaultQuery,
		}, {
			name:     "ids",
			args:     api.ListIndividualsOptions{IDs: containers.NewStringSet("1", "2")},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND id IN ($1,$2) ORDER BY created_at`,
			wantArgs: []interface{}{"1", "2"},
		}, {
			name:     "full name",
			args:     api.ListIndividualsOptions{FullName: "John"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (full_name ILIKE $1 OR preferred_name ILIKE $2) ORDER BY created_at`,
			wantArgs: []interface{}{"%John%", "%John%"},
		}, {
			name:     "address",
			args:     api.ListIndividualsOptions{Address: "123 Main St"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND address ILIKE $1 ORDER BY created_at`,
			wantArgs: []interface{}{"%123 Main St%"},
		}, {
			name:     "genders",
			args:     api.ListIndividualsOptions{Genders: containers.NewSet[api.Gender](api.GenderMale, api.GenderFemale)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND gender IN ($1,$2) ORDER BY created_at`,
			wantArgs: []interface{}{"female", "male"},
		}, {
			name:     "all genders",
			args:     api.ListIndividualsOptions{Genders: api.AllGenders()},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND gender IN ($1,$2,$3,$4) ORDER BY created_at`,
			wantArgs: []interface{}{"female", "male", "other", "prefers_not_to_say"},
		}, {
			name:     "birth date from",
			args:     api.ListIndividualsOptions{BirthDateFrom: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND birth_date >= $1 ORDER BY created_at`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:    "zero birth date from",
			args:    api.ListIndividualsOptions{BirthDateFrom: &zeroTime},
			wantSql: defaultQuery,
		}, {
			name:     "birth date to",
			args:     api.ListIndividualsOptions{BirthDateTo: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND birth_date <= $1 ORDER BY created_at`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:    "zero birth date to",
			args:    api.ListIndividualsOptions{BirthDateFrom: &zeroTime},
			wantSql: defaultQuery,
		}, {
			name:     "phone number",
			args:     api.ListIndividualsOptions{PhoneNumber: "1234567890"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (normalized_phone_number_1 ILIKE $1 OR normalized_phone_number_2 ILIKE $1 OR normalized_phone_number_3 ILIKE $1) ORDER BY created_at`,
			wantArgs: []interface{}{"%1234567890%"},
		}, {
			name:     "email",
			args:     api.ListIndividualsOptions{Email: "info@email.com"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (email_1 = $1 OR email_2 = $1 OR email_3 = $1) ORDER BY created_at`,
			wantArgs: []interface{}{"info@email.com"},
		}, {
			name:     "presents protection concerns",
			args:     api.ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND presents_protection_concerns = $1 ORDER BY created_at`,
			wantArgs: []interface{}{true},
		}, {
			name:     "does not presents protection concerns",
			args:     api.ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND presents_protection_concerns = $1 ORDER BY created_at`,
			wantArgs: []interface{}{false},
		}, {
			name:     "is minor",
			args:     api.ListIndividualsOptions{IsMinor: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_minor = $1 ORDER BY created_at`,
			wantArgs: []interface{}{true},
		}, {
			name:     "is not minor",
			args:     api.ListIndividualsOptions{IsMinor: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_minor = $1 ORDER BY created_at`,
			wantArgs: []interface{}{false},
		}, {
			name:     "country_id",
			args:     api.ListIndividualsOptions{CountryID: "1"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND country_id = $1 ORDER BY created_at`,
			wantArgs: []interface{}{"1"},
		}, {
			name:     "displacement_statuses",
			args:     api.ListIndividualsOptions{DisplacementStatuses: containers.NewSet[api.DisplacementStatus](api.DisplacementStatusIDP, api.DisplacementStatusRefugee)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND displacement_status IN ($1,$2) ORDER BY created_at`,
			wantArgs: []interface{}{"idp", "refugee"},
		}, {
			name:     "all displacement_statuses",
			args:     api.ListIndividualsOptions{DisplacementStatuses: api.AllDisplacementStatuses()},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND displacement_status IN ($1,$2,$3,$4,$5,$6) ORDER BY created_at`,
			wantArgs: []interface{}{"host_community", "idp", "non_displaced", "other", "refugee", "returnee"},
		}, {
			name:     "skip",
			args:     api.ListIndividualsOptions{Skip: 10},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL ORDER BY created_at OFFSET 10`,
			wantArgs: nil,
		}, {
			name:     "take",
			args:     api.ListIndividualsOptions{Take: 10},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL ORDER BY created_at LIMIT 10`,
			wantArgs: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, args := newGetAllIndividualsSQLQuery("postgres", tt.args).build()
			assert.Equal(t, tt.wantSql, sql)
			assert.Equal(t, tt.wantArgs, args)
		})
	}
}
