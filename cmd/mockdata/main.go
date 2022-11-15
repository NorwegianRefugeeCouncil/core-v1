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

func randomGender() string {
	var g []string
	for _, d := range api.AllGenders().Items() {
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
		constants.FileColumnIndividualCollectionAgentName,
		constants.FileColumnIndividualCollectionAgentTitle,
		constants.FileColumnIndividualCollectionTime,
		constants.FileColumnIndividualComments,
		constants.FileColumnIndividualCommunicationDisabilityLevel,
		constants.FileColumnIndividualCommunityID,
		constants.FileColumnIndividualDisplacementStatus,
		constants.FileColumnIndividualEmail1,
		constants.FileColumnIndividualEmail2,
		constants.FileColumnIndividualEmail3,
		constants.FileColumnIndividualFullName,
		constants.FileColumnIndividualFreeField1,
		constants.FileColumnIndividualFreeField2,
		constants.FileColumnIndividualFreeField3,
		constants.FileColumnIndividualFreeField4,
		constants.FileColumnIndividualFreeField5,
		constants.FileColumnIndividualGender,
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
		displacementStatus := randomDisplacementStatus()

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

		var fullName = strings.ToUpper(f.LastName()) + ", " + f.FirstName()
		var preferredName = fullName
		if randBool(5) {
			preferredName = f.Name()
		}

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

		gender := randomGender()
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
			preferredName = ""
		}
		presentsProtectionConcerns := randomBool()

		spokenLanguage1 := randomLanguage()
		spokenLanguage2 := randomLanguage()
		spokenLanguage3 := randomLanguage()

		if err := writer.Write([]string{
			address,
			age,
			birthDate,
			cognitiveDisabilityLevel,
			collectionAdministrativeArea1,
			collectionAdministrativeArea2,
			collectionAdministrativeArea3,
			collectionAgentName,
			collectionAgentTitle,
			collectionTime,
			comments,
			communicationDisabilityLevel,
			communityId,
			displacementStatus,
			email1,
			email2,
			email3,
			fullName,
			freeField1,
			freeField2,
			freeField3,
			freeField4,
			freeField5,
			gender,
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
