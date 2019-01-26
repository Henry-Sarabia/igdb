package igdb

//go:generate gomodifytags -file $GOFILE -struct Platform -add-tags json -w

// Platform represents the hardware used to run the game
// or game delivery network.
// For more information visit: https://api-docs.igdb.com/#platform
type Platform struct {
	Abbreviation    string           `json:"abbreviation"`
	AlternativeName string           `json:"alternative_name"`
	Category        PlatformCategory `json:"category"`
	CreatedAt       int              `json:"created_at"`
	Generation      int              `json:"generation"`
	Name            string           `json:"name"`
	PlatformLogo    int              `json:"platform_logo"`
	ProductFamily   int              `json:"product_family"`
	Slug            string           `json:"slug"`
	Summary         string           `json:"summary"`
	UpdatedAt       int              `json:"updated_at"`
	URL             string           `json:"url"`
	Versions        []int            `json:"versions"`
	Websites        []int            `json:"websites"`
}

//go:generate stringer -type=PlatformCategory

type PlatformCategory int

const (
	PlatformConsole PlatformCategory = iota + 1
	PlatformArcade
	PlatformPlatform
	PlatformOperatingSystem
	PlatformPortableConsole
	PlatformComputer
)
