package igdb

//go:generate gomodifytags -file $GOFILE -struct GameVersion -add-tags json -w

// GameVersion provides details about game editions and versions.
// For more information visit: https://api-docs.igdb.com/#game-version
type GameVersion struct {
	CreatedAt int    `json:"created_at"`
	Features  []int  `json:"features"`
	Game      int    `json:"game"`
	Games     []int  `json:"games"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}
