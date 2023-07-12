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
		templateName                        = "error.gohtml"
		formParamFile                       = "file"
		formParamDeduplicationType          = "deduplicationType"
		formParamDeduplicationLogicOperator = "deduplicationLogicOperator"
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
		fileErrors := []api.FileError{}

		err = api.UnmarshallRecordsFromFile(&records, formFile, filename)
		if err != nil {
			l.Error("failed to parse file", zap.Error(err))
			renderError("Failed to parse input file: "+err.Error(), nil)
			return
		}

		colMapping, fileErrors := api.GetColumnMapping(records, &fields)

		if fileErrors != nil {
			renderError("Could not parse uploaded file", fileErrors)
			return
		}

		fileErrors = api.UnmarshalIndividualsTabularData(records, &individuals, colMapping, &UPLOAD_LIMIT)

		if fileErrors != nil {
			renderError("Could not parse uploaded file", fileErrors)
			return
		}

		df := api.GetDataframeFromRecords(records)
		df = api.AddIndexColumn(df) // adding indices to the records, so we can recognize them in the filtered results

		fileErrors = api.FindDuplicatesInUUIDColumn(df)
		if fileErrors != nil {
			renderError("Could not parse uploaded file due to duplicates in the id column", fileErrors)
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
		deduplicationLogicOperator := r.MultipartForm.Value[formParamDeduplicationLogicOperator]

		if len(deduplicationTypes) > 0 {
			optionNames, err := deduplication.GetDeduplicationTypeNames(deduplicationTypes)
			if err != nil {
				l.Error("invalid deduplication type", zap.String("deduplication_type", strings.Join(deduplicationTypes, ",")), zap.Error(err))
				renderError(fmt.Sprintf("Invalid deduplication type: %s", strings.Join(deduplicationTypes, ",")), nil)
				return
			}

			duplicatesScores := api.FindDuplicatesInUpload(optionNames, df, deduplicationLogicOperator[0])
			errors := api.FormatFileDeduplicationErrors(duplicatesScores, optionNames, records, colMapping)
			if len(errors) > 0 {
				if errors != nil {
					renderError(fmt.Sprintf("Found %d duplicates within your uploaded file: ", len(errors)), errors)
					return
				}
			}

			duplicatesInDB, err := individualRepo.FindDuplicates(ctx, individuals, optionNames, deduplicationLogicOperator[0])
			if err != nil {
				renderError("An error occurred while trying to check for duplicates: "+err.Error(), nil)
				return
			}

			dbDuplicationErrors := api.FormatDbDeduplicationErrors(duplicatesInDB, optionNames, df, deduplicationLogicOperator[0])
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
