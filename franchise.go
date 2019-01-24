package igdb

//go:generate gomodifytags -file $GOFILE -struct Franchise -add-tags json -w

// Franchise is a list of video game franchises such as Star Wars.
// For more information visit: https://api-docs.igdb.com/#franchise
type Franchise struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	Url       string `json:"url"`
}
