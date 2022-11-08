package api

import (
	"html/template"
	"net/url"
	"time"

	"github.com/nrc-no/notcore/internal/containers"
)

type ListIndividualsOptions struct {
	Address                        string
	AgeFrom                        *int
	AgeTo                          *int
	BirthDateFrom                  *time.Time
	BirthDateTo                    *time.Time
	CognitiveDisabilityLevel       DisabilityLevel
	CollectionAdministrativeArea1  string
	CollectionAdministrativeArea2  string
	CollectionAdministrativeArea3  string
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
	Genders                        containers.Set[Gender]
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
	IdentificationContext          string
	InternalID                     string
	IsHeadOfCommunity              *bool
	IsHeadOfHousehold              *bool
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

func (o ListIndividualsOptions) QueryParams() template.HTML {
	params := newListIndividualsOptionsEncoder(o, time.Now()).encode()
	u := url.URL{Path: "/countries/" + o.CountryID + "/individuals"}
	u.RawQuery = params.Encode()
	return template.HTML(u.String())
}

func NewIndividualListFromURLValues(values url.Values, into *ListIndividualsOptions) error {
	parser := listIndividualsOptionsDecoder{
		out:    into,
		values: values,
		now:    time.Now(),
	}
	return parser.parse()
}
