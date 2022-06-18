package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
)

func HandleDownload(repo db.IndividualRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ret, err := repo.GetAll(r.Context(), api.GetAllOptions{})
		if err != nil {
			http.Error(w, "failed to get records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		csvEncoder := csv.NewWriter(w)
		defer csvEncoder.Flush()

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=records.csv"))
		w.Header().Set("Content-Type", "text/csv")

		if err := csvEncoder.Write([]string{
			"id",
			"full_name",
			"preferred_name",
			"displacement_status",
			"email",
			"address",
			"phone_number",
			"birth_date",
			"is_minor",
			"gender",
			"presents_protection_concerns",
			"physical_impairment",
			"sensory_impairment",
			"mental_impairment",
			"countries",
		}); err != nil {
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
				strings.Join(individual.Countries, ","),
			}); err != nil {
				http.Error(w, "failed to write record: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

	})
}
