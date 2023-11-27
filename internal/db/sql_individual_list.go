package db

import (
	"fmt"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/constants"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
)

type getAllIndividualsSQLQuery struct {
	*strings.Builder
	driverName string
	a          []interface{}
}

func newGetAllIndividualsSQLQuery(driverName string, options api.ListIndividualsOptions) *getAllIndividualsSQLQuery {
	qry := &getAllIndividualsSQLQuery{
		Builder:    &strings.Builder{},
		driverName: driverName,
	}
	qry = qry.
		writeString("SELECT * FROM individual_registrations WHERE deleted_at IS NULL").
		withInactive(options.Inactive).
		withAddress(options.Address).
		withAgeFrom(options.AgeFrom).
		withAgeTo(options.AgeTo).
		withBirthDateFrom(options.BirthDateFrom).
		withBirthDateTo(options.BirthDateTo).
		withCognitiveDisabilityLevel(options.CognitiveDisabilityLevel).
		withCollectionAdministrativeArea1(options.CollectionAdministrativeArea1).
		withCollectionAdministrativeArea2(options.CollectionAdministrativeArea2).
		withCollectionAdministrativeArea3(options.CollectionAdministrativeArea3).
		withCollectionOffice(options.CollectionOffice).
		withCollectionAgentName(options.CollectionAgentName).
		withCollectionAgentTitle(options.CollectionAgentTitle).
		withCollectionTimeFrom(options.CollectionTimeFrom).
		withCollectionTimeTo(options.CollectionTimeTo).
		withCommunityID(options.CommunityID).
		withCountryID(options.CountryID).
		withCreatedAtFrom(options.CreatedAtFrom).
		withCreatedAtTo(options.CreatedAtTo).
		withDisplacementStatuses(options.DisplacementStatuses).
		withEmail(options.Email).
		withFreeField1(options.FreeField1).
		withFreeField2(options.FreeField2).
		withFreeField3(options.FreeField3).
		withFreeField4(options.FreeField4).
		withFreeField5(options.FreeField5).
		withFullName(options.FullName).
		withMothersName(options.MothersName).
		withSexes(options.Sexes).
		withHasCognitiveDisability(options.HasCognitiveDisability).
		withHasCommunicationDisability(options.HasCommunicationDisability).
		withHasConsentedToRgpd(options.HasConsentedToRGPD).
		withHasConsentedToReferral(options.HasConsentedToReferral).
		withHasDisability(options.HasDisability).
		withHasHearingDisability(options.HasHearingDisability).
		withHasMobilityDisability(options.HasMobilityDisability).
		withHasSelfCareDisability(options.HasSelfCareDisability).
		withHasVisionDisability(options.HasVisionDisability).
		withHearingDisabilityLevel(options.HearingDisabilityLevel).
		withHouseholdID(options.HouseholdID).
		withIds(options.IDs).
		withIdentificationNumber(options.IdentificationNumber).
		withEngagementContext(options.EngagementContext).
		withInternalID(options.InternalID).
		withIsHeadOfCommunity(options.IsHeadOfCommunity).
		withIsHeadOfHousehold(options.IsHeadOfHousehold).
		withIsFemaleHeadedHousehold(options.IsFemaleHeadedHousehold).
		withIsMinorHeadedHousehold(options.IsMinorHeadedHousehold).
		withIsMinor(options.IsMinor).
		withBoolean(options.IsChildAtRisk, constants.DBColumnIndividualIsChildAtRisk).
		withBoolean(options.IsWomanAtRisk, constants.DBColumnIndividualIsWomanAtRisk).
		withBoolean(options.IsElderAtRisk, constants.DBColumnIndividualIsElderAtRisk).
		withBoolean(options.IsPregnant, constants.DBColumnIndividualIsPregnant).
		withBoolean(options.IsLactating, constants.DBColumnIndividualIsLactating).
		withBoolean(options.IsSeparatedChild, constants.DBColumnIndividualIsSeparatedChild).
		withBoolean(options.IsSingleParent, constants.DBColumnIndividualIsSingleParent).
		withBoolean(options.HasMedicalCondition, constants.DBColumnIndividualHasMedicalCondition).
		withBoolean(options.NeedsLegalAndPhysicalProtection, constants.DBColumnIndividualNeedsLegalAndPhysicalProtection).
		withMobilityDisabilityLevel(options.MobilityDisabilityLevel).
		withNationality(options.Nationality).
		withPhoneNumber(options.PhoneNumber).
		withPreferredContactMethod(options.PreferredContactMethod).
		withPreferredCommunicationLanguage(options.PreferredCommunicationLanguage).
		withPrefersToRemainAnonymous(options.PrefersToRemainAnonymous).
		withPresentsProtectionConcerns(options.PresentsProtectionConcerns).
		withPWDComments(options.PWDComments).
		withVulnerabilityComments(options.VulnerabilityComments).
		withSelfCareDisabilityLevel(options.SelfCareDisabilityLevel).
		withServiceCC(options.ServiceCC, options.ServiceRequestedDateFrom, options.ServiceRequestedDateTo, options.ServiceDeliveredDateFrom, options.ServiceDeliveredDateTo).
		withSpokenLanguage(options.SpokenLanguage).
		withUpdatedAtFrom(options.UpdatedAtFrom).
		withUpdatedAtTo(options.UpdatedAtTo).
		withVisionDisabilityLevel(options.VisionDisabilityLevel).

		// these must be in that order
		withSort(options.Sort).
		withOffset(options.Skip).
		withLimit(options.Take)

	return qry
}

