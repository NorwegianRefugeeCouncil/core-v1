package utils

import "strings"

func TrimString(s string) string {
	return strings.Trim(s, " \t\n\r")
}
