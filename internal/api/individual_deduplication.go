package api

import (
	"errors"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"strings"
)

func FindDuplicatesInUUIDColumn(df dataframe.DataFrame) []FileError {
	filteredDf := df.Select([]string{indexColumnName, constants.FileColumnIndividualID, constants.FileColumnIndividualLastName})
	fileErrors := []FileError{}

	duplicatesPerId := getDuplicateUUIDs(filteredDf)

	for id := range duplicatesPerId {
		participants := []string{}
		for _, row := range duplicatesPerId[id].Items() {
			participants = append(participants, fmt.Sprintf("Last name: %s - (Row %d)",
				filteredDf.Select(constants.FileColumnIndividualLastName).Elem(row, 0).String(),
				row+2),
			)
		}

		fileErrors = append(fileErrors, FileError{
			Message: fmt.Sprintf("%s share the same id: %s", strings.Join(participants, ", "), id),
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
		uuid := df.Select(constants.FileColumnIndividualID).Elem(i, 0).String()
		duplicates := ExcludeSelfFromDataframe(df, i).Filter(dataframe.F{
			Colname:    constants.FileColumnIndividualID,
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

func FindDuplicatesInUpload(optionNames []deduplication.DeduplicationTypeName, df dataframe.DataFrame, deduplicationLogicOperator string) []containers.Set[int] {
	duplicateScores := []containers.Set[int]{}
	for i := 0; i < df.Nrow(); i++ {
		duplicateScores = append(duplicateScores, containers.NewSet[int]())
		getDuplicationScoresForRecord(optionNames, df, i, duplicateScores[i], deduplicationLogicOperator)
	}
	return duplicateScores
}

func getDuplicationScoresForRecord(optionNames []deduplication.DeduplicationTypeName, df dataframe.DataFrame, currentIndex int, duplicates containers.Set[int], deduplicationLogicOperator string) {
	// the duplicationScore is a metric to determine if the record is a duplicate, it counts how many sub-criteria have been fulfilled
	duplicationScore := make([]int, df.Nrow())

	zeros := []int{}
	for i := 0; i < df.Nrow(); i++ {
		zeros = append(zeros, 0)
	}
	copy(duplicationScore, zeros)

	// e.g. IDs, Names, FullName
	for _, optionName := range optionNames {
		option := deduplication.DeduplicationTypes[optionName]
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
		if deduplicationLogicOperator == deduplication.LOGICAL_OPERATOR_OR {
			if duplicationScore[r] > 0 {
				duplicates.Add(r)
			}
		} else {
			if duplicationScore[r] == len(optionNames) {
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

func FormatDbDeduplicationErrors(duplicates []*Individual, deduplicationTypes []deduplication.DeduplicationTypeName, df dataframe.DataFrame, deduplicationLogicOperator string) []FileError {
	duplicateErrors := make([]FileError, 0)

	for d := 0; d < len(duplicates); d++ {
		scores := make([]int, df.Nrow())
		databaseValues := map[string]interface{}{}
		errorList := make([]error, 0)
		filteredDf := df

		for _, deduplicationType := range deduplicationTypes {
			filters := []dataframe.F{}
			for _, column := range deduplication.DeduplicationTypes[deduplicationType].Config.Columns {
				value, err := duplicates[d].GetFieldValue(constants.IndividualDBToFileMap[column])
				if err != nil {
					errorList = append(errorList, errors.New(fmt.Sprintf("Unknown value for %s", column)))
				} else if value != "" {
					databaseValues[column] = value
				}
			}
			for column, value := range databaseValues {
				filters = append(filters, dataframe.F{
					Colname:    column,
					Comparando: value,
					Comparator: series.Eq,
				})
			}

			if deduplication.DeduplicationTypes[deduplicationType].Config.Condition == deduplication.LOGICAL_OPERATOR_OR {
				filteredDf = filteredDf.Filter(filters...)
			} else {
				for f := range filters {
					filteredDf = filteredDf.Filter(filters[f])
				}
			}

			for f := 0; f < filteredDf.Nrow(); f++ {
				rowNumber, err := filteredDf.Select(indexColumnName).Elem(f, 0).Int()
				if err == nil {
					scores[rowNumber] = scores[rowNumber] + 1
				}
			}
		}
		for f := 0; f < filteredDf.Nrow(); f++ {
			rowNumber, err := filteredDf.Select(indexColumnName).Elem(f, 0).Int()
			if err == nil {
				if deduplicationLogicOperator == deduplication.LOGICAL_OPERATOR_OR {
					if scores[rowNumber] > 0 {
						for column, dbValue := range databaseValues {
							fileValue := filteredDf.Select(column).Elem(f, 0).String()
							errorList = append(errorList, fmt.Errorf(":: %s :: Database value: %s | File value: %s", column, dbValue, fileValue))
						}
					}
				} else {
					if scores[rowNumber] == len(deduplicationTypes) {
						for column, dbValue := range databaseValues {
							fileValue := filteredDf.Select(column).Elem(f, 0).String()
							errorList = append(errorList, fmt.Errorf(":: %s :: database value: %s | file value: %s", column, dbValue, fileValue))
						}
					}
				}
				if len(errorList) > 0 {
					duplicateErrors = append(duplicateErrors, FileError{
						fmt.Sprintf("Last name %s - Row %d is a duplicate of the participant %s with the id %s",
							filteredDf.Select(constants.FileColumnIndividualLastName).Elem(f, 0),
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

func FormatFileDeduplicationErrors(duplicateMap []containers.Set[int], deduplicationTypes []deduplication.DeduplicationTypeName, records [][]string, columnMapping map[string]int) []FileError {
	duplicateErrors := make([]FileError, 0)
	columnNames := make([]string, 0)
	for _, deduplicationType := range deduplicationTypes {
		for _, column := range deduplication.DeduplicationTypes[deduplicationType].Config.Columns {
			columnNames = append(columnNames, column)
		}
	}

	for originalIndex, duplicates := range duplicateMap {
		for _, duplicateIndex := range duplicates.Items() {
			errorList := make([]error, 0)
			for _, column := range columnNames {
				originalValue := records[originalIndex+1][columnMapping[column]]
				duplicateValue := records[duplicateIndex+1][columnMapping[column]]
				if !(originalValue == "" && duplicateValue == "") {
					errorList = append(errorList, fmt.Errorf(":: %s :: Row %d: %s | Row %d: %s",
						column,
						originalIndex+2,
						originalValue,
						duplicateIndex+2,
						duplicateValue),
					)
				}
			}
			duplicateErrors = append(duplicateErrors, FileError{
				fmt.Sprintf("Last name %s - Row %d and Last name: %s - Row %d in your file are duplicates",
					records[originalIndex+1][columnMapping[constants.FileColumnIndividualLastName]],
					originalIndex+2,
					records[duplicateIndex+1][columnMapping[constants.FileColumnIndividualLastName]],
					duplicateIndex+2,
				),
				errorList,
			})
		}
	}
	return duplicateErrors
}
