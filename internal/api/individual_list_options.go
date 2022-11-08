package api

import (
	"html/template"
	"net/url"
	"time"

	"github.com/nrc-no/notcore/internal/containers"
)

type ListIndividualsOptions struct {
	Address                    string
	IDs                        containers.StringSet
	BirthDateFrom              *time.Time
	BirthDateTo                *time.Time
	AgeFrom                    *int
	AgeTo                      *int
	CountryID                  string
	DisplacementStatuses       containers.Set[DisplacementStatus]
	Email                      string
	FullName                   string
	Genders                    containers.Set[Gender]
	IsMinor                    *bool
	PhoneNumber                string
	PresentsProtectionConcerns *bool
	Skip                       int
	Take                       int
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
