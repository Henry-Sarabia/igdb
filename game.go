package igdb

// GameCategory corresponds to the IGDB
// enumerated type Game Category which
// describes the type of game content.
// GameCategory implements the Stringer
// interface.
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

// External contains information from
// an external service IDs.
type External struct {
	Steam string `json:"steam"`
}

// Game contains information about a game stored in the IGDB.
// See https://igdb.github.io/api/endpoints/game/ for more information.
type Game struct {
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
	Slug                 string        `json:"slug"`
	URL                  URL           `json:"url"`
	CreatedAt            int           `json:"created_at"` //unix epoch
	UpdatedAt            int           `json:"updated_at"` //unix epoch
	Summary              string        `json:"summary"`
	Storyline            string        `json:"storyline"`
	Collection           int           `json:"collection"`
	Franchise            int           `json:"franchise"`
	Hypes                int           `json:"hypes"`
	Popularity           float64       `json:"popularity"`
	Rating               float64       `json:"rating"`
	RatingCount          int           `json:"rating_count"`
	AggregateRating      float64       `json:"aggregated_rating"`
	AggregateRatingCount int           `json:"aggregated_rating_count"`
	TotalRating          float64       `json:"total_rating"`
	TotalRatingCount     int           `json:"total_rating_count"`
	WeightedRating       float64       `json:"weighted_rating"`
	Game                 int           `json:"game"`
	Developers           []int         `json:"developers"`
	Publishers           []int         `json:"publishers"`
	Engines              []int         `json:"game_engines"`
	Category             GameCategory  `json:"category"`
	TimeToBeat           BeatTime      `json:"time_to_beat"`
	PlayerPerspectives   []int         `json:"player_perspectives"`
	GameModes            []int         `json:"game_modes"`
	Keywords             []int         `json:"keywords"`
	Themes               []int         `json:"themes"`
	Genres               []int         `json:"genres"`
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
	DLCs                 []int         `json:"dlcs"`
	Expansions           []int         `json:"expansions"`
	Standalone           []int         `json:"standalone_expansions"`
	Bundles              []int         `json:"bundles"`
	SimilarGames         []int         `json:"games"`
	Follows              interface{}   `json:"follows"`
	PulseCount           interface{}   `json:"pulse_count"`
	External             External      `json:"external"`
	MultiplayerModes     interface{}   `json:"multiplayer_modes"`
	Franchises           []int         `json:"franchises"`
}

// GetGame gets IGDB information for a game identified by their unique IGDB ID.
func (c *Client) GetGame(id int, opts ...OptionFunc) (*Game, error) {
	url, err := c.singleURL(GameEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var g []Game

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// GetGames gets IGDB information for a list of games identified by a list of their unique IGDB IDs.
func (c *Client) GetGames(ids []int, opts ...OptionFunc) ([]*Game, error) {
	url, err := c.multiURL(GameEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var g []*Game

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// SearchGames searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchGames(qry string, opts ...OptionFunc) ([]*Game, error) {
	url, err := c.searchURL(GameEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var g []*Game

	err = c.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// GameCategory implements the stringer interface
// by matching its code with the IGDBs enumerated type
// Game Category and returns the category as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/game-category/
func (g GameCategory) String() string {
	switch g {
	case 0:
		return "Main Game"
	case 1:
		return "DLC / Addon"
	case 2:
		return "Expansion"
	case 3:
		return "Bundle"
	case 4:
		return "Standalone Expansion"
	default:
		return "Undefined"
	}
}
