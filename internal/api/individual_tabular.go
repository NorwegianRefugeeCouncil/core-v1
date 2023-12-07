package api

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/pkg/logutils"
	"golang.org/x/exp/slices"
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
		t := locales.GetTranslator()
		fileNameParts := strings.Split(filename, ".")
		fileType := fileNameParts[len(fileNameParts)-1]
		err := errors.New(t("error_file_type", fileType))
		return err
	}
}

func GetColumnMapping(header []string, fields *[]string) (map[string]int, []FileError) {
	headerInternal := locales.GetTranslationKeys(header)
	colMapping := map[string]int{}
	errs := []error{}
	for i, col := range headerInternal {
		field, ok := constants.IndividualFileToDBMap[col]
		if !ok {
			ok = slices.Contains(constants.IndividualSystemFileColumns, col)
			if ok {
				continue
			}
			errs = append(errs, errors.New(locales.GetTranslator()("error_unknown_column_detail", logutils.Escape(col))))
		}
		*fields = append(*fields, field)
		colMapping[col] = i
	}
	if len(errs) > 0 {
		t := locales.GetTranslator()
		return nil, []FileError{{
			Message: t("error_unknown_column"),
			Err:     errs,
		}}
	}
	return colMapping, nil
}

func UnmarshalIndividualsTabularData(data [][]string, individuals *[]*Individual, colMapping map[string]int, rowLimit *int) []FileError {
	if rowLimit != nil && len(data[1:]) > *rowLimit {
		return []FileError{{locales.GetTranslator()("error_upload_limit", len(data[1:]), *rowLimit), nil}}
	}
	var fileErrors []FileError

	for row, cols := range data[1:] {
		individual := &Individual{}
		var rowErrors []error
		for _, err := range individual.unmarshalTabularData(colMapping, cols) {
			rowErrors = append(rowErrors, err)
		}
		if len(rowErrors) > 0 {
			t := locales.GetTranslator()
			fileErrors = append(fileErrors, FileError{
				Message: t("error_row_parse_fail", row+2),
				Err:     rowErrors,
			})
		}
		*individuals = append(*individuals, individual)
	}

	return fileErrors
}

