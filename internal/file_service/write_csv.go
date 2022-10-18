package file_service

import (
	"context"
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func WriteIndividualCSV(ctx context.Context, w http.ResponseWriter, individuals []*api.Individual) error {
	l := logging.NewLogger(ctx)

	csvEncoder := csv.NewWriter(w)
	defer csvEncoder.Flush()

	w.Header().Set("Content-Disposition", "attachment; filename=records.csv")
	w.Header().Set("Content-Type", "text/csv")

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
		l.Error("failed to write header", zap.Error(err))
		return err
	}

	for _, individual := range individuals {
		var birthDate string
		if individual.BirthDate != nil {
			birthDate = individual.BirthDate.Format("2006-01-02")
		}
		if err := csvEncoder.Write([]string{
			individual.ID,
			individual.FullName,
			individual.PreferredName,
			individual.DisplacementStatus,
			individual.Email,
			individual.Address,
			individual.PhoneNumber,
			birthDate,
			strconv.FormatBool(individual.IsMinor),
			individual.Gender,
			strconv.FormatBool(individual.PresentsProtectionConcerns),
			individual.PhysicalImpairment,
			individual.SensoryImpairment,
			individual.MentalImpairment,
			individual.CountryID,
		}); err != nil {
			l.Error("failed to write record", zap.Error(err))
			return err
		}
	}

	return nil
}
