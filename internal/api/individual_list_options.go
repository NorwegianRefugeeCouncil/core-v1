package api

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/api/enumTypes"

	"github.com/nrc-no/notcore/internal/containers"
)

type ListIndividualsOptions struct {
	Inactive                        *bool
	Address                         string
	AgeFrom                         *int
	AgeTo                           *int
	BirthDateFrom                   *time.Time
	BirthDateTo                     *time.Time
	CognitiveDisabilityLevel        enumTypes.DisabilityLevel
	CollectionAdministrativeArea1   string
	CollectionAdministrativeArea2   string
	CollectionAdministrativeArea3   string
	CollectionOffice                string
	CollectionAgentName             string
	CollectionAgentTitle            string
	CollectionTimeFrom              *time.Time
	CollectionTimeTo                *time.Time
	CommunityID                     string
	CountryID                       string
	CreatedAtFrom                   *time.Time
	CreatedAtTo                     *time.Time
	DisplacementStatuses            containers.Set[enumTypes.DisplacementStatus]
	Email                           string
	FreeField1                      string
	FreeField2                      string
	FreeField3                      string
	FreeField4                      string
	FreeField5                      string
	FullName                        string
	Sexes                           containers.Set[enumTypes.Sex]
	HasCognitiveDisability          *bool
	HasCommunicationDisability      *bool
	HasConsentedToRGPD              *bool
	HasConsentedToReferral          *bool
	HasDisability                   *bool
	HasHearingDisability            *bool
	HasMobilityDisability           *bool
	HasSelfCareDisability           *bool
	HasVisionDisability             *bool
	HearingDisabilityLevel          enumTypes.DisabilityLevel
	HouseholdID                     string
	IDs                             containers.StringSet
	IdentificationNumber            string
	EngagementContext               containers.Set[enumTypes.EngagementContext]
	InternalID                      string
	IsHeadOfCommunity               *bool
	IsHeadOfHousehold               *bool
	IsFemaleHeadedHousehold         *bool
	IsMinorHeadedHousehold          *bool
	IsMinor                         *bool
	IsChildAtRisk                   *bool
	IsWomanAtRisk                   *bool
	IsElderAtRisk                   *bool
	IsPregnant                      *bool
	IsLactating                     *bool
	IsSeparatedChild                *bool
	IsSingleParent                  *bool
	HasMedicalCondition             *bool
	NeedsLegalAndPhysicalProtection *bool
	MobilityDisabilityLevel         enumTypes.DisabilityLevel
	MothersName                     string
	Nationality                     string
	PhoneNumber                     string
	PreferredContactMethod          string
	PreferredCommunicationLanguage  string
	PrefersToRemainAnonymous        *bool
	PresentsProtectionConcerns      *bool
	PWDComments                     string
	VulnerabilityComments           string
	SelfCareDisabilityLevel         enumTypes.DisabilityLevel
	SpokenLanguage                  string
	UpdatedAtFrom                   *time.Time
	UpdatedAtTo                     *time.Time
	Skip                            int
	Take                            int
	Sort                            SortTerms
	ServiceCC                       containers.Set[enumTypes.ServiceCC]
	ServiceRequestedDateFrom        *time.Time
	ServiceRequestedDateTo          *time.Time
	ServiceDeliveredDateFrom        *time.Time
	ServiceDeliveredDateTo          *time.Time
	ServiceType                     string
	Service                         string
	ServiceSubService               string
	ServiceLocation                 string
	ServiceDonor                    string
	ServiceProjectName              string
	ServiceAgentName                string
	VisionDisabilityLevel           enumTypes.DisabilityLevel
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
	return o.Inactive != nil && *o.Inactive
}

func (o ListIndividualsOptions) IsNotInactiveSelected() bool {
	return o.Inactive != nil && !*o.Inactive
}

func (o ListIndividualsOptions) IsIsChildAtRiskSelected() bool {
	return o.IsChildAtRisk != nil && *o.IsChildAtRisk
}

