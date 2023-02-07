package mockdata

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/manveru/faker"
	"github.com/nrc-no/notcore/internal/api"
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
		"nationalId",
		"other",
	)
}

func randomDisabilityLevel() string {
	var g []string
	for _, d := range api.AllDisabilityLevels().Items() {
		g = append(g, string(d))
	}
	return pick(g...)
}

func randomSex() string {
	var g []string
	for _, d := range api.AllSexes().Items() {
		g = append(g, string(d))
	}
	return pick(g...)
}

func randomDisplacementStatus() string {
	var ds []string
	for _, d := range api.AllDisplacementStatuses().Items() {
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
		"sms",
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
		constants.FileColumnIndividualAddress,
		constants.FileColumnIndividualAge,
		constants.FileColumnIndividualBirthDate,
		constants.FileColumnIndividualCognitiveDisabilityLevel,
		constants.FileColumnIndividualCollectionAdministrativeArea1,
		constants.FileColumnIndividualCollectionAdministrativeArea2,
		constants.FileColumnIndividualCollectionAdministrativeArea3,
		constants.FileColumnIndividualCollectionOffice,
		constants.FileColumnIndividualCollectionAgentName,
		constants.FileColumnIndividualCollectionAgentTitle,
		constants.FileColumnIndividualCollectionTime,
		constants.FileColumnIndividualComments,
		constants.FileColumnIndividualCommunicationDisabilityLevel,
		constants.FileColumnIndividualCommunityID,
		constants.FileColumnIndividualCommunitySize,
		constants.FileColumnIndividualDisplacementStatus,
		constants.FileColumnIndividualDisplacementStatusComment,
		constants.FileColumnIndividualEmail1,
		constants.FileColumnIndividualEmail2,
		constants.FileColumnIndividualEmail3,
		constants.FileColumnIndividualFullName,
		constants.FileColumnIndividualFirstName,
		constants.FileColumnIndividualMiddleName,
		constants.FileColumnIndividualLastName,
		constants.FileColumnIndividualMothersName,
		constants.FileColumnIndividualFreeField1,
		constants.FileColumnIndividualFreeField2,
		constants.FileColumnIndividualFreeField3,
		constants.FileColumnIndividualFreeField4,
		constants.FileColumnIndividualFreeField5,
		constants.FileColumnIndividualSex,
		constants.FileColumnIndividualHasCognitiveDisability,
		constants.FileColumnIndividualHasCommunicationDisability,
		constants.FileColumnIndividualHasConsentedToRGPD,
		constants.FileColumnIndividualHasConsentedToReferral,
		constants.FileColumnIndividualHasHearingDisability,
		constants.FileColumnIndividualHasMobilityDisability,
		constants.FileColumnIndividualHasSelfCareDisability,
		constants.FileColumnIndividualHasVisionDisability,
		constants.FileColumnIndividualHearingDisabilityLevel,
		constants.FileColumnIndividualHouseholdID,
		constants.FileColumnIndividualHouseholdSize,
		constants.FileColumnIndividualIdentificationType1,
		constants.FileColumnIndividualIdentificationTypeExplanation1,
		constants.FileColumnIndividualIdentificationNumber1,
		constants.FileColumnIndividualIdentificationType2,
		constants.FileColumnIndividualIdentificationTypeExplanation2,
		constants.FileColumnIndividualIdentificationNumber2,
		constants.FileColumnIndividualIdentificationType3,
		constants.FileColumnIndividualIdentificationTypeExplanation3,
		constants.FileColumnIndividualIdentificationNumber3,
		constants.FileColumnIndividualEngagementContext,
		constants.FileColumnIndividualInternalID,
		constants.FileColumnIndividualIsHeadOfCommunity,
		constants.FileColumnIndividualIsHeadOfHousehold,
		constants.FileColumnIndividualIsFemaleHeadedHousehold,
		constants.FileColumnIndividualIsMinorHeadedHousehold,
		constants.FileColumnIndividualIsMinor,
		constants.FileColumnIndividualMobilityDisabilityLevel,
		constants.FileColumnIndividualNationality1,
		constants.FileColumnIndividualNationality2,
		constants.FileColumnIndividualPhoneNumber1,
		constants.FileColumnIndividualPhoneNumber2,
		constants.FileColumnIndividualPhoneNumber3,
		constants.FileColumnIndividualPreferredContactMethod,
		constants.FileColumnIndividualPreferredContactMethodComments,
		constants.FileColumnIndividualPreferredName,
		constants.FileColumnIndividualPreferredCommunicationLanguage,
		constants.FileColumnIndividualPrefersToRemainAnonymous,
		constants.FileColumnIndividualPresentsProtectionConcerns,
		constants.FileColumnIndividualSelfCareDisabilityLevel,
		constants.FileColumnIndividualSpokenLanguage1,
		constants.FileColumnIndividualSpokenLanguage2,
		constants.FileColumnIndividualSpokenLanguage3,
		constants.FileColumnIndividualVisionDisabilityLevel,
		constants.FileColumnIndividualServiceCC1,
		constants.FileColumnIndividualServiceRequestedDate1,
		constants.FileColumnIndividualServiceDeliveredDate1,
		constants.FileColumnIndividualServiceCC2,
		constants.FileColumnIndividualServiceRequestedDate2,
		constants.FileColumnIndividualServiceDeliveredDate2,
		constants.FileColumnIndividualServiceCC3,
		constants.FileColumnIndividualServiceRequestedDate3,
		constants.FileColumnIndividualServiceDeliveredDate3,
		constants.FileColumnIndividualServiceCC4,
		constants.FileColumnIndividualServiceRequestedDate4,
		constants.FileColumnIndividualServiceDeliveredDate4,
		constants.FileColumnIndividualServiceCC5,
		constants.FileColumnIndividualServiceRequestedDate5,
		constants.FileColumnIndividualServiceDeliveredDate5,
		constants.FileColumnIndividualServiceCC6,
		constants.FileColumnIndividualServiceRequestedDate6,
		constants.FileColumnIndividualServiceDeliveredDate6,
		constants.FileColumnIndividualServiceCC7,
		constants.FileColumnIndividualServiceRequestedDate7,
		constants.FileColumnIndividualServiceDeliveredDate7,
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
		hasCognitiveDisability := randomBool()
		hasCommunicationDisability := randomBool()
		hasConsentedToRgpd := randomBool()
		hasConsentedToReferral := randomBool()
		hasHearingDisability := randomBool()
		hasMobilityDisability := randomBool()
		hasSelfCareDisability := randomBool()
		hasVisionDisability := randomBool()

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
			isHeadOfHousehold = randomBool()
		}

		isMinorHeadedHousehold := randomBool()
		isFemaleHeadedHousehold := randomBool()

		isMinor := randomBool()

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
		}
		presentsProtectionConcerns := randomBool()

		spokenLanguage1 := randomLanguage()
		spokenLanguage2 := randomLanguage()
		spokenLanguage3 := randomLanguage()

		serviceCC1 := string(api.ServiceCCShelter)
		serviceRequested1 := randomDate()
		serviceDelivered1 := randomDate()

		serviceCC2 := string(api.ServiceCCWash)
		serviceRequested2 := randomDate()
		serviceDelivered2 := randomDate()

		serviceCC3 := string(api.ServiceCCProtection)
		serviceRequested3 := randomDate()
		serviceDelivered3 := randomDate()

		serviceCC4 := string(api.ServiceCCEducation)
		serviceRequested4 := randomDate()
		serviceDelivered4 := randomDate()

		serviceCC5 := string(api.ServiceCCICLA)
		serviceRequested5 := randomDate()
		serviceDelivered5 := randomDate()

		serviceCC6 := string(api.ServiceCCLFS)
		serviceRequested6 := randomDate()
		serviceDelivered6 := randomDate()

		serviceCC7 := string(api.ServiceCCCVA)
		serviceRequested7 := randomDate()
		serviceDelivered7 := randomDate()

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
