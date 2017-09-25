package igdb

import (
	"strconv"
	"strings"
)

// Review type
type Review struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	URL            URL    `json:"url"`
	CreatedAt      int    `json:"created_at"` //unix epoch
	UpdatedAt      int    `json:"updated_at"` //unix epoch
	Game           int    `json:"game"`
	Category       int    `json:"category"` // Documentation is missing
	Likes          int    `json:"likes"`
	Views          int    `json:"views"`
	RatingCategory int    `json:"rating_category"` // Documenation is missing
	Platform       int    `json:"platform"`
	Video          string `json:"video"`
	Introduction   string `json:"introduction"`
	Content        string `json:"content"`
	Conclusion     string `json:"conclusion"`
	PositivePoints string `json:"positive_points"`
	NegativePoints string `json:"negative_points"`
}

// GetReview gets IGDB information for a review identified by their unique IGDB ID.
func (c *Client) GetReview(id int, opts ...OptionFunc) (*Review, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "reviews/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var r []Review

	err := c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return &r[0], nil
}

// GetReviews gets IGDB information for a list of reviews identified by a list of their unique IGDB IDs.
func (c *Client) GetReviews(ids []int, opts ...OptionFunc) ([]*Review, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToStrings(ids)
	url := c.rootURL + "reviews/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var r []*Review

	err := c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// SearchReviews searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchReviews(qry string, opts ...OptionFunc) ([]*Review, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "reviews/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var r []*Review

	err := c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
