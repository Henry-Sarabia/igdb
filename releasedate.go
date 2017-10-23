package igdb

// ReleaseDate contains information on
// an IGDB entry for a particular release
// date. ReleaseDate is used primarily as
// an extension to the Game endpoint.
// ReleaseDate does not support the search
// functionality.
type ReleaseDate struct {
	ID          int          `json:"id"`
	Game        int          `json:"game"`
	ReleaseDate int          `json:"ReleaseDate"`
	Category    DateCategory `json:"category"`
	Platform    int          `json:"platform"`
	Human       string       `json:"human"`
	UpdatedAt   int          `json:"updated_at"` // Unix time in milliseconds unspecified
	CreatedAt   int          `json:"created_at"` // Unix time in milliseconds unspecified
	Date        int          `json:"date"`       // Unix time in milliseconds
	Region      int          `json:"region"`
	Year        int          `json:"y"`
	Month       int          `json:"m"`
}

// GetReleaseDate returns a single ReleaseDate identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have an
// effect due to GetReleaseDate only returning a single ReleaseDate object and not
// a list of ReleaseDates.
func (c *Client) GetReleaseDate(id int, opts ...OptionFunc) (*ReleaseDate, error) {
	url, err := c.singleURL(ReleaseDateEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var r []ReleaseDate

	err = c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return &r[0], nil
}

// GetReleaseDates returns a list of ReleaseDates identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
func (c *Client) GetReleaseDates(ids []int, opts ...OptionFunc) ([]*ReleaseDate, error) {
	url, err := c.multiURL(ReleaseDateEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var r []*ReleaseDate

	err = c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
