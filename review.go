package igdb

// Review contains information on an
// IGDB entry for a review on a particular
// video game.
type Review struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	URL            URL    `json:"url"`
	CreatedAt      int    `json:"created_at"` // Unix time in milliseconds
	UpdatedAt      int    `json:"updated_at"` // Unix time in milliseconds
	Game           int    `json:"game"`
	Category       int    `json:"category"` // Missing official documentation
	Likes          int    `json:"likes"`
	Views          int    `json:"views"`
	RatingCategory int    `json:"rating_category"` // Missing official documentation
	Platform       int    `json:"platform"`
	Video          string `json:"video"`
	Introduction   string `json:"introduction"`
	Content        string `json:"content"`
	Conclusion     string `json:"conclusion"`
	PositivePoints string `json:"positive_points"`
	NegativePoints string `json:"negative_points"`
}

// GetReview gets IGDB information for a Review identified by its unique
// IGDB ID. GetReview returns a single Review identified by the provided
// IGDB ID. Functional options may be provided but sorting and pagination
// will not have an effect due to GetReview only returning a single Review
// object and not a list of Reviews.
func (c *Client) GetReview(id int, opts ...OptionFunc) (*Review, error) {
	url, err := c.singleURL(ReviewEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var r []Review

	err = c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return &r[0], nil
}

// GetReviews returns a list of Reviews identified by the provided list of IGDB
// IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of Reviews
// based solely on the provided options.
func (c *Client) GetReviews(ids []int, opts ...OptionFunc) ([]*Review, error) {
	url, err := c.multiURL(ReviewEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var r []*Review

	err = c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// SearchReviews returns a list of Reviews found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchReviews(qry string, opts ...OptionFunc) ([]*Review, error) {
	url, err := c.searchURL(ReviewEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var r []*Review

	err = c.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
