package api

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
)

func newListIndividualsOptionsEncoder(values ListIndividualsOptions, now time.Time) *listIndividualsOptionsEncoder {
	return &listIndividualsOptionsEncoder{
		values: values,
		now:    now,
	}
}

type listIndividualsOptionsEncoder struct {
	out    url.Values
	values ListIndividualsOptions
	now    time.Time
}

func (p *listIndividualsOptionsEncoder) encode() url.Values {
	p.out = url.Values{}
	fns := []func(){
		p.encodeInactive,
		p.encodeAddress,
		p.encodeAgeFrom,
		p.encodeAgeTo,
		p.encodeBirthDateFrom,
		p.encodeBirthDateTo,
		p.encodeCognitiveDisabilityLevel,
		p.encodeCollectionAdministrativeArea1,
		p.encodeCollectionAdministrativeArea2,
		p.encodeCollectionAdministrativeArea3,
		p.encodeCollectionOffice,
		p.encodeCollectionAgentName,
		p.encodeCollectionAgentTitle,
		p.encodeCollectionTimeFrom,
		p.encodeCollectionTimeTo,
		p.encodeCommunityID,
		p.encodeCreatedAtFrom,
		p.encodeCreatedAtTo,
		p.encodeDisplacementStatuses,
		p.encodeEmail,
		p.encodeFreeField1,
		p.encodeFreeField2,
		p.encodeFreeField3,
		p.encodeFreeField4,
		p.encodeFreeField5,
		p.encodeFullName,
		p.encodeMothersName,
		p.encodeSexes,
		p.encodeHasCognitiveDisability,
		p.encodeHasCommunicationDisability,
		p.encodeHasConsentedToRGPD,
		p.encodeHasConsentedToReferral,
		p.encodeHasDisability,
		p.encodeHasHearingDisability,
		p.encodeHasMobilityDisability,
		p.encodeHasSelfCareDisability,
		p.encodeHasVisionDisability,
		p.encodeHearingDisabilityLevel,
		p.encodeHouseholdID,
		p.encodeID,
		p.encodeIdentificationNumber,
		p.encodeEngagementContext,
		p.encodeInternalID,
		p.encodeIsHeadOfCommunity,
		p.encodeIsHeadOfHousehold,
		p.encodeIsFemaleHeadedHousehold,
		p.encodeIsMinorHeadedHousehold,
		p.encodeIsMinor,
		p.encodeMobilityDisabilityLevel,
		p.encodeNationality,
		p.encodePhoneNumber,
		p.encodePreferredContactMethod,
		p.encodePreferredCommunicationLanguage,
		p.encodePrefersToRemainAnonymous,
		p.encodePresentsProtectionConcerns,
		p.encodeSelfCareDisabilityLevel,
		p.encodeServiceCC,
		p.encodeServiceRequestedDateFrom,
		p.encodeServiceRequestedDateTo,
		p.encodeServiceDeliveredDateFrom,
		p.encodeServiceDeliveredDateTo,
		p.encodeSpokenLanguage,
		p.encodeUpdatedAtFrom,
		p.encodeUpdatedAtTo,
		p.encodeSkip,
		p.encodeTake,
		p.encodeVisionDisabilityLevel,
		p.encodeSort,
	}
	for _, fn := range fns {
		fn()
	}
	return p.out
}

func (p *listIndividualsOptionsEncoder) encodeInactive() {
	if p.values.Inactive != nil {
		p.out.Add(constants.FormParamsGetIndividualsActive, strconv.FormatBool(*p.values.Inactive))
	}
}

func (p *listIndividualsOptionsEncoder) encodeAddress() {
	if len(p.values.Address) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsAddress, p.values.Address)
	}
}

func (p *listIndividualsOptionsEncoder) encodeAgeFrom() {
	if p.values.AgeFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsAgeFrom, strconv.Itoa(*p.values.AgeFrom))
	}
}

func (p *listIndividualsOptionsEncoder) encodeAgeTo() {
	if p.values.AgeTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsAgeTo, strconv.Itoa(*p.values.AgeTo))
	}
}

func (p *listIndividualsOptionsEncoder) encodeBirthDateFrom() {
	if p.values.BirthDateFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsBirthDateFrom, p.values.BirthDateFrom.Format("2006-01-02"))
	}
}

