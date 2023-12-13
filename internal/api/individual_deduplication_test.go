package api

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testRecords = [][]string{
	{
		"index",
		constants.DBColumnIndividualIdentificationNumber1,
		constants.DBColumnIndividualIdentificationNumber2,
		constants.DBColumnIndividualIdentificationNumber3,
		constants.DBColumnIndividualFirstName,
		constants.DBColumnIndividualMiddleName,
		constants.DBColumnIndividualLastName,
		constants.DBColumnIndividualNativeName,
		constants.DBColumnIndividualFullName,
	},
	{"0", "1", "2", "3", "A", "B", "C", "D", "E"},
	{"1", "4", "5", "6", "F", "G", "H", "I", "J"},
	{"2", "1", "", "", "A", "B", "C", "", ""},
	{"3", "2", "", "", "A", "B", "", "", ""},
	{"4", "3", "", "", "A", "B", "", "", ""},
	{"5", "", "5", "", "A", "Z", "C", "", "E"},
	{"6", "3", "2", "", "A", "B", "C", "D", ""},
	{"7", "1", "2", "3", "A", "B", "C", "D", "K"},
	{"8", "", "", "5", "A", "B", "C", "D", ""},
	{"9", "7", "", "9", "A", "", "", "", ""},
}

var df = dataframe.LoadRecords(testRecords)

func TestGetDuplicationScoresForRecord(t *testing.T) {
	tests := []struct {
		name                       string
		deduplicationTypes         []deduplication.DeduplicationTypeName
		deduplicationLogicOperator string
		records                    dataframe.DataFrame
		index                      int
		want                       containers.Set[int]
	}{
		{
			name:                       "check IDs, AND",
			records:                    df,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			index:                      0,
			want:                       containers.NewSet[int](2, 3, 4, 6, 7),
		},
		{
			name:                       "check Names, AND",
			records:                    df,
			index:                      0,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want:                       containers.NewSet[int](6, 7, 8),
		},
		{
			name:                       "check Names and IDs, AND",
			records:                    df,
			index:                      0,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want:                       containers.NewSet[int](6, 7),
		},
		{
			name:                       "check Full Name, AND",
			records:                    df,
			index:                      0,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want:                       containers.NewSet[int](5),
		},
		{
			name:                       "check Names and IDs and Full Name, AND",
			records:                    df,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want:                       containers.NewSet[int](),
		},
		{
			name:                       "check IDs, OR",
			records:                    df,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			index:                      0,
			want:                       containers.NewSet[int](2, 3, 4, 6, 7),
		},
		{
			name:                       "check Names, OR",
			records:                    df,
			index:                      0,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want:                       containers.NewSet[int](6, 7, 8),
		},
		{
			name:                       "check Names and IDs, OR",
			records:                    df,
			index:                      0,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want:                       containers.NewSet[int](2, 3, 4, 6, 7, 8),
		},
		{
			name:                       "check Full Name, OR",
			records:                    df,
			index:                      0,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want:                       containers.NewSet[int](5),
		},
		{
			name:                       "check Names and IDs and Full Name, OR",
			records:                    df,
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want:                       containers.NewSet[int](2, 3, 4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicates := []containers.Set[int]{0: containers.NewSet[int]()}
			getDuplicationScoresForRecord(tt.deduplicationTypes, tt.records, tt.index, duplicates[0], tt.deduplicationLogicOperator)
			assert.Equal(t, tt.want, duplicates[0])
		})
	}
}

