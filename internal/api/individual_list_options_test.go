package api

import (
	"net/url"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/locales"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func TestNewIndividualListFromURLValues(t *testing.T) {
	locales.LoadTranslations()
	locales.Init()

	tests := []struct {
		name    string
		args    url.Values
		want    ListIndividualsOptions
		wantErr bool
	}{
		{
			name: "inactive",
			args: url.Values{"inactive": []string{"true"}},
			want: ListIndividualsOptions{Inactive: pointers.Bool(true)},
		}, {
			name: "address",
			args: url.Values{"address": []string{"address"}},
			want: ListIndividualsOptions{Address: "address"},
		}, {
			name: "ageFrom",
			args: url.Values{"age_from": []string{"10"}},
			want: ListIndividualsOptions{AgeFrom: pointers.Int(10)},
		}, {
			name:    "ageFrom (invalid)",
			args:    url.Values{"age_from": []string{"abc"}},
			wantErr: true,
		}, {
			name: "ageTo",
			args: url.Values{"age_to": []string{"10"}},
			want: ListIndividualsOptions{AgeTo: pointers.Int(10)},
		}, {
			name:    "ageTo (invalid)",
			args:    url.Values{"age_to": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "birthDateFrom",
			args: url.Values{"birth_date_from": []string{"2009-01-01"}},
			want: ListIndividualsOptions{BirthDateFrom: pointers.Time(time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC))},
		}, {
			name:    "birthDateFrom (invalid)",
			args:    url.Values{"birth_date_from": []string{"abc"}},
			wantErr: true,
		}, {
			name: "birthDateTo",
			args: url.Values{"birth_date_to": []string{"2009-01-01"}},
			want: ListIndividualsOptions{BirthDateTo: pointers.Time(time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC))},
		}, {
			name:    "birthDateTo (invalid)",
			args:    url.Values{"birth_date_to": []string{"abc"}},
			wantErr: true,
		}, {
			name: "cognitiveDisabilityLevel",
			args: url.Values{"cognitive_disability_level": []string{"mild"}},
			want: ListIndividualsOptions{CognitiveDisabilityLevel: enumTypes.DisabilityLevelMild},
		}, {
			name:    "cognitiveDisabilityLevel (invalid)",
			args:    url.Values{"cognitive_disability_level": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "collectionAdministrativeArea1",
			args: url.Values{"collection_administrative_area_1": []string{"area1"}},
			want: ListIndividualsOptions{CollectionAdministrativeArea1: "area1"},
		}, {
			name: "collectionAdministrativeArea2",
			args: url.Values{"collection_administrative_area_2": []string{"area2"}},
			want: ListIndividualsOptions{CollectionAdministrativeArea2: "area2"},
		}, {
			name: "collectionAdministrativeArea3",
			args: url.Values{"collection_administrative_area_3": []string{"area3"}},
			want: ListIndividualsOptions{CollectionAdministrativeArea3: "area3"},
		}, {
			name: "collectionOffice",
			args: url.Values{"collection_office": []string{"collection office"}},
			want: ListIndividualsOptions{CollectionOffice: "collection office"},
		}, {
			name: "collectionAgentName",
			args: url.Values{"collection_agent_name": []string{"amy"}},
			want: ListIndividualsOptions{CollectionAgentName: "amy"},
		}, {
			name: "collectionAgentTile",
			args: url.Values{"collection_agent_title": []string{"admin"}},
			want: ListIndividualsOptions{CollectionAgentTitle: "admin"},
		}, {
			name: "collectionTimeFrom",
			args: url.Values{"collection_time_from": []string{"2020-01-01"}},
			want: ListIndividualsOptions{CollectionTimeFrom: pointers.Time(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))},
		}, {
			name:    "collectionTimeFrom (invalid)",
			args:    url.Values{"collection_time_from": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "collectionTimeTo",
			args: url.Values{"collection_time_to": []string{"2020-01-01"}},
			want: ListIndividualsOptions{CollectionTimeTo: pointers.Time(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))},
		}, {
			name:    "collectionTimeTo (invalid)",
			args:    url.Values{"collection_time_to": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "communityId",
			args: url.Values{"community_id": []string{"communityId"}},
			want: ListIndividualsOptions{CommunityID: "communityId"},
		}, {
			name: "createdAtFrom",
			args: url.Values{"created_at_from": []string{"2020-01-01T10:30:05Z"}},
			want: ListIndividualsOptions{CreatedAtFrom: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
		}, {
			name:    "createdAtFrom (invalid)",
			args:    url.Values{"created_at_from": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "createdAtTo",
			args: url.Values{"created_at_to": []string{"2020-01-01T10:30:05Z"}},
			want: ListIndividualsOptions{CreatedAtTo: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
		}, {
			name:    "createdAtTo (invalid)",
			args:    url.Values{"created_at_to": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "displacementStatus",
			args: url.Values{"displacement_status": []string{"idp"}},
			want: ListIndividualsOptions{DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus](enumTypes.DisplacementStatusIDP)},
		}, {
			name: "displacementStatus (multiple)",
			args: url.Values{"displacement_status": []string{"idp", "refugee"}},
			want: ListIndividualsOptions{DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus](enumTypes.DisplacementStatusIDP, enumTypes.DisplacementStatusRefugee)},
		}, {
			name:    "displacementStatus (invalid)",
			args:    url.Values{"displacement_status": []string{"invalidd"}},
			wantErr: true,
		}, {
			name: "email",
			args: url.Values{"email": []string{"email"}},
			want: ListIndividualsOptions{Email: "email"},
		}, {
			name: "freeField1",
			args: url.Values{"free_field_1": []string{"freeField1"}},
			want: ListIndividualsOptions{FreeField1: "freeField1"},
		}, {
			name: "freeField2",
			args: url.Values{"free_field_2": []string{"freeField2"}},
			want: ListIndividualsOptions{FreeField2: "freeField2"},
		}, {
			name: "freeField3",
			args: url.Values{"free_field_3": []string{"freeField3"}},
			want: ListIndividualsOptions{FreeField3: "freeField3"},
		}, {
			name: "freeField4",
			args: url.Values{"free_field_4": []string{"freeField4"}},
			want: ListIndividualsOptions{FreeField4: "freeField4"},
		}, {
			name: "freeField5",
			args: url.Values{"free_field_5": []string{"freeField5"}},
			want: ListIndividualsOptions{FreeField5: "freeField5"},
		}, {
			name: "fullName",
			args: url.Values{"full_name": []string{"name"}},
			want: ListIndividualsOptions{FullName: "name"},
		}, {
			name: "sex",
			args: url.Values{"sex": []string{"female"}},
			want: ListIndividualsOptions{Sexes: containers.NewSet[enumTypes.Sex](enumTypes.SexFemale)},
		}, {
			name: "sex (multiple)",
			args: url.Values{"sex": []string{"female", "male"}},
			want: ListIndividualsOptions{Sexes: containers.NewSet[enumTypes.Sex](enumTypes.SexFemale, enumTypes.SexMale)},
		}, {
			name:    "sex (invalid)",
			args:    url.Values{"sex": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasCognitiveDisability",
			args: url.Values{"has_cognitive_disability": []string{"true"}},
			want: ListIndividualsOptions{HasCognitiveDisability: pointers.Bool(true)},
		}, {
			name: "hasCognitiveDisability (false)",
			args: url.Values{"has_cognitive_disability": []string{"false"}},
			want: ListIndividualsOptions{HasCognitiveDisability: pointers.Bool(false)},
		}, {
			name:    "hasCognitiveDisability (invalid)",
			args:    url.Values{"has_cognitive_disability": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasCommunicationDisability",
			args: url.Values{"has_communication_disability": []string{"true"}},
			want: ListIndividualsOptions{HasCommunicationDisability: pointers.Bool(true)},
		}, {
			name: "hasCommunicationDisability (false)",
			args: url.Values{"has_communication_disability": []string{"false"}},
			want: ListIndividualsOptions{HasCommunicationDisability: pointers.Bool(false)},
		}, {
			name: "hasConsentedToRgpd",
			args: url.Values{"has_consented_to_rgpd": []string{"true"}},
			want: ListIndividualsOptions{HasConsentedToRGPD: pointers.Bool(true)},
		}, {
			name: "hasConsentedToRgpd (false)",
			args: url.Values{"has_consented_to_rgpd": []string{"false"}},
			want: ListIndividualsOptions{HasConsentedToRGPD: pointers.Bool(false)},
		}, {
			name:    "hasConsentedToRgpd (invalid)",
			args:    url.Values{"has_consented_to_rgpd": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasConsentedToReferral",
			args: url.Values{"has_consented_to_referral": []string{"true"}},
			want: ListIndividualsOptions{HasConsentedToReferral: pointers.Bool(true)},
		}, {
			name: "hasConsentedToReferral (false)",
			args: url.Values{"has_consented_to_referral": []string{"false"}},
			want: ListIndividualsOptions{HasConsentedToReferral: pointers.Bool(false)},
		}, {
			name:    "hasConsentedToReferral (invalid)",
			args:    url.Values{"has_consented_to_referral": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasDisability",
			args: url.Values{"has_disability": []string{"true"}},
			want: ListIndividualsOptions{HasDisability: pointers.Bool(true)},
		}, {
			name: "hasDisability (false)",
			args: url.Values{"has_disability": []string{"false"}},
			want: ListIndividualsOptions{HasDisability: pointers.Bool(false)},
		}, {
			name:    "hasDisability (invalid)",
			args:    url.Values{"has_disability": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasHearingDisability",
			args: url.Values{"has_hearing_disability": []string{"true"}},
			want: ListIndividualsOptions{HasHearingDisability: pointers.Bool(true)},
		}, {
			name: "hasHearingDisability (false)",
			args: url.Values{"has_hearing_disability": []string{"false"}},
			want: ListIndividualsOptions{HasHearingDisability: pointers.Bool(false)},
		}, {
			name:    "hasHearingDisability (invalid)",
			args:    url.Values{"has_hearing_disability": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasMobilityDisability",
			args: url.Values{"has_mobility_disability": []string{"true"}},
			want: ListIndividualsOptions{HasMobilityDisability: pointers.Bool(true)},
		}, {
			name: "hasMobilityDisability (false)",
			args: url.Values{"has_mobility_disability": []string{"false"}},
			want: ListIndividualsOptions{HasMobilityDisability: pointers.Bool(false)},
		}, {
			name:    "hasMobilityDisability (invalid)",
			args:    url.Values{"has_mobility_disability": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasSelfCareDisability",
			args: url.Values{"has_selfcare_disability": []string{"true"}},
			want: ListIndividualsOptions{HasSelfCareDisability: pointers.Bool(true)},
		}, {
			name: "hasSelfCareDisability (false)",
			args: url.Values{"has_selfcare_disability": []string{"false"}},
			want: ListIndividualsOptions{HasSelfCareDisability: pointers.Bool(false)},
		}, {
			name:    "hasSelfCareDisability (invalid)",
			args:    url.Values{"has_selfcare_disability": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hasVisionDisability",
			args: url.Values{"has_vision_disability": []string{"true"}},
			want: ListIndividualsOptions{HasVisionDisability: pointers.Bool(true)},
		}, {
			name: "hasVisionDisability (false)",
			args: url.Values{"has_vision_disability": []string{"false"}},
			want: ListIndividualsOptions{HasVisionDisability: pointers.Bool(false)},
		}, {
			name:    "hasVisionDisability (invalid)",
			args:    url.Values{"has_vision_disability": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "hearingDisabilityLevel",
			args: url.Values{"hearing_disability_level": []string{"mild"}},
			want: ListIndividualsOptions{HearingDisabilityLevel: enumTypes.DisabilityLevelMild},
		}, {
			name:    "hearingDisabilityLevel (invalid)",
			args:    url.Values{"hearing_disability_level": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "householdId",
			args: url.Values{"household_id": []string{"household-id"}},
			want: ListIndividualsOptions{HouseholdID: "household-id"},
		}, {
			name: "id",
			args: url.Values{"id": []string{"id"}},
			want: ListIndividualsOptions{IDs: containers.NewStringSet("id")},
		}, {
			name: "id (multiple)",
			args: url.Values{"id": []string{"id1", "id2"}},
			want: ListIndividualsOptions{IDs: containers.NewStringSet("id1", "id2")},
		}, {
			name: "identificationNumber",
			args: url.Values{"identification_number": []string{"identification-number"}},
			want: ListIndividualsOptions{IdentificationNumber: "identification-number"},
		}, {
			name: "engagementContext",
			args: url.Values{"engagement_context": []string{"inOffice"}},
			want: ListIndividualsOptions{EngagementContext: containers.NewSet[enumTypes.EngagementContext](enumTypes.EngagementContextInOffice)},
		}, {
			name: "internalId",
			args: url.Values{"internal_id": []string{"internal-id"}},
			want: ListIndividualsOptions{InternalID: "internal-id"},
		}, {
			name: "isHeadOfCommunity",
			args: url.Values{"is_head_of_community": []string{"true"}},
			want: ListIndividualsOptions{IsHeadOfCommunity: pointers.Bool(true)},
		}, {
			name: "isHeadOfCommunity (false)",
			args: url.Values{"is_head_of_community": []string{"false"}},
			want: ListIndividualsOptions{IsHeadOfCommunity: pointers.Bool(false)},
		}, {
			name:    "isHeadOfCommunity (invalid)",
			args:    url.Values{"is_head_of_community": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "isHeadOfHousehold",
			args: url.Values{"is_head_of_household": []string{"true"}},
			want: ListIndividualsOptions{IsHeadOfHousehold: pointers.Bool(true)},
		}, {
			name: "isHeadOfHousehold (false)",
			args: url.Values{"is_head_of_household": []string{"false"}},
			want: ListIndividualsOptions{IsHeadOfHousehold: pointers.Bool(false)},
		}, {
			name:    "isHeadOfHousehold (invalid)",
			args:    url.Values{"is_head_of_household": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "isFemaleHeadedHousehold",
			args: url.Values{"is_female_headed_household": []string{"true"}},
			want: ListIndividualsOptions{IsFemaleHeadedHousehold: pointers.Bool(true)},
		}, {
			name: "isFemaleHeadedHousehold (false)",
			args: url.Values{"is_female_headed_household": []string{"false"}},
			want: ListIndividualsOptions{IsFemaleHeadedHousehold: pointers.Bool(false)},
		}, {
			name:    "isFemaleHeadedHousehold (invalid)",
			args:    url.Values{"is_female_headed_household": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "isMinorHeadedHousehold",
			args: url.Values{"is_minor_headed_household": []string{"true"}},
			want: ListIndividualsOptions{IsMinorHeadedHousehold: pointers.Bool(true)},
		}, {
			name: "isMinorHeadedHousehold (false)",
			args: url.Values{"is_minor_headed_household": []string{"false"}},
			want: ListIndividualsOptions{IsMinorHeadedHousehold: pointers.Bool(false)},
		}, {
			name:    "isMinorHeadedHousehold (invalid)",
			args:    url.Values{"is_minor_headed_household": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "isMinor",
			args: url.Values{"is_minor": []string{"true"}},
			want: ListIndividualsOptions{IsMinor: pointers.Bool(true)},
		}, {
			name: "isMinor (false)",
			args: url.Values{"is_minor": []string{"false"}},
			want: ListIndividualsOptions{IsMinor: pointers.Bool(false)},
		}, {
			name:    "isMinor (invalid)",
			args:    url.Values{"is_minor": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "mobilityDisabilityLevel",
			args: url.Values{"mobility_disability_level": []string{"mild"}},
			want: ListIndividualsOptions{MobilityDisabilityLevel: enumTypes.DisabilityLevelMild},
		}, {
			name:    "mobilityDisabilityLevel (invalid)",
			args:    url.Values{"mobility_disability_level": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "nationality",
			args: url.Values{"nationality": []string{"nationality"}},
			want: ListIndividualsOptions{Nationality: "nationality"},
		}, {
			name: "phoneNumber",
			args: url.Values{"phone_number": []string{"phone-number"}},
			want: ListIndividualsOptions{PhoneNumber: "phone-number"},
		}, {
			name: "preferredContactMethod",
			args: url.Values{"preferred_contact_method": []string{"preferred-contact-method"}},
			want: ListIndividualsOptions{PreferredContactMethod: "preferred-contact-method"},
		}, {
			name: "preferredCommunicationLanguage",
			args: url.Values{"preferred_communication_language": []string{"preferred-communication-language"}},
			want: ListIndividualsOptions{PreferredCommunicationLanguage: "preferred-communication-language"},
		}, {
			name: "prefersToRemainAnonymous",
			args: url.Values{"prefers_to_remain_anonymous": []string{"true"}},
			want: ListIndividualsOptions{PrefersToRemainAnonymous: pointers.Bool(true)},
		}, {
			name: "prefersToRemainAnonymous (false)",
			args: url.Values{"prefers_to_remain_anonymous": []string{"false"}},
			want: ListIndividualsOptions{PrefersToRemainAnonymous: pointers.Bool(false)},
		}, {
			name:    "prefersToRemainAnonymous (invalid)",
			args:    url.Values{"prefers_to_remain_anonymous": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "presentsProtectionConcerns",
			args: url.Values{"presents_protection_concerns": []string{"true"}},
			want: ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(true)},
		}, {
			name: "presentsProtectionConcerns (false)",
			args: url.Values{"presents_protection_concerns": []string{"false"}},
			want: ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(false)},
		}, {
			name:    "presentsProtectionConcerns (invalid)",
			args:    url.Values{"presents_protection_concerns": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "selfCareDisabilityLevel",
			args: url.Values{"selfcare_disability_level": []string{"mild"}},
			want: ListIndividualsOptions{SelfCareDisabilityLevel: enumTypes.DisabilityLevelMild},
		}, {
			name:    "selfCareDisabilityLevel (invalid)",
			args:    url.Values{"selfcare_disability_level": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "serviceType",
			args: url.Values{"service_type": []string{"service-type"}},
			want: ListIndividualsOptions{ServiceType: "service-type"},
		}, {
			name: "service",
			args: url.Values{"service": []string{"service"}},
			want: ListIndividualsOptions{Service: "service"},
		}, {
			name: "serviceSubService",
			args: url.Values{"service_sub_service": []string{"sub-service"}},
			want: ListIndividualsOptions{ServiceSubService: "sub-service"},
		}, {
			name: "serviceLocation",
			args: url.Values{"service_location": []string{"location"}},
			want: ListIndividualsOptions{ServiceLocation: "location"},
		}, {
			name: "serviceDonor",
			args: url.Values{"service_donor": []string{"donor"}},
			want: ListIndividualsOptions{ServiceDonor: "donor"},
		}, {
			name: "serviceProjectName",
			args: url.Values{"service_project_name": []string{"project-name"}},
			want: ListIndividualsOptions{ServiceProjectName: "project-name"},
		}, {
			name: "serviceAgentName",
			args: url.Values{"service_agent_name": []string{"agent-name"}},
			want: ListIndividualsOptions{ServiceAgentName: "agent-name"},
		}, {
			name: "spokenLanguage",
			args: url.Values{"spoken_language": []string{"spoken-language"}},
			want: ListIndividualsOptions{SpokenLanguage: "spoken-language"},
		}, {
			name: "updatedAtFrom",
			args: url.Values{"updated_at_from": []string{"2020-01-01T10:30:05Z"}},
			want: ListIndividualsOptions{UpdatedAtFrom: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
		}, {
			name:    "updatedAtFrom (invalid)",
			args:    url.Values{"updated_at_from": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "updatedAtTo",
			args: url.Values{"updated_at_to": []string{"2020-01-01T10:30:05Z"}},
			want: ListIndividualsOptions{UpdatedAtTo: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
		}, {
			name:    "updatedAtTo (invalid)",
			args:    url.Values{"updated_at_to": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "skip",
			args: url.Values{"skip": []string{"1"}},
			want: ListIndividualsOptions{Skip: 1},
		}, {
			name:    "skip (invalid)",
			args:    url.Values{"skip": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "sort (asc)",
			args: url.Values{"sort": []string{"full_name"}},
			want: ListIndividualsOptions{Sort: SortTerms{{Field: "full_name", Direction: SortDirectionAscending}}},
		}, {
			name: "sort (desc)",
			args: url.Values{"sort": []string{"-age"}},
			want: ListIndividualsOptions{Sort: SortTerms{{Field: "age", Direction: SortDirectionDescending}}},
		}, {
			name: "sort (multiple)",
			args: url.Values{"sort": []string{"full_name,-age"}},
			want: ListIndividualsOptions{Sort: SortTerms{{Field: "full_name", Direction: SortDirectionAscending}, {Field: "age", Direction: SortDirectionDescending}}},
		}, {
			name:    "sort (duplicate)",
			args:    url.Values{"sort": []string{"full_name,-full_name"}},
			wantErr: true,
		}, {
			name:    "sort (invalid)",
			args:    url.Values{"sort": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "sort (empty)",
			args: url.Values{"sort": []string{""}},
			want: ListIndividualsOptions{},
		}, {
			name: "take",
			args: url.Values{"take": []string{"1"}},
			want: ListIndividualsOptions{Take: 1},
		}, {
			name:    "take (invalid)",
			args:    url.Values{"take": []string{"invalid"}},
			wantErr: true,
		}, {
			name: "visionDisabilityLevel",
			args: url.Values{"vision_disability_level": []string{"mild"}},
			want: ListIndividualsOptions{VisionDisabilityLevel: enumTypes.DisabilityLevelMild},
		}, {
			name:    "visionDisabilityLevel (invalid)",
			args:    url.Values{"vision_disability_level": []string{"invalid"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ret ListIndividualsOptions
			err := NewIndividualListFromURLValues(tt.args, &ret)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tt.want, ret)
		})
	}

}

func TestListIndividualsOptions_QueryParams(t *testing.T) {

	const countryId = "usa"
	tests := []struct {
		name string
		o    ListIndividualsOptions
		want string
	}{
		{
			name: "inactive",
			o:    ListIndividualsOptions{CountryID: countryId, Inactive: pointers.Bool(true)},
			want: "/countries/usa/participants?inactive=true",
		}, {
			name: "address",
			o:    ListIndividualsOptions{CountryID: countryId, Address: "address"},
			want: "/countries/usa/participants?address=address",
		}, {
			name: "ageFrom",
			o:    ListIndividualsOptions{CountryID: countryId, AgeFrom: pointers.Int(1)},
			want: "/countries/usa/participants?age_from=1",
		}, {
			name: "ageTo",
			o:    ListIndividualsOptions{CountryID: countryId, AgeTo: pointers.Int(1)},
			want: "/countries/usa/participants?age_to=1",
		}, {
			name: "birthDateFrom",
			o:    ListIndividualsOptions{CountryID: countryId, BirthDateFrom: pointers.Time(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))},
			want: "/countries/usa/participants?birth_date_from=2000-01-01",
		}, {
			name: "birthDateTo",
			o:    ListIndividualsOptions{CountryID: countryId, BirthDateTo: pointers.Time(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))},
			want: "/countries/usa/participants?birth_date_to=2000-01-01",
		}, {
			name: "cognitiveDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, CognitiveDisabilityLevel: enumTypes.DisabilityLevelNone},
			want: "/countries/usa/participants?cognitive_disability_level=none",
		}, {
			name: "collectionAdministrativeArea1",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAdministrativeArea1: "collectionAdministrativeArea1"},
			want: "/countries/usa/participants?collection_administrative_area_1=collectionAdministrativeArea1",
		}, {
			name: "collectionAdministrativeArea2",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAdministrativeArea2: "collectionAdministrativeArea2"},
			want: "/countries/usa/participants?collection_administrative_area_2=collectionAdministrativeArea2",
		}, {
			name: "collectionAdministrativeArea3",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAdministrativeArea3: "collectionAdministrativeArea3"},
			want: "/countries/usa/participants?collection_administrative_area_3=collectionAdministrativeArea3",
		}, {
			name: "collectionOffice",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionOffice: "collectionOffice"},
			want: "/countries/usa/participants?collection_office=collectionOffice",
		}, {
			name: "collectionAgentName",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAgentName: "collectionAgentName"},
			want: "/countries/usa/participants?collection_agent_name=collectionAgentName",
		}, {
			name: "collectionAgentTitle",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAgentTitle: "collectionAgentTitle"},
			want: "/countries/usa/participants?collection_agent_title=collectionAgentTitle",
		}, {
			name: "collectionTimeFrom",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionTimeFrom: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/participants?collection_time_from=2000-01-01",
		}, {
			name: "collectionTimeTo",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionTimeTo: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/participants?collection_time_to=2000-01-01",
		}, {
			name: "communityID",
			o:    ListIndividualsOptions{CountryID: countryId, CommunityID: "communityID"},
			want: "/countries/usa/participants?community_id=communityID",
		}, {
			name: "createdAtFrom",
			o:    ListIndividualsOptions{CountryID: countryId, CreatedAtFrom: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/participants?created_at_from=2000-01-01T10%3A30%3A05Z",
		}, {
			name: "createdAtTo",
			o:    ListIndividualsOptions{CountryID: countryId, CreatedAtTo: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/participants?created_at_to=2000-01-01T10%3A30%3A05Z",
		}, {
			name: "displacement status",
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus](enumTypes.DisplacementStatusIDP)},
			want: "/countries/usa/participants?displacement_status=idp",
		}, {
			name: "displacement status (multiple)",
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: containers.NewSet[enumTypes.DisplacementStatus](enumTypes.DisplacementStatusIDP, enumTypes.DisplacementStatusRefugee)},
			want: "/countries/usa/participants?displacement_status=idp&displacement_status=refugee",
		}, {
			name: "email",
			o:    ListIndividualsOptions{CountryID: countryId, Email: "email"},
			want: "/countries/usa/participants?email=email",
		}, {
			name: "free field 1",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField1: "freeField1"},
			want: "/countries/usa/participants?free_field_1=freeField1",
		}, {
			name: "free field 2",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField2: "freeField2"},
			want: "/countries/usa/participants?free_field_2=freeField2",
		}, {
			name: "free field 3",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField3: "freeField3"},
			want: "/countries/usa/participants?free_field_3=freeField3",
		}, {
			name: "free field 4",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField4: "freeField4"},
			want: "/countries/usa/participants?free_field_4=freeField4",
		}, {
			name: "free field 5",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField5: "freeField5"},
			want: "/countries/usa/participants?free_field_5=freeField5",
		}, {
			name: "fullName",
			o:    ListIndividualsOptions{CountryID: countryId, FullName: "fullName"},
			want: "/countries/usa/participants?full_name=fullName",
		}, {
			name: "sex",
			o:    ListIndividualsOptions{CountryID: countryId, Sexes: containers.NewSet[enumTypes.Sex](enumTypes.SexMale)},
			want: "/countries/usa/participants?sex=male",
		}, {
			name: "sex (multiple)",
			o:    ListIndividualsOptions{CountryID: countryId, Sexes: containers.NewSet[enumTypes.Sex](enumTypes.SexMale, enumTypes.SexFemale)},
			want: "/countries/usa/participants?sex=female&sex=male",
		}, {
			name: "hasCognitiveDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasCognitiveDisability: pointers.Bool(true)},
			want: "/countries/usa/participants?has_cognitive_disability=true",
		}, {
			name: "hasCognitiveDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasCognitiveDisability: pointers.Bool(false)},
			want: "/countries/usa/participants?has_cognitive_disability=false",
		}, {
			name: "hasCommunicationDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasCommunicationDisability: pointers.Bool(true)},
			want: "/countries/usa/participants?has_communication_disability=true",
		}, {
			name: "hasCommunicationDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasCommunicationDisability: pointers.Bool(false)},
			want: "/countries/usa/participants?has_communication_disability=false",
		}, {
			name: "hasConstentedToRGPD",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToRGPD: pointers.Bool(true)},
			want: "/countries/usa/participants?has_consented_to_rgpd=true",
		}, {
			name: "hasConstentedToRGPD (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToRGPD: pointers.Bool(false)},
			want: "/countries/usa/participants?has_consented_to_rgpd=false",
		}, {
			name: "hasConsentedToReferral",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToReferral: pointers.Bool(true)},
			want: "/countries/usa/participants?has_consented_to_referral=true",
		}, {
			name: "hasConsentedToReferral (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToReferral: pointers.Bool(false)},
			want: "/countries/usa/participants?has_consented_to_referral=false",
		}, {
			name: "hasHearingDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasHearingDisability: pointers.Bool(true)},
			want: "/countries/usa/participants?has_hearing_disability=true",
		}, {
			name: "hasHearingDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasHearingDisability: pointers.Bool(false)},
			want: "/countries/usa/participants?has_hearing_disability=false",
		}, {
			name: "hasMobilityDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasMobilityDisability: pointers.Bool(true)},
			want: "/countries/usa/participants?has_mobility_disability=true",
		}, {
			name: "hasMobilityDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasMobilityDisability: pointers.Bool(false)},
			want: "/countries/usa/participants?has_mobility_disability=false",
		}, {
			name: "hasSelfCareDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasSelfCareDisability: pointers.Bool(true)},
			want: "/countries/usa/participants?has_selfcare_disability=true",
		}, {
			name: "hasSelfCareDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasSelfCareDisability: pointers.Bool(false)},
			want: "/countries/usa/participants?has_selfcare_disability=false",
		}, {
			name: "hasVisionDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasVisionDisability: pointers.Bool(true)},
			want: "/countries/usa/participants?has_vision_disability=true",
		}, {
			name: "hasVisionDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasVisionDisability: pointers.Bool(false)},
			want: "/countries/usa/participants?has_vision_disability=false",
		}, {
			name: "hearingDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, HearingDisabilityLevel: enumTypes.DisabilityLevelMild},
			want: "/countries/usa/participants?hearing_disability_level=mild",
		}, {
			name: "householdID",
			o:    ListIndividualsOptions{CountryID: countryId, HouseholdID: "householdId"},
			want: "/countries/usa/participants?household_id=householdId",
		}, {
			name: "ids",
			o:    ListIndividualsOptions{CountryID: countryId, IDs: containers.NewStringSet("id1", "id2")},
			want: "/countries/usa/participants?id=id1&id=id2",
		}, {
			name: "identificationNumber",
			o:    ListIndividualsOptions{CountryID: countryId, IdentificationNumber: "identificationNumber"},
			want: "/countries/usa/participants?identification_number=identificationNumber",
		}, {
			name: "engagementContext",
			o:    ListIndividualsOptions{CountryID: countryId, EngagementContext: containers.NewSet[enumTypes.EngagementContext](enumTypes.EngagementContextInOffice)},
			want: "/countries/usa/participants?engagement_context=inOffice",
		}, {
			name: "internalID",
			o:    ListIndividualsOptions{CountryID: countryId, InternalID: "internalID"},
			want: "/countries/usa/participants?internal_id=internalID",
		}, {
			name: "isHeadOfCommunity",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfCommunity: pointers.Bool(true)},
			want: "/countries/usa/participants?is_head_of_community=true",
		}, {
			name: "isHeadOfCommunity (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfCommunity: pointers.Bool(false)},
			want: "/countries/usa/participants?is_head_of_community=false",
		}, {
			name: "isHeadOfHousehold",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfHousehold: pointers.Bool(true)},
			want: "/countries/usa/participants?is_head_of_household=true",
		}, {
			name: "isHeadOfHousehold (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfHousehold: pointers.Bool(false)},
			want: "/countries/usa/participants?is_head_of_household=false",
		}, {
			name: "isFemaleHeadedHousehold",
			o:    ListIndividualsOptions{CountryID: countryId, IsFemaleHeadedHousehold: pointers.Bool(true)},
			want: "/countries/usa/participants?is_female_headed_household=true",
		}, {
			name: "isFemaleHeadedHousehold (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsFemaleHeadedHousehold: pointers.Bool(false)},
			want: "/countries/usa/participants?is_female_headed_household=false",
		}, {
			name: "isMinorHeadedHousehold",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinorHeadedHousehold: pointers.Bool(true)},
			want: "/countries/usa/participants?is_minor_headed_household=true",
		}, {
			name: "isMinorHeadedHousehold (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinorHeadedHousehold: pointers.Bool(false)},
			want: "/countries/usa/participants?is_minor_headed_household=false",
		}, {
			name: "isMinor",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinor: pointers.Bool(true)},
			want: "/countries/usa/participants?is_minor=true",
		}, {
			name: "isMinor (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinor: pointers.Bool(false)},
			want: "/countries/usa/participants?is_minor=false",
		}, {
			name: "mobilityDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, MobilityDisabilityLevel: enumTypes.DisabilityLevelMild},
			want: "/countries/usa/participants?mobility_disability_level=mild",
		}, {
			name: "nationality",
			o:    ListIndividualsOptions{CountryID: countryId, Nationality: "nationality"},
			want: "/countries/usa/participants?nationality=nationality",
		}, {
			name: "phoneNumber",
			o:    ListIndividualsOptions{CountryID: countryId, PhoneNumber: "phoneNumber"},
			want: "/countries/usa/participants?phone_number=phoneNumber",
		}, {
			name: "preferredContactMethod",
			o:    ListIndividualsOptions{CountryID: countryId, PreferredContactMethod: "phone"},
			want: "/countries/usa/participants?preferred_contact_method=phone",
		}, {
			name: "preferredCommunicationLanguage",
			o:    ListIndividualsOptions{CountryID: countryId, PreferredCommunicationLanguage: "en"},
			want: "/countries/usa/participants?preferred_communication_language=en",
		}, {
			name: "prefersToRemainAnonymous",
			o:    ListIndividualsOptions{CountryID: countryId, PrefersToRemainAnonymous: pointers.Bool(true)},
			want: "/countries/usa/participants?prefers_to_remain_anonymous=true",
		}, {
			name: "prefersToRemainAnonymous (false)",
			o:    ListIndividualsOptions{CountryID: countryId, PrefersToRemainAnonymous: pointers.Bool(false)},
			want: "/countries/usa/participants?prefers_to_remain_anonymous=false",
		}, {
			name: "presentsProtectionConcerns",
			o:    ListIndividualsOptions{CountryID: countryId, PresentsProtectionConcerns: pointers.Bool(true)},
			want: "/countries/usa/participants?presents_protection_concerns=true",
		}, {
			name: "presentsProtectionConcerns (false)",
			o:    ListIndividualsOptions{CountryID: countryId, PresentsProtectionConcerns: pointers.Bool(false)},
			want: "/countries/usa/participants?presents_protection_concerns=false",
		}, {
			name: "selfCareDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, SelfCareDisabilityLevel: enumTypes.DisabilityLevelMild},
			want: "/countries/usa/participants?selfcare_disability_level=mild",
		}, {
			name: "serviceType",
			o:    ListIndividualsOptions{CountryID: countryId, ServiceType: "serviceType"},
			want: "/countries/usa/participants?service_type=serviceType",
		}, {
			name: "service",
			o:    ListIndividualsOptions{CountryID: countryId, Service: "service"},
			want: "/countries/usa/participants?service=service",
		}, {
			name: "subService",
			o:    ListIndividualsOptions{CountryID: countryId, ServiceSubService: "subService"},
			want: "/countries/usa/participants?service_sub_service=subService",
		}, {
			name: "serviceLocation",
			o:    ListIndividualsOptions{CountryID: countryId, ServiceLocation: "serviceLocation"},
			want: "/countries/usa/participants?service_location=serviceLocation",
		}, {
			name: "serviceDonor",
			o:    ListIndividualsOptions{CountryID: countryId, ServiceDonor: "serviceDonor"},
			want: "/countries/usa/participants?service_donor=serviceDonor",
		}, {
			name: "serviceProjectName",
			o:    ListIndividualsOptions{CountryID: countryId, ServiceProjectName: "serviceProjectName"},
			want: "/countries/usa/participants?service_project_name=serviceProjectName",
		}, {
			name: "serviceAgentName",
			o:    ListIndividualsOptions{CountryID: countryId, ServiceAgentName: "serviceAgentName"},
			want: "/countries/usa/participants?service_agent_name=serviceAgentName",
		}, {
			name: "spokenLanguage",
			o:    ListIndividualsOptions{CountryID: countryId, SpokenLanguage: "en"},
			want: "/countries/usa/participants?spoken_language=en",
		}, {
			name: "updatedAtFrom",
			o:    ListIndividualsOptions{CountryID: countryId, UpdatedAtFrom: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/participants?updated_at_from=2020-01-01T10%3A30%3A05Z",
		}, {
			name: "updatedAtTo",
			o:    ListIndividualsOptions{CountryID: countryId, UpdatedAtTo: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/participants?updated_at_to=2020-01-01T10%3A30%3A05Z",
		}, {
			name: "skip",
			o:    ListIndividualsOptions{CountryID: countryId, Skip: 1},
			want: "/countries/usa/participants?skip=1",
		}, {
			name: "take",
			o:    ListIndividualsOptions{CountryID: countryId, Take: 1},
			want: "/countries/usa/participants?take=1",
		}, {
			name: "visionDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, VisionDisabilityLevel: enumTypes.DisabilityLevelMild},
			want: "/countries/usa/participants?vision_disability_level=mild",
		}, {
			name: "empty",
			o:    ListIndividualsOptions{CountryID: countryId},
			want: "/countries/usa/participants",
		}, {
			name: "sort (asc)",
			o:    ListIndividualsOptions{CountryID: countryId, Sort: SortTerms{{Field: "column", Direction: SortDirectionAscending}}},
			want: "/countries/usa/participants?sort=column",
		}, {
			name: "sort (desc)",
			o:    ListIndividualsOptions{CountryID: countryId, Sort: SortTerms{{Field: "column", Direction: SortDirectionDescending}}},
			want: "/countries/usa/participants?sort=-column",
		}, {
			name: "sort (multiple)",
			o: ListIndividualsOptions{CountryID: countryId, Sort: SortTerms{
				{Field: "column", Direction: SortDirectionAscending},
				{Field: "column2", Direction: SortDirectionDescending}},
			},
			want: "/countries/usa/participants?sort=column%2C-column2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.o.QueryParams()
			want := tt.want
			assert.Equal(t, want, got)
		})
	}
}
