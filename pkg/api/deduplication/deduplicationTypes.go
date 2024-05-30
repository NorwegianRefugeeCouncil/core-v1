package deduplication

import (
	"fmt"

	"github.com/nrc-no/notcore/internal/constants"
)

type DeduplicationTypeName string
type LogicOperator string

const (
	DeduplicationTypeNameIds          DeduplicationTypeName = "Ids"
	DeduplicationTypeNamePhoneNumbers DeduplicationTypeName = "PhoneNumbers"
	DeduplicationTypeNameMothersName  DeduplicationTypeName = "MothersName"
	DeduplicationTypeNameEmails       DeduplicationTypeName = "Emails"
	DeduplicationTypeNameNames        DeduplicationTypeName = "Names"
	DeduplicationTypeNameFullName     DeduplicationTypeName = "FullName"
	DeduplicationTypeNameFreeField1   DeduplicationTypeName = "FreeField1"
	DeduplicationTypeNameFreeField2   DeduplicationTypeName = "FreeField2"
	DeduplicationTypeNameFreeField3   DeduplicationTypeName = "FreeField3"
	DeduplicationTypeNameFreeField4   DeduplicationTypeName = "FreeField4"
	DeduplicationTypeNameFreeField5   DeduplicationTypeName = "FreeField5"
	DeduplicationTypeNameBirthdate    DeduplicationTypeName = "Birthdate"
)

const (
	LOGICAL_OPERATOR_OR  LogicOperator = "OR"
	LOGICAL_OPERATOR_AND LogicOperator = "AND"
)

type DeduplicationTypeValue struct {
	Columns          []string
	Condition        LogicOperator
	QueryAnd         string
	QueryOr          string
	QueryNotAllEmpty string
}

type DeduplicationType struct {
	ID     DeduplicationTypeName
	Config DeduplicationTypeValue
	Label  string
	Order  int
}

type DeduplicationConfig struct {
	Operator LogicOperator
	Types    []DeduplicationType
}

