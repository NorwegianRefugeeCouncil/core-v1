package handlers

import (
	"context"
	"encoding/csv"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func UploadHandler(repo db.IndividualRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const (
			formParamFile = "file"
		)

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		// todo: find sensible max memory value
		maxMemory := int64(1024 * 1024 * 1024)
		if err := r.ParseMultipartForm(maxMemory); err != nil {
			l.Error("failed to parse multipart form", zap.Error(err))
			http.Error(w, "failed to parse form: "+err.Error(), http.StatusInternalServerError)
			return
		}

		formFile, _, err := r.FormFile(formParamFile)
		if err != nil {
			l.Error("failed to get form file", zap.Error(err))
			http.Error(w, "failed to parse input file: "+err.Error(), http.StatusBadRequest)
			return
		}

		fields, individuals, err := parseIndividualsCSV(ctx, formFile)
		if err != nil {
			l.Error("failed to parse csv", zap.Error(err))
			http.Error(w, "failed to parse csv: "+err.Error(), http.StatusBadRequest)
			return
		}

		_, err = repo.PutMany(r.Context(), individuals, fields)
		if err != nil {
			l.Error("failed to put individuals", zap.Error(err))
			http.Error(w, "failed to put records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/individuals", http.StatusSeeOther)

	})
}

func parseIndividualsCSV(ctx context.Context, reader io.Reader) ([]string, []*api.Individual, error) {

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
			fields = cols
			for j, col := range cols {
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
