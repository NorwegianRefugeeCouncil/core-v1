package api

import (
	"fmt"
	"html/template"
	"net/url"
	"strings"
	"time"
)

type GetAllOptions struct {
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
	Sort                       SortOptions
}

type SortDirection string

const (
	SortDirectionNone       SortDirection = "none"
	SortDirectionAscending  SortDirection = "ascending"
	SortDirectionDescending SortDirection = "descending"
)

type SortOption struct {
	Field     string
	Direction SortDirection
}

type SortOptions []SortOption

func (o GetAllOptions) IsMinorSelected() bool {
	return o.IsMinor != nil && *o.IsMinor
}

func (o GetAllOptions) IsNotMinorSelected() bool {
	return o.IsMinor != nil && !*o.IsMinor
}

func (o GetAllOptions) IsPresentsProtectionConcernsSelected() bool {
	return o.PresentsProtectionConcerns != nil && *o.PresentsProtectionConcerns
}

func (o GetAllOptions) IsNotPresentsProtectionConcernsSelected() bool {
	return o.PresentsProtectionConcerns != nil && !*o.PresentsProtectionConcerns
}

func (o GetAllOptions) AgeFrom() int {
	if o.BirthDateTo == nil {
		return 0
	}
	now := time.Now()
	return now.Year() - o.BirthDateTo.Year() - 1
}

func (o GetAllOptions) AgeTo() int {
	if o.BirthDateFrom == nil {
		return 0
	}
	now := time.Now()
	return now.Year() - o.BirthDateFrom.Year() - 1
}

func (o GetAllOptions) NextPage() GetAllOptions {
	ret := o
	ret.Skip += ret.Take
	return ret
}

func (o GetAllOptions) PreviousPage() GetAllOptions {
	ret := o
	ret.Skip -= ret.Take
	if ret.Skip < 0 {
		ret.Skip = 0
	}
	return ret
}

func (o GetAllOptions) FirstPage() GetAllOptions {
	ret := o
	ret.Skip = 0
	return ret
}

func (o GetAllOptions) QueryParams() template.HTML {
	params := o.buildUrlValues()
	u := url.URL{
		Path: "/countries/" + o.CountryID + "/individuals",
	}
	u.RawQuery = params.Encode()
	return template.HTML(u.String())
}

func (o GetAllOptions) GetSortURLForField(field string, direction SortDirection) template.HTML {
	newOptions := o
	newOptions.Sort = []SortOption{
		{Field: field, Direction: direction},
	}
	return newOptions.QueryParams()
}

func (o GetAllOptions) GetSortDirectionForField(field string) SortDirection {
	for _, o := range o.Sort {
		if o.Field == field {
			return o.Direction
		}
	}
	return SortDirectionNone
}

func (o GetAllOptions) buildUrlValues() *url.Values {
	var params = &url.Values{}
	if len(o.FullName) != 0 {
		params.Add("full_name", o.FullName)
	}
	if len(o.IDs) != 0 {
		params.Add("ids", strings.Join(o.IDs, ","))
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
	if len(o.Sort) != 0 {
		for _, s := range o.Sort {
			params.Add("sort_by", s.Field)
			params.Add("sort_direction", string(s.Direction))
		}
	}
	return params
}
