package igdb

// PlatformService handles all the API
// calls for the IGDB Platform endpoint.
type PlatformService service

// Platform contains information on an IGDB entry for the particular
// hardware used to run a game or game delivery network.
//
// For more information, visit: https://igdb.github.io/api/endpoints/platform/
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

// PlatformVersion contains information on an IGDB
// entry for a particular version of a Platform.
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

// PlatformDate contains information on the
// release date for a particular PlatformVersion.
type PlatformDate struct {
	Date   int        `json:"date"` // Unix time in milliseconds
	Region RegionCode `json:"region"`
}

// PlatformCompany contains information on an IGDB entry
// for a company that worked on a particular platform.
type PlatformCompany struct {
	Company int    `json:"company"`
	Comment string `json:"comment"`
}

// Get returns a single Platform identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Platforms, an error is returned.
func (ps *PlatformService) Get(id int, opts ...FuncOption) (*Platform, error) {
	url, err := ps.client.singleURL(PlatformEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var p []Platform

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// List returns a list of Platforms identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Platforms based solely on the provided
// options. Any ID that does not match a Platform is ignored. If none of the IDs
// match a Platform, an error is returned.
func (ps *PlatformService) List(ids []int, opts ...FuncOption) ([]*Platform, error) {
	url, err := ps.client.multiURL(PlatformEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var p []*Platform

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Search returns a list of Platforms found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Platforms are found using the provided query, an error is returned.
func (ps *PlatformService) Search(qry string, opts ...FuncOption) ([]*Platform, error) {
	url, err := ps.client.searchURL(PlatformEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var p []*Platform

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Count returns the number of Platforms available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Platforms to count.
func (ps *PlatformService) Count(opts ...FuncOption) (int, error) {
	ct, err := ps.client.getEndpointCount(PlatformEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Platform object.
func (ps *PlatformService) ListFields() ([]string, error) {
	fl, err := ps.client.getEndpointFieldList(PlatformEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
