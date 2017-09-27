package igdb

// GameMode is
type GameMode struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Games     []int  `json:"games"`
}

// GetGameMode gets IGDB information for a game mode identified by its unique IGDB ID.
func (c *Client) GetGameMode(id int, opts ...OptionFunc) (*GameMode, error) {
	url := c.singleURL(GameModeEndpoint, id, opts...)

	var g []GameMode

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// GetGameModes gets IGDB information for a list of game modes identified by their
// unique IGDB IDs.
func (c *Client) GetGameModes(ids []int, opts ...OptionFunc) ([]*GameMode, error) {
	url := c.multiURL(GameModeEndpoint, ids, opts...)

	var g []*GameMode

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// SearchGameModes searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchGameModes(qry string, opts ...OptionFunc) ([]*GameMode, error) {
	url := c.searchURL(GameModeEndpoint, qry, opts...)

	var g []*GameMode

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
