package igdb

//go:generate gomodifytags -file $GOFILE -struct Feed -add-tags json -w

// Feed items are a social feed of status updates, media, and news articles.
// For more information visit: https://api-docs.igdb.com/#feed
type Feed struct {
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

//go:generate stringer -type=FeedCategory

type FeedCategory int

const (
	FeedPulseArticle FeedCategory = iota + 1
	FeedComingSoon
	FeedNewTrailer
	_
	FeedUserContributedItem
	FeedUserContributionsItem
	FeedPageContributedItem
)
