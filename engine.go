package igdb

import (
	"strconv"
	"strings"
)

// Engine is
type Engine struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Logo      Image  `json:"logo"`
	Games     []int  `json:"games"`
	Companies []int  `json:"companies"`
	Platforms []int  `json:"platforms"`
}

// GetEngine gets IGDB information for a game engine identified by its unique IGDB ID.
func (c *Client) GetEngine(id int, opts ...OptionFunc) (*Engine, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "game_engines/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var eng []Engine

	err := c.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return &eng[0], nil
}

// GetEngines gets IGDB information for a list of game engines identified by their
// unique IGDB IDs.
func (c *Client) GetEngines(ids []int, opts ...OptionFunc) ([]*Engine, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToStrings(ids)
	url := c.rootURL + "game_engines/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var eng []*Engine

	err := c.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return eng, nil
}

// SearchEngines searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchEngines(qry string, opts ...OptionFunc) ([]*Engine, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "game_engines/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var eng []*Engine

	err := c.get(url, &eng)
	if err != nil {
		return nil, err
	}

	return eng, nil
}
