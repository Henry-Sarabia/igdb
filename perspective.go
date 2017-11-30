package igdb

// PerspectiveService handles all the API
// calls for the IGDB Perspective endpoint.
type PerspectiveService service

// Perspective contains information on an IGDB entry for a particular
// player perspective (e.g. first-person or virtual reality).
type Perspective struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// Get returns a single Perspective identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Perspectives, an error is returned.
func (ps *PerspectiveService) Get(id int, opts ...OptionFunc) (*Perspective, error) {
	url, err := ps.client.singleURL(PerspectiveEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var p []Perspective

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// List returns a list of Perspectives identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Perspectives based solely on the provided
// options. Any ID that does not match a Perspective is ignored. If none of the IDs
// match a Perspective, an error is returned.
func (ps *PerspectiveService) List(ids []int, opts ...OptionFunc) ([]*Perspective, error) {
	url, err := ps.client.multiURL(PerspectiveEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var p []*Perspective

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Search returns a list of Perspectives found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Perspectives are found using the provided query, an error is returned.
func (ps *PerspectiveService) Search(qry string, opts ...OptionFunc) ([]*Perspective, error) {
	url, err := ps.client.searchURL(PerspectiveEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var p []*Perspective

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Count returns the number of Perspectives available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which Perspectives to count.
func (ps *PerspectiveService) Count(opts ...OptionFunc) (int, error) {
	ct, err := ps.client.GetEndpointCount(PerspectiveEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Perspective object.
func (ps *PerspectiveService) ListFields() ([]string, error) {
	fl, err := ps.client.GetEndpointFieldList(PerspectiveEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
