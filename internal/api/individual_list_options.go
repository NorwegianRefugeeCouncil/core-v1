package api

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/containers"
)

type ListIndividualsOptions struct {
	Inactive                       *bool
	Address                        string
	AgeFrom                        *int
	AgeTo                          *int
	BirthDateFrom                  *time.Time
	BirthDateTo                    *time.Time
	CognitiveDisabilityLevel       DisabilityLevel
	CollectionAdministrativeArea1  string
	CollectionAdministrativeArea2  string
	CollectionAdministrativeArea3  string
	CollectionOffice               string
	CollectionAgentName            string
	CollectionAgentTitle           string
	CollectionTimeFrom             *time.Time
	CollectionTimeTo               *time.Time
	CommunityID                    string
	CountryID                      string
	CreatedAtFrom                  *time.Time
	CreatedAtTo                    *time.Time
	DisplacementStatuses           containers.Set[DisplacementStatus]
	Email                          string
	FreeField1                     string
	FreeField2                     string
	FreeField3                     string
	FreeField4                     string
	FreeField5                     string
	FullName                       string
	Sexes                          containers.Set[Sex]
	HasCognitiveDisability         *bool
	HasCommunicationDisability     *bool
	HasConsentedToRGPD             *bool
	HasConsentedToReferral         *bool
	HasHearingDisability           *bool
	HasMobilityDisability          *bool
	HasSelfCareDisability          *bool
	HasVisionDisability            *bool
	HearingDisabilityLevel         DisabilityLevel
	HouseholdID                    string
	IDs                            containers.StringSet
	IdentificationNumber           string
	EngagementContext              containers.Set[EngagementContext]
	InternalID                     string
	IsHeadOfCommunity              *bool
	IsHeadOfHousehold              *bool
	IsFemaleHeadedHousehold        *bool
	IsMinorHeadedHousehold         *bool
	IsMinor                        *bool
	MobilityDisabilityLevel        DisabilityLevel
	Nationality                    string
	PhoneNumber                    string
	PreferredContactMethod         string
	PreferredCommunicationLanguage string
	PrefersToRemainAnonymous       *bool
	PresentsProtectionConcerns     *bool
	SelfCareDisabilityLevel        DisabilityLevel
	SpokenLanguage                 string
	UpdatedAtFrom                  *time.Time
	UpdatedAtTo                    *time.Time
	Skip                           int
	Take                           int
	Sort                           SortTerms
	ServiceCC1                     ServiceCC
	ServiceRequestedDate1From      *time.Time
	ServiceRequestedDate1To        *time.Time
	ServiceDeliveredDate1From      *time.Time
	ServiceDeliveredDate1To        *time.Time
	ServiceCC2                     ServiceCC
	ServiceRequestedDate2From      *time.Time
	ServiceRequestedDate2To        *time.Time
	ServiceDeliveredDate2From      *time.Time
	ServiceDeliveredDate2To        *time.Time
	ServiceCC3                     ServiceCC
	ServiceRequestedDate3From      *time.Time
	ServiceRequestedDate3To        *time.Time
	ServiceDeliveredDate3From      *time.Time
	ServiceDeliveredDate3To        *time.Time
	ServiceCC4                     ServiceCC
	ServiceRequestedDate4From      *time.Time
	ServiceRequestedDate4To        *time.Time
	ServiceDeliveredDate4From      *time.Time
	ServiceDeliveredDate4To        *time.Time
	ServiceCC5                     ServiceCC
	ServiceRequestedDate5From      *time.Time
	ServiceRequestedDate5To        *time.Time
	ServiceDeliveredDate5From      *time.Time
	ServiceDeliveredDate5To        *time.Time
	ServiceCC6                     ServiceCC
	ServiceRequestedDate6From      *time.Time
	ServiceRequestedDate6To        *time.Time
	ServiceDeliveredDate6From      *time.Time
	ServiceDeliveredDate6To        *time.Time
	ServiceCC7                     ServiceCC
	ServiceRequestedDate7From      *time.Time
	ServiceRequestedDate7To        *time.Time
	ServiceDeliveredDate7From      *time.Time
	ServiceDeliveredDate7To        *time.Time
	VisionDisabilityLevel          DisabilityLevel
}

type SortDirection string

const (
	SortDirectionNone       SortDirection = ""
	SortDirectionAscending  SortDirection = "ascending"
	SortDirectionDescending SortDirection = "descending"
)

type SortTerm struct {
	Field     string
	Direction SortDirection
}

type SortTerms []SortTerm

func (s SortTerms) MarshalQuery() string {
	var query string
	for i, term := range s {
		if i > 0 {
			query += ","
		}
		if term.Direction == SortDirectionAscending {
			query += term.Field
		} else {
			query += "-" + term.Field
		}
	}
	return query
}

func (s *SortTerms) UnmarshalQuery(query string) error {
	terms := make(SortTerms, 0)
	parts := strings.Split(query, ",")
	seenColumns := containers.NewStringSet()
	for _, term := range parts {
		column, direction, err := s.parseTerm(term)
		if err != nil {
			return err
		}
		if !sortableColumns.Contains(column) {
			return fmt.Errorf("invalid sort column: %s", column)
		}
		if seenColumns.Contains(column) {
			return fmt.Errorf("duplicate sort column: %s", column)
		}
		seenColumns.Add(column)
		terms = append(terms, SortTerm{
			Field:     column,
			Direction: direction,
		})
	}
	*s = terms
	return nil
}

