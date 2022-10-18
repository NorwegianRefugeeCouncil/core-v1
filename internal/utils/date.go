package utils

import "time"

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
