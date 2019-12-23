package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Game -add-tags json -w

// Game contains information on an IGDB entry for a particular video game.
// For more information visit: https://api-docs.igdb.com/#game
type Game struct {
	ID                    int          `json:"id,omitempty"`
	AgeRatings            []int        `json:"age_ratings,omitempty"`
	AggregatedRating      float64      `json:"aggregated_rating,omitempty"`
	AggregatedRatingCount int          `json:"aggregated_rating_count,omitempty"`
	AlternativeNames      []int        `json:"alternative_names,omitempty"`
	Artworks              []int        `json:"artworks,omitempty"`
	Bundles               []int        `json:"bundles,omitempty"`
	Category              GameCategory `json:"category,omitempty"`
	Collection            int          `json:"collection,omitempty"`
	Cover                 int          `json:"cover,omitempty"`
	CreatedAt             int          `json:"created_at,omitempty"`
	DLCS                  []int        `json:"dlcs,omitempty"`
	Expansions            []int        `json:"expansions,omitempty"`
	ExternalGames         []int        `json:"external_games,omitempty"`
	FirstReleaseDate      int          `json:"first_release_date,omitempty"`
	Follows               int          `json:"follows,omitempty"`
	Franchise             int          `json:"franchise,omitempty"`
	Franchises            []int        `json:"franchises,omitempty"`
	GameEngines           []int        `json:"game_engines,omitempty"`
	GameModes             []int        `json:"game_modes,omitempty"`
	Genres                []int        `json:"genres,omitempty"`
	Hypes                 int          `json:"hypes,omitempty"`
	InvolvedCompanies     []int        `json:"involved_companies,omitempty"`
	Keywords              []int        `json:"keywords,omitempty"`
	MultiplayerModes      []int        `json:"multiplayer_modes,omitempty"`
	Name                  string       `json:"name,omitempty"`
	ParentGame            int          `json:"parent_game,omitempty"`
	Platforms             []int        `json:"platforms,omitempty"`
	PlayerPerspectives    []int        `json:"player_perspectives,omitempty"`
	Popularity            float64      `json:"popularity,omitempty"`
	PulseCount            int          `json:"pulse_count,omitempty"`
	Rating                float64      `json:"rating,omitempty"`
	RatingCount           int          `json:"rating_count,omitempty"`
	ReleaseDates          []int        `json:"release_dates,omitempty"`
	Screenshots           []int        `json:"screenshots,omitempty"`
	SimilarGames          []int        `json:"similar_games,omitempty"`
	Slug                  string       `json:"slug,omitempty"`
	StandaloneExpansions  []int        `json:"standalone_expansions,omitempty"`
	Status                GameStatus   `json:"status,omitempty"`
	Storyline             string       `json:"storyline,omitempty"`
	Summary               string       `json:"summary,omitempty"`
	Tags                  []Tag        `json:"tags,omitempty"`
	Themes                []int        `json:"themes,omitempty"`
	TimeToBeat            int          `json:"time_to_beat,omitempty"`
	TotalRating           float64      `json:"total_rating,omitempty"`
	TotalRatingCount      int          `json:"total_rating_count,omitempty"`
	UpdatedAt             int          `json:"updated_at,omitempty"`
	URL                   string       `json:"url,omitempty"`
	VersionParent         int          `json:"version_parent,omitempty"`
	VersionTitle          string       `json:"version_title,omitempty"`
	Videos                []int        `json:"videos,omitempty"`
	Websites              []int        `json:"websites,omitempty"`
}

// GameCategory specifies a type of game content.
type GameCategory int

//go:generate stringer -type=GameCategory,GameStatus

// Expected GameCategory enums from the IGDB.
const (
	MainGame GameCategory = iota
	DLCAddon
	Expansion
	Bundle
	StandaloneExpansion
)

// GameStatus specifies the release status of a specific game.
type GameStatus int

// Expected GameStatus enums from the IGDB.
const (
	StatusReleased GameStatus = iota
	_
	StatusAlpha
	StatusBeta
	StatusEarlyAccess
	StatusOffline
	StatusCancelled
)

// GameService handles all the API
// calls for the IGDB Game endpoint.
type GameService service

// Get returns a single Game identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Games, an error is returned.
func (gs *GameService) Get(id int, opts ...Option) (*Game, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var g []*Game

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Game with ID %v", id)
	}

	return g[0], nil
}

// List returns a list of Games identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Game is ignored. If none of the IDs
// match a Game, an error is returned.
func (gs *GameService) List(ids []int, opts ...Option) ([]*Game, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var g []*Game

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := gs.client.get(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Games with IDs %v", ids)
	}

	return g, nil
}

// Index returns an index of Games based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Games can
// be found using the provided options, an error is returned.
func (gs *GameService) Index(opts ...Option) ([]*Game, error) {
	var g []*Game

	err := gs.client.get(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Games")
	}

	return g, nil
}

// Search returns a list of Games found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Games are found using the provided query, an error is returned.
func (gs *GameService) Search(qry string, opts ...Option) ([]*Game, error) {
	var g []*Game

	opts = append(opts, setSearch(qry))
	err := gs.client.get(gs.end, &g, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Game with query %s", qry)
	}

	return g, nil
}

// Count returns the number of Games available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Games to count.
func (gs *GameService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Games")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Game object.
func (gs *GameService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Game fields")
	}

	return f, nil
}
