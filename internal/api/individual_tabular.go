package api

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/nrc-no/notcore/pkg/logutils"
	"io"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/slices"
)

// Unmarshal

type FileError struct {
	Message string
	Err     []error
}

func UnmarshalIndividualsCSV(reader io.Reader, individuals *[]*Individual, fields *[]string) ([]FileError, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	if err != nil {
		return []FileError{}, err
	}
	return UnmarshalIndividualsTabularData(records, individuals, fields), err
}

func UnmarshalIndividualsExcel(reader io.Reader, individuals *[]*Individual, fields *[]string) ([]FileError, error) {
	f, err := excelize.OpenReader(reader)
	var fileErrors []FileError
	if err != nil {
		return fileErrors, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		err := errors.New("no sheets found")
		return fileErrors, err
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return fileErrors, err
	}
	if len(rows) == 0 {
		err := errors.New("no rows found")
		return fileErrors, err
	}

	return UnmarshalIndividualsTabularData(rows, individuals, fields), err
}

func UnmarshalIndividualsTabularData(data [][]string, individuals *[]*Individual, fields *[]string) []FileError {
	colMapping := map[string]int{}
	headerRow := data[0]
	var fileErrors []FileError
	var columnErrors []error
	for i, col := range headerRow {
		col = trimString(col)
		field, ok := constants.IndividualFileToDBMap[col]
		if !ok {
			ok = constants.IndividualSystemFileColumns.Contains(col)
			if ok {
				continue
			}
			columnErrors = append(columnErrors, fmt.Errorf("unknown column: \"%s\"	", logutils.Escape(col)))
		}
		*fields = append(*fields, field)
		col = trimString(col)
		colMapping[strings.Trim(col, " \n\t\r")] = i
	}
	if len(columnErrors) > 0 {
		fileErrors = append(fileErrors, FileError{
			Message: fmt.Sprintf("Unknown columns"),
			Err:     columnErrors,
		})
	}

	for row, cols := range data[1:] {
		individual := &Individual{}
		var rowErrors []error
		for _, err := range individual.unmarshalTabularData(colMapping, cols) {
			rowErrors = append(rowErrors, err)
		}
		if len(rowErrors) > 0 {
			fileErrors = append(fileErrors, FileError{
				Message: fmt.Sprintf("Parsing row #%d has lead to an error", row+2),
				Err:     rowErrors,
			})
		}
		*individuals = append(*individuals, individual)
	}

	return fileErrors
}

