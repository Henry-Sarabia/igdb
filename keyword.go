package igdb

// KeywordService handles all the API
// calls for the IGDB Keyword endpoint.
type KeywordService service

// Keyword contains information on an IGDB entry for a particular keyword.
// Keywords are words or phrases that get tagged to a game
// (e.g. "world war 2" or "steampunk").
//
// For more information, visit: https://igdb.github.io/api/endpoints/keyword/
type Keyword struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// Get returns a single Keyword identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Keywords, an error is returned.
func (ks *KeywordService) Get(id int, opts ...FuncOption) (*Keyword, error) {
	url, err := ks.client.singleURL(KeywordEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var kw []Keyword

	err = ks.client.get(url, &kw)
	if err != nil {
		return nil, err
	}

	return &kw[0], nil
}

// List returns a list of Keywords identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Keywords based solely on the provided
// options. Any ID that does not match a Keyword is ignored. If none of the IDs
// match a Keyword, an error is returned.
func (ks *KeywordService) List(ids []int, opts ...FuncOption) ([]*Keyword, error) {
	url, err := ks.client.multiURL(KeywordEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var kw []*Keyword

	err = ks.client.get(url, &kw)
	if err != nil {
		return nil, err
	}

	return kw, nil
}

// Search returns a list of Keywords found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Keywords are found using the provided query, an error is returned.
func (ks *KeywordService) Search(qry string, opts ...FuncOption) ([]*Keyword, error) {
	url, err := ks.client.searchURL(KeywordEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var kw []*Keyword

	err = ks.client.get(url, &kw)
	if err != nil {
		return nil, err
	}

	return kw, nil
}

// Count returns the number of Keywords available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which Keywords to count.
func (ks *KeywordService) Count(opts ...FuncOption) (int, error) {
	ct, err := ks.client.getEndpointCount(KeywordEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Keyword object.
func (ks *KeywordService) ListFields() ([]string, error) {
	fl, err := ks.client.getEndpointFieldList(KeywordEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
