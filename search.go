package igdb

//go:generate gomodifytags -file $GOFILE -struct Search -add-tags json -w

// Search is it's own endpoint and also contains: Characters, Collections
// Games, People, Platforms, and Themes.
type Search struct {
	AlternativeName string  `json:"alternative_name"`
	Character       int     `json:"character"`
	Collection      int     `json:"collection"`
	Company         int     `json:"company"`
	Description     string  `json:"description"`
	Game            int     `json:"game"`
	Name            string  `json:"name"`
	Person          int     `json:"person"`
	Platform        int     `json:"platform"`
	Popularity      float64 `json:"popularity"`
	PublishedAt     int     `json:"published_at"`
	TestDummy       int     `json:"test_dummy"`
	Theme           int     `json:"theme"`
}
