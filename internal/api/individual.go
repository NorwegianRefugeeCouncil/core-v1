package api

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"time"

	"github.com/lib/pq"
)

type Individual struct {
	ID                         string     `db:"id"`
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
	Countries                  Countries  `db:"countries"`
}

type Countries []string

func (c *Countries) Scan(value interface{}) error {
	var ret = make([]string, 0)
	if err := pq.Array(&ret).Scan(value); err != nil {
		return err
	}
	*c = ret
	return nil
}

func (c Countries) Value() (driver.Value, error) {
	return pq.Array([]string(c)).Value()
}

var AllndividualFields = []string{
	"id",
	"full_name",
	"phone_number",
	"normalized_phone_number",
	"email",
	"address",
	"birth_date",
	"gender",
	"displacement_status",
	"preferred_name",
	"is_minor",
	"presents_protection_concerns",
	"physical_impairment",
	"sensory_impairment",
	"mental_impairment",
	"countries",
}

func (i Individual) GetFieldValue(field string) (interface{}, error) {
	switch field {
	case "id":
		return i.ID, nil
	case "full_name":
		return i.FullName, nil
	case "phone_number":
		return i.PhoneNumber, nil
	case "normalized_phone_number":
		return i.NormalizedPhoneNumber, nil
	case "email":
		return i.Email, nil
	case "address":
		return i.Address, nil
	case "birth_date":
		return i.BirthDate, nil
	case "gender":
		return i.Gender, nil
	case "displacement_status":
		return i.DisplacementStatus, nil
	case "preferred_name":
		return i.PreferredName, nil
	case "is_minor":
		return i.IsMinor, nil
	case "presents_protection_concerns":
		return i.PresentsProtectionConcerns, nil
	case "physical_impairment":
		return i.PhysicalImpairment, nil
	case "sensory_impairment":
		return i.SensoryImpairment, nil
	case "mental_impairment":
		return i.MentalImpairment, nil
	case "countries":
		return i.Countries, nil
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

type GetAllOptions struct {
	Genders                    []string
	BirthDateFrom              *time.Time
	BirthDateTo                *time.Time
	PhoneNumber                string
	Address                    string
	Take                       int
	Skip                       int
	Email                      string
	FullName                   string
	IsMinor                    *bool
	PresentsProtectionConcerns *bool
	Countries                  []string
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
	if o.BirthDateFrom == nil {
		return 0
	}
	now := time.Now()
	return now.Year() - o.BirthDateTo.Year() - 1
}

func (o GetAllOptions) AgeTo() int {
	if o.BirthDateTo == nil {
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

	u := url.URL{
		Path: "/individuals",
	}
	u.RawQuery = params.Encode()
	return template.HTML(u.String())
}
