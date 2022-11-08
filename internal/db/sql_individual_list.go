package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
)

type getAllIndividualsSQLQuery struct {
	*strings.Builder
	driverName string
	a          []interface{}
}

func newGetAllIndividualsSQLQuery(driverName string, options api.ListIndividualsOptions) *getAllIndividualsSQLQuery {
	qry := &getAllIndividualsSQLQuery{
		Builder:    &strings.Builder{},
		driverName: driverName,
	}
	qry = qry.
		writeString("SELECT * FROM individual_registrations WHERE deleted_at IS NULL").
		withIds(options.IDs).
		withFullName(options.FullName).
		withAddress(options.Address).
		withGenders(options.Genders).
		withBirthDateFrom(options.BirthDateFrom).
		withBirthDateTo(options.BirthDateTo).
		withAgeFrom(options.AgeFrom).
		withAgeTo(options.AgeTo).
		withPhoneNumber(options.PhoneNumber).
		withEmail(options.Email).
		withIsMinor(options.IsMinor).
		withPresentsProtectionConcerns(options.PresentsProtectionConcerns).
		withCountryID(options.CountryID).
		withDisplacementStatuses(options.DisplacementStatuses).
		withFreeField1(options.FreeField1).
		withFreeField2(options.FreeField2).
		withFreeField3(options.FreeField3).
		withFreeField4(options.FreeField4).
		withFreeField5(options.FreeField5).
		withOrderBy("created_at").
		withLimit(options.Take).
		withOffset(options.Skip)
	return qry
}

func (g *getAllIndividualsSQLQuery) writeString(str string) *getAllIndividualsSQLQuery {
	_, _ = g.Builder.WriteString(str)
	return g
}

func (g *getAllIndividualsSQLQuery) writeArg(arg interface{}) *getAllIndividualsSQLQuery {
	g.a = append(g.a, arg)
	return g.writeLastArg()
}

func (g *getAllIndividualsSQLQuery) writeArgNum(i int) *getAllIndividualsSQLQuery {
	g.writeString(fmt.Sprintf("$%d", i))
	return g
}

func (g *getAllIndividualsSQLQuery) writeLastArg() *getAllIndividualsSQLQuery {
	return g.writeArgNum(len(g.a))
}

