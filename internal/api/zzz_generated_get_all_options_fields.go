// Code generated by gen.go. DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"time"
)

type GetAllOptions_Field uint

const (
	GetAllOptions_Field_Address GetAllOptions_Field = iota
	GetAllOptions_Field_IDs
	GetAllOptions_Field_BirthDateFrom
	GetAllOptions_Field_BirthDateTo
	GetAllOptions_Field_CountryID
	GetAllOptions_Field_DisplacementStatuses
	GetAllOptions_Field_Email
	GetAllOptions_Field_FullName
	GetAllOptions_Field_Genders
	GetAllOptions_Field_IsMinor
	GetAllOptions_Field_PhoneNumber
	GetAllOptions_Field_PresentsProtectionConcerns
	GetAllOptions_Field_Skip
	GetAllOptions_Field_Take
)

const _GetAllOptions_Field_name = "addressidsbirthDateFrombirthDateTocountryIddisplacementStatusesemailfullNamegendersisMinorphoneNumberpresentsProtectionConcernsskiptake"

var _GetAllOptions_Field_index = [...]uint{0, 7, 10, 23, 34, 43, 63, 68, 76, 83, 90, 101, 127, 131, 135}

func (i GetAllOptions_Field) String() string {
	if i >= GetAllOptions_Field(len(_GetAllOptions_Field_index)-1) {
		return "GetAllOptions_Field(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GetAllOptions_Field_name[_GetAllOptions_Field_index[i]:_GetAllOptions_Field_index[i+1]]
}

func (i GetAllOptions_Field) MarshalJSON() ([]byte, error) {
	str := i.String()
	return json.Marshal(str)
}

func (i *GetAllOptions_Field) UnmarshalJSON(data []byte) error {
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return err
	}
	parsed, err := Parse_GetAllOptions_Field(strVal)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

func (i GetAllOptions_Field) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *GetAllOptions_Field) UnmarshalText(data []byte) error {
	str := string(data)
	parsed, err := Parse_GetAllOptions_Field(str)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

func All_GetAllOptions_Fields() []GetAllOptions_Field {
	return []GetAllOptions_Field{
		GetAllOptions_Field_Address,
		GetAllOptions_Field_IDs,
		GetAllOptions_Field_BirthDateFrom,
		GetAllOptions_Field_BirthDateTo,
		GetAllOptions_Field_CountryID,
		GetAllOptions_Field_DisplacementStatuses,
		GetAllOptions_Field_Email,
		GetAllOptions_Field_FullName,
		GetAllOptions_Field_Genders,
		GetAllOptions_Field_IsMinor,
		GetAllOptions_Field_PhoneNumber,
		GetAllOptions_Field_PresentsProtectionConcerns,
		GetAllOptions_Field_Skip,
		GetAllOptions_Field_Take,
	}
}

func Parse_GetAllOptions_Field(value string) (GetAllOptions_Field, error) {
	switch value {
	case "address":
		return GetAllOptions_Field_Address, nil
	case "ids":
		return GetAllOptions_Field_IDs, nil
	case "birthDateFrom":
		return GetAllOptions_Field_BirthDateFrom, nil
	case "birthDateTo":
		return GetAllOptions_Field_BirthDateTo, nil
	case "countryId":
		return GetAllOptions_Field_CountryID, nil
	case "displacementStatuses":
		return GetAllOptions_Field_DisplacementStatuses, nil
	case "email":
		return GetAllOptions_Field_Email, nil
	case "fullName":
		return GetAllOptions_Field_FullName, nil
	case "genders":
		return GetAllOptions_Field_Genders, nil
	case "isMinor":
		return GetAllOptions_Field_IsMinor, nil
	case "phoneNumber":
		return GetAllOptions_Field_PhoneNumber, nil
	case "presentsProtectionConcerns":
		return GetAllOptions_Field_PresentsProtectionConcerns, nil
	case "skip":
		return GetAllOptions_Field_Skip, nil
	case "take":
		return GetAllOptions_Field_Take, nil
	default:
		return 0, fmt.Errorf("unknown GetAllOptions field %s", value)
	}
}

func Get_GetAllOptions_FieldValue(i *GetAllOptions, field GetAllOptions_Field) (interface{}, error) {
	switch field {
	case GetAllOptions_Field_Address:
		return i.Address, nil
	case GetAllOptions_Field_IDs:
		return i.IDs, nil
	case GetAllOptions_Field_BirthDateFrom:
		return i.BirthDateFrom, nil
	case GetAllOptions_Field_BirthDateTo:
		return i.BirthDateTo, nil
	case GetAllOptions_Field_CountryID:
		return i.CountryID, nil
	case GetAllOptions_Field_DisplacementStatuses:
		return i.DisplacementStatuses, nil
	case GetAllOptions_Field_Email:
		return i.Email, nil
	case GetAllOptions_Field_FullName:
		return i.FullName, nil
	case GetAllOptions_Field_Genders:
		return i.Genders, nil
	case GetAllOptions_Field_IsMinor:
		return i.IsMinor, nil
	case GetAllOptions_Field_PhoneNumber:
		return i.PhoneNumber, nil
	case GetAllOptions_Field_PresentsProtectionConcerns:
		return i.PresentsProtectionConcerns, nil
	case GetAllOptions_Field_Skip:
		return i.Skip, nil
	case GetAllOptions_Field_Take:
		return i.Take, nil
	default:
		return nil, fmt.Errorf("unknown field %s", field)
	}
}

func Set_GetAllOptions_FieldValue(i *GetAllOptions, field GetAllOptions_Field, value interface{}) error {
	switch field {
	case GetAllOptions_Field_Address:
		if v, ok := value.(string); ok {
			i.Address = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Address, value)
	case GetAllOptions_Field_IDs:
		if v, ok := value.([]string); ok {
			i.IDs = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.IDs, value)
	case GetAllOptions_Field_BirthDateFrom:
		if v, ok := value.(*time.Time); ok {
			i.BirthDateFrom = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.BirthDateFrom, value)
	case GetAllOptions_Field_BirthDateTo:
		if v, ok := value.(*time.Time); ok {
			i.BirthDateTo = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.BirthDateTo, value)
	case GetAllOptions_Field_CountryID:
		if v, ok := value.(string); ok {
			i.CountryID = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.CountryID, value)
	case GetAllOptions_Field_DisplacementStatuses:
		if v, ok := value.([]string); ok {
			i.DisplacementStatuses = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.DisplacementStatuses, value)
	case GetAllOptions_Field_Email:
		if v, ok := value.(string); ok {
			i.Email = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Email, value)
	case GetAllOptions_Field_FullName:
		if v, ok := value.(string); ok {
			i.FullName = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.FullName, value)
	case GetAllOptions_Field_Genders:
		if v, ok := value.([]string); ok {
			i.Genders = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Genders, value)
	case GetAllOptions_Field_IsMinor:
		if v, ok := value.(*bool); ok {
			i.IsMinor = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.IsMinor, value)
	case GetAllOptions_Field_PhoneNumber:
		if v, ok := value.(string); ok {
			i.PhoneNumber = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.PhoneNumber, value)
	case GetAllOptions_Field_PresentsProtectionConcerns:
		if v, ok := value.(*bool); ok {
			i.PresentsProtectionConcerns = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.PresentsProtectionConcerns, value)
	case GetAllOptions_Field_Skip:
		if v, ok := value.(int); ok {
			i.Skip = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Skip, value)
	case GetAllOptions_Field_Take:
		if v, ok := value.(int); ok {
			i.Take = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Take, value)
	default:
		return fmt.Errorf("unknown field %s", field)
	}
}

type GetAllOptions_Builder struct {
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

func New_GetAllOptions_Builder() *GetAllOptions_Builder {
	return &GetAllOptions_Builder{}
}
func (b *GetAllOptions_Builder) WithAddress(value string) *GetAllOptions_Builder {
	b.Address = value
	return b
}
func (b *GetAllOptions_Builder) WithIDs(value []string) *GetAllOptions_Builder {
	b.IDs = value
	return b
}
func (b *GetAllOptions_Builder) WithBirthDateFrom(value *time.Time) *GetAllOptions_Builder {
	b.BirthDateFrom = value
	return b
}
func (b *GetAllOptions_Builder) WithBirthDateTo(value *time.Time) *GetAllOptions_Builder {
	b.BirthDateTo = value
	return b
}
func (b *GetAllOptions_Builder) WithCountryID(value string) *GetAllOptions_Builder {
	b.CountryID = value
	return b
}
func (b *GetAllOptions_Builder) WithDisplacementStatuses(value []string) *GetAllOptions_Builder {
	b.DisplacementStatuses = value
	return b
}
func (b *GetAllOptions_Builder) WithEmail(value string) *GetAllOptions_Builder {
	b.Email = value
	return b
}
func (b *GetAllOptions_Builder) WithFullName(value string) *GetAllOptions_Builder {
	b.FullName = value
	return b
}
func (b *GetAllOptions_Builder) WithGenders(value []string) *GetAllOptions_Builder {
	b.Genders = value
	return b
}
func (b *GetAllOptions_Builder) WithIsMinor(value *bool) *GetAllOptions_Builder {
	b.IsMinor = value
	return b
}
func (b *GetAllOptions_Builder) WithPhoneNumber(value string) *GetAllOptions_Builder {
	b.PhoneNumber = value
	return b
}
func (b *GetAllOptions_Builder) WithPresentsProtectionConcerns(value *bool) *GetAllOptions_Builder {
	b.PresentsProtectionConcerns = value
	return b
}
func (b *GetAllOptions_Builder) WithSkip(value int) *GetAllOptions_Builder {
	b.Skip = value
	return b
}
func (b *GetAllOptions_Builder) WithTake(value int) *GetAllOptions_Builder {
	b.Take = value
	return b
}

func (b *GetAllOptions_Builder) Build() *GetAllOptions {
	ret := &GetAllOptions{
		Address:                    b.Address,
		IDs:                        b.IDs,
		BirthDateFrom:              b.BirthDateFrom,
		BirthDateTo:                b.BirthDateTo,
		CountryID:                  b.CountryID,
		DisplacementStatuses:       b.DisplacementStatuses,
		Email:                      b.Email,
		FullName:                   b.FullName,
		Genders:                    b.Genders,
		IsMinor:                    b.IsMinor,
		PhoneNumber:                b.PhoneNumber,
		PresentsProtectionConcerns: b.PresentsProtectionConcerns,
		Skip:                       b.Skip,
		Take:                       b.Take,
	}
	return ret
}
