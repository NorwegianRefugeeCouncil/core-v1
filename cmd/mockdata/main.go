package mockdata

import (
	"encoding/csv"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/manveru/faker"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/utils"
)

func randomCountry() string {
	idx := rand.Intn(len(constants.Countries))
	return constants.Countries[idx].ISO3166Alpha3
}

func randomLanguage() string {
	idx := rand.Intn(len(constants.Languages))
	return constants.Languages[idx].ID
}

func randomIdentificationType() string {
	return pick(
		"",
		"passport",
		"national_id",
		"other",
	)
}

func randomDisabilityLevel() string {
	var g []string
	for _, d := range enumTypes.AllDisabilityLevels().Items() {
		g = append(g, string(d))
	}
	return pick(g...)
}

func randomSex() string {
	var g []string
	for _, d := range enumTypes.AllSexes().Items() {
		g = append(g, string(d))
	}
	return pick(g...)
}

func randomDisplacementStatus() string {
	var ds []string
	for _, d := range enumTypes.AllDisplacementStatuses().Items() {
		ds = append(ds, string(d))
	}
	return pick(ds...)
}

func randomDate() string {
	now := time.Now()
	yearFrom := now.Year() - rand.Intn(10) - 1
	yearTo := now.Year()
	year := rand.Intn(yearTo-yearFrom) + yearFrom

	month := rand.Intn(11) + 1
	day := rand.Intn(27) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
}

func pick(items ...string) string {
	return items[rand.Intn(len(items))]
}

func randomBool() string {
	if randBool(50) {
		return "true"
	}
	return "false"
}

func randomOptionalBool() string {
	r := rand.Intn(100)
	if r > 66 {
		return "true"
	} else if r > 33 {
		return "false"
	}
	return ""
}

func randomEngagementContext() string {
	return pick(
		"",
		"houseVisit",
		"fieldActivity",
		"inOffice",
		"remoteChannels",
		"other",
	)
}

func randomContactMethod() string {
	return pick(
		"phone",
		"email",
		"whatsapp",
		"visit",
		"other",
	)
}

