package igdb

//go:generate gomodifytags -file $GOFILE -struct GameMode -add-tags json -w

// GameMode represents a video game mode such as single or multi player.
// For more information visit: https://api-docs.igdb.com/#game-mode
type GameMode struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}
