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
		writeString("SELECT * FROM individuals WHERE deleted_at IS NULL").
		withIds(options.IDs).
		withFullName(options.FullName).
		withAddress(options.Address).
		withGenders(options.Genders).
		withBirthDateFrom(options.BirthDateFrom).
		withBirthDateTo(options.BirthDateTo).
		withPhoneNumber(options.PhoneNumber).
		withEmail(options.Email).
		withIsMinor(options.IsMinor).
		withPresentsProtectionConcerns(options.PresentsProtectionConcerns).
		withCountryID(options.CountryID).
		withDisplacementStatuses(options.DisplacementStatuses).
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
	g.writeString(fmt.Sprintf("$%d", len(g.a)+1))
	g.a = append(g.a, arg)
	return g
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
		g.writeString(" AND normalized_phone_number LIKE ").writeArg("%" + normalizedPhoneNumber + "%")
	} else if g.driverName == "postgres" {
		g.writeString(" AND normalized_phone_number ILIKE ").writeArg("%" + normalizedPhoneNumber + "%")
	}
	return g
}

func (g *getAllIndividualsSQLQuery) withEmail(email string) *getAllIndividualsSQLQuery {
	if len(email) == 0 {
		return g
	}
	normalizedEmail := strings.ToLower(email)
	g.writeString(" AND email = ").writeArg(normalizedEmail)
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
	g.writeString(" AND gender IN (")
	for i, ds := range displacementStatuses.Items() {
		if i != 0 {
			g.writeString(",")
		}
		g.writeArg(string(ds))
	}
	g.writeString(")")
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
