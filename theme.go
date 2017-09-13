package igdb

import (
	"strconv"
	"strings"
)

// Theme type
type Theme struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Games     []int  `json:"games"`
}

// GetTheme gets IGDB information for a theme identified by their unique IGDB ID.
func (c *Client) GetTheme(id int, opts ...OptionFunc) (*Theme, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "themes/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var t []Theme

	err := c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return &t[0], nil
}

// GetThemes gets IGDB information for a list of themes identified by a list of their unique IGDB IDs.
func (c *Client) GetThemes(ids []int, opts ...OptionFunc) ([]*Theme, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToString(ids)
	url := c.rootURL + "themes/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var t []*Theme

	err := c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchThemes searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchThemes(qry string, opts ...OptionFunc) ([]*Theme, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "themes/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var t []*Theme

	err := c.get(url, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
