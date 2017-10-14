package igdb

// Engine is
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

// GetEngine gets IGDB information for a game engine identified by its unique IGDB ID.
func (c *Client) GetEngine(id int, opts ...OptionFunc) (*Engine, error) {
	url, err := c.singleURL(EngineEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var eng []Engine

	err = c.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return &eng[0], nil
}

// GetEngines gets IGDB information for a list of game engines identified by their
// unique IGDB IDs.
func (c *Client) GetEngines(ids []int, opts ...OptionFunc) ([]*Engine, error) {
	url, err := c.multiURL(EngineEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var eng []*Engine

	err = c.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return eng, nil
}

// SearchEngines searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchEngines(qry string, opts ...OptionFunc) ([]*Engine, error) {
	url, err := c.searchURL(EngineEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var eng []*Engine

	err = c.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return eng, nil
}
