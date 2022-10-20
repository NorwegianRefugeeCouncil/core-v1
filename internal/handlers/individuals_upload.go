package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func UploadHandler(individualRepo db.IndividualRepo) http.Handler {

	const (
		formParamFile = "file"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		authIntf, err := utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		allowedCountryIDs := authIntf.GetCountryIDsWithReadWritePermissions()
		if len(allowedCountryIDs) == 0 {
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

		filename := r.MultipartForm.File[formParamFile][0].Filename

		formFile, _, err := r.FormFile(formParamFile)
		if err != nil {
			l.Error("failed to get form file", zap.Error(err))
			http.Error(w, "failed to parse input file: "+err.Error(), http.StatusBadRequest)
			return
		}

		var individuals []*api.Individual
		var fields []string

		if strings.HasSuffix(filename, ".csv") {
			if err = api.UnmarshalIndividualsCSV(formFile, &individuals, &fields); err != nil {
				l.Error("failed to parse csv", zap.Error(err))
				http.Error(w, "failed to parse csv: "+err.Error(), http.StatusBadRequest)
				return
			}
		} else if strings.HasSuffix(filename, ".xlsx") || strings.HasSuffix(filename, ".xls") {
			if err = api.UnmarshalIndividualsExcel(formFile, &individuals, &fields); err != nil {
				l.Error("failed to parse excel file", zap.Error(err))
				http.Error(w, "failed to parse excel file: "+err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			l.Error(fmt.Sprintf("unsupported content type: %s", r.Header.Get("Content-Type")))
			http.Error(w, "invalid content type", http.StatusBadRequest)
			return
		}

		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		for _, individual := range individuals {
			individual.CountryID = selectedCountryID
			fields = append(fields, "country_id")
		}

		for _, individual := range individuals {
			if !authIntf.CanReadWriteToCountryID(individual.CountryID) {
				l.Warn("user does not have permission to upload individuals to country", zap.String("country_id", individual.CountryID))
				http.Error(w, "You are not allowed to upload to country: "+individual.CountryID, http.StatusForbidden)
				return
			}
		}

		_, err = individualRepo.PutMany(r.Context(), individuals, fields)
		if err != nil {
			l.Error("failed to put individuals", zap.Error(err))
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
