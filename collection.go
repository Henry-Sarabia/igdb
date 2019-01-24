package igdb

//go:generate gomodifytags -file $GOFILE -struct Collection -add-tags json -w

// Collection represents a video game series.
// For more information visit: https://api-docs.igdb.com/#collection
type Collection struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}
