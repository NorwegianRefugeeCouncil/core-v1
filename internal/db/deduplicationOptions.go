package db

import (
	"github.com/nrc-no/notcore/internal/constants"
)

type DeduplicationOptionName string

const (
	DeduplicationOptionNameIds          DeduplicationOptionName = "Ids"
	DeduplicationOptionNamePhoneNumbers DeduplicationOptionName = "PhoneNumbers"
	DeduplicationOptionNameEmails       DeduplicationOptionName = "Emails"
	DeduplicationOptionNameNames        DeduplicationOptionName = "Names"
	DeduplicationOptionNameFullName     DeduplicationOptionName = "FullName"
	DeduplicationOptionNameFreeField1   DeduplicationOptionName = "FreeField1"
	DeduplicationOptionNameFreeField2   DeduplicationOptionName = "FreeField2"
	DeduplicationOptionNameFreeField3   DeduplicationOptionName = "FreeField3"
	DeduplicationOptionNameFreeField4   DeduplicationOptionName = "FreeField4"
	DeduplicationOptionNameFreeField5   DeduplicationOptionName = "FreeField5"
)

const (
	LOGICAL_OPERATOR_OR  string = "OR"
	LOGICAL_OPERATOR_AND string = "AND"
)

type DeduplicationOptionValue struct {
	Columns   []string
	Condition string
}

type DeduplicationOption struct {
	ID      DeduplicationOptionName
	Value   DeduplicationOptionValue
	Label   string
	Default bool
}

var DeduplicationOptions = map[DeduplicationOptionName]DeduplicationOption{
	DeduplicationOptionNamePhoneNumbers: {
		ID:      DeduplicationOptionNamePhoneNumbers,
		Label:   "Phone numbers",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber3},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameEmails: {
		ID:      DeduplicationOptionNameEmails,
		Label:   "E-Mails",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail3},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameIds: {
		ID:      DeduplicationOptionNameIds,
		Label:   "IDs",
		Default: true,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber3},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameNames: {
		ID:      DeduplicationOptionNameNames,
		Label:   "Names",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualFirstName, constants.DBColumnIndividualMiddleName, constants.DBColumnIndividualLastName},
			Condition: LOGICAL_OPERATOR_AND,
		},
	},
	DeduplicationOptionNameFullName: {
		ID:      DeduplicationOptionNameFullName,
		Label:   "Full Name",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualFullName},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameFreeField1: {
		ID:      DeduplicationOptionNameFreeField1,
		Label:   "Free Field 1",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualFreeField1},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameFreeField2: {
		ID:      DeduplicationOptionNameFreeField2,
		Label:   "Free Field 2",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualFreeField2},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameFreeField3: {
		ID:      DeduplicationOptionNameFreeField3,
		Label:   "Free Field 3",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualFreeField3},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameFreeField4: {
		ID:      DeduplicationOptionNameFreeField4,
		Label:   "Free Field 4",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualFreeField4},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
	DeduplicationOptionNameFreeField5: {
		ID:      DeduplicationOptionNameFreeField5,
		Label:   "Free Field 5",
		Default: false,
		Value: DeduplicationOptionValue{
			Columns:   []string{constants.DBColumnIndividualFreeField5},
			Condition: LOGICAL_OPERATOR_OR,
		},
	},
}

var deduplicationOptionNames = map[string]DeduplicationOptionName{
	"Ids":          DeduplicationOptionNameIds,
	"PhoneNumbers": DeduplicationOptionNamePhoneNumbers,
	"Emails":       DeduplicationOptionNameEmails,
	"Names":        DeduplicationOptionNameNames,
	"FullName":     DeduplicationOptionNameFullName,
	"FreeField1":   DeduplicationOptionNameFreeField1,
	"FreeField2":   DeduplicationOptionNameFreeField2,
	"FreeField3":   DeduplicationOptionNameFreeField3,
	"FreeField4":   DeduplicationOptionNameFreeField4,
	"FreeField5":   DeduplicationOptionNameFreeField5,
}

func ParseDeduplicationOptionName(str string) (DeduplicationOptionName, bool) {
	c, ok := deduplicationOptionNames[str]
	return c, ok
}