func (g *getAllIndividualsSQLQuery) withInactive(inactive *bool) *getAllIndividualsSQLQuery {
	if inactive == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualInactive + " = ").writeArg(*inactive)
	return g
}

func (g *getAllIndividualsSQLQuery) withAddress(address string) *getAllIndividualsSQLQuery {
	if len(address) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND " + constants.DBColumnIndividualAddress + " LIKE ").writeArg("%" + address + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND " + constants.DBColumnIndividualAddress + " ILIKE ").writeArg("%" + address + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withAgeFrom(from *int) *getAllIndividualsSQLQuery {
	if from == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualAge + " >= ").writeArg(*from)
	return g
}

func (g *getAllIndividualsSQLQuery) withAgeTo(to *int) *getAllIndividualsSQLQuery {
	if to == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualAge + " <= ").writeArg(*to)
	return g
}

func (g *getAllIndividualsSQLQuery) withBirthDateFrom(from *time.Time) *getAllIndividualsSQLQuery {
	zero := &time.Time{}
	if from == nil || from.IsZero() || from == zero {
		return g
	}
	g.writeString(" AND birth_date >= ").writeArg(from)
	return g
}

func (g *getAllIndividualsSQLQuery) withBirthDateTo(to *time.Time) *getAllIndividualsSQLQuery {
	zero := &time.Time{}
	if to == nil || to.IsZero() || to == zero {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualBirthDate + " <= ").writeArg(to)
	return g
}

func (g *getAllIndividualsSQLQuery) withCognitiveDisabilityLevel(c enumTypes.DisabilityLevel) *getAllIndividualsSQLQuery {
	if c == enumTypes.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCognitiveDisabilityLevel + " = ").writeArg(string(c))
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAdministrativeArea1(area string) *getAllIndividualsSQLQuery {
	if len(area) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCollectionAdministrativeArea1 + " = ").writeArg(area)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAdministrativeArea2(area string) *getAllIndividualsSQLQuery {
	if len(area) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCollectionAdministrativeArea2 + " = ").writeArg(area)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAdministrativeArea3(area string) *getAllIndividualsSQLQuery {
	if len(area) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCollectionAdministrativeArea3 + " = ").writeArg(area)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionOffice(office string) *getAllIndividualsSQLQuery {
	if len(office) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCollectionOffice + " = ").writeArg(office)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAgentName(name string) *getAllIndividualsSQLQuery {
	if len(name) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND " + constants.DBColumnIndividualCollectionAgentName + " LIKE ")
		g.writeArg("%" + name + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND " + constants.DBColumnIndividualCollectionAgentName + " ILIKE ")
		g.writeArg("%" + name + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAgentTitle(title string) *getAllIndividualsSQLQuery {
	if len(title) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCollectionAgentTitle + " = ").writeArg(title)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionTimeFrom(d *time.Time) *getAllIndividualsSQLQuery {
	if d == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCollectionTime + " >= ").writeArg(d)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionTimeTo(d *time.Time) *getAllIndividualsSQLQuery {
	if d == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCollectionTime + " <= ").writeArg(d)
	return g
}

func (g *getAllIndividualsSQLQuery) withCommunityID(id string) *getAllIndividualsSQLQuery {
	if len(id) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCommunityID + " = ").writeArg(id)
	return g
}

func (g *getAllIndividualsSQLQuery) withCountryID(countryID string) *getAllIndividualsSQLQuery {
	if len(countryID) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCountryID + " = ").writeArg(countryID)
	return g
}

