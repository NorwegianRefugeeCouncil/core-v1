package api

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
	"time"
)

var dateFormat = "2006-01-02"
var minBirthdate = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

func ParseDate(s string) (*time.Time, error) {
	if s != "" {
		date, err := time.Parse(dateFormat, s)
		if err != nil {
			return nil, err
		}
		return &date, nil
	}
	return nil, nil
}

func ParseBirthdate(s string) (*time.Time, error) {
	if s != "" {
		date, err := time.Parse(dateFormat, s)
		if err != nil {
			return nil, fmt.Errorf("%s: %s is invalid: %w", constants.FileColumnIndividualBirthDate, date, err)
		}
		if date.Before(minBirthdate) {
			return nil, fmt.Errorf("%s: %s is before %s", constants.FileColumnIndividualBirthDate, date, minBirthdate)
		}
		return &date, nil
	}
	return nil, nil
}

func ParseAge(s string) (*int, error) {
	if s != "" {
		age, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", constants.FileColumnIndividualAge, err)
		}
		if age < 0 {
			return nil, fmt.Errorf("%s: %d is negative", constants.FileColumnIndividualAge, age)
		}
		return &age, nil
	}
	return nil, nil
}

func trimString(s string) string {
	return strings.Trim(s, " \t\n\r")
}

func NormalizePhoneNumber(phoneNumber string) string {
	ret := ""
	for _, c := range phoneNumber {
		if c >= '0' && c <= '9' {
			ret += string(c)
		}
	}
	return ret
}

func normalizeEmail(email string) string {
	return strings.ToLower(email)
}

func makeIndexSetWithSkip(size int, skip int) containers.Set[int] {
	a := containers.Set[int]{}
	for i := 0; i < size; i++ {
		a.Add(i)
	}
	a.Remove(skip)
	return a
}

func getTimeFormatForField(field string) string {
	switch field {
	case constants.DBColumnIndividualUpdatedAt:
		return time.RFC3339
	case constants.DBColumnIndividualCreatedAt:
		return time.RFC3339
	case constants.DBColumnIndividualDeletedAt:
		return time.RFC3339
	default:
		return "2006-01-02"
	}
}

var TRUE_VALUES = []string{"true", "yes", "1"}
var FALSE_VALUES = []string{"false", "no", "0"}

func getValidatedBoolean(value string) (bool, error) {
	isExplicitlyTrue := slices.Contains(TRUE_VALUES, strings.ToLower(value))
	isExplicitlyFalse := slices.Contains(FALSE_VALUES, strings.ToLower(value))
	if !isExplicitlyTrue && !isExplicitlyFalse {
		return false, fmt.Errorf("invalid boolean value \"%s\". Valid values are: \"%s\", \"%s\"", value, strings.Join(TRUE_VALUES, "\", \""), strings.Join(FALSE_VALUES, "\", \""))
	}
	return isExplicitlyTrue, nil
}

func stringArrayToInterfaceArray(row []string) []interface{} {
	var result []interface{}
	for _, col := range row {
		result = append(result, col)
	}
	return result
}
