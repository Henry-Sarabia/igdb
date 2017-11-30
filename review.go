package igdb

// ReviewService handles all the API
// calls for the IGDB Review endpoint.
type ReviewService service

// Review contains information on an IGDB entry
// for a review on a particular video game.
//
// For more information, visit: https://igdb.github.io/api/endpoints/review/
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

// Get returns a single Review identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Reviews, an error is returned.
func (rs *ReviewService) Get(id int, opts ...OptionFunc) (*Review, error) {
	url, err := rs.client.singleURL(ReviewEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var r []Review

	err = rs.client.get(url, &r)
	if err != nil {
		return nil, err
	}

	return &r[0], nil
}

// List returns a list of Reviews identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Reviews based solely on the provided
// options. Any ID that does not match a Review is ignored. If none of the IDs
// match a Review, an error is returned.
func (rs *ReviewService) List(ids []int, opts ...OptionFunc) ([]*Review, error) {
	url, err := rs.client.multiURL(ReviewEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var r []*Review

	err = rs.client.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Search returns a list of Reviews found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Reviews are found using the provided query, an error is returned.
func (rs *ReviewService) Search(qry string, opts ...OptionFunc) ([]*Review, error) {
	url, err := rs.client.searchURL(ReviewEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var r []*Review

	err = rs.client.get(url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Count returns the number of Reviews available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which Reviews to count.
func (rs *ReviewService) Count(opts ...OptionFunc) (int, error) {
	ct, err := rs.client.getEndpointCount(ReviewEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Review object.
func (rs *ReviewService) ListFields() ([]string, error) {
	fl, err := rs.client.getEndpointFieldList(ReviewEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
