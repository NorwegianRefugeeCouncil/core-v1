package constants

import (
	_ "embed"
	"encoding/json"
)

//go:embed countries.json
var countriesJson string

type Country struct {
	// ISO3166Alpha2 is the ISO 3166-1 alpha-2 code of the country.
	ISO3166Alpha2 string `json:"iso3166Alpha2"`
	// ISO3166Alpha3 is the ISO 3166-1 alpha-3 code of the country.
	ISO3166Alpha3 string `json:"iso3166Alpha3"`
	// ISO3166Numeric is the ISO 3166-1 numeric code of the country.
	ISO3166Numeric string `json:"iso3166Numeric"`
	// ISO31662 is the ISO 3166-2 code of the country.
	ISO31662 string `json:"iso31662"`
	// Region is the region of the country.
	Region string `json:"region"`
	// RegionCode is the region code of the country.
	RegionCode string `json:"regionCode"`
	// IntermediateRegion is the intermediate region of the country.
	IntermediateRegion string `json:"intermediateRegion"`
	// SubRegion is the subregion of the country.
	SubRegion string `json:"subRegion"`
	// SubRegionCode is the subregion code of the country.
	SubRegionCode string `json:"subRegionCode"`
	// IntermediateRegionCode is the intermediate region code of the country.
	IntermediateRegionCode string `json:"intermediateRegionCode"`
	// Name is the name of the country.
	Name string `json:"name"`
}

func init() {
	err := json.Unmarshal([]byte(countriesJson), &Countries)
	if err != nil {
		panic(err)
	}
}

var Countries []Country