func (s *SortTerms) parseTerm(term string) (string, SortDirection, error) {
	if term == "" {
		return "", "", fmt.Errorf("empty term")
	}
	var direction = SortDirectionAscending
	if term[0] == '-' {
		direction = SortDirectionDescending
		term = term[1:]
	}
	if len(term) == 0 {
		return "", "", fmt.Errorf("invalid sort term: %s", term)
	}
	return term, direction, nil
}

func (o ListIndividualsOptions) IsInactiveSelected() bool {
	return *o.Inactive
}

func (o ListIndividualsOptions) IsMinorSelected() bool {
	return o.IsMinor != nil && *o.IsMinor
}

func (o ListIndividualsOptions) IsMinorHeadedHouseholdSelected() bool {
	return o.IsMinorHeadedHousehold != nil && *o.IsMinorHeadedHousehold
}

func (o ListIndividualsOptions) IsFemaleHeadedHouseholdSelected() bool {
	return o.IsFemaleHeadedHousehold != nil && *o.IsFemaleHeadedHousehold
}

func (o ListIndividualsOptions) IsNotMinorSelected() bool {
	return o.IsMinor != nil && !*o.IsMinor
}

func (o ListIndividualsOptions) IsPresentsProtectionConcernsSelected() bool {
	return o.PresentsProtectionConcerns != nil && *o.PresentsProtectionConcerns
}

func (o ListIndividualsOptions) IsNotPresentsProtectionConcernsSelected() bool {
	return o.PresentsProtectionConcerns != nil && !*o.PresentsProtectionConcerns
}

func (o ListIndividualsOptions) IsHasVisionDisabilitySelected() bool {
	return o.HasVisionDisability != nil && *o.HasVisionDisability
}

func (o ListIndividualsOptions) IsNotHasVisionDisabilitySelected() bool {
	return o.HasVisionDisability != nil && !*o.HasVisionDisability
}

func (o ListIndividualsOptions) IsHasCognitiveDisabilitySelected() bool {
	return o.HasCognitiveDisability != nil && *o.HasCognitiveDisability
}

func (o ListIndividualsOptions) IsNotHasCognitiveDisabilitySelected() bool {
	return o.HasCognitiveDisability != nil && !*o.HasCognitiveDisability
}

func (o ListIndividualsOptions) IsHasCommunicationDisabilitySelected() bool {
	return o.HasCommunicationDisability != nil && *o.HasCommunicationDisability
}

func (o ListIndividualsOptions) IsNotHasCommunicationDisabilitySelected() bool {
	return o.HasCommunicationDisability != nil && !*o.HasCommunicationDisability
}

func (o ListIndividualsOptions) IsHasHearingDisabilitySelected() bool {
	return o.HasHearingDisability != nil && *o.HasHearingDisability
}

func (o ListIndividualsOptions) IsNotHasHearingDisabilitySelected() bool {
	return o.HasHearingDisability != nil && !*o.HasHearingDisability
}

func (o ListIndividualsOptions) IsHasMobilityDisabilitySelected() bool {
	return o.HasMobilityDisability != nil && *o.HasMobilityDisability
}

func (o ListIndividualsOptions) IsNotHasMobilityDisabilitySelected() bool {
	return o.HasMobilityDisability != nil && !*o.HasMobilityDisability
}

func (o ListIndividualsOptions) IsHasSelfCareDisabilitySelected() bool {
	return o.HasSelfCareDisability != nil && *o.HasSelfCareDisability
}

func (o ListIndividualsOptions) IsNotHasSelfCareDisabilitySelected() bool {
	return o.HasSelfCareDisability != nil && !*o.HasSelfCareDisability
}

func (o ListIndividualsOptions) NextPage() ListIndividualsOptions {
	ret := o
	ret.Skip += ret.Take
	return ret
}

func (o ListIndividualsOptions) PreviousPage() ListIndividualsOptions {
	ret := o
	ret.Skip -= ret.Take
	if ret.Skip < 0 {
		ret.Skip = 0
	}
	return ret
}

func (o ListIndividualsOptions) FirstPage() ListIndividualsOptions {
	ret := o
	ret.Skip = 0
	return ret
}

func (o ListIndividualsOptions) WithTake(take int) ListIndividualsOptions {
	ret := o
	ret.Take = take
	return ret
}

func (o ListIndividualsOptions) WithSort(field string, direction string) ListIndividualsOptions {
	o.Sort = SortTerms{
		{
			Field:     field,
			Direction: SortDirection(direction),
		},
	}
	return o
}

func (o ListIndividualsOptions) GetSortDirection(field string) SortDirection {
	for _, term := range o.Sort {
		if term.Field == field {
			return term.Direction
		}
	}
	return SortDirectionNone
}

func (o ListIndividualsOptions) QueryParams() string {
	params := newListIndividualsOptionsEncoder(o, time.Now()).encode()
	u := url.URL{Path: "/countries/" + o.CountryID + "/participants"}
	u.RawQuery = params.Encode()
	return u.String()
}

func NewIndividualListFromURLValues(values url.Values, into *ListIndividualsOptions) error {
	parser := listIndividualsOptionsDecoder{
		out:    into,
		values: values,
		now:    time.Now(),
	}
	return parser.parse()
}
