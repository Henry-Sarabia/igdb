package igdb

import (
	"strconv"
	"strings"
)

// Person type
type Person struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	URL         URL         `json:"url"`
	CreatedAt   int         `json:"created_at"` //unix epoch
	UpdatedAt   int         `json:"updated_at"` //unix epoch
	DOB         int         `json:"dob"`
	Gender      GenderCode  `json:"gender"`
	Country     CountryCode `json:"country"`
	Mugshot     Image       `json:"mug_shot"`
	Bio         string      `json:"bio"`
	Description string      `json:"description"`
	Games       []int       `json:"games"`
	Characters  []int       `json:"characters"`
	VoiceActed  []int       `json:"voice_acted"`
}

// GetPerson gets IGDB information for a person identified by its unique IGDB ID.
func (c *Client) GetPerson(id int, opts ...OptionFunc) (*Person, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "people/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var p []Person

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPersons gets IGDB information for a list of people identified by their
// unique IGDB IDs.
func (c *Client) GetPersons(ids []int, opts ...OptionFunc) ([]*Person, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToString(ids)
	url := c.rootURL + "people/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var p []*Person

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPersons searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchPersons(qry string, opts ...OptionFunc) ([]*Person, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "people/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var p []*Person

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
