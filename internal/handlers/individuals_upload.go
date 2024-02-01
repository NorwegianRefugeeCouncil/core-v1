package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"go.uber.org/zap"
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
			t   = locales.GetTranslator()
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
			renderError(t("error_failed_to_parse_file_v0", err.Error()), nil)
			return
		}

		var individuals []*api.Individual
		var fields []string
		var records [][]string
		fileErrors := []api.FileError{}

		err = api.UnmarshallRecordsFromFile(&records, formFile, filename)
		if err != nil {
			l.Error("failed to parse file", zap.Error(err))
			renderError(t("error_failed_to_parse_file_v0", err.Error()), nil)
			return
		}

		colMapping, fileErrors := api.GetColumnMapping(records[0], &fields)

		if fileErrors != nil {
			renderError(t("error_failed_to_parse_file"), fileErrors)
			return
		}

		fileErrors = api.UnmarshalIndividualsTabularData(records, &individuals, colMapping, &UPLOAD_LIMIT)

		if fileErrors != nil {
			renderError(t("error_failed_to_parse_file"), fileErrors)
			return
		}

		deduplicationTypes := r.MultipartForm.Value[formParamDeduplicationType]
		deduplicationLogicOperator := deduplication.LogicOperator(r.MultipartForm.Value[formParamDeduplicationLogicOperator][0])
		deduplicationConfig, err := deduplication.GetDeduplicationConfig(deduplicationTypes, deduplicationLogicOperator)

		mandatoryColumns := []string{constants.DBColumnIndividualLastName}
		var idColumnExistsInFile bool
		if _, idColumnExistsInFile = colMapping[constants.DBColumnIndividualID]; idColumnExistsInFile {
			mandatoryColumns = append(mandatoryColumns, constants.DBColumnIndividualID)
		}

		df, err := api.CreateDataframeFromRecords(records, deduplicationConfig.Types, mandatoryColumns)
		if err != nil {
			l.Error("failed to get dataframe from records", zap.Error(err))
			renderError(t("error_failed_to_parse_file"), []api.FileError{
				{
					Message: t("error_deduplication_preparation"),
					Err:     []error{err},
				},
			})
			return
		}

		if idColumnExistsInFile {
			fileErrors = api.FindDuplicatesInUUIDColumn(df)
			if fileErrors != nil {
				renderError(t("error_file_with_duplicate_uuids"), fileErrors)
				return
			}
		}

		selectedCountryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country id", zap.Error(err))
			renderError(t("error_no_selected_country"), nil)
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
			renderError(t("error_load_participants", err.Error()), nil)
			return
		}

		invalidIndividualIds := validateIndividualsExistInCountry(individualIds, existingIndividuals, selectedCountryID)

		if len(invalidIndividualIds) > 0 {
			l.Warn("user trying to update individuals that don't exist or are in the wrong country", zap.Strings("individual_ids", invalidIndividualIds))
			renderError(t("error_nonexistent_participant", strings.Join(invalidIndividualIds, ",")), nil)
			return
		}

		if len(deduplicationConfig.Types) > 0 {
			duplicatesScores := api.FindDuplicatesInUpload(deduplicationConfig, df)
			errors := api.FormatFileDeduplicationErrors(duplicatesScores, deduplicationConfig, records, colMapping)
			if len(errors) > 0 {
				if errors != nil {
					renderError(t("error_found_duplicates_in_file", len(errors)), errors)
					return
				}
			}

			duplicatesInDB, err := individualRepo.FindDuplicates(ctx, df, deduplicationConfig)
			if err != nil {
				renderError(t("error_deduplication_fail", err.Error()), nil)
				return
			}

			dbDuplicationErrors := api.FormatDbDeduplicationErrors(duplicatesInDB, df, deduplicationConfig)
			if len(dbDuplicationErrors) > 0 {
				renderError(t("error_found_duplicates_in_db", len(dbDuplicationErrors)), dbDuplicationErrors)
				return
			}
		}

		_, err = individualRepo.PutMany(r.Context(), individuals, fieldSet)
		if err != nil {
			l.Error("failed to put individuals", zap.Error(err))
			renderError(t("error_upload_fail", err.Error()), nil)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/countries/%s/participants", selectedCountryID), http.StatusSeeOther)

		return
	})
}
