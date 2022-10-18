package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func HandleDownload(
	userRepo db.IndividualRepo,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		ret, err := userRepo.GetAll(ctx, api.GetAllOptions{})
		if err != nil {
			l.Error("failed to get individuals", zap.Error(err))
			http.Error(w, "failed to get records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		csvEncoder := csv.NewWriter(w)
		defer csvEncoder.Flush()

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=records.csv"))
		w.Header().Set("Content-Type", "text/csv")

		if err := csvEncoder.Write([]string{
			csvHeaderIndividualID,
			csvHeaderIndividualFullName,
			csvHeaderIndividualPreferredName,
			csvHeaderIndividualDisplacementStatus,
			csvHeaderIndividualEmail,
			csvHeaderIndividualAddress,
			csvHeaderIndividualPhoneNumber,
			csvHeaderIndividualBirthDate,
			csvHeaderIndividualIsMinor,
			csvHeaderIndividualGender,
			csvHeaderIndividualPresentsProtectionConcerns,
			csvHeaderIndividualPhysicalImpairment,
			csvHeaderIndividualSensoryImpairment,
			csvHeaderIndividualMentalImpairment,
			csvHeaderIndividualCountryID,
		}); err != nil {
			l.Error("failed to write header", zap.Error(err))
			http.Error(w, "failed to write header: "+err.Error(), http.StatusInternalServerError)
			return
		}

		for _, individual := range ret {
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
				http.Error(w, "failed to write record: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

	})
}
