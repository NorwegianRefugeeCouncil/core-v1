package db

import (
	"fmt"
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
		withAddress(options.Address).
		withAgeFrom(options.AgeFrom).
		withAgeTo(options.AgeTo).
		withBirthDateFrom(options.BirthDateFrom).
		withBirthDateTo(options.BirthDateTo).
		withCognitiveDisabilityLevel(options.CognitiveDisabilityLevel).
		withCollectionAdministrativeArea1(options.CollectionAdministrativeArea1).
		withCollectionAdministrativeArea2(options.CollectionAdministrativeArea2).
		withCollectionAdministrativeArea3(options.CollectionAdministrativeArea3).
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
		withGenders(options.Genders).
		withHasCognitiveDisability(options.HasCognitiveDisability).
		withHasCommunicationDisability(options.HasCommunicationDisability).
		withHasConsentedToRgpd(options.HasConsentedToRGPD).
		withHasConsentedToReferral(options.HasConsentedToReferral).
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
		withIsMinor(options.IsMinor).
		withMobilityDisabilityLevel(options.MobilityDisabilityLevel).
		withNationality(options.Nationality).
		withPhoneNumber(options.PhoneNumber).
		withPreferredContactMethod(options.PreferredContactMethod).
		withPreferredCommunicationLanguage(options.PreferredCommunicationLanguage).
		withPrefersToRemainAnonymous(options.PrefersToRemainAnonymous).
		withPresentsProtectionConcerns(options.PresentsProtectionConcerns).
		withSelfCareDisabilityLevel(options.SelfCareDisabilityLevel).
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

func (g *getAllIndividualsSQLQuery) withAddress(address string) *getAllIndividualsSQLQuery {
	if len(address) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND address LIKE ").writeArg("%" + address + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND address ILIKE ").writeArg("%" + address + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withAgeFrom(from *int) *getAllIndividualsSQLQuery {
	if from == nil {
		return g
	}
	g.writeString(" AND age >= ").writeArg(*from)
	return g
}

func (g *getAllIndividualsSQLQuery) withAgeTo(to *int) *getAllIndividualsSQLQuery {
	if to == nil {
		return g
	}
	g.writeString(" AND age <= ").writeArg(*to)
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
	g.writeString(" AND birth_date <= ").writeArg(to)
	return g
}

func (g *getAllIndividualsSQLQuery) withCognitiveDisabilityLevel(c api.DisabilityLevel) *getAllIndividualsSQLQuery {
	if c == api.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND cognitive_disability_level = ").writeArg(string(c))
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAdministrativeArea1(area string) *getAllIndividualsSQLQuery {
	if len(area) == 0 {
		return g
	}
	g.writeString(" AND collection_administrative_area_1 = ").writeArg(area)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAdministrativeArea2(area string) *getAllIndividualsSQLQuery {
	if len(area) == 0 {
		return g
	}
	g.writeString(" AND collection_administrative_area_2 = ").writeArg(area)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAdministrativeArea3(area string) *getAllIndividualsSQLQuery {
	if len(area) == 0 {
		return g
	}
	g.writeString(" AND collection_administrative_area_3 = ").writeArg(area)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAgentName(name string) *getAllIndividualsSQLQuery {
	if len(name) == 0 {
		return g
	}
	g.writeString(" AND collection_agent_name = ").writeArg(name)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionAgentTitle(title string) *getAllIndividualsSQLQuery {
	if len(title) == 0 {
		return g
	}
	g.writeString(" AND collection_agent_title = ").writeArg(title)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionTimeFrom(d *time.Time) *getAllIndividualsSQLQuery {
	if d == nil {
		return g
	}
	g.writeString(" AND collection_time >= ").writeArg(d)
	return g
}

func (g *getAllIndividualsSQLQuery) withCollectionTimeTo(d *time.Time) *getAllIndividualsSQLQuery {
	if d == nil {
		return g
	}
	g.writeString(" AND collection_time <= ").writeArg(d)
	return g
}

func (g *getAllIndividualsSQLQuery) withCommunityID(id string) *getAllIndividualsSQLQuery {
	if len(id) == 0 {
		return g
	}
	g.writeString(" AND community_id = ").writeArg(id)
	return g
}

func (g *getAllIndividualsSQLQuery) withCountryID(countryID string) *getAllIndividualsSQLQuery {
	if len(countryID) == 0 {
		return g
	}
	g.writeString(" AND country_id = ").writeArg(countryID)
	return g
}

func (g *getAllIndividualsSQLQuery) withCreatedAtFrom(t *time.Time) *getAllIndividualsSQLQuery {
	if t == nil {
		return g
	}
	g.writeString(" AND created_at >= ").writeArg(t)
	return g
}

func (g *getAllIndividualsSQLQuery) withCreatedAtTo(t *time.Time) *getAllIndividualsSQLQuery {
	if t == nil {
		return g
	}
	g.writeString(" AND created_at <= ").writeArg(t)
	return g
}

func (g *getAllIndividualsSQLQuery) withDisplacementStatuses(displacementStatuses containers.Set[api.DisplacementStatus]) *getAllIndividualsSQLQuery {
	if displacementStatuses.IsEmpty() {
		return g
	}
	g.writeString(" AND displacement_status IN (")
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
	g.writeString("email_1 = ").writeArg(normalizedEmail).writeString(" OR ")
	g.writeString("email_2 = ").writeLastArg().writeString(" OR ")
	g.writeString("email_3 = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField1(freeField1 string) *getAllIndividualsSQLQuery {
	if len(freeField1) == 0 {
		return g
	}
	g.writeString(" AND free_field_1 = ").writeArg(freeField1)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField2(freeField2 string) *getAllIndividualsSQLQuery {
	if len(freeField2) == 0 {
		return g
	}
	g.writeString(" AND free_field_2 = ").writeArg(freeField2)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField3(freeField3 string) *getAllIndividualsSQLQuery {
	if len(freeField3) == 0 {
		return g
	}
	g.writeString(" AND free_field_3 = ").writeArg(freeField3)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField4(freeField4 string) *getAllIndividualsSQLQuery {
	if len(freeField4) == 0 {
		return g
	}
	g.writeString(" AND free_field_4 = ").writeArg(freeField4)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField5(freeField5 string) *getAllIndividualsSQLQuery {
	if len(freeField5) == 0 {
		return g
	}
	g.writeString(" AND free_field_5 = ").writeArg(freeField5)
	return g
}

func (g *getAllIndividualsSQLQuery) withFullName(name string) *getAllIndividualsSQLQuery {
	if len(name) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND (full_name LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR preferred_name LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(")")
	} else if g.driverName == "postgres" {
		g.writeString(" AND (full_name ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR preferred_name ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(")")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withGenders(genders containers.Set[api.Gender]) *getAllIndividualsSQLQuery {
	if len(genders) == 0 {
		return g
	}
	g.writeString(" AND gender IN (")
	for i, gender := range genders.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(string(gender))
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withHasCognitiveDisability(hasCognitiveDisability *bool) *getAllIndividualsSQLQuery {
	if hasCognitiveDisability == nil {
		return g
	}
	g.writeString(" AND has_cognitive_disability = ").writeArg(*hasCognitiveDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasCommunicationDisability(hasCommunicationDisability *bool) *getAllIndividualsSQLQuery {
	if hasCommunicationDisability == nil {
		return g
	}
	g.writeString(" AND has_communication_disability = ").writeArg(*hasCommunicationDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasConsentedToRgpd(hasConsentedToRgpd *bool) *getAllIndividualsSQLQuery {
	if hasConsentedToRgpd == nil {
		return g
	}
	g.writeString(" AND has_consented_to_rgpd = ").writeArg(*hasConsentedToRgpd)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasConsentedToReferral(hasConsentedToReferral *bool) *getAllIndividualsSQLQuery {
	if hasConsentedToReferral == nil {
		return g
	}
	g.writeString(" AND has_consented_to_referral = ").writeArg(*hasConsentedToReferral)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasHearingDisability(hasHearingDisability *bool) *getAllIndividualsSQLQuery {
	if hasHearingDisability == nil {
		return g
	}
	g.writeString(" AND has_hearing_disability = ").writeArg(*hasHearingDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasMobilityDisability(hasMobilityDisability *bool) *getAllIndividualsSQLQuery {
	if hasMobilityDisability == nil {
		return g
	}
	g.writeString(" AND has_mobility_disability = ").writeArg(*hasMobilityDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasSelfCareDisability(hasSelfCareDisability *bool) *getAllIndividualsSQLQuery {
	if hasSelfCareDisability == nil {
		return g
	}
	g.writeString(" AND has_selfcare_disability = ").writeArg(*hasSelfCareDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHasVisionDisability(hasVisionDisability *bool) *getAllIndividualsSQLQuery {
	if hasVisionDisability == nil {
		return g
	}
	g.writeString(" AND has_vision_disability = ").writeArg(*hasVisionDisability)
	return g
}

func (g *getAllIndividualsSQLQuery) withHearingDisabilityLevel(hearingDisabilityLevel api.DisabilityLevel) *getAllIndividualsSQLQuery {
	if hearingDisabilityLevel == api.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND hearing_disability_level = ").writeArg(string(hearingDisabilityLevel))
	return g
}

func (g *getAllIndividualsSQLQuery) withHouseholdID(householdID string) *getAllIndividualsSQLQuery {
	if len(householdID) == 0 {
		return g
	}
	g.writeString(" AND household_id = ").writeArg(householdID)
	return g
}

func (g *getAllIndividualsSQLQuery) withIds(ids containers.StringSet) *getAllIndividualsSQLQuery {
	if len(ids) == 0 {
		return g
	}
	g.writeString(" AND id IN (")
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
	g.writeString("identification_number_1 = ").writeArg(identificationNumber)
	g.writeString(" OR identification_number_2 = ").writeLastArg()
	g.writeString(" OR identification_number_3 = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withEngagementContext(engagementContext containers.Set[api.EngagementContext]) *getAllIndividualsSQLQuery {
	if engagementContext.IsEmpty() {
		return g
	}
	g.writeString(" AND engagement_context IN (")
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
	g.writeString(" AND internal_id = ").writeArg(internalID)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsHeadOfCommunity(isHeadOfCommunity *bool) *getAllIndividualsSQLQuery {
	if isHeadOfCommunity == nil {
		return g
	}
	g.writeString(" AND is_head_of_community = ").writeArg(*isHeadOfCommunity)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsHeadOfHousehold(isHeadOfHousehold *bool) *getAllIndividualsSQLQuery {
	if isHeadOfHousehold == nil {
		return g
	}
	g.writeString(" AND is_head_of_household = ").writeArg(*isHeadOfHousehold)
	return g
}

func (g *getAllIndividualsSQLQuery) withIsMinor(isMinor *bool) *getAllIndividualsSQLQuery {
	if isMinor == nil {
		return g
	}
	if *isMinor {
		g.writeString(" AND is_minor = ").writeArg(true)
	} else {
		g.writeString(" AND is_minor = ").writeArg(false)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withMobilityDisabilityLevel(mobilityDisabilityLevel api.DisabilityLevel) *getAllIndividualsSQLQuery {
	if mobilityDisabilityLevel == api.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND mobility_disability_level = ").writeArg(string(mobilityDisabilityLevel))
	return g
}

func (g *getAllIndividualsSQLQuery) withNationality(nationality string) *getAllIndividualsSQLQuery {
	if len(nationality) == 0 {
		return g
	}
	g.writeString(" AND (")
	g.writeString("nationality_1 = ").writeArg(nationality)
	g.writeString(" OR nationality_2 = ").writeLastArg()
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
		g.writeString(" normalized_phone_number_1 LIKE ").writeArg("%" + normalizedPhoneNumber + "%").writeString(" OR ")
		g.writeString(" normalized_phone_number_2 LIKE ").writeLastArg().writeString(" OR ")
		g.writeString(" normalized_phone_number_3 LIKE ").writeLastArg()
		g.writeString(")")
	} else if g.driverName == "postgres" {
		g.writeString(" AND (")
		g.writeString("normalized_phone_number_1 ILIKE ").writeArg("%" + normalizedPhoneNumber + "%").writeString(" OR ")
		g.writeString("normalized_phone_number_2 ILIKE ").writeLastArg().writeString(" OR ")
		g.writeString("normalized_phone_number_3 ILIKE ").writeLastArg()
		g.writeString(")")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withPreferredContactMethod(preferredContactMethod string) *getAllIndividualsSQLQuery {
	if len(preferredContactMethod) == 0 {
		return g
	}
	g.writeString(" AND preferred_contact_method = ").writeArg(preferredContactMethod)
	return g
}

func (g *getAllIndividualsSQLQuery) withPreferredCommunicationLanguage(language string) *getAllIndividualsSQLQuery {
	if len(language) == 0 {
		return g
	}
	g.writeString(" AND preferred_communication_language = ").writeArg(language)
	return g
}

func (g *getAllIndividualsSQLQuery) withPrefersToRemainAnonymous(prefersToRemainAnonymous *bool) *getAllIndividualsSQLQuery {
	if prefersToRemainAnonymous == nil {
		return g
	}
	g.writeString(" AND prefers_to_remain_anonymous = ").writeArg(*prefersToRemainAnonymous)
	return g
}

func (g *getAllIndividualsSQLQuery) withPresentsProtectionConcerns(presentsProtectionConcerns *bool) *getAllIndividualsSQLQuery {
	if presentsProtectionConcerns == nil {
		return g
	}
	if *presentsProtectionConcerns {
		g.writeString(" AND presents_protection_concerns = ").writeArg(true)
	} else {
		g.writeString(" AND presents_protection_concerns = ").writeArg(false)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withSelfCareDisabilityLevel(selfCareDisabilityLevel api.DisabilityLevel) *getAllIndividualsSQLQuery {
	if selfCareDisabilityLevel == api.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND selfcare_disability_level = ").writeArg(string(selfCareDisabilityLevel))
	return g
}

func (g *getAllIndividualsSQLQuery) withSpokenLanguage(language string) *getAllIndividualsSQLQuery {
	if len(language) == 0 {
		return g
	}
	g.writeString(" AND (")
	g.writeString("spoken_language_1 = ").writeArg(language)
	g.writeString(" OR spoken_language_2 = ").writeLastArg()
	g.writeString(" OR spoken_language_3 = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withUpdatedAtFrom(updatedAtFrom *time.Time) *getAllIndividualsSQLQuery {
	if updatedAtFrom == nil {
		return g
	}
	g.writeString(" AND updated_at >= ").writeArg(*updatedAtFrom)
	return g
}

func (g *getAllIndividualsSQLQuery) withUpdatedAtTo(updatedAtTo *time.Time) *getAllIndividualsSQLQuery {
	if updatedAtTo == nil {
		return g
	}
	g.writeString(" AND updated_at <= ").writeArg(*updatedAtTo)
	return g
}

func (g *getAllIndividualsSQLQuery) withVisionDisabilityLevel(visionDisabilityLevel api.DisabilityLevel) *getAllIndividualsSQLQuery {
	if visionDisabilityLevel == api.DisabilityLevelUnspecified {
		return g
	}
	g.writeString(" AND vision_disability_level = ").writeArg(string(visionDisabilityLevel))
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
