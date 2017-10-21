package igdb

// Franchise contains information on an
// IGDB entry for a particular video game
// franchise.
type Franchise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetFranchise returns a single Franchise identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetFranchise only returning a single Franchise object and
// not a list of Franchises.
func (c *Client) GetFranchise(id int, opts ...OptionFunc) (*Franchise, error) {
	url, err := c.singleURL(FranchiseEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var f []Franchise

	err = c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return &f[0], nil
}

// GetFranchises returns a list of Franchises identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of Franchises based
// solely on the provided options.
func (c *Client) GetFranchises(ids []int, opts ...OptionFunc) ([]*Franchise, error) {
	url, err := c.multiURL(FranchiseEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var f []*Franchise

	err = c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// SearchFranchises returns a list of Franchises found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchFranchises(qry string, opts ...OptionFunc) ([]*Franchise, error) {
	url, err := c.searchURL(FranchiseEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var f []*Franchise

	err = c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
