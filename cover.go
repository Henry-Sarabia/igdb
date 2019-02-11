package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// Cover represents the cover art for a specific video game.
// For more information visit: https://api-docs.igdb.com/#cover
type Cover struct {
	Image
	ID   int `json:"id"`
	Game int `json:"game"`
}

// CoverService handles all the API calls for the IGDB Cover endpoint.
type CoverService service

// Get returns a single Cover identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Covers, an error is returned.
func (cs *CoverService) Get(id int, opts ...FuncOption) (*Cover, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var cov []*Cover

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.get(cs.end, &cov, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Cover with ID %v", id)
	}

	return cov[0], nil
}

// List returns a list of Covers identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Cover is ignored. If none of the IDs
// match a Cover, an error is returned.
func (cs *CoverService) List(ids []int, opts ...FuncOption) ([]*Cover, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var cov []*Cover

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := cs.client.get(cs.end, &cov, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Covers with IDs %v", ids)
	}

	return cov, nil
}

// Index returns an index of Covers based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Covers can
// be found using the provided options, an error is returned.
func (cs *CoverService) Index(opts ...FuncOption) ([]*Cover, error) {
	var cov []*Cover

	err := cs.client.get(cs.end, &cov, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Covers")
	}

	return cov, nil
}

// Count returns the number of Covers available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Covers to count.
func (cs *CoverService) Count(opts ...FuncOption) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Covers")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Cover object.
func (cs *CoverService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Cover fields")
	}

	return f, nil
}
