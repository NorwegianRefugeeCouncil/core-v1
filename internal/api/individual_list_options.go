package api

import (
	"fmt"
	"html/template"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
)

type ListIndividualsOptions struct {
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

func (o ListIndividualsOptions) AgeFrom() int {
	if o.BirthDateTo == nil {
		return 0
	}
	now := time.Now()
	return now.Year() - o.BirthDateTo.Year() - 1
}

func (o ListIndividualsOptions) AgeTo() int {
	if o.BirthDateFrom == nil {
		return 0
	}
	now := time.Now()
	return now.Year() - o.BirthDateFrom.Year() - 1
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
	var params = &url.Values{}
	if len(o.FullName) != 0 {
		params.Add(constants.FormParamsGetIndividualsFullName, o.FullName)
	}
	if len(o.IDs) != 0 {
		for _, id := range o.IDs {
			params.Add(constants.FormParamsGetIndividualsID, id)
		}
	}
	if len(o.Address) != 0 {
		params.Add(constants.FormParamsGetIndividualsAddress, o.Address)
	}
	if len(o.Email) != 0 {
		params.Add(constants.FormParamsGetIndividualsEmail, o.Email)
	}
	if len(o.PhoneNumber) != 0 {
		params.Add(constants.FormParamsGetIndividualsPhoneNumber, o.PhoneNumber)
	}
	if o.Take != 0 {
		params.Add(constants.FormParamsGetIndividualsTake, fmt.Sprintf("%d", o.Take))
	}
	if o.Skip != 0 {
		params.Add(constants.FormParamsGetIndividualsSkip, fmt.Sprintf("%d", o.Skip))
	}
	if len(o.Genders) != 0 {
		for _, g := range o.Genders {
			params.Add(constants.FormParamsGetIndividualsGender, g)
		}
	}
	if o.BirthDateFrom != nil {
		params.Add(constants.FormParamsGetIndividualsAgeFrom, fmt.Sprintf("%d", o.AgeTo()))
	}
	if o.BirthDateTo != nil {
		params.Add(constants.FormParamsGetIndividualsAgeTo, fmt.Sprintf("%d", o.AgeFrom()))
	}

	if o.IsMinorSelected() {
		params.Add(constants.FormParamsGetIndividualsIsMinor, strconv.FormatBool(true))
	} else if o.IsNotMinorSelected() {
		params.Add(constants.FormParamsGetIndividualsIsMinor, strconv.FormatBool(false))
	}

	if o.IsPresentsProtectionConcernsSelected() {
		params.Add(constants.FormParamsGetIndividualsPresentsProtectionConcerns, strconv.FormatBool(true))
	} else if o.IsNotPresentsProtectionConcernsSelected() {
		params.Add(constants.FormParamsGetIndividualsPresentsProtectionConcerns, strconv.FormatBool(false))
	}
	if len(o.DisplacementStatuses) != 0 {
		params.Add(constants.FormParamsGetIndividualsDisplacementStatus, strings.Join(o.DisplacementStatuses, ","))
	}
	u := url.URL{
		Path: "/countries/" + o.CountryID + "/individuals",
	}
	u.RawQuery = params.Encode()
	return template.HTML(u.String())
}

func parseQryParamInt(strValue string) (int, error) {
	if len(strValue) != 0 {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			return 0, err
		}
		return intValue, nil
	}
	return 0, nil
}

func NewIndividualListFromURLValues(values url.Values, out *ListIndividualsOptions) error {

	var err error
	out.Take, err = parseQryParamInt(values.Get(constants.FormParamsGetIndividualsTake))
	if err != nil {
		return err
	}
	if out.Take <= 0 || out.Take > 100 {
		out.Take = 20
	}

	out.Skip, err = parseQryParamInt(values.Get(constants.FormParamsGetIndividualsSkip))
	if err != nil {
		return err
	}
	if out.Skip < 0 {
		out.Skip = 0
	}

	out.Email = values.Get(constants.FormParamsGetIndividualsEmail)
	out.FullName = values.Get(constants.FormParamsGetIndividualsFullName)
	out.PhoneNumber = values.Get(constants.FormParamsGetIndividualsPhoneNumber)
	out.Address = values.Get(constants.FormParamsGetIndividualsAddress)
	out.Genders = values[constants.FormParamsGetIndividualsGender]

	switch values.Get(constants.FormParamsGetIndividualsIsMinor) {
	case strconv.FormatBool(true):
		isMinor := true
		out.IsMinor = &isMinor
	case strconv.FormatBool(false):
		isMinor := false
		out.IsMinor = &isMinor
	}

	switch values.Get(constants.FormParamsGetIndividualsPresentsProtectionConcerns) {
	case strconv.FormatBool(true):
		presentsProtectionConcerns := true
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	case strconv.FormatBool(false):
		presentsProtectionConcerns := false
		out.PresentsProtectionConcerns = &presentsProtectionConcerns
	}

	ageFromStr := values.Get(constants.FormParamsGetIndividualsAgeFrom)
	if len(ageFromStr) != 0 {
		ageFrom, err := parseQryParamInt(values.Get(constants.FormParamsGetIndividualsAgeFrom))
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageFrom+1)*365)
		out.BirthDateTo = &yearsAgo
	}
	ageToStr := values.Get(constants.FormParamsGetIndividualsAgeTo)
	if len(ageToStr) != 0 {
		ageTo, err := parseQryParamInt(values.Get(constants.FormParamsGetIndividualsAgeTo))
		if err != nil {
			return err
		}
		yearsAgo := time.Now().AddDate(0, 0, -(ageTo+1)*365)
		out.BirthDateFrom = &yearsAgo
	}
	out.CountryID = values.Get(constants.FormParamsGetIndividualsCountryID)
	displacementStatuses := values[constants.FormParamsGetIndividualsDisplacementStatus]
	var displacementStatusMap = map[string]bool{}
	for _, s := range displacementStatuses {
		if displacementStatusMap[s] {
			continue
		}
		displacementStatusMap[s] = true
		out.DisplacementStatuses = append(out.DisplacementStatuses, s)
	}

	idValues := values[constants.FormParamsGetIndividualsID]
	idSet := containers.NewStringSet()
	for _, v := range idValues {
		parts := strings.Split(v, ",")
		for _, p := range parts {
			if len(p) == 0 {
				continue
			}
			idSet.Add(p)
		}
	}
	out.IDs = idSet.Items()
	return nil
}
