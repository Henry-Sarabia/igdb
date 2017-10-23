package igdb

// AltName contains information on an
// alternative name for an IGDB object.
type AltName struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// CompletionTime contains the time to complete
// a particular video game measured in seconds.
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
// ID for a particular Youtube video.
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

// Game contains information on an IGDB
// entry for a particular video game.
type Game struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	Slug                 string         `json:"slug"`
	URL                  URL            `json:"url"`
	CreatedAt            int            `json:"created_at"` // Unix time in milliseconds
	UpdatedAt            int            `json:"updated_at"` // Unix time in milliseconds
	Summary              string         `json:"summary"`
	Storyline            string         `json:"storyline"`
	Collection           int            `json:"collection"`
	Franchise            int            `json:"franchise"`
	Hypes                int            `json:"hypes"`
	Popularity           float64        `json:"popularity"`
	Rating               float64        `json:"rating"`
	RatingCount          int            `json:"rating_count"`
	AggregateRating      float64        `json:"aggregated_rating"`
	AggregateRatingCount int            `json:"aggregated_rating_count"`
	TotalRating          float64        `json:"total_rating"`
	TotalRatingCount     int            `json:"total_rating_count"`
	WeightedRating       float64        `json:"weighted_rating"`
	Game                 int            `json:"game"`
	Developers           []int          `json:"developers"`
	Publishers           []int          `json:"publishers"`
	Engines              []int          `json:"game_engines"`
	Category             GameCategory   `json:"category"`
	TimeToBeat           CompletionTime `json:"time_to_beat"`
	PlayerPerspectives   []int          `json:"player_perspectives"`
	GameModes            []int          `json:"game_modes"`
	Keywords             []int          `json:"keywords"`
	Themes               []int          `json:"themes"`
	Genres               []int          `json:"genres"`
	FirstReleaseDate     int            `json:"first_release_date"` // Unix time in milliseconds
	Status               GameStatus     `json:"status"`
	ReleaseDates         []ReleaseDate  `json:"release_dates"`
	AlternativeNames     []AltName      `json:"alternative_names"`
	Screenshots          []Image        `json:"screenshots"`
	Videos               []YoutubeVideo `json:"videos"`
	Covers               Image          `json:"cover"`
	ESRB                 ESRB           `json:"esrb"`
	PEGI                 PEGI           `json:"pegi"`
	Websites             []Website      `json:"websites"`
	Tags                 []Tag          `json:"tags"`
	DLCs                 []int          `json:"dlcs"`
	Expansions           []int          `json:"expansions"`
	Standalone           []int          `json:"standalone_expansions"`
	Bundles              []int          `json:"bundles"`
	SimilarGames         []int          `json:"games"`
	Follows              interface{}    `json:"follows"`
	PulseCount           interface{}    `json:"pulse_count"`
	External             External       `json:"external"`
	MultiplayerModes     interface{}    `json:"multiplayer_modes"`
	Franchises           []int          `json:"franchises"`
	Platforms            []int          `json:"platforms"`
}

// GetGame returns a single Game identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will
// not have an effect due to GetGame only returning a single Game
// object and not a list of Games.
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

// GetGames returns a list of Games identified by the provided list of IGDB
// IDs. Provide functional options to filter, sort, and paginate the results.
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

// SearchGames returns a list of Games found by searching the IGDB using the provided
// query. Provide functional options to filter, sort, and paginate the results.
// Providing an empty query will instead retrieve an index of Collections based
// solely on the provided options.
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
