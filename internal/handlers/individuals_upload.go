package handlers

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"go.uber.org/zap"
	"net/http"
	"strings"
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
		var records [][]string

		err = api.UnmarshallRecordsFromFile(&records, formFile, filename)
		if err != nil {
			l.Error("failed to parse file", zap.Error(err))
			renderError("Failed to parse input file: "+err.Error(), nil)
			return
		}

		fileErrors := api.UnmarshalIndividualsTabularData(records, &individuals, &fields, &UPLOAD_LIMIT)

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

		if len(deduplicationTypes) > 0 {
			optionNames, err := deduplication.GetDeduplicationTypeNames(deduplicationTypes)
			if err != nil {
				l.Error("invalid deduplication type", zap.String("deduplication_type", strings.Join(deduplicationTypes, ",")), zap.Error(err))
				renderError(fmt.Sprintf("Invalid deduplication type: %s", strings.Join(deduplicationTypes, ",")), nil)
				return
			}

			duplicatesInFile := api.FindDuplicatesInUpload(optionNames, records)
			if len(duplicatesInFile) > 0 {
				errors := api.FormatFileDeduplicationErrors(duplicatesInFile, optionNames, records)
				if errors != nil {
					renderError("Found duplicates within your uploaded file: ", errors)
					return
				}
			}

			duplicatesInDB, err := individualRepo.FindDuplicates(ctx, individuals, optionNames)
			if err != nil {
				renderError("An error occurred while trying to check for duplicates: "+err.Error(), nil)
				return
			}

			dbDuplicationErrors := api.FormatDbDeduplicationErrors(duplicatesInDB, optionNames)
			if len(dbDuplicationErrors) > 0 {
				renderError(fmt.Sprintf("%d duplicates found in database", len(dbDuplicationErrors)), dbDuplicationErrors)
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
