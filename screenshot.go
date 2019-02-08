package igdb

//go:generate gomodifytags -file $GOFILE -struct Screenshot -add-tags json -w

// Screenshot represents a screenshot of a particular game.
// For more information visit: https://api-docs.igdb.com/#screenshot
type Screenshot struct {
	Image
	Game int `json:"game"`
}
