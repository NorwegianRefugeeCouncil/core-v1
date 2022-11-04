package logutils

import "strings"

func Escape(s string) string {
	var out = strings.Builder{}
	for _, r := range s {
		switch r {
		case '\n':
			out.WriteString("\\n")
		case '\t':
			out.WriteString("\\t")
		case '\r':
			out.WriteString("\\r")
		default:
			out.WriteRune(r)
		}
	}
	return out.String()
}
