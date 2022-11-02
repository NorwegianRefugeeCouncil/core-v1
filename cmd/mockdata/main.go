package mockdata

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/manveru/faker"
	"github.com/nrc-no/notcore/internal/constants"
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
	switch rand.Intn(3) {
	case 0:
		return "national_id"
	case 1:
		return "passport"
	case 2:
		return "other"
	default:
		panic("unreachable")
	}
}

func randomDisabilityLevel() string {
	switch rand.Intn(3) {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	default:
		panic("unreachable")
	}
}

func randomGender() string {
	switch rand.Intn(3) {
	case 0:
		return "male"
	case 1:
		return "female"
	case 2:
		return "other"
	case 3:
		return "prefer_not_to_say"
	default:
		panic("unreachable")
	}
}

func randomDisplacementStatus() string {
	switch rand.Intn(2) {
	case 0:
		return "idp"
	case 1:
		return "host_community"
	case 2:
		return "refugee"
	default:
		panic("unreachable")
	}
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

func randomBool() string {
	if randBool(50) {
		return "true"
	}
	return "false"
}

func randomIdentificationContext() string {
	switch rand.Intn(4) {
	case 0:
		return "houseVisit"
	case 1:
		return "fieldActivity"
	case 2:
		return "inOffice"
	case 3:
		return "remoteChannels"
	case 4:
		return "other"
	default:
		panic("unreachable")
	}
}

func randomContactMethod() string {
	switch rand.Intn(5) {
	case 0:
		return "phone"
	case 1:
		return "email"
	case 2:
		return "sms"
	case 3:
		return "whatsapp"
	case 4:
		return "other"
	default:
		panic("unreachable")
	}
}

func Generate(count uint) error {

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
		constants.FileColumnIndividualBirthDate,
		constants.FileColumnIndividualCognitiveDisabilityLevel,
		constants.FileColumnIndividualCollectionAdministrativeArea1,
		constants.FileColumnIndividualCollectionAdministrativeArea2,
		constants.FileColumnIndividualCollectionAdministrativeArea3,
		constants.FileColumnIndividualCollectionAgentName,
		constants.FileColumnIndividualCollectionAgentTitle,
		constants.FileColumnIndividualCollectionTime,
		constants.FileColumnIndividualCommunicationDisabilityLevel,
		constants.FileColumnIndividualCommunityID,
		constants.FileColumnIndividualDisplacementStatus,
		constants.FileColumnIndividualEmail,
		constants.FileColumnIndividualFullName,
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
		constants.FileColumnIndividualIdentificationContext,
		constants.FileColumnIndividualInternalID,
		constants.FileColumnIndividualIsHeadOfCommunity,
		constants.FileColumnIndividualIsHeadOfHousehold,
		constants.FileColumnIndividualIsMinor,
		constants.FileColumnIndividualMobilityDisabilityLevel,
		constants.FileColumnIndividualNationality1,
		constants.FileColumnIndividualNationality2,
		constants.FileColumnIndividualPhoneNumber,
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

		cognitiveDisabilityLevel := randomDisabilityLevel()
		collectionAdministrativeArea1 := f.Country()
		collectionAdministrativeArea2 := f.State()
		collectionAdministrativeArea3 := f.City()
		collectionAgentName := f.Name()
		collectionAgentTitle := f.JobTitle()
		collectionTime := randomDate()
		communicationDisabilityLevel := randomDisabilityLevel()
		communityId := ""
		if randBool(20) {
			communityId = uuid.New().String()
		}
		displacementStatus := randomDisplacementStatus()
		var email string
		if randBool(80) {
			email = f.Email()
		}
		var fullName = f.Name()
		var preferredName = fullName
		if randBool(5) {
			preferredName = f.Name()
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
		householdId := ""
		if randBool(20) {
			householdId = uuid.New().String()
		}
		identificationType1 := randomIdentificationType()
		identificationTypeExplanation1 := ""
		identificationNumber1 := strconv.Itoa(rand.Intn(1000000000))
		identificationType2 := randomIdentificationType()
		identificationTypeExplanation2 := ""
		identificationNumber2 := strconv.Itoa(rand.Intn(1000000000))
		identificationType3 := randomIdentificationType()
		identificationTypeExplanation3 := ""
		identificationNumber3 := strconv.Itoa(rand.Intn(1000000000))
		identificationContext := randomIdentificationContext()
		internalId := strconv.Itoa(rand.Intn(1000000000))
		isHeadOfCommunity := randomBool()
		isHeadOfHousehold := randomBool()
		isMinor := randomBool()
		mobilityDisabilityLevel := randomDisabilityLevel()
		nationality1 := randomCountry()
		nationality2 := randomCountry()
		var phoneNumber string
		if randBool(80) {
			phoneNumber = f.PhoneNumber()
		}
		preferredContactMethod := randomContactMethod()
		preferredContactMethodComments := ""
		preferredCommunicationLanguage := randomLanguage()
		prefersToRemainAnonymous := randomBool()
		presentsProtectionConcerns := randomBool()
		selfCareDisabilityLevel := randomDisabilityLevel()
		spokenLanguage1 := randomLanguage()
		spokenLanguage2 := randomLanguage()
		spokenLanguage3 := randomLanguage()
		visionDisabilityLevel := randomDisabilityLevel()

		if err := writer.Write([]string{
			address,
			birthDate,
			cognitiveDisabilityLevel,
			collectionAdministrativeArea1,
			collectionAdministrativeArea2,
			collectionAdministrativeArea3,
			collectionAgentName,
			collectionAgentTitle,
			collectionTime,
			communicationDisabilityLevel,
			communityId,
			displacementStatus,
			email,
			fullName,
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
			identificationContext,
			internalId,
			isHeadOfCommunity,
			isHeadOfHousehold,
			isMinor,
			mobilityDisabilityLevel,
			nationality1,
			nationality2,
			phoneNumber,
			preferredContactMethod,
			preferredContactMethodComments,
			preferredName,
			preferredCommunicationLanguage,
			prefersToRemainAnonymous,
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

func randBool(probability int) bool {
	return rand.Intn(100) < probability
}
