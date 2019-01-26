package igdb

//go:generate gomodifytags -file $GOFILE -struct Title -add-tags json -w

// Title represents a particular job title in the game industry.
// For more information visit: https://api-docs.igdb.com/#title
type Title struct {
	CreatedAt   int    `json:"created_at"`
	Description string `json:"description"`
	Games       []int  `json:"games"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	UpdatedAt   int    `json:"updated_at"`
	URL         string `json:"url"`
}
