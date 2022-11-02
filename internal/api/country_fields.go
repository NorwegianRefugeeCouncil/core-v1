package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Country_Field uint8

const (
	Country_Field_ID Country_Field = iota
	Country_Field_Code
	Country_Field_Name
	Country_Field_JwtGroup
)

const _Country_Field_name = "idcodenamejwtGroup"

var _Country_Field_index = [...]uint8{0, 2, 6, 10, 18}

func (i Country_Field) String() string {
	if i >= Country_Field(len(_Country_Field_index)-1) {
		return "Country_Field(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Country_Field_name[_Country_Field_index[i]:_Country_Field_index[i+1]]
}

func (i Country_Field) MarshalJSON() ([]byte, error) {
	str := i.String()
	return json.Marshal(str)
}

func (i *Country_Field) UnmarshalJSON(data []byte) error {
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return err
	}
	parsed, err := Parse_Country_Field(strVal)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

func (i Country_Field) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *Country_Field) UnmarshalText(data []byte) error {
	str := string(data)
	parsed, err := Parse_Country_Field(str)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

func All_Country_Fields() []Country_Field {
	return []Country_Field{
		Country_Field_ID,
		Country_Field_Code,
		Country_Field_Name,
		Country_Field_JwtGroup,
	}
}

func Parse_Country_Field(value string) (Country_Field, error) {
	switch value {
	case "id":
		return Country_Field_ID, nil
	case "code":
		return Country_Field_Code, nil
	case "name":
		return Country_Field_Name, nil
	case "jwtGroup":
		return Country_Field_JwtGroup, nil
	default:
		return 0, fmt.Errorf("unknown Country field %s", value)
	}
}

func Get_Country_FieldValue(i *Country, field Country_Field) (interface{}, error) {
	switch field {
	case Country_Field_ID:
		return i.ID, nil
	case Country_Field_Code:
		return i.Code, nil
	case Country_Field_Name:
		return i.Name, nil
	case Country_Field_JwtGroup:
		return i.JwtGroup, nil
	default:
		return nil, fmt.Errorf("unknown field %s", field)
	}
}

func Set_Country_FieldValue(i *Country, field Country_Field, value interface{}) error {
	switch field {
	case Country_Field_ID:
		if v, ok := value.(string); ok {
			i.ID = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.ID, value)
	case Country_Field_Code:
		if v, ok := value.(string); ok {
			i.Code = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Code, value)
	case Country_Field_Name:
		if v, ok := value.(string); ok {
			i.Name = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.Name, value)
	case Country_Field_JwtGroup:
		if v, ok := value.(string); ok {
			i.JwtGroup = v
			return nil
		}
		return fmt.Errorf("invalid value type for field %s. Expected %T but got %T", field, i.JwtGroup, value)
	default:
		return fmt.Errorf("unknown field %s", field)
	}
}

type Country_Builder struct {
	ID       string
	Code     string
	Name     string
	JwtGroup string
}

func New_Country_Builder() *Country_Builder {
	return &Country_Builder{}
}
func (b *Country_Builder) WithID(value string) *Country_Builder {
	b.ID = value
	return b
}
func (b *Country_Builder) WithCode(value string) *Country_Builder {
	b.Code = value
	return b
}
func (b *Country_Builder) WithName(value string) *Country_Builder {
	b.Name = value
	return b
}
func (b *Country_Builder) WithJwtGroup(value string) *Country_Builder {
	b.JwtGroup = value
	return b
}

func (b *Country_Builder) Build() *Country {
	ret := &Country{
		ID:       b.ID,
		Code:     b.Code,
		Name:     b.Name,
		JwtGroup: b.JwtGroup,
	}
	return ret
}