func TestFindDuplicatesInUpload(t *testing.T) {
	tests := []struct {
		name                       string
		deduplicationTypes         []deduplication.DeduplicationTypeName
		deduplicationLogicOperator string
		want                       []containers.Set[int]
	}{
		{
			name:                       "check IDs, AND",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want: []containers.Set[int]{
				0: containers.NewSet[int](2, 3, 4, 6, 7),
				1: containers.NewSet[int](5, 8),
				2: containers.NewSet[int](0, 7),
				3: containers.NewSet[int](0, 6, 7),
				4: containers.NewSet[int](0, 6, 7),
				5: containers.NewSet[int](1, 8),
				6: containers.NewSet[int](0, 3, 4, 7),
				7: containers.NewSet[int](0, 2, 3, 4, 6),
				8: containers.NewSet[int](1, 5),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Names, AND",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want: []containers.Set[int]{
				0: containers.NewSet[int](6, 7, 8),
				1: containers.NewSet[int](),
				2: containers.NewSet[int](),
				3: containers.NewSet[int](4),
				4: containers.NewSet[int](3),
				5: containers.NewSet[int](),
				6: containers.NewSet[int](0, 7, 8),
				7: containers.NewSet[int](0, 6, 8),
				8: containers.NewSet[int](0, 6, 7),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Names and IDs, AND",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want: []containers.Set[int]{
				0: containers.NewSet[int](6, 7),
				1: containers.NewSet[int](),
				2: containers.NewSet[int](),
				3: containers.NewSet[int](),
				4: containers.NewSet[int](),
				5: containers.NewSet[int](),
				6: containers.NewSet[int](0, 7),
				7: containers.NewSet[int](0, 6),
				8: containers.NewSet[int](),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Full Name, AND",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want: []containers.Set[int]{
				0: containers.NewSet[int](5),
				1: containers.NewSet[int](),
				2: containers.NewSet[int](),
				3: containers.NewSet[int](),
				4: containers.NewSet[int](),
				5: containers.NewSet[int](0),
				6: containers.NewSet[int](),
				7: containers.NewSet[int](),
				8: containers.NewSet[int](),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Names and IDs and Full Name, AND",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_AND,
			want: []containers.Set[int]{
				0: containers.NewSet[int](),
				1: containers.NewSet[int](),
				2: containers.NewSet[int](),
				3: containers.NewSet[int](),
				4: containers.NewSet[int](),
				5: containers.NewSet[int](),
				6: containers.NewSet[int](),
				7: containers.NewSet[int](),
				8: containers.NewSet[int](),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check IDs, OR",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want: []containers.Set[int]{
				0: containers.NewSet[int](2, 3, 4, 6, 7),
				1: containers.NewSet[int](5, 8),
				2: containers.NewSet[int](0, 7),
				3: containers.NewSet[int](0, 6, 7),
				4: containers.NewSet[int](0, 6, 7),
				5: containers.NewSet[int](1, 8),
				6: containers.NewSet[int](0, 3, 4, 7),
				7: containers.NewSet[int](0, 2, 3, 4, 6),
				8: containers.NewSet[int](1, 5),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Names, OR",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want: []containers.Set[int]{
				0: containers.NewSet[int](6, 7, 8),
				1: containers.NewSet[int](),
				2: containers.NewSet[int](),
				3: containers.NewSet[int](4),
				4: containers.NewSet[int](3),
				5: containers.NewSet[int](),
				6: containers.NewSet[int](0, 7, 8),
				7: containers.NewSet[int](0, 6, 8),
				8: containers.NewSet[int](0, 6, 7),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Names and IDs, OR",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want: []containers.Set[int]{
				0: containers.NewSet[int](2, 3, 4, 6, 7, 8),
				1: containers.NewSet[int](5, 8),
				2: containers.NewSet[int](0, 7),
				3: containers.NewSet[int](0, 4, 6, 7),
				4: containers.NewSet[int](0, 3, 6, 7),
				5: containers.NewSet[int](1, 8),
				6: containers.NewSet[int](0, 3, 4, 7, 8),
				7: containers.NewSet[int](0, 2, 3, 4, 6, 8),
				8: containers.NewSet[int](0, 1, 5, 6, 7),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Full Name, OR",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want: []containers.Set[int]{
				0: containers.NewSet[int](5),
				1: containers.NewSet[int](),
				2: containers.NewSet[int](),
				3: containers.NewSet[int](),
				4: containers.NewSet[int](),
				5: containers.NewSet[int](0),
				6: containers.NewSet[int](),
				7: containers.NewSet[int](),
				8: containers.NewSet[int](),
				9: containers.NewSet[int](),
			},
		},
		{
			name:                       "check Names and IDs and Full Name, OR",
			deduplicationTypes:         []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameFullName},
			deduplicationLogicOperator: deduplication.LOGICAL_OPERATOR_OR,
			want: []containers.Set[int]{
				0: containers.NewSet[int](2, 3, 4, 5, 6, 7, 8),
				1: containers.NewSet[int](5, 8),
				2: containers.NewSet[int](0, 7),
				3: containers.NewSet[int](0, 4, 6, 7),
				4: containers.NewSet[int](0, 3, 6, 7),
				5: containers.NewSet[int](0, 1, 8),
				6: containers.NewSet[int](0, 3, 4, 7, 8),
				7: containers.NewSet[int](0, 2, 3, 4, 6, 8),
				8: containers.NewSet[int](0, 1, 5, 6, 7),
				9: containers.NewSet[int](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicates := FindDuplicatesInUpload(tt.deduplicationTypes, df, tt.deduplicationLogicOperator)
			assert.Equal(t, tt.want, duplicates)
		})
	}
}

func TestFindDuplicatesInUUIDColumn(t *testing.T) {
	tests := []struct {
		name    string
		records [][]string
		want    map[string]containers.Set[int]
	}{
		{
			name: "",
			records: [][]string{
				{"index", "id"},
				{"0", "1"},
				{"1", "2"},
				{"2", "2"},
				{"3", "2"},
				{"4", "5"},
				{"5", "1"},
				{"6", "5"},
				{"7", "8"},
			},
			want: map[string]containers.Set[int]{
				"1": containers.NewSet[int](0, 5),
				"2": containers.NewSet[int](1, 2, 3),
				"5": containers.NewSet[int](4, 6),
			},
		},
	}
	for _, tt := range tests {
		var testDf = dataframe.LoadRecords(tt.records)
		t.Run(tt.name, func(t *testing.T) {
			duplicates := getDuplicateUUIDs(testDf)
			assert.Equal(t, tt.want, duplicates)
		})
	}
}