func (p *listIndividualsOptionsEncoder) encodeBirthDateTo() {
	if p.values.BirthDateTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsBirthDateTo, p.values.BirthDateTo.Format("2006-01-02"))
	}
}

func (p *listIndividualsOptionsEncoder) encodeServiceRequestedDateFrom() {
	if p.values.ServiceRequestedDateFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsServiceRequestedDateFrom, p.values.ServiceRequestedDateFrom.Format("2006-01-02"))
	}
}

func (p *listIndividualsOptionsEncoder) encodeServiceRequestedDateTo() {
	if p.values.ServiceRequestedDateTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsServiceRequestedDateTo, p.values.ServiceRequestedDateTo.Format("2006-01-02"))
	}
}

func (p *listIndividualsOptionsEncoder) encodeServiceDeliveredDateFrom() {
	if p.values.ServiceDeliveredDateFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsServiceDeliveredDateFrom, p.values.ServiceDeliveredDateFrom.Format("2006-01-02"))
	}
}

func (p *listIndividualsOptionsEncoder) encodeServiceDeliveredDateTo() {
	if p.values.ServiceDeliveredDateTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsServiceDeliveredDateTo, p.values.ServiceDeliveredDateTo.Format("2006-01-02"))
	}
}

func (p *listIndividualsOptionsEncoder) encodeCognitiveDisabilityLevel() {
	if p.values.CognitiveDisabilityLevel != DisabilityLevelUnspecified {
		p.out.Add(constants.FormParamsGetIndividualsCognitiveDisabilityLevel, string(p.values.CognitiveDisabilityLevel))
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionAdministrativeArea1() {
	if len(p.values.CollectionAdministrativeArea1) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsCollectionAdministrativeArea1, p.values.CollectionAdministrativeArea1)
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionAdministrativeArea2() {
	if len(p.values.CollectionAdministrativeArea2) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsCollectionAdministrativeArea2, p.values.CollectionAdministrativeArea2)
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionAdministrativeArea3() {
	if len(p.values.CollectionAdministrativeArea3) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsCollectionAdministrativeArea3, p.values.CollectionAdministrativeArea3)
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionOffice() {
	if len(p.values.CollectionOffice) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsCollectionOffice, p.values.CollectionOffice)
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionAgentName() {
	if len(p.values.CollectionAgentName) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsCollectionAgentName, p.values.CollectionAgentName)
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionAgentTitle() {
	if len(p.values.CollectionAgentTitle) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsCollectionAgentTitle, p.values.CollectionAgentTitle)
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionTimeFrom() {
	if p.values.CollectionTimeFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsCollectionTimeFrom, p.values.CollectionTimeFrom.Format(time.RFC3339))
	}
}

func (p *listIndividualsOptionsEncoder) encodeCollectionTimeTo() {
	if p.values.CollectionTimeTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsCollectionTimeTo, p.values.CollectionTimeTo.Format(time.RFC3339))
	}
}

func (p *listIndividualsOptionsEncoder) encodeCommunityID() {
	if len(p.values.CommunityID) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsCommunityID, p.values.CommunityID)
	}
}

func (p *listIndividualsOptionsEncoder) encodeCreatedAtFrom() {
	if p.values.CreatedAtFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsCreatedAtFrom, p.values.CreatedAtFrom.Format(time.RFC3339))
	}
}

func (p *listIndividualsOptionsEncoder) encodeCreatedAtTo() {
	if p.values.CreatedAtTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsCreatedAtTo, p.values.CreatedAtTo.Format(time.RFC3339))
	}
}

func (p *listIndividualsOptionsEncoder) encodeDisplacementStatuses() {
	if len(p.values.DisplacementStatuses) > 0 {
		for _, ds := range p.values.DisplacementStatuses.Items() {
			p.out.Add(constants.FormParamsGetIndividualsDisplacementStatus, string(ds))
		}
	}
}

func (p *listIndividualsOptionsEncoder) encodeEmail() {
	if len(p.values.Email) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsEmail, p.values.Email)
	}
}

func (p *listIndividualsOptionsEncoder) encodeFreeField1() {
	if len(p.values.FreeField1) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsFreeField1, p.values.FreeField1)
	}
}

func (p *listIndividualsOptionsEncoder) encodeFreeField2() {
	if len(p.values.FreeField2) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsFreeField2, p.values.FreeField2)
	}
}

