package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Review -add-tags json -w

type Review struct {
	ID             int            `json:"id"`
	Category       ReviewCategory `json:"category"`
	Conclusion     string         `json:"conclusion"`
	Content        string         `json:"content"`
	CreatedAt      int            `json:"created_at"`
	Game           int            `json:"game"`
	Introduction   string         `json:"introduction"`
	Likes          int            `json:"likes"`
	NegativePoints string         `json:"negative_points"`
	Platform       int            `json:"platform"`
	PositivePoints string         `json:"positive_points"`
	Slug           string         `json:"slug"`
	Title          string         `json:"title"`
	UpdatedAt      int            `json:"updated_at"`
	URL            string         `json:"url"`
	User           int            `json:"user"`
	UserRating     int            `json:"user_rating"`
	Video          int            `json:"video"`
	Views          int            `json:"views"`
}

//go:generate stringer -type=ReviewCategory

type ReviewCategory int

const (
	ReviewText ReviewCategory = iota + 1
	ReviewVid
)

// ReviewService handles all the API calls for the IGDB Review endpoint.
type ReviewService service

// Get returns a single Review identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Reviews, an error is returned.
func (rs *ReviewService) Get(id int, opts ...Option) (*Review, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var rev []*Review

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := rs.client.get(rs.end, &rev, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Review with ID %v", id)
	}

	return rev[0], nil
}

// List returns a list of Reviews identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Review is ignored. If none of the IDs
// match a Review, an error is returned.
func (rs *ReviewService) List(ids []int, opts ...Option) ([]*Review, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var rev []*Review

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := rs.client.get(rs.end, &rev, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Reviews with IDs %v", ids)
	}

	return rev, nil
}

// Index returns an index of Reviews based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Reviews can
// be found using the provided options, an error is returned.
func (rs *ReviewService) Index(opts ...Option) ([]*Review, error) {
	var rev []*Review

	err := rs.client.get(rs.end, &rev, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Reviews")
	}

	return rev, nil
}

// Count returns the number of Reviews available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Reviews to count.
func (rs *ReviewService) Count(opts ...Option) (int, error) {
	ct, err := rs.client.getCount(rs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Reviews")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Review object.
func (rs *ReviewService) Fields() ([]string, error) {
	f, err := rs.client.getFields(rs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Review fields")
	}

	return f, nil
}
