package igdb

// ThemeService handles all the API
// calls for the IGDB Theme endpoint.
type ThemeService service

// Theme contains information on an IGDB entry for a
// particular video game theme (e.g. Fantasy or Horror).
//
// For more information, visit: https://igdb.github.io/api/endpoints/theme/
type Theme struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// Get returns a single Theme identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Themes, an error is returned.
func (ts *ThemeService) Get(id int, opts ...FuncOption) (*Theme, error) {
	url, err := ts.client.singleURL(ThemeEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var th []Theme

	err = ts.client.get(url, &th)
	if err != nil {
		return nil, err
	}

	return &th[0], nil
}

// List returns a list of Themes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Themes based solely on the provided
// options. Any ID that does not match a Theme is ignored. If none of the IDs
// match a Theme, an error is returned.
func (ts *ThemeService) List(ids []int, opts ...FuncOption) ([]*Theme, error) {
	url, err := ts.client.multiURL(ThemeEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var th []*Theme

	err = ts.client.get(url, &th)
	if err != nil {
		return nil, err
	}

	return th, nil
}

// Search returns a list of Themes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Themes are found using the provided query, an error is returned.
func (ts *ThemeService) Search(qry string, opts ...FuncOption) ([]*Theme, error) {
	url, err := ts.client.searchURL(ThemeEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var th []*Theme

	err = ts.client.get(url, &th)
	if err != nil {
		return nil, err
	}

	return th, nil
}

// Count returns the number of Themes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Themes to count.
func (ts *ThemeService) Count(opts ...FuncOption) (int, error) {
	ct, err := ts.client.getEndpointCount(ThemeEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Theme object.
func (ts *ThemeService) ListFields() ([]string, error) {
	fl, err := ts.client.getEndpointFieldList(ThemeEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