func (i *Individual) unmarshalTabularData(colMapping map[string]int, cols []string) []error {
	var errors []error
	for field, idx := range colMapping {
		switch field {
		case constants.FileColumnIndividualID:
			i.ID = cols[idx]
		case constants.FileColumnIndividualActive:
			i.Active = isTrue(cols[idx])
		case constants.FileColumnIndividualAddress:
			i.Address = cols[idx]
		case constants.FileColumnIndividualAge:
			var ageStr = cols[idx]
			if ageStr == "" {
				continue
			}
			age, err := strconv.Atoi(ageStr)
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.Age = &age
		case constants.FileColumnIndividualBirthDate:
			var birthDate *time.Time
			birthDate, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.BirthDate = birthDate
		case constants.FileColumnIndividualCognitiveDisabilityLevel:
			disabilityLevel, err := ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.CognitiveDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualCollectionAdministrativeArea1:
			i.CollectionAdministrativeArea1 = cols[idx]
		case constants.FileColumnIndividualCollectionAdministrativeArea2:
			i.CollectionAdministrativeArea2 = cols[idx]
		case constants.FileColumnIndividualCollectionAdministrativeArea3:
			i.CollectionAdministrativeArea3 = cols[idx]
		case constants.FileColumnIndividualCollectionOffice:
			i.CollectionOffice = cols[idx]
		case constants.FileColumnIndividualCollectionAgentName:
			i.CollectionAgentName = cols[idx]
		case constants.FileColumnIndividualCollectionAgentTitle:
			i.CollectionAgentTitle = cols[idx]
		case constants.FileColumnIndividualComments:
			i.Comments = cols[idx]
		case constants.FileColumnIndividualCollectionTime:
			var collectionTime *time.Time
			collectionTime, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.CollectionTime = *collectionTime
		case constants.FileColumnIndividualCommunicationDisabilityLevel:
			disabilityLevel, err := ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.CommunicationDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualCommunityID:
			i.CommunityID = cols[idx]
		case constants.FileColumnIndividualDisplacementStatus:
			displacementStatus, err := ParseDisplacementStatus(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.DisplacementStatus = displacementStatus
		case constants.FileColumnIndividualDisplacementStatusComment:
			i.DisplacementStatusComment = cols[idx]
		case constants.FileColumnIndividualEmail1:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, err)
					break
				}
				i.Email1 = email.Address
			}
		case constants.FileColumnIndividualEmail2:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, err)
					break
				}
				i.Email2 = email.Address
			}
		case constants.FileColumnIndividualEmail3:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, err)
					break
				}
				i.Email3 = email.Address
			}
		case constants.FileColumnIndividualFullName:
			i.FullName = cols[idx]
		case constants.FileColumnIndividualFirstName:
			i.FirstName = cols[idx]
		case constants.FileColumnIndividualMiddleName:
			i.MiddleName = cols[idx]
		case constants.FileColumnIndividualLastName:
			i.LastName = cols[idx]
		case constants.FileColumnIndividualFreeField1:
			i.FreeField1 = cols[idx]
		case constants.FileColumnIndividualFreeField2:
			i.FreeField2 = cols[idx]
		case constants.FileColumnIndividualFreeField3:
			i.FreeField3 = cols[idx]
		case constants.FileColumnIndividualFreeField4:
			i.FreeField4 = cols[idx]
		case constants.FileColumnIndividualFreeField5:
			i.FreeField5 = cols[idx]
		case constants.FileColumnIndividualSex:
			i.Sex = Sex(cols[idx])
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
			disabilityLevel, err := ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.HearingDisabilityLevel = disabilityLevel
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
		case constants.FileColumnIndividualEngagementContext:
			i.EngagementContext = cols[idx]
		case constants.FileColumnIndividualInternalID:
			i.InternalID = cols[idx]
		case constants.FileColumnIndividualIsHeadOfCommunity:
			i.IsHeadOfCommunity = isTrue(cols[idx])
		case constants.FileColumnIndividualIsHeadOfHousehold:
			i.IsHeadOfHousehold = isTrue(cols[idx])
		case constants.FileColumnIndividualIsFemaleHeadedHousehold:
			i.IsFemaleHeadedHousehold = isTrue(cols[idx])
		case constants.FileColumnIndividualIsMinorHeadedHousehold:
			i.IsMinorHeadedHousehold = isTrue(cols[idx])
		case constants.FileColumnIndividualIsMinor:
			i.IsMinor = isTrue(cols[idx])
		case constants.FileColumnIndividualMobilityDisabilityLevel:
			disabilityLevel, err := ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.MobilityDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualNationality1:
			i.Nationality1 = cols[idx]
		case constants.FileColumnIndividualNationality2:
			i.Nationality2 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber1:
			i.PhoneNumber1 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber2:
			i.PhoneNumber2 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber3:
			i.PhoneNumber3 = cols[idx]
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
			disabilityLevel, err := ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.SelfCareDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualSpokenLanguage1:
			i.SpokenLanguage1 = cols[idx]
		case constants.FileColumnIndividualSpokenLanguage2:
			i.SpokenLanguage2 = cols[idx]
		case constants.FileColumnIndividualSpokenLanguage3:
			i.SpokenLanguage3 = cols[idx]
		case constants.FileColumnIndividualVisionDisabilityLevel:
			disabilityLevel, err := ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.VisionDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualServiceCC1:
			cc, err := ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceCC1 = cc
		case constants.FileColumnIndividualServiceRequestedDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceRequestedDate1 = date
		case constants.FileColumnIndividualServiceDeliveredDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceDeliveredDate1 = date
		case constants.FileColumnIndividualServiceComments1:
			i.ServiceComments1 = cols[idx]
		case constants.FileColumnIndividualServiceCC2:
			cc, err := ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceCC2 = cc
		case constants.FileColumnIndividualServiceRequestedDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceRequestedDate2 = date
		case constants.FileColumnIndividualServiceDeliveredDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceDeliveredDate2 = date
		case constants.FileColumnIndividualServiceComments2:
			i.ServiceComments2 = cols[idx]
		case constants.FileColumnIndividualServiceCC3:
			cc, err := ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceCC3 = cc
		case constants.FileColumnIndividualServiceRequestedDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceRequestedDate3 = date
		case constants.FileColumnIndividualServiceDeliveredDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceDeliveredDate3 = date
		case constants.FileColumnIndividualServiceComments3:
			i.ServiceComments3 = cols[idx]
		case constants.FileColumnIndividualServiceCC4:
			cc, err := ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceCC4 = cc
		case constants.FileColumnIndividualServiceRequestedDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceRequestedDate4 = date
		case constants.FileColumnIndividualServiceDeliveredDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceDeliveredDate4 = date
		case constants.FileColumnIndividualServiceComments4:
			i.ServiceComments4 = cols[idx]
		case constants.FileColumnIndividualServiceCC5:
			cc, err := ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceCC5 = cc
		case constants.FileColumnIndividualServiceRequestedDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceRequestedDate5 = date
		case constants.FileColumnIndividualServiceDeliveredDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceDeliveredDate5 = date
		case constants.FileColumnIndividualServiceComments5:
			i.ServiceComments5 = cols[idx]
		case constants.FileColumnIndividualServiceCC6:
			cc, err := ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceCC6 = cc
		case constants.FileColumnIndividualServiceRequestedDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceRequestedDate6 = date
		case constants.FileColumnIndividualServiceDeliveredDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceDeliveredDate6 = date
		case constants.FileColumnIndividualServiceComments6:
			i.ServiceComments6 = cols[idx]
		case constants.FileColumnIndividualServiceCC7:
			cc, err := ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceCC7 = cc
		case constants.FileColumnIndividualServiceRequestedDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceRequestedDate7 = date
		case constants.FileColumnIndividualServiceDeliveredDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.ServiceDeliveredDate7 = date
		case constants.FileColumnIndividualServiceComments7:
			i.ServiceComments7 = cols[idx]
		}
	}
	if len(errors) > 0 {
		return errors
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

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.SetSheetName("Sheet1", sheetName)

	streamWriter, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}

	if err := streamWriter.SetRow("A1", stringArrayToInterfaceArray(constants.IndividualFileColumns)); err != nil {
		return err
	}

	for idx, individual := range individuals {
		row, err := individual.marshalTabularData()
		if err != nil {
			return err
		}
		if err := streamWriter.SetRow(fmt.Sprintf("A%d", idx+2), stringArrayToInterfaceArray(row)); err != nil {
			return err
		}
	}

	if err := streamWriter.Flush(); err != nil {
		return err
	}

	if err := f.Write(w); err != nil {
		return err
	}

	return nil
}

