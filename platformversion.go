package igdb

//go:generate gomodifytags -file $GOFILE -struct PlatformVersion -add-tags json -w

// PlatformVersion represents a particular version of a platform.
// For more information visit: https://api-docs.igdb.com/#platform-version
type PlatformVersion struct {
	Companies                   []int  `json:"companies"`
	Connectivity                string `json:"connectivity"`
	CPU                         string `json:"cpu"`
	Graphics                    string `json:"graphics"`
	MainManufacturer            int    `json:"main_manufacturer"`
	Media                       string `json:"media"`
	Memory                      string `json:"memory"`
	Name                        string `json:"name"`
	OS                          string `json:"os"`
	Output                      string `json:"output"`
	PlatformLogo                int    `json:"platform_logo"`
	PlatformVersionReleaseDates []int  `json:"platform_version_release_dates"`
	Resolutions                 string `json:"resolutions"`
	Slug                        string `json:"slug"`
	Sound                       string `json:"sound"`
	Storage                     string `json:"storage"`
	Summary                     string `json:"summary"`
	URL                         string `json:"url"`
}
