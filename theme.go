package igdb

// Theme type
type Theme struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetTheme gets IGDB information for a theme identified by their unique IGDB ID.
func (c *Client) GetTheme(id int, opts ...OptionFunc) (*Theme, error) {
	url, err := c.singleURL(ThemeEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var t []Theme

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return &t[0], nil
}

// GetThemes gets IGDB information for a list of themes identified by a list of their unique IGDB IDs.
func (c *Client) GetThemes(ids []int, opts ...OptionFunc) ([]*Theme, error) {
	url, err := c.multiURL(ThemeEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var t []*Theme

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchThemes searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchThemes(qry string, opts ...OptionFunc) ([]*Theme, error) {
	url, err := c.searchURL(ThemeEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var t []*Theme

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
