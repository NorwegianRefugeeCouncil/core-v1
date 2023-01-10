package api

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
)

type listIndividualsOptionsDecoder struct {
	out    *ListIndividualsOptions
	values url.Values
	now    time.Time
}

func (p *listIndividualsOptionsDecoder) parse() error {
	fns := []func() error{
		p.parseActive,
		p.parseAddress,
		p.parseAgeFrom,
		p.parseAgeTo,
		p.parseBirthDateFrom,
		p.parseBirthDateTo,
		p.parseCognitiveDisabilityLevel,
		p.parseCollectionAdministrativeArea1,
		p.parseCollectionAdministrativeArea2,
		p.parseCollectionAdministrativeArea3,
		p.parseCollectionOffice,
		p.parseCollectionAgentName,
		p.parseCollectionAgentTitle,
		p.parseCollectionTimeFrom,
		p.parseCollectionTimeTo,
		p.parseCommunityID,
		p.parseCountryID,
		p.parseCreatedAtFrom,
		p.parseCreatedAtTo,
		p.parseDisplacementStatuses,
		p.parseEmail,
		p.parseFreeField1,
		p.parseFreeField2,
		p.parseFreeField3,
		p.parseFreeField4,
		p.parseFreeField5,
		p.parseFullName,
		p.parseSexes,
		p.parseHasCognitiveDisability,
		p.parseHasCommunicationDisability,
		p.parseHasConsentedToRgpd,
		p.parseHasConsentedToReferral,
		p.parseHasHearingDisability,
		p.parseHasMobilityDisability,
		p.parseHasSelfCareDisability,
		p.parseHasVisionDisability,
		p.parseHearingDisabilityLevel,
		p.parseHouseholdID,
		p.parseIDs,
		p.parseIdentificationNumber,
		p.parseEngagementContext,
		p.parseInternalID,
		p.parseIsHeadOfCommunity,
		p.parseIsHeadOfHousehold,
		p.parseIsFemaleHeadedHousehold,
		p.parseIsMinorHeadedHousehold,
		p.parseIsMinor,
		p.parseMobilityDisabilityLevel,
		p.parseNationality,
		p.parsePhoneNumber,
		p.parsePreferredContactMethod,
		p.parsePreferredCommunicationLanguage,
		p.parsePrefersToRemainAnonymous,
		p.parsePresentsProtectionConcerns,
		p.parseSelfCareDisabilityLevel,
		p.parseSpokenLanguage,
		p.parseUpdatedAtFrom,
		p.parseUpdatedAtTo,
		p.parseSkip,
		p.parseTake,
		p.parseVisionDisabilityLevel,
		p.parseSort,
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func (p *listIndividualsOptionsDecoder) parseActive() error {
	var err error
	p.out.Active, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsActive))
	return err
}

func (p *listIndividualsOptionsDecoder) parseAddress() error {
	p.out.Address = p.values.Get(constants.FormParamsGetIndividualsAddress)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseAgeFrom() error {
	ageFrom, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeFrom))
	if err != nil || ageFrom == nil {
		return err
	}
	p.out.AgeFrom = ageFrom
	return nil
}

func (p *listIndividualsOptionsDecoder) parseAgeTo() error {
	ageTo, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeTo))
	if err != nil || ageTo == nil {
		return err
	}
	p.out.AgeTo = ageTo
	return nil
}

func (p *listIndividualsOptionsDecoder) parseBirthDateFrom() error {
	birthDateFrom, err := parseOptionalDate(p.values.Get(constants.FormParamsGetIndividualsBirthDateFrom))
	if err != nil || birthDateFrom == nil {
		return err
	}
	p.out.BirthDateFrom = birthDateFrom
	return nil
}

