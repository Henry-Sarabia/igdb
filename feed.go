package igdb

import (
	"strconv"
	"strings"
)

// FeedCategory codes
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
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "feeds/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var f []Feed

	err := c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return &f[0], nil
}

// GetFeeds gets IGDB information for a list of game engines identified by their
// unique IGDB IDs.
func (c *Client) GetFeeds(ids []int, opts ...OptionFunc) ([]*Feed, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToString(ids)
	url := c.rootURL + "feeds/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var f []*Feed

	err := c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
