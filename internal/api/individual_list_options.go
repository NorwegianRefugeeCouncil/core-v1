package api

import (
	"fmt"
	"html/template"
	"net/url"
	"strings"
	"time"
)

type IndividualListOptions struct {
	Address                    string
	IDs                        []string
	BirthDateFrom              *time.Time
	BirthDateTo                *time.Time
	CountryID                  string
	DisplacementStatuses       []string
	Email                      string
	FullName                   string
	Genders                    []string
	IsMinor                    *bool
	PhoneNumber                string
	PresentsProtectionConcerns *bool
	Skip                       int
	Take                       int
}

func (o IndividualListOptions) IsMinorSelected() bool {
	return o.IsMinor != nil && *o.IsMinor
}

func (o IndividualListOptions) IsNotMinorSelected() bool {
	return o.IsMinor != nil && !*o.IsMinor
}

func (o IndividualListOptions) IsPresentsProtectionConcernsSelected() bool {
	return o.PresentsProtectionConcerns != nil && *o.PresentsProtectionConcerns
}

func (o IndividualListOptions) IsNotPresentsProtectionConcernsSelected() bool {
	return o.PresentsProtectionConcerns != nil && !*o.PresentsProtectionConcerns
}

func (o IndividualListOptions) AgeFrom() int {
	if o.BirthDateTo == nil {
		return 0
	}
	now := time.Now()
	return now.Year() - o.BirthDateTo.Year() - 1
}

func (o IndividualListOptions) AgeTo() int {
	if o.BirthDateFrom == nil {
		return 0
	}
	now := time.Now()
	return now.Year() - o.BirthDateFrom.Year() - 1
}

func (o IndividualListOptions) NextPage() IndividualListOptions {
	ret := o
	ret.Skip += ret.Take
	return ret
}

func (o IndividualListOptions) PreviousPage() IndividualListOptions {
	ret := o
	ret.Skip -= ret.Take
	if ret.Skip < 0 {
		ret.Skip = 0
	}
	return ret
}

func (o IndividualListOptions) FirstPage() IndividualListOptions {
	ret := o
	ret.Skip = 0
	return ret
}

func (o IndividualListOptions) WithTake(take int) IndividualListOptions {
	ret := o
	ret.Take = take
	return ret
}

func (o IndividualListOptions) QueryParams() template.HTML {
	var params = &url.Values{}
	if len(o.FullName) != 0 {
		params.Add("full_name", o.FullName)
	}
	if len(o.IDs) != 0 {
		for _, id := range o.IDs {
			params.Add("id", id)
		}
	}
	if len(o.Address) != 0 {
		params.Add("address", o.Address)
	}
	if len(o.Email) != 0 {
		params.Add("email", o.Email)
	}
	if len(o.PhoneNumber) != 0 {
		params.Add("phone_number", o.PhoneNumber)
	}
	if o.Take != 0 {
		params.Add("take", fmt.Sprintf("%d", o.Take))
	}
	if o.Skip != 0 {
		params.Add("skip", fmt.Sprintf("%d", o.Skip))
	}
	if len(o.Genders) != 0 {
		for _, g := range o.Genders {
			params.Add("gender", g)
		}
	}
	if o.BirthDateFrom != nil {
		params.Add("age_from", fmt.Sprintf("%d", o.AgeTo()))
	}
	if o.BirthDateTo != nil {
		params.Add("age_to", fmt.Sprintf("%d", o.AgeFrom()))
	}

	if o.IsMinorSelected() {
		params.Add("is_minor", "true")
	} else if o.IsNotMinorSelected() {
		params.Add("is_minor", "false")
	}

	if o.IsPresentsProtectionConcernsSelected() {
		params.Add("presents_protection_concerns", "true")
	} else if o.IsNotPresentsProtectionConcernsSelected() {
		params.Add("presents_protection_concerns", "false")
	}
	if len(o.DisplacementStatuses) != 0 {
		params.Add("displacement_status", strings.Join(o.DisplacementStatuses, ","))
	}
	u := url.URL{
		Path: "/countries/" + o.CountryID + "/individuals",
	}
	u.RawQuery = params.Encode()
	return template.HTML(u.String())
}
