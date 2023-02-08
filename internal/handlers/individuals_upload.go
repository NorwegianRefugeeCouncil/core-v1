package handlers

import (
	"errors"
	"fmt"
	"github.com/nrc-no/notcore/internal/constants"
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

var UPLOAD_LIMIT = 10000

func HandleUpload(renderer Renderer, individualRepo db.IndividualRepo) http.Handler {

	const (
		templateName               = "error.gohtml"
		formParamFile              = "file"
		formParamDeduplicationType = "deduplicationType"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			l   = logging.NewLogger(ctx)
		)

		renderError := func(title string, fileErrors []api.FileError) {
			renderer.RenderView(w, r, templateName, map[string]interface{}{
				"Errors": fileErrors,
				"Title":  title,
			})
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
			renderError("failed to parse input file: "+err.Error(), nil)
			return
		}

		var individuals []*api.Individual
		var fields []string

		fileErrors, records, err := api.UnmarshalIndividualsFile(filename, formFile, &individuals, &fields)
		if err != nil {
			l.Error("failed to parse file", zap.Error(err))
			renderError("Failed to parse input file: "+err.Error(), nil)
		}
		if fileErrors != nil {
			renderError("Could not parse uploaded file", fileErrors)
			return
		}

		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			renderError("Could not detect selected country. Please select a country from the dropdown.", nil)
			return
		}

		fieldSet := containers.NewStringSet(fields...)
		fieldSet.Add("country_id")

		var individualIds = containers.NewStringSet()
		for _, individual := range individuals {
			individual.CountryID = selectedCountryID
			if len(individual.ID) > 0 {
				individualIds.Add(individual.ID)
			}
		}

		existingIndividuals, err := individualRepo.GetAll(ctx, api.ListIndividualsOptions{IDs: individualIds, CountryID: selectedCountryID})
		if err != nil {
			l.Error("failed to get existing individuals", zap.Error(err))
			renderError("Could not load list of participants: "+err.Error(), nil)
			return
		}

		invalidIndividualIds := validateIndividualsExistInCountry(individualIds, existingIndividuals, selectedCountryID)
		if len(invalidIndividualIds) > 0 {
			l.Warn("user trying to update individuals that don't exist or are in the wrong country", zap.Strings("individual_ids", invalidIndividualIds))
			renderError(fmt.Sprintf("Could not update participants %s, they do not exist in the database for the selected country.", strings.Join(invalidIndividualIds, ",")), nil)
			return
		}

		deduplicationTypes := r.MultipartForm.Value[formParamDeduplicationType]

		// TODO find duplicates within the file
		if len(deduplicationTypes) > 0 {

			structProperties := make([]string, 0)
			fileColumns := make([]string, 0)
			for _, d := range deduplicationTypes {
				dt, ok := db.ParseString(d)
				if ok {
					for _, vc := range db.DeduplicationOptions[dt].Value.Columns {
						fileColumns = append(fileColumns, constants.IndividualDBToFileMap[vc])
						structProperties = append(structProperties, constants.IndividualFileToStructMap[vc])
					}
				} else {
					l.Error("invalid deduplication type", zap.String("deduplication_type", d))
					renderError("Invalid deduplication type: "+d, nil)
					return
				}
			}

			duplicatesInFile := api.FindDuplicatesInUpload(fileColumns, records)

			if duplicatesInFile.Len() > 0 {
				errors := make([]error, 0)
				for _, duplicate := range duplicatesInFile.Items() {
					errors = append(errors, fmt.Errorf(duplicate))
				}
				renderError(
					"Found duplicates within your uploaded file: ",
					[]api.FileError{{
						Message: fmt.Sprintf("When checking for duplicates in the columns %s, we found the following duplicated values: ", deduplicationTypes),
						Err:     errors}},
				)
				return
			}

			duplicates, err := individualRepo.FindDuplicates(r.Context(), individuals, deduplicationTypes)
			if err != nil {
				renderError("An error occurred while trying to check for duplicates: "+err.Error(), nil)
				return
			}

			// TODO make this a function
			if len(duplicates) > 0 || duplicatesInFile.Len() > 0 {
				duplicateErrors := make([]api.FileError, 0, len(duplicates))
				for _, duplicate := range duplicates {
					errorList := make([]error, 0)
					for _, field := range deduplicationTypes {
						f, ok := db.ParseString(field)
						if !ok {
							renderError("Can not deduplicate by "+field, nil)
						}
						for _, col := range db.DeduplicationOptions[f].Value.Columns {
							val, err := duplicate.GetFieldValue(constants.IndividualDBToFileMap[col])
							if err != nil {
								errorList = append(errorList, errors.New(fmt.Sprintf("Unknown value for %s", col)))
							} else if val != "" {
								errorList = append(errorList, errors.New(fmt.Sprintf("Duplicate value for %s: %s", col, val)))
							}
						}
					}
					duplicateErrors = append(duplicateErrors, api.FileError{
						Message: fmt.Sprintf("Participant %s has values that are duplicated in your upload", duplicate.ID),
						Err:     errorList,
					})
				}
				renderError(fmt.Sprintf("%d duplicates found in database", len(duplicates)), duplicateErrors)
				return
			}
		}

		_, err = individualRepo.PutMany(r.Context(), individuals, fieldSet)
		if err != nil {
			l.Error("failed to put individuals", zap.Error(err))
			renderError("Could not upload participant data: "+err.Error(), nil)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/countries/%s/participants", selectedCountryID), http.StatusSeeOther)

		return
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
