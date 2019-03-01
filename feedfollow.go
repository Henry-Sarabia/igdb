package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct FeedFollow -add-tags json -w

// FeedFollow represents the following of social feed composed of
// status updates, media, and news articles.
// For more information visit: https://api-docs.igdb.com/#feed-follow
type FeedFollow struct {
	ID          int          `json:"id"`
	CreatedAt   int          `json:"created_at"`
	Feed        FeedCategory `json:"feed"`
	PublishedAt int          `json:"published_at"`
	UpdatedAt   int          `json:"updated_at"`
	User        int          `json:"user"`
}

// FeedFollowService handles all the API calls for the IGDB FeedFollow endpoint.
type FeedFollowService service

// Get returns a single FeedFollow identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any FeedFollows, an error is returned.
func (fs *FeedFollowService) Get(id int, opts ...Option) (*FeedFollow, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ff []*FeedFollow

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := fs.client.get(fs.end, &ff, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get FeedFollow with ID %v", id)
	}

	return ff[0], nil
}

// List returns a list of FeedFollows identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a FeedFollow is ignored. If none of the IDs
// match a FeedFollow, an error is returned.
func (fs *FeedFollowService) List(ids []int, opts ...Option) ([]*FeedFollow, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ff []*FeedFollow

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := fs.client.get(fs.end, &ff, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get FeedFollows with IDs %v", ids)
	}

	return ff, nil
}

// Index returns an index of FeedFollows based solely on the provided functional
// options used to sort, filter, and paginate the results. If no FeedFollows can
// be found using the provided options, an error is returned.
func (fs *FeedFollowService) Index(opts ...Option) ([]*FeedFollow, error) {
	var ff []*FeedFollow

	err := fs.client.get(fs.end, &ff, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of FeedFollows")
	}

	return ff, nil
}

// Count returns the number of FeedFollows available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which FeedFollows to count.
func (fs *FeedFollowService) Count(opts ...Option) (int, error) {
	ct, err := fs.client.getCount(fs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count FeedFollows")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB FeedFollow object.
func (fs *FeedFollowService) Fields() ([]string, error) {
	f, err := fs.client.getFields(fs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get FeedFollow fields")
	}

	return f, nil
}
