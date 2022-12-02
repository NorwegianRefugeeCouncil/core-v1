package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func generateUniqueDownloadFileNameForCountryAndExtension(selectedCountryID string, format string) string {
	fileName := fmt.Sprintf("%s_%s.%s", selectedCountryID, uuid.New().String(), format)
	return fileName
}

func assertValidFileNameForCountry(fileName, wantCountryID string) (string, error) {
	parts := strings.Split(fileName, "_")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid file name")
	}
	countryID, err := uuid.Parse(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid file name")
	}
	if countryID.String() != wantCountryID {
		return "", fmt.Errorf("invalid file name")
	}

	secondParts := strings.Split(parts[1], ".")
	if len(secondParts) != 2 {
		return "", fmt.Errorf("invalid file name")
	}

	ext := secondParts[1]
	if !isValidFileExtension(ext) {
		return "", fmt.Errorf("invalid file name")
	}

	_, err = uuid.Parse(secondParts[0])
	if err != nil {
		return "", fmt.Errorf("invalid file name")
	}

	return "download." + ext, nil
}

func isValidFileExtension(ext string) bool {
	return ext == "csv" || ext == "xlsx"
}

func HandleDownload(
	userRepo db.IndividualRepo,
) http.Handler {

	const queryParamFormat = "format"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			http.Error(w, "couldn't get selected country id: "+err.Error(), http.StatusInternalServerError)
			return
		}

		file := r.URL.Query().Get("file")
		if file != "" {

			// at this point, the file was already created and ready for download.

			resultFileName, err := assertValidFileNameForCountry(file, selectedCountryID)
			if err != nil {
				l.Error("invalid file name", zap.Error(err))
				http.Error(w, "invalid file name: "+err.Error(), http.StatusBadRequest)
				return
			}
			filePath := path.Join("/tmp", file)

			// check that the file already exists
			_, err = os.Stat(filePath)
			if err != nil {
				l.Error("failed to get file info", zap.Error(err))
				http.Error(w, "failed to get file info: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// open the file
			file, err := os.Open(filePath)
			if err != nil {
				l.Error("failed to open file", zap.Error(err))
				http.Error(w, "failed to open file: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// serve the file
			http.ServeContent(w, r, resultFileName, time.Now(), file)
			return
		}

		// the file was not created yet, so we need to create it, and redirect
		// the request to the same endpoint, but with the file query param.

		if err := r.ParseForm(); err != nil {
			l.Error("failed to parse form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var getAllOptions api.ListIndividualsOptions
		if err := api.NewIndividualListFromURLValues(r.Form, &getAllOptions); err != nil {
			l.Error("failed to parse options", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		getAllOptions.CountryID = selectedCountryID
		getAllOptions.Take = 0
		getAllOptions.Skip = 0

		format := r.URL.Query().Get(queryParamFormat)
		if format == "" {
			format = "xlsx"
		}

		if !isValidFileExtension(format) {
			l.Error("invalid format")
			http.Error(w, "invalid format", http.StatusBadRequest)
			return
		}

		ret, err := userRepo.GetAll(ctx, getAllOptions)
		if err != nil {
			l.Error("failed to get individuals", zap.Error(err))
			http.Error(w, "failed to get records: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fileName := generateUniqueDownloadFileNameForCountryAndExtension(selectedCountryID, format)
		downloadFile, err := os.Create(path.Join("/tmp", fileName))
		if err != nil {
			l.Error("failed to create temp file", zap.Error(err))
			http.Error(w, "failed to create temp file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			if err := downloadFile.Close(); err != nil {
				l.Error("failed to close temp file", zap.Error(err))
			}
		}()

		if format == "xlsx" {
			w.Header().Set("Content-Disposition", "attachment; filename=records.xlsx")
			w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

			if err := api.MarshalIndividualsExcel(downloadFile, ret); err != nil {
				l.Error("failed to write xlsx", zap.Error(err))
				http.Error(w, "failed to write xlsx: "+err.Error(), http.StatusInternalServerError)
				return
			}

		}

		if format == "csv" {
			w.Header().Set("Content-Disposition", "attachment; filename=records.csv")
			w.Header().Set("Content-Type", "text/csv")

			if err := api.MarshalIndividualsCSV(downloadFile, ret); err != nil {
				l.Error("failed to write csv", zap.Error(err))
				http.Error(w, "failed to write csv: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		go func() {
			// dirty hack to clear out the file after 30 minutes
			time.Sleep(30 * time.Minute)
			if err := os.Remove(downloadFile.Name()); err != nil {
				l.Error("failed to remove temp file", zap.Error(err))
			}
		}()

		redirectPath := fmt.Sprintf("/countries/%s/individuals/download?file=%s",
			selectedCountryID,
			path.Base(downloadFile.Name()),
		)
		http.Redirect(w, r, redirectPath, http.StatusSeeOther)

	})
}
