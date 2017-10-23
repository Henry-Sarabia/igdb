package igdb

// Collection contains information on an
// IGDB entry for a particular video game
// series.
type Collection struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetCollection returns a single Collection identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetCollection only returning a single Collection object and
// not a list of Collections.
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

// GetCollections returns a list of Collections identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
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

// SearchCollections returns a list of Collections found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
// Providing an empty query will instead retrieve an index of Collections based solely on
// the provided options.
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
