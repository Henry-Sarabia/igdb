package igdb

import "strconv"

// ID is an unsigned 64-bit integer
type ID int

// URL is
type URL string

// StatusCode codes
type StatusCode int

// Tag codes??
type Tag int //uint32

// AltName is
type AltName struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
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

// toString simply returns the string equivalent of a given ID.
func (id ID) string() string {
	return strconv.Itoa(int(id))
}

// idsString returns the slice of string equivalents of a given
// list of IDs.
func idsString(ids []ID) []string {
	var str []string
	for _, id := range ids {
		str = append(str, id.string())
	}
	return str
}
