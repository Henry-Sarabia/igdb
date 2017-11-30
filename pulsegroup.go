package igdb

// PulseGroupService handles all the API
// calls for the IGDB PulseGroup endpoint.
type PulseGroupService service

// PulseGroup contains information on an
// IGDB entry for a group of news articles.
type PulseGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         URL    `json:"url"`
	CreatedAt   int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt   int    `json:"updated_at"` // Unix time in milliseconds
	PublishedAt int    `json:"published_at"`
	Category    int    `json:"category"`
	Tags        []Tag  `json:"tags"`
	Pulses      []int  `json:"pulses"`
	Game        int    `json:"game"`
}

// Get returns a single PulseGroup identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PulseGroups, an error is returned.
func (pgs *PulseGroupService) Get(id int, opts ...OptionFunc) (*PulseGroup, error) {
	url, err := pgs.client.singleURL(PulseGroupEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var pg []PulseGroup

	err = pgs.client.get(url, &pg)
	if err != nil {
		return nil, err
	}

	return &pg[0], nil
}

// List returns a list of PulseGroups identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of PulseGroups based solely on the provided
// options. Any ID that does not match a PulseGroup is ignored. If none of the IDs
// match a PulseGroup, an error is returned.
func (pgs *PulseGroupService) List(ids []int, opts ...OptionFunc) ([]*PulseGroup, error) {
	url, err := pgs.client.multiURL(PulseGroupEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var pg []*PulseGroup

	err = pgs.client.get(url, &pg)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

// Search returns a list of PulseGroups found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no PulseGroups are found using the provided query, an error is returned.
func (pgs *PulseGroupService) Search(qry string, opts ...OptionFunc) ([]*PulseGroup, error) {
	url, err := pgs.client.searchURL(PulseGroupEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var pg []*PulseGroup

	err = pgs.client.get(url, &pg)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

// Count returns the number of PulseGroups available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which PulseGroups to count.
func (pgs *PulseGroupService) Count(opts ...OptionFunc) (int, error) {
	ct, err := pgs.client.getEndpointCount(PulseGroupEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB PulseGroup object.
func (pgs *PulseGroupService) ListFields() ([]string, error) {
	fl, err := pgs.client.getEndpointFieldList(PulseGroupEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
