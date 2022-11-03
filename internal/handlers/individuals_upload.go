package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
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

		fieldSet := containers.NewStringSet(fields...)
		fieldSet.Add("country_id")
		fields = fieldSet.Items()

		var individualIds []string
		for _, individual := range individuals {
			individual.CountryID = selectedCountryID
			if len(individual.ID) > 0 {
				individualIds = append(individualIds, individual.ID)
			}
		}

		existingIndividuals, err := individualRepo.GetAll(ctx, api.GetAllOptions{IDs: individualIds})
		if err != nil {
			l.Error("failed to get existing individuals", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		invalidIndividualIds := validateIndividualsExistInCountry(individualIds, existingIndividuals, selectedCountryID)
		if len(invalidIndividualIds) > 0 {
			l.Warn("user trying to update individuals that don't exist or are in the wrong country", zap.Strings("individual_ids", invalidIndividualIds))
			http.Error(w, fmt.Sprintf("individuals not found: %v", invalidIndividualIds), http.StatusNotFound)
			return
		}

		_, err = individualRepo.PutMany(r.Context(), individuals, fields)
		if err != nil {
			l.Error("failed to put individuals", zap.Error(err))
			http.Error(w, "failed to put records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals", selectedCountryID), http.StatusSeeOther)
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
