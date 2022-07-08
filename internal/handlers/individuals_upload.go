package handlers

import (
	"context"
	"encoding/csv"
	"github.com/nrc-no/notcore/internal/clients/zanzibar"
	"github.com/nrc-no/notcore/internal/containers"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func UploadHandler(client *zanzibar.ZanzibarClient, individualRepo db.IndividualRepo) http.Handler {

	const (
		formParamFile = "file"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		countryIDsWithWritePermission := utils.GetCountryIDsWithPermission(ctx, "write")
		if !utils.IsGlobalAdmin(ctx) && len(countryIDsWithWritePermission) == 0 {
			l.Warn("User does not have permission to upload individuals")
			http.Error(w, "You are not allowed to upload", http.StatusForbidden)
			return
		}

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

		allowedCountryCodes, err := client.CheckPermittedLocations(ctx, zanzibar.LocationType_Country)

		var validIndividuals = []*api.Individual{}
		forbiddenCountryCodes := containers.NewStringSet()
		allowedCountrySet := containers.NewStringSet(allowedCountryCodes...)

		for _, individual := range individuals {
			for _, countryCode := range individual.Countries {
				if !utils.IsGlobalAdmin(ctx) && !allowedCountrySet.Contains(countryCode) {
					forbiddenCountryCodes.Add(countryCode)
				} else {
					validIndividuals = append(validIndividuals, individual)
				}
			}
		}
		if !forbiddenCountryCodes.IsEmpty() {
			forbiddenCountryCodesStr := strings.Join(forbiddenCountryCodes.Items(), ", ")
			l.Warn("user does not have permission to upload individuals to country", zap.String("country", forbiddenCountryCodesStr))
			// TODO show error in FE
			//http.Error(w, "You are not allowed to upload to countries: "+forbiddenCountryCodesStr, http.StatusForbidden)
		}

		_, err = individualRepo.PutMany(r.Context(), validIndividuals, fields)
		if err != nil {
			l.Error("failed to put individuals", zap.Error(err))
			http.Error(w, "failed to put records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = client.AddIndividualsToLocation(ctx, zanzibar.LocationType_Country, validIndividuals)
		if err != nil {
			l.Error("failed to add individuals to zanzibar graph", zap.Error(err))
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
			fields = make([]string, len(cols))
			for i, col := range cols {
				fields[i] = trimString(col)
			}
			for j, col := range cols {
				col = trimString(col)
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
