package igdb

// Title type
type Title struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         URL    `json:"url"`
	CreatedAt   int    `json:"created_at"` //unix epoch
	UpdatedAt   int    `json:"updated_at"` //unix epoch
	Description string `json:"description"`
	Games       []int  `json:"games"`
}
