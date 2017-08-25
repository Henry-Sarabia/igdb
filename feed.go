package igdb

// FeedCategory codes
type FeedCategory int

// Feed is
type Feed struct {
	ID        ID           `json:"id"`
	URL       URL          `json:"url"`
	CreatedAt int          `json:"created_at"` //unix epoch
	UpdatedAt int          `json:"updated_at"` //unix epoch
	Content   string       `json:"content"`
	Category  FeedCategory `json:"category"`
	User      ID           `json:"user"`
	Games     []ID         `json:"games"`
	Title     string       `json:"title"`
	LikeCount int          `json:"feed_likes_count"`
}
