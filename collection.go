package main

// Collection is
type Collection struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Games     []ID   `json:"games"`
}
