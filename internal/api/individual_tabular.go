package api

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/slices"
)

// Unmarshal

func UnmarshalIndividualsCSV(reader io.Reader, individuals *[]*Individual, fields *[]string) error {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	return UnmarshalIndividualsTabularData(records, individuals, fields)
}

func UnmarshalIndividualsExcel(reader io.Reader, individuals *[]*Individual, fields *[]string) error {
	f, err := excelize.OpenReader(reader)

	if err != nil {
		return err
	}

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		err := errors.New("no sheets found")
		return err
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		err := errors.New("no rows found")
		return err
	}

	return UnmarshalIndividualsTabularData(rows, individuals, fields)
}

func UnmarshalIndividualsTabularData(data [][]string, individuals *[]*Individual, fields *[]string) error {
	colMapping := map[string]int{}
	headerRow := data[0]
	for i, col := range headerRow {
		*fields = append(*fields, constants.IndividualFileToDBMap[trimString(col)])
		col = trimString(col)
		colMapping[strings.Trim(col, " \n\t\r")] = i
	}

	for _, cols := range data[1:] {
		individual := &Individual{}
		if err := individual.unmarshalTabularData(colMapping, cols); err != nil {
			return err
		}
		*individuals = append(*individuals, individual)
	}

	return nil
}

func (i *Individual) unmarshalTabularData(colMapping map[string]int, cols []string) error {
	var err error
	for field, idx := range colMapping {
		switch field {
		case constants.FileColumnIndividualAddress:
			i.Address = cols[idx]
		case constants.FileColumnIndividualBirthDate:
			i.BirthDate, err = ParseDate(cols[idx])
		case constants.FileColumnIndividualCognitiveDisabilityLevel:
			i.CognitiveDisabilityLevel, err = ParseDisabilityLevel(cols[idx])
		case constants.FileColumnIndividualCollectionAdministrativeArea1:
			i.CollectionAdministrativeArea1 = cols[idx]
		case constants.FileColumnIndividualCollectionAdministrativeArea2:
			i.CollectionAdministrativeArea2 = cols[idx]
		case constants.FileColumnIndividualCollectionAdministrativeArea3:
			i.CollectionAdministrativeArea3 = cols[idx]
		case constants.FileColumnIndividualCollectionAgentName:
			i.CollectionAgentName = cols[idx]
		case constants.FileColumnIndividualCollectionAgentTitle:
			i.CollectionAgentTitle = cols[idx]
		case constants.FileColumnIndividualCollectionTime:
			var collectionTime *time.Time
			collectionTime, err = ParseDate(cols[idx])
			if err != nil && collectionTime != nil {
				i.CollectionTime = *collectionTime
			}
		case constants.FileColumnIndividualCommunicationDisabilityLevel:
			i.CommunicationDisabilityLevel, err = ParseDisabilityLevel(cols[idx])
		case constants.FileColumnIndividualCommunityID:
			i.CommunityID = cols[idx]
		case constants.FileColumnIndividualDisplacementStatus:
			i.DisplacementStatus = DisplacementStatus(cols[idx])
		case constants.FileColumnIndividualEmail:
			i.Email = cols[idx]
		case constants.FileColumnIndividualFullName:
			i.FullName = cols[idx]
		case constants.FileColumnIndividualGender:
			i.Gender = Gender(cols[idx])
		case constants.FileColumnIndividualHasCognitiveDisability:
			i.HasCognitiveDisability = isTrue(cols[idx])
		case constants.FileColumnIndividualHasCommunicationDisability:
			i.HasCommunicationDisability = isTrue(cols[idx])
		case constants.FileColumnIndividualHasConsentedToRGPD:
			i.HasConsentedToRGPD = isTrue(cols[idx])
		case constants.FileColumnIndividualHasConsentedToReferral:
			i.HasConsentedToReferral = isTrue(cols[idx])
		case constants.FileColumnIndividualHasHearingDisability:
			i.HasHearingDisability = isTrue(cols[idx])
		case constants.FileColumnIndividualHasMobilityDisability:
			i.HasMobilityDisability = isTrue(cols[idx])
		case constants.FileColumnIndividualHasSelfCareDisability:
			i.HasSelfCareDisability = isTrue(cols[idx])
		case constants.FileColumnIndividualHasVisionDisability:
			i.HasVisionDisability = isTrue(cols[idx])
		case constants.FileColumnIndividualHearingDisabilityLevel:
			i.HearingDisabilityLevel, err = ParseDisabilityLevel(cols[idx])
		case constants.FileColumnIndividualHouseholdID:
			i.HouseholdID = cols[idx]
		case constants.FileColumnIndividualIdentificationType1:
			i.IdentificationType1 = cols[idx]
		case constants.FileColumnIndividualIdentificationTypeExplanation1:
			i.IdentificationTypeExplanation1 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber1:
			i.IdentificationNumber1 = cols[idx]
		case constants.FileColumnIndividualIdentificationType2:
			i.IdentificationType2 = cols[idx]
		case constants.FileColumnIndividualIdentificationTypeExplanation2:
			i.IdentificationTypeExplanation2 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber2:
			i.IdentificationNumber2 = cols[idx]
		case constants.FileColumnIndividualIdentificationType3:
			i.IdentificationType3 = cols[idx]
		case constants.FileColumnIndividualIdentificationTypeExplanation3:
			i.IdentificationTypeExplanation3 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber3:
			i.IdentificationNumber3 = cols[idx]
		case constants.FileColumnIndividualIdentificationContext:
			i.IdentificationContext = cols[idx]
		case constants.FileColumnIndividualInternalID:
			i.InternalID = cols[idx]
		case constants.FileColumnIndividualIsHeadOfCommunity:
			i.IsHeadOfCommunity = isTrue(cols[idx])
		case constants.FileColumnIndividualIsHeadOfHousehold:
			i.IsHeadOfHousehold = isTrue(cols[idx])
		case constants.FileColumnIndividualIsMinor:
			i.IsMinor = isTrue(cols[idx])
		case constants.FileColumnIndividualMobilityDisabilityLevel:
			i.MobilityDisabilityLevel, err = ParseDisabilityLevel(cols[idx])
		case constants.FileColumnIndividualNationality1:
			i.Nationality1 = cols[idx]
		case constants.FileColumnIndividualNationality2:
			i.Nationality2 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber:
			i.PhoneNumber = cols[idx]
		case constants.FileColumnIndividualPreferredContactMethod:
			i.PreferredContactMethod = cols[idx]
		case constants.FileColumnIndividualPreferredContactMethodComments:
			i.PreferredContactMethodComments = cols[idx]
		case constants.FileColumnIndividualPreferredName:
			i.PreferredName = cols[idx]
		case constants.FileColumnIndividualPreferredCommunicationLanguage:
			i.PreferredCommunicationLanguage = cols[idx]
		case constants.FileColumnIndividualPrefersToRemainAnonymous:
			i.PrefersToRemainAnonymous = isTrue(cols[idx])
		case constants.FileColumnIndividualPresentsProtectionConcerns:
			i.PresentsProtectionConcerns = isTrue(cols[idx])
		case constants.FileColumnIndividualSelfCareDisabilityLevel:
			i.SelfCareDisabilityLevel, err = ParseDisabilityLevel(cols[idx])
		case constants.FileColumnIndividualSpokenLanguage1:
			i.SpokenLanguage1 = cols[idx]
		case constants.FileColumnIndividualSpokenLanguage2:
			i.SpokenLanguage2 = cols[idx]
		case constants.FileColumnIndividualSpokenLanguage3:
			i.SpokenLanguage3 = cols[idx]
		case constants.FileColumnIndividualVisionDisabilityLevel:
			i.VisionDisabilityLevel, err = ParseDisabilityLevel(cols[idx])
		}
	}
	if err != nil {
		return err
	}
	i.Normalize()
	return nil
}

