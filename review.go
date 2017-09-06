package igdb

// Review type
type Review struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	URL            URL    `json:"url"`
	CreatedAt      int    `json:"created_at"` //unix epoch
	UpdatedAt      int    `json:"updated_at"` //unix epoch
	Game           int    `json:"game"`
	Category       int    `json:"category"` // Documentation is missing
	Likes          int    `json:"likes"`
	Views          int    `json:"views"`
	RatingCategory int    `json:"rating_category"` // Documenation is missing
	Platform       int    `json:"platform"`
	Video          string `json:"video"`
	Introduction   string `json:"introduction"`
	Content        string `json:"content"`
	Conclusion     string `json:"conclusion"`
	PositivePoints string `json:"positive_points"`
	NegativePoints string `json:"negative_points"`
}
