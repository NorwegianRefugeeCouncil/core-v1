package zanzibar

type LocationType string

const (
	LocationType_Global  LocationType = "global"
	LocationType_Region  LocationType = "region"
	LocationType_Country LocationType = "country"
	LocationType_Area    LocationType = "area"
)

func (x LocationType) String() string {
	return x.String()
}
