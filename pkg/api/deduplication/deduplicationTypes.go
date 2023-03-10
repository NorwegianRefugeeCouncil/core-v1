package deduplication

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/constants"
)

type DeduplicationTypeName string

const (
	DeduplicationTypeNameIds          DeduplicationTypeName = "Ids"
	DeduplicationTypeNamePhoneNumbers DeduplicationTypeName = "PhoneNumbers"
	DeduplicationTypeNameEmails       DeduplicationTypeName = "Emails"
	DeduplicationTypeNameNames        DeduplicationTypeName = "Names"
	DeduplicationTypeNameFullName     DeduplicationTypeName = "FullName"
	DeduplicationTypeNameFreeField1   DeduplicationTypeName = "FreeField1"
	DeduplicationTypeNameFreeField2   DeduplicationTypeName = "FreeField2"
	DeduplicationTypeNameFreeField3   DeduplicationTypeName = "FreeField3"
	DeduplicationTypeNameFreeField4   DeduplicationTypeName = "FreeField4"
	DeduplicationTypeNameFreeField5   DeduplicationTypeName = "FreeField5"
)

const (
	LOGICAL_OPERATOR_OR  string = "OR"
	LOGICAL_OPERATOR_AND string = "AND"
)

type DeduplicationTypeValue struct {
	Columns   []string
	Condition string
}

type DeduplicationType struct {
	ID    DeduplicationTypeName
	Value DeduplicationTypeValue
	Label string
	Order int
}

var DeduplicationTypes = map[DeduplicationTypeName]DeduplicationType{
	DeduplicationTypeNamePhoneNumbers: {
		ID:    DeduplicationTypeNamePhoneNumbers,
		Label: "Phone numbers",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber3},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 4,
	},
	DeduplicationTypeNameEmails: {
		ID:    DeduplicationTypeNameEmails,
		Label: "E-Mails",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail3},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 2,
	},
	DeduplicationTypeNameIds: {
		ID:    DeduplicationTypeNameIds,
		Label: "Identification numbers",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber3},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 0,
	},
	DeduplicationTypeNameNames: {
		ID:    DeduplicationTypeNameNames,
		Label: "Names (First, Middle, Last, Native)",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFirstName, constants.DBColumnIndividualMiddleName, constants.DBColumnIndividualLastName, constants.DBColumnIndividualNativeName},
			Condition: LOGICAL_OPERATOR_AND,
		},
		Order: 8,
	},
	DeduplicationTypeNameFullName: {
		ID:    DeduplicationTypeNameFullName,
		Label: "Full Name",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFullName},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 6,
	},
	DeduplicationTypeNameFreeField1: {
		ID:    DeduplicationTypeNameFreeField1,
		Label: "Free Field 1",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField1},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 1,
	},
	DeduplicationTypeNameFreeField2: {
		ID:    DeduplicationTypeNameFreeField2,
		Label: "Free Field 2",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField2},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 3,
	},
	DeduplicationTypeNameFreeField3: {
		ID:    DeduplicationTypeNameFreeField3,
		Label: "Free Field 3",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField3},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 5,
	},
	DeduplicationTypeNameFreeField4: {
		ID:    DeduplicationTypeNameFreeField4,
		Label: "Free Field 4",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField4},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 7,
	},
	DeduplicationTypeNameFreeField5: {
		ID:    DeduplicationTypeNameFreeField5,
		Label: "Free Field 5",
		Value: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField5},
			Condition: LOGICAL_OPERATOR_OR,
		},
		Order: 9,
	},
}

var deduplicationTypeNames = map[string]DeduplicationTypeName{
	"Ids":          DeduplicationTypeNameIds,
	"PhoneNumbers": DeduplicationTypeNamePhoneNumbers,
	"Emails":       DeduplicationTypeNameEmails,
	"Names":        DeduplicationTypeNameNames,
	"FullName":     DeduplicationTypeNameFullName,
	"FreeField1":   DeduplicationTypeNameFreeField1,
	"FreeField2":   DeduplicationTypeNameFreeField2,
	"FreeField3":   DeduplicationTypeNameFreeField3,
	"FreeField4":   DeduplicationTypeNameFreeField4,
	"FreeField5":   DeduplicationTypeNameFreeField5,
}

func GetDeduplicationTypeNames(deduplicationTypes []string) ([]DeduplicationTypeName, error) {
	optionNames := make([]DeduplicationTypeName, 0)
	fileColumns := make([]string, 0)
	for _, d := range deduplicationTypes {
		dt, ok := deduplicationTypeNames[d]
		optionNames = append(optionNames, dt)
		if ok {
			for _, vc := range DeduplicationTypes[dt].Value.Columns {
				fileColumns = append(fileColumns, constants.IndividualDBToFileMap[vc])
			}
		} else {
			return nil, fmt.Errorf("invalid deduplication type: %s", d)
		}
	}
	return optionNames, nil
}
