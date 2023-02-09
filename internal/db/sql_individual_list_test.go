package db

import (
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func Test_newGetAllIndividualsSQLQuery(t *testing.T) {
	someDate, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	zeroTime := time.Time{}
	const defaultQuery = `SELECT * FROM individual_registrations WHERE deleted_at IS NULL`
	tests := []struct {
		name     string
		args     api.ListIndividualsOptions
		wantSql  string
		wantArgs []interface{}
	}{
		{
			name:    "empty",
			args:    api.ListIndividualsOptions{},
			wantSql: defaultQuery,
		}, {
			name:     "inactive (true)",
			args:     api.ListIndividualsOptions{Inactive: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND inactive = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "inactive (false)",
			args:     api.ListIndividualsOptions{Inactive: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND inactive = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "address",
			args:     api.ListIndividualsOptions{Address: "123 Main St"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND address ILIKE $1`,
			wantArgs: []interface{}{"%123 Main St%"},
		}, {
			name:     "ageFrom",
			args:     api.ListIndividualsOptions{AgeFrom: pointers.Int(18)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND age >= $1`,
			wantArgs: []interface{}{18},
		}, {
			name:     "ageTo",
			args:     api.ListIndividualsOptions{AgeTo: pointers.Int(18)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND age <= $1`,
			wantArgs: []interface{}{18},
		}, {
			name:     "birthDateFrom",
			args:     api.ListIndividualsOptions{BirthDateFrom: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND birth_date >= $1`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:    "birthDateFrom (zero)",
			args:    api.ListIndividualsOptions{BirthDateFrom: &zeroTime},
			wantSql: defaultQuery,
		}, {
			name:     "birthDateTo",
			args:     api.ListIndividualsOptions{BirthDateTo: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND birth_date <= $1`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:    "birthDateTo (zero)",
			args:    api.ListIndividualsOptions{BirthDateFrom: &zeroTime},
			wantSql: defaultQuery,
		}, {
			name:     "cognitiveDisabilityLevel",
			args:     api.ListIndividualsOptions{CognitiveDisabilityLevel: enumTypes.DisabilityLevelMild},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND cognitive_disability_level = $1`,
			wantArgs: []interface{}{"mild"},
		}, {
			name:     "collectionAdministrativeArea1",
			args:     api.ListIndividualsOptions{CollectionAdministrativeArea1: "area1"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_administrative_area_1 = $1`,
			wantArgs: []interface{}{"area1"},
		}, {
			name:     "collectionAdministrativeArea2",
			args:     api.ListIndividualsOptions{CollectionAdministrativeArea2: "area2"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_administrative_area_2 = $1`,
			wantArgs: []interface{}{"area2"},
		}, {
			name:     "collectionAdministrativeArea3",
			args:     api.ListIndividualsOptions{CollectionAdministrativeArea3: "area3"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_administrative_area_3 = $1`,
			wantArgs: []interface{}{"area3"},
		}, {
			name:     "collectionOffice",
			args:     api.ListIndividualsOptions{CollectionOffice: "office"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_office = $1`,
			wantArgs: []interface{}{"office"},
		}, {
			name:     "collectionAgentName",
			args:     api.ListIndividualsOptions{CollectionAgentName: "agent"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_agent_name = $1`,
			wantArgs: []interface{}{"agent"},
		}, {
			name:     "collectionAgentTitle",
			args:     api.ListIndividualsOptions{CollectionAgentTitle: "title"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_agent_title = $1`,
			wantArgs: []interface{}{"title"},
		}, {
			name:     "collectionTimeFrom",
			args:     api.ListIndividualsOptions{CollectionTimeFrom: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_time >= $1`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:     "collectionTimeTo",
			args:     api.ListIndividualsOptions{CollectionTimeTo: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND collection_time <= $1`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:     "communityId",
			args:     api.ListIndividualsOptions{CommunityID: "communityId"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND community_id = $1`,
			wantArgs: []interface{}{"communityId"},
		}, {
			name:     "countryId",
			args:     api.ListIndividualsOptions{CountryID: "countryId"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND country_id = $1`,
			wantArgs: []interface{}{"countryId"},
		}, {
			name:     "createdAtFrom",
			args:     api.ListIndividualsOptions{CreatedAtFrom: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND created_at >= $1`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:     "createdAtTo",
			args:     api.ListIndividualsOptions{CreatedAtTo: &someDate},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND created_at <= $1`,
			wantArgs: []interface{}{&someDate},
		}, {
			name:     "displacementStatus (single)",
			args:     api.ListIndividualsOptions{DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus](enumTypes.DisplacementStatusIDP, enumTypes.DisplacementStatusRefugee)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND displacement_status IN ($1,$2)`,
			wantArgs: []interface{}{"idp", "refugee"},
		}, {
			name:     "displacementStatus (all)",
			args:     api.ListIndividualsOptions{DisplacementStatuses: enumTypes.AllDisplacementStatuses()},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND displacement_status IN ($1,$2,$3,$4,$5,$6,$7)`,
			wantArgs: []interface{}{"asylum_seeker", "host_community", "idp", "non_displaced", "other", "refugee", "returnee"},
		}, {
			name:    "displacementStatus (none)",
			args:    api.ListIndividualsOptions{DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus]()},
			wantSql: `SELECT * FROM individual_registrations WHERE deleted_at IS NULL`,
		}, {
			name:     "email",
			args:     api.ListIndividualsOptions{Email: "email"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (email_1 = $1 OR email_2 = $1 OR email_3 = $1)`,
			wantArgs: []interface{}{"email"},
		}, {
			name:     "freeField1",
			args:     api.ListIndividualsOptions{FreeField1: "freeField1"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND free_field_1 = $1`,
			wantArgs: []interface{}{"freeField1"},
		}, {
			name:     "freeField2",
			args:     api.ListIndividualsOptions{FreeField2: "freeField2"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND free_field_2 = $1`,
			wantArgs: []interface{}{"freeField2"},
		}, {
			name:     "freeField3",
			args:     api.ListIndividualsOptions{FreeField3: "freeField3"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND free_field_3 = $1`,
			wantArgs: []interface{}{"freeField3"},
		}, {
			name:     "freeField4",
			args:     api.ListIndividualsOptions{FreeField4: "freeField4"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND free_field_4 = $1`,
			wantArgs: []interface{}{"freeField4"},
		}, {
			name:     "freeField5",
			args:     api.ListIndividualsOptions{FreeField5: "freeField5"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND free_field_5 = $1`,
			wantArgs: []interface{}{"freeField5"},
		}, {
			name:     "full name",
			args:     api.ListIndividualsOptions{FullName: "John"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (full_name ILIKE $1 OR preferred_name ILIKE $2 OR first_name ILIKE $3 OR middle_name ILIKE $4 OR last_name ILIKE $5 OR native_name ILIKE $6)`,
			wantArgs: []interface{}{"%John%", "%John%", "%John%", "%John%", "%John%", "%John%"},
		}, {
			name:     "mothers name",
			args:     api.ListIndividualsOptions{MothersName: "Jane Doe"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND mothers_name ILIKE $1`,
			wantArgs: []interface{}{"%Jane Doe%"},
		}, {
			name:     "sex (single)",
			args:     api.ListIndividualsOptions{Sexes: containers.NewSet[enumTypes.Sex](enumTypes.SexMale, enumTypes.SexFemale)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND sex IN ($1,$2)`,
			wantArgs: []interface{}{"female", "male"},
		}, {
			name:     "sex (all)",
			args:     api.ListIndividualsOptions{Sexes: enumTypes.AllSexes()},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND sex IN ($1,$2,$3,$4)`,
			wantArgs: []interface{}{"female", "male", "other", "prefers_not_to_say"},
		}, {
			name:     "hasCognitiveDisability",
			args:     api.ListIndividualsOptions{HasCognitiveDisability: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_cognitive_disability = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasCognitiveDisability (false)",
			args:     api.ListIndividualsOptions{HasCognitiveDisability: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_cognitive_disability = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasCommunicationDisability",
			args:     api.ListIndividualsOptions{HasCommunicationDisability: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_communication_disability = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasCommunicationDisability (false)",
			args:     api.ListIndividualsOptions{HasCommunicationDisability: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_communication_disability = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasConsentedToRGPD",
			args:     api.ListIndividualsOptions{HasConsentedToRGPD: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_consented_to_rgpd = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasConsentedToRGPD (false)",
			args:     api.ListIndividualsOptions{HasConsentedToRGPD: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_consented_to_rgpd = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasConsentedToReferral",
			args:     api.ListIndividualsOptions{HasConsentedToReferral: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_consented_to_referral = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasConsentedToReferral (false)",
			args:     api.ListIndividualsOptions{HasConsentedToReferral: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_consented_to_referral = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasDisability",
			args:     api.ListIndividualsOptions{HasDisability: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND inactive = false AND has_disability = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasDisability (false)",
			args:     api.ListIndividualsOptions{HasDisability: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND inactive = false AND has_disability = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasHearingDisability",
			args:     api.ListIndividualsOptions{HasHearingDisability: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_hearing_disability = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasHearingDisability (false)",
			args:     api.ListIndividualsOptions{HasHearingDisability: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_hearing_disability = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasMobilityDisability",
			args:     api.ListIndividualsOptions{HasMobilityDisability: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_mobility_disability = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasMobilityDisability (false)",
			args:     api.ListIndividualsOptions{HasMobilityDisability: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_mobility_disability = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasSelfCareDisability",
			args:     api.ListIndividualsOptions{HasSelfCareDisability: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_selfcare_disability = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasSelfCareDisability (false)",
			args:     api.ListIndividualsOptions{HasSelfCareDisability: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_selfcare_disability = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hasVisionDisability",
			args:     api.ListIndividualsOptions{HasVisionDisability: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_vision_disability = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "hasVisionDisability (false)",
			args:     api.ListIndividualsOptions{HasVisionDisability: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND has_vision_disability = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "hearingDisabilityLevel",
			args:     api.ListIndividualsOptions{HearingDisabilityLevel: enumTypes.DisabilityLevelMild},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND hearing_disability_level = $1`,
			wantArgs: []interface{}{"mild"},
		}, {
			name:     "householdId",
			args:     api.ListIndividualsOptions{HouseholdID: "household-id"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND household_id = $1`,
			wantArgs: []interface{}{"household-id"},
		}, {
			name:     "id (single)",
			args:     api.ListIndividualsOptions{IDs: containers.NewStringSet("1")},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND id IN ($1)`,
			wantArgs: []interface{}{"1"},
		}, {
			name:     "id (multiple)",
			args:     api.ListIndividualsOptions{IDs: containers.NewStringSet("1", "2")},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND id IN ($1,$2)`,
			wantArgs: []interface{}{"1", "2"},
		}, {
			name:     "identificationNumber",
			args:     api.ListIndividualsOptions{IdentificationNumber: "123456789"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (identification_number_1 = $1 OR identification_number_2 = $1 OR identification_number_3 = $1)`,
			wantArgs: []interface{}{"123456789"},
		}, {
			name:     "engagementContext",
			args:     api.ListIndividualsOptions{EngagementContext: containers.NewSet[enumTypes.EngagementContext](enumTypes.EngagementContextInOffice)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND engagement_context IN ($1)`,
			wantArgs: []interface{}{"inOffice"},
		}, {
			name:     "internalID",
			args:     api.ListIndividualsOptions{InternalID: "internal-id"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND internal_id = $1`,
			wantArgs: []interface{}{"internal-id"},
		}, {
			name:     "isHeadOfCommunity",
			args:     api.ListIndividualsOptions{IsHeadOfCommunity: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_head_of_community = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "isHeadOfCommunity (false)",
			args:     api.ListIndividualsOptions{IsHeadOfCommunity: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_head_of_community = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "isHeadOfHousehold",
			args:     api.ListIndividualsOptions{IsHeadOfHousehold: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_head_of_household = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "isHeadOfHousehold (false)",
			args:     api.ListIndividualsOptions{IsHeadOfHousehold: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_head_of_household = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "isFemaleHeadedHousehold",
			args:     api.ListIndividualsOptions{IsFemaleHeadedHousehold: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_female_headed_household = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "isFemaleHeadedHousehold (false)",
			args:     api.ListIndividualsOptions{IsFemaleHeadedHousehold: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_female_headed_household = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "isMinorHeadedHousehold",
			args:     api.ListIndividualsOptions{IsMinorHeadedHousehold: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_minor_headed_household = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "isMinorHeadedHousehold (false)",
			args:     api.ListIndividualsOptions{IsMinorHeadedHousehold: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_minor_headed_household = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "isMinor",
			args:     api.ListIndividualsOptions{IsMinor: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_minor = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "isMinor (false)",
			args:     api.ListIndividualsOptions{IsMinor: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND is_minor = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "mobilityDisabilityLevel",
			args:     api.ListIndividualsOptions{MobilityDisabilityLevel: enumTypes.DisabilityLevelMild},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND mobility_disability_level = $1`,
			wantArgs: []interface{}{"mild"},
		}, {
			name:     "nationality",
			args:     api.ListIndividualsOptions{Nationality: "nationality"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (nationality_1 = $1 OR nationality_2 = $1)`,
			wantArgs: []interface{}{"nationality"},
		}, {
			name:     "phone number",
			args:     api.ListIndividualsOptions{PhoneNumber: "1234567890"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (normalized_phone_number_1 ILIKE $1 OR normalized_phone_number_2 ILIKE $1 OR normalized_phone_number_3 ILIKE $1)`,
			wantArgs: []interface{}{"%1234567890%"},
		}, {
			name:     "preferredContactMehtod",
			args:     api.ListIndividualsOptions{PreferredContactMethod: "contactMethod"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND preferred_contact_method = $1`,
			wantArgs: []interface{}{"contactMethod"},
		}, {
			name:     "preferredCommunicationLanguage",
			args:     api.ListIndividualsOptions{PreferredCommunicationLanguage: "language"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND preferred_communication_language = $1`,
			wantArgs: []interface{}{"language"},
		}, {
			name:     "prefersToRemainAnonymous",
			args:     api.ListIndividualsOptions{PrefersToRemainAnonymous: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND prefers_to_remain_anonymous = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "prefersToRemainAnonymous (false)",
			args:     api.ListIndividualsOptions{PrefersToRemainAnonymous: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND prefers_to_remain_anonymous = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "presentsProtectionConcerns",
			args:     api.ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(true)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND presents_protection_concerns = $1`,
			wantArgs: []interface{}{true},
		}, {
			name:     "presentsProtectionConcerns (false)",
			args:     api.ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(false)},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND presents_protection_concerns = $1`,
			wantArgs: []interface{}{false},
		}, {
			name:     "selfCareDisabilityLevel",
			args:     api.ListIndividualsOptions{SelfCareDisabilityLevel: enumTypes.DisabilityLevelMild},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND selfcare_disability_level = $1`,
			wantArgs: []interface{}{"mild"},
		}, {
			name:     "spokenLanguage",
			args:     api.ListIndividualsOptions{SpokenLanguage: "language"},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND (spoken_language_1 = $1 OR spoken_language_2 = $1 OR spoken_language_3 = $1)`,
			wantArgs: []interface{}{"language"},
		}, {
			name:     "updatedAtFrom",
			args:     api.ListIndividualsOptions{UpdatedAtFrom: pointers.Time(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND updated_at >= $1`,
			wantArgs: []interface{}{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		}, {
			name:     "updatedAtTo",
			args:     api.ListIndividualsOptions{UpdatedAtTo: pointers.Time(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND updated_at <= $1`,
			wantArgs: []interface{}{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		}, {
			name:     "visionDisabilityLevel",
			args:     api.ListIndividualsOptions{VisionDisabilityLevel: enumTypes.DisabilityLevelMild},
			wantSql:  `SELECT * FROM individual_registrations WHERE deleted_at IS NULL AND vision_disability_level = $1`,
			wantArgs: []interface{}{"mild"},
		}, {
			name:    "skip",
			args:    api.ListIndividualsOptions{Skip: 10},
			wantSql: `SELECT * FROM individual_registrations WHERE deleted_at IS NULL OFFSET 10`,
		}, {
			name:    "take",
			args:    api.ListIndividualsOptions{Take: 10},
			wantSql: `SELECT * FROM individual_registrations WHERE deleted_at IS NULL LIMIT 10`,
		}, {
			name:    "skip and take",
			args:    api.ListIndividualsOptions{Skip: 10, Take: 10},
			wantSql: `SELECT * FROM individual_registrations WHERE deleted_at IS NULL OFFSET 10 LIMIT 10`,
		}, {
			name:    "sort (asc)",
			args:    api.ListIndividualsOptions{Sort: api.SortTerms{{Field: "id", Direction: api.SortDirectionAscending}}},
			wantSql: `SELECT * FROM individual_registrations WHERE deleted_at IS NULL ORDER BY id ASC`,
		}, {
			name:    "sort (desc)",
			args:    api.ListIndividualsOptions{Sort: api.SortTerms{{Field: "id", Direction: api.SortDirectionDescending}}},
			wantSql: `SELECT * FROM individual_registrations WHERE deleted_at IS NULL ORDER BY id DESC`,
		}, {
			name: "sort (multiple)",
			args: api.ListIndividualsOptions{Sort: api.SortTerms{
				{Field: "id", Direction: api.SortDirectionAscending},
				{Field: "full_name", Direction: api.SortDirectionDescending},
			}},
			wantSql: `SELECT * FROM individual_registrations WHERE deleted_at IS NULL ORDER BY id ASC, full_name DESC`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, args := newGetAllIndividualsSQLQuery("postgres", tt.args).build()
			assert.Equal(t, tt.wantSql, sql)
			assert.Equal(t, tt.wantArgs, args)
		})
	}
}
