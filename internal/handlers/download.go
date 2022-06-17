package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"

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

		csvEncoder.Write([]string{
			"id",
			"full_name",
			"email",
			"address",
			"phone_number",
			"birth_date",
			"gender",
		})

		for _, individual := range ret {
			var birthDate string
			if individual.BirthDate != nil {
				birthDate = individual.BirthDate.Format("2006-01-02")
			}
			csvEncoder.Write([]string{
				individual.ID,
				individual.FullName,
				individual.Email,
				individual.Address,
				individual.PhoneNumber,
				birthDate,
				individual.Gender,
			})
		}

	})
}
