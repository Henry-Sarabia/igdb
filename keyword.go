package igdb

//go:generate gomodifytags -file $GOFILE -struct Keyword -add-tags json -w

// Keyword represents a word or phrase that get tagged to a game
// such as "World War 2" or "Steampunk".
// For more information visit: https://api-docs.igdb.com/#keyword
type Keyword struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	Url       string `json:"url"`
}
