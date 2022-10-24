package api

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/nrc-no/notcore/internal/constants"
)

func (i *Individual) UnmarshalFormData(v url.Values) error {
	var err error
	parseBool := func(s string) (bool, error) {
		switch s {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			return false, fmt.Errorf("invalid bool %q", s)
		}
	}
	i.IsMinor = v.Get(constants.FormParamIndividualIsMinor) == "true"
	i.PresentsProtectionConcerns = v.Get(constants.FormParamIndividualPresentsProtectionConcerns) == "true"
	i.CountryID = v.Get(constants.FormParamIndividualCountryID)
	i.Address = v.Get(constants.FormParamIndividualAddress)
	if i.BirthDate, err = ParseDate(v.Get(constants.FormParamIndividualBirthDate)); err != nil {
		return err
	}
	if i.CognitiveDisabilityLevel, err = ParseDisabilityLevel(v.Get(constants.FormParamIndividualCognitiveDisabilityLevel)); err != nil {
		return err
	}
	i.CollectionAdministrativeArea1 = v.Get(constants.FormParamIndividualCollectionAdministrativeArea1)
	i.CollectionAdministrativeArea2 = v.Get(constants.FormParamIndividualCollectionAdministrativeArea2)
	i.CollectionAdministrativeArea3 = v.Get(constants.FormParamIndividualCollectionAdministrativeArea3)
	i.CollectionAgentID = v.Get(constants.FormParamIndividualCollectionAgentID)
	if i.CollectionTime, err = ParseDate(v.Get(constants.FormParamIndividualCollectionTime)); err != nil {
		return err
	}
	if i.CommunicationDisabilityLevel, err = ParseDisabilityLevel(v.Get(constants.FormParamIndividualCommunicationDisabilityLevel)); err != nil {
		return err
	}
	i.CommunityID = v.Get(constants.FormParamIndividualCommunityID)
	i.CountryID = v.Get(constants.FormParamIndividualCountryID)
	i.DisplacementStatus = v.Get(constants.FormParamIndividualDisplacementStatus)
	i.Email = v.Get(constants.FormParamIndividualEmail)
	i.FullName = v.Get(constants.FormParamIndividualFullName)
	i.Gender = v.Get(constants.FormParamIndividualGender)
	if i.HasCognitiveDisability, err = parseBool(v.Get(constants.FormParamIndividualHasCognitiveDisability)); err != nil {
		return err
	}
	if i.HasCommunicationDisability, err = parseBool(v.Get(constants.FormParamIndividualHasCommunicationDisability)); err != nil {
		return err
	}
	if i.HasConsentedToRGPD, err = parseBool(v.Get(constants.FormParamIndividualHasConsentedToRGPD)); err != nil {
		return err
	}
	if i.HasConsentedToReferral, err = parseBool(v.Get(constants.FormParamIndividualHasConsentedToReferral)); err != nil {
		return err
	}
	if i.HasHearingDisability, err = parseBool(v.Get(constants.FormParamIndividualHasHearingDisability)); err != nil {
		return err
	}
	if i.HasMobilityDisability, err = parseBool(v.Get(constants.FormParamIndividualHasMobilityDisability)); err != nil {
		return err
	}
	if i.HasSelfCareDisability, err = parseBool(v.Get(constants.FormParamIndividualHasSelfCareDisability)); err != nil {
		return err
	}
	if i.HasVisionDisability, err = parseBool(v.Get(constants.FormParamIndividualHasVisionDisability)); err != nil {
		return err
	}
	if i.HearingDisabilityLevel, err = ParseDisabilityLevel(v.Get(constants.FormParamIndividualHearingDisabilityLevel)); err != nil {
		return err
	}
	i.HouseholdID = v.Get(constants.FormParamIndividualHouseholdID)
	i.ID = v.Get(constants.FormParamIndividualID)
	i.IdentificationType1 = v.Get(constants.FormParamIndividualIdentificationType1)
	i.IdentificationTypeExplanation1 = v.Get(constants.FormParamIndividualIdentificationTypeExplanation1)
	i.IdentificationNumber1 = v.Get(constants.FormParamIndividualIdentificationNumber1)
	i.IdentificationType2 = v.Get(constants.FormParamIndividualIdentificationType2)
	i.IdentificationTypeExplanation2 = v.Get(constants.FormParamIndividualIdentificationTypeExplanation2)
	i.IdentificationNumber2 = v.Get(constants.FormParamIndividualIdentificationNumber2)
	i.IdentificationType3 = v.Get(constants.FormParamIndividualIdentificationType3)
	i.IdentificationTypeExplanation3 = v.Get(constants.FormParamIndividualIdentificationTypeExplanation3)
	i.IdentificationNumber3 = v.Get(constants.FormParamIndividualIdentificationNumber3)
	i.IdentificationContext = v.Get(constants.FormParamIndividualIdentificationContext)
	i.InternalID = v.Get(constants.FormParamIndividualInternalID)
	if i.IsActive, err = parseBool(v.Get(constants.FormParamIndividualIsActive)); err != nil {
		return err
	}
	if i.IsHeadOfCommunity, err = parseBool(v.Get(constants.FormParamIndividualIsHeadOfCommunity)); err != nil {
		return err
	}
	if i.IsHeadOfHousehold, err = parseBool(v.Get(constants.FormParamIndividualIsHeadOfHousehold)); err != nil {
		return err
	}
	if i.IsMinor, err = parseBool(v.Get(constants.FormParamIndividualIsMinor)); err != nil {
		return err
	}
	if i.MobilityDisabilityLevel, err = ParseDisabilityLevel(v.Get(constants.FormParamIndividualMobilityDisabilityLevel)); err != nil {
		return err
	}
	i.Nationality1 = v.Get(constants.FormParamIndividualNationality1)
	i.Nationality2 = v.Get(constants.FormParamIndividualNationality2)
	i.PhoneNumber = v.Get(constants.FormParamIndividualPhoneNumber)
	i.PreferredContactMethod = v.Get(constants.FormParamIndividualPreferredContactMethod)
	i.PreferredContactMethodComments = v.Get(constants.FormParamIndividualPreferredContactMethodComments)
	i.PreferredName = v.Get(constants.FormParamIndividualPreferredName)
	i.PreferredCommunicationLanguage = v.Get(constants.FormParamIndividualPreferredCommunicationLanguage)
	if i.PrefersToRemainAnonymous, err = parseBool(v.Get(constants.FormParamIndividualPrefersToRemainAnonymous)); err != nil {
		return err
	}
	if i.PresentsProtectionConcerns, err = parseBool(v.Get(constants.FormParamIndividualPresentsProtectionConcerns)); err != nil {
		return err
	}
	if i.RecordedAge, err = strconv.Atoi(v.Get(constants.FormParamIndividualRecordedAge)); err != nil {
		return err
	}
	if i.SelfCareDisabilityLevel, err = ParseDisabilityLevel(v.Get(constants.FormParamIndividualSelfCareDisabilityLevel)); err != nil {
		return err
	}
	i.SpokenLanguage1 = v.Get(constants.FormParamIndividualSpokenLanguage1)
	i.SpokenLanguage2 = v.Get(constants.FormParamIndividualSpokenLanguage2)
	i.SpokenLanguage3 = v.Get(constants.FormParamIndividualSpokenLanguage3)
	if i.VisionDisabilityLevel, err = ParseDisabilityLevel(v.Get(constants.FormParamIndividualVisionDisabilityLevel)); err != nil {
		return err
	}
	i.Normalize()
	return nil
}
