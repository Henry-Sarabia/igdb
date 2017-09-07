package igdb

import (
	"strconv"
	"strings"
)

// GenderCode codes
type GenderCode int

// SpeciesCode codes
type SpeciesCode int

// Character is
type Character struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
	URL       URL         `json:"url"`
	CreatedAt int         `json:"created_at"`
	UpdatedAt int         `json:"updated_at"`
	Mugshot   Image       `json:"mug_shot"`
	Gender    GenderCode  `json:"gender"`
	AKAs      []string    `json:"akas"`
	Species   SpeciesCode `json:"species"`
	Games     []int       `json:"games"`
	People    []int       `json:"people"`
}

// GetCharacter gets IGDB information for a character identified by its unique IGDB ID.
func (c *Client) GetCharacter(id int, opts ...OptionFunc) (*Character, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "characters/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var ch []Character

	err := c.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return &ch[0], nil
}

// GetCharacters gets IGDB information for a list of characters identified by their
// unique IGDB IDs.
func (c *Client) GetCharacters(ids []int, opts ...OptionFunc) ([]*Character, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToString(ids)
	url := c.rootURL + "characters/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var ch []*Character

	err := c.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

// SearchCharacters searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchCharacters(qry string, opts ...OptionFunc) ([]*Character, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "characters/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var ch []*Character

	err := c.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
