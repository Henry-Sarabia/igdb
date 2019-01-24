package igdb

//go:generate gomodifytags -file $GOFILE -struct GameEngine -add-tags json -w

// GameEngine represents a video game engine such as Unreal Engine.
// For more information visit: https://api-docs.igdb.com/#game-engine
type GameEngine struct {
	Companies   []int  `json:"companies"`
	CreatedAt   int    `json:"created_at"`
	Description string `json:"description"`
	Logo        int    `json:"logo"`
	Name        string `json:"name"`
	Platforms   []int  `json:"platforms"`
	Slug        string `json:"slug"`
	UpdatedAt   int    `json:"updated_at"`
	URL         string `json:"url"`
}
