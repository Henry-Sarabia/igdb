package igdb

// Feed contains information on an IGDB
// entry for a social feed composed of
// status updates, media, and news articles.
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

// GetFeed returns a single Feed identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will
// not have an effect due to GetFeed only returning a single Feed object
// and not a list of Feeds.
func (c *Client) GetFeed(id int, opts ...OptionFunc) (*Feed, error) {
	url, err := c.singleURL(FeedEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var f []Feed

	err = c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return &f[0], nil
}

// GetFeeds returns a list of Feeds identified by the provided list of IGDB
// IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of Feeds based
// solely on the provided options.
func (c *Client) GetFeeds(ids []int, opts ...OptionFunc) ([]*Feed, error) {
	url, err := c.multiURL(FeedEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var f []*Feed

	err = c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
