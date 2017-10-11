package igdb

// FeedCategory corresponds to the IGDB
// enumerated feed item category which
// describes the type of feed item in
// a particular feed. FeedCategory
// implements the Stringer interface.
type FeedCategory int

// Feed is
type Feed struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	URL         URL          `json:"url"`
	Slug        string       `json:"slug"`
	CreatedAt   int          `json:"created_at"` //unix epoch
	UpdatedAt   int          `json:"updated_at"` //unix epoch
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

// GetFeed gets IGDB information for a feed identified by its unique IGDB ID.
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

// GetFeeds gets IGDB information for a list of game engines identified by their
// unique IGDB IDs.
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

// FeedCategory's string method matches the code with
//  IGDB's Feed Item Category enumerated type
// and return the category as a string. Codes with
// no match will returned "Undefined".

// FeedCategory implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type Feed Item Category and returns the category
// as a string. Codes with no match will return
// "Undefined".
func (f FeedCategory) String() string {
	switch f {
	case 1:
		return "Pulse Article"
	case 2:
		return "Coming Soon"
	case 3:
		return "New Trailer"
	case 5:
		return "User Contributed Item"
	case 6:
		return "User Contributions Item"
	case 7:
		return "Page Contributed Item"
	default:
		return "Undefined"
	}
}
