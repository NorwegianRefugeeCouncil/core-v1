package api

import (
	"encoding/csv"
	"errors"
	"fmt"
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
	return UnmarshalIndividualsTabularData(records, individuals, fields)
}

func UnmarshalIndividualsExcel(reader io.Reader, individuals *[]*Individual, fields *[]string) error {
	f, err := excelize.OpenReader(reader)

	if err != nil {
		return err
	}

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

	return UnmarshalIndividualsTabularData(rows, individuals, fields)
}

func UnmarshalIndividualsTabularData(data [][]string, individuals *[]*Individual, fields *[]string) error {
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
		}
	}
	i.Normalize()
	return nil
}

// Marshal

func MarshalIndividualsCSV(w io.Writer, individuals []*Individual) error {
	csvEncoder := csv.NewWriter(w)
	defer csvEncoder.Flush()

	if err := csvEncoder.Write(constants.IndividualFileColumns); err != nil {
		return err
	}

	for _, individual := range individuals {
		row, err := individual.marshalTabularData()
		if err != nil {
			return err
		}
		if err := csvEncoder.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func MarshalIndividualsExcel(w io.Writer, individuals []*Individual) error {
	const sheetName = "Individuals"

	f := excelize.NewFile()

	sheet := f.NewSheet(sheetName)

	if err := f.SetSheetRow(sheetName, "A1", &constants.IndividualFileColumns); err != nil {
		return err
	}

	for i, individual := range individuals {
		row, err := individual.marshalTabularData()
		if err != nil {
			return err
		}
		if err := f.SetSheetRow(sheetName, fmt.Sprintf("A%d", i+2), &row); err != nil {
			return err
		}
	}

	f.SetActiveSheet(sheet)

	if err := f.Write(w); err != nil {
		return err
	}

	return nil
}

func (i *Individual) marshalTabularData() ([]string, error) {
	row := make([]string, len(constants.IndividualFileColumns))
	for j, col := range constants.IndividualFileColumns {
		value, err := i.GetFieldValue(constants.IndividualFileToDBMap[col])
		if err != nil {
			return nil, err
		}

		switch col {
		case constants.FileColumnIndividualBirthDate:
			var birthDate string
			if i.BirthDate != nil {
				birthDate = i.BirthDate.Format("2006-01-02")
			}
			row[j] = birthDate
		case constants.FileColumnIndividualIsMinor, constants.FileColumnIndividualPresentsProtectionConcerns:
			row[j] = strconv.FormatBool(value.(bool))
		default:
			row[j] = value.(string)
		}
	}
	return row, nil
}

var TRUE_VALUES = []string{"true", "yes", "1"}

func isTrue(value string) bool {
	return slices.Contains(TRUE_VALUES, strings.ToLower(value))
}
