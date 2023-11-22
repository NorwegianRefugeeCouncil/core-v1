package api

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
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
		name    string
		config  deduplication.DeduplicationConfig
		records dataframe.DataFrame
		index   int
		want    containers.Set[int]
	}{
		{
			name:    "check IDs, AND",
			records: df,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
			index: 0,
			want:  containers.NewSet[int](2, 3, 4, 6, 7),
		},
		{
			name:    "check Names, AND",
			records: df,
			index:   0,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames]},
			},
			want: containers.NewSet[int](6, 7, 8),
		},
		{
			name:    "check Names and IDs, AND",
			records: df,
			index:   0,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
			want: containers.NewSet[int](6, 7),
		},
		{
			name:    "check Full Name, AND",
			records: df,
			index:   0,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName]},
			},
			want: containers.NewSet[int](5),
		},
		{
			name:    "check Names and IDs and Full Name, AND",
			records: df,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			want: containers.NewSet[int](),
		},
		{
			name:    "check IDs, OR",
			records: df,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds]},
			},
			index: 0,
			want:  containers.NewSet[int](2, 3, 4, 6, 7),
		},
		{
			name:    "check Names, OR",
			records: df,
			index:   0,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames]},
			},
			want: containers.NewSet[int](6, 7, 8),
		},
		{
			name:    "check Names and IDs, OR",
			records: df,
			index:   0,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
			want: containers.NewSet[int](2, 3, 4, 6, 7, 8),
		},
		{
			name:    "check Full Name, OR",
			records: df,
			index:   0,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName]},
			},
			want: containers.NewSet[int](5),
		},
		{
			name:    "check Names and IDs and Full Name, OR",
			records: df,
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			want: containers.NewSet[int](2, 3, 4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicates := []containers.Set[int]{0: containers.NewSet[int]()}
			getDuplicationScoresForRecord(tt.config, tt.records, tt.index, duplicates[0])
			assert.Equal(t, tt.want, duplicates[0])
		})
	}
}

func TestFindDuplicatesInUpload(t *testing.T) {
	tests := []struct {
		name   string
		config deduplication.DeduplicationConfig
		want   []containers.Set[int]
	}{
		{
			name: "check IDs, AND",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds]},
			},
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
			name: "check Names, AND",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames]},
			},
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
			name: "check Names and IDs, AND",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
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
			name: "check Full Name, AND",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName]},
			},
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
			name: "check Names and IDs and Full Name, AND",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
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
			name: "check IDs, OR",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds]},
			},
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
			name: "check Names, OR",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames]},
			},
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
			name: "check Names and IDs, OR",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
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
			name: "check Full Name, OR",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName]},
			},
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
			name: "check Names and IDs and Full Name, OR",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
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
		{
			name: "no deduplication configured",
			config: deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{},
			},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicates := FindDuplicatesInUpload(tt.config, df)
			assert.Equal(t, tt.want, duplicates)
		})
	}
}

func TestFindDuplicatesInUUIDColumn(t *testing.T) {
	tests := []struct {
		name string
		df   dataframe.DataFrame
		want map[string]containers.Set[int]
	}{
		{
			name: "",
			df: dataframe.New(
				series.New([]string{"1", "2", "2", "2", "5", "1", "5", "8", ""}, series.String, "id"),
				series.New([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}, series.String, "index"),
			),
			want: map[string]containers.Set[int]{
				"1": containers.NewSet[int](0, 5),
				"2": containers.NewSet[int](1, 2, 3),
				"5": containers.NewSet[int](4, 6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicates := getDuplicateUUIDs(tt.df)
			assert.Equal(t, tt.want, duplicates)
		})
	}
}

func TestGetDataframeFromRecords(t *testing.T) {
	tests := []struct {
		name      string
		records   [][]string
		config    []deduplication.DeduplicationType
		mandatory []string
		want      dataframe.DataFrame
		wantErr   bool
	}{
		{
			name: "with id column",
			records: [][]string{
				{"other", "id", "full_name"},
				{"0", "id1", ""},
				{"2", "id3", "full"},
				{"4", "id5", "name"},
			},
			config: []deduplication.DeduplicationType{
				deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
			},
			mandatory: []string{"id"},
			want: dataframe.New(
				series.New([]string{"", "full", "name"}, series.String, "full_name"),
				series.New([]string{"id1", "id3", "id5"}, series.String, "id"),
				series.New([]string{"0", "1", "2"}, series.String, "index"),
			),
			wantErr: false,
		},
		{
			name: "without id column",
			records: [][]string{
				{"other", "full_name"},
				{"0", ""},
				{"2", "full"},
				{"4", "name"},
			},
			config: []deduplication.DeduplicationType{
				deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
			},
			mandatory: []string{},
			want: dataframe.New(
				series.New([]string{"", "full", "name"}, series.String, "full_name"),
				series.New([]string{"0", "1", "2"}, series.String, "index"),
			),
			wantErr: false,
		},
		{
			name: "no deduplication type",
			records: [][]string{
				{"index", "full_name"},
				{"0", ""},
				{"2", "full"},
				{"4", "name"},
			},
			mandatory: []string{},
			config:    []deduplication.DeduplicationType{},
			want:      dataframe.DataFrame{},
			wantErr:   false,
		},
		{
			name: "error",
			records: [][]string{
				{"index", "full_name"},
				{"0", ""},
				{"2", "full"},
				{"4", "name"},
			},
			mandatory: []string{"birth_date"},
			config: []deduplication.DeduplicationType{
				deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
			},
			want:    dataframe.DataFrame{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDf, err := CreateDataframeFromRecords(tt.records, tt.config, tt.mandatory)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, testDf)
			}
		})
	}
}

func TestGetRecordsFromIndividual(t *testing.T) {
	tests := []struct {
		name       string
		individual Individual
		config     []deduplication.DeduplicationType
		mandatory  []string
		want       [][]string
	}{
		{
			name: "without mandatory columns",
			individual: Individual{
				FullName: "A B",
			},
			config: []deduplication.DeduplicationType{
				deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
			},
			mandatory: []string{},
			want: [][]string{
				{"full_name"},
				{"A B"},
			},
		},
		{
			name: "with mandatory columns",
			individual: Individual{
				FullName: "A B",
				ID:       "id",
			},
			config: []deduplication.DeduplicationType{
				deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
			},
			mandatory: []string{"id"},
			want: [][]string{
				{"full_name", "id"},
				{"A B", "id"},
			},
		},
		{
			name: "with redundant mandatory columns",
			individual: Individual{
				FullName: "A B",
				ID:       "id",
			},
			config: []deduplication.DeduplicationType{
				deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
			},
			mandatory: []string{"id", "full_name"},
			want: [][]string{
				{"full_name", "id"},
				{"A B", "id"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			records := GetRecordsFromIndividual(tt.config, &tt.individual, tt.mandatory)
			assert.Equal(t, tt.want, records)
		})
	}
}
