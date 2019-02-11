package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Feed -add-tags json -w

// Feed items are a social feed of status updates, media, and news articles.
// For more information visit: https://api-docs.igdb.com/#feed
type Feed struct {
	ID             int          `json:"id"`
	Category       FeedCategory `json:"category"`
	Content        string       `json:"content"`
	CreatedAt      int          `json:"created_at"`
	FeedLikesCount int          `json:"feed_likes_count"`
	FeedVideo      int          `json:"feed_video"`
	Games          []int        `json:"games"`
	Meta           string       `json:"meta"`
	PublishedAt    int          `json:"published_at"`
	Pulse          int          `json:"pulse"`
	Slug           string       `json:"slug"`
	Title          string       `json:"title"`
	UID            string       `json:"uid"`
	UpdatedAt      int          `json:"updated_at"`
	URL            string       `json:"url"`
	User           int          `json:"user"`
}

type FeedCategory int

//go:generate stringer -type=FeedCategory

const (
	FeedPulseArticle FeedCategory = iota + 1
	FeedComingSoon
	FeedNewTrailer
	_
	FeedUserContributedItem
	FeedUserContributionsItem
	FeedPageContributedItem
)

// FeedService handles all the API calls for the IGDB Feed endpoint.
type FeedService service

// Get returns a single Feed identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Feeds, an error is returned.
func (fs *FeedService) Get(id int, opts ...FuncOption) (*Feed, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var feed []*Feed

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := fs.client.get(fs.end, &feed, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Feed with ID %v", id)
	}

	return feed[0], nil
}

// List returns a list of Feeds identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Feed is ignored. If none of the IDs
// match a Feed, an error is returned.
func (fs *FeedService) List(ids []int, opts ...FuncOption) ([]*Feed, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var feed []*Feed

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := fs.client.get(fs.end, &feed, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Feeds with IDs %v", ids)
	}

	return feed, nil
}

// Index returns an index of Feeds based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Feeds can
// be found using the provided options, an error is returned.
func (fs *FeedService) Index(opts ...FuncOption) ([]*Feed, error) {
	var feed []*Feed

	err := fs.client.get(fs.end, &feed, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Feeds")
	}

	return feed, nil
}

// Count returns the number of Feeds available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Feeds to count.
func (fs *FeedService) Count(opts ...FuncOption) (int, error) {
	ct, err := fs.client.getCount(fs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Feeds")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Feed object.
func (fs *FeedService) Fields() ([]string, error) {
	f, err := fs.client.getFields(fs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Feed fields")
	}

	return f, nil
}
