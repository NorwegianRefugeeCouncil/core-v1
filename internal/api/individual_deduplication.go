package api

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"golang.org/x/exp/slices"
)

func FindDuplicatesInUUIDColumn(df dataframe.DataFrame) []FileError {
	filteredDf := df.Select([]string{indexColumnName, constants.DBColumnIndividualID, constants.DBColumnIndividualLastName})
	fileErrors := []FileError{}

	duplicatesPerId := getDuplicateUUIDs(filteredDf)

	for id := range duplicatesPerId {
		participants := []string{}
		for _, row := range duplicatesPerId[id].Items() {
			participants = append(participants,
				locales.GetTranslator()("error_sharing_uuids_detail",
					filteredDf.Select(constants.DBColumnIndividualLastName).Elem(row, 0).String(),
					row+2,
				),
			)
		}

		fileErrors = append(fileErrors, FileError{
			Message: locales.GetTranslator()("error_sharing_uuids", strings.Join(participants, ", "), id),
		})
	}
	if len(fileErrors) > 0 {
		return fileErrors
	}
	return nil
}

func getDuplicateUUIDs(df dataframe.DataFrame) map[string]containers.Set[int] {
	duplicatesPerId := map[string]containers.Set[int]{}
	for i := 0; i < df.Nrow(); i++ {
		uuid := df.Select(constants.DBColumnIndividualID).Elem(i, 0).String()
		if uuid == "" {
			continue
		}

		duplicates := df.FilterAggregation(dataframe.And,
			dataframe.F{
				Colname:    constants.DBColumnIndividualID,
				Comparando: uuid,
				Comparator: series.In,
			}, dataframe.F{
				Colname:    indexColumnName,
				Comparando: i,
				Comparator: series.Neq,
			})

		for d := 0; d < duplicates.Nrow(); d++ {
			rowNumber, err := duplicates.Select(indexColumnName).Elem(d, 0).Int()
			if err == nil {
				if duplicatesPerId[uuid] != nil {
					duplicatesPerId[uuid].Add(rowNumber)
				} else {
					duplicatesPerId[uuid] = containers.NewSet[int](rowNumber)
				}
			}
		}
	}
	return duplicatesPerId
}

func FindDuplicatesInUpload(config deduplication.DeduplicationConfig, df dataframe.DataFrame) []containers.Set[int] {
	duplicateScores := []containers.Set[int]{}
	for i := 0; i < df.Nrow(); i++ {
		duplicateScores = append(duplicateScores, containers.NewSet[int]())
		getDuplicationScoresForRecord(config, df, i, duplicateScores[i])
	}
	return duplicateScores
}

func getDuplicationScoresForRecord(config deduplication.DeduplicationConfig, df dataframe.DataFrame, currentIndex int, duplicates containers.Set[int]) {
	// the duplicationScore is a metric to determine if the record is a duplicate, it counts how many sub-criteria have been fulfilled
	duplicationScore := make([]int, df.Nrow())

	zeros := []int{}
	for i := 0; i < df.Nrow(); i++ {
		zeros = append(zeros, 0)
	}
	copy(duplicationScore, zeros)

	// e.g. IDs, Names, FullName
	for _, option := range config.Types {
		// the duplicationScoreByType is a metric to determine if the record is a duplicate for the current sub-criterion,
		// it counts how many sub-criteria for the deduplication type have been fulfilled
		duplicationScoreByType := make([]int, df.Nrow())
		copy(duplicationScoreByType, zeros)

		// when the condition is OR, we need to compare all elements in the respective columns to all other elements
		if option.Config.Condition == deduplication.LOGICAL_OPERATOR_OR {
			getOrDuplicationScore(duplicationScoreByType, duplicationScore, df, currentIndex, option)
		} else {
			getAndDuplicationScore(duplicationScore, df, currentIndex, option)
		}
	}
	for r := range duplicationScore {
		if config.Operator == deduplication.LOGICAL_OPERATOR_OR {
			if duplicationScore[r] > 0 {
				duplicates.Add(r)
			}
		} else {
			if duplicationScore[r] == len(config.Types) {
				duplicates.Add(r)
			}
		}
	}
}

