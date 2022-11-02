package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"time"
)

type Individual_Field uint8

const (
	Individual_Field_ID Individual_Field = iota
	Individual_Field_CountryID
	Individual_Field_FullName
	Individual_Field_PhoneNumber
	Individual_Field_NormalizedPhoneNumber
	Individual_Field_Email
	Individual_Field_Address
	Individual_Field_BirthDate
	Individual_Field_Gender
	Individual_Field_DisplacementStatus
	Individual_Field_PreferredName
	Individual_Field_IsMinor
	Individual_Field_PresentsProtectionConcerns
	Individual_Field_PhysicalImpairment
	Individual_Field_SensoryImpairment
	Individual_Field_MentalImpairment
	Individual_Field_CreatedAt
	Individual_Field_UpdatedAt
	Individual_Field_DeletedAt
)

const _Individual_Field_name = "idcountryIdfullNamephoneNumbernormalizedPhoneNumberemailaddressbirthDategenderdisplacementStatuspreferredNameisMinorpresentsProtectionConcernsphysicalImpairmentsensoryImpairmentmentalImpairmentcreatedAtupdatedAtdeletedAt"

var _Individual_Field_index = [...]uint8{0, 2, 11, 19, 30, 51, 56, 63, 72, 78, 96, 109, 116, 142, 160, 177, 193, 202, 211, 220}