func getTimeFormatForField(field string) string {
	switch field {
	case constants.DBColumnIndividualUpdatedAt:
		return time.RFC3339
	case constants.DBColumnIndividualCreatedAt:
		return time.RFC3339
	case constants.DBColumnIndividualDeletedAt:
		return time.RFC3339
	default:
		return "2006-01-02"
	}
}

func (i *Individual) marshalTabularData() ([]string, error) {
	row := make([]string, len(constants.IndividualFileColumns))
	for j, col := range constants.IndividualFileColumns {
		field, ok := constants.IndividualDBToFileMap[col]
		if !ok {
			return nil, fmt.Errorf("unknown column %s", col) // should not happen but we never know.
		}
		value, err := i.GetFieldValue(field)
		if err != nil {
			return nil, err
		}

		switch v := value.(type) {
		case bool:
			row[j] = strconv.FormatBool(v)
		case *bool:
			if v != nil {
				row[j] = strconv.FormatBool(*v)
			}
		case int:
			row[j] = strconv.Itoa(v)
		case *int:
			if v != nil {
				row[j] = strconv.Itoa(*value.(*int))
			}
		case string:
			row[j] = v
		case time.Time:
			row[j] = v.Format(getTimeFormatForField(field))
		case *time.Time:
			if v != nil {
				row[j] = v.Format(getTimeFormatForField(field))
			}
		case DisabilityLevel:
			row[j] = string(v)
		case DisplacementStatus:
			row[j] = string(v)
		case ServiceCC:
			row[j] = string(v)
		case Sex:
			row[j] = string(v)
		default:
			row[j] = fmt.Sprintf("%v", v)
		}
	}
	return row, nil
}

var TRUE_VALUES = []string{"true", "yes", "1"}

func isTrue(value string) bool {
	return slices.Contains(TRUE_VALUES, strings.ToLower(value))
}

func stringArrayToInterfaceArray(row []string) []interface{} {
	var result []interface{}
	for _, col := range row {
		result = append(result, col)
	}
	return result
}