func (o ListIndividualsOptions) IsNotIsChildAtRiskSelected() bool {
	return o.IsChildAtRisk != nil && !*o.IsChildAtRisk
}

func (o ListIndividualsOptions) IsIsWomanAtRiskSelected() bool {
	return o.IsWomanAtRisk != nil && *o.IsWomanAtRisk
}

func (o ListIndividualsOptions) IsNotIsWomanAtRiskSelected() bool {
	return o.IsWomanAtRisk != nil && !*o.IsWomanAtRisk
}

func (o ListIndividualsOptions) IsIsElderAtRiskSelected() bool {
	return o.IsElderAtRisk != nil && *o.IsElderAtRisk
}

func (o ListIndividualsOptions) IsNotIsElderAtRiskSelected() bool {
	return o.IsElderAtRisk != nil && !*o.IsElderAtRisk
}

func (o ListIndividualsOptions) IsIsSeparatedChildSelected() bool {
	return o.IsSeparatedChild != nil && *o.IsSeparatedChild
}

func (o ListIndividualsOptions) IsNotIsSeparatedChildSelected() bool {
	return o.IsSeparatedChild != nil && !*o.IsSeparatedChild
}

func (o ListIndividualsOptions) IsIsSingleParentSelected() bool {
	return o.IsSingleParent != nil && *o.IsSingleParent
}

func (o ListIndividualsOptions) IsNotIsSingleParentSelected() bool {
	return o.IsSingleParent != nil && !*o.IsSingleParent
}

func (o ListIndividualsOptions) IsIsPregnantSelected() bool {
	return o.IsPregnant != nil && *o.IsPregnant
}

func (o ListIndividualsOptions) IsNotIsPregnantSelected() bool {
	return o.IsPregnant != nil && !*o.IsPregnant
}

func (o ListIndividualsOptions) IsIsLactatingSelected() bool {
	return o.IsLactating != nil && *o.IsLactating
}

func (o ListIndividualsOptions) IsNotIsLactatingSelected() bool {
	return o.IsLactating != nil && !*o.IsLactating
}

func (o ListIndividualsOptions) IsHasMedicalConditionSelected() bool {
	return o.HasMedicalCondition != nil && *o.HasMedicalCondition
}

func (o ListIndividualsOptions) IsNotHasMedicalConditionSelected() bool {
	return o.HasMedicalCondition != nil && !*o.HasMedicalCondition
}

func (o ListIndividualsOptions) IsNeedsLegalAndPhysicalProtectionSelected() bool {
	return o.NeedsLegalAndPhysicalProtection != nil && *o.NeedsLegalAndPhysicalProtection
}

func (o ListIndividualsOptions) IsNotNeedsLegalAndPhysicalProtectionSelected() bool {
	return o.NeedsLegalAndPhysicalProtection != nil && !*o.NeedsLegalAndPhysicalProtection
}

func (o ListIndividualsOptions) IsMinorHeadedHouseholdSelected() bool {
	return o.IsMinorHeadedHousehold != nil && *o.IsMinorHeadedHousehold
}

func (o ListIndividualsOptions) IsNotMinorHeadedHouseholdSelected() bool {
	return o.IsMinorHeadedHousehold != nil && !*o.IsMinorHeadedHousehold
}

func (o ListIndividualsOptions) IsFemaleHeadedHouseholdSelected() bool {
	return o.IsFemaleHeadedHousehold != nil && *o.IsFemaleHeadedHousehold
}

func (o ListIndividualsOptions) IsNotFemaleHeadedHouseholdSelected() bool {
	return o.IsFemaleHeadedHousehold != nil && !*o.IsFemaleHeadedHousehold
}

func (o ListIndividualsOptions) IsMinorSelected() bool {
	return o.IsMinor != nil && *o.IsMinor
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

func (o ListIndividualsOptions) IsHasDisabilitySelected() bool {
	return o.HasDisability != nil && *o.HasDisability
}

func (o ListIndividualsOptions) IsNotHasDisabilitySelected() bool {
	return o.HasDisability != nil && !*o.HasDisability
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
