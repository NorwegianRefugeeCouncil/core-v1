package api

import (
	"strings"
	"time"
)

var dateFormat = "2006-01-02"

func ParseDate(s string) (*time.Time, error) {
	if s != "" {
		birthDate, err := time.Parse(dateFormat, s)
		if err != nil {
			return nil, err
		}
		return &birthDate, nil
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
