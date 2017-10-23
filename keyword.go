package igdb

// Keyword contains information on an IGDB
// entry for a particular keyword. Keywords
// are words or phrases that get tagged to
// a game (e.g. "world war 2" or "steampunk").
type Keyword struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetKeyword returns a single Keyword identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not
// have an effect due to GetKeyword only returning a single Keyword object
// and not a list of Keywords.
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

// GetKeywords returns a list of Keywords identified by the provided list of IGDB
// IDs. Provide functional options to filter, sort, and paginate the results.
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

// SearchKeywords returns a list of Keywords found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
// Providing an empty query will instead retrieve an index of Keywords based solely on
// the provided options.
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
