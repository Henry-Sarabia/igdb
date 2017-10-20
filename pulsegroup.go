package igdb

// PulseGroup contains information on an
// IGDB entry for a group of news articles.
type PulseGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         URL    `json:"url"`
	CreatedAt   int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt   int    `json:"updated_at"` // Unix time in milliseconds
	PublishedAt int    `json:"published_at"`
	Category    int    `json:"category"`
	Tags        []Tag  `json:"tags"`
	Pulses      []int  `json:"pulses"`
	Game        int    `json:"game"`
}

// GetPulseGroup gets IGDB information for a PulseGroup identified by its unique
// IGDB ID. GetPulseGroup returns a single PulseGroup identified by the provided
// IGDB ID. Functional options may be provided but sorting and pagination will
// not have an effect due to GetPulseGroup only returning a single PulseGroup
// object and not a list of PulseGroups.
func (c *Client) GetPulseGroup(id int, opts ...OptionFunc) (*PulseGroup, error) {
	url, err := c.singleURL(PulseGroupEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []PulseGroup

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPulseGroups returns a list of PulseGroups identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of PulseGroups based
// solely on the provided options.
func (c *Client) GetPulseGroups(ids []int, opts ...OptionFunc) ([]*PulseGroup, error) {
	url, err := c.multiURL(PulseGroupEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*PulseGroup

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPulseGroups returns a list of PulseGroups found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchPulseGroups(qry string, opts ...OptionFunc) ([]*PulseGroup, error) {
	url, err := c.searchURL(PulseGroupEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*PulseGroup

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
