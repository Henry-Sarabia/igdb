package igdb

// RegionCode codes
type RegionCode int

// PlatformDate type
type PlatformDate struct {
	Date   int        `json:"date"` //unix epoch
	Region RegionCode `json:"region"`
}

// PlatformCompany type
type PlatformCompany struct {
	Company int    `json:"company"`
	Comment string `json:"comment"`
}

// PlatformVersion type
type PlatformVersion struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	OS            string            `json:"os"`
	Media         string            `json:"media"`
	Memory        string            `json:"memory"`
	Online        string            `json:"online"`
	Output        string            `json:"output"`
	Storage       string            `json:"storage"`
	Graphics      string            `json:"graphics"`
	Resolutions   string            `json:"resolutions"`
	Connectivity  string            `json:"connectivity"`
	Logo          Image             `json:"logo"`
	Summary       string            `json:"summary"`
	ReleaseDates  []PlatformDate    `json:"release_dates"`
	Developers    []PlatformCompany `json:"developers"`
	Manufacturers []PlatformCompany `json:"manufacturers"`
}

// Platform type
type Platform struct {
	ID         int               `json:"id"`
	Name       string            `json:"name"`
	Slug       string            `json:"slug"`
	URL        URL               `json:"url"`
	CreatedAt  int               `json:"created_at"` //unix epoch
	UpdatedAt  int               `json:"updated_at"` //unix epoch
	Logo       Image             `json:"logo"`
	Website    string            `json:"website"`
	Summary    string            `json:"summary"`
	AltName    string            `json:"alternative_name"`
	Generation int               `json:"generation"`
	Games      []int             `json:"games"`
	Version    []PlatformVersion `json:"version"`
}
