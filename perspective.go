package igdb

// Perspective contains information on
// an IGDB entry for a particular player
// perspective (e.g. first-person or
// virtual reality).
type Perspective struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetPerspective returns a single Perspective identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have an
// effect due to GetPerspective only returning a single Perspective object and not
// a list of Perspectives.
func (c *Client) GetPerspective(id int, opts ...OptionFunc) (*Perspective, error) {
	url, err := c.singleURL(PerspectiveEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []Perspective

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPerspectives returns a list of Perspectives identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
func (c *Client) GetPerspectives(ids []int, opts ...OptionFunc) ([]*Perspective, error) {
	url, err := c.multiURL(PerspectiveEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Perspective

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPerspectives returns a list of Perspectives found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
// Providing an empty query will instead retrieve an index of Perspectives based solely on
// the provided options.
func (c *Client) SearchPerspectives(qry string, opts ...OptionFunc) ([]*Perspective, error) {
	url, err := c.searchURL(PerspectiveEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Perspective

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
