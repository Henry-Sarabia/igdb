package igdb

//go:generate gomodifytags -file $GOFILE -struct PlatformVersionReleaseDate -add-tags json -w

// PlatformVersionReleaseDate describes a platform release date.
// Used to dig deeper into release dates, platforms, and versions.
// For more information visit: https://api-docs.igdb.com/#platform-version-release-date
type PlatformVersionReleaseDate struct {
	Category        DateCategory   `json:"category"`
	CreatedAt       int            `json:"created_at"`
	Date            int            `json:"date"`
	Human           string         `json:"human"`
	M               int            `json:"m"`
	PlatformVersion int            `json:"platform_version"`
	Region          RegionCategory `json:"region"`
	UpdatedAt       int            `json:"updated_at"`
	Y               int            `json:"y"`
}
