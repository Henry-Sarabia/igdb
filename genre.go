package igdb

// Genre contains information on an IGDB
// entry for a particular video game genre.
type Genre struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetGenre returns a single Genre identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will
// not have an effect due to GetGenre only returning a single Genre
// object and not a list of Genres.
func (c *Client) GetGenre(id int, opts ...OptionFunc) (*Genre, error) {
	url, err := c.singleURL(GenreEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var g []Genre

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// GetGenres returns a list of Genres identified by the provided list of IGDB
// IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of Genres based
// solely on the provided options.
func (c *Client) GetGenres(ids []int, opts ...OptionFunc) ([]*Genre, error) {
	url, err := c.multiURL(GenreEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var g []*Genre

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// SearchGenres returns a list of Genres found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchGenres(qry string, opts ...OptionFunc) ([]*Genre, error) {
	url, err := c.searchURL(GenreEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var g []*Genre

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
