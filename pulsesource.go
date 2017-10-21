package igdb

// PulseSource contains information
// on an IGDB entry for a specific
// news source.
type PulseSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Game int    `json:"game"`
	Page int    `json:"page"`
}

// GetPulseSource returns a single PulseSource identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have an
// effect due to GetPulseSource only returning a single PulseSource object and not
// a list of PulseSources.
func (c *Client) GetPulseSource(id int, opts ...OptionFunc) (*PulseSource, error) {
	url, err := c.singleURL(PulseSourceEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []PulseSource

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPulseSources returns a list of PulseSources identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of PulseSources based
// solely on the provided options.
func (c *Client) GetPulseSources(ids []int, opts ...OptionFunc) ([]*PulseSource, error) {
	url, err := c.multiURL(PulseSourceEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*PulseSource

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPulseSources returns a list of PulseSources found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchPulseSources(qry string, opts ...OptionFunc) ([]*PulseSource, error) {
	url, err := c.searchURL(PulseSourceEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*PulseSource

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