func (p *listIndividualsOptionsEncoder) encodeFreeField3() {
	if len(p.values.FreeField3) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsFreeField3, p.values.FreeField3)
	}
}

func (p *listIndividualsOptionsEncoder) encodeFreeField4() {
	if len(p.values.FreeField4) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsFreeField4, p.values.FreeField4)
	}
}

func (p *listIndividualsOptionsEncoder) encodeFreeField5() {
	if len(p.values.FreeField5) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsFreeField5, p.values.FreeField5)
	}
}

func (p *listIndividualsOptionsEncoder) encodeFullName() {
	if len(p.values.FullName) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsFullName, p.values.FullName)
	}
}

func (p *listIndividualsOptionsEncoder) encodeMothersName() {
	if len(p.values.MothersName) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsMothersName, p.values.MothersName)
	}
}

func (p *listIndividualsOptionsEncoder) encodeSexes() {
	if len(p.values.Sexes) > 0 {
		for _, g := range p.values.Sexes.Items() {
			p.out.Add(constants.FormParamsGetIndividualsSex, string(g))
		}
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasCognitiveDisability() {
	if p.values.HasCognitiveDisability != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasCognitiveDisability, strconv.FormatBool(*p.values.HasCognitiveDisability))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasCommunicationDisability() {
	if p.values.HasCommunicationDisability != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasCommunicationDisability, strconv.FormatBool(*p.values.HasCommunicationDisability))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasConsentedToRGPD() {
	if p.values.HasConsentedToRGPD != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasConsentedToRgpd, strconv.FormatBool(*p.values.HasConsentedToRGPD))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasConsentedToReferral() {
	if p.values.HasConsentedToReferral != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasConsentedToReferral, strconv.FormatBool(*p.values.HasConsentedToReferral))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasDisability() {
	if p.values.HasDisability != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasDisability, strconv.FormatBool(*p.values.HasDisability))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasHearingDisability() {
	if p.values.HasHearingDisability != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasHearingDisability, strconv.FormatBool(*p.values.HasHearingDisability))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasMobilityDisability() {
	if p.values.HasMobilityDisability != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasMobilityDisability, strconv.FormatBool(*p.values.HasMobilityDisability))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasSelfCareDisability() {
	if p.values.HasSelfCareDisability != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasSelfCareDisability, strconv.FormatBool(*p.values.HasSelfCareDisability))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHasVisionDisability() {
	if p.values.HasVisionDisability != nil {
		p.out.Add(constants.FormParamsGetIndividualsHasVisionDisability, strconv.FormatBool(*p.values.HasVisionDisability))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHearingDisabilityLevel() {
	if p.values.HearingDisabilityLevel != DisabilityLevelUnspecified {
		p.out.Add(constants.FormParamsGetIndividualsHearingDisabilityLevel, string(p.values.HearingDisabilityLevel))
	}
}

func (p *listIndividualsOptionsEncoder) encodeHouseholdID() {
	if len(p.values.HouseholdID) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsHouseholdID, p.values.HouseholdID)
	}
}

func (p *listIndividualsOptionsEncoder) encodeID() {
	if p.values.IDs.Len() != 0 {
		for _, id := range p.values.IDs.Items() {
			p.out.Add(constants.FormParamsGetIndividualsID, id)
		}
	}
}

func (p *listIndividualsOptionsEncoder) encodeIdentificationNumber() {
	if len(p.values.IdentificationNumber) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsIdentificationNumber, p.values.IdentificationNumber)
	}
}

func (p *listIndividualsOptionsEncoder) encodeEngagementContext() {
	if len(p.values.EngagementContext) > 0 {
		for _, ds := range p.values.EngagementContext.Items() {
			p.out.Add(constants.FormParamsGetIndividualsEngagementContext, string(ds))
		}
	}
}

func (p *listIndividualsOptionsEncoder) encodeInternalID() {
	if len(p.values.InternalID) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsInternalID, p.values.InternalID)
	}
}

func (p *listIndividualsOptionsEncoder) encodeIsHeadOfCommunity() {
	if p.values.IsHeadOfCommunity != nil {
		p.out.Add(constants.FormParamsGetIndividualsIsHeadOfCommunity, strconv.FormatBool(*p.values.IsHeadOfCommunity))
	}
}

func (p *listIndividualsOptionsEncoder) encodeIsHeadOfHousehold() {
	if p.values.IsHeadOfHousehold != nil {
		p.out.Add(constants.FormParamsGetIndividualsIsHeadOfHousehold, strconv.FormatBool(*p.values.IsHeadOfHousehold))
	}
}

