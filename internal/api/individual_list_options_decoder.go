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
	ageFrom, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeTo))
	if err != nil || ageFrom == nil {
		return err
	}
	yearsAgo := calculateBirthDateFromAge(*ageFrom, p.now)
	p.out.BirthDateFrom = &yearsAgo
	return nil
}

func (p *listIndividualsOptionsDecoder) parseBirthDateTo() error {
	ageTo, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeFrom))
	if err != nil || ageTo == nil {
		return err
	}
	yearsAgo := calculateBirthDateFromAge(*ageTo, p.now)
	p.out.BirthDateTo = &yearsAgo
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

func calculateBirthDateFromAge(age int, now time.Time) time.Time {
	return now.AddDate(0, 0, -(age+1)*365).Truncate(24 * time.Hour)
}

func calculateAgeFromBirthDate(birthDate time.Time, now time.Time) int {
	return now.Year() - birthDate.Year() - 1 // -1 because we want to round down
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
