package igdb

// Engine contains information on an
// IGDB entry for a particular video
// game engine.
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

// GetEngine returns a single Engine identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetEngine only returning a single Engine object and
// not a list of Engines.
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

// GetEngines returns a list of Engines identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
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

// SearchEngines returns a list of Engines found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
// Providing an empty query will instead retrieve an index of Engines based solely
// on the provided options.
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
