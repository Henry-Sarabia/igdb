package igdb

// DateCategory code
type DateCategory int

// Year is
type Year int

// Month is
type Month int

// ReleaseDate hold information about date of release, platforms, and versions
type ReleaseDate struct {
	ID        int          `json:"id"`
	Game      int          `json:"game"`
	Category  DateCategory `json:"category"`
	Platform  int          `json:"platform"`
	Human     string       `json:"human"`
	UpdatedAt int          `json:"updated_at"` //unix epoch unspecified
	CreatedAt int          `json:"created_at"` //unix epoch unspecified
	Date      int          `json:"date"`       //unix epoch
	Region    int          `json:"region"`
	Year      Year         `json:"y"`
	Month     Month        `json:"m"`
}
