package constants

// The language data was obtained from https://iso639-3.sil.org/code_tables/download_tables

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed languages.json
var languagesJson string

//go:embed language_names.json
var languageNamesJson string

//go:embed language_retirements.json
var languageRetirements string

//go:embed language_macro_mapping.json
var macroLanguageMapping string

func init() {
	if err := json.Unmarshal([]byte(languagesJson), &Languages); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(languageNamesJson), &LanguageNames); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(languageRetirements), &LanguageRetirements); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(macroLanguageMapping), &MacroLanguageMappings); err != nil {
		panic(err)
	}

	for _, l := range Languages {
		LanguagesByCode[l.ID] = l
	}
}

var LanguagesByCode = make(map[string]Language)

type Language struct {
	ID       string `json:"id"`
	ISO6392B string `json:"iso6392B"`
	ISO6392T string `json:"iso6392T"`
	ISO6391  string `json:"iso6391"`
	Scope    string `json:"scope"`
	Type     string `json:"type"`
	Name     string `json:"name"`
}

type LanguageName struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InvertedName string `json:"invertedName"`
}

type LanguageRetirement struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	RetirementReason RetirementReason `json:"retirementReason"`
	RetirementRemedy string           `json:"retirementRemedy"`
	ChangeTo         string           `json:"changeTo"`
	EffectiveDate    string           `json:"effectiveDate"`
}

type MacroLanguageMapping struct {
	MacroLanguageID          string         `json:"macroLanguageId"`
	IndividualLanguageID     string         `json:"individualLanguageId"`
	IndividualLanguageStatus LanguageStatus `json:"individualLanguageStatus"`
}

var Languages []Language
var LanguageNames []LanguageName
var LanguageRetirements []LanguageRetirement
var MacroLanguageMappings []MacroLanguageMapping

type LanguageScope uint

const (
	LanguageScopeIndividual LanguageScope = iota
	LanguageScopeMacro
	LanguageScopeSpecial
)

func (l *LanguageScope) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"I"`:
		*l = LanguageScopeIndividual
	case `"M"`:
		*l = LanguageScopeMacro
	case `"S"`:
		*l = LanguageScopeSpecial
	default:
		return fmt.Errorf("invalid language scope: %s", string(data))
	}
	return nil
}

type LanguageType uint

const (
	LanguageTypeAncient LanguageType = iota
	LanguageTypeConstructed
	LanguageTypeSpecial
	LanguageTypeExtinct
	LanguageTypeHistorical
	LanguageTypeLiving
)

func (l *LanguageType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"A"`:
		*l = LanguageTypeAncient
	case `"C"`:
		*l = LanguageTypeConstructed
	case `"S"`:
		*l = LanguageTypeSpecial
	case `"E"`:
		*l = LanguageTypeExtinct
	case `"H"`:
		*l = LanguageTypeHistorical
	case `"L"`:
		*l = LanguageTypeLiving
	default:
		return fmt.Errorf("invalid language type: %s", string(data))
	}
	return nil
}

type RetirementReason uint

const (
	RetirementReasonChange RetirementReason = iota
	RetirementReasonDuplicate
	RetirementReasonNonExistent
	RetirementReasonSplit
	RetirementReasonMerge
)

func (r *RetirementReason) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"C"`:
		*r = RetirementReasonChange
	case `"D"`:
		*r = RetirementReasonDuplicate
	case `"N"`:
		*r = RetirementReasonNonExistent
	case `"S"`:
		*r = RetirementReasonSplit
	case `"M"`:
		*r = RetirementReasonMerge
	default:
		return fmt.Errorf("invalid retirement reason: %s", string(data))
	}
	return nil
}

type LanguageStatus uint

const (
	LanguageStatusActive LanguageStatus = iota
	LanguageStatusRetired
)

func (l *LanguageStatus) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"A"`:
		*l = LanguageStatusActive
	case `"R"`:
		*l = LanguageStatusRetired
	default:
		return fmt.Errorf("invalid language status: %s", string(data))
	}
	return nil
}
