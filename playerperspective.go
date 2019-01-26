package igdb

//go:generate gomodifytags -file $GOFILE -struct PlayerPerspective -add-tags json -w

// PlayerPerspective describes the view or perspective of the player in a video game.
// For more information visit: https://api-docs.igdb.com/#player-perspective
type PlayerPerspective struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}
