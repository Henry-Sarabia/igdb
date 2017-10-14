package igdb

// Keyword type
type Keyword struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetKeyword gets IGDB information for a keyword identified by its unique IGDB ID.
func (c *Client) GetKeyword(id int, opts ...OptionFunc) (*Keyword, error) {
	url, err := c.singleURL(KeywordEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var k []Keyword

	err = c.get(url, &k)
	if err != nil {
		return nil, err
	}

	return &k[0], nil
}

// GetKeywords gets IGDB information for a list of keywords identified by their
// unique IGDB IDs.
func (c *Client) GetKeywords(ids []int, opts ...OptionFunc) ([]*Keyword, error) {
	url, err := c.multiURL(KeywordEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var k []*Keyword

	err = c.get(url, &k)
	if err != nil {
		return nil, err
	}

	return k, nil
}

// SearchKeywords searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchKeywords(qry string, opts ...OptionFunc) ([]*Keyword, error) {
	url, err := c.searchURL(KeywordEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var k []*Keyword

	err = c.get(url, &k)
	if err != nil {
		return nil, err
	}

	return k, nil
}
