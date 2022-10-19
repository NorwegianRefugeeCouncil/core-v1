package api

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/constants"
)

func UnmarshalIndividualsCSV(reader io.Reader, values *[]*Individual) ([]string, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var fields []string
	colMapping := map[string]int{}
	headerRow := records[0]
	fields = make([]string, len(headerRow))
	for i, col := range headerRow {
		fields[i] = trimString(col)
		col = trimString(col)
		colMapping[strings.Trim(col, " \n\t\r")] = i
	}

	for _, cols := range records[1:] {
		individual := &Individual{}
		if err := individual.UnmarshalCSV(colMapping, cols); err != nil {
			return nil, err
		}
		*values = append(*values, individual)
	}

	return fields, nil
}

func (i *Individual) UnmarshalCSV(colMapping map[string]int, cols []string) error {
	var err error
	for field, idx := range colMapping {
		switch field {
		case constants.FileHeaderIndividualID:
			i.ID = cols[idx]
		case constants.FileHeaderIndividualFullName:
			i.FullName = cols[idx]
		case constants.FileHeaderIndividualPreferredName:
			i.PreferredName = cols[idx]
		case constants.FileHeaderIndividualDisplacementStatus:
			i.DisplacementStatus = cols[idx]
		case constants.FileHeaderIndividualPhoneNumber:
			i.PhoneNumber = cols[idx]
		case constants.FileHeaderIndividualEmail:
			i.Email = cols[idx]
		case constants.FileHeaderIndividualAddress:
			i.Address = cols[idx]
		case constants.FileHeaderIndividualGender:
			i.Gender = cols[idx]
		case constants.FileHeaderIndividualBirthDate:
			i.BirthDate, err = ParseDate(cols[idx])
			if err != nil {
				return err
			}
		case constants.FileHeaderIndividualIsMinor:
			i.IsMinor = cols[idx] == "true"
		case constants.FileHeaderIndividualPresentsProtectionConcerns:
			i.PresentsProtectionConcerns = cols[idx] == "true"
		case constants.FileHeaderIndividualPhysicalImpairment:
			i.PhysicalImpairment = cols[idx]
		case constants.FileHeaderIndividualSensoryImpairment:
			i.SensoryImpairment = cols[idx]
		case constants.FileHeaderIndividualMentalImpairment:
			i.MentalImpairment = cols[idx]
		case constants.FileHeaderIndividualCountryID:
			i.CountryID = cols[idx]
		}

	}
	i.Normalize()
	return nil
}

func MarshalIndividualsCSV(w io.Writer, individuals []*Individual) error {
	csvEncoder := csv.NewWriter(w)
	defer csvEncoder.Flush()

	if err := csvEncoder.Write([]string{
		constants.FileHeaderIndividualID,
		constants.FileHeaderIndividualFullName,
		constants.FileHeaderIndividualPreferredName,
		constants.FileHeaderIndividualDisplacementStatus,
		constants.FileHeaderIndividualEmail,
		constants.FileHeaderIndividualAddress,
		constants.FileHeaderIndividualPhoneNumber,
		constants.FileHeaderIndividualBirthDate,
		constants.FileHeaderIndividualIsMinor,
		constants.FileHeaderIndividualGender,
		constants.FileHeaderIndividualPresentsProtectionConcerns,
		constants.FileHeaderIndividualPhysicalImpairment,
		constants.FileHeaderIndividualSensoryImpairment,
		constants.FileHeaderIndividualMentalImpairment,
		constants.FileHeaderIndividualCountryID,
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
