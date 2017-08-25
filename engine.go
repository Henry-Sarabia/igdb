package igdb

// Engine is
type Engine struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Logo      Image  `json:"logo"`
	Games     []ID   `json:"games"`
	Companies []ID   `json:"companies"`
	Platforms []ID   `json:"platforms"`
}
