package igdb

// Title contains information on an
// IGDB entry for a particular job
// title in the video game industry.
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

// GetTitle gets IGDB information for a Title identified by its unique
// IGDB ID. GetTitle returns a single Title identified by the provided
// IGDB ID. Functional options may be provided but sorting and pagination
// will not have an effect due to GetTitle only returning a single Title
// object and not a list of Titles.
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

// GetTitles returns a list of Titles identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the
// results. Providing an empty list of IDs will instead retrieve an index
// of Titles based solely on the provided options.
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

// SearchTitles returns a list of Titles found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
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
