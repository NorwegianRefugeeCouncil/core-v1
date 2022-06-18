package handlers

import (
	"encoding/csv"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
)

func UploadHandler(repo db.IndividualRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// todo: find sensible max memory value
		maxMemory := int64(1024 * 1024 * 1024)
		if err := r.ParseMultipartForm(maxMemory); err != nil {
			http.Error(w, "failed", http.StatusInternalServerError)
			return
		}

		formFile, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "failed", http.StatusBadRequest)
			return
		}

		fields, individuals, err := parseIndividualsCSV(formFile)
		if err != nil {
			http.Error(w, "failed", http.StatusBadRequest)
			return
		}

		_, err = repo.PutMany(r.Context(), individuals, fields)
		if err != nil {
			http.Error(w, "failed to put records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/individuals", http.StatusSeeOther)

	})
}

func parseIndividualsCSV(reader io.Reader) ([]string, []*api.Individual, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	var fields []string
	var individuals = make([]*api.Individual, len(records)-1)
	colMapping := map[string]int{}
	for i, cols := range records {
		if i == 0 {
			fields = cols
			for j, col := range cols {
				colMapping[strings.Trim(col, " \n\t\r")] = j
			}
		} else {
			individual, err := parseIndividualCsvRow(colMapping, cols)
			if err != nil {
				return nil, nil, err
			}
			individuals[i-1] = individual
		}
	}
	return fields, individuals, nil
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
