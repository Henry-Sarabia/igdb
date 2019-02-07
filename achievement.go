package igdb

import (
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct Achievement -add-tags json -w

// AchievementService handles all the API calls for the IGDB
// Achievement endpoint.
// This endpoint is only available for the IGDB Pro tier or above.
type AchievementService service

// Achievement data for specific games for specific platforms
// (currently limited to achievements from Steam, Playstation, and XBox).
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

// AchievementRank specifies an achievement's rank ranging
// from bronze to platinum.
type AchievementRank int

const (
	RankBronze AchievementRank = iota + 1
	RankSilver
	RankGold
	RankPlatinum
)

// AchievementCategory specifies an achievement's native platform.
type AchievementCategory int

const (
	AchievementPlaystation AchievementCategory = iota + 1
	AchievementXbox
	AchievementSteam
)

// AchievementLanguage specifices an achievement's language.
type AchievementLanguage int

const (
	LanguageEurope AchievementLanguage = iota + 1
	LanguageNorthAmerica
	LanguageAustralia
	LanguageNewZealand
	LanguageJapan
	LanguageChina
	LanguageAsia
	LanguageWorldwide
	LanguageHongKong
	LanguageSouthKorea
)

// Index returns an index of Achievements based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Achievements can
// be found using the provided options, an error is returned.
func (as *AchievementService) Index(opts ...FuncOption) ([]*Achievement, error) {
	var ach []*Achievement

	err := as.client.get(as.end, &ach, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Achievements")
	}

	return ach, nil
}

// Count returns the number of Achievements available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Achievements to count.
func (as *AchievementService) Count(opts ...FuncOption) (int, error) {
	ct, err := as.client.getCount(as.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Achievements")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Achievement object.
func (as *AchievementService) Fields() ([]string, error) {
	f, err := as.client.getFields(as.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Achievement fields")
	}

	return f, nil
}
