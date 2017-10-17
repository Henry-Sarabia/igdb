package igdb

// PulseSource type
type PulseSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Game int    `json:"game"`
	Page int    `json:"page"`
}

// GetPulseSource gets IGDB information for a pulse source identified by its unique IGDB ID.
func (c *Client) GetPulseSource(id int, opts ...optionFunc) (*PulseSource, error) {
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

// GetPulseSources gets IGDB information for a list of pulse sources identified by their
// unique IGDB IDs.
func (c *Client) GetPulseSources(ids []int, opts ...optionFunc) ([]*PulseSource, error) {
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

// SearchPulseSources searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchPulseSources(qry string, opts ...optionFunc) ([]*PulseSource, error) {
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
