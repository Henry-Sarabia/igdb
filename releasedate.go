package igdb

import (
	"strconv"
	"strings"
)

// DateCategory code
type DateCategory int

// Year is
type Year int

// Month is
type Month int

// ReleaseDate hold information about date of release, platforms, and versions
type ReleaseDate struct {
	ID          int          `json:"id"`
	Game        int          `json:"game"`
	ReleaseDate int          `json:"ReleaseDate"`
	Category    DateCategory `json:"category"`
	Platform    int          `json:"platform"`
	Human       string       `json:"human"`
	UpdatedAt   int          `json:"updated_at"` //unix epoch unspecified
	CreatedAt   int          `json:"created_at"` //unix epoch unspecified
	Date        int          `json:"date"`       //unix epoch
	Region      int          `json:"region"`
	Year        Year         `json:"y"`
	Month       Month        `json:"m"`
}

// GetReleaseDate gets IGDB information for a release date identified by their unique IGDB ID.
func (c *Client) GetReleaseDate(id int, opts ...OptionFunc) (*ReleaseDate, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "release_dates/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var r []ReleaseDate

	err := c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return &r[0], nil
}

// GetReleaseDates gets IGDB information for a list of release dates identified by a list of their unique IGDB IDs.
func (c *Client) GetReleaseDates(ids []int, opts ...OptionFunc) ([]*ReleaseDate, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToString(ids)
	url := c.rootURL + "release_dates/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var r []*ReleaseDate

	err := c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
