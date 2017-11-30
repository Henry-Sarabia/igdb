package igdb

// FranchiseService handles all the API
// calls for the IGDB Franchise endpoint.
type FranchiseService service

// Franchise contains information on an IGDB
// entry for a particular video game franchise.
type Franchise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// Get returns a single Franchise identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Franchises, an error is returned.
func (fs *FranchiseService) Get(id int, opts ...OptionFunc) (*Franchise, error) {
	url, err := fs.client.singleURL(FranchiseEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var f []Franchise

	err = fs.client.get(url, &f)
	if err != nil {
		return nil, err
	}

	return &f[0], nil
}

// List returns a list of Franchises identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Franchises based solely on the provided
// options. Any ID that does not match a Franchise is ignored. If none of the IDs
// match a Franchise, an error is returned.
func (fs *FranchiseService) List(ids []int, opts ...OptionFunc) ([]*Franchise, error) {
	url, err := fs.client.multiURL(FranchiseEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var f []*Franchise

	err = fs.client.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Search returns a list of Franchises found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Franchises are found using the provided query, an error is returned.
func (fs *FranchiseService) Search(qry string, opts ...OptionFunc) ([]*Franchise, error) {
	url, err := fs.client.searchURL(FranchiseEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var f []*Franchise

	err = fs.client.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Count returns the number of Franchises available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which Franchises to count.
func (fs *FranchiseService) Count(opts ...OptionFunc) (int, error) {
	ct, err := fs.client.getEndpointCount(FranchiseEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Franchise object.
func (fs *FranchiseService) ListFields() ([]string, error) {
	fl, err := fs.client.getEndpointFieldList(FranchiseEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
