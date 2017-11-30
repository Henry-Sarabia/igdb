package igdb

// EngineService handles all the API
// calls for the IGDB Engine endpoint.
type EngineService service

// Engine contains information on an IGDB
// entry for a particular video game engine.
type Engine struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Logo      Image  `json:"logo"`
	Games     []int  `json:"games"`
	Companies []int  `json:"companies"`
	Platforms []int  `json:"platforms"`
}

// Get returns a single Engine identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Engines, an error is returned.
func (es *EngineService) Get(id int, opts ...OptionFunc) (*Engine, error) {
	url, err := es.client.singleURL(EngineEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var eng []Engine

	err = es.client.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return &eng[0], nil
}

// List returns a list of Engines identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Engines based solely on the provided
// options. Any ID that does not match a Engine is ignored. If none of the IDs
// match a Engine, an error is returned.
func (es *EngineService) List(ids []int, opts ...OptionFunc) ([]*Engine, error) {
	url, err := es.client.multiURL(EngineEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var eng []*Engine

	err = es.client.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return eng, nil
}

// Search returns a list of Engines found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Engines are found using the provided query, an error is returned.
func (es *EngineService) Search(qry string, opts ...OptionFunc) ([]*Engine, error) {
	url, err := es.client.searchURL(EngineEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var eng []*Engine

	err = es.client.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return eng, nil
}

// Count returns the number of Engines available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which Engines to count.
func (es *EngineService) Count(opts ...OptionFunc) (int, error) {
	ct, err := es.client.getEndpointCount(EngineEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Engine object.
func (es *EngineService) ListFields() ([]string, error) {
	fl, err := es.client.getEndpointFieldList(EngineEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
