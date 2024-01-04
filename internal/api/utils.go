package api

import (
	"errors"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/locales"
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
			t := locales.GetTranslator()
			col := t(constants.FileColumnIndividualBirthDate)
			return nil, errors.New(t("error_parse_birthdate_invalid", col, s, err))
		}
		if date.Before(minBirthdate) {
			t := locales.GetTranslator()
			col := t(constants.FileColumnIndividualBirthDate)
			return nil, errors.New(t("error_parse_birthdate_minimum", col, date, minBirthdate))
		}
		return &date, nil
	}
	return nil, nil
}

func ParseAge(s string) (*int, error) {
	if s != "" {
		age, err := strconv.Atoi(s)
		if err != nil {
			t := locales.GetTranslator()
			col := t(constants.FileColumnIndividualAge)
			return nil, fmt.Errorf("%s: %w", col, err)
		}
		if age < 0 {
			t := locales.GetTranslator()
			col := t(constants.FileColumnIndividualAge)
			return nil, errors.New(t("error_parse_age", col, age))
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

func stringArrayToInterfaceArray(row []string) []interface{} {
	var result []interface{}
	for _, col := range row {
		result = append(result, col)
	}
	return result
}

var indexColumnName = "index"

func AddIndexColumn(df dataframe.DataFrame) dataframe.DataFrame {
	indexes := []int{}
	for i := 0; i < df.Nrow(); i++ {
		indexes = append(indexes, i)
	}
	dfMutated := df.Mutate(series.New(indexes, series.String, indexColumnName))
	return dfMutated
}

var TRUE_VALUES = []string{"true", "yes", "1"}

func isExplicitlyTrue(value string) bool {
	return slices.Contains(TRUE_VALUES, strings.ToLower(value))
}
