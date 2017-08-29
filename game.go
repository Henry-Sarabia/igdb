package igdb

import "net/url"

// GameCategory is a code associated with a game
// denoting its category. The codes are as follows:
// 0 - Main Game
// 1 - DLC / Addon
// 2 - Expansion
// 3 - Bundle
// 4 - Standalone Expansion
type GameCategory int

// BeatTime is the time to beat a game
// measured in seconds.
type BeatTime struct {
	Hastly     int `json:"hastly"`
	Normally   int `json:"normally"`
	Completely int `json:"completely"`
}

// ESRB containts the rating and synopsis
// of a game from the Entertainment Software
// Rating Board.
type ESRB struct {
	Rating   int    `json:"rating"`
	Synopsis string `json:"synopsis"`
}

// PEGI contains the rating and synopsis
// of a game from the Pan European Game
// Information system.
type PEGI struct {
	Rating   int    `json:"rating"`
	Synopsis string `json:"synopsis"`
}

// Website contains information about
// a website referenced in the IGDB.
type Website struct {
	Category int `json:"category"` //codes
	URL      URL `json:"url"`
}

// Game contains information about a game stored in the IGDB.
// See https://igdb.github.io/api/endpoints/game/ for more information.
type Game struct {
	ID                   ID            `json:"id"`
	Name                 string        `json:"name"`
	Slug                 string        `json:"slug"`
	URL                  URL           `json:"url"`
	CreatedAt            int           `json:"created_at"` //unix epoch
	UpdatedAt            int           `json:"updated_at"` //unix epoch
	Summary              string        `json:"summary"`
	Storyline            string        `json:"storyline"`
	Collection           ID            `json:"collection"`
	Franchise            ID            `json:"franchise"`
	Hypes                int           `json:"hypes"`
	Popularity           float64       `json:"popularity"`
	Rating               float64       `json:"rating"`
	RatingCount          int           `json:"raing_count"`
	AggregateRating      float64       `json:"aggregated_rating"`
	AggregateRatingCount int           `json:"aggregated_rating_count"`
	TotalRating          float64       `json:"total_rating"`
	TotalRatingCount     int           `json:"total_rating_count"`
	WeightedRating       float64       `json:"weighted_rating"`
	Game                 ID            `json:"game"`
	Developers           []ID          `json:"developers"`
	Publishers           []ID          `json:"publishers"`
	Engines              []ID          `json:"game_engines"`
	Category             GameCategory  `json:"category"`
	TimeToBeat           BeatTime      `json:"time_to_beat"`
	PlayerPerspectives   []ID          `json:"player_perspectives"`
	GameModes            []ID          `json:"game_modes"`
	Keywords             []ID          `json:"keywords"`
	Themes               []ID          `json:"themes"`
	Genres               []ID          `json:"genres"`
	FirstReleaseDate     int           `json:"first_release_date"` //unix epoch
	Status               StatusCode    `json:"status"`
	ReleaseDates         []ReleaseDate `json:"release_dates"`
	AlternativeNames     []AltName     `json:"alternative_names"`
	Screenshots          []Image       `json:"screenshots"`
	Videos               []Video       `json:"videos"`
	Covers               Image         `json:"cover"`
	ESRB                 ESRB          `json:"esrb"`
	PEGI                 PEGI          `json:"pegi"`
	Websites             []Website     `json:"websites"`
	Tags                 []Tag         `json:"tags"`
	DLCs                 []ID          `json:"dlcs"`
	Expansions           []ID          `json:"expansions"`
	Standalone           []ID          `json:"standalone_expansions"`
	Bundles              []ID          `json:"bundles"`
	SimilarGames         []ID          `json:"games"`
}

// GetGame gets IGDB information for a game identified by their unique IGDB ID.
func (c *Client) GetGame(id ID, opts ...OptionFunc) (*Game, error) {
	opt := Options{Values: url.Values{}}

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := rootURL + "games/" + id.toString()
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var g []Game

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// SearchGames searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchGames(qry string, opts ...OptionFunc) ([]*Game, error) {
	opt := Options{Values: url.Values{}}

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := rootURL + "games/?search=" + qry + "&fields=*"
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var g []*Game

	err := c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
