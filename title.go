package igdb

// Title type
type Title struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         URL    `json:"url"`
	CreatedAt   int    `json:"created_at"` //unix epoch
	UpdatedAt   int    `json:"updated_at"` //unix epoch
	Description string `json:"description"`
	Games       []int  `json:"games"`
}

// GetTitle gets IGDB information for a title identified by their unique IGDB ID.
func (c *Client) GetTitle(id int, opts ...OptionFunc) (*Title, error) {
	url := c.singleURL(TitleEndpoint, id, opts...)

	var t []Title

	err := c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return &t[0], nil
}

// GetTitles gets IGDB information for a list of titles identified by a list of their unique IGDB IDs.
func (c *Client) GetTitles(ids []int, opts ...OptionFunc) ([]*Title, error) {
	url := c.multiURL(TitleEndpoint, ids, opts...)

	var t []*Title

	err := c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchTitles searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchTitles(qry string, opts ...OptionFunc) ([]*Title, error) {
	url := c.searchURL(TitleEndpoint, qry, opts...)

	var t []*Title

	err := c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
