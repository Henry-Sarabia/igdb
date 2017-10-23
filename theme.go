package igdb

// Theme contains information on an IGDB
// entry for a particular video game theme
// (e.g. Fantasy or Horror).
type Theme struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int    `json:"updated_at"` // Unix time in milliseconds
	Games     []int  `json:"games"`
}

// GetTheme returns a single Theme identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will
// not have an effect due to GetTheme only returning a single Theme
// object and not a list of Themes.
func (c *Client) GetTheme(id int, opts ...OptionFunc) (*Theme, error) {
	url, err := c.singleURL(ThemeEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var t []Theme

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return &t[0], nil
}

// GetThemes returns a list of Themes identified by the provided list of IGDB
// IDs. Provide functional options to filter, sort, and paginate the results.
func (c *Client) GetThemes(ids []int, opts ...OptionFunc) ([]*Theme, error) {
	url, err := c.multiURL(ThemeEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var t []*Theme

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchThemes returns a list of Themes found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the
// results. Providing an empty query will instead retrieve an index of Themes
// based solely on the provided options.
func (c *Client) SearchThemes(qry string, opts ...OptionFunc) ([]*Theme, error) {
	url, err := c.searchURL(ThemeEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var t []*Theme

	err = c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
