package igdb

// GenreService handles all the API
// calls for the IGDB Genre endpoint.
type GenreService service

// Genre contains information on an IGDB
// entry for a particular video game genre.
//
// For more information, visit: https://igdb.github.io/api/endpoints/genre/
type Genre struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// Get returns a single Genre identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Genres, an error is returned.
func (gs *GenreService) Get(id int, opts ...FuncOption) (*Genre, error) {
	url, err := gs.client.singleURL(GenreEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var g []Genre

	err = gs.client.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// List returns a list of Genres identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Genres based solely on the provided
// options. Any ID that does not match a Genre is ignored. If none of the IDs
// match a Genre, an error is returned.
func (gs *GenreService) List(ids []int, opts ...FuncOption) ([]*Genre, error) {
	url, err := gs.client.multiURL(GenreEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var g []*Genre

	err = gs.client.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Search returns a list of Genres found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Genres are found using the provided query, an error is returned.
func (gs *GenreService) Search(qry string, opts ...FuncOption) ([]*Genre, error) {
	url, err := gs.client.searchURL(GenreEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var g []*Genre

	err = gs.client.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Count returns the number of Genres available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which Genres to count.
func (gs *GenreService) Count(opts ...FuncOption) (int, error) {
	ct, err := gs.client.getEndpointCount(GenreEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Genre object.
func (gs *GenreService) ListFields() ([]string, error) {
	fl, err := gs.client.getEndpointFieldList(GenreEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
