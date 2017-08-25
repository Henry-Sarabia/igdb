package igdb

// ID is an unsigned 64-bit integer
type ID int

// URL is
type URL string

// Metadata stores
type Metadata struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
}

// Image is a struct that holds the ID to reach the image along with its dimensions
type Image struct {
	URL    URL    `json:"url"`
	ID     string `json:"cloudinary_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Video is a struct that holds the name of a video along with its ID.
type Video struct {
	Name string `json:"name"`
	ID   string `json:"video_id"` // Youtube slug
}
