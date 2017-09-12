package igdb

import (
	"strconv"
	"strings"
)

// Genre type
type Genre struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Games     []int  `json:"games"`
}

// GetGenre gets IGDB information for a genre identified by its unique IGDB ID.
func (c *Client) GetGenre(id int, opts ...OptionFunc) (*Genre, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "genres/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var g []Genre

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// GetGenres gets IGDB information for a list of genres identified by their
// unique IGDB IDs.
func (c *Client) GetGenres(ids []int, opts ...OptionFunc) ([]*Genre, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToString(ids)
	url := c.rootURL + "genres/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var g []*Genre

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// SearchGenres searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchGenres(qry string, opts ...OptionFunc) ([]*Genre, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "genres/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var g []*Genre

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
