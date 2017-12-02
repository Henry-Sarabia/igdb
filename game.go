package igdb

// GameService handles all the API
// calls for the IGDB Game endpoint.
type GameService service

// Game contains information on an IGDB entry for a particular video game.
//
// For more information, visit: https://igdb.github.io/api/endpoints/game/
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
	VersionParent        int            `json:"version_parent"`
	VersionTitle         interface{}    `json:"version_title"`
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
func (gs *GameService) Get(id int, opts ...FuncOption) (*Game, error) {
	url, err := gs.client.singleURL(GameEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var g []Game

	err = gs.client.get(url, &g)
	if err != nil {
		return nil, err
	}

	return &g[0], nil
}

// List returns a list of Games identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Games based solely on the provided
// options. Any ID that does not match a Game is ignored. If none of the IDs
// match a Game, an error is returned.
func (gs *GameService) List(ids []int, opts ...FuncOption) ([]*Game, error) {
	url, err := gs.client.multiURL(GameEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var g []*Game

	err = gs.client.get(url, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Search returns a list of Games found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Games are found using the provided query, an error is returned.
func (gs *GameService) Search(qry string, opts ...FuncOption) ([]*Game, error) {
	url, err := gs.client.searchURL(GameEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var g []*Game

	err = gs.client.get(url, &g)
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