func (p *listIndividualsOptionsEncoder) encodeIsFemaleHeadedHousehold() {
	if p.values.IsFemaleHeadedHousehold != nil {
		p.out.Add(constants.FormParamsGetIndividualsIsFemaleHeadedHousehold, strconv.FormatBool(*p.values.IsFemaleHeadedHousehold))
	}
}

func (p *listIndividualsOptionsEncoder) encodeIsMinorHeadedHousehold() {
	if p.values.IsMinorHeadedHousehold != nil {
		p.out.Add(constants.FormParamsGetIndividualsIsMinorHeadedHousehold, strconv.FormatBool(*p.values.IsMinorHeadedHousehold))
	}
}

func (p *listIndividualsOptionsEncoder) encodeIsMinor() {
	if p.values.IsMinor != nil {
		p.out.Add(constants.FormParamsGetIndividualsIsMinor, strconv.FormatBool(*p.values.IsMinor))
	}
}

func (p *listIndividualsOptionsEncoder) encodeMobilityDisabilityLevel() {
	if p.values.MobilityDisabilityLevel != DisabilityLevelUnspecified {
		p.out.Add(constants.FormParamsGetIndividualsMobilityDisabilityLevel, string(p.values.MobilityDisabilityLevel))
	}
}

func (p *listIndividualsOptionsEncoder) encodeServiceCC() {
	if len(p.values.ServiceCC) > 0 {
		for _, ds := range p.values.ServiceCC.Items() {
			p.out.Add(constants.FormParamsGetIndividualsServiceCC, string(ds))
		}
	}
}

func (p *listIndividualsOptionsEncoder) encodeNationality() {
	if len(p.values.Nationality) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsNationality, p.values.Nationality)
	}
}

func (p *listIndividualsOptionsEncoder) encodePhoneNumber() {
	if len(p.values.PhoneNumber) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsPhoneNumber, p.values.PhoneNumber)
	}
}

func (p *listIndividualsOptionsEncoder) encodePreferredContactMethod() {
	if len(p.values.PreferredContactMethod) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsPreferredContactMethod, p.values.PreferredContactMethod)
	}
}

func (p *listIndividualsOptionsEncoder) encodePreferredCommunicationLanguage() {
	if len(p.values.PreferredCommunicationLanguage) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsPreferredCommunicationLanguage, p.values.PreferredCommunicationLanguage)
	}
}

func (p *listIndividualsOptionsEncoder) encodePrefersToRemainAnonymous() {
	if p.values.PrefersToRemainAnonymous != nil {
		p.out.Add(constants.FormParamsGetIndividualsPrefersToRemainAnonymous, strconv.FormatBool(*p.values.PrefersToRemainAnonymous))
	}
}

func (p *listIndividualsOptionsEncoder) encodePresentsProtectionConcerns() {
	if p.values.PresentsProtectionConcerns != nil {
		p.out.Add(constants.FormParamsGetIndividualsPresentsProtectionConcerns, strconv.FormatBool(*p.values.PresentsProtectionConcerns))
	}
}

func (p *listIndividualsOptionsEncoder) encodeSelfCareDisabilityLevel() {
	if p.values.SelfCareDisabilityLevel != DisabilityLevelUnspecified {
		p.out.Add(constants.FormParamsGetIndividualsSelfCareDisabilityLevel, string(p.values.SelfCareDisabilityLevel))
	}
}

func (p *listIndividualsOptionsEncoder) encodeSpokenLanguage() {
	if len(p.values.SpokenLanguage) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsSpokenLanguage, p.values.SpokenLanguage)
	}
}

func (p *listIndividualsOptionsEncoder) encodeUpdatedAtFrom() {
	if p.values.UpdatedAtFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsUpdatedAtFrom, p.values.UpdatedAtFrom.Format(time.RFC3339))
	}
}

func (p *listIndividualsOptionsEncoder) encodeUpdatedAtTo() {
	if p.values.UpdatedAtTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsUpdatedAtTo, p.values.UpdatedAtTo.Format(time.RFC3339))
	}
}

func (p *listIndividualsOptionsEncoder) encodeVisionDisabilityLevel() {
	if p.values.VisionDisabilityLevel != DisabilityLevelUnspecified {
		p.out.Add(constants.FormParamsGetIndividualsVisionDisabilityLevel, string(p.values.VisionDisabilityLevel))
	}
}

