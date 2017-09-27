package igdb

// PulseSource type
type PulseSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Game int    `json:"game"`
	Page int    `json:"page"`
}

// GetPulseSource gets IGDB information for a pulse source identified by its unique IGDB ID.
func (c *Client) GetPulseSource(id int, opts ...OptionFunc) (*PulseSource, error) {
	url := c.singleURL(PulseSourceEndpoint, id, opts...)

	var p []PulseSource

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPulseSources gets IGDB information for a list of pulse sources identified by their
// unique IGDB IDs.
func (c *Client) GetPulseSources(ids []int, opts ...OptionFunc) ([]*PulseSource, error) {
	url := c.multiURL(PulseSourceEndpoint, ids, opts...)

	var p []*PulseSource

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPulseSources searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchPulseSources(qry string, opts ...OptionFunc) ([]*PulseSource, error) {
	url := c.searchURL(PulseSourceEndpoint, qry, opts...)

	var p []*PulseSource

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
