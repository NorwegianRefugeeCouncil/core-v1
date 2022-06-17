package utils

import "strings"

func NormalizePhoneNumber(phoneNumber string) string {
	ret := ""
	for _, c := range phoneNumber {
		if c >= '0' && c <= '9' {
			ret += string(c)
		}
	}
	return ret
}

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}