func (i Individual_Field) String() string {
	if i >= Individual_Field(len(_Individual_Field_index)-1) {
		return "Individual_Field(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Individual_Field_name[_Individual_Field_index[i]:_Individual_Field_index[i+1]]
}

func (i Individual_Field) MarshalJSON() ([]byte, error) {
	str := i.String()
	return json.Marshal(str)
}

func (i *Individual_Field) UnmarshalJSON(data []byte) error {
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return err
	}
	parsed, err := Parse_Individual_Field(strVal)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

func (i Individual_Field) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *Individual_Field) UnmarshalText(data []byte) error {
	str := string(data)
	parsed, err := Parse_Individual_Field(str)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

func All_Individual_Fields() []Individual_Field {
	return []Individual_Field{
		Individual_Field_ID,
		Individual_Field_CountryID,
		Individual_Field_FullName,
		Individual_Field_PhoneNumber,
		Individual_Field_NormalizedPhoneNumber,
		Individual_Field_Email,
		Individual_Field_Address,
		Individual_Field_BirthDate,
		Individual_Field_Gender,
		Individual_Field_DisplacementStatus,
		Individual_Field_PreferredName,
		Individual_Field_IsMinor,
		Individual_Field_PresentsProtectionConcerns,
		Individual_Field_PhysicalImpairment,
		Individual_Field_SensoryImpairment,
		Individual_Field_MentalImpairment,
		Individual_Field_CreatedAt,
		Individual_Field_UpdatedAt,
		Individual_Field_DeletedAt,
	}
}

func Parse_Individual_Field(value string) (Individual_Field, error) {
	switch value {
	case "id":
		return Individual_Field_ID, nil
	case "countryId":
		return Individual_Field_CountryID, nil
	case "fullName":
		return Individual_Field_FullName, nil
	case "phoneNumber":
		return Individual_Field_PhoneNumber, nil
	case "normalizedPhoneNumber":
		return Individual_Field_NormalizedPhoneNumber, nil
	case "email":
		return Individual_Field_Email, nil
	case "address":
		return Individual_Field_Address, nil
	case "birthDate":
		return Individual_Field_BirthDate, nil
	case "gender":
		return Individual_Field_Gender, nil
	case "displacementStatus":
		return Individual_Field_DisplacementStatus, nil
	case "preferredName":
		return Individual_Field_PreferredName, nil
	case "isMinor":
		return Individual_Field_IsMinor, nil
	case "presentsProtectionConcerns":
		return Individual_Field_PresentsProtectionConcerns, nil
	case "physicalImpairment":
		return Individual_Field_PhysicalImpairment, nil
	case "sensoryImpairment":
		return Individual_Field_SensoryImpairment, nil
	case "mentalImpairment":
		return Individual_Field_MentalImpairment, nil
	case "createdAt":
		return Individual_Field_CreatedAt, nil
	case "updatedAt":
		return Individual_Field_UpdatedAt, nil
	case "deletedAt":
		return Individual_Field_DeletedAt, nil
	default:
		return 0, fmt.Errorf("unknown Individual field %s", value)
	}
}

func Get_Individual_FieldValue(i *Individual, field Individual_Field) (interface{}, error) {
	switch field {
	case Individual_Field_ID:
		return i.ID, nil
	case Individual_Field_CountryID:
		return i.CountryID, nil
	case Individual_Field_FullName:
		return i.FullName, nil
	case Individual_Field_PhoneNumber:
		return i.PhoneNumber, nil
	case Individual_Field_NormalizedPhoneNumber:
		return i.NormalizedPhoneNumber, nil
	case Individual_Field_Email:
		return i.Email, nil
	case Individual_Field_Address:
		return i.Address, nil
	case Individual_Field_BirthDate:
		return i.BirthDate, nil
	case Individual_Field_Gender:
		return i.Gender, nil
	case Individual_Field_DisplacementStatus:
		return i.DisplacementStatus, nil
	case Individual_Field_PreferredName:
		return i.PreferredName, nil
	case Individual_Field_IsMinor:
		return i.IsMinor, nil
	case Individual_Field_PresentsProtectionConcerns:
		return i.PresentsProtectionConcerns, nil
	case Individual_Field_PhysicalImpairment:
		return i.PhysicalImpairment, nil
	case Individual_Field_SensoryImpairment:
		return i.SensoryImpairment, nil
	case Individual_Field_MentalImpairment:
		return i.MentalImpairment, nil
	case Individual_Field_CreatedAt:
		return i.CreatedAt, nil
	case Individual_Field_UpdatedAt:
		return i.UpdatedAt, nil
	case Individual_Field_DeletedAt:
		return i.DeletedAt, nil
	default:
		return nil, fmt.Errorf("unknown field %s", field)
	}
}

func Set_Individual_FieldValue(i *Individual, field Individual_Field, value interface{}) error {
	switch field {
	case Individual_Field_ID:
		if v, ok := value.(string); ok {
			i.ID = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.ID, value)
	case Individual_Field_CountryID:
		if v, ok := value.(string); ok {
			i.CountryID = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.CountryID, value)
	case Individual_Field_FullName:
		if v, ok := value.(string); ok {
			i.FullName = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.FullName, value)
	case Individual_Field_PhoneNumber:
		if v, ok := value.(string); ok {
			i.PhoneNumber = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.PhoneNumber, value)
	case Individual_Field_NormalizedPhoneNumber:
		if v, ok := value.(string); ok {
			i.NormalizedPhoneNumber = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.NormalizedPhoneNumber, value)
	case Individual_Field_Email:
		if v, ok := value.(string); ok {
			i.Email = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Email, value)
	case Individual_Field_Address:
		if v, ok := value.(string); ok {
			i.Address = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Address, value)
	case Individual_Field_BirthDate:
		if v, ok := value.(*time.Time); ok {
			i.BirthDate = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.BirthDate, value)
	case Individual_Field_Gender:
		if v, ok := value.(string); ok {
			i.Gender = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Gender, value)
	case Individual_Field_DisplacementStatus:
		if v, ok := value.(string); ok {
			i.DisplacementStatus = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.DisplacementStatus, value)
	case Individual_Field_PreferredName:
		if v, ok := value.(string); ok {
			i.PreferredName = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.PreferredName, value)
	case Individual_Field_IsMinor:
		if v, ok := value.(bool); ok {
			i.IsMinor = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.IsMinor, value)
	case Individual_Field_PresentsProtectionConcerns:
		if v, ok := value.(bool); ok {
			i.PresentsProtectionConcerns = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.PresentsProtectionConcerns, value)
	case Individual_Field_PhysicalImpairment:
		if v, ok := value.(string); ok {
			i.PhysicalImpairment = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.PhysicalImpairment, value)
	case Individual_Field_SensoryImpairment:
		if v, ok := value.(string); ok {
			i.SensoryImpairment = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.SensoryImpairment, value)
	case Individual_Field_MentalImpairment:
		if v, ok := value.(string); ok {
			i.MentalImpairment = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.MentalImpairment, value)
	case Individual_Field_CreatedAt:
		if v, ok := value.(time.Time); ok {
			i.CreatedAt = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.CreatedAt, value)
	case Individual_Field_UpdatedAt:
		if v, ok := value.(time.Time); ok {
			i.UpdatedAt = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.UpdatedAt, value)
	case Individual_Field_DeletedAt:
		if v, ok := value.(*time.Time); ok {
			i.DeletedAt = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.DeletedAt, value)
	default:
		return fmt.Errorf("unknown field %s", field)
	}
}

type Individual_Builder struct {
	ID                         string
	CountryID                  string
	FullName                   string
	PhoneNumber                string
	NormalizedPhoneNumber      string
	Email                      string
	Address                    string
	BirthDate                  *time.Time
	Gender                     string
	DisplacementStatus         string
	PreferredName              string
	IsMinor                    bool
	PresentsProtectionConcerns bool
	PhysicalImpairment         string
	SensoryImpairment          string
	MentalImpairment           string
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	DeletedAt                  *time.Time
}

func New_Individual_Builder() *Individual_Builder {
	return &Individual_Builder{}
}
func (b *Individual_Builder) WithID(value string) *Individual_Builder {
	b.ID = value
	return b
}
func (b *Individual_Builder) WithCountryID(value string) *Individual_Builder {
	b.CountryID = value
	return b
}
func (b *Individual_Builder) WithFullName(value string) *Individual_Builder {
	b.FullName = value
	return b
}
func (b *Individual_Builder) WithPhoneNumber(value string) *Individual_Builder {
	b.PhoneNumber = value
	return b
}
func (b *Individual_Builder) WithNormalizedPhoneNumber(value string) *Individual_Builder {
	b.NormalizedPhoneNumber = value
	return b
}
func (b *Individual_Builder) WithEmail(value string) *Individual_Builder {
	b.Email = value
	return b
}
func (b *Individual_Builder) WithAddress(value string) *Individual_Builder {
	b.Address = value
	return b
}
func (b *Individual_Builder) WithBirthDate(value *time.Time) *Individual_Builder {
	b.BirthDate = value
	return b
}
func (b *Individual_Builder) WithGender(value string) *Individual_Builder {
	b.Gender = value
	return b
}
func (b *Individual_Builder) WithDisplacementStatus(value string) *Individual_Builder {
	b.DisplacementStatus = value
	return b
}
func (b *Individual_Builder) WithPreferredName(value string) *Individual_Builder {
	b.PreferredName = value
	return b
}
func (b *Individual_Builder) WithIsMinor(value bool) *Individual_Builder {
	b.IsMinor = value
	return b
}
func (b *Individual_Builder) WithPresentsProtectionConcerns(value bool) *Individual_Builder {
	b.PresentsProtectionConcerns = value
	return b
}
func (b *Individual_Builder) WithPhysicalImpairment(value string) *Individual_Builder {
	b.PhysicalImpairment = value
	return b
}
func (b *Individual_Builder) WithSensoryImpairment(value string) *Individual_Builder {
	b.SensoryImpairment = value
	return b
}
func (b *Individual_Builder) WithMentalImpairment(value string) *Individual_Builder {
	b.MentalImpairment = value
	return b
}
func (b *Individual_Builder) WithCreatedAt(value time.Time) *Individual_Builder {
	b.CreatedAt = value
	return b
}
func (b *Individual_Builder) WithUpdatedAt(value time.Time) *Individual_Builder {
	b.UpdatedAt = value
	return b
}
func (b *Individual_Builder) WithDeletedAt(value *time.Time) *Individual_Builder {
	b.DeletedAt = value
	return b
}

func (b *Individual_Builder) Build() *Individual {
	ret := &Individual{
		ID:                         b.ID,
		CountryID:                  b.CountryID,
		FullName:                   b.FullName,
		PhoneNumber:                b.PhoneNumber,
		NormalizedPhoneNumber:      b.NormalizedPhoneNumber,
		Email:                      b.Email,
		Address:                    b.Address,
		BirthDate:                  b.BirthDate,
		Gender:                     b.Gender,
		DisplacementStatus:         b.DisplacementStatus,
		PreferredName:              b.PreferredName,
		IsMinor:                    b.IsMinor,
		PresentsProtectionConcerns: b.PresentsProtectionConcerns,
		PhysicalImpairment:         b.PhysicalImpairment,
		SensoryImpairment:          b.SensoryImpairment,
		MentalImpairment:           b.MentalImpairment,
		CreatedAt:                  b.CreatedAt,
		UpdatedAt:                  b.UpdatedAt,
		DeletedAt:                  b.DeletedAt,
	}
	return ret
}
