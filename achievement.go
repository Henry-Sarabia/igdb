package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Achievement -add-tags json -w

// Achievement data for specific games for specific platforms
// (currently limited to achievements from Steam, Playstation, and XBox).
// For more information visit: https://api-docs.igdb.com/#achievement
type Achievement struct {
	ID               int                 `json:"id"`
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

// AchievementRank specifies an achievement's rank ranging
// from bronze to platinum.
type AchievementRank int

//go:generate stringer -type=AchievementRank,AchievementCategory,AchievementLanguage

// Expected AchievementRank enums from the IGDB.
const (
	RankBronze AchievementRank = iota + 1
	RankSilver
	RankGold
	RankPlatinum
)

// AchievementCategory specifies an achievement's native platform.
type AchievementCategory int

// Expected AchievementCategory enums from the IGDB.
const (
	AchievementPlaystation AchievementCategory = iota + 1
	AchievementXbox
	AchievementSteam
)

// AchievementLanguage specifies an achievement's language.
type AchievementLanguage int

// Expected AchievementLanguage enums from the IGDB.
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

// AchievementService handles all the API calls for the IGDB
// Achievement endpoint.
// This endpoint is only available for the IGDB Pro tier or above.
type AchievementService service

// Get returns a single Achievement identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Achievements, an error is returned.
func (as *AchievementService) Get(id int, opts ...Option) (*Achievement, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ach []*Achievement

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := as.client.get(as.end, &ach, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Achievement with ID %v", id)
	}

	return ach[0], nil
}

// List returns a list of Achievements identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Achievement is ignored. If none of the IDs
// match a Achievement, an error is returned.
func (as *AchievementService) List(ids []int, opts ...Option) ([]*Achievement, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ach []*Achievement

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := as.client.get(as.end, &ach, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Achievements with IDs %v", ids)
	}

	return ach, nil
}

// Index returns an index of Achievements based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Achievements can
// be found using the provided options, an error is returned.
func (as *AchievementService) Index(opts ...Option) ([]*Achievement, error) {
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
func (as *AchievementService) Count(opts ...Option) (int, error) {
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
