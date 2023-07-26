package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
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

func assertValidFileNameForCountry(fileName, wantCountryID string) (string, string, error) {
	parts := strings.Split(fileName, "_")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid file name")
	}
	countryID, err := uuid.Parse(parts[0])
	if err != nil {
		return "", "", fmt.Errorf("invalid file name")
	}
	if countryID.String() != wantCountryID {
		return "", "", fmt.Errorf("invalid file name")
	}

	secondParts := strings.Split(parts[1], ".")
	if len(secondParts) != 2 {
		return "", "", fmt.Errorf("invalid file name")
	}

	ext := secondParts[1]
	if !isValidFileExtension(ext) {
		return "", "", fmt.Errorf("invalid file name")
	}

	_, err = uuid.Parse(secondParts[0])
	if err != nil {
		return "", "", fmt.Errorf("invalid file name")
	}

	return "download." + ext, ext, nil
}

func isValidFileExtension(ext string) bool {
	return ext == "csv" || ext == "xlsx"
}

func setContentTypeForExtension(w http.ResponseWriter, ext string) {
	switch ext {
	case "csv":
		w.Header().Set("Content-Type", "text/csv")
	case "xlsx":
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	}
}

func HandleDownload(
	userRepo db.IndividualRepo,
	azureStorageClient *azblob.Client,
	containerName string,
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
			resultFileName, resultFileExtension, err := assertValidFileNameForCountry(file, selectedCountryID)
			if err != nil {
				l.Error("invalid file name", zap.Error(err))
				http.Error(w, "invalid file name: "+err.Error(), http.StatusBadRequest)
				return
			}

			rangeHeader := r.Header.Get("Range")
			downloadStreamOptions := &azblob.DownloadStreamOptions{}

			if rangeHeader != "" {
				offset, count, err := parseRangeHeader(rangeHeader)
				if err != nil {
					l.Error("invalid range header", zap.Error(err))
					http.Error(w, "invalid range header: "+err.Error(), http.StatusBadRequest)
					return
				}

				downloadStreamOptions.Range = azblob.HTTPRange{
					Offset: offset,
					Count:  count,
				}
			}

			downloadStream, err := azureStorageClient.DownloadStream(ctx, containerName, file, downloadStreamOptions)
			if err != nil {
				l.Error("failed to download file", zap.Error(err))
				http.Error(w, "failed to download file: "+err.Error(), http.StatusInternalServerError)
				return
			}

			b, err := ioutil.ReadAll(downloadStream.Body)
			if err != nil {
				l.Error("failed to read file", zap.Error(err))
				http.Error(w, "failed to read file: "+err.Error(), http.StatusInternalServerError)
				return
			}

			setContentTypeForExtension(w, resultFileExtension)
			http.ServeContent(w, r, resultFileName, time.Time{}, bytes.NewReader(b))
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
			if err := os.Remove(downloadFile.Name()); err != nil {
				l.Error("failed to remove temp file", zap.Error(err))
			}
		}()

		if format == "xlsx" {
			if err := api.MarshalIndividualsExcel(downloadFile, ret); err != nil {
				l.Error("failed to write xlsx", zap.Error(err))
				http.Error(w, "failed to write xlsx: "+err.Error(), http.StatusInternalServerError)
				return
			}

		}

		if format == "csv" {
			if err := api.MarshalIndividualsCSV(downloadFile, ret); err != nil {
				l.Error("failed to write csv", zap.Error(err))
				http.Error(w, "failed to write csv: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		_, err = azureStorageClient.UploadFile(ctx, containerName, fileName, downloadFile, nil)
		if err != nil {
			l.Error("failed to upload file", zap.Error(err))
			http.Error(w, "failed to upload file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		redirectPath := fmt.Sprintf("/countries/%s/participants/download?file=%s",
			selectedCountryID,
			path.Base(fileName),
		)
		http.Redirect(w, r, redirectPath, http.StatusSeeOther)
	})
}

func parseRangeHeader(rangeHeader string) (int64, int64, error) {
	match, _ := regexp.MatchString(`bytes=\d+-\d+`, rangeHeader)
	if !match {
		return 0, 0, fmt.Errorf("invalid range header")
	}
	rangeHeader = strings.ReplaceAll(rangeHeader, "bytes=", "")
	parts := strings.Split(rangeHeader, "-")
	offset, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid range header")
	}
	count, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid range header")
	}
	return offset, count, nil
}
