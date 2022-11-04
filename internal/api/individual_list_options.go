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

func NewIndividualListFromURLValues(values url.Values, out *ListIndividualsOptions) error {
	parser := listIndividualsOptionsParser{
		out:    out,
		values: values,
		now:    time.Now(),
	}
	return parser.parse()
}

type listIndividualsOptionsParser struct {
	out    *ListIndividualsOptions
	values url.Values
	now    time.Time
}

func (p *listIndividualsOptionsParser) parse() error {
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

func (p *listIndividualsOptionsParser) parseSkip() (err error) {
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

func (p *listIndividualsOptionsParser) parseTake() (err error) {
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

func (p *listIndividualsOptionsParser) parseFullName() error {
	p.out.FullName = p.values.Get(constants.FormParamsGetIndividualsFullName)
	return nil
}

func (p *listIndividualsOptionsParser) parseCountryID() error {
	p.out.CountryID = p.values.Get(constants.FormParamsGetIndividualsCountryID)
	return nil
}

func (p *listIndividualsOptionsParser) parseEmail() error {
	p.out.Email = p.values.Get(constants.FormParamsGetIndividualsEmail)
	return nil
}

func (p *listIndividualsOptionsParser) parsePhoneNumber() error {
	p.out.PhoneNumber = p.values.Get(constants.FormParamsGetIndividualsPhoneNumber)
	return nil
}

func (p *listIndividualsOptionsParser) parseAddress() error {
	p.out.Address = p.values.Get(constants.FormParamsGetIndividualsAddress)
	return nil
}

func (p *listIndividualsOptionsParser) parseIDs() error {
	idSet := containers.NewStringSet()
	idSet.Add(p.values[constants.FormParamsGetIndividualsID]...)
	if idSet.IsEmpty() {
		return nil
	}
	p.out.IDs = idSet.Items()
	return nil
}

func (p *listIndividualsOptionsParser) parseGenders() error {
	genderSet := containers.NewStringSet()
	genderSet.Add(p.values[constants.FormParamsGetIndividualsGender]...)
	if genderSet.IsEmpty() {
		return nil
	}
	p.out.Genders = genderSet.Items()
	return nil
}

func (p *listIndividualsOptionsParser) parseDisplacementStatuses() error {
	dsSet := containers.NewStringSet()
	dsSet.Add(p.values[constants.FormParamsGetIndividualsDisplacementStatus]...)
	if dsSet.IsEmpty() {
		return nil
	}
	p.out.DisplacementStatuses = dsSet.Items()
	return nil
}

func (p *listIndividualsOptionsParser) parseBirthDateFrom() error {
	ageFrom, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeTo))
	if err != nil || ageFrom == nil {
		return err
	}
	yearsAgo := calculateBirthDateFromAge(*ageFrom, p.now)
	p.out.BirthDateFrom = &yearsAgo
	return nil
}

func (p *listIndividualsOptionsParser) parseBirthDateTo() error {
	ageTo, err := parseOptionalInt(p.values.Get(constants.FormParamsGetIndividualsAgeFrom))
	if err != nil || ageTo == nil {
		return err
	}
	yearsAgo := calculateBirthDateFromAge(*ageTo, p.now)
	p.out.BirthDateTo = &yearsAgo
	return nil
}

func (p *listIndividualsOptionsParser) parseIsMinor() (err error) {
	p.out.IsMinor, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsIsMinor))
	return err
}

func (p *listIndividualsOptionsParser) parsePresentsProtectionConcerns() (err error) {
	p.out.PresentsProtectionConcerns, err = parseOptionalBool(p.values.Get(constants.FormParamsGetIndividualsPresentsProtectionConcerns))
	return err
}

func calculateBirthDateFromAge(age int, now time.Time) time.Time {
	return now.AddDate(0, 0, -(age+1)*365).Truncate(24 * time.Hour)
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
