package api

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
)

func newListIndividualsOptionsEncoder(values ListIndividualsOptions, now time.Time) *listIndividualsOptionsEncoder {
	return &listIndividualsOptionsEncoder{
		values: values,
		now:    now,
	}
}

type listIndividualsOptionsEncoder struct {
	out    url.Values
	values ListIndividualsOptions
	now    time.Time
}

func (p *listIndividualsOptionsEncoder) encode() url.Values {
	p.out = url.Values{}
	fns := []func(){
		p.encodeFullName,
		p.encodeAddress,
		p.encodeIDs,
		p.encodeEmail,
		p.encodePhoneNumber,
		p.encodeSkip,
		p.encodeTake,
		p.encodeGenders,
		p.encodeBirthDateFrom,
		p.encodeBirthDateTo,
		p.encodeIsMinor,
		p.encodePresentsProtectionConcerns,
		p.encodeDisplacementStatuses,
	}
	for _, fn := range fns {
		fn()
	}
	return p.out
}

func (p *listIndividualsOptionsEncoder) encodeFullName() {
	if len(p.values.FullName) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsFullName, p.values.FullName)
	}
}

func (p *listIndividualsOptionsEncoder) encodeIDs() {
	if len(p.values.IDs) != 0 {
		for _, id := range p.values.IDs.Items() {
			p.out.Add(constants.FormParamsGetIndividualsID, id)
		}
	}
}

func (p *listIndividualsOptionsEncoder) encodeAddress() {
	if len(p.values.Address) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsAddress, p.values.Address)
	}
}

func (p *listIndividualsOptionsEncoder) encodeEmail() {
	if len(p.values.Email) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsEmail, p.values.Email)
	}
}

func (p *listIndividualsOptionsEncoder) encodePhoneNumber() {
	if len(p.values.PhoneNumber) != 0 {
		p.out.Add(constants.FormParamsGetIndividualsPhoneNumber, p.values.PhoneNumber)
	}
}

func (p *listIndividualsOptionsEncoder) encodeTake() {
	if p.values.Take != 0 {
		p.out.Add(constants.FormParamsGetIndividualsTake, fmt.Sprintf("%d", p.values.Take))
	}
}

func (p *listIndividualsOptionsEncoder) encodeSkip() {
	if p.values.Skip != 0 {
		p.out.Add(constants.FormParamsGetIndividualsSkip, fmt.Sprintf("%d", p.values.Skip))
	}
}

func (p *listIndividualsOptionsEncoder) encodeGenders() {
	if len(p.values.Genders) > 0 {
		for _, g := range p.values.Genders.Items() {
			p.out.Add(constants.FormParamsGetIndividualsGender, string(g))
		}
	}
}

func (p *listIndividualsOptionsEncoder) encodeBirthDateFrom() {
	if p.values.BirthDateFrom != nil {
		p.out.Add(constants.FormParamsGetIndividualsAgeTo, strconv.Itoa(calculateAgeFromBirthDate(*p.values.BirthDateFrom, p.now)))
	}
}

func (p *listIndividualsOptionsEncoder) encodeBirthDateTo() {
	if p.values.BirthDateTo != nil {
		p.out.Add(constants.FormParamsGetIndividualsAgeFrom, strconv.Itoa(calculateAgeFromBirthDate(*p.values.BirthDateTo, p.now)))
	}
}

func (p *listIndividualsOptionsEncoder) encodeIsMinor() {
	if p.values.IsMinor != nil {
		p.out.Add(constants.FormParamsGetIndividualsIsMinor, strconv.FormatBool(*p.values.IsMinor))
	}
}

func (p *listIndividualsOptionsEncoder) encodePresentsProtectionConcerns() {
	if p.values.PresentsProtectionConcerns != nil {
		p.out.Add(constants.FormParamsGetIndividualsPresentsProtectionConcerns, strconv.FormatBool(*p.values.PresentsProtectionConcerns))
	}
}

func (p *listIndividualsOptionsEncoder) encodeDisplacementStatuses() {
	if len(p.values.DisplacementStatuses) > 0 {
		for _, ds := range p.values.DisplacementStatuses.Items() {
			p.out.Add(constants.FormParamsGetIndividualsDisplacementStatus, string(ds))
		}
	}
}
