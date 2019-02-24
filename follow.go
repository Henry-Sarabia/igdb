package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Follow -add-tags json -w

type Follow struct {
	ID   int `json:"id"`
	Game int `json:"game"`
	User int `json:"user"`
}

// FollowService handles all the API calls for the IGDB Follow endpoint.
type FollowService service

// Get returns a single Follow identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Follows, an error is returned.
func (fs *FollowService) Get(id int, opts ...Option) (*Follow, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var f []*Follow

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := fs.client.get(fs.end, &f, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Follow with ID %v", id)
	}

	return f[0], nil
}

// List returns a list of Follows identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Follow is ignored. If none of the IDs
// match a Follow, an error is returned.
func (fs *FollowService) List(ids []int, opts ...Option) ([]*Follow, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var f []*Follow

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := fs.client.get(fs.end, &f, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Follows with IDs %v", ids)
	}

	return f, nil
}

// Index returns an index of Follows based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Follows can
// be found using the provided options, an error is returned.
func (fs *FollowService) Index(opts ...Option) ([]*Follow, error) {
	var f []*Follow

	err := fs.client.get(fs.end, &f, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Follows")
	}

	return f, nil
}

// Count returns the number of Follows available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Follows to count.
func (fs *FollowService) Count(opts ...Option) (int, error) {
	ct, err := fs.client.getCount(fs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Follows")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Follow object.
func (fs *FollowService) Fields() ([]string, error) {
	f, err := fs.client.getFields(fs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Follow fields")
	}

	return f, nil
}
