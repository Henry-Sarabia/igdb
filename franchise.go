package igdb

import (
	"strconv"
	"strings"
)

// Franchise is
type Franchise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Games     []int  `json:"games"`
}

// GetFranchise gets IGDB information for a franchise identified by its unique IGDB ID.
func (c *Client) GetFranchise(id int, opts ...OptionFunc) (*Franchise, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "franchises/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var f []Franchise

	err := c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return &f[0], nil
}

// GetFranchises gets IGDB information for a list of franchises identified by their
// unique IGDB IDs.
func (c *Client) GetFranchises(ids []int, opts ...OptionFunc) ([]*Franchise, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToStrings(ids)
	url := c.rootURL + "franchises/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var f []*Franchise

	err := c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// SearchFranchises searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchFranchises(qry string, opts ...OptionFunc) ([]*Franchise, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "franchises/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var f []*Franchise

	err := c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