func (i *Individual) unmarshalTabularData(colMapping map[string]int, cols []string) []error {
	var errs []error
	t := locales.GetTranslator()
	if len(cols) <= len(colMapping) {
		filler := make([]string, len(colMapping)-len(cols))
		cols = append(cols, filler...)
	}
	for field, idx := range colMapping {
		switch field {
		case constants.DBColumnIndividualID:
			i.ID = cols[idx]
		case constants.DBColumnIndividualInactive:
			i.Inactive = isExplicitlyTrue(cols[idx])
		case constants.DBColumnIndividualAddress:
			i.Address = cols[idx]
		case constants.DBColumnIndividualAge:
			age, err := ParseAge(cols[idx])
			if err != nil {
				errs = append(errs, err)
				break
			}
			i.Age = age
		case constants.DBColumnIndividualBirthDate:
			var birthDate *time.Time
			birthDate, err := ParseBirthdate(cols[idx])
			if err != nil {
				errs = append(errs, err)
				break
			}
			i.BirthDate = birthDate
		case constants.DBColumnIndividualCognitiveDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualCognitiveDisabilityLevel, err, enumTypes.AllDisabilityLevels().String())))
				break
			}
			i.CognitiveDisabilityLevel = disabilityLevel
		case constants.DBColumnIndividualCollectionAdministrativeArea1:
			i.CollectionAdministrativeArea1 = cols[idx]
		case constants.DBColumnIndividualCollectionAdministrativeArea2:
			i.CollectionAdministrativeArea2 = cols[idx]
		case constants.DBColumnIndividualCollectionAdministrativeArea3:
			i.CollectionAdministrativeArea3 = cols[idx]
		case constants.DBColumnIndividualCollectionOffice:
			i.CollectionOffice = cols[idx]
		case constants.DBColumnIndividualCollectionAgentName:
			i.CollectionAgentName = cols[idx]
		case constants.DBColumnIndividualCollectionAgentTitle:
			i.CollectionAgentTitle = cols[idx]
		case constants.DBColumnIndividualComments:
			i.Comments = cols[idx]
		case constants.DBColumnIndividualCollectionTime:
			var collectionTime *time.Time
			collectionTime, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualCollectionTime, err))
				break
			}
			if collectionTime != nil {
				i.CollectionTime = *collectionTime
			}
		case constants.DBColumnIndividualCommunicationDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualCommunicationDisabilityLevel, err, enumTypes.AllDisabilityLevels().String())))
				break
			}
			i.CommunicationDisabilityLevel = disabilityLevel
		case constants.DBColumnIndividualCommunityID:
			i.CommunityID = cols[idx]
		case constants.DBColumnIndividualCommunitySize:
			var communitySizeStr = cols[idx]
			if communitySizeStr == "" {
				continue
			}
			communitySize, err := strconv.Atoi(communitySizeStr)
			if err != nil {
				errs = append(errs, err)
				break
			}
			i.CommunitySize = &communitySize
		case constants.DBColumnIndividualDisplacementStatus:
			displacementStatus, err := enumTypes.ParseDisplacementStatus(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualDisplacementStatus, err, enumTypes.AllDisplacementStatuses().String())))
				break
			}
			i.DisplacementStatus = displacementStatus
		case constants.DBColumnIndividualDisplacementStatusComment:
			i.DisplacementStatusComment = cols[idx]
		case constants.DBColumnIndividualEmail1:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail1, err))
					break
				}
				i.Email1 = email.Address
			}
		case constants.DBColumnIndividualEmail2:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail2, err))
					break
				}
				i.Email2 = email.Address
			}
		case constants.DBColumnIndividualEmail3:
			if cols[idx] != "" {
				email, err := mail.ParseAddress(cols[idx])
				if err != nil {
					errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualEmail3, err))
					break
				}
				i.Email3 = email.Address
			}
		case constants.DBColumnIndividualFullName:
			i.FullName = cols[idx]
		case constants.DBColumnIndividualFirstName:
			i.FirstName = cols[idx]
		case constants.DBColumnIndividualMiddleName:
			i.MiddleName = cols[idx]
		case constants.DBColumnIndividualLastName:
			i.LastName = cols[idx]
		case constants.DBColumnIndividualNativeName:
			i.NativeName = cols[idx]
		case constants.DBColumnIndividualMothersName:
			i.MothersName = cols[idx]
		case constants.DBColumnIndividualFreeField1:
			i.FreeField1 = cols[idx]
		case constants.DBColumnIndividualFreeField2:
			i.FreeField2 = cols[idx]
		case constants.DBColumnIndividualFreeField3:
			i.FreeField3 = cols[idx]
		case constants.DBColumnIndividualFreeField4:
			i.FreeField4 = cols[idx]
		case constants.DBColumnIndividualFreeField5:
			i.FreeField5 = cols[idx]
		case constants.DBColumnIndividualSex:
			sex, err := enumTypes.ParseSex(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualSex, err, enumTypes.AllSexes().String())))
				break
			}
			i.Sex = sex
		case constants.DBColumnIndividualHasMedicalCondition:
			hasMedicalCondition, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasMedicalCondition = hasMedicalCondition.BoolPtr()
		case constants.DBColumnIndividualNeedsLegalAndPhysicalProtection:
			needsLegalAndPhysicalProtection, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.NeedsLegalAndPhysicalProtection = needsLegalAndPhysicalProtection.BoolPtr()
		case constants.DBColumnIndividualIsChildAtRisk:
			isChildAtRisk, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsChildAtRisk = isChildAtRisk.BoolPtr()
		case constants.DBColumnIndividualIsWomanAtRisk:
			isWomanAtRisk, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsWomanAtRisk = isWomanAtRisk.BoolPtr()
		case constants.DBColumnIndividualIsElderAtRisk:
			isElderAtRisk, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsElderAtRisk = isElderAtRisk.BoolPtr()
		case constants.DBColumnIndividualIsLactating:
			isLactating, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsLactating = isLactating.BoolPtr()
		case constants.DBColumnIndividualIsPregnant:
			isPregnant, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsPregnant = isPregnant.BoolPtr()
		case constants.DBColumnIndividualIsSingleParent:
			isSingleParent, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsSingleParent = isSingleParent.BoolPtr()
		case constants.DBColumnIndividualIsSeparatedChild:
			isSeparatedChild, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsSeparatedChild = isSeparatedChild.BoolPtr()
		case constants.DBColumnIndividualHasCognitiveDisability:
			hasCognitiveDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCognitiveDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasCognitiveDisability = hasCognitiveDisability.BoolPtr()
		case constants.DBColumnIndividualHasCommunicationDisability:
			hasCommunicationDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasCommunicationDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasCommunicationDisability = hasCommunicationDisability.BoolPtr()
		case constants.DBColumnIndividualHasConsentedToRGPD:
			hasConsentedToRGPD, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasConsentedToRGPD, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasConsentedToRGPD = hasConsentedToRGPD.BoolPtr()
		case constants.DBColumnIndividualHasConsentedToReferral:
			hasConsentedToReferral, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasConsentedToReferral, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasConsentedToReferral = hasConsentedToReferral.BoolPtr()
		case constants.DBColumnIndividualHasDisability:
			hasDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasDisability = hasDisability.BoolPtr()
		case constants.DBColumnIndividualHasHearingDisability:
			hasHearingDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasHearingDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasHearingDisability = hasHearingDisability.BoolPtr()
		case constants.DBColumnIndividualHasMobilityDisability:
			hasMobilityDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasMobilityDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasMobilityDisability = hasMobilityDisability.BoolPtr()
		case constants.DBColumnIndividualHasSelfCareDisability:
			hasSelfCareDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasSelfCareDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasSelfCareDisability = hasSelfCareDisability.BoolPtr()
		case constants.DBColumnIndividualHasVisionDisability:
			hasVisionDisability, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHasVisionDisability, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.HasVisionDisability = hasVisionDisability.BoolPtr()
		case constants.DBColumnIndividualHearingDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualHearingDisabilityLevel, err, enumTypes.AllDisabilityLevels().String())))
				break
			}
			i.HearingDisabilityLevel = disabilityLevel
		case constants.DBColumnIndividualHouseholdID:
			i.HouseholdID = cols[idx]
		case constants.DBColumnIndividualHouseholdSize:
			var householdSizeStr = cols[idx]
			if householdSizeStr == "" {
				continue
			}
			householdSize, err := strconv.Atoi(householdSizeStr)
			if err != nil {
				errs = append(errs, err)
				break
			}
			i.HouseholdSize = &householdSize
		case constants.DBColumnIndividualIdentificationType1:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIdentificationType1, err, enumTypes.AllIdentificationTypes().String())))
				break
			}
			i.IdentificationType1 = identificationType
		case constants.DBColumnIndividualIdentificationTypeExplanation1:
			i.IdentificationTypeExplanation1 = cols[idx]
		case constants.DBColumnIndividualIdentificationNumber1:
			i.IdentificationNumber1 = cols[idx]
		case constants.DBColumnIndividualIdentificationType2:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIdentificationType2, err, enumTypes.AllIdentificationTypes().String())))
				break
			}
			i.IdentificationType2 = identificationType
		case constants.DBColumnIndividualIdentificationTypeExplanation2:
			i.IdentificationTypeExplanation2 = cols[idx]
		case constants.DBColumnIndividualIdentificationNumber2:
			i.IdentificationNumber2 = cols[idx]
		case constants.DBColumnIndividualIdentificationType3:
			identificationType, err := enumTypes.ParseIdentificationType(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIdentificationType3, err, enumTypes.AllIdentificationTypes().String())))
				break
			}
			i.IdentificationType3 = identificationType
		case constants.DBColumnIndividualIdentificationTypeExplanation3:
			i.IdentificationTypeExplanation3 = cols[idx]
		case constants.DBColumnIndividualIdentificationNumber3:
			i.IdentificationNumber3 = cols[idx]
		case constants.DBColumnIndividualEngagementContext:
			engagementContext, err := enumTypes.ParseEngagementContext(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualEngagementContext, err, enumTypes.AllEngagementContexts().String())))
				break
			}
			i.EngagementContext = engagementContext
		case constants.DBColumnIndividualInternalID:
			i.InternalID = cols[idx]
		case constants.DBColumnIndividualIsHeadOfCommunity:
			isHeadOfCommunity, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIsHeadOfCommunity, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsHeadOfCommunity = isHeadOfCommunity.BoolPtr()
		case constants.DBColumnIndividualIsHeadOfHousehold:
			isHeadOfHousehold, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIsHeadOfHousehold, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsHeadOfHousehold = isHeadOfHousehold.BoolPtr()
		case constants.DBColumnIndividualIsFemaleHeadedHousehold:
			isFemaleHeadedHousehold, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIsFemaleHeadedHousehold, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsFemaleHeadedHousehold = isFemaleHeadedHousehold.BoolPtr()
		case constants.DBColumnIndividualIsMinorHeadedHousehold:
			isMinorHeadedHousehold, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIsMinorHeadedHousehold, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsMinorHeadedHousehold = isMinorHeadedHousehold.BoolPtr()
		case constants.DBColumnIndividualIsMinor:
			isMinor, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualIsMinor, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.IsMinor = isMinor.BoolPtr()
		case constants.DBColumnIndividualMobilityDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualMobilityDisabilityLevel, err, enumTypes.AllDisabilityLevels().String())))
				break
			}
			i.MobilityDisabilityLevel = disabilityLevel
		case constants.DBColumnIndividualNationality1:
			if cols[idx] != "" {
				if c := constants.CountriesByCode[cols[idx]].Name; c != "" {
					i.Nationality1 = cols[idx]
				} else if c := constants.CountriesByName[cols[idx]].Name; c != "" {
					i.Nationality1 = constants.CountriesByName[cols[idx]].ISO3166Alpha3
				} else {
					errs = append(errs, errors.New(t("error_invalid_value_nationality_hint", constants.FileColumnIndividualNationality1, cols[idx])))
					break
				}
			}
		case constants.DBColumnIndividualNationality2:
			if cols[idx] != "" {
				if c := constants.CountriesByCode[cols[idx]].Name; c != "" {
					i.Nationality2 = cols[idx]
				} else if c := constants.CountriesByName[cols[idx]].Name; c != "" {
					i.Nationality2 = constants.CountriesByName[cols[idx]].ISO3166Alpha3
				} else {
					errs = append(errs, errors.New(t("error_invalid_value_nationality_hint", constants.FileColumnIndividualNationality2, cols[idx])))
					break
				}
			}
		case constants.DBColumnIndividualPhoneNumber1:
			i.PhoneNumber1 = cols[idx]
		case constants.DBColumnIndividualPhoneNumber2:
			i.PhoneNumber2 = cols[idx]
		case constants.DBColumnIndividualPhoneNumber3:
			i.PhoneNumber3 = cols[idx]
		case constants.DBColumnIndividualPreferredContactMethod:
			preferredContactMethod, err := enumTypes.ParseContactMethod(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualPreferredContactMethod, err, enumTypes.AllContactMethods().String())))
				break
			}
			i.PreferredContactMethod = preferredContactMethod
		case constants.DBColumnIndividualPreferredContactMethodComments:
			i.PreferredContactMethodComments = cols[idx]
		case constants.DBColumnIndividualPreferredName:
			i.PreferredName = cols[idx]
		case constants.DBColumnIndividualPreferredCommunicationLanguage:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.PreferredCommunicationLanguage = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.PreferredCommunicationLanguage = constants.LanguagesByName[cols[idx]].ID
				} else {
					errs = append(errs, errors.New(t("error_invalid_value", constants.DBColumnIndividualPreferredCommunicationLanguage, cols[idx])))
					break
				}
			}
		case constants.DBColumnIndividualPrefersToRemainAnonymous:
			prefersToRemainAnonymous, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualPrefersToRemainAnonymous, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.PrefersToRemainAnonymous = prefersToRemainAnonymous.BoolPtr()
		case constants.DBColumnIndividualPresentsProtectionConcerns:
			presentsProtectionConcerns, err := enumTypes.ParseOptionalBoolean(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualPresentsProtectionConcerns, err, enumTypes.AllOptionalBooleans().String())))
			}
			i.PresentsProtectionConcerns = presentsProtectionConcerns.BoolPtr()
		case constants.DBColumnIndividualPWDComments:
			i.PWDComments = cols[idx]
		case constants.DBColumnIndividualVulnerabilityComments:
			i.VulnerabilityComments = cols[idx]
		case constants.DBColumnIndividualSelfCareDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualSelfCareDisabilityLevel, err, enumTypes.AllDisabilityLevels().String())))
				break
			}
			i.SelfCareDisabilityLevel = disabilityLevel
		case constants.DBColumnIndividualSpokenLanguage1:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage1 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage1 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errs = append(errs, errors.New(t("error_invalid_value", constants.DBColumnIndividualSpokenLanguage1, cols[idx])))
					break
				}
			}
		case constants.DBColumnIndividualSpokenLanguage2:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage2 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage2 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errs = append(errs, errors.New(t("error_invalid_value", constants.DBColumnIndividualSpokenLanguage2, cols[idx])))
					break
				}
			}
		case constants.DBColumnIndividualSpokenLanguage3:
			if cols[idx] != "" {
				if l := constants.LanguagesByCode[cols[idx]].Name; l != "" {
					i.SpokenLanguage3 = cols[idx]
				} else if l := constants.LanguagesByName[cols[idx]].Name; l != "" {
					i.SpokenLanguage3 = constants.LanguagesByName[cols[idx]].ID
				} else {
					errs = append(errs, errors.New(t("error_invalid_value", constants.DBColumnIndividualSpokenLanguage3, cols[idx])))
					break
				}
			}
		case constants.DBColumnIndividualVisionDisabilityLevel:
			disabilityLevel, err := enumTypes.ParseDisabilityLevel(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualVisionDisabilityLevel, err, enumTypes.AllDisabilityLevels().String())))
				break
			}
			i.VisionDisabilityLevel = disabilityLevel
		case constants.DBColumnIndividualServiceCC1:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC1, err, enumTypes.AllServiceCCs().String())))
				break
			}
			i.ServiceCC1 = cc
		case constants.DBColumnIndividualServiceRequestedDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate1, err))
				break
			}
			i.ServiceRequestedDate1 = date
		case constants.DBColumnIndividualServiceDeliveredDate1:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate1, err))
				break
			}
			i.ServiceDeliveredDate1 = date
		case constants.DBColumnIndividualServiceComments1:
			i.ServiceComments1 = cols[idx]
		case constants.DBColumnIndividualServiceCC2:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC2, err, enumTypes.AllServiceCCs().String())))
				break
			}
			i.ServiceCC2 = cc
		case constants.DBColumnIndividualServiceRequestedDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate2, err))
				break
			}
			i.ServiceRequestedDate2 = date
		case constants.DBColumnIndividualServiceDeliveredDate2:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate2, err))
				break
			}
			i.ServiceDeliveredDate2 = date
		case constants.DBColumnIndividualServiceComments2:
			i.ServiceComments2 = cols[idx]
		case constants.DBColumnIndividualServiceCC3:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC3, err, enumTypes.AllServiceCCs().String())))
				break
			}
			i.ServiceCC3 = cc
		case constants.DBColumnIndividualServiceRequestedDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate3, err))
				break
			}
			i.ServiceRequestedDate3 = date
		case constants.DBColumnIndividualServiceDeliveredDate3:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate3, err))
				break
			}
			i.ServiceDeliveredDate3 = date
		case constants.DBColumnIndividualServiceComments3:
			i.ServiceComments3 = cols[idx]
		case constants.DBColumnIndividualServiceCC4:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC4, err, enumTypes.AllServiceCCs().String())))
				break
			}
			i.ServiceCC4 = cc
		case constants.DBColumnIndividualServiceRequestedDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate4, err))
				break
			}
			i.ServiceRequestedDate4 = date
		case constants.DBColumnIndividualServiceDeliveredDate4:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate4, err))
				break
			}
			i.ServiceDeliveredDate4 = date
		case constants.DBColumnIndividualServiceComments4:
			i.ServiceComments4 = cols[idx]
		case constants.DBColumnIndividualServiceCC5:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC5, err, enumTypes.AllServiceCCs().String())))
				break
			}
			i.ServiceCC5 = cc
		case constants.DBColumnIndividualServiceRequestedDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate5, err))
				break
			}
			i.ServiceRequestedDate5 = date
		case constants.DBColumnIndividualServiceDeliveredDate5:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate5, err))
				break
			}
			i.ServiceDeliveredDate5 = date
		case constants.DBColumnIndividualServiceComments5:
			i.ServiceComments5 = cols[idx]
		case constants.DBColumnIndividualServiceCC6:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC6, err, enumTypes.AllServiceCCs().String())))
				break
			}
			i.ServiceCC6 = cc
		case constants.DBColumnIndividualServiceRequestedDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate6, err))
				break
			}
			i.ServiceRequestedDate6 = date
		case constants.DBColumnIndividualServiceDeliveredDate6:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate6, err))
				break
			}
			i.ServiceDeliveredDate6 = date
		case constants.DBColumnIndividualServiceComments6:
			i.ServiceComments6 = cols[idx]
		case constants.DBColumnIndividualServiceCC7:
			cc, err := enumTypes.ParseServiceCC(cols[idx])
			if err != nil {
				errs = append(errs, errors.New(t("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC7, err, enumTypes.AllServiceCCs().String())))
				break
			}
			i.ServiceCC7 = cc
		case constants.DBColumnIndividualServiceRequestedDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceRequestedDate7, err))
				break
			}
			i.ServiceRequestedDate7 = date
		case constants.DBColumnIndividualServiceDeliveredDate7:
			var date *time.Time
			date, err := ParseDate(cols[idx])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", constants.FileColumnIndividualServiceDeliveredDate7, err))
				break
			}
			i.ServiceDeliveredDate7 = date
		case constants.DBColumnIndividualServiceComments7:
			i.ServiceComments7 = cols[idx]
		}
	}
	if len(errs) > 0 {
		return errs
	}
	i.Normalize()
	return nil
}

// Marshal

func MarshalIndividualsCSV(w io.Writer, individuals []*Individual) error {
	csvEncoder := csv.NewWriter(w)
	defer csvEncoder.Flush()

	if err := csvEncoder.Write(locales.TranslateSlice(constants.IndividualFileColumns)); err != nil {
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
	//c := constants.FileColumnsLocalized

	if err := streamWriter.SetRow("A1", stringArrayToInterfaceArray(locales.TranslateSlice(constants.IndividualFileColumns))); err != nil {
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
		field, ok := constants.IndividualFileToDBMap[col]
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
		case enumTypes.OptionalBoolean:
			row[j] = string(v)
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
		case enumTypes.EngagementContext:
			row[j] = string(v)
		case enumTypes.IdentificationType:
			row[j] = string(v)
		case enumTypes.ContactMethod:
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
