package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"time"
)

type Individual struct {
	ID                    string     `db:"id"`
	FullName              string     `db:"full_name"`
	PhoneNumber           string     `db:"phone_number"`
	NormalizedPhoneNumber string     `db:"normalized_phone_number"`
	Email                 string     `db:"email"`
	Address               string     `db:"address"`
	BirthDate             *time.Time `db:"birth_date"`
	Gender                string     `db:"gender"`
}

func (i Individual) String() string {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

type GetAllOptions struct {
	Genders       []string
	BirthDateFrom *time.Time
	BirthDateTo   *time.Time
	PhoneNumber   string
	Address       string
	Take          int
	Skip          int
	Email         string
	FullName      string
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

	u := url.URL{
		Path: "/individuals",
	}
	u.RawQuery = params.Encode()
	return template.HTML(u.String())
}
