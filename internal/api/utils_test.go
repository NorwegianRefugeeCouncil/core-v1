package api

import (
	"errors"
	"github.com/nrc-no/notcore/internal/containers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMakeIndexSetWithSkip(t *testing.T) {
	tests := []struct {
		name string
		size int
		skip int
		want containers.Set[int]
	}{
		{
			name: "skip 1",
			size: 5,
			skip: 1,
			want: containers.NewSet[int](0, 2, 3, 4),
		}, {
			name: "skip out of range",
			size: 5,
			skip: 6,
			want: containers.NewSet[int](0, 1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := makeIndexSetWithSkip(tt.size, tt.skip)
			assert.Equal(t, tt.want, set)
		})
	}
}

func TestParseDate(t *testing.T) {

	t.Run("empty string", func(t *testing.T) {
		result, err := ParseDate("")
		assert.Nil(t, result)
		assert.Nil(t, err)
	})

	tests := []struct {
		name    string
		value   string
		want    time.Time
		error   error
		wantErr bool
	}{
		{
			name:    "2006-01-02",
			value:   "2006-01-02",
			want:    time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "2006-13-40",
			value:   "2006-13-40",
			error:   errors.New("parsing time \"2006-13-40\": month out of range"),
			wantErr: true,
		},
		{
			name:    "01.02.2006",
			value:   "01.02.2006",
			error:   errors.New("parsing time \"01.02.2006\" as \"2006-01-02\": cannot parse \"2.2006\" as \"2006\""),
			wantErr: true,
		},
		{
			name:    "01. Feb 2006",
			value:   "01. Feb 2006",
			error:   errors.New("parsing time \"01. Feb 2006\" as \"2006-01-02\": cannot parse \"Feb 2006\" as \"2006\""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.error.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &tt.want, result)
			}
		})
	}
}

func TestBirthParseDate(t *testing.T) {

	t.Run("empty string", func(t *testing.T) {
		result, err := ParseBirthdate("")
		assert.Nil(t, result)
		assert.Nil(t, err)
	})

	tests := []struct {
		name    string
		value   string
		want    time.Time
		error   error
		wantErr bool
	}{
		{
			name:    "2006-01-02",
			value:   "2006-01-02",
			want:    time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "1806-01-02",
			value:   "1806-01-02",
			error:   errors.New("birth_date: 1806-01-02 00:00:00 +0000 UTC is before 1900-01-01 00:00:00 +0000 UTC"),
			wantErr: true,
		},
		{
			name:    "2006-13-40",
			value:   "2006-13-40",
			error:   errors.New("birth_date: 0001-01-01 00:00:00 +0000 UTC is invalid: parsing time \"2006-13-40\": month out of range"),
			wantErr: true,
		},
		{
			name:    "01.02.2006",
			value:   "01.02.2006",
			error:   errors.New("birth_date: 0001-01-01 00:00:00 +0000 UTC is invalid: parsing time \"01.02.2006\" as \"2006-01-02\": cannot parse \"2.2006\" as \"2006\""),
			wantErr: true,
		},
		{
			name:    "01. Feb 2006",
			value:   "01. Feb 2006",
			error:   errors.New("birth_date: 0001-01-01 00:00:00 +0000 UTC is invalid: parsing time \"01. Feb 2006\" as \"2006-01-02\": cannot parse \"Feb 2006\" as \"2006\""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseBirthdate(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.error.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &tt.want, result)
			}
		})
	}
}

func TestParseAge(t *testing.T) {

	t.Run("empty string", func(t *testing.T) {
		result, err := ParseAge("")
		assert.Nil(t, result)
		assert.Nil(t, err)
	})

	t.Run("not a number", func(t *testing.T) {
		result, err := ParseAge("abc")
		assert.Nil(t, result)
		assert.Equal(t, err.Error(), errors.New("age: strconv.Atoi: parsing \"abc\": invalid syntax").Error())
	})

	tests := []struct {
		name    string
		value   string
		want    int
		error   error
		wantErr bool
	}{
		{
			name:    "34",
			value:   "34",
			want:    34,
			wantErr: false,
		},
		{
			name:    "-1",
			value:   "-1",
			error:   errors.New("age: -1 is negative"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAge(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.error.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &tt.want, result)
			}
		})
	}
}

func TestGetValidatedBoolean(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "\"true\"",
			value: "true",
			want:  true,
		}, {
			name:  "\"yes\"",
			value: "yes",
			want:  true,
		}, {
			name:  "\"1\"",
			value: "1",
			want:  true,
		}, {
			name:  "\"false\"",
			value: "false",
			want:  false,
		}, {
			name:  "\"no\"",
			value: "no",
			want:  false,
		}, {
			name:  "\"0\"",
			value: "0",
			want:  false,
		}, {
			name:  "\"nope\"",
			value: "nope",
			want:  false,
		}, {
			name:  "\"yeah\"",
			value: "yeah",
			want:  false,
		}, {
			name:  "empty string",
			value: "",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isExplicitlyTrue(tt.value))
		})
	}
}
