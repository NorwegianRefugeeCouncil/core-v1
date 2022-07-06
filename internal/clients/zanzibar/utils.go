package zanzibar

type LocationType int

const (
	LocationType_Global LocationType = iota
	LocationType_Region
	LocationType_Country
	LocationType_Area
)

func (l LocationType) String() string {
	return [...]string{"global", "region", "country", "area"}[l]
}
