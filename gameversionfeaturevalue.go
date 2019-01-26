package igdb

//go:generate gomodifytags -file $GOFILE -struct GameVersionFeatureValue -add-tags json -w

// GameVersionFeatureValue represents the bool/text value of a particular feature.
// For more information visit: https://api-docs.igdb.com/#game-version-feature-value
type GameVersionFeatureValue struct {
	Game            int                     `json:"game"`
	GameFeature     int                     `json:"game_feature"`
	IncludedFeature VersionFeatureInclusion `json:"included_feature"`
	Note            string                  `json:"note"`
}

//go:generate stringer -type=VersionFeatureInclusion

type VersionFeatureInclusion int

const (
	VersionFeatureNotIncluded VersionFeatureInclusion = iota
	VersionFeatureIncluded
	VersionFeaturePreOrderOnly
)