func Generate(count uint) error {

	var householdIds []string
	var householdCount = utils.Max(2, int(count)/5)
	for i := 0; i < householdCount; i++ {
		householdIds = append(householdIds, strconv.Itoa(rand.Intn(1000000)))
	}

	var communityIds []string
	var communityIdCount = utils.Max(2, int(count)/100)
	for i := 0; i < communityIdCount; i++ {
		communityIds = append(communityIds, strconv.Itoa(rand.Intn(1000000)))
	}

	f, err := faker.New("en")
	if err != nil {
		return err
	}

	file, err := os.Create("generated.csv")
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	if err := writer.Write([]string{
		constants.DBColumnIndividualAddress,
		constants.DBColumnIndividualAge,
		constants.DBColumnIndividualBirthDate,
		constants.DBColumnIndividualCognitiveDisabilityLevel,
		constants.DBColumnIndividualCollectionAdministrativeArea1,
		constants.DBColumnIndividualCollectionAdministrativeArea2,
		constants.DBColumnIndividualCollectionAdministrativeArea3,
		constants.DBColumnIndividualCollectionOffice,
		constants.DBColumnIndividualCollectionAgentName,
		constants.DBColumnIndividualCollectionAgentTitle,
		constants.DBColumnIndividualCollectionTime,
		constants.DBColumnIndividualComments,
		constants.DBColumnIndividualCommunicationDisabilityLevel,
		constants.DBColumnIndividualCommunityID,
		constants.DBColumnIndividualCommunitySize,
		constants.DBColumnIndividualDisplacementStatus,
		constants.DBColumnIndividualDisplacementStatusComment,
		constants.DBColumnIndividualEmail1,
		constants.DBColumnIndividualEmail2,
		constants.DBColumnIndividualEmail3,
		constants.DBColumnIndividualFullName,
		constants.DBColumnIndividualFirstName,
		constants.DBColumnIndividualMiddleName,
		constants.DBColumnIndividualLastName,
		constants.DBColumnIndividualNativeName,
		constants.DBColumnIndividualMothersName,
		constants.DBColumnIndividualFreeField1,
		constants.DBColumnIndividualFreeField2,
		constants.DBColumnIndividualFreeField3,
		constants.DBColumnIndividualFreeField4,
		constants.DBColumnIndividualFreeField5,
		constants.DBColumnIndividualSex,
		constants.DBColumnIndividualHasCognitiveDisability,
		constants.DBColumnIndividualHasCommunicationDisability,
		constants.DBColumnIndividualHasConsentedToRGPD,
		constants.DBColumnIndividualHasConsentedToReferral,
		constants.DBColumnIndividualHasDisability,
		constants.DBColumnIndividualHasHearingDisability,
		constants.DBColumnIndividualHasMobilityDisability,
		constants.DBColumnIndividualHasSelfCareDisability,
		constants.DBColumnIndividualHasVisionDisability,
		constants.DBColumnIndividualHearingDisabilityLevel,
		constants.DBColumnIndividualHouseholdID,
		constants.DBColumnIndividualHouseholdSize,
		constants.DBColumnIndividualIdentificationType1,
		constants.DBColumnIndividualIdentificationTypeExplanation1,
		constants.DBColumnIndividualIdentificationNumber1,
		constants.DBColumnIndividualIdentificationType2,
		constants.DBColumnIndividualIdentificationTypeExplanation2,
		constants.DBColumnIndividualIdentificationNumber2,
		constants.DBColumnIndividualIdentificationType3,
		constants.DBColumnIndividualIdentificationTypeExplanation3,
		constants.DBColumnIndividualIdentificationNumber3,
		constants.DBColumnIndividualEngagementContext,
		constants.DBColumnIndividualInternalID,
		constants.DBColumnIndividualIsHeadOfCommunity,
		constants.DBColumnIndividualIsHeadOfHousehold,
		constants.DBColumnIndividualIsFemaleHeadedHousehold,
		constants.DBColumnIndividualIsMinorHeadedHousehold,
		constants.DBColumnIndividualIsMinor,
		constants.DBColumnIndividualIsChildAtRisk,
		constants.DBColumnIndividualIsWomanAtRisk,
		constants.DBColumnIndividualIsElderAtRisk,
		constants.DBColumnIndividualIsSeparatedChild,
		constants.DBColumnIndividualIsSingleParent,
		constants.DBColumnIndividualIsPregnant,
		constants.DBColumnIndividualIsLactating,
		constants.DBColumnIndividualHasMedicalCondition,
		constants.DBColumnIndividualNeedsLegalAndPhysicalProtection,
		constants.DBColumnIndividualVulnerabilityComments,
		constants.DBColumnIndividualMobilityDisabilityLevel,
		constants.DBColumnIndividualNationality1,
		constants.DBColumnIndividualNationality2,
		constants.DBColumnIndividualPhoneNumber1,
		constants.DBColumnIndividualPhoneNumber2,
		constants.DBColumnIndividualPhoneNumber3,
		constants.DBColumnIndividualPreferredContactMethod,
		constants.DBColumnIndividualPreferredContactMethodComments,
		constants.DBColumnIndividualPreferredName,
		constants.DBColumnIndividualPreferredCommunicationLanguage,
		constants.DBColumnIndividualPrefersToRemainAnonymous,
		constants.DBColumnIndividualPresentsProtectionConcerns,
		constants.DBColumnIndividualPWDComments,
		constants.DBColumnIndividualSelfCareDisabilityLevel,
		constants.DBColumnIndividualSpokenLanguage1,
		constants.DBColumnIndividualSpokenLanguage2,
		constants.DBColumnIndividualSpokenLanguage3,
		constants.DBColumnIndividualVisionDisabilityLevel,
		constants.DBColumnIndividualServiceCC1,
		constants.DBColumnIndividualServiceRequestedDate1,
		constants.DBColumnIndividualServiceDeliveredDate1,
		constants.DBColumnIndividualServiceCC2,
		constants.DBColumnIndividualServiceRequestedDate2,
		constants.DBColumnIndividualServiceDeliveredDate2,
		constants.DBColumnIndividualServiceCC3,
		constants.DBColumnIndividualServiceRequestedDate3,
		constants.DBColumnIndividualServiceDeliveredDate3,
		constants.DBColumnIndividualServiceCC4,
		constants.DBColumnIndividualServiceRequestedDate4,
		constants.DBColumnIndividualServiceDeliveredDate4,
		constants.DBColumnIndividualServiceCC5,
		constants.DBColumnIndividualServiceRequestedDate5,
		constants.DBColumnIndividualServiceDeliveredDate5,
		constants.DBColumnIndividualServiceCC6,
		constants.DBColumnIndividualServiceRequestedDate6,
		constants.DBColumnIndividualServiceDeliveredDate6,
		constants.DBColumnIndividualServiceCC7,
		constants.DBColumnIndividualServiceRequestedDate7,
		constants.DBColumnIndividualServiceDeliveredDate7,
		constants.DBColumnIndividualInactive,
	}); err != nil {
		return err
	}
	for i := 0; i < int(count); i++ {

		var address string
		if randBool(80) {
			address = f.StreetAddress() + "\n" + f.PostCode() + " " + f.City() + "\n" + f.Country()
		}
		var birthDate string
		if randBool(80) {
			start, _ := time.Parse("2006-01-02", "1900-01-01")
			end := time.Now()
			birthDate = start.Add(time.Duration(rand.Int63n(end.Unix()-start.Unix())) * time.Second).Format("2006-01-02")
		}

		var age string
		if randBool(80) {
			age = strconv.Itoa(rand.Intn(100))
		}

		collectionAdministrativeArea1 := f.Country()
		collectionAdministrativeArea2 := f.State()
		collectionAdministrativeArea3 := f.City()
		collectionOffice := f.City() + " Office"
		collectionAgentName := f.Name()
		collectionAgentTitle := f.JobTitle()
		collectionTime := randomDate()

		var comments string
		if randBool(50) {
			comments = randomText(f)
		}

		communityId := ""
		if randBool(80) {
			communityId = pick(communityIds...)
		}

		var communitySize string
		if randBool(80) {
			communitySize = strconv.Itoa(rand.Intn(10))
		}

		displacementStatus := randomDisplacementStatus()
		displacementStatusComment := ""

		if displacementStatus == "other" {
			displacementStatusComment = randomText(f)
		}
		hasDisability := randomOptionalBool()
		pwdComments := randomText(f)

		var email1 string
		if randBool(80) {
			email1 = f.Email()
		}
		var email2 string
		if email1 != "" && randBool(40) {
			email2 = f.Email()
		}
		var email3 string
		if email2 != "" && randBool(40) {
			email3 = f.Email()
		}

		firstName := f.FirstName()
		middleName := f.FirstName()
		lastName := f.LastName()
		nativeName := f.Name()
		var fullName = firstName + " " + middleName + " " + lastName
		var preferredName = firstName + " " + lastName
		if randBool(5) {
			preferredName = f.Name()
		}

		mothersName := f.Name()

		var freeField1 string
		if randBool(30) {
			freeField1 = strconv.Itoa(rand.Intn(1000000))
		}
		var freeField2 string
		if randBool(30) {
			freeField2 = strconv.Itoa(rand.Intn(1000000))
		}
		var freeField3 string
		if randBool(30) {
			freeField3 = strconv.Itoa(rand.Intn(1000000))
		}
		var freeField4 string
		if randBool(30) {
			freeField4 = strconv.Itoa(rand.Intn(1000000))
		}
		var freeField5 string
		if randBool(30) {
			freeField5 = strconv.Itoa(rand.Intn(1000000))
		}

		sex := randomSex()
		hasCognitiveDisability := randomOptionalBool()
		hasCommunicationDisability := randomOptionalBool()
		hasConsentedToRgpd := randomOptionalBool()
		hasConsentedToReferral := randomOptionalBool()
		hasHearingDisability := randomOptionalBool()
		hasMobilityDisability := randomOptionalBool()
		hasSelfCareDisability := randomOptionalBool()
		hasVisionDisability := randomOptionalBool()

		hearingDisabilityLevel := randomDisabilityLevel()
		if hasHearingDisability == "false" {
			hearingDisabilityLevel = "none"
		}
		cognitiveDisabilityLevel := randomDisabilityLevel()
		if hasCognitiveDisability == "false" {
			cognitiveDisabilityLevel = "none"
		}
		communicationDisabilityLevel := randomDisabilityLevel()
		if hasCommunicationDisability == "false" {
			communicationDisabilityLevel = "none"
		}
		mobilityDisabilityLevel := randomDisabilityLevel()
		if hasMobilityDisability == "false" {
			mobilityDisabilityLevel = "none"
		}
		selfCareDisabilityLevel := randomDisabilityLevel()
		if hasSelfCareDisability == "false" {
			selfCareDisabilityLevel = "none"
		}
		visionDisabilityLevel := randomDisabilityLevel()
		if hasVisionDisability == "false" {
			visionDisabilityLevel = "none"
		}

		householdId := ""
		if randBool(80) {
			householdId = pick(householdIds...)
		}

		var householdSize string
		if randBool(80) {
			householdSize = strconv.Itoa(rand.Intn(10))
		}

		identificationType1 := randomIdentificationType()
		identificationTypeExplanation1 := ""
		var identificationNumber1 string
		if len(identificationType1) != 0 {
			identificationNumber1 = strconv.Itoa(rand.Intn(1000000000))
		}
		if identificationType1 == "other" {
			identificationTypeExplanation1 = randomText(f)
		}

		identificationType2 := randomIdentificationType()
		identificationTypeExplanation2 := ""

		var identificationNumber2 string
		if len(identificationType2) != 0 {
			identificationNumber2 = strconv.Itoa(rand.Intn(1000000000))
		}
		if identificationType2 == "other" {
			identificationTypeExplanation2 = strings.Join(f.Paragraphs(rand.Intn(3)+1, false), "\n\n")
		}

		identificationType3 := randomIdentificationType()
		identificationTypeExplanation3 := ""

		var identificationNumber3 string
		if len(identificationType3) != 0 {
			identificationNumber3 = strconv.Itoa(rand.Intn(1000000000))
		}
		if identificationType3 == "other" {
			identificationTypeExplanation3 = strings.Join(f.Paragraphs(rand.Intn(3)+1, false), "\n\n")
		}

		engagementContext := randomEngagementContext()
		internalId := strconv.Itoa(rand.Intn(1000000000))
		isHeadOfCommunity := "false"
		if communityId != "" && randBool(5) {
			isHeadOfCommunity = "true"
		}

		isHeadOfHousehold := "false"
		if householdId != "" {
			isHeadOfHousehold = randomOptionalBool()
		}

		isMinorHeadedHousehold := randomOptionalBool()
		isFemaleHeadedHousehold := randomOptionalBool()

		isMinor := randomOptionalBool()
		isChildAtRisk := randomOptionalBool()
		isWomanAtRisk := randomOptionalBool()
		isElderAtRisk := randomOptionalBool()
		isSeparatedChild := randomOptionalBool()
		isSingleParent := randomOptionalBool()
		isPregnant := randomOptionalBool()
		isLactating := randomOptionalBool()
		hasMedicalCondition := randomOptionalBool()
		needsLegalAndPhysicalProtection := randomOptionalBool()
		vulnerabilityComments := randomText(f)

		nationality1 := randomCountry()
		nationality2 := randomCountry()
		var phoneNumber1 string
		if randBool(80) {
			phoneNumber1 = f.PhoneNumber()
		}
		var phoneNumber2 string
		if phoneNumber1 != "" && randBool(40) {
			phoneNumber2 = f.PhoneNumber()
		}
		var phoneNumber3 string
		if phoneNumber2 != "" && randBool(40) {
			phoneNumber3 = f.PhoneNumber()
		}
		preferredContactMethod := randomContactMethod()
		preferredContactMethodComments := ""
		preferredCommunicationLanguage := randomLanguage()
		prefersToRemainAnonymous := randBool(5)
		if prefersToRemainAnonymous {
			fullName = ""
			firstName = ""
			middleName = ""
			lastName = ""
			preferredName = ""
			nativeName = ""
		}
		presentsProtectionConcerns := randomBool()

		spokenLanguage1 := randomLanguage()
		spokenLanguage2 := randomLanguage()
		spokenLanguage3 := randomLanguage()

		serviceCC1 := string(enumTypes.ServiceCCShelter)
		serviceRequested1 := randomDate()
		serviceDelivered1 := randomDate()

		serviceCC2 := string(enumTypes.ServiceCCWash)
		serviceRequested2 := randomDate()
		serviceDelivered2 := randomDate()

		serviceCC3 := string(enumTypes.ServiceCCProtection)
		serviceRequested3 := randomDate()
		serviceDelivered3 := randomDate()

		serviceCC4 := string(enumTypes.ServiceCCEducation)
		serviceRequested4 := randomDate()
		serviceDelivered4 := randomDate()

		serviceCC5 := string(enumTypes.ServiceCCICLA)
		serviceRequested5 := randomDate()
		serviceDelivered5 := randomDate()

		serviceCC6 := string(enumTypes.ServiceCCLFS)
		serviceRequested6 := randomDate()
		serviceDelivered6 := randomDate()

		serviceCC7 := string(enumTypes.ServiceCCCVA)
		serviceRequested7 := randomDate()
		serviceDelivered7 := randomDate()

		inactive := randBool(10)

		if err := writer.Write([]string{
			address,
			age,
			birthDate,
			cognitiveDisabilityLevel,
			collectionAdministrativeArea1,
			collectionAdministrativeArea2,
			collectionAdministrativeArea3,
			collectionOffice,
			collectionAgentName,
			collectionAgentTitle,
			collectionTime,
			comments,
			communicationDisabilityLevel,
			communityId,
			communitySize,
			displacementStatus,
			displacementStatusComment,
			email1,
			email2,
			email3,
			fullName,
			firstName,
			middleName,
			lastName,
			nativeName,
			mothersName,
			freeField1,
			freeField2,
			freeField3,
			freeField4,
			freeField5,
			sex,
			hasCognitiveDisability,
			hasCommunicationDisability,
			hasConsentedToRgpd,
			hasConsentedToReferral,
			hasDisability,
			hasHearingDisability,
			hasMobilityDisability,
			hasSelfCareDisability,
			hasVisionDisability,
			hearingDisabilityLevel,
			householdId,
			householdSize,
			identificationType1,
			identificationTypeExplanation1,
			identificationNumber1,
			identificationType2,
			identificationTypeExplanation2,
			identificationNumber2,
			identificationType3,
			identificationTypeExplanation3,
			identificationNumber3,
			engagementContext,
			internalId,
			isHeadOfCommunity,
			isHeadOfHousehold,
			isFemaleHeadedHousehold,
			isMinorHeadedHousehold,
			isMinor,
			isChildAtRisk,
			isWomanAtRisk,
			isElderAtRisk,
			isSeparatedChild,
			isSingleParent,
			isPregnant,
			isLactating,
			hasMedicalCondition,
			needsLegalAndPhysicalProtection,
			vulnerabilityComments,
			mobilityDisabilityLevel,
			nationality1,
			nationality2,
			phoneNumber1,
			phoneNumber2,
			phoneNumber3,
			preferredContactMethod,
			preferredContactMethodComments,
			preferredName,
			preferredCommunicationLanguage,
			strconv.FormatBool(prefersToRemainAnonymous),
			presentsProtectionConcerns,
			pwdComments,
			selfCareDisabilityLevel,
			spokenLanguage1,
			spokenLanguage2,
			spokenLanguage3,
			visionDisabilityLevel,
			serviceCC1,
			serviceRequested1,
			serviceDelivered1,
			serviceCC2,
			serviceRequested2,
			serviceDelivered2,
			serviceCC3,
			serviceRequested3,
			serviceDelivered3,
			serviceCC4,
			serviceRequested4,
			serviceDelivered4,
			serviceCC5,
			serviceRequested5,
			serviceDelivered5,
			serviceCC6,
			serviceRequested6,
			serviceDelivered6,
			serviceCC7,
			serviceRequested7,
			serviceDelivered7,
			strconv.FormatBool(inactive),
		}); err != nil {
			return err
		}

	}

	writer.Flush()
	return nil

}

func randomText(f *faker.Faker) string {
	return strings.Join(f.Paragraphs(rand.Intn(3)+1, false), "\n\n")
}

func randBool(probability int) bool {
	return rand.Intn(100) < probability
}