// Marshal

func MarshalIndividualsCSV(w io.Writer, individuals []*Individual) error {
	csvEncoder := csv.NewWriter(w)
	defer csvEncoder.Flush()

	if err := csvEncoder.Write(constants.IndividualFileColumns); err != nil {
		return err
	}

	for _, individual := range individuals {
		row, err := individual.marshalTabularData()
		if err != nil {
			return err
		}
		if err := csvEncoder.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func MarshalIndividualsExcel(w io.Writer, individuals []*Individual) error {
	const sheetName = "Individuals"

	f := excelize.NewFile()

	sheet := f.NewSheet(sheetName)

	if err := f.SetSheetRow(sheetName, "A1", &constants.IndividualFileColumns); err != nil {
		return err
	}

	for i, individual := range individuals {
		row, err := individual.marshalTabularData()
		if err != nil {
			return err
		}
		if err := f.SetSheetRow(sheetName, fmt.Sprintf("A%d", i+2), &row); err != nil {
			return err
		}
	}

	f.SetActiveSheet(sheet)

	if err := f.Write(w); err != nil {
		return err
	}

	return nil
}

func (i *Individual) marshalTabularData() ([]string, error) {
	row := make([]string, len(constants.IndividualFileColumns))
	for j, col := range constants.IndividualFileColumns {
		value, err := i.GetFieldValue(constants.IndividualFileToDBMap[col])
		if err != nil {
			return nil, err
		}

		switch col {
		case constants.FileColumnIndividualBirthDate:
			var birthDate string
			if i.BirthDate != nil {
				birthDate = i.BirthDate.Format("2006-01-02")
			}
			row[j] = birthDate
		case constants.FileColumnIndividualIsMinor, constants.FileColumnIndividualPresentsProtectionConcerns:
			row[j] = strconv.FormatBool(value.(bool))
		default:
			row[j] = value.(string)
		}
	}
	return row, nil
}

var TRUE_VALUES = []string{"true", "yes", "1"}

func isTrue(value string) bool {
	return slices.Contains(TRUE_VALUES, strings.ToLower(value))
}
