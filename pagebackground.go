package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

// PageBackground represents the background image of a specific page.
// For more information visit: https://api-docs.igdb.com/#page-background
type PageBackground struct {
	Image
	ID int `json:"id,omitempty"`
}

// PageBackgroundService handles all the API calls for the IGDB PageBackground endpoint.
type PageBackgroundService service

// Get returns a single PageBackground identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PageBackgrounds, an error is returned.
func (ps *PageBackgroundService) Get(id int, opts ...Option) (*PageBackground, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var bg []*PageBackground

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &bg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PageBackground with ID %v", id)
	}

	return bg[0], nil
}

// List returns a list of PageBackgrounds identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PageBackground is ignored. If none of the IDs
// match a PageBackground, an error is returned.
func (ps *PageBackgroundService) List(ids []int, opts ...Option) ([]*PageBackground, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var bg []*PageBackground

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &bg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PageBackgrounds with IDs %v", ids)
	}

	return bg, nil
}

// Index returns an index of PageBackgrounds based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PageBackgrounds can
// be found using the provided options, an error is returned.
func (ps *PageBackgroundService) Index(opts ...Option) ([]*PageBackground, error) {
	var bg []*PageBackground

	err := ps.client.get(ps.end, &bg, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PageBackgrounds")
	}

	return bg, nil
}

// Count returns the number of PageBackgrounds available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PageBackgrounds to count.
func (ps *PageBackgroundService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PageBackgrounds")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PageBackground object.
func (ps *PageBackgroundService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PageBackground fields")
	}

	return f, nil
}
