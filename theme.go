package igdb

//go:generate gomodifytags -file $GOFILE -struct Theme -add-tags json -w

// Theme represents a particular video game theme.
// For more information visit: https://api-docs.igdb.com/#theme
type Theme struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}
