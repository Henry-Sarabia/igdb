package igdb

import (
	"strconv"
	"strings"
)

// GameMode is
type GameMode struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Games     []int  `json:"games"`
}

// GetGameMode gets IGDB information for a game mode identified by its unique IGDB ID.
func (c *Client) GetGameMode(id int, opts ...OptionFunc) (*GameMode, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "game_modes/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var g []GameMode

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// GetGameModes gets IGDB information for a list of game modes identified by their
// unique IGDB IDs.
func (c *Client) GetGameModes(ids []int, opts ...OptionFunc) ([]*GameMode, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToStrings(ids)
	url := c.rootURL + "game_modes/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var g []*GameMode

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// SearchGameModes searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchGameModes(qry string, opts ...OptionFunc) ([]*GameMode, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "game_modes/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var g []*GameMode

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
