package igdb

// GameModeService handles all the API
// calls for the IGDB GameMode endpoint.
type GameModeService service

// GameMode contains information on an IGDB entry for a particular game mode
// (e.g. single player, multiplayer).
//
// For more information, visit: https://igdb.github.io/api/endpoints/game-mode/
type GameMode struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// Get returns a single GameMode identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameModes, an error is returned.
func (gms *GameModeService) Get(id int, opts ...OptionFunc) (*GameMode, error) {
	url, err := gms.client.singleURL(GameModeEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var gm []GameMode

	err = gms.client.get(url, &gm)
	if err != nil {
		return nil, err
	}

	return &gm[0], nil
}

// List returns a list of GameModes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of GameModes based solely on the provided
// options. Any ID that does not match a GameMode is ignored. If none of the IDs
// match a GameMode, an error is returned.
func (gms *GameModeService) List(ids []int, opts ...OptionFunc) ([]*GameMode, error) {
	url, err := gms.client.multiURL(GameModeEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var gm []*GameMode

	err = gms.client.get(url, &gm)
	if err != nil {
		return nil, err
	}

	return gm, nil
}

// Search returns a list of GameModes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no GameModes are found using the provided query, an error is returned.
func (gms *GameModeService) Search(qry string, opts ...OptionFunc) ([]*GameMode, error) {
	url, err := gms.client.searchURL(GameModeEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var gm []*GameMode

	err = gms.client.get(url, &gm)
	if err != nil {
		return nil, err
	}

	return gm, nil
}

// Count returns the number of GameModes available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which GameModes to count.
func (gms *GameModeService) Count(opts ...OptionFunc) (int, error) {
	ct, err := gms.client.getEndpointCount(GameModeEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB GameMode object.
func (gms *GameModeService) ListFields() ([]string, error) {
	fl, err := gms.client.getEndpointFieldList(GameModeEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
