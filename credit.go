package igdb

import (
	"strconv"
	"strings"
)

// CreditCategory codes
type CreditCategory int

// Credit is
type Credit struct {
	ID                    int            `json:"id"`
	Name                  string         `json:"name"`
	Slug                  string         `json:"slug"`
	URL                   URL            `json:"url"`
	CreatedAt             int            `json:"created_at"` //unix epoch
	UpdatedAt             int            `json:"updated_at"` //unix epoch
	Game                  int            `json:"game"`
	Category              CreditCategory `json:"category"`
	Company               int            `json:"company"`
	Position              int            `json:"position"`
	Person                int            `json:"person"`
	Character             int            `json:"character"`
	Title                 int            `json:"title"`
	Country               CountryCode    `json:"country"`
	CreditedName          string         `json:"credited_name"`
	CharacterCreditedName string         `json:"character_credited_name"`
	PersonTitle           interface{}    `json:"person_title"`
}

// GetCredit gets IGDB information for a credit identified by its unique IGDB ID.
func (c *Client) GetCredit(id int, opts ...OptionFunc) (*Credit, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "credits/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var cr []Credit

	err := c.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return &cr[0], nil
}

// GetCredits gets IGDB information for a list of credits identified by their
// unique IGDB IDs.
func (c *Client) GetCredits(ids []int, opts ...OptionFunc) ([]*Credit, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToStrings(ids)
	url := c.rootURL + "credits/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var cr []*Credit

	err := c.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}

// SearchCredits searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchCredits(qry string, opts ...OptionFunc) ([]*Credit, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "credits/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var cr []*Credit

	err := c.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}
