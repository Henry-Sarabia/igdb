package main

// StatusCode codes
type StatusCode int

// DateCategory code
type DateCategory int

// GameCategory code
type GameCategory int

// Year is
type Year int

// Month is
type Month int

// Tag codes??
type Tag int //uint32

// BeatTime is the time to beat a given game measured in seconds.
type BeatTime struct {
	Hastly     int `json:"hastly"`
	Normally   int `json:"normally"`
	Completely int `json:"completely"`
}

// ReleaseDate hold information about date of release, platforms, and versions
type ReleaseDate struct {
	ID        ID           `json:"id"`
	Game      ID           `json:"game"`
	Category  DateCategory `json:"category"`
	Platform  ID           `json:"platform"`
	Human     string       `json:"human"`
	UpdatedAt int          `json:"updated_at"` //unix epoch??
	CreatedAt int          `json:"created_at"` //unix epoch??
	Date      int          `json:"date"`       //unix epoch
	Region    ID           `json:"region"`
	Year      Year         `json:"y"`
	Month     Month        `json:"m"`
}

// AltName is
type AltName struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// ESRB is
type ESRB struct {
	Rating   int    `json:"rating"`
	Synopsis string `json:"synopsis"`
}

// PEGI is
type PEGI struct {
	Rating   int    `json:"rating"`
	Synopsis string `json:"synopsis"`
}

// Website is
type Website struct {
	Category int `json:"category"` //codes
	URL      URL `json:"url"`
}

// Game is
type Game struct {
	ID                   ID            `json:"id"`
	Name                 string        `json:"name"`
	Slug                 string        `json:"slug"`
	URL                  string        `json:"url"`
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
	Covers               []Image       `json:"cover"`
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
