package igdb

// PulseSourceService handles all the API
// calls for the IGDB PulseSource endpoint.
type PulseSourceService service

// PulseSource contains information on
// an IGDB entry for a specific news source.
type PulseSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Game int    `json:"game"`
	Page int    `json:"page"`
}

// Get returns a single PulseSource identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PulseSources, an error is returned.
func (pss *PulseSourceService) Get(id int, opts ...OptionFunc) (*PulseSource, error) {
	url, err := pss.client.singleURL(PulseSourceEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var ps []PulseSource

	err = pss.client.get(url, &ps)
	if err != nil {
		return nil, err
	}

	return &ps[0], nil
}

// List returns a list of PulseSources identified by the provided list of IGDB IDs.
// Provide functional options to filter, sort, and paginate the results. Omitting
// IDs will instead retrieve an index of PulseSources based solely on the provided
// options. Any ID that does not match a PulseSource is ignored. If none of the IDs
// match a PulseSource, an error is returned.
func (pss *PulseSourceService) List(ids []int, opts ...OptionFunc) ([]*PulseSource, error) {
	url, err := pss.client.multiURL(PulseSourceEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var ps []*PulseSource

	err = pss.client.get(url, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// Search returns a list of PulseSources found by searching the IGDB using the provided
// query. Provide functional options to filter, sort, and paginate the results. If
// no PulseSources are found using the provided query, an error is returned.
func (pss *PulseSourceService) Search(qry string, opts ...OptionFunc) ([]*PulseSource, error) {
	url, err := pss.client.searchURL(PulseSourceEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var ps []*PulseSource

	err = pss.client.get(url, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// Count returns the number of PulseSources available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which PulseSources to count.
func (pss *PulseSourceService) Count(opts ...OptionFunc) (int, error) {
	ct, err := pss.client.GetEndpointCount(PulseSourceEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB PulseSource object.
func (pss *PulseSourceService) ListFields() ([]string, error) {
	fl, err := pss.client.GetEndpointFieldList(PulseSourceEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
