package igdb

// PlatformDate contains information
// on the release date for a particular
// PlatformVersion.
type PlatformDate struct {
	Date   int        `json:"date"` // Unix time in milliseconds
	Region RegionCode `json:"region"`
}

// PlatformCompany contains information
// on an IGDB entry for a company that
// worked on a particular platform.
type PlatformCompany struct {
	Company int    `json:"company"`
	Comment string `json:"comment"`
}

// PlatformVersion contains information on an
// IGDB entry for a particular version of a
// Platform.
type PlatformVersion struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	URL           URL               `json:"url"`
	Manufacturer  string            `json:"manufacturer"`
	OS            string            `json:"os"`
	CPU           string            `json:"cpu"`
	Media         string            `json:"media"`
	Memory        string            `json:"memory"`
	Online        string            `json:"online"`
	Output        string            `json:"output"`
	Storage       string            `json:"storage"`
	Graphics      string            `json:"graphics"`
	Resolutions   string            `json:"resolutions"`
	Connectivity  string            `json:"connectivity"`
	Sound         string            `json:"sound"`
	Logo          Image             `json:"logo"`
	Summary       string            `json:"summary"`
	ReleaseDates  []PlatformDate    `json:"release_dates"`
	Developers    []PlatformCompany `json:"developers"`
	Manufacturers []PlatformCompany `json:"manufacturers"`
}

// Platform contains information on an IGDB entry
// for the particular hardware used to run a game
// or game delivery network.
type Platform struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	URL           URL               `json:"url"`
	CreatedAt     int               `json:"created_at"` // Unix time in milliseconds
	UpdatedAt     int               `json:"updated_at"` // Unix time in milliseconds
	Shortcut      string            `json:"shortcut"`
	Logo          Image             `json:"logo"`
	Website       string            `json:"website"`
	Summary       string            `json:"summary"`
	AltName       string            `json:"alternative_name"`
	Generation    int               `json:"generation"`
	Category      int               `json:"category"`
	ProductFamily int               `json:"product_family"`
	Games         []int             `json:"games"`
	Versions      []PlatformVersion `json:"versions"`
}

// GetPlatform returns a single Platform identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not 
// have an effect due to GetPlatform only returning a single Platform object
// and not a list of Platforms.
func (c *Client) GetPlatform(id int, opts ...OptionFunc) (*Platform, error) {
	url, err := c.singleURL(PlatformEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []Platform

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPlatforms returns a list of Platforms identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of Platforms based
// solely on the provided options.
func (c *Client) GetPlatforms(ids []int, opts ...OptionFunc) ([]*Platform, error) {
	url, err := c.multiURL(PlatformEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Platform

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPlatforms returns a list of Platforms found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchPlatforms(qry string, opts ...OptionFunc) ([]*Platform, error) {
	url, err := c.searchURL(PlatformEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Platform

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