func (g *getAllIndividualsSQLQuery) withCreatedAtFrom(t *time.Time) *getAllIndividualsSQLQuery {
	if t == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCreatedAt + " >= ").writeArg(t)
	return g
}

func (g *getAllIndividualsSQLQuery) withCreatedAtTo(t *time.Time) *getAllIndividualsSQLQuery {
	if t == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualCreatedAt + " <= ").writeArg(t)
	return g
}

func (g *getAllIndividualsSQLQuery) withDisplacementStatuses(displacementStatuses containers.Set[enumTypes.DisplacementStatus]) *getAllIndividualsSQLQuery {
	if displacementStatuses.IsEmpty() {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualDisplacementStatus + " IN (")
	for i, ds := range displacementStatuses.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(string(ds))
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withEmail(email string) *getAllIndividualsSQLQuery {
	if len(email) == 0 {
		return g
	}
	normalizedEmail := strings.ToLower(email)
	g.writeString(" AND (")
	g.writeString(constants.DBColumnIndividualEmail1 + " = ").writeArg(normalizedEmail).writeString(" OR ")
	g.writeString(constants.DBColumnIndividualEmail2 + " = ").writeLastArg().writeString(" OR ")
	g.writeString(constants.DBColumnIndividualEmail3 + " = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField1(freeField1 string) *getAllIndividualsSQLQuery {
	if len(freeField1) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualFreeField1 + " = ").writeArg(freeField1)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField2(freeField2 string) *getAllIndividualsSQLQuery {
	if len(freeField2) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualFreeField2 + " = ").writeArg(freeField2)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField3(freeField3 string) *getAllIndividualsSQLQuery {
	if len(freeField3) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualFreeField3 + " = ").writeArg(freeField3)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField4(freeField4 string) *getAllIndividualsSQLQuery {
	if len(freeField4) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualFreeField4 + " = ").writeArg(freeField4)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField5(freeField5 string) *getAllIndividualsSQLQuery {
	if len(freeField5) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualFreeField5 + " = ").writeArg(freeField5)
	return g
}

func (g *getAllIndividualsSQLQuery) withMothersName(mothersName string) *getAllIndividualsSQLQuery {
	if len(mothersName) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND " + constants.DBColumnIndividualMothersName + " LIKE ")
		g.writeArg("%" + mothersName + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND " + constants.DBColumnIndividualMothersName + " ILIKE ")
		g.writeArg("%" + mothersName + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withFullName(name string) *getAllIndividualsSQLQuery {
	if len(name) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND (" + constants.DBColumnIndividualFullName + " LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualPreferredName + " LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualFirstName + " LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualMiddleName + " LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualLastName + " LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualNativeName + " LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(")")
	} else if g.driverName == "postgres" {
		g.writeString(" AND (" + constants.DBColumnIndividualFullName + " ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualPreferredName + " ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualFirstName + " ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualMiddleName + " ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualLastName + " ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR " + constants.DBColumnIndividualNativeName + " ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(")")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withSexes(sexes containers.Set[enumTypes.Sex]) *getAllIndividualsSQLQuery {
	if len(sexes) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualSex + " IN (")
	for i, sex := range sexes.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(string(sex))
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withHasCognitiveDisability(hasCognitiveDisability *bool) *getAllIndividualsSQLQuery {
	if hasCognitiveDisability == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasCognitiveDisability + " = ").writeArg(*hasCognitiveDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasCommunicationDisability(hasCommunicationDisability *bool) *getAllIndividualsSQLQuery {
	if hasCommunicationDisability == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasCommunicationDisability + " = ").writeArg(*hasCommunicationDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasConsentedToRgpd(hasConsentedToRgpd *bool) *getAllIndividualsSQLQuery {
	if hasConsentedToRgpd == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasConsentedToRGPD + " = ").writeArg(*hasConsentedToRgpd)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasConsentedToReferral(hasConsentedToReferral *bool) *getAllIndividualsSQLQuery {
	if hasConsentedToReferral == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasConsentedToReferral + " = ").writeArg(*hasConsentedToReferral)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasHearingDisability(hasHearingDisability *bool) *getAllIndividualsSQLQuery {
	if hasHearingDisability == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasHearingDisability + " = ").writeArg(*hasHearingDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasDisability(hasDisability *bool) *getAllIndividualsSQLQuery {
	if hasDisability == nil {
		return g
	}
	g.writeString(" AND has_disability = ").writeArg(*hasDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasMobilityDisability(hasMobilityDisability *bool) *getAllIndividualsSQLQuery {
	if hasMobilityDisability == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasMobilityDisability + " = ").writeArg(*hasMobilityDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasSelfCareDisability(hasSelfCareDisability *bool) *getAllIndividualsSQLQuery {
	if hasSelfCareDisability == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasSelfCareDisability + " = ").writeArg(*hasSelfCareDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasVisionDisability(hasVisionDisability *bool) *getAllIndividualsSQLQuery {
	if hasVisionDisability == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHasVisionDisability + " = ").writeArg(*hasVisionDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHearingDisabilityLevel(hearingDisabilityLevel enumTypes.DisabilityLevel) *getAllIndividualsSQLQuery {
	if hearingDisabilityLevel == enumTypes.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHearingDisabilityLevel + " = ").writeArg(string(hearingDisabilityLevel))
	return g
}

func (g *getAllIndividualsSQLQuery) withHouseholdID(householdID string) *getAllIndividualsSQLQuery {
	if len(householdID) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualHouseholdID + " = ").writeArg(householdID)
	return g
}

func (g *getAllIndividualsSQLQuery) withIds(ids containers.StringSet) *getAllIndividualsSQLQuery {
	if ids.Len() == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualID + " IN (")
	for i, id := range ids.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(id)
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withIdentificationNumber(identificationNumber string) *getAllIndividualsSQLQuery {
	if len(identificationNumber) == 0 {
		return g
	}
	g.writeString(" AND (")
	g.writeString(constants.DBColumnIndividualIdentificationNumber1 + " = ").writeArg(identificationNumber)
	g.writeString(" OR " + constants.DBColumnIndividualIdentificationNumber2 + " = ").writeLastArg()
	g.writeString(" OR " + constants.DBColumnIndividualIdentificationNumber3 + " = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withServiceCC(serviceCC containers.Set[enumTypes.ServiceCC], requestedFrom *time.Time, requestedTo *time.Time, deliveredFrom *time.Time, deliveredTo *time.Time) *getAllIndividualsSQLQuery {
	zero := &time.Time{}
	requestedFromIsUndefined := requestedFrom == nil || requestedFrom.IsZero() || requestedFrom == zero
	requestedToIsUndefined := requestedTo == nil || requestedTo.IsZero() || requestedTo == zero
	deliveredFromIsUndefined := deliveredFrom == nil || deliveredFrom.IsZero() || deliveredFrom == zero
	deliveredToIsUndefined := deliveredTo == nil || deliveredTo.IsZero() || deliveredTo == zero

	if requestedFromIsUndefined && requestedToIsUndefined && deliveredFromIsUndefined && deliveredToIsUndefined && serviceCC.IsEmpty() {
		return g
	}

	items := "'"
	for i, ds := range serviceCC.Items() {
		if i != 0 {
			items += "','"
		}
		items += string(ds)
	}
	items += "'"

	query := " AND ("
	for i := 1; i <= 7; i++ {
		if i != 1 {
			query += " OR "
		}
		query += "("
		var conditions []string
		if !serviceCC.IsEmpty() {
			conditions = append(conditions, fmt.Sprintf("service_cc_%d IN (%s)", i, items))
		}
		if !requestedToIsUndefined {
			conditions = append(conditions, fmt.Sprintf("service_requested_date_%d <= '%s'", i, requestedTo.Format("2006-01-02")))
		}
		if !requestedFromIsUndefined {
			conditions = append(conditions, fmt.Sprintf("service_requested_date_%d >= '%s'", i, requestedFrom.Format("2006-01-02")))
		}
		if !deliveredToIsUndefined {
			conditions = append(conditions, fmt.Sprintf("service_delivered_date_%d <= '%s'", i, deliveredTo.Format("2006-01-02")))
		}
		if !deliveredFromIsUndefined {
			conditions = append(conditions, fmt.Sprintf("service_delivered_date_%d >= '%s'", i, deliveredFrom.Format("2006-01-02")))
		}
		query += strings.Join(conditions, " AND ")
		query += ")"
	}
	query += ")"
	g.writeString(query)
	return g
}

func (g *getAllIndividualsSQLQuery) withEngagementContext(engagementContext containers.Set[enumTypes.EngagementContext]) *getAllIndividualsSQLQuery {
	if engagementContext.IsEmpty() {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualEngagementContext + " IN (")
	for i, ds := range engagementContext.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(string(ds))
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withInternalID(internalID string) *getAllIndividualsSQLQuery {
	if len(internalID) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualInternalID + " = ").writeArg(internalID)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsHeadOfCommunity(isHeadOfCommunity *bool) *getAllIndividualsSQLQuery {
	if isHeadOfCommunity == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualIsHeadOfCommunity + " = ").writeArg(*isHeadOfCommunity)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsHeadOfHousehold(isHeadOfHousehold *bool) *getAllIndividualsSQLQuery {
	if isHeadOfHousehold == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualIsHeadOfHousehold + " = ").writeArg(*isHeadOfHousehold)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsFemaleHeadedHousehold(isFemaleHeadedHousehold *bool) *getAllIndividualsSQLQuery {
	if isFemaleHeadedHousehold == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualIsFemaleHeadedHousehold + " = ").writeArg(*isFemaleHeadedHousehold)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsMinorHeadedHousehold(isMinorHeadedHousehold *bool) *getAllIndividualsSQLQuery {
	if isMinorHeadedHousehold == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualIsMinorHeadedHousehold + " = ").writeArg(*isMinorHeadedHousehold)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsMinor(isMinor *bool) *getAllIndividualsSQLQuery {
	if isMinor == nil {
		return g
	}
	if *isMinor {
		g.writeString(" AND " + constants.DBColumnIndividualIsMinor + " = ").writeArg(true)
	} else {
		g.writeString(" AND " + constants.DBColumnIndividualIsMinor + " = ").writeArg(false)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withBoolean(b *bool, dbColumnName string) *getAllIndividualsSQLQuery {
	if b == nil {
		return g
	}
	if *b {
		g.writeString(" AND " + dbColumnName + " = ").writeArg(true)
	} else {
		g.writeString(" AND " + dbColumnName + " = ").writeArg(false)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withMobilityDisabilityLevel(mobilityDisabilityLevel enumTypes.DisabilityLevel) *getAllIndividualsSQLQuery {
	if mobilityDisabilityLevel == enumTypes.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualMobilityDisabilityLevel + " = ").writeArg(string(mobilityDisabilityLevel))
	return g
}

func (g *getAllIndividualsSQLQuery) withNationality(nationality string) *getAllIndividualsSQLQuery {
	if len(nationality) == 0 {
		return g
	}
	g.writeString(" AND (")
	g.writeString(constants.DBColumnIndividualNationality1 + " = ").writeArg(nationality)
	g.writeString(" OR " + constants.DBColumnIndividualNationality2 + " = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withPhoneNumber(phoneNumber string) *getAllIndividualsSQLQuery {
	if len(phoneNumber) == 0 {
		return g
	}
	normalizedPhoneNumber := api.NormalizePhoneNumber(phoneNumber)
	if g.driverName == "sqlite" {
		g.writeString(" AND (")
		g.writeString(" " + constants.DBColumnIndividualNormalizedPhoneNumber1 + " LIKE ").writeArg("%" + normalizedPhoneNumber + "%").writeString(" OR ")
		g.writeString(" " + constants.DBColumnIndividualNormalizedPhoneNumber2 + " LIKE ").writeLastArg().writeString(" OR ")
		g.writeString(" " + constants.DBColumnIndividualNormalizedPhoneNumber3 + " LIKE ").writeLastArg()
		g.writeString(")")
	} else if g.driverName == "postgres" {
		g.writeString(" AND (")
		g.writeString(constants.DBColumnIndividualNormalizedPhoneNumber1 + " ILIKE ").writeArg("%" + normalizedPhoneNumber + "%").writeString(" OR ")
		g.writeString(constants.DBColumnIndividualNormalizedPhoneNumber2 + " ILIKE ").writeLastArg().writeString(" OR ")
		g.writeString(constants.DBColumnIndividualNormalizedPhoneNumber3 + " ILIKE ").writeLastArg()
		g.writeString(")")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withPreferredContactMethod(preferredContactMethod string) *getAllIndividualsSQLQuery {
	if len(preferredContactMethod) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualPreferredContactMethod + " = ").writeArg(preferredContactMethod)
	return g
}

func (g *getAllIndividualsSQLQuery) withPreferredCommunicationLanguage(language string) *getAllIndividualsSQLQuery {
	if len(language) == 0 {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualPreferredCommunicationLanguage + " = ").writeArg(language)
	return g
}

func (g *getAllIndividualsSQLQuery) withPrefersToRemainAnonymous(prefersToRemainAnonymous *bool) *getAllIndividualsSQLQuery {
	if prefersToRemainAnonymous == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualPrefersToRemainAnonymous + " = ").writeArg(*prefersToRemainAnonymous)
	return g
}

func (g *getAllIndividualsSQLQuery) withPresentsProtectionConcerns(presentsProtectionConcerns *bool) *getAllIndividualsSQLQuery {
	if presentsProtectionConcerns == nil {
		return g
	}
	if *presentsProtectionConcerns {
		g.writeString(" AND " + constants.DBColumnIndividualPresentsProtectionConcerns + " = ").writeArg(true)
	} else {
		g.writeString(" AND " + constants.DBColumnIndividualPresentsProtectionConcerns + " = ").writeArg(false)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withPWDComments(pwdComments string) *getAllIndividualsSQLQuery {
	if len(pwdComments) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND " + constants.DBColumnIndividualPWDComments + " LIKE ")
		g.writeArg("%" + pwdComments + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND " + constants.DBColumnIndividualPWDComments + " ILIKE ")
		g.writeArg("%" + pwdComments + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withVulnerabilityComments(vulnerabilityComments string) *getAllIndividualsSQLQuery {
	if len(vulnerabilityComments) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND " + constants.DBColumnIndividualVulnerabilityComments + " LIKE ")
		g.writeArg("%" + vulnerabilityComments + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND " + constants.DBColumnIndividualVulnerabilityComments + " ILIKE ")
		g.writeArg("%" + vulnerabilityComments + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withSelfCareDisabilityLevel(selfCareDisabilityLevel enumTypes.DisabilityLevel) *getAllIndividualsSQLQuery {
	if selfCareDisabilityLevel == enumTypes.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualSelfCareDisabilityLevel + " = ").writeArg(string(selfCareDisabilityLevel))
	return g
}

func (g *getAllIndividualsSQLQuery) withSpokenLanguage(language string) *getAllIndividualsSQLQuery {
	if len(language) == 0 {
		return g
	}
	g.writeString(" AND (")
	g.writeString(constants.DBColumnIndividualSpokenLanguage1 + " = ").writeArg(language)
	g.writeString(" OR " + constants.DBColumnIndividualSpokenLanguage2 + " = ").writeLastArg()
	g.writeString(" OR " + constants.DBColumnIndividualSpokenLanguage3 + " = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withUpdatedAtFrom(updatedAtFrom *time.Time) *getAllIndividualsSQLQuery {
	if updatedAtFrom == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualUpdatedAt + " >= ").writeArg(*updatedAtFrom)
	return g
}

func (g *getAllIndividualsSQLQuery) withUpdatedAtTo(updatedAtTo *time.Time) *getAllIndividualsSQLQuery {
	if updatedAtTo == nil {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualUpdatedAt + " <= ").writeArg(*updatedAtTo)
	return g
}

func (g *getAllIndividualsSQLQuery) withVisionDisabilityLevel(visionDisabilityLevel enumTypes.DisabilityLevel) *getAllIndividualsSQLQuery {
	if visionDisabilityLevel == enumTypes.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND " + constants.DBColumnIndividualVisionDisabilityLevel + " = ").writeArg(string(visionDisabilityLevel))
	return g
}

func (g *getAllIndividualsSQLQuery) withOrderBy(field string) *getAllIndividualsSQLQuery {
	if len(field) == 0 {
		return g
	}
	g.writeString(" ORDER BY ").writeString(field)
	return g
}

func (g *getAllIndividualsSQLQuery) withLimit(limit int) *getAllIndividualsSQLQuery {
	if limit == 0 {
		return g
	}
	g.writeString(fmt.Sprintf(" LIMIT %d", limit))
	return g
}

func (g *getAllIndividualsSQLQuery) withOffset(offset int) *getAllIndividualsSQLQuery {
	if offset == 0 {
		return g
	}
	g.writeString(fmt.Sprintf(" OFFSET %d", offset))
	return g
}

func (g *getAllIndividualsSQLQuery) withSort(sortTerms api.SortTerms) *getAllIndividualsSQLQuery {
	if len(sortTerms) == 0 {
		return g
	}
	g.writeString(" ORDER BY ")
	for i, sortTerm := range sortTerms {
		if i > 0 {
			g.writeString(", ")
		}
		g.writeString(sortTerm.Field)
		if sortTerm.Direction == api.SortDirectionDescending {
			g.writeString(" DESC")
		} else {
			g.writeString(" ASC")
		}
	}
	return g
}

func (g *getAllIndividualsSQLQuery) build() (string, []interface{}) {
	return g.sql(), g.sqlArgs()
}

func (g *getAllIndividualsSQLQuery) sql() string {
	return g.Builder.String()
}

func (g *getAllIndividualsSQLQuery) sqlArgs() []interface{} {
	return g.a
}

func (g *getAllIndividualsSQLQuery) writeString(str string) *getAllIndividualsSQLQuery {
	_, _ = g.Builder.WriteString(str)
	return g
}

func (g *getAllIndividualsSQLQuery) writeArg(arg interface{}) *getAllIndividualsSQLQuery {
	g.a = append(g.a, arg)
	return g.writeLastArg()
}

func (g *getAllIndividualsSQLQuery) writeArgNum(i int) *getAllIndividualsSQLQuery {
	g.writeString(fmt.Sprintf("$%d", i))
	return g
}

func (g *getAllIndividualsSQLQuery) writeLastArg() *getAllIndividualsSQLQuery {
	return g.writeArgNum(len(g.a))
}

func (g *getAllIndividualsSQLQuery) writeStringArgs(sep string, args ...string) *getAllIndividualsSQLQuery {
	for i, arg := range args {
		if i > 0 {
			g.writeString(sep)
		}
		g.writeArg(arg)
	}
	return g
}
