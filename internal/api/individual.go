package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
)

type Individual struct {
	ID                         string             `json:"id" db:"id"`
	CountryID                  string             `json:"countryId" db:"country_id"`
	FullName                   string             `json:"fullName" db:"full_name"`
	PhoneNumber                string             `json:"phoneNumber" db:"phone_number"`
	NormalizedPhoneNumber      string             `json:"normalizedPhoneNumber" db:"normalized_phone_number"`
	Email                      string             `json:"email" db:"email"`
	Address                    string             `json:"address" db:"address"`
	BirthDate                  *time.Time         `json:"birthDate" db:"birth_date"`
	Gender                     Gender             `json:"gender" db:"gender"`
	DisplacementStatus         DisplacementStatus `json:"displacementStatus" db:"displacement_status"`
	PreferredName              string             `json:"preferredName" db:"preferred_name"`
	IsMinor                    bool               `json:"isMinor" db:"is_minor"`
	PresentsProtectionConcerns bool               `json:"presentsProtectionConcerns" db:"presents_protection_concerns"`
	PhysicalImpairment         string             `json:"physicalImpairment" db:"physical_impairment"`
	SensoryImpairment          string             `json:"sensoryImpairment" db:"sensory_impairment"`
	MentalImpairment           string             `json:"mentalImpairment" db:"mental_impairment"`
	CreatedAt                  time.Time          `json:"createdAt" db:"created_at"`
	UpdatedAt                  time.Time          `json:"updatedAt" db:"updated_at"`
	DeletedAt                  *time.Time         `json:"deletedAt" db:"deleted_at"`
}

type IndividualList struct {
	Items []*Individual `json:"items"`
}

func (i *Individual) GetFieldValue(field string) (interface{}, error) {
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
	case constants.DBColumnIndividualCreatedAt:
		return i.CreatedAt, nil
	case constants.DBColumnIndividualUpdatedAt:
		return i.UpdatedAt, nil
	case constants.DBColumnIndividualDeletedAt:
		return i.DeletedAt, nil
	default:
		return nil, fmt.Errorf("unknown field: %s", field)
	}
}

func (i *Individual) String() string {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func (i *Individual) Normalize() {
	i.ID = trimString(i.ID)
	i.FullName = trimString(i.FullName)
	i.PreferredName = trimString(i.PreferredName)
	if i.PreferredName == "" {
		i.PreferredName = i.FullName
	}
	i.Email = trimString(normalizeEmail(i.Email))
	i.PhoneNumber = trimString(i.PhoneNumber)
	i.Address = trimString(i.Address)
	i.NormalizedPhoneNumber = NormalizePhoneNumber(i.PhoneNumber)
	i.PhysicalImpairment = trimString(i.PhysicalImpairment)
	i.MentalImpairment = trimString(i.MentalImpairment)
	i.SensoryImpairment = trimString(i.SensoryImpairment)
}