func (p *listIndividualsOptionsDecoder) parseBirthDateTo() error {
	birthDateTo, err := parseOptionalDate(p.values.Get(constants.FormParamsGetIndividualsBirthDateTo))
	if err != nil || birthDateTo == nil {
		return err
	}
	p.out.BirthDateTo = birthDateTo
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCognitiveDisabilityLevel() error {
	var err error
	p.out.CognitiveDisabilityLevel, err = ParseDisabilityLevel(p.values.Get(constants.FormParamsGetIndividualsCognitiveDisabilityLevel))
	return err
}

func (p *listIndividualsOptionsDecoder) parseCollectionAdministrativeArea1() error {
	p.out.CollectionAdministrativeArea1 = p.values.Get(constants.FormParamsGetIndividualsCollectionAdministrativeArea1)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCollectionAdministrativeArea2() error {
	p.out.CollectionAdministrativeArea2 = p.values.Get(constants.FormParamsGetIndividualsCollectionAdministrativeArea2)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCollectionAdministrativeArea3() error {
	p.out.CollectionAdministrativeArea3 = p.values.Get(constants.FormParamsGetIndividualsCollectionAdministrativeArea3)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCollectionOffice() error {
	p.out.CollectionOffice = p.values.Get(constants.FormParamsGetIndividualsCollectionOffice)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCollectionAgentName() error {
	p.out.CollectionAgentName = p.values.Get(constants.FormParamsGetIndividualsCollectionAgentName)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCollectionAgentTitle() error {
	p.out.CollectionAgentTitle = p.values.Get(constants.FormParamsGetIndividualsCollectionAgentTitle)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCollectionTimeFrom() error {
	var err error
	p.out.CollectionTimeFrom, err = parseOptionalDateTime(p.values.Get(constants.FormParamsGetIndividualsCollectionTimeFrom))
	return err
}

func (p *listIndividualsOptionsDecoder) parseCollectionTimeTo() error {
	var err error
	p.out.CollectionTimeTo, err = parseOptionalDateTime(p.values.Get(constants.FormParamsGetIndividualsCollectionTimeTo))
	return err
}

func (p *listIndividualsOptionsDecoder) parseCommunityID() error {
	p.out.CommunityID = p.values.Get(constants.FormParamsGetIndividualsCommunityID)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCountryID() error {
	p.out.CountryID = p.values.Get(constants.FormParamsGetIndividualsCountryID)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCreatedAtFrom() error {
	var err error
	p.out.CreatedAtFrom, err = parseOptionalDateTime(p.values.Get(constants.FormParamsGetIndividualsCreatedAtFrom))
	return err
}

func (p *listIndividualsOptionsDecoder) parseCreatedAtTo() error {
	var err error
	p.out.CreatedAtTo, err = parseOptionalDateTime(p.values.Get(constants.FormParamsGetIndividualsCreatedAtTo))
	return err
}

func (p *listIndividualsOptionsDecoder) parseDisplacementStatuses() error {
	if len(p.values[constants.FormParamsGetIndividualsDisplacementStatus]) == 0 {
		return nil
	}
	dsSet := containers.NewSet[DisplacementStatus]()
	for _, ds := range p.values[constants.FormParamsGetIndividualsDisplacementStatus] {
		parsedDs, err := ParseDisplacementStatus(ds)
		if err != nil {
			return err
		}
		dsSet.Add(parsedDs)
	}
	p.out.DisplacementStatuses = dsSet
	return nil
}

func (p *listIndividualsOptionsDecoder) parseEmail() error {
	p.out.Email = p.values.Get(constants.FormParamsGetIndividualsEmail)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseFreeField1() error {
	p.out.FreeField1 = p.values.Get(constants.FormParamsGetIndividualsFreeField1)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseFreeField2() error {
	p.out.FreeField2 = p.values.Get(constants.FormParamsGetIndividualsFreeField2)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseFreeField3() error {
	p.out.FreeField3 = p.values.Get(constants.FormParamsGetIndividualsFreeField3)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseFreeField4() error {
	p.out.FreeField4 = p.values.Get(constants.FormParamsGetIndividualsFreeField4)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseFreeField5() error {
	p.out.FreeField5 = p.values.Get(constants.FormParamsGetIndividualsFreeField5)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseFullName() error {
	p.out.FullName = p.values.Get(constants.FormParamsGetIndividualsFullName)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseSexes() error {
	if len(p.values[constants.FormParamsGetIndividualsSex]) == 0 {
		return nil
	}
	sexSet := containers.NewSet[Sex]()
	for _, g := range p.values[constants.FormParamsGetIndividualsSex] {
		sex, err := ParseSex(g)
		if err != nil {
			return err
		}
		sexSet.Add(sex)
	}
	p.out.Sexes = sexSet
	return nil
}

func (p *listIndividualsOptionsDecoder) parseHasCognitiveDisability() error {
	var err error
	p.out.HasCognitiveDisability, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasCognitiveDisability))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHasCommunicationDisability() error {
	var err error
	p.out.HasCommunicationDisability, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasCommunicationDisability))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHasConsentedToRgpd() error {
	var err error
	p.out.HasConsentedToRGPD, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasConsentedToRgpd))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHasConsentedToReferral() error {
	var err error
	p.out.HasConsentedToReferral, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasConsentedToReferral))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHasHearingDisability() error {
	var err error
	p.out.HasHearingDisability, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasHearingDisability))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHasMobilityDisability() error {
	var err error
	p.out.HasMobilityDisability, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasMobilityDisability))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHasSelfCareDisability() error {
	var err error
	p.out.HasSelfCareDisability, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasSelfCareDisability))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHasVisionDisability() error {
	var err error
	p.out.HasVisionDisability, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsHasVisionDisability))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHearingDisabilityLevel() error {
	var err error
	p.out.HearingDisabilityLevel, err = ParseDisabilityLevel(p.values.Get(constants.FormParamsGetIndividualsHearingDisabilityLevel))
	return err
}

func (p *listIndividualsOptionsDecoder) parseHouseholdID() error {
	p.out.HouseholdID = p.values.Get(constants.FormParamsGetIndividualsHouseholdID)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseIDs() error {
	if len(p.values[constants.FormParamsGetIndividualsID]) == 0 {
		return nil
	}
	idSet := containers.NewStringSet()
	idSet.Add(p.values[constants.FormParamsGetIndividualsID]...)
	if idSet.IsEmpty() {
		return nil
	}
	p.out.IDs = idSet
	return nil
}

func (p *listIndividualsOptionsDecoder) parseIdentificationNumber() error {
	p.out.IdentificationNumber = p.values.Get(constants.FormParamsGetIndividualsIdentificationNumber)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseEngagementContext() error {
	if len(p.values[constants.FormParamsGetIndividualsEngagementContext]) == 0 {
		return nil
	}
	ecSet := containers.NewSet[EngagementContext]()
	for _, ec := range p.values[constants.FormParamsGetIndividualsEngagementContext] {
		parsedEc, err := ParseEngagementContext(ec)
		if err != nil {
			return err
		}
		ecSet.Add(parsedEc)
	}
	p.out.EngagementContext = ecSet
	return nil
}

func (p *listIndividualsOptionsDecoder) parseInternalID() error {
	p.out.InternalID = p.values.Get(constants.FormParamsGetIndividualsInternalID)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseIsHeadOfCommunity() error {
	var err error
	p.out.IsHeadOfCommunity, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsIsHeadOfCommunity))
	return err
}

func (p *listIndividualsOptionsDecoder) parseIsHeadOfHousehold() error {
	var err error
	p.out.IsHeadOfHousehold, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsIsHeadOfHousehold))
	return err
}

func (p *listIndividualsOptionsDecoder) parseIsFemaleHeadedHousehold() error {
	var err error
	p.out.IsFemaleHeadedHousehold, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsIsFemaleHeadedHousehold))
	return err
}

func (p *listIndividualsOptionsDecoder) parseIsMinorHeadedHousehold() error {
	var err error
	p.out.IsMinorHeadedHousehold, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsIsMinorHeadedHousehold))
	return err
}

func (p *listIndividualsOptionsDecoder) parseIsMinor() (err error) {
	p.out.IsMinor, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsIsMinor))
	return err
}

func (p *listIndividualsOptionsDecoder) parseMobilityDisabilityLevel() error {
	var err error
	p.out.MobilityDisabilityLevel, err = ParseDisabilityLevel(p.values.Get(constants.FormParamsGetIndividualsMobilityDisabilityLevel))
	return err
}

func (p *listIndividualsOptionsDecoder) parseNationality() error {
	p.out.Nationality = p.values.Get(constants.FormParamsGetIndividualsNationality)
	return nil
}

func (p *listIndividualsOptionsDecoder) parsePhoneNumber() error {
	p.out.PhoneNumber = p.values.Get(constants.FormParamsGetIndividualsPhoneNumber)
	return nil
}

func (p *listIndividualsOptionsDecoder) parsePreferredContactMethod() error {
	p.out.PreferredContactMethod = p.values.Get(constants.FormParamsGetIndividualsPreferredContactMethod)
	return nil
}

func (p *listIndividualsOptionsDecoder) parsePreferredCommunicationLanguage() error {
	p.out.PreferredCommunicationLanguage = p.values.Get(constants.FormParamsGetIndividualsPreferredCommunicationLanguage)
	return nil
}

func (p *listIndividualsOptionsDecoder) parsePrefersToRemainAnonymous() error {
	var err error
	p.out.PrefersToRemainAnonymous, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsPrefersToRemainAnonymous))
	return err
}

