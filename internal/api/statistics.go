package api

type CountryStatistics struct {
	TotalCount            int                       `json:"totalCount"`
	ChildrenCount         int                       `json:"childrenCount"`
	WomenAndChildrenCount int                       `json:"womenAndChildrenCount"`
	ElderlyCount          int                       `json:"elderlyCount"`
	ByGenderAge           map[string]map[string]int `json:"genderAge"`
	ByDisplacementStatus  map[string]int            `json:"displacementStatus"`
	ByPhysicalImpairment  map[string]int            `json:"physicalImpairment"`
	ByMentalImpairment    map[string]int            `json:"mentalImpairment"`
	BySensoryImpairment   map[string]int            `json:"sensoryImpairment"`
}
