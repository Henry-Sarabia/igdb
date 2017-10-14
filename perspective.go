package igdb

// Perspective type
type Perspective struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetPerspective gets IGDB information for a player perspective identified by its unique IGDB ID.
func (c *Client) GetPerspective(id int, opts ...OptionFunc) (*Perspective, error) {
	url, err := c.singleURL(PerspectiveEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []Perspective

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPerspectives gets IGDB information for a list of player perspectives identified by their
// unique IGDB IDs.
func (c *Client) GetPerspectives(ids []int, opts ...OptionFunc) ([]*Perspective, error) {
	url, err := c.multiURL(PerspectiveEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Perspective

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPerspectives searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchPerspectives(qry string, opts ...OptionFunc) ([]*Perspective, error) {
	url, err := c.searchURL(PerspectiveEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Perspective

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
