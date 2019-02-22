package igdb

//go:generate gomodifytags -file $GOFILE -struct SocialMetric -add-tags json -w

type SocialMetric struct {
	Category           SocialMetricCategory `json:"category"`
	CreatedAt          int                  `json:"created_at"`
	SocialMetricSource int                  `json:"social_metric_source"`
	Value              int                  `json:"value"`
}

//go:generate stringer -type=SocialMetricCategory

type SocialMetricCategory int

const (
	SocialFollows SocialMetricCategory = iota + 1
	SocialLikes
	SocialHates
	SocialShares
	SocialViews
	SocialComments
	SocialFavorites
)
