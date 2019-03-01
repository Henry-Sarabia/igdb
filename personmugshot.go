package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

// PersonMugshot represents the mugshot of
// a person in the video game industry.
// For more information visit: https://api-docs.igdb.com/#person-mug-shot
type PersonMugshot struct {
	Image
	ID int `json:"id"`
}

// PersonMugshotService handles all the API calls for the IGDB PersonMugshot endpoint.
type PersonMugshotService service

// Get returns a single PersonMugshot identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PersonMugshots, an error is returned.
func (ps *PersonMugshotService) Get(id int, opts ...Option) (*PersonMugshot, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var mug []*PersonMugshot

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &mug, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PersonMugshot with ID %v", id)
	}

	return mug[0], nil
}

// List returns a list of PersonMugshots identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PersonMugshot is ignored. If none of the IDs
// match a PersonMugshot, an error is returned.
func (ps *PersonMugshotService) List(ids []int, opts ...Option) ([]*PersonMugshot, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var mug []*PersonMugshot

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &mug, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PersonMugshots with IDs %v", ids)
	}

	return mug, nil
}

// Index returns an index of PersonMugshots based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PersonMugshots can
// be found using the provided options, an error is returned.
func (ps *PersonMugshotService) Index(opts ...Option) ([]*PersonMugshot, error) {
	var mug []*PersonMugshot

	err := ps.client.get(ps.end, &mug, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PersonMugshots")
	}

	return mug, nil
}

// Count returns the number of PersonMugshots available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PersonMugshots to count.
func (ps *PersonMugshotService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PersonMugshots")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PersonMugshot object.
func (ps *PersonMugshotService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PersonMugshot fields")
	}

	return f, nil
}
