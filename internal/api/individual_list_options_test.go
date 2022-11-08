package api

import (
	"html/template"
	"net/url"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func TestNewIndividualListFromURLValues(t *testing.T) {

	tests := []struct {
		name    string
		args    url.Values
		want    ListIndividualsOptions
		wantErr bool
	}{
		{
			name: "empty",
			args: url.Values{},
			want: ListIndividualsOptions{},
		}, {
			name: constants.FormParamsGetIndividualsSkip,
			args: url.Values{constants.FormParamsGetIndividualsSkip: []string{"1"}},
			want: ListIndividualsOptions{Skip: 1},
		}, {
			name:    "invalid skip",
			args:    url.Values{constants.FormParamsGetIndividualsSkip: []string{"abc"}},
			wantErr: true,
		}, {
			name:    "negative skip",
			args:    url.Values{constants.FormParamsGetIndividualsSkip: []string{"-10"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsTake,
			args: url.Values{constants.FormParamsGetIndividualsTake: []string{"1"}},
			want: ListIndividualsOptions{Take: 1},
		}, {
			name:    "invalid take",
			args:    url.Values{constants.FormParamsGetIndividualsTake: []string{"abc"}},
			wantErr: true,
		}, {
			name:    "negative take",
			args:    url.Values{constants.FormParamsGetIndividualsTake: []string{"-10"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsFullName,
			args: url.Values{constants.FormParamsGetIndividualsFullName: []string{"fullName"}},
			want: ListIndividualsOptions{FullName: "fullName"},
		}, {
			name: constants.FormParamsGetIndividualsAddress,
			args: url.Values{constants.FormParamsGetIndividualsAddress: []string{"address"}},
			want: ListIndividualsOptions{Address: "address"},
		}, {
			name: constants.FormParamsGetIndividualsEmail,
			args: url.Values{constants.FormParamsGetIndividualsEmail: []string{"email"}},
			want: ListIndividualsOptions{Email: "email"},
		}, {
			name: constants.FormParamsGetIndividualsPhoneNumber,
			args: url.Values{constants.FormParamsGetIndividualsPhoneNumber: []string{"phoneNumber"}},
			want: ListIndividualsOptions{PhoneNumber: "phoneNumber"},
		}, {
			name: constants.FormParamsGetIndividualsCountryID,
			args: url.Values{constants.FormParamsGetIndividualsCountryID: []string{"countryID"}},
			want: ListIndividualsOptions{CountryID: "countryID"},
		}, {
			name: constants.FormParamsGetIndividualsBirthDateFrom,
			args: url.Values{constants.FormParamsGetIndividualsBirthDateFrom: []string{"2009-01-01"}},
			want: ListIndividualsOptions{BirthDateFrom: pointers.Time(time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC))},
		}, {
			name:    "invalid birth date from",
			args:    url.Values{constants.FormParamsGetIndividualsBirthDateFrom: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsBirthDateTo,
			args: url.Values{constants.FormParamsGetIndividualsBirthDateTo: []string{"2009-01-01"}},
			want: ListIndividualsOptions{BirthDateTo: pointers.Time(time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC))},
		}, {
			name:    "invalid birth date to",
			args:    url.Values{constants.FormParamsGetIndividualsBirthDateTo: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsAgeFrom,
			args: url.Values{constants.FormParamsGetIndividualsAgeFrom: []string{"1"}},
			want: ListIndividualsOptions{AgeFrom: pointers.Int(1)},
		}, {
			name:    "invalid age from",
			args:    url.Values{constants.FormParamsGetIndividualsAgeFrom: []string{"abc"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsAgeTo,
			args: url.Values{constants.FormParamsGetIndividualsAgeTo: []string{"1"}},
			want: ListIndividualsOptions{AgeTo: pointers.Int(1)},
		}, {
			name:    "invalid age to",
			args:    url.Values{constants.FormParamsGetIndividualsAgeTo: []string{"abc"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsIsMinor,
			args: url.Values{constants.FormParamsGetIndividualsIsMinor: []string{"true"}},
			want: ListIndividualsOptions{IsMinor: pointers.Bool(true)},
		}, {
			name: constants.FormParamsGetIndividualsIsMinor,
			args: url.Values{constants.FormParamsGetIndividualsIsMinor: []string{"false"}},
			want: ListIndividualsOptions{IsMinor: pointers.Bool(false)},
		}, {
			name:    "invalid isMinor",
			args:    url.Values{constants.FormParamsGetIndividualsIsMinor: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsPresentsProtectionConcerns,
			args: url.Values{constants.FormParamsGetIndividualsPresentsProtectionConcerns: []string{"true"}},
			want: ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(true)},
		}, {
			name: constants.FormParamsGetIndividualsPresentsProtectionConcerns,
			args: url.Values{constants.FormParamsGetIndividualsPresentsProtectionConcerns: []string{"false"}},
			want: ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(false)},
		}, {
			name:    "invalid presentsProtectionConcerns",
			args:    url.Values{constants.FormParamsGetIndividualsPresentsProtectionConcerns: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsID,
			args: url.Values{constants.FormParamsGetIndividualsID: []string{"1"}},
			want: ListIndividualsOptions{IDs: containers.NewStringSet("1")},
		}, {
			name: "multiple ids",
			args: url.Values{constants.FormParamsGetIndividualsID: []string{"1", "2", "2", "3"}},
			want: ListIndividualsOptions{IDs: containers.NewStringSet("1", "2", "3")},
		}, {
			name: constants.FormParamsGetIndividualsDisplacementStatus,
			args: url.Values{constants.FormParamsGetIndividualsDisplacementStatus: []string{"idp", "idp", "refugee"}},
			want: ListIndividualsOptions{DisplacementStatuses: containers.NewSet[DisplacementStatus](DisplacementStatusIDP, DisplacementStatusRefugee)},
		}, {
			name: constants.FormParamsGetIndividualsFreeField1,
			args: url.Values{constants.FormParamsGetIndividualsFreeField1: []string{"freeField1"}},
			want: ListIndividualsOptions{FreeField1: "freeField1"},
		}, {
			name: constants.FormParamsGetIndividualsFreeField2,
			args: url.Values{constants.FormParamsGetIndividualsFreeField2: []string{"freeField2"}},
			want: ListIndividualsOptions{FreeField2: "freeField2"},
		}, {
			name: constants.FormParamsGetIndividualsFreeField3,
			args: url.Values{constants.FormParamsGetIndividualsFreeField3: []string{"freeField3"}},
			want: ListIndividualsOptions{FreeField3: "freeField3"},
		}, {
			name: constants.FormParamsGetIndividualsFreeField4,
			args: url.Values{constants.FormParamsGetIndividualsFreeField4: []string{"freeField4"}},
			want: ListIndividualsOptions{FreeField4: "freeField4"},
		}, {
			name: constants.FormParamsGetIndividualsFreeField5,
			args: url.Values{constants.FormParamsGetIndividualsFreeField5: []string{"freeField5"}},
			want: ListIndividualsOptions{FreeField5: "freeField5"},
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
			name: "address",
			o:    ListIndividualsOptions{CountryID: countryId, Address: "address"},
			want: "/countries/usa/individuals?address=address",
		}, {
			name: "ageFrom",
			o:    ListIndividualsOptions{CountryID: countryId, AgeFrom: pointers.Int(1)},
			want: "/countries/usa/individuals?age_from=1",
		}, {
			name: "ageTo",
			o:    ListIndividualsOptions{CountryID: countryId, AgeTo: pointers.Int(1)},
			want: "/countries/usa/individuals?age_to=1",
		}, {
			name: "birthDateFrom",
			o:    ListIndividualsOptions{CountryID: countryId, BirthDateFrom: pointers.Time(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))},
			want: "/countries/usa/individuals?birth_date_from=2000-01-01",
		}, {
			name: "birthDateTo",
			o:    ListIndividualsOptions{CountryID: countryId, BirthDateTo: pointers.Time(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))},
			want: "/countries/usa/individuals?birth_date_to=2000-01-01",
		}, {
			name: "cognitiveDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, CognitiveDisabilityLevel: DisabilityLevelNone},
			want: "/countries/usa/individuals?cognitive_disability_level=none",
		}, {
			name: "collectionAdministrativeArea1",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAdministrativeArea1: "collectionAdministrativeArea1"},
			want: "/countries/usa/individuals?collection_administrative_area_1=collectionAdministrativeArea1",
		}, {
			name: "collectionAdministrativeArea2",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAdministrativeArea2: "collectionAdministrativeArea2"},
			want: "/countries/usa/individuals?collection_administrative_area_2=collectionAdministrativeArea2",
		}, {
			name: "collectionAdministrativeArea3",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAdministrativeArea3: "collectionAdministrativeArea3"},
			want: "/countries/usa/individuals?collection_administrative_area_3=collectionAdministrativeArea3",
		}, {
			name: "collectionAgentName",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAgentName: "collectionAgentName"},
			want: "/countries/usa/individuals?collection_agent_name=collectionAgentName",
		}, {
			name: "collectionAgentTitle",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionAgentTitle: "collectionAgentTitle"},
			want: "/countries/usa/individuals?collection_agent_title=collectionAgentTitle",
		}, {
			name: "collectionTimeFrom",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionTimeFrom: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/individuals?collection_time_from=2000-01-01T10%3A30%3A05Z",
		}, {
			name: "collectionTimeTo",
			o:    ListIndividualsOptions{CountryID: countryId, CollectionTimeTo: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/individuals?collection_time_to=2000-01-01T10%3A30%3A05Z",
		}, {
			name: "communityID",
			o:    ListIndividualsOptions{CountryID: countryId, CommunityID: "communityID"},
			want: "/countries/usa/individuals?community_id=communityID",
		}, {
			name: "createdAtFrom",
			o:    ListIndividualsOptions{CountryID: countryId, CreatedAtFrom: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/individuals?created_at_from=2000-01-01T10%3A30%3A05Z",
		}, {
			name: "createdAtTo",
			o:    ListIndividualsOptions{CountryID: countryId, CreatedAtTo: pointers.Time(time.Date(2000, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/individuals?created_at_to=2000-01-01T10%3A30%3A05Z",
		}, {
			name: "displacement status",
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: containers.NewSet[DisplacementStatus](DisplacementStatusIDP)},
			want: "/countries/usa/individuals?displacement_status=idp",
		}, {
			name: "displacement status (multiple)",
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: containers.NewSet[DisplacementStatus](DisplacementStatusIDP, DisplacementStatusRefugee)},
			want: "/countries/usa/individuals?displacement_status=idp&displacement_status=refugee",
		}, {
			name: "email",
			o:    ListIndividualsOptions{CountryID: countryId, Email: "email"},
			want: "/countries/usa/individuals?email=email",
		}, {
			name: "free field 1",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField1: "freeField1"},
			want: "/countries/usa/individuals?free_field_1=freeField1",
		}, {
			name: "free field 2",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField2: "freeField2"},
			want: "/countries/usa/individuals?free_field_2=freeField2",
		}, {
			name: "free field 3",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField3: "freeField3"},
			want: "/countries/usa/individuals?free_field_3=freeField3",
		}, {
			name: "free field 4",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField4: "freeField4"},
			want: "/countries/usa/individuals?free_field_4=freeField4",
		}, {
			name: "free field 5",
			o:    ListIndividualsOptions{CountryID: countryId, FreeField5: "freeField5"},
			want: "/countries/usa/individuals?free_field_5=freeField5",
		}, {
			name: "fullName",
			o:    ListIndividualsOptions{CountryID: countryId, FullName: "fullName"},
			want: "/countries/usa/individuals?full_name=fullName",
		}, {
			name: "gender",
			o:    ListIndividualsOptions{CountryID: countryId, Genders: containers.NewSet[Gender]("male")},
			want: "/countries/usa/individuals?gender=male",
		}, {
			name: "gender (multiple)",
			o:    ListIndividualsOptions{CountryID: countryId, Genders: containers.NewSet[Gender](GenderMale, GenderFemale)},
			want: "/countries/usa/individuals?gender=female&gender=male",
		}, {
			name: "hasCognitiveDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasCognitiveDisability: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_cognitive_disability=true",
		}, {
			name: "hasCognitiveDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasCognitiveDisability: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_cognitive_disability=false",
		}, {
			name: "hasCommunicationDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasCommunicationDisability: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_communication_disability=true",
		}, {
			name: "hasCommunicationDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasCommunicationDisability: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_communication_disability=false",
		}, {
			name: "hasConstentedToRGPD",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToRGPD: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_consented_to_rgpd=true",
		}, {
			name: "hasConstentedToRGPD (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToRGPD: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_consented_to_rgpd=false",
		}, {
			name: "hasConsentedToReferral",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToReferral: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_consented_to_referral=true",
		}, {
			name: "hasConsentedToReferral (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasConsentedToReferral: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_consented_to_referral=false",
		}, {
			name: "hasHearingDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasHearingDisability: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_hearing_disability=true",
		}, {
			name: "hasHearingDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasHearingDisability: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_hearing_disability=false",
		}, {
			name: "hasMobilityDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasMobilityDisability: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_mobility_disability=true",
		}, {
			name: "hasMobilityDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasMobilityDisability: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_mobility_disability=false",
		}, {
			name: "hasSelfCareDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasSelfCareDisability: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_selfcare_disability=true",
		}, {
			name: "hasSelfCareDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasSelfCareDisability: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_selfcare_disability=false",
		}, {
			name: "hasVisionDisability",
			o:    ListIndividualsOptions{CountryID: countryId, HasVisionDisability: pointers.Bool(true)},
			want: "/countries/usa/individuals?has_vision_disability=true",
		}, {
			name: "hasVisionDisability (false)",
			o:    ListIndividualsOptions{CountryID: countryId, HasVisionDisability: pointers.Bool(false)},
			want: "/countries/usa/individuals?has_vision_disability=false",
		}, {
			name: "hearingDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, HearingDisabilityLevel: DisabilityLevelMild},
			want: "/countries/usa/individuals?hearing_disability_level=mild",
		}, {
			name: "householdID",
			o:    ListIndividualsOptions{CountryID: countryId, HouseholdID: "householdId"},
			want: "/countries/usa/individuals?household_id=householdId",
		}, {
			name: "ids",
			o:    ListIndividualsOptions{CountryID: countryId, IDs: containers.NewStringSet("id1", "id2")},
			want: "/countries/usa/individuals?id=id1&id=id2",
		}, {
			name: "identificationNumber",
			o:    ListIndividualsOptions{CountryID: countryId, IdentificationNumber: "identificationNumber"},
			want: "/countries/usa/individuals?identification_number=identificationNumber",
		}, {
			name: "identificationContext",
			o:    ListIndividualsOptions{CountryID: countryId, IdentificationContext: "identificationContext"},
			want: "/countries/usa/individuals?identification_context=identificationContext",
		}, {
			name: "internalID",
			o:    ListIndividualsOptions{CountryID: countryId, InternalID: "internalID"},
			want: "/countries/usa/individuals?internal_id=internalID",
		}, {
			name: "isHeadOfCommunity",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfCommunity: pointers.Bool(true)},
			want: "/countries/usa/individuals?is_head_of_community=true",
		}, {
			name: "isHeadOfCommunity (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfCommunity: pointers.Bool(false)},
			want: "/countries/usa/individuals?is_head_of_community=false",
		}, {
			name: "isHeadOfHousehold",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfHousehold: pointers.Bool(true)},
			want: "/countries/usa/individuals?is_head_of_household=true",
		}, {
			name: "isHeadOfHousehold (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsHeadOfHousehold: pointers.Bool(false)},
			want: "/countries/usa/individuals?is_head_of_household=false",
		}, {
			name: "isMinor",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinor: pointers.Bool(true)},
			want: "/countries/usa/individuals?is_minor=true",
		}, {
			name: "isMinor (false)",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinor: pointers.Bool(false)},
			want: "/countries/usa/individuals?is_minor=false",
		}, {
			name: "mobilityDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, MobilityDisabilityLevel: DisabilityLevelMild},
			want: "/countries/usa/individuals?mobility_disability_level=mild",
		}, {
			name: "nationality",
			o:    ListIndividualsOptions{CountryID: countryId, Nationality: "nationality"},
			want: "/countries/usa/individuals?nationality=nationality",
		}, {
			name: "phoneNumber",
			o:    ListIndividualsOptions{CountryID: countryId, PhoneNumber: "phoneNumber"},
			want: "/countries/usa/individuals?phone_number=phoneNumber",
		}, {
			name: "preferredContactMethod",
			o:    ListIndividualsOptions{CountryID: countryId, PreferredContactMethod: "phone"},
			want: "/countries/usa/individuals?preferred_contact_method=phone",
		}, {
			name: "preferredCommunicationLanguage",
			o:    ListIndividualsOptions{CountryID: countryId, PreferredCommunicationLanguage: "en"},
			want: "/countries/usa/individuals?preferred_communication_language=en",
		}, {
			name: "prefersToRemainAnonymous",
			o:    ListIndividualsOptions{CountryID: countryId, PrefersToRemainAnonymous: pointers.Bool(true)},
			want: "/countries/usa/individuals?prefers_to_remain_anonymous=true",
		}, {
			name: "prefersToRemainAnonymous (false)",
			o:    ListIndividualsOptions{CountryID: countryId, PrefersToRemainAnonymous: pointers.Bool(false)},
			want: "/countries/usa/individuals?prefers_to_remain_anonymous=false",
		}, {
			name: "presentsProtectionConcerns",
			o:    ListIndividualsOptions{CountryID: countryId, PresentsProtectionConcerns: pointers.Bool(true)},
			want: "/countries/usa/individuals?presents_protection_concerns=true",
		}, {
			name: "presentsProtectionConcerns (false)",
			o:    ListIndividualsOptions{CountryID: countryId, PresentsProtectionConcerns: pointers.Bool(false)},
			want: "/countries/usa/individuals?presents_protection_concerns=false",
		}, {
			name: "selfCareDisabilityLevel",
			o:    ListIndividualsOptions{CountryID: countryId, SelfCareDisabilityLevel: DisabilityLevelMild},
			want: "/countries/usa/individuals?selfcare_disability_level=mild",
		}, {
			name: "spokenLanguage",
			o:    ListIndividualsOptions{CountryID: countryId, SpokenLanguage: "en"},
			want: "/countries/usa/individuals?spoken_language=en",
		}, {
			name: "updatedAtFrom",
			o:    ListIndividualsOptions{CountryID: countryId, UpdatedAtFrom: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/individuals?updated_at_from=2020-01-01T10%3A30%3A05Z",
		}, {
			name: "updatedAtTo",
			o:    ListIndividualsOptions{CountryID: countryId, UpdatedAtTo: pointers.Time(time.Date(2020, 1, 1, 10, 30, 5, 0, time.UTC))},
			want: "/countries/usa/individuals?updated_at_to=2020-01-01T10%3A30%3A05Z",
		}, {
			name: "skip",
			o:    ListIndividualsOptions{CountryID: countryId, Skip: 1},
			want: "/countries/usa/individuals?skip=1",
		}, {
			name: "take",
			o:    ListIndividualsOptions{CountryID: countryId, Take: 1},
			want: "/countries/usa/individuals?take=1",
		}, {
			name: "empty",
			o:    ListIndividualsOptions{CountryID: countryId},
			want: "/countries/usa/individuals",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.o.QueryParams()
			want := template.HTML(tt.want)
			assert.Equal(t, want, got)
		})
	}
}