func getOrDuplicationScore(scoresByType []int, totalScores []int, df dataframe.DataFrame, currentIndex int, option deduplication.DeduplicationType) {

	// e.g. identification_number_1, identification_number_2, identification_number_3
	for _, column := range option.Config.Columns {
		// the whole column, which we are about to use, including the line number, so we can map scores properly
		thisColumn := df.Select([]string{column, indexColumnName})

		// empty values are not considered duplicates
		currentValue := thisColumn.Elem(currentIndex, 0).String()
		if currentValue == "" {
			continue
		}

		filters := []dataframe.F{}
		// check for duplicates of the current value within its own column
		filters = append(filters, dataframe.F{
			Colname:    column,
			Comparando: currentValue,
			Comparator: series.In,
		})

		// if there are multiple columns to check, we also filter all the other columns
		if len(option.Config.Columns) > 1 {
			for _, c := range option.Config.Columns {
				if c != column {
					filters = append(filters, dataframe.F{
						Colname:    c,
						Comparando: currentValue,
						Comparator: series.In,
					})
				}
			}
		}
		result := df.FilterAggregation(dataframe.Or, filters...)

		result = result.Filter(dataframe.F{
			Colname:    indexColumnName,
			Comparando: currentIndex,
			Comparator: series.Neq,
		})

		for r := 0; r < result.Nrow(); r++ {
			index, err := result.Select(indexColumnName).Elem(r, 0).Int()
			if err == nil {
				scoresByType[index]++
			}
		}
	}

	// if any of the sub-criteria have been fulfilled, this counts as a duplicate for this deduplicationType
	for r := range scoresByType {
		if scoresByType[r] > 0 {
			totalScores[r]++
		}
	}
}

func getAndDuplicationScore(totalScores []int, df dataframe.DataFrame, currentIndex int, option deduplication.DeduplicationType) {
	// we can exclude the current row, to prevent a false positive
	others := df.Filter(dataframe.F{
		Colname:    indexColumnName,
		Comparando: currentIndex,
		Comparator: series.Neq,
	})

	filters := []dataframe.F{}

	// e.g. first_name, middle_name, last_name, native_name
	for _, column := range option.Config.Columns {
		current := df.Select(column).Elem(currentIndex, 0)

		filters = append(filters, dataframe.F{
			Colidx:     0,
			Colname:    column,
			Comparando: current,
			Comparator: series.Eq,
		})
	}

	filteredDf := others.FilterAggregation(dataframe.And, filters...)

	// if there are any rows left, they count as duplicates within the current deduplicationType
	for r := 0; r < filteredDf.Nrow(); r++ {
		ind, err := filteredDf.Select(indexColumnName).Elem(r, 0).Int()
		if err == nil {
			totalScores[ind]++
		}
	}
}

func FormatDbDeduplicationErrors(duplicateMap map[int][]*Individual, individuals []*Individual, config deduplication.DeduplicationConfig) []FileError {
	duplicateErrors := make([]FileError, 0)

	columnNames := make([]string, 0)
	t := locales.GetTranslator()
	for _, deduplicationType := range config.Types {
		for _, column := range deduplicationType.Config.Columns {
			columnNames = append(columnNames, column)
		}
	}

	for originalIndex, duplicates := range duplicateMap {
		for _, ind := range duplicates {
			errorList := make([]error, 0)
			for _, column := range columnNames {
				originalValue, err := individuals[originalIndex].GetFieldValue(column)
				if err != nil {
					log.Fatalln()
				}
				duplicateValue, err := ind.GetFieldValue(column)
				if err != nil {
					log.Fatalln()
				}
				if !(originalValue == "" && duplicateValue == "") {
					errorList = append(errorList, errors.New(t("error_db_duplicate_detail", column, duplicateValue, originalValue)))
				}
			}
			duplicateErrors = append(duplicateErrors, FileError{
				t("error_db_duplicate",
					individuals[originalIndex].LastName,
					originalIndex+2,
					ind.LastName,
					ind.ID,
				),
				errorList,
			})
		}
	}
	return duplicateErrors
}

