package api

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/slices"
)

// Unmarshal

func UnmarshalIndividualsCSV(reader io.Reader, individuals *[]*Individual, fields *[]string) error {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	return unmarshalIndividualsTabularData(records, individuals, fields)
}

func UnmarshalIndividualsExcel(reader io.Reader, individuals *[]*Individual, fields *[]string) error {
	f, err := excelize.OpenReader(reader)

	if err != nil {
		return err
	}

	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 	}
	// }()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		err := errors.New("no sheets found")
		return err
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		err := errors.New("no rows found")
		return err
	}

	return unmarshalIndividualsTabularData(rows, individuals, fields)
}

func unmarshalIndividualsTabularData(data [][]string, individuals *[]*Individual, fields *[]string) error {
	colMapping := map[string]int{}
	headerRow := data[0]
	for i, col := range headerRow {
		*fields = append(*fields, constants.IndividualFileToDBMap[trimString(col)])
		col = trimString(col)
		colMapping[strings.Trim(col, " \n\t\r")] = i
	}

	for _, cols := range data[1:] {
		individual := &Individual{}
		if err := individual.unmarshalTabularData(colMapping, cols); err != nil {
			return err
		}
		*individuals = append(*individuals, individual)
	}

	return nil
}

func (i *Individual) unmarshalTabularData(colMapping map[string]int, cols []string) error {
	var err error
	for field, idx := range colMapping {
		switch field {
		case constants.FileColumnIndividualID:
			i.ID = cols[idx]
		case constants.FileColumnIndividualFullName:
			i.FullName = cols[idx]
		case constants.FileColumnIndividualPreferredName:
			i.PreferredName = cols[idx]
		case constants.FileColumnIndividualDisplacementStatus:
			i.DisplacementStatus = cols[idx]
		case constants.FileColumnIndividualPhoneNumber:
			i.PhoneNumber = cols[idx]
		case constants.FileColumnIndividualEmail:
			i.Email = cols[idx]
		case constants.FileColumnIndividualAddress:
			i.Address = cols[idx]
		case constants.FileColumnIndividualGender:
			i.Gender = cols[idx]
		case constants.FileColumnIndividualBirthDate:
			i.BirthDate, err = ParseDate(cols[idx])
			if err != nil {
				return err
			}
		case constants.FileColumnIndividualIsMinor:
			i.IsMinor = isTrue(cols[idx])
		case constants.FileColumnIndividualPresentsProtectionConcerns:
			i.PresentsProtectionConcerns = isTrue(cols[idx])
		case constants.FileColumnIndividualPhysicalImpairment:
			i.PhysicalImpairment = cols[idx]
		case constants.FileColumnIndividualSensoryImpairment:
			i.SensoryImpairment = cols[idx]
		case constants.FileColumnIndividualMentalImpairment:
			i.MentalImpairment = cols[idx]
		case constants.FileColumnIndividualCountryID:
			i.CountryID = cols[idx]
		}

	}
	i.Normalize()
	return nil
}

// Marshal

func MarshalIndividualsCSV(w io.Writer, individuals []*Individual) error {
	csvEncoder := csv.NewWriter(w)
	defer csvEncoder.Flush()

	if err := csvEncoder.Write([]string{
		constants.FileColumnIndividualID,
		constants.FileColumnIndividualFullName,
		constants.FileColumnIndividualPreferredName,
		constants.FileColumnIndividualDisplacementStatus,
		constants.FileColumnIndividualEmail,
		constants.FileColumnIndividualAddress,
		constants.FileColumnIndividualPhoneNumber,
		constants.FileColumnIndividualBirthDate,
		constants.FileColumnIndividualIsMinor,
		constants.FileColumnIndividualGender,
		constants.FileColumnIndividualPresentsProtectionConcerns,
		constants.FileColumnIndividualPhysicalImpairment,
		constants.FileColumnIndividualSensoryImpairment,
		constants.FileColumnIndividualMentalImpairment,
		constants.FileColumnIndividualCountryID,
	}); err != nil {
		return err
	}

	for _, individual := range individuals {
		if err := individual.MarshalCSV(csvEncoder); err != nil {
			return err
		}
	}

	return nil
}

func (i *Individual) MarshalCSV(csvEncoder *csv.Writer) error {
	var birthDate string

	if i.BirthDate != nil {
		birthDate = i.BirthDate.Format("2006-01-02")
	}

	if err := csvEncoder.Write([]string{
		i.ID,
		i.FullName,
		i.PreferredName,
		i.DisplacementStatus,
		i.Email,
		i.Address,
		i.PhoneNumber,
		birthDate,
		strconv.FormatBool(i.IsMinor),
		i.Gender,
		strconv.FormatBool(i.PresentsProtectionConcerns),
		i.PhysicalImpairment,
		i.SensoryImpairment,
		i.MentalImpairment,
		i.CountryID,
	}); err != nil {
		return err
	}

	return nil
}

var TRUE_VALUES = []string{"true", "yes", "1"}

func isTrue(value string) bool {
	return slices.Contains(TRUE_VALUES, strings.ToLower(value))
}
