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

func FindDuplicatesInUpload(optionNames []deduplication.DeduplicationTypeName, records [][]string) map[int]containers.Set[int] {
	df := dataframe.LoadRecords(records,
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.String),
		dataframe.HasHeader(true),
	)

	duplicates := map[int]containers.Set[int]{}
	for i := 0; i < df.Nrow(); i++ {
		duplicates[i] = containers.NewSet[int]()
	}
	for i := 0; i < df.Nrow(); i++ {
		GetDuplicationScoresForRecord(optionNames, df, i, duplicates[i])
		// if all sub-criteria of all deduplicationTypes have been fulfilled, this counts as a duplicate

	}

	return duplicates
}

func GetDuplicationScoresForRecord(optionNames []deduplication.DeduplicationTypeName, df dataframe.DataFrame, currentIndex int, duplicates containers.Set[int]) {
	// the duplicationScore is a metric to determine if the record is a duplicate, it counts how many sub-criteria have been fulfilled
	duplicationScore := make([]int, df.Nrow())

	indexes := []int{}
	zeros := []int{}
	for i := 0; i < df.Nrow(); i++ {
		indexes = append(indexes, i)
		zeros = append(zeros, 0)
	}
	copy(duplicationScore, zeros)
	df = df.Mutate(series.New(indexes, series.String, "index"))

	// e.g. IDs, Names, FullName
	for _, optionName := range optionNames {
		option := deduplication.DeduplicationTypes[optionName]
		// the duplicationScoreByType is a metric to determine if the record is a duplicate for the current sub-criterion,
		// it counts how many sub-criteria for the deduplication type have been fulfilled
		duplicationScoreByType := make([]int, df.Nrow())
		copy(duplicationScoreByType, zeros)

		// when the condition is OR, we need to compare all elements in the respective columns to all other elements
		if option.Config.Condition == deduplication.LOGICAL_OPERATOR_OR {

			// e.g. identification_number_1, identification_number_2, identification_number_3
			for _, column := range option.Config.Columns {
				// the whole column, which we are about to use, including the line number, so we can map scores properly
				thisColumn := df.Select([]string{column, "index"})

				// empty values are not considered duplicates
				currentValue := thisColumn.Elem(currentIndex, 0).String()
				if currentValue == "" {
					continue
				}

				// we can exclude the current row, to prevent a false positive
				otherElements := makeIndexSetWithSkip(thisColumn.Nrow(), currentIndex).Items()

				// check for duplicates of the current value within its own column
				result := thisColumn.Subset(otherElements).Filter(dataframe.F{
					Colname:    column,
					Comparando: currentValue,
					Comparator: series.In,
				})
				for r := 0; r < result.Nrow(); r++ {
					index, err := result.Select("index").Elem(r, 0).Int()
					if err == nil {
						duplicationScoreByType[index]++
					}
				}

				// if there are multiple columns to check, we'll check the next one as well. we're wrapping around the indices, so that all combinations are checked
				// NOTE: this only works this way because all types with OR only have 3 columns at most
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
						index, err := result.Select("index").Elem(r, 0).Int()
						if err == nil {
							duplicationScoreByType[index]++
						}
					}
				}
			}

			// if any of the sub-criteria have been fulfilled, this counts as a duplicate for this deduplicationType
			for r := range duplicationScoreByType {
				if duplicationScoreByType[r] > 0 {
					duplicationScore[r]++
				}
			}
		} else {
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
				ind, err := result.Select("index").Elem(r, 0).Int()
				if err == nil {
					duplicationScore[ind]++
				}
			}
		}
	}
	for r := range duplicationScore {
		if duplicationScore[r] == len(optionNames) {
			duplicates.Add(r)
		}
	}
}

func FormatDbDeduplicationErrors(duplicates []*Individual, deduplicationTypes []deduplication.DeduplicationTypeName) []FileError {
	duplicateErrors := make([]FileError, 0)
	for _, duplicate := range duplicates {
		errorList := make([]error, 0)
		for _, deduplicationType := range deduplicationTypes {
			for _, col := range deduplication.DeduplicationTypes[deduplicationType].Config.Columns {
				val, err := duplicate.GetFieldValue(constants.IndividualDBToFileMap[col])
				if err != nil {
					errorList = append(errorList, errors.New(fmt.Sprintf("Unknown value for %s", col)))
				} else if val != "" {
					errorList = append(errorList, errors.New(fmt.Sprintf("Duplicate value for %s: %s", col, val)))
				}
			}
		}
		duplicateErrors = append(duplicateErrors, FileError{
			fmt.Sprintf("Participant %s has values that are duplicated in your upload", duplicate.ID),
			errorList,
		})
	}
	return duplicateErrors
}

func FormatFileDeduplicationErrors(duplicates map[int]containers.Set[int], deduplicationTypes []deduplication.DeduplicationTypeName, records [][]string) []FileError {
	deduplicationTypesStrings := make([]string, 0)
	for _, deduplicationType := range deduplicationTypes {
		for _, col := range deduplication.DeduplicationTypes[deduplicationType].Config.Columns {
			deduplicationTypesStrings = append(deduplicationTypesStrings, col)
		}
	}

	errors := make([]error, 0)
	for d, duplicate := range duplicates {
		if duplicate.Len() > 0 {
			// we add 2 to the row numbers, 1 for the header, and 1 because excel starts counting at 1
			excelIndices := make([]string, 0)
			for _, dl := range duplicate.Items() {
				excelIndices = append(excelIndices, fmt.Sprintf("%d", dl+2))
			}
			errors = append(errors, fmt.Errorf("Row %d has duplicates in the rows %s", d+2, strings.Join(excelIndices, ", ")))
		}
	}

	if len(errors) == 0 {
		return nil
	}
	return []FileError{{
		fmt.Sprintf("We found the following duplicates when checking for the following columns: %s", strings.Join(deduplicationTypesStrings, ", ")),
		errors}}
}