func (p *listIndividualsOptionsDecoder) parsePresentsProtectionConcerns() (err error) {
	p.out.PresentsProtectionConcerns, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsPresentsProtectionConcerns))
	return err
}

func (p *listIndividualsOptionsDecoder) parseSelfCareDisabilityLevel() error {
	var err error
	p.out.SelfCareDisabilityLevel, err = ParseDisabilityLevel(p.values.Get(constants.FormParamsGetIndividualsSelfCareDisabilityLevel))
	return err
}

func (p *listIndividualsOptionsDecoder) parseSpokenLanguage() error {
	p.out.SpokenLanguage = p.values.Get(constants.FormParamsGetIndividualsSpokenLanguage)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseUpdatedAtFrom() error {
	var err error
	p.out.UpdatedAtFrom, err = parseOptionalDateTime(p.values.Get(constants.FormParamsGetIndividualsUpdatedAtFrom))
	return err
}

func (p *listIndividualsOptionsDecoder) parseUpdatedAtTo() error {
	var err error
	p.out.UpdatedAtTo, err = parseOptionalDateTime(p.values.Get(constants.FormParamsGetIndividualsUpdatedAtTo))
	return err
}

func (p *listIndividualsOptionsDecoder) parseVisionDisabilityLevel() (err error) {
	p.out.VisionDisabilityLevel, err = ParseDisabilityLevel(p.values.Get(constants.FormParamsGetIndividualsVisionDisabilityLevel))
	return err
}

func (p *listIndividualsOptionsDecoder) parseSkip() (err error) {
	var skip *int
	skip, err = parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsSkip))
	if err != nil || skip == nil {
		return err
	}
	p.out.Skip = *skip
	if p.out.Skip < 0 {
		return fmt.Errorf("skip must be greater or equal to 0")
	}
	return nil
}

