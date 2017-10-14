package igdb

// Franchise is
type Franchise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetFranchise gets IGDB information for a franchise identified by its unique IGDB ID.
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

// GetFranchises gets IGDB information for a list of franchises identified by their
// unique IGDB IDs.
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

// SearchFranchises searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
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
