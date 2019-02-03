package igdb

//go:generate gomodifytags -file $GOFILE -struct GameVersionFeature -add-tags json -w

// GameVersionFeature represents features and descriptions of what makes
// each version/edition different from their main game.
// For more information visit: https://api-docs.igdb.com/#game-version-feature
type GameVersionFeature struct {
	Category    VersionFeatureCategory `json:"category"`
	Description string                 `json:"description"`
	Position    int                    `json:"position"`
	Title       string                 `json:"title"`
	Values      []int                  `json:"values"`
}

//go:generate stringer -type=VersionFeatureCategory

type VersionFeatureCategory int

const (
	VersionFeatureBoolean VersionFeatureCategory = iota
	VersionFeatureDescription
)
