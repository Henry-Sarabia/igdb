package igdb

import (
	"strconv"
)

// GameService handles all the API
// calls for the IGDB Game endpoint.
type GameService service

// Game contains information on an IGDB entry for a particular video game.
//
// For more information, visit: https://igdb.github.io/api/endpoints/game/
type Game struct {
	ID                   int            `json:"id,omitempty"`
	Name                 string         `json:"name,omitempty"`
	Slug                 string         `json:"slug,omitempty"`
	URL                  URL            `json:"url,omitempty"`
	CreatedAt            int            `json:"created_at,omitempty"` // Unix time in milliseconds
	UpdatedAt            int            `json:"updated_at,omitempty"` // Unix time in milliseconds
	Summary              string         `json:"summary,omitempty"`
	Storyline            string         `json:"storyline,omitempty"`
	Collection           int            `json:"collection,omitempty"`
	Franchise            int            `json:"franchise,omitempty"`
	Hypes                int            `json:"hypes,omitempty"`
	Popularity           float64        `json:"popularity,omitempty"`
	Rating               float64        `json:"rating,omitempty"`
	RatingCount          int            `json:"rating_count,omitempty"`
	AggregateRating      float64        `json:"aggregated_rating,omitempty"`
	AggregateRatingCount int            `json:"aggregated_rating_count,omitempty"`
	TotalRating          float64        `json:"total_rating,omitempty"`
	TotalRatingCount     int            `json:"total_rating_count,omitempty"`
	WeightedRating       float64        `json:"weighted_rating,omitempty"`
	Game                 int            `json:"game,omitempty"`
	VersionParent        int            `json:"version_parent,omitempty"`
	VersionTitle         interface{}    `json:"version_title,omitempty"`
	Developers           []int          `json:"developers,omitempty"`
	Publishers           []int          `json:"publishers,omitempty"`
	Engines              []int          `json:"game_engines,omitempty"`
	Category             GameCategory   `json:"category,omitempty"`
	TimeToBeat           CompletionTime `json:"time_to_beat,omitempty"`
	PlayerPerspectives   []int          `json:"player_perspectives,omitempty"`
	GameModes            []int          `json:"game_modes,omitempty"`
	Keywords             []int          `json:"keywords,omitempty"`
	Themes               []int          `json:"themes,omitempty"`
	Genres               []int          `json:"genres,omitempty"`
	FirstReleaseDate     int            `json:"first_release_date,omitempty"` // Unix time in milliseconds
	Status               GameStatus     `json:"status,omitempty"`
	AlternativeNames     []AltName      `json:"alternative_names,omitempty"`
	Screenshots          []Image        `json:"screenshots,omitempty"`
	Videos               []YoutubeVideo `json:"videos,omitempty"`
	Cover                Image          `json:"cover,omitempty"`
	ESRB                 ESRB           `json:"esrb,omitempty"`
	PEGI                 PEGI           `json:"pegi,omitempty"`
	Websites             []Website      `json:"websites,omitempty"`
	Tags                 []Tag          `json:"tags,omitempty"`
	DLCs                 []int          `json:"dlcs,omitempty"`
	Expansions           []int          `json:"expansions,omitempty"`
	Standalone           []int          `json:"standalone_expansions,omitempty"`
	Bundles              []int          `json:"bundles,omitempty"`
	SimilarGames         []int          `json:"games,omitempty"`
	Follows              interface{}    `json:"follows,omitempty"`
	PulseCount           interface{}    `json:"pulse_count,omitempty"`
	External             External       `json:"external,omitempty"`
	MultiplayerModes     interface{}    `json:"multiplayer_modes,omitempty"`
	Franchises           []int          `json:"franchises,omitempty"`
	Platforms            []int          `json:"platforms,omitempty"`
}

// AltName contains information on an
// alternative name for an IGDB object.
type AltName struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// CompletionTime contains the time to complete
// a particular video game. This time is measured
// in seconds.
type CompletionTime struct {
	Hastly     int `json:"hastly"`
	Normally   int `json:"normally"`
	Completely int `json:"completely"`
}

// ESRB contains the rating and synopsis
// for a particular video game given by
// the Entertainment Software Rating Board.
type ESRB struct {
	Rating   ESRBCode `json:"rating"`
	Synopsis string   `json:"synopsis"`
}

// External contains information for
// connecting external service IDs to
// the IGDB for a particular object.
type External struct {
	Steam string `json:"steam"`
}

// PEGI contains the rating and synopsis
// for a particular video game given by
// the Pan European Game Information organization.
type PEGI struct {
	Rating   PEGICode `json:"rating"`
	Synopsis string   `json:"synopsis"`
}

// YoutubeVideo contains the name and
// ID of a  Youtube video.
type YoutubeVideo struct {
	Name string `json:"name"`
	ID   string `json:"video_id"` // Youtube slug
}

// Website contains address and category
// information on a website referenced
// in the IGDB.
type Website struct {
	Category WebsiteCategory `json:"category"`
	URL      URL             `json:"url"`
}

// Get returns a single Game identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Games, an error is returned.
//func (gs *GameService) Get(id int, opts ...FuncOption) (*Game, error) {
//	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
//	req, err := gs.client.request(GameEndpoint, opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	var g []*Game
//
//	err = gs.client.send(req, &g)
//	if err != nil {
//		return nil, errors.Wrap(err, "cannot make request")
//	}
//
//	return g[0], nil
//}
func (gs *GameService) Get(id int, opts ...FuncOption) (*Game, error) {
	var g []*Game

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(GameEndpoint, &g, opts...)
	if err != nil {
		return nil, err
	}

	return g[0], nil
}

// List returns a list of Games identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of Games based solely on the provided
// options. Any ID that does not match a Game is ignored. If none of the IDs
// match a Game, an error is returned.
func (gs *GameService) List(ids []int, opts ...FuncOption) ([]*Game, error) {
	req, err := gs.client.request(GameEndpoint, opts...)
	if err != nil {
		return nil, err
	}

	var g []*Game

	err = gs.client.send(req, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Search returns a list of Games found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Games are found using the provided query, an error is returned.
//TODO: remember that Search also has its own endpoint
func (gs *GameService) Search(qry string, opts ...FuncOption) ([]*Game, error) {
	opts = append(opts, setSearch(qry))
	req, err := gs.client.request(GameEndpoint, opts...)
	if err != nil {
		return nil, err
	}

	var g []*Game

	err = gs.client.send(req, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Count returns the number of Games available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Games to count.
func (gs *GameService) Count(opts ...FuncOption) (int, error) {
	ct, err := gs.client.getEndpointCount(GameEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Game object.
func (gs *GameService) ListFields() ([]string, error) {
	fl, err := gs.client.getEndpointFieldList(GameEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
