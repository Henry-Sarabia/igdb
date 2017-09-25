package igdb

import (
	"strconv"
	"strings"
)

// Keyword type
type Keyword struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Games     []int  `json:"games"`
}

// GetKeyword gets IGDB information for a keyword identified by its unique IGDB ID.
func (c *Client) GetKeyword(id int, opts ...OptionFunc) (*Keyword, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "keywords/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var k []Keyword

	err := c.get(url, &k)
	if err != nil {
		return nil, err
	}

	return &k[0], nil
}

// GetKeywords gets IGDB information for a list of keywords identified by their
// unique IGDB IDs.
func (c *Client) GetKeywords(ids []int, opts ...OptionFunc) ([]*Keyword, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToStrings(ids)
	url := c.rootURL + "keywords/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var k []*Keyword

	err := c.get(url, &k)
	if err != nil {
		return nil, err
	}

	return k, nil
}

// SearchKeywords searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchKeywords(qry string, opts ...OptionFunc) ([]*Keyword, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "keywords/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var k []*Keyword

	err := c.get(url, &k)
	if err != nil {
		return nil, err
	}

	return k, nil
}
