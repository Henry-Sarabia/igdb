package igdb

//go:generate gomodifytags -file $GOFILE -struct Genre -add-tags json -w

// Genre represents the genre of a particular video game.
// For more information visit: https://api-docs.igdb.com/#genre
type Genre struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}
