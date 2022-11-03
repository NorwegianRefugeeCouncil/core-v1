package db

import (
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func Test_newGetAllIndividualsSQLQuery(t *testing.T) {
	someDate, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	zeroTime := time.Time{}
	const defaultQuery = `SELECT * FROM individuals WHERE deleted_at IS NULL ORDER BY created_at`
	tests := []struct {
		name     string
		args     api.GetAllOptions
		wantSql  string
		wantArgs []interface{}
	}{
		{
			name:    "empty",
			args:    api.GetAllOptions{},
			wantSql: defaultQuery,
		}, {
			name:     "ids",
			args:     api.GetAllOptions{IDs: []string{"1", "2"}},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND id IN ($1,$2) ORDER BY created_at`,
			wantArgs: []interface{}{"1", "2"},
		}, {
			name:     "full name",
			args:     api.GetAllOptions{FullName: "John"},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND (full_name ILIKE $1 OR preferred_name ILIKE $2) ORDER BY created_at`,
			wantArgs: []interface{}{"%John%", "%John%"},
		}, {
			name:     "address",
			args:     api.GetAllOptions{Address: "123 Main St"},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND address ILIKE $1 ORDER BY created_at`,
			wantArgs: []interface{}{"%123 Main St%"},
		}, {
			name:     "genders",
			args:     api.GetAllOptions{Genders: []string{"male", "female"}},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND gender IN ($1,$2) ORDER BY created_at`,
			wantArgs: []interface{}{"male", "female"},
		}, {
			name:     "birth date from",
			args:     api.GetAllOptions{BirthDateFrom: &someDate},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND birth_date >= $1 ORDER BY created_at`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:    "zero birth date from",
			args:    api.GetAllOptions{BirthDateFrom: &zeroTime},
			wantSql: defaultQuery,
		}, {
			name:     "birth date to",
			args:     api.GetAllOptions{BirthDateTo: &someDate},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND birth_date <= $1 ORDER BY created_at`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:    "zero birth date to",
			args:    api.GetAllOptions{BirthDateFrom: &zeroTime},
			wantSql: defaultQuery,
		}, {
			name:     "phone number",
			args:     api.GetAllOptions{PhoneNumber: "1234567890"},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND normalized_phone_number ILIKE $1 ORDER BY created_at`,
			wantArgs: []interface{}{"%1234567890%"},
		}, {
			name:     "email",
			args:     api.GetAllOptions{Email: "info@email.com"},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND email = $1 ORDER BY created_at`,
			wantArgs: []interface{}{"info@email.com"},
		}, {
			name:     "presents protection concerns",
			args:     api.GetAllOptions{PresentsProtectionConcerns: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND presents_protection_concerns = $1 ORDER BY created_at`,
			wantArgs: []interface{}{true},
		}, {
			name:     "does not presents protection concerns",
			args:     api.GetAllOptions{PresentsProtectionConcerns: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND presents_protection_concerns = $1 ORDER BY created_at`,
			wantArgs: []interface{}{false},
		}, {
			name:     "is minor",
			args:     api.GetAllOptions{IsMinor: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND is_minor = $1 ORDER BY created_at`,
			wantArgs: []interface{}{true},
		}, {
			name:     "is not minor",
			args:     api.GetAllOptions{IsMinor: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND is_minor = $1 ORDER BY created_at`,
			wantArgs: []interface{}{false},
		}, {
			name:     "country_id",
			args:     api.GetAllOptions{CountryID: "1"},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND country_id = $1 ORDER BY created_at`,
			wantArgs: []interface{}{"1"},
		}, {
			name:     "displacement_statuses",
			args:     api.GetAllOptions{DisplacementStatuses: []string{"1", "2"}},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL AND displacement_status IN ($1,$2) ORDER BY created_at`,
			wantArgs: []interface{}{"1", "2"},
		}, {
			name:     "skip",
			args:     api.GetAllOptions{Skip: 10},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL ORDER BY created_at OFFSET 10`,
			wantArgs: nil,
		}, {
			name:     "take",
			args:     api.GetAllOptions{Take: 10},
			wantSql:  `SELECT * FROM individuals WHERE deleted_at IS NULL ORDER BY created_at LIMIT 10`,
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
