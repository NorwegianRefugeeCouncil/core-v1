package api

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/constants"
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
			return nil, fmt.Errorf("%s: %w is invalid", constants.FileColumnIndividualAge, err)
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
