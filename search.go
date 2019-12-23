package igdb

import "github.com/pkg/errors"

//go:generate gomodifytags -file $GOFILE -struct SearchResult -add-tags json -w

// SearchResult represents a result from searching the IGDB.
// It can contain: Characters, Collections Games, People, Platforms, and Themes.
type SearchResult struct {
	AlternativeName string  `json:"alternative_name,omitempty"`
	Character       int     `json:"character,omitempty"`
	Collection      int     `json:"collection,omitempty"`
	Company         int     `json:"company,omitempty"`
	Description     string  `json:"description,omitempty"`
	Game            int     `json:"game,omitempty"`
	Name            string  `json:"name,omitempty"`
	Person          int     `json:"person,omitempty"`
	Platform        int     `json:"platform,omitempty"`
	Popularity      float64 `json:"popularity,omitempty"`
	PublishedAt     int     `json:"published_at,omitempty"`
	TestDummy       int     `json:"test_dummy,omitempty"`
	Theme           int     `json:"theme,omitempty"`
}

// Search returns a list of SearchResults using the provided query. Provide functional
// options to sort, filter, and paginate the results. If no results are found, an error
// is returned.
// Search can only search through Characters, Collections, Games, People, Platforms,
// and Themes.
func (c *Client) Search(qry string, opts ...Option) ([]*SearchResult, error) {
	var res []*SearchResult

	opts = append(opts, setSearch(qry))
	err := c.get(EndpointSearch, &res, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot perform search with query %s", qry)
	}

	return res, nil
}