var DeduplicationTypes = map[DeduplicationTypeName]DeduplicationType{
	DeduplicationTypeNamePhoneNumbers: {
		ID:    DeduplicationTypeNamePhoneNumbers,
		Label: "deduplication_type_phone_numbers",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber3},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("(ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s = '' AND ti.%s = '' AND ti.%s ='')",
				constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber3,
				constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber3,
				constants.DBColumnIndividualPhoneNumber3,
				constants.DBColumnIndividualPhoneNumber3, constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber3, constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber3, constants.DBColumnIndividualPhoneNumber3,
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber3),
			QueryOr: fmt.Sprintf("(ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s))",
				constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber3,
				constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber3,
				constants.DBColumnIndividualPhoneNumber3,
				constants.DBColumnIndividualPhoneNumber3, constants.DBColumnIndividualPhoneNumber1,
				constants.DBColumnIndividualPhoneNumber3, constants.DBColumnIndividualPhoneNumber2,
				constants.DBColumnIndividualPhoneNumber3, constants.DBColumnIndividualPhoneNumber3),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != '' OR ti.%s != '' OR ti.%s != ''",
				constants.DBColumnIndividualPhoneNumber1, constants.DBColumnIndividualPhoneNumber2, constants.DBColumnIndividualPhoneNumber3),
		},
		Order: 4,
	},
	DeduplicationTypeNameEmails: {
		ID:    DeduplicationTypeNameEmails,
		Label: "deduplication_type_emails",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail3},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("(ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s = '' AND ti.%s = '' AND ti.%s ='')",
				constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail3,
				constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail3,
				constants.DBColumnIndividualEmail3,
				constants.DBColumnIndividualEmail3, constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail3, constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail3, constants.DBColumnIndividualEmail3,
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail3),
			QueryOr: fmt.Sprintf("(ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s))",
				constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail3,
				constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail3,
				constants.DBColumnIndividualEmail3,
				constants.DBColumnIndividualEmail3, constants.DBColumnIndividualEmail1,
				constants.DBColumnIndividualEmail3, constants.DBColumnIndividualEmail2,
				constants.DBColumnIndividualEmail3, constants.DBColumnIndividualEmail3),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != '' OR ti.%s != '' OR ti.%s != ''",
				constants.DBColumnIndividualEmail1, constants.DBColumnIndividualEmail2, constants.DBColumnIndividualEmail3),
		},
		Order: 2,
	},
	DeduplicationTypeNameIds: {
		ID:    DeduplicationTypeNameIds,
		Label: "deduplication_type_id_numbers",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber3},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("(ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s = '' AND ti.%s = '' AND ti.%s ='')",
				constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber3,
				constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber3,
				constants.DBColumnIndividualIdentificationNumber3,
				constants.DBColumnIndividualIdentificationNumber3, constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber3, constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber3, constants.DBColumnIndividualIdentificationNumber3,
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber3),
			QueryOr: fmt.Sprintf("(ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s)) OR (ti.%s != '' AND (ti.%s = ir.%s OR ti.%s = ir.%s OR ti.%s = ir.%s))",
				constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber3,
				constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber3,
				constants.DBColumnIndividualIdentificationNumber3,
				constants.DBColumnIndividualIdentificationNumber3, constants.DBColumnIndividualIdentificationNumber1,
				constants.DBColumnIndividualIdentificationNumber3, constants.DBColumnIndividualIdentificationNumber2,
				constants.DBColumnIndividualIdentificationNumber3, constants.DBColumnIndividualIdentificationNumber3),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != '' OR ti.%s != '' OR ti.%s != ''",
				constants.DBColumnIndividualIdentificationNumber1, constants.DBColumnIndividualIdentificationNumber2, constants.DBColumnIndividualIdentificationNumber3),
		},
		Order: 0,
	},
	DeduplicationTypeNameNames: {
		ID:    DeduplicationTypeNameNames,
		Label: "deduplication_type_names",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFirstName, constants.DBColumnIndividualMiddleName, constants.DBColumnIndividualLastName, constants.DBColumnIndividualNativeName},
			Condition: LOGICAL_OPERATOR_AND,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s AND ti.%s = ir.%s AND ti.%s = ir.%s AND ti.%s = ir.%s",
				constants.DBColumnIndividualFirstName, constants.DBColumnIndividualFirstName,
				constants.DBColumnIndividualMiddleName, constants.DBColumnIndividualMiddleName,
				constants.DBColumnIndividualLastName, constants.DBColumnIndividualLastName,
				constants.DBColumnIndividualNativeName, constants.DBColumnIndividualNativeName),
			QueryOr: fmt.Sprintf("ti.%s = ir.%s AND ti.%s = ir.%s AND ti.%s = ir.%s AND ti.%s = ir.%s AND (ti.%s != '' OR ti.%s != '' OR ti.%s != '' OR ti.%s != '')",
				constants.DBColumnIndividualFirstName, constants.DBColumnIndividualFirstName,
				constants.DBColumnIndividualMiddleName, constants.DBColumnIndividualMiddleName,
				constants.DBColumnIndividualLastName, constants.DBColumnIndividualLastName,
				constants.DBColumnIndividualNativeName, constants.DBColumnIndividualNativeName,
				constants.DBColumnIndividualFirstName, constants.DBColumnIndividualMiddleName,
				constants.DBColumnIndividualLastName, constants.DBColumnIndividualNativeName),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != '' OR ti.%s != '' OR ti.%s != '' OR ti.%s != ''",
				constants.DBColumnIndividualFirstName, constants.DBColumnIndividualMiddleName, constants.DBColumnIndividualLastName, constants.DBColumnIndividualNativeName),
		},
		Order: 10,
	},
	DeduplicationTypeNameBirthdate: {
		ID:    DeduplicationTypeNameBirthdate,
		Label: "deduplication_type_birthday",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualBirthDate},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualBirthDate, constants.DBColumnIndividualBirthDate),
			QueryOr: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualBirthDate, constants.DBColumnIndividualBirthDate),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s IS NOT NULL", constants.DBColumnIndividualBirthDate), 

		},
		Order: 11,
	},
	DeduplicationTypeNameMothersName: {
		ID:    DeduplicationTypeNameMothersName,
		Label: "deduplication_type_mothers_name",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualMothersName},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualMothersName, constants.DBColumnIndividualMothersName),
			QueryOr: fmt.Sprintf("ti.%s != '' AND ti.%s = ir.%s",
				constants.DBColumnIndividualMothersName, constants.DBColumnIndividualMothersName, constants.DBColumnIndividualMothersName),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != ''",	constants.DBColumnIndividualMothersName),
		},
		Order: 8,
	},
	DeduplicationTypeNameFullName: {
		ID:    DeduplicationTypeNameFullName,
		Label: "deduplication_type_full_name",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFullName},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualFullName, constants.DBColumnIndividualFullName),
			QueryOr: fmt.Sprintf("ti.%s != '' AND ti.%s = ir.%s",
				constants.DBColumnIndividualFullName, constants.DBColumnIndividualFullName, constants.DBColumnIndividualFullName),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != ''",	constants.DBColumnIndividualFullName),
		},
		Order: 6,
	},
	DeduplicationTypeNameFreeField1: {
		ID:    DeduplicationTypeNameFreeField1,
		Label: "deduplication_type_free_field_1",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField1},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField1, constants.DBColumnIndividualFreeField1),
			QueryOr: fmt.Sprintf("ti.%s != '' AND ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField1, constants.DBColumnIndividualFreeField1, constants.DBColumnIndividualFreeField1),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != ''",	constants.DBColumnIndividualFreeField1),
		},
		Order: 1,
	},
	DeduplicationTypeNameFreeField2: {
		ID:    DeduplicationTypeNameFreeField2,
		Label: "deduplication_type_free_field_2",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField2},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField2, constants.DBColumnIndividualFreeField2),
			QueryOr: fmt.Sprintf("ti.%s != '' AND ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField2, constants.DBColumnIndividualFreeField2, constants.DBColumnIndividualFreeField2),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != ''",	constants.DBColumnIndividualFreeField2),
		},
		Order: 3,
	},
	DeduplicationTypeNameFreeField3: {
		ID:    DeduplicationTypeNameFreeField3,
		Label: "deduplication_type_free_field_3",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField3},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField3, constants.DBColumnIndividualFreeField3),
			QueryOr: fmt.Sprintf("ti.%s != '' AND ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField3, constants.DBColumnIndividualFreeField3, constants.DBColumnIndividualFreeField3),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != ''",	constants.DBColumnIndividualFreeField3),
		},
		Order: 5,
	},
	DeduplicationTypeNameFreeField4: {
		ID:    DeduplicationTypeNameFreeField4,
		Label: "deduplication_type_free_field_4",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField4},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField4, constants.DBColumnIndividualFreeField4),
			QueryOr: fmt.Sprintf("ti.%s != '' AND ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField4, constants.DBColumnIndividualFreeField4, constants.DBColumnIndividualFreeField4),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != ''",	constants.DBColumnIndividualFreeField4),
		},
		Order: 7,
	},
	DeduplicationTypeNameFreeField5: {
		ID:    DeduplicationTypeNameFreeField5,
		Label: "deduplication_type_free_field_5",
		Config: DeduplicationTypeValue{
			Columns:   []string{constants.DBColumnIndividualFreeField5},
			Condition: LOGICAL_OPERATOR_OR,
			QueryAnd: fmt.Sprintf("ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField5, constants.DBColumnIndividualFreeField5),
			QueryOr: fmt.Sprintf("ti.%s != '' AND ti.%s = ir.%s",
				constants.DBColumnIndividualFreeField5, constants.DBColumnIndividualFreeField5, constants.DBColumnIndividualFreeField5),
			QueryNotAllEmpty: fmt.Sprintf("ti.%s != ''",	constants.DBColumnIndividualFreeField5),
		},
		Order: 9,
	},
}

var deduplicationTypeNames = map[string]DeduplicationTypeName{
	"Ids":          DeduplicationTypeNameIds,
	"PhoneNumbers": DeduplicationTypeNamePhoneNumbers,
	"Emails":       DeduplicationTypeNameEmails,
	"Names":        DeduplicationTypeNameNames,
	"MothersName":  DeduplicationTypeNameMothersName,
	"FullName":     DeduplicationTypeNameFullName,
	"FreeField1":   DeduplicationTypeNameFreeField1,
	"FreeField2":   DeduplicationTypeNameFreeField2,
	"FreeField3":   DeduplicationTypeNameFreeField3,
	"FreeField4":   DeduplicationTypeNameFreeField4,
	"FreeField5":   DeduplicationTypeNameFreeField5,
	"Birthdate":    DeduplicationTypeNameBirthdate,
}

func GetDeduplicationConfig(deduplicationTypes []string, operator LogicOperator) (DeduplicationConfig, error) {
	optionNames := make([]DeduplicationType, 0)
	for _, d := range deduplicationTypes {
		dt := DeduplicationTypes[DeduplicationTypeName(d)]
		optionNames = append(optionNames, dt)
	}
	return DeduplicationConfig{
		operator, optionNames,
	}, nil
}
