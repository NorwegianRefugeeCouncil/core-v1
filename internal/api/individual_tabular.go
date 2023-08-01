package api

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/pkg/logutils"
	"io"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/xuri/excelize/v2"
)

type FileError struct {
	Message string
	Err     []error
}

// Unmarshal

func UnmarshalRecordsFromCSV(records *[][]string, reader io.Reader) error {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	csvReader.Comma = ','
	csvReader.LazyQuotes = false
	csvReader.Comment = 0
	output, err := csvReader.ReadAll()
	if err == nil {
		*records = output
	}
	return err
}

func UnmarshalRecordsFromExcel(records *[][]string, reader io.Reader) error {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

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
	f.Close()
	if err == nil {
		*records = rows
	}

	for i, record := range *records {
		header := (*records)[0]
		diff := len(header) - len(record)
		if diff > 0 {
			filler := make([]string, diff)
			(*records)[i] = append((*records)[i], filler...)
		}
	}
	return err
}

func UnmarshallRecordsFromFile(records *[][]string, reader io.Reader, filename string) error {
	if strings.HasSuffix(filename, ".csv") {
		return UnmarshalRecordsFromCSV(records, reader)
	} else if strings.HasSuffix(filename, ".xlsx") || strings.HasSuffix(filename, ".xls") {
		return UnmarshalRecordsFromExcel(records, reader)
	} else {
		fileNameParts := strings.Split(filename, ".")
		fileType := fileNameParts[len(fileNameParts)-1]

		err := errors.New(fmt.Sprintf("Could not process uploaded file of filetype %s, please upload a .csv or a .xls(x) file.", fileType))
		return err
	}
}

func GetColumnMapping(data [][]string, fields *[]string) (map[string]int, []FileError) {
	colMapping := map[string]int{}
	errors := []error{}
	headerRow := data[0]
	for i, col := range headerRow {
		col = trimString(col)
		field, ok := constants.IndividualFileToDBMap[col]
		if !ok {
			ok = constants.IndividualSystemFileColumns.Contains(col)
			if ok {
				continue
			}
			errors = append(errors, fmt.Errorf("column: \"%s\"	", logutils.Escape(col)))
		}
		*fields = append(*fields, field)
		col = trimString(col)
		colMapping[strings.Trim(col, " \n\t\r")] = i
	}
	if len(errors) > 0 {
		return nil, []FileError{FileError{
			Message: fmt.Sprintf("Unknown columns"),
			Err:     errors,
		}}
	}
	return colMapping, nil
}