func FormatFileDeduplicationErrors(duplicateMap []containers.Set[int], individuals []*Individual, config deduplication.DeduplicationConfig) []FileError {
	alertedOn := make(map[int]containers.Set[int])
	uniqueDuplicates := make(map[int]containers.Set[int])
	for i, duplicates := range duplicateMap {
		if alertedOn[i] == nil {
			alertedOn[i] = containers.NewSet[int]()
		}
		for _, duplicateIndex := range duplicates.Items() {
			if alertedOn[duplicateIndex] == nil {
				alertedOn[duplicateIndex] = containers.NewSet[int]()
			}
			if alertedOn[duplicateIndex].Contains(i) || alertedOn[i].Contains(duplicateIndex) {
				continue
			}
			alertedOn[duplicateIndex].Add(i)
			alertedOn[i].Add(duplicateIndex)
			if uniqueDuplicates[i] == nil {
				uniqueDuplicates[i] = containers.NewSet[int]()
			}
			uniqueDuplicates[i].Add(duplicateIndex)
		}
	}

	duplicateErrors := make([]FileError, 0)

	columnNames := make([]string, 0)
	t := locales.GetTranslator()
	for _, deduplicationType := range config.Types {
		for _, column := range deduplicationType.Config.Columns {
			columnNames = append(columnNames, column)
		}
	}

	for originalIndex, duplicates := range uniqueDuplicates {
		for _, duplicateIndex := range duplicates.Items() {
			errorList := make([]error, 0)
			for _, column := range columnNames {
				originalValue, err := individuals[originalIndex].GetFieldValue(column)
				if err != nil {
					log.Fatalln()
				}
				duplicateValue, err := individuals[duplicateIndex].GetFieldValue(column)
				if err != nil {
					log.Fatalln()
				}
				if !(originalValue == "" && duplicateValue == "") {
					errorList = append(errorList, errors.New(t("error_file_duplicate_detail",
						column,
						originalIndex+2,
						originalValue,
						duplicateIndex+2,
						duplicateValue)),
					)
				}
			}
			duplicateErrors = append(duplicateErrors, FileError{
				t("error_file_duplicate",
					individuals[originalIndex].LastName,
					originalIndex+2,
					individuals[duplicateIndex].LastName,
					duplicateIndex+2,
				),
				errorList,
			})
		}
	}
	return duplicateErrors
}

func CreateDataframeFromRecords(records [][]string, deduplicationTypes []deduplication.DeduplicationType, mandatory []string) (dataframe.DataFrame, error) {
	dbCols, err := locales.GetDBColumns(records[0])
	if err != nil {
		return dataframe.DataFrame{}, err
	}

	columnsOfInterest := []string{}
	if len(deduplicationTypes) == 0 && len(mandatory) == 0 {
		return dataframe.DataFrame{}, nil
	}
	for _, deduplicationType := range deduplicationTypes {
		columnsOfInterest = append(columnsOfInterest, deduplicationType.Config.Columns...)
	}
	for _, mandatoryColumn := range mandatory {
		if !slices.Contains(columnsOfInterest, mandatoryColumn) {
			columnsOfInterest = append(columnsOfInterest, mandatoryColumn)
		}
	}

	df := dataframe.LoadRecords(records,
		dataframe.Names(dbCols...),
		dataframe.DetectTypes(true),
		dataframe.DefaultType(series.String),
		dataframe.HasHeader(true),
	).Select(columnsOfInterest)

	if df.Err != nil {
		return dataframe.DataFrame{}, df.Err
	}

	df = AddIndexColumn(df) // adding indices to the records, so we can recognize them in the filtered results

	return df, nil
}

func GetRecordsFromIndividual(deduplicationTypes []deduplication.DeduplicationType, individual *Individual, mandatory []string) [][]string {
	var record [][]string
	var header []string
	var values []string
	for _, dType := range deduplicationTypes {
		for _, field := range dType.Config.Columns {
			header, values = collectValues(individual, field, header, values)
		}
	}
	for _, field := range mandatory {
		if slices.Contains(header, field) {
			continue
		}
		header, values = collectValues(individual, field, header, values)
	}
	record = append(record, header)
	record = append(record, values)
	return record
}

func collectValues(individual *Individual, field string, header []string, values []string) ([]string, []string) {
	v, err := individual.GetFieldValue(field)
	if err != nil {
		return nil, nil
	}
	header = append(header, field)
	switch v.(type) {
	case string:
		values = append(values, v.(string))
	case *time.Time:
		if v.(*time.Time) != nil {
			values = append(values, v.(*time.Time).Format("2006-01-02"))
		} else {
			values = append(values, "")
		}
	}
	return header, values
}
