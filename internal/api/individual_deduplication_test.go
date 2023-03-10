package api_test

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testRecords = [][]string{
	{
		constants.FileColumnIndividualIdentificationNumber1,
		constants.FileColumnIndividualIdentificationNumber2,
		constants.FileColumnIndividualIdentificationNumber3,
		constants.FileColumnIndividualFirstName,
		constants.FileColumnIndividualMiddleName,
		constants.FileColumnIndividualLastName,
		constants.FileColumnIndividualNativeName,
		constants.FileColumnIndividualFullName,
	},
	{"1", "2", "3", "A", "B", "C", "D", "E"}, //0
	{"4", "5", "6", "F", "G", "H", "I", "J"}, //1
	{"1", "", "", "A", "B", "C", "", ""},     //2
	{"2", "", "", "A", "B", "", "", ""},      //3
	{"3", "", "", "A", "B", "", "", ""},      //4
	{"", "5", "", "A", "Z", "C", "", "E"},    //5
	{"3", "2", "", "A", "B", "C", "D", ""},   //6
	{"1", "2", "3", "A", "B", "C", "D", "K"}, //7
	{"", "", "5", "A", "B", "C", "D", ""},    //8
	{"7", "", "9", "A", "", "", "", ""},      //9
}

var df = dataframe.LoadRecords(testRecords)

func TestGetDuplicationScoresForRecord(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes []deduplication.DeduplicationTypeName
		records            dataframe.DataFrame
		index              int
		want               containers.Set[int]
	}{
		{
			name:               "check IDs",
			records:            df,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			index:              0,
			want:               containers.NewSet[int](2, 3, 4, 6, 7),
		},
		{
			name:               "check Names",
			records:            df,
			index:              0,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames},
			want:               containers.NewSet[int](6, 7, 8),
		},
		{
			name:               "check Names and IDs",
			records:            df,
			index:              0,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds},
			want:               containers.NewSet[int](6, 7),
		},
		{
			name:               "check Full Name",
			records:            df,
			index:              0,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameFullName},
			want:               containers.NewSet[int](5),
		},
		{
			name:               "check Names and IDs and Full Name",
			records:            df,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameFullName},
			want:               containers.NewSet[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicates := map[int]containers.Set[int]{0: containers.NewSet[int]()}
			api.GetDuplicationScoresForRecord(tt.deduplicationTypes, tt.records, tt.index, duplicates[0])
			assert.Equal(t, tt.want, duplicates[0])
		})
	}
}

func TestFindDuplicatesInUpload(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes []deduplication.DeduplicationTypeName
		records            [][]string
		want               map[int]containers.Set[int]
	}{
		{
			name:               "check IDs",
			records:            testRecords,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			want: map[int]containers.Set[int]{
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
			name:               "check Names",
			records:            testRecords,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames},
			want: map[int]containers.Set[int]{
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
			name:               "check Names and IDs",
			records:            testRecords,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds},
			want: map[int]containers.Set[int]{
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
			name:               "check Full Name",
			records:            testRecords,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameFullName},
			want: map[int]containers.Set[int]{
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
			name:               "check Names and IDs and Full Name",
			records:            testRecords,
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameFullName},
			want: map[int]containers.Set[int]{
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
			duplicates := api.FindDuplicatesInUpload(tt.deduplicationTypes, tt.records)
			assert.Equal(t, tt.want, duplicates)
		})
	}
}