func (p *listIndividualsOptionsDecoder) parseTake() (err error) {
	var take *int
	take, err = parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsTake))
	if err != nil || take == nil {
		return err
	}
	p.out.Take = *take
	if p.out.Take < 0 {
		return fmt.Errorf("take must be greater or equal to 0")
	}
	return nil
}

func (p *listIndividualsOptionsDecoder) parseSort() (err error) {
	var sort = p.values.Get(constants.FormParamsGetIndividualsSort)
	if len(sort) == 0 {
		return nil
	}
	var out = &SortTerms{}
	if err := out.UnmarshalQuery(sort); err != nil {
		return err
	}
	p.out.Sort = *out
	return err
}

func parseOptionalInt(strValue string) (*int, error) {
	if len(strValue) == 0 {
		return nil, nil
	}
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}

func parseOptionalDate(strValue string) (*time.Time, error) {
	if len(strValue) == 0 {
		return nil, nil
	}
	dateValue, err := time.Parse("2006-01-02", strValue)
	if err != nil {
		return nil, err
	}
	return &dateValue, nil
}

func parseOptionalDateTime(strValue string) (*time.Time, error) {
	if len(strValue) == 0 {
		return nil, nil
	}
	dateTimeValue, err := time.Parse(time.RFC3339, strValue)
	if err != nil {
		return nil, err
	}
	return &dateTimeValue, nil
}

func parseOptionalBool(val string) (*bool, error) {
	if len(val) == 0 {
		return nil, nil
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &boolVal, nil
}