func (g *getAllIndividualsSQLQuery) writeStringArgs(sep string, args ...string) *getAllIndividualsSQLQuery {
	for i, arg := range args {
		if i > 0 {
			g.writeString(sep)
		}
		g.writeArg(arg)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withIds(ids containers.StringSet) *getAllIndividualsSQLQuery {
	if len(ids) == 0 {
		return g
	}
	g.writeString(" AND id IN (")
	for i, id := range ids.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(id)
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withFullName(name string) *getAllIndividualsSQLQuery {
	if len(name) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND (full_name LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR preferred_name LIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(")")
	} else if g.driverName == "postgres" {
		g.writeString(" AND (full_name ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(" OR preferred_name ILIKE ")
		g.writeArg("%" + name + "%")
		g.writeString(")")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withAddress(address string) *getAllIndividualsSQLQuery {
	if len(address) == 0 {
		return g
	}
	if g.driverName == "sqlite" {
		g.writeString(" AND address LIKE ").writeArg("%" + address + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND address ILIKE ").writeArg("%" + address + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withGenders(genders containers.Set[api.Gender]) *getAllIndividualsSQLQuery {
	if len(genders) == 0 {
		return g
	}
	g.writeString(" AND gender IN (")
	for i, gender := range genders.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(string(gender))
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withBirthDateFrom(from *time.Time) *getAllIndividualsSQLQuery {
	zero := &time.Time{}
	if from == nil || from.IsZero() || from == zero {
		return g
	}
	g.writeString(" AND birth_date >= ").writeArg(from)
	return g
}

func (g *getAllIndividualsSQLQuery) withAgeFrom(from *int) *getAllIndividualsSQLQuery {
	if from == nil {
		return g
	}
	g.writeString(" AND age >= ").writeArg(from)
	return g
}

func (g *getAllIndividualsSQLQuery) withAgeTo(to *int) *getAllIndividualsSQLQuery {
	if to == nil {
		return g
	}
	g.writeString(" AND age <= ").writeArg(to)
	return g
}

func (g *getAllIndividualsSQLQuery) withBirthDateTo(to *time.Time) *getAllIndividualsSQLQuery {
	zero := &time.Time{}
	if to == nil || to.IsZero() || to == zero {
		return g
	}
	g.writeString(" AND birth_date <= ").writeArg(to)
	return g
}

func (g *getAllIndividualsSQLQuery) withPhoneNumber(phoneNumber string) *getAllIndividualsSQLQuery {
	if len(phoneNumber) == 0 {
		return g
	}
	normalizedPhoneNumber := api.NormalizePhoneNumber(phoneNumber)
	if g.driverName == "sqlite" {
		g.writeString(" AND (")
		g.writeString(" normalized_phone_number_1 LIKE ").writeArg("%" + normalizedPhoneNumber + "%").writeString(" OR ")
		g.writeString(" normalized_phone_number_2 LIKE ").writeLastArg().writeString(" OR ")
		g.writeString(" normalized_phone_number_3 LIKE ").writeLastArg()
		g.writeString(")")
	} else if g.driverName == "postgres" {
		g.writeString(" AND (")
		g.writeString("normalized_phone_number_1 ILIKE ").writeArg("%" + normalizedPhoneNumber + "%").writeString(" OR ")
		g.writeString("normalized_phone_number_2 ILIKE ").writeLastArg().writeString(" OR ")
		g.writeString("normalized_phone_number_3 ILIKE ").writeLastArg()
		g.writeString(")")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withEmail(email string) *getAllIndividualsSQLQuery {
	if len(email) == 0 {
		return g
	}
	normalizedEmail := strings.ToLower(email)
	g.writeString(" AND (")
	g.writeString("email_1 = ").writeArg(normalizedEmail).writeString(" OR ")
	g.writeString("email_2 = ").writeLastArg().writeString(" OR ")
	g.writeString("email_3 = ").writeLastArg()
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withIsMinor(isMinor *bool) *getAllIndividualsSQLQuery {
	if isMinor == nil {
		return g
	}
	if *isMinor {
		g.writeString(" AND is_minor = ").writeArg(true)
	} else {
		g.writeString(" AND is_minor = ").writeArg(false)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withPresentsProtectionConcerns(presentsProtectionConcerns *bool) *getAllIndividualsSQLQuery {
	if presentsProtectionConcerns == nil {
		return g
	}
	if *presentsProtectionConcerns {
		g.writeString(" AND presents_protection_concerns = ").writeArg(true)
	} else {
		g.writeString(" AND presents_protection_concerns = ").writeArg(false)
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withCountryID(countryID string) *getAllIndividualsSQLQuery {
	if len(countryID) == 0 {
		return g
	}
	g.writeString(" AND country_id = ").writeArg(countryID)
	return g
}

func (g *getAllIndividualsSQLQuery) withDisplacementStatuses(displacementStatuses containers.Set[api.DisplacementStatus]) *getAllIndividualsSQLQuery {
	if len(displacementStatuses) == 0 {
		return g
	}
	g.writeString(" AND displacement_status IN (")
	for i, ds := range displacementStatuses.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(string(ds))
	}
	g.writeString(")")
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField1(freeField1 string) *getAllIndividualsSQLQuery {
	if len(freeField1) == 0 {
		return g
	}
	g.writeString(" AND free_field_1 = ").writeArg(freeField1)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField2(freeField2 string) *getAllIndividualsSQLQuery {
	if len(freeField2) == 0 {
		return g
	}
	g.writeString(" AND free_field_2 = ").writeArg(freeField2)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField3(freeField3 string) *getAllIndividualsSQLQuery {
	if len(freeField3) == 0 {
		return g
	}
	g.writeString(" AND free_field_3 = ").writeArg(freeField3)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField4(freeField4 string) *getAllIndividualsSQLQuery {
	if len(freeField4) == 0 {
		return g
	}
	g.writeString(" AND free_field_4 = ").writeArg(freeField4)
	return g
}

func (g *getAllIndividualsSQLQuery) withFreeField5(freeField5 string) *getAllIndividualsSQLQuery {
	if len(freeField5) == 0 {
		return g
	}
	g.writeString(" AND free_field_5 = ").writeArg(freeField5)
	return g
}

func (g *getAllIndividualsSQLQuery) withOrderBy(field string) *getAllIndividualsSQLQuery {
	if len(field) == 0 {
		return g
	}
	g.writeString(" ORDER BY ").writeString(field)
	return g
}

func (g *getAllIndividualsSQLQuery) withLimit(limit int) *getAllIndividualsSQLQuery {
	if limit == 0 {
		return g
	}
	g.writeString(fmt.Sprintf(" LIMIT %d", limit))
	return g
}

func (g *getAllIndividualsSQLQuery) withOffset(offset int) *getAllIndividualsSQLQuery {
	if offset == 0 {
		return g
	}
	g.writeString(fmt.Sprintf(" OFFSET %d", offset))
	return g
}

func (g *getAllIndividualsSQLQuery) build() (string, []interface{}) {
	return g.sql(), g.sqlArgs()
}

func (g *getAllIndividualsSQLQuery) sql() string {
	return g.Builder.String()
}

func (g *getAllIndividualsSQLQuery) sqlArgs() []interface{} {
	return g.a
}
