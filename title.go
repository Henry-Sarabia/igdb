package igdb

// TitleService handles all the API
// calls for the IGDB Title endpoint.
type TitleService service

// Title contains information on an IGDB entry for
// a particular job title in the video game industry.
//
// For more information, visit: https://igdb.github.io/api/endpoints/title/
type Title struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         URL    `json:"url"`
	CreatedAt   int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt   int    `json:"updated_at"` // Unix time in milliseconds
	Description string `json:"description"`
	Games       []int  `json:"games"`
}

// Get returns a single Title identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Titles, an error is returned.
func (ts *TitleService) Get(id int, opts ...FuncOption) (*Title, error) {
	url, err := ts.client.singleURL(TitleEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var t []Title

	err = ts.client.get(url, &t)
	if err != nil {
		return nil, err
	}

	return &t[0], nil
}

// List returns a list of Titles identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of Titles based solely on the provided
// options. Any ID that does not match a Title is ignored. If none of the IDs
// match a Title, an error is returned.
func (ts *TitleService) List(ids []int, opts ...FuncOption) ([]*Title, error) {
	url, err := ts.client.multiURL(TitleEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var t []*Title

	err = ts.client.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Search returns a list of Titles found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Titles are found using the provided query, an error is returned.
func (ts *TitleService) Search(qry string, opts ...FuncOption) ([]*Title, error) {
	url, err := ts.client.searchURL(TitleEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var t []*Title

	err = ts.client.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Count returns the number of Titles available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Titles to count.
func (ts *TitleService) Count(opts ...FuncOption) (int, error) {
	ct, err := ts.client.getEndpointCount(TitleEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Title object.
func (ts *TitleService) ListFields() ([]string, error) {
	fl, err := ts.client.getEndpointFieldList(TitleEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
