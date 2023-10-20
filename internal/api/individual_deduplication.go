package api

import (
	"errors"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"golang.org/x/exp/slices"
	"strings"
	"time"
)

func FindDuplicatesInUUIDColumn(df dataframe.DataFrame) []FileError {
	filteredDf := df.Select([]string{indexColumnName, constants.DBColumnIndividualID, constants.DBColumnIndividualLastName})
	fileErrors := []FileError{}
	df.Col(indexColumnName).Records()

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
		duplicates := ExcludeSelfFromDataframe(df, i).Filter(dataframe.F{
			Colname:    constants.DBColumnIndividualID,
			Comparando: uuid,
			Comparator: series.In,
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

		newDf := ExcludeSelfFromDataframe(thisColumn, currentIndex)

		// check for duplicates of the current value within its own column
		result := newDf.Filter(dataframe.F{
			Colname:    column,
			Comparando: currentValue,
			Comparator: series.In,
		})
		for r := 0; r < result.Nrow(); r++ {
			index, err := result.Select(indexColumnName).Elem(r, 0).Int()
			if err == nil {
				scoresByType[index]++
			}
		}

		//df.Select([]string{column, indexColumnName}).

		// if there are multiple columns to check, we also filter all the other columns
		if len(option.Config.Columns) > 1 {
			filters := []dataframe.F{}
			for _, c := range option.Config.Columns {
				if c != column {
					filters = append(filters, dataframe.F{
						Colname:    c,
						Comparando: currentValue,
						Comparator: series.In,
					})
				}
			}
			result = df.Filter(filters...)

			for r := 0; r < result.Nrow(); r++ {
				index, err := result.Select(indexColumnName).Elem(r, 0).Int()
				if err == nil {
					scoresByType[index]++
				}
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
	otherElements := makeIndexSetWithSkip(df.Nrow(), currentIndex).Items()
	// we're chaining the filters, such that only complete combinations of values are counted
	result := df.Subset(otherElements).Filter()

	// e.g. first_name, middle_name, last_name, native_name
	for _, column := range option.Config.Columns {
		current := df.Select(column).Elem(currentIndex, 0)
		result = result.Filter(dataframe.F{
			Colidx:     0,
			Colname:    column,
			Comparando: current,
			Comparator: series.Eq,
		})
	}
	// if there are any rows left, they count as duplicates within the current deduplicationType
	for r := 0; r < result.Nrow(); r++ {
		ind, err := result.Select(indexColumnName).Elem(r, 0).Int()
		if err == nil {
			totalScores[ind]++
		}
	}
}

type AggregationType struct {
	Aggregation dataframe.Aggregation
	Filters     []dataframe.F
}

func FormatDbDeduplicationErrors(duplicates []*Individual, df dataframe.DataFrame, config deduplication.DeduplicationConfig) []FileError {
	duplicateErrors := make([]FileError, 0)
	t := locales.GetTranslator()

	for d := 0; d < len(duplicates); d++ {
		scores := make([]int, df.Nrow())
		databaseValues := map[string]interface{}{}
		filteredDf := df
		at := []AggregationType{}

		for _, deduplicationType := range config.Types {
			for _, column := range deduplicationType.Config.Columns {
				value, err := duplicates[d].GetFieldValue(constants.IndividualFileToDBMap[column])
				if err == nil {

					switch value.(type) {
					case string:
						if value.(string) != "" {
							databaseValues[column] = value.(string)
						}
					case *time.Time:
						if value.(*time.Time) != nil {
							databaseValues[column] = value.(*time.Time).Format("2006-01-02")
						}
					}
				}
			}
		}
		for _, deduplicationType := range config.Types {
			filters := []dataframe.F{}

			for column, value := range databaseValues {
				filters = append(filters, dataframe.F{
					Colname:    column,
					Comparando: value,
					Comparator: series.Eq,
				})
			}
			if deduplicationType.Config.Condition == deduplication.LOGICAL_OPERATOR_OR {
				at = append(at, AggregationType{
					Aggregation: dataframe.Or,
					Filters:     filters,
				})
			} else {
				at = append(at, AggregationType{
					Aggregation: dataframe.And,
					Filters:     filters,
				})
			}

			if deduplicationType.Config.Condition == deduplication.LOGICAL_OPERATOR_OR {
				filteredDf = filteredDf.FilterAggregation(dataframe.Or, filters...)
			} else {
				filteredDf = filteredDf.FilterAggregation(dataframe.And, filters...)
			}

			for f := 0; f < filteredDf.Nrow(); f++ {
				rowNumber, err := filteredDf.Select(indexColumnName).Elem(f, 0).Int()
				if err == nil {
					scores[rowNumber] = scores[rowNumber] + 1
				}
			}
			if config.Operator == deduplication.LOGICAL_OPERATOR_OR {
				if filteredDf.Nrow() > 0 {
					break
				} else {
					filteredDf = df
				}
			}
		}

		for f := 0; f < filteredDf.Nrow(); f++ {
			errorList := make([]error, 0)
			rowNumber, err := filteredDf.Select(indexColumnName).Elem(f, 0).Int()
			if err == nil {
				if config.Operator == deduplication.LOGICAL_OPERATOR_OR {
					if scores[rowNumber] > 0 {
						for column, dbValue := range databaseValues {
							fileValue := filteredDf.Select(column).Elem(f, 0).String()
							errorList = append(errorList, errors.New(t("error_db_duplicate_detail", column, dbValue, fileValue)))
						}
					}
				} else {
					if scores[rowNumber] == len(config.Types) {
						for column, dbValue := range databaseValues {
							fileValue := filteredDf.Select(column).Elem(f, 0).String()
							errorList = append(errorList, errors.New(t("error_db_duplicate_detail", column, dbValue, fileValue)))
						}
					}
				}

				if len(errorList) > 0 {
					duplicateErrors = append(duplicateErrors, FileError{
						t("error_db_duplicate",
							filteredDf.Select(constants.IndividualFileToDBMap[constants.FileColumnIndividualLastName]).Elem(f, 0),
							rowNumber+2,
							duplicates[d].LastName,
							duplicates[d].ID,
						),
						errorList,
					})
				}
			}
		}
	}
	return duplicateErrors
}

func FormatFileDeduplicationErrors(duplicateMap []containers.Set[int], config deduplication.DeduplicationConfig, records [][]string, columnMapping map[string]int) []FileError {
	duplicateErrors := make([]FileError, 0)
	alertedOn := containers.Set[int]{}
	columnNames := make([]string, 0)
	t := locales.GetTranslator()
	for _, deduplicationType := range config.Types {
		for _, column := range deduplicationType.Config.Columns {
			columnNames = append(columnNames, column)
		}
	}

	for originalIndex, duplicates := range duplicateMap {
		for _, duplicateIndex := range duplicates.Items() {
			if alertedOn.Contains(duplicateIndex) {
				continue
			}
			errorList := make([]error, 0)
			for _, column := range columnNames {
				originalValue := records[originalIndex+1][columnMapping[column]]
				duplicateValue := records[duplicateIndex+1][columnMapping[column]]
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
					records[originalIndex+1][columnMapping[constants.FileColumnIndividualLastName]],
					originalIndex+2,
					records[duplicateIndex+1][columnMapping[constants.FileColumnIndividualLastName]],
					duplicateIndex+2,
				),
				errorList,
			})
		}
		alertedOn.Add(originalIndex)
		alertedOn.Add(duplicates.Items()...)
	}
	return duplicateErrors
}

func GetDataframeFromRecords(records [][]string, deduplicationTypes []deduplication.DeduplicationType) (dataframe.DataFrame, error) {
	columnsOfInterest := []string{}
	if slices.Contains(records[0], constants.FileColumnIndividualID) {
		columnsOfInterest = append(columnsOfInterest, constants.FileColumnIndividualID)
	}
	for _, deduplicationType := range deduplicationTypes {
		columnsOfInterest = append(columnsOfInterest, deduplicationType.Config.Columns...)
	}

	df := dataframe.LoadRecords(records,
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.String),
		dataframe.HasHeader(true),
	).Select(columnsOfInterest)

	if df.Err != nil {
		return dataframe.DataFrame{}, df.Err
	}
	return df, nil
}

func GetRecordsFromIndividual(deduplicationTypes []deduplication.DeduplicationType, individual *Individual) [][]string {
	var record [][]string
	var header []string
	var values []string
	for _, dType := range deduplicationTypes {
		for _, field := range dType.Config.Columns {
			value, err := individual.GetFieldValue(field)
			if err != nil {
				continue
			}
			header = append(header, field)
			values = append(values, value.(string))
		}
	}
	record = append(record, header)
	record = append(record, values)
	return record
}
