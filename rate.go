package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Rate -add-tags json -w

// Rate represents a user's rating.
// For more information visit: https://api-docs.igdb.com/#rate
type Rate struct {
	ID     int     `json:"id"`
	Rating float64 `json:"rating"`
	User   int     `json:"user"`
}

// RateService handles all the API calls for the IGDB Rate endpoint.
type RateService service

// Get returns a single Rate identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Rates, an error is returned.
func (rs *RateService) Get(id int, opts ...Option) (*Rate, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var rate []*Rate

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := rs.client.get(rs.end, &rate, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Rate with ID %v", id)
	}

	return rate[0], nil
}

// List returns a list of Rates identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Rate is ignored. If none of the IDs
// match a Rate, an error is returned.
func (rs *RateService) List(ids []int, opts ...Option) ([]*Rate, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var rate []*Rate

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := rs.client.get(rs.end, &rate, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Rates with IDs %v", ids)
	}

	return rate, nil
}

// Index returns an index of Rates based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Rates can
// be found using the provided options, an error is returned.
func (rs *RateService) Index(opts ...Option) ([]*Rate, error) {
	var rate []*Rate

	err := rs.client.get(rs.end, &rate, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Rates")
	}

	return rate, nil
}

// Count returns the number of Rates available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Rates to count.
func (rs *RateService) Count(opts ...Option) (int, error) {
	ct, err := rs.client.getCount(rs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Rates")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Rate object.
func (rs *RateService) Fields() ([]string, error) {
	f, err := rs.client.getFields(rs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Rate fields")
	}

	return f, nil
}
