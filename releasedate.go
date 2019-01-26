package igdb

//go:generate gomodifytags -file $GOFILE -struct ReleaseDate -add-tags json -w

// ReleaseDate represents the release date for a particular game.
// Used to dig deeper into release dates, platforms, and versions.
// For more information visit: https://api-docs.igdb.com/#release-date
type ReleaseDate struct {
	Category  DateCategory   `json:"category"`
	CreatedAt int            `json:"created_at"`
	Date      int            `json:"date"`
	Game      int            `json:"game"`
	Human     string         `json:"human"`
	M         int            `json:"m"`
	Platform  int            `json:"platform"`
	Region    RegionCategory `json:"region"`
	UpdatedAt int            `json:"updated_at"`
	Y         int            `json:"y"`
}

//go:generate stringer -type=DateCategory,RegionCategory

// DateCategory represents the IGDB enumerated type Date Category which
// describes the format of a release date. Use the Stringer interface to
// access the corresponding Date Category values as strings.
type DateCategory int

const (
	DateYYYYMMMMDD DateCategory = iota
	DateYYYYMMMM
	DateYYYY
	DateYYYYQ1
	DateYYYYQ2
	DateYYYYQ3
	DateYYYYQ4
	DateTBD
)

type RegionCategory int

const (
	RegionEurope RegionCategory = iota + 1
	RegionNorthAmerica
	RegionAustralia
	RegionNewZealand
	RegionJapan
	RegionChina
	RegionAsia
	RegionWorldwide
)
