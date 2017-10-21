package igdb

// GameMode contains information on an
// IGDB entry for a particular game mode
// (e.g. single player, multiplayer).
type GameMode struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetGameMode returns a single GameMode identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetGameMode only returning a single GameMode object and
// not a list of GameModes.
func (c *Client) GetGameMode(id int, opts ...OptionFunc) (*GameMode, error) {
	url, err := c.singleURL(GameModeEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var g []GameMode

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// GetGameModes returns a list of GameModes identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of GameModes based
// solely on the provided options.
func (c *Client) GetGameModes(ids []int, opts ...OptionFunc) ([]*GameMode, error) {
	url, err := c.multiURL(GameModeEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var g []*GameMode

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// SearchGameModes returns a list of GameModes found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchGameModes(qry string, opts ...OptionFunc) ([]*GameMode, error) {
	url, err := c.searchURL(GameModeEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var g []*GameMode

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
