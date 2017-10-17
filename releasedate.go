package igdb

// ReleaseDate hold information about date of release, platforms, and versions
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

// GetReleaseDate gets IGDB information for a release date identified by their unique IGDB ID.
func (c *Client) GetReleaseDate(id int, opts ...optionFunc) (*ReleaseDate, error) {
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

// GetReleaseDates gets IGDB information for a list of release dates identified by a list of their unique IGDB IDs.
func (c *Client) GetReleaseDates(ids []int, opts ...optionFunc) ([]*ReleaseDate, error) {
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
