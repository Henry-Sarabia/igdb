package igdb

// FeedService handles all the API
// calls for the IGDB Feed endpoint.
type FeedService service

// Feed contains information on an IGDB entry for a social feed composed of
// status updates, media, and news articles. Feed does not support the search
// function.
//
// For more information, visit: https://igdb.github.io/api/endpoints/feed/
type Feed struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	URL         URL          `json:"url"`
	Slug        string       `json:"slug"`
	CreatedAt   int          `json:"created_at"` // Unix time in milliseconds
	UpdatedAt   int          `json:"updated_at"` // Unix time in milliseconds
	PublishedAt int          `json:"published_at"`
	Content     string       `json:"content"`
	Category    FeedCategory `json:"category"`
	User        int          `json:"user"`
	Games       []int        `json:"games"`
	Title       string       `json:"title"`
	LikeCount   int          `json:"feed_likes_count"`
	FeedVideo   interface{}  `json:"feed_video"`
	Meta        string       `json:"meta"`
	Pulse       int          `json:"pulse"`
	UID         string       `json:"uid"`
}

// Get returns a single Feed identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Feeds, an error is returned.
func (fs *FeedService) Get(id int, opts ...FuncOption) (*Feed, error) {
	url, err := fs.client.singleURL(FeedEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var f []Feed

	err = fs.client.get(url, &f)
	if err != nil {
		return nil, err
	}

	return &f[0], nil
}

// List returns a list of Feeds identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of Feeds based solely on the provided
// options. Any ID that does not match a Feed is ignored. If none of the IDs
// match a Feed, an error is returned.
func (fs *FeedService) List(ids []int, opts ...FuncOption) ([]*Feed, error) {
	url, err := fs.client.multiURL(FeedEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var f []*Feed

	err = fs.client.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Count returns the number of Feeds available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Feeds to count.
func (fs *FeedService) Count(opts ...FuncOption) (int, error) {
	ct, err := fs.client.getEndpointCount(FeedEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Feed object.
func (fs *FeedService) ListFields() ([]string, error) {
	fl, err := fs.client.getEndpointFieldList(FeedEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
