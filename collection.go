package igdb

// Collection is
type Collection struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Games     []int  `json:"games"`
}

// GetCollection gets IGDB information for a collection identified by its unique IGDB ID.
func (c *Client) GetCollection(id int, opts ...OptionFunc) (*Collection, error) {
	url, err := c.singleURL(CollectionEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var col []Collection

	err = c.get(url, &col)
	if err != nil {
		return nil, err
	}

	return &col[0], nil
}

// GetCollections gets IGDB information for a list of collections identified by their
// unique IGDB IDs.
func (c *Client) GetCollections(ids []int, opts ...OptionFunc) ([]*Collection, error) {
	url, err := c.multiURL(CollectionEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var col []*Collection

	err = c.get(url, &col)
	if err != nil {
		return nil, err
	}

	return col, nil
}

// SearchCollections searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchCollections(qry string, opts ...OptionFunc) ([]*Collection, error) {
	url, err := c.searchURL(CollectionEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var col []*Collection

	err = c.get(url, &col)
	if err != nil {
		return nil, err
	}

	return col, nil
}
