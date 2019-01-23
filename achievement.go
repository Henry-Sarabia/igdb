package igdb

//go:generate gomodifytags -file $GOFILE -struct Achievement -add-tags json -w

// Achievement data for specific games for specific platforms
// (currently limited to achievements from Steam, Playstation,
// and XBox).
// For more information visit: https://api-docs.igdb.com/#achievement
type Achievement struct {
	AchievementIcon  int                 `json:"achievement_icon"`
	Category         AchievementCategory `json:"category"`
	CreatedAt        int                 `json:"created_at"`
	Description      string              `json:"description"`
	ExternalID       string              `json:"external_id"`
	Game             int                 `json:"game"`
	Language         AchievementLanguage `json:"language"`
	Name             string              `json:"name"`
	OwnersPercentage float64             `json:"owners_percentage"`
	Rank             AchievementRank     `json:"rank"`
	Slug             string              `json:"slug"`
	Tags             []Tag               `json:"tags"`
	UpdatedAt        int                 `json:"updated_at"`
}

//go:generate stringer -type=AchievementRank,AchievementCategory,AchievementLanguage

type AchievementRank int

const (
	Bronze AchievementRank = iota + 1
	Silver
	Gold
	Platinum
)

type AchievementCategory int

const (
	Playstation AchievementCategory = iota + 1
	Xbox
	Steam
)

type AchievementLanguage int

const (
	Europe AchievementLanguage = iota + 1
	NorthAmerica
	Australia
	NewZealand
	Japan
	China
	Asia
	Worldwide
	HongKong
	SouthKorea
)
