package handlers

import (
	"encoding/csv"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/utils"
)

func UploadHandler(repo db.IndividualRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseMultipartForm(1000000); err != nil {
			http.Error(w, "failed", http.StatusInternalServerError)
			return
		}

		formFile, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "failed", http.StatusBadRequest)
			return
		}

		csvReader := csv.NewReader(formFile)
		csvReader.TrimLeadingSpace = true
		records, err := csvReader.ReadAll()
		if err != nil {
			http.Error(w, "failed to read records: "+err.Error(), http.StatusBadRequest)
			return
		}

		var individuals = make([]*api.Individual, len(records)-1)

		colMapping := map[string]int{}
		for i, cols := range records {
			if i == 0 {
				for j, col := range cols {
					colMapping[strings.Trim(col, " \n\t\r")] = j
				}
			} else {
				var individual = &api.Individual{}
				for field, idx := range colMapping {
					if field == "id" {
						individual.ID = cols[idx]
					} else if field == "full_name" {
						individual.FullName = cols[idx]
					} else if field == "address" {
						individual.Address = cols[idx]
					} else if field == "phone_number" {
						individual.PhoneNumber = cols[idx]
					} else if field == "email" {
						individual.Email = utils.NormalizeEmail(cols[idx])
					} else if field == "gender" {
						individual.Gender = cols[idx]
					} else if field == "birth_date" {
						birthDateStr := cols[idx]
						if len(birthDateStr) > 0 {
							birthDate, err := time.Parse("2006-01-02", birthDateStr)
							if err != nil {
								http.Error(w, "failed to parse birth date: "+err.Error(), http.StatusBadRequest)
								return
							}
							individual.BirthDate = &birthDate
						}
					}
				}
				individuals[i-1] = individual
				individual.NormalizedPhoneNumber = utils.NormalizePhoneNumber(individual.PhoneNumber)
			}
		}

		_, err = repo.PutMany(r.Context(), individuals)
		if err != nil {
			http.Error(w, "failed to put records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/individuals", http.StatusSeeOther)

	})
}

func parseQryParamInt(r *http.Request, key string) (int, error) {
	strValue := r.URL.Query().Get(key)
	if len(strValue) != 0 {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			return 0, err
		}
		return intValue, nil
	}
	return 0, nil
}