func UnmarshalIndividualsTabularData(data [][]string, individuals *[]*Individual, colMapping map[string]int, rowLimit *int) []FileError {

	if rowLimit != nil && len(data[1:]) > *rowLimit {
		return []FileError{{fmt.Sprintf("Your file contains %d participants, which exceeds the upload limit of %d participants at a time.", len(data[1:]), *rowLimit), nil}}
	}
	var fileErrors []FileError

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
	if len(cols) <= len(colMapping) {
		filler := make([]string, len(colMapping)-len(cols))
		cols = append(cols, filler...)
	}
	for field, idx := range colMapping {
		switch field {
		case constants.FileColumnIndividualID:
			i.ID = cols[idx]
		case constants.FileColumnIndividualInactive:
			inactive, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualInactive, err))
			}
			i.Inactive = inactive
		case constants.FileColumnIndividualAddress:
			i.Address = cols[idx]
		case constants.FileColumnIndividualAge:
			age, err := ParseAge(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.Age = age
		case constants.FileColumnIndividualBirthDate:
			var birthDate *time.Time
			birthDate, err := ParseBirthdate(cols[idx])
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.BirthDate = birthDate
		case constants.FileColumnIndividualCognitiveDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualCognitiveDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
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
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualCollectionTime, err))
				break
			}
			if collectionTime != nil {
				i.CollectionTime = *collectionTime
			}
		case constants.FileColumnIndividualCommunicationDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualCommunicationDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.CommunicationDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualCommunityID:
			i.CommunityID = cols[idx]
		case constants.FileColumnIndividualCommunitySize:
			var communitySizeStr = cols[idx]
			if communitySizeStr == "" {
				continue
			}
			communitySize, err := strconv.Atoi(communitySizeStr)
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.CommunitySize = &communitySize
		case constants.FileColumnIndividualDisplacementStatus:
			displacementStatus, err := enumTypes.ParseDisplacementStatus(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualDisplacementStatus, err, enumTypes.AllDisplacementStatuses().String()))
				break
			}
			i.DisplacementStatus = displacementStatus
		case constants.FileColumnIndividualDisplacementStatusComment:
			i.DisplacementStatusComment = cols[idx]
		case constants.FileColumnIndividualEmail1:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail1, err))
					break
				}
				i.Email1 = email.Address
			}
		case constants.FileColumnIndividualEmail2:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail2, err))
					break
				}
				i.Email2 = email.Address
			}
		case constants.FileColumnIndividualEmail3:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail3, err))
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
		case constants.FileColumnIndividualNativeName:
			i.NativeName = cols[idx]
		case constants.FileColumnIndividualMothersName:
			i.MothersName = cols[idx]
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
			sex, err := enumTypes.ParseSex(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualSex, err, enumTypes.AllSexes().String()))
				break
			}
			i.Sex = sex
		case constants.FileColumnIndividualHasCognitiveDisability:
			hasCognitiveDisability, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasCognitiveDisability, err))
			}
			i.HasCognitiveDisability = hasCognitiveDisability
		case constants.FileColumnIndividualHasCommunicationDisability:
			hasCommunicationDisability, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasCommunicationDisability, err))
			}
			i.HasCommunicationDisability = hasCommunicationDisability
		case constants.FileColumnIndividualHasConsentedToRGPD:
			hasConsentedToRGPD, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasConsentedToRGPD, err))
			}
			i.HasConsentedToRGPD = hasConsentedToRGPD
		case constants.FileColumnIndividualHasConsentedToReferral:
			hasConsentedToReferral, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasConsentedToReferral, err))
			}
			i.HasConsentedToReferral = hasConsentedToReferral
		case constants.FileColumnIndividualHasDisability:
			hasDisability, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasDisability, err))
			}
			i.HasDisability = hasDisability
		case constants.FileColumnIndividualHasHearingDisability:
			hasHearingDisability, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasHearingDisability, err))
			}
			i.HasHearingDisability = hasHearingDisability
		case constants.FileColumnIndividualHasMobilityDisability:
			hasMobilityDisability, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasMobilityDisability, err))
			}
			i.HasMobilityDisability = hasMobilityDisability
		case constants.FileColumnIndividualHasSelfCareDisability:
			hasSelfCareDisability, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasSelfCareDisability, err))
			}
			i.HasSelfCareDisability = hasSelfCareDisability
		case constants.FileColumnIndividualHasVisionDisability:
			hasVisionDisability, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualHasVisionDisability, err))
			}
			i.HasVisionDisability = hasVisionDisability
		case constants.FileColumnIndividualHearingDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualHearingDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.HearingDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualHouseholdID:
			i.HouseholdID = cols[idx]
		case constants.FileColumnIndividualHouseholdSize:
			var householdSizeStr = cols[idx]
			if householdSizeStr == "" {
				continue
			}
			householdSize, err := strconv.Atoi(householdSizeStr)
			if err != nil {
				errors = append(errors, err)
				break
			}
			i.CommunitySize = &householdSize
		case constants.FileColumnIndividualIdentificationType1:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIdentificationType1, err, enumTypes.AllIdentificationTypes().String()))
				break
			}
			i.IdentificationType1 = identificationType
		case constants.FileColumnIndividualIdentificationTypeExplanation1:
			i.IdentificationTypeExplanation1 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber1:
			i.IdentificationNumber1 = cols[idx]
		case constants.FileColumnIndividualIdentificationType2:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIdentificationType2, err, enumTypes.AllIdentificationTypes().String()))
				break
			}
			i.IdentificationType2 = identificationType
		case constants.FileColumnIndividualIdentificationTypeExplanation2:
			i.IdentificationTypeExplanation2 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber2:
			i.IdentificationNumber2 = cols[idx]
		case constants.FileColumnIndividualIdentificationType3:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualIdentificationType3, err, enumTypes.AllIdentificationTypes().String()))
				break
			}
			i.IdentificationType3 = identificationType
		case constants.FileColumnIndividualIdentificationTypeExplanation3:
			i.IdentificationTypeExplanation3 = cols[idx]
		case constants.FileColumnIndividualIdentificationNumber3:
			i.IdentificationNumber3 = cols[idx]
		case constants.FileColumnIndividualEngagementContext:
			engagementContext, err := enumTypes.ParseEngagementContext(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualEngagementContext, err, enumTypes.AllEngagementContexts().String()))
				break
			}
			i.EngagementContext = engagementContext
		case constants.FileColumnIndividualInternalID:
			i.InternalID = cols[idx]
		case constants.FileColumnIndividualIsHeadOfCommunity:
			isHeadOfCommunity, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualIsHeadOfCommunity, err))
			}
			i.IsHeadOfCommunity = isHeadOfCommunity
		case constants.FileColumnIndividualIsHeadOfHousehold:
			isHeadOfHousehold, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualIsHeadOfHousehold, err))
			}
			i.IsHeadOfHousehold = isHeadOfHousehold
		case constants.FileColumnIndividualIsFemaleHeadedHousehold:
			isFemaleHeadedHousehold, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualIsFemaleHeadedHousehold, err))
			}
			i.IsFemaleHeadedHousehold = isFemaleHeadedHousehold
		case constants.FileColumnIndividualIsMinorHeadedHousehold:
			isMinorHeadedHousehold, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualIsMinorHeadedHousehold, err))
			}
			i.IsMinorHeadedHousehold = isMinorHeadedHousehold
		case constants.FileColumnIndividualIsMinor:
			isMinor, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualIsMinor, err))
			}
			i.IsMinor = isMinor
		case constants.FileColumnIndividualMobilityDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualMobilityDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.MobilityDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualNationality1:
			if cols[idx] != "" {
				if c := constants.CountriesByCode[cols[idx]].Name; c != "" {
					i.Nationality1 = cols[idx]
				} else if c := constants.CountriesByName[cols[idx]].Name; c != "" {
					i.Nationality1 = constants.CountriesByName[cols[idx]].ISO3166Alpha3
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\". valid values adhere to the ISO3166Alpha3 norm", constants.FileColumnIndividualNationality1, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualNationality2:
			if cols[idx] != "" {
				if c := constants.CountriesByCode[cols[idx]].Name; c != "" {
					i.Nationality2 = cols[idx]
				} else if c := constants.CountriesByName[cols[idx]].Name; c != "" {
					i.Nationality2 = constants.CountriesByName[cols[idx]].ISO3166Alpha3
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\". valid values adhere to the ISO3166Alpha3 norm", constants.FileColumnIndividualNationality2, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualPhoneNumber1:
			i.PhoneNumber1 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber2:
			i.PhoneNumber2 = cols[idx]
		case constants.FileColumnIndividualPhoneNumber3:
			i.PhoneNumber3 = cols[idx]
		case constants.FileColumnIndividualPreferredContactMethod:
			preferredContactMethod, err := enumTypes.ParseContactMethod(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualPreferredContactMethod, err, enumTypes.AllContactMethods().String()))
				break
			}
			i.PreferredContactMethod = preferredContactMethod
		case constants.FileColumnIndividualPreferredContactMethodComments:
			i.PreferredContactMethodComments = cols[idx]
		case constants.FileColumnIndividualPreferredName:
			i.PreferredName = cols[idx]
		case constants.FileColumnIndividualPreferredCommunicationLanguage:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.PreferredCommunicationLanguage = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.PreferredCommunicationLanguage = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualPreferredCommunicationLanguage, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualPrefersToRemainAnonymous:
			prefersToRemainAnonymous, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualPrefersToRemainAnonymous, err))
			}
			i.PrefersToRemainAnonymous = prefersToRemainAnonymous
		case constants.FileColumnIndividualPresentsProtectionConcerns:
			presentsProtectionConcerns, err := getValidatedBoolean(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualPresentsProtectionConcerns, err))
			}
			i.PresentsProtectionConcerns = presentsProtectionConcerns
		case constants.FileColumnIndividualPWDComments:
			i.PWDComments = cols[idx]
		case constants.FileColumnIndividualSelfCareDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualSelfCareDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.SelfCareDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualSpokenLanguage1:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage1 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage1 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualSpokenLanguage1, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualSpokenLanguage2:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage2 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage2 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualSpokenLanguage2, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualSpokenLanguage3:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage3 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage3 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errors = append(errors, fmt.Errorf("%s: invalid value \"%s\"", constants.FileColumnIndividualSpokenLanguage3, cols[idx]))
					break
				}
			}
		case constants.FileColumnIndividualVisionDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualVisionDisabilityLevel, err, enumTypes.AllDisabilityLevels().String()))
				break
			}
			i.VisionDisabilityLevel = disabilityLevel
		case constants.FileColumnIndividualServiceCC1:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC1, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC1 = cc
		case constants.FileColumnIndividualServiceRequestedDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate1, err))
				break
			}
			i.ServiceRequestedDate1 = date
		case constants.FileColumnIndividualServiceDeliveredDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate1, err))
				break
			}
			i.ServiceDeliveredDate1 = date
		case constants.FileColumnIndividualServiceComments1:
			i.ServiceComments1 = cols[idx]
		case constants.FileColumnIndividualServiceCC2:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC2, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC2 = cc
		case constants.FileColumnIndividualServiceRequestedDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate2, err))
				break
			}
			i.ServiceRequestedDate2 = date
		case constants.FileColumnIndividualServiceDeliveredDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate2, err))
				break
			}
			i.ServiceDeliveredDate2 = date
		case constants.FileColumnIndividualServiceComments2:
			i.ServiceComments2 = cols[idx]
		case constants.FileColumnIndividualServiceCC3:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC3, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC3 = cc
		case constants.FileColumnIndividualServiceRequestedDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate3, err))
				break
			}
			i.ServiceRequestedDate3 = date
		case constants.FileColumnIndividualServiceDeliveredDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate3, err))
				break
			}
			i.ServiceDeliveredDate3 = date
		case constants.FileColumnIndividualServiceComments3:
			i.ServiceComments3 = cols[idx]
		case constants.FileColumnIndividualServiceCC4:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC4, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC4 = cc
		case constants.FileColumnIndividualServiceRequestedDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate4, err))
				break
			}
			i.ServiceRequestedDate4 = date
		case constants.FileColumnIndividualServiceDeliveredDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate4, err))
				break
			}
			i.ServiceDeliveredDate4 = date
		case constants.FileColumnIndividualServiceComments4:
			i.ServiceComments4 = cols[idx]
		case constants.FileColumnIndividualServiceCC5:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC5, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC5 = cc
		case constants.FileColumnIndividualServiceRequestedDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate5, err))
				break
			}
			i.ServiceRequestedDate5 = date
		case constants.FileColumnIndividualServiceDeliveredDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate5, err))
				break
			}
			i.ServiceDeliveredDate5 = date
		case constants.FileColumnIndividualServiceComments5:
			i.ServiceComments5 = cols[idx]
		case constants.FileColumnIndividualServiceCC6:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC6, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC6 = cc
		case constants.FileColumnIndividualServiceRequestedDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate6, err))
				break
			}
			i.ServiceRequestedDate6 = date
		case constants.FileColumnIndividualServiceDeliveredDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate6, err))
				break
			}
			i.ServiceDeliveredDate6 = date
		case constants.FileColumnIndividualServiceComments6:
			i.ServiceComments6 = cols[idx]
		case constants.FileColumnIndividualServiceCC7:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w. valid values are %s", constants.FileColumnIndividualServiceCC7, err, enumTypes.AllServiceCCs().String()))
				break
			}
			i.ServiceCC7 = cc
		case constants.FileColumnIndividualServiceRequestedDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate7, err))
				break
			}
			i.ServiceRequestedDate7 = date
		case constants.FileColumnIndividualServiceDeliveredDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errors = append(errors, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate7, err))
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
			if (field == constants.DBColumnIndividualNationality1 || field == constants.DBColumnIndividualNationality2) && v != "" {
				row[j] = constants.CountriesByCode[v].Name
				break
			}
			if (field == constants.DBColumnIndividualSpokenLanguage1 || field == constants.DBColumnIndividualSpokenLanguage2 || field == constants.DBColumnIndividualSpokenLanguage3 || field == constants.DBColumnIndividualPreferredCommunicationLanguage) && v != "" {
				row[j] = constants.LanguagesByCode[v].Name
				break
			}
			row[j] = v
		case time.Time:
			row[j] = v.Format(getTimeFormatForField(field))
		case *time.Time:
			if v != nil {
				row[j] = v.Format(getTimeFormatForField(field))
			}
		case enumTypes.DisabilityLevel:
			row[j] = string(v)
		case enumTypes.DisplacementStatus:
			row[j] = string(v)
		case enumTypes.ServiceCC:
			row[j] = string(v)
		case enumTypes.Sex:
			row[j] = string(v)
		default:
			row[j] = fmt.Sprintf("%v", v)
		}
	}
	return row, nil
}
