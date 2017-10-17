package igdb

// Title type
type Title struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         URL    `json:"url"`
	CreatedAt   int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt   int    `json:"updated_at"` // Unix time in milliseconds
	Description string `json:"description"`
	Games       []int  `json:"games"`
}

// GetTitle gets IGDB information for a title identified by their unique IGDB ID.
func (c *Client) GetTitle(id int, opts ...OptionFunc) (*Title, error) {
	url, err := c.singleURL(TitleEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var t []Title

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return &t[0], nil
}

// GetTitles gets IGDB information for a list of titles identified by a list of their unique IGDB IDs.
func (c *Client) GetTitles(ids []int, opts ...OptionFunc) ([]*Title, error) {
	url, err := c.multiURL(TitleEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var t []*Title

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchTitles searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchTitles(qry string, opts ...OptionFunc) ([]*Title, error) {
	url, err := c.searchURL(TitleEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var t []*Title

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
