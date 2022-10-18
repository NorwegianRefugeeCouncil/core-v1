package file_service

import (
	"context"
	"encoding/csv"
	"io"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/internal/validation"
	"go.uber.org/zap"
)

func ParseIndividualsCSV(ctx context.Context, reader io.Reader) ([]string, []*api.Individual, error) {
	l := logging.NewLogger(ctx)

	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		l.Error("failed to read csv", zap.Error(err))
		return nil, nil, err
	}
	var fields []string
	var individuals = make([]*api.Individual, len(records)-1)
	colMapping := map[string]int{}
	for i, cols := range records {
		if i == 0 {
			fields = make([]string, len(cols))
			for i, col := range cols {
				fields[i] = utils.TrimString(col)
			}
			for j, col := range cols {
				col = utils.TrimString(col)
				colMapping[strings.Trim(col, " \n\t\r")] = j
			}
		} else {
			individual, err := parseIndividualCsvRow(colMapping, cols)
			if err != nil {
				l.Error("failed to parse individual row", zap.Error(err))
				return nil, nil, err
			}
			individuals[i-1] = individual
		}
	}
	return fields, individuals, nil
}

func parseIndividualCsvRow(colMapping map[string]int, cols []string) (*api.Individual, error) {
	var err error
	var individual = &api.Individual{}
	for field, idx := range colMapping {
		switch field {
		case constants.FileHeaderIndividualID:
			individual.ID = cols[idx]
		case constants.FileHeaderIndividualFullName:
			individual.FullName = cols[idx]
		case constants.FileHeaderIndividualPreferredName:
			individual.PreferredName = cols[idx]
		case constants.FileHeaderIndividualDisplacementStatus:
			individual.DisplacementStatus = cols[idx]
		case constants.FileHeaderIndividualPhoneNumber:
			individual.PhoneNumber = cols[idx]
		case constants.FileHeaderIndividualEmail:
			individual.Email = cols[idx]
		case constants.FileHeaderIndividualAddress:
			individual.Address = cols[idx]
		case constants.FileHeaderIndividualGender:
			individual.Gender = cols[idx]
		case constants.FileHeaderIndividualBirthDate:
			individual.BirthDate, err = utils.ParseDate(cols[idx])
			if err != nil {
				return nil, err
			}
		case constants.FileHeaderIndividualIsMinor:
			individual.IsMinor = cols[idx] == "true"
		case constants.FileHeaderIndividualPresentsProtectionConcerns:
			individual.PresentsProtectionConcerns = cols[idx] == "true"
		case constants.FileHeaderIndividualPhysicalImpairment:
			individual.PhysicalImpairment = cols[idx]
		case constants.FileHeaderIndividualSensoryImpairment:
			individual.SensoryImpairment = cols[idx]
		case constants.FileHeaderIndividualMentalImpairment:
			individual.MentalImpairment = cols[idx]
		case constants.FileHeaderIndividualCountryID:
			individual.CountryID = cols[idx]
		}

	}
	validation.NormalizeIndividual(individual)
	return individual, nil
}
