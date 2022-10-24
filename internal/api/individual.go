package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
)

type Individual struct {
	ID                         string     `db:"id"`
	CountryID                  string     `db:"country_id"`
	FullName                   string     `db:"full_name"`
	PhoneNumber                string     `db:"phone_number"`
	NormalizedPhoneNumber      string     `db:"normalized_phone_number"`
	Email                      string     `db:"email"`
	Address                    string     `db:"address"`
	BirthDate                  *time.Time `db:"birth_date"`
	Gender                     string     `db:"gender"`
	DisplacementStatus         string     `db:"displacement_status"`
	PreferredName              string     `db:"preferred_name"`
	IsMinor                    bool       `db:"is_minor"`
	PresentsProtectionConcerns bool       `db:"presents_protection_concerns"`
	PhysicalImpairment         string     `db:"physical_impairment"`
	SensoryImpairment          string     `db:"sensory_impairment"`
	MentalImpairment           string     `db:"mental_impairment"`
}

func (i Individual) GetFieldValue(field string) (interface{}, error) {
	switch field {
	case constants.DBColumnIndividualAddress:
		return i.Address, nil
	case constants.DBColumnIndividualBirthDate:
		return i.BirthDate, nil
	case constants.DBColumnIndividualCountryID:
		return i.CountryID, nil
	case constants.DBColumnIndividualDisplacementStatus:
		return i.DisplacementStatus, nil
	case constants.DBColumnIndividualEmail:
		return i.Email, nil
	case constants.DBColumnIndividualFullName:
		return i.FullName, nil
	case constants.DBColumnIndividualGender:
		return i.Gender, nil
	case constants.DBColumnIndividualID:
		return i.ID, nil
	case constants.DBColumnIndividualIsMinor:
		return i.IsMinor, nil
	case constants.DBColumnIndividualMentalImpairment:
		return i.MentalImpairment, nil
	case constants.DBColumnIndividualNormalizedPhoneNumber:
		return i.NormalizedPhoneNumber, nil
	case constants.DBColumnIndividualPhoneNumber:
		return i.PhoneNumber, nil
	case constants.DBColumnIndividualPhysicalImpairment:
		return i.PhysicalImpairment, nil
	case constants.DBColumnIndividualPreferredName:
		return i.PreferredName, nil
	case constants.DBColumnIndividualPresentsProtectionConcerns:
		return i.PresentsProtectionConcerns, nil
	case constants.DBColumnIndividualSensoryImpairment:
		return i.SensoryImpairment, nil
	default:
		return nil, fmt.Errorf("unknown field: %s", field)
	}
}

func (i Individual) String() string {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func (individual *Individual) Normalize() {
	individual.ID = trimString(individual.ID)
	individual.FullName = trimString(individual.FullName)
	individual.PreferredName = trimString(individual.PreferredName)
	if individual.PreferredName == "" {
		individual.PreferredName = individual.FullName
	}
	individual.DisplacementStatus = trimString(individual.DisplacementStatus)
	individual.Email = trimString(normalizeEmail(individual.Email))
	individual.PhoneNumber = trimString(individual.PhoneNumber)
	individual.Address = trimString(individual.Address)
	individual.Gender = trimString(individual.Gender)
	individual.NormalizedPhoneNumber = NormalizePhoneNumber(individual.PhoneNumber)
	individual.PhysicalImpairment = trimString(individual.PhysicalImpairment)
	individual.MentalImpairment = trimString(individual.MentalImpairment)
	individual.SensoryImpairment = trimString(individual.SensoryImpairment)
}

type GetAllOptions struct {
	Address                    string
	ID                         string
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
	var params = &url.Values{}
	if len(o.FullName) != 0 {
		params.Add("full_name", o.FullName)
	}
	if len(o.ID) != 0 {
		params.Add("id", o.ID)
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

	if *o.IsMinor {
		params.Add("is_minor", "true")
	} else {
		params.Add("is_minor", "false")
	}
	if o.IsPresentsProtectionConcernsSelected() {
		params.Add("presents_protection_concerns", "true")
	} else if o.IsNotPresentsProtectionConcernsSelected() {
		params.Add("presents_protection_concerns", "false")
	}
	if len(o.CountryID) != 0 {
		params.Add("country_id", o.CountryID)
	}
	if len(o.DisplacementStatuses) != 0 {
		params.Add("displacement_status", strings.Join(o.DisplacementStatuses, ","))
	}
	u := url.URL{
		Path: "/individuals",
	}
	u.RawQuery = params.Encode()
	return template.HTML(u.String())
}
