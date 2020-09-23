package igdb

import "github.com/pkg/errors"

//go:generate gomodifytags -file $GOFILE -struct SearchResult -add-tags json -w

// SearchResult represents a result from searching the IGDB.
// It can contain: Characters, Collections Games, People, Platforms, and Themes.
type SearchResult struct {
	AlternativeName string `json:"alternative_name"`
	Character       int    `json:"character"`
	Collection      int    `json:"collection"`
	Company         int    `json:"company"`
	Description     string `json:"description"`
	Game            int    `json:"game"`
	Name            string `json:"name"`
	Person          int    `json:"person"`
	Platform        int    `json:"platform"`
	PublishedAt     int    `json:"published_at"`
	TestDummy       int    `json:"test_dummy"`
	Theme           int    `json:"theme"`
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
