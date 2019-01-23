package igdb

//go:generate gomodifytags -file $GOFILE -struct AchievementIcon -add-tags json -w

// AchievementIcon is an icon for a specific achievement.
// For more information visit: https://api-docs.igdb.com/#achievement-icon
type AchievementIcon struct {
	AlphaChannel bool   `json:"alpha_channel"`
	Animated     bool   `json:"animated"`
	Height       int    `json:"height"`
	ImageID      string `json:"image_id"`
	URL          string `json:"url"`
	Width        int    `json:"width"`
}