func (p *listIndividualsOptionsEncoder) encodeSkip() {
	if p.values.Skip != 0 {
		p.out.Add(constants.FormParamsGetIndividualsSkip, fmt.Sprintf("%d", p.values.Skip))
	}
}

func (p *listIndividualsOptionsEncoder) encodeSort() {
	if len(p.values.Sort) == 0 {
		return
	}
	p.out.Add(constants.FormParamsGetIndividualsSort, p.values.Sort.MarshalQuery())
}

func (p *listIndividualsOptionsEncoder) encodeTake() {
	if p.values.Take != 0 {
		p.out.Add(constants.FormParamsGetIndividualsTake, fmt.Sprintf("%d", p.values.Take))
	}
}

var sortableColumns = containers.NewStringSet(
	constants.DBColumnIndividualAddress,
	constants.DBColumnIndividualAge,
	constants.DBColumnIndividualBirthDate,
	constants.DBColumnIndividualCognitiveDisabilityLevel,
	constants.DBColumnIndividualCollectionAdministrativeArea1,
	constants.DBColumnIndividualCollectionAdministrativeArea2,
	constants.DBColumnIndividualCollectionAdministrativeArea3,
	constants.DBColumnIndividualCollectionAgentName,
	constants.DBColumnIndividualCollectionAgentTitle,
	constants.DBColumnIndividualCollectionOffice,
	constants.DBColumnIndividualCollectionTime,
	constants.DBColumnIndividualCommunicationDisabilityLevel,
	constants.DBColumnIndividualCommunityID,
	constants.DBColumnIndividualCreatedAt,
	constants.DBColumnIndividualDisplacementStatus,
	constants.DBColumnIndividualEmail1,
	constants.DBColumnIndividualEmail2,
	constants.DBColumnIndividualEmail3,
	constants.DBColumnIndividualEngagementContext,
	constants.DBColumnIndividualFirstName,
	constants.DBColumnIndividualFreeField1,
	constants.DBColumnIndividualFreeField2,
	constants.DBColumnIndividualFreeField3,
	constants.DBColumnIndividualFreeField4,
	constants.DBColumnIndividualFreeField5,
	constants.DBColumnIndividualFullName,
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
	constants.DBColumnIndividualID,
	constants.DBColumnIndividualIdentificationNumber1,
	constants.DBColumnIndividualIdentificationNumber2,
	constants.DBColumnIndividualIdentificationNumber3,
	constants.DBColumnIndividualIdentificationType1,
	constants.DBColumnIndividualIdentificationType2,
	constants.DBColumnIndividualIdentificationType3,
	constants.DBColumnIndividualInactive,
	constants.DBColumnIndividualInternalID,
	constants.DBColumnIndividualIsFemaleHeadedHousehold,
	constants.DBColumnIndividualIsHeadOfCommunity,
	constants.DBColumnIndividualIsHeadOfHousehold,
	constants.DBColumnIndividualIsMinor,
	constants.DBColumnIndividualIsMinorHeadedHousehold,
	constants.DBColumnIndividualLastName,
	constants.DBColumnIndividualMiddleName,
	constants.DBColumnIndividualMobilityDisabilityLevel,
	constants.DBColumnIndividualMothersName,
	constants.DBColumnIndividualNationality1,
	constants.DBColumnIndividualNationality2,
	constants.DBColumnIndividualNativeName,
	constants.DBColumnIndividualPhoneNumber1,
	constants.DBColumnIndividualPhoneNumber2,
	constants.DBColumnIndividualPhoneNumber3,
	constants.DBColumnIndividualPreferredCommunicationLanguage,
	constants.DBColumnIndividualPreferredContactMethod,
	constants.DBColumnIndividualPrefersToRemainAnonymous,
	constants.DBColumnIndividualPresentsProtectionConcerns,
	constants.DBColumnIndividualSelfCareDisabilityLevel,
	constants.DBColumnIndividualSex,
	constants.DBColumnIndividualSpokenLanguage1,
	constants.DBColumnIndividualSpokenLanguage2,
	constants.DBColumnIndividualSpokenLanguage3,
	constants.DBColumnIndividualUpdatedAt,
	constants.DBColumnIndividualVisionDisabilityLevel,
)
