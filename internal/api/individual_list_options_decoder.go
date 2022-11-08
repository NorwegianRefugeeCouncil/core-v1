package api

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
)

type listIndividualsOptionsDecoder struct {
	out    *ListIndividualsOptions
	values url.Values
	now    time.Time
}

func (p *listIndividualsOptionsDecoder) parse() error {
	fns := []func() error{
		p.parseSkip,
		p.parseTake,
		p.parseCountryID,
		p.parseFullName,
		p.parseEmail,
		p.parsePhoneNumber,
		p.parseAddress,
		p.parseIDs,
		p.parseGenders,
		p.parseDisplacementStatuses,
		p.parseBirthDateFrom,
		p.parseBirthDateTo,
		p.parseAgeFrom,
		p.parseAgeTo,
		p.parseIsMinor,
		p.parsePresentsProtectionConcerns,
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func (p *listIndividualsOptionsDecoder) parseSkip() (err error) {
	var skip *int
	skip, err = parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsSkip))
	if err != nil || skip == nil {
		return err
	}
	p.out.Skip = *skip
	if p.out.Skip < 0 {
		return fmt.Errorf("skip must be greater or equal to 0")
	}
	return nil
}

func (p *listIndividualsOptionsDecoder) parseTake() (err error) {
	var take *int
	take, err = parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsTake))
	if err != nil || take == nil {
		return err
	}
	p.out.Take = *take
	if p.out.Take < 0 {
		return fmt.Errorf("take must be greater or equal to 0")
	}
	return nil
}

func (p *listIndividualsOptionsDecoder) parseFullName() error {
	p.out.FullName = p.values.Get(constants.FormParamsGetIndividualsFullName)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseCountryID() error {
	p.out.CountryID = p.values.Get(constants.FormParamsGetIndividualsCountryID)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseEmail() error {
	p.out.Email = p.values.Get(constants.FormParamsGetIndividualsEmail)
	return nil
}

func (p *listIndividualsOptionsDecoder) parsePhoneNumber() error {
	p.out.PhoneNumber = p.values.Get(constants.FormParamsGetIndividualsPhoneNumber)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseAddress() error {
	p.out.Address = p.values.Get(constants.FormParamsGetIndividualsAddress)
	return nil
}

func (p *listIndividualsOptionsDecoder) parseIDs() error {
	if len(p.values[constants.FormParamsGetIndividualsID]) == 0 {
		return nil
	}
	idSet := containers.NewStringSet()
	idSet.Add(p.values[constants.FormParamsGetIndividualsID]...)
	if idSet.IsEmpty() {
		return nil
	}
	p.out.IDs = idSet
	return nil
}

func (p *listIndividualsOptionsDecoder) parseGenders() error {
	if len(p.values[constants.FormParamsGetIndividualsGender]) == 0 {
		return nil
	}
	genderSet := containers.NewSet[Gender]()
	for _, g := range p.values[constants.FormParamsGetIndividualsGender] {
		gender, err := ParseGender(g)
		if err != nil {
			return err
		}
		genderSet.Add(gender)
	}
	p.out.Genders = genderSet
	return nil
}

func (p *listIndividualsOptionsDecoder) parseDisplacementStatuses() error {
	if len(p.values[constants.FormParamsGetIndividualsDisplacementStatus]) == 0 {
		return nil
	}
	dsSet := containers.NewSet[DisplacementStatus]()
	for _, ds := range p.values[constants.FormParamsGetIndividualsDisplacementStatus] {
		parsedDs, err := ParseDisplacementStatus(ds)
		if err != nil {
			return err
		}
		dsSet.Add(parsedDs)
	}
	p.out.DisplacementStatuses = dsSet
	return nil
}

func (p *listIndividualsOptionsDecoder) parseBirthDateFrom() error {
	birthDateFrom, err := parseOptionalDate(p.values.Get(constants.FormParamsGetIndividualsBirthDateFrom))
	if err != nil || birthDateFrom == nil {
		return err
	}
	p.out.BirthDateFrom = birthDateFrom
	return nil
}

func (p *listIndividualsOptionsDecoder) parseBirthDateTo() error {
	birthDateTo, err := parseOptionalDate(p.values.Get(constants.FormParamsGetIndividualsBirthDateTo))
	if err != nil || birthDateTo == nil {
		return err
	}
	p.out.BirthDateTo = birthDateTo
	return nil
}

func (p *listIndividualsOptionsDecoder) parseAgeFrom() error {
	ageFrom, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeFrom))
	if err != nil || ageFrom == nil {
		return err
	}
	p.out.AgeFrom = ageFrom
	return nil
}

func (p *listIndividualsOptionsDecoder) parseAgeTo() error {
	ageTo, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeTo))
	if err != nil || ageTo == nil {
		return err
	}
	p.out.AgeTo = ageTo
	return nil
}

func (p *listIndividualsOptionsDecoder) parseIsMinor() (err error) {
	p.out.IsMinor, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsIsMinor))
	return err
}

func (p *listIndividualsOptionsDecoder) parsePresentsProtectionConcerns() (err error) {
	p.out.PresentsProtectionConcerns, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsPresentsProtectionConcerns))
	return err
}

func parseOptionalInt(strValue string) (*int, error) {
	if len(strValue) == 0 {
		return nil, nil
	}
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}

func parseOptionalDate(strValue string) (*time.Time, error) {
	if len(strValue) == 0 {
		return nil, nil
	}
	dateValue, err := time.Parse("2006-01-02", strValue)
	if err != nil {
		return nil, err
	}
	return &dateValue, nil
}

func parseOptionalBool(val string) (*bool, error) {
	if len(val) == 0 {
		return nil, nil
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &boolVal, nil
}
