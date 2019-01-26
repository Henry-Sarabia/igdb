package igdb

//go:generate gomodifytags -file $GOFILE -struct TimeToBeat -add-tags json -w

// TimeToBeat represents the average completion times for a particular game.
// For more information: https://api-docs.igdb.com/#time-to-beat
type TimeToBeat struct {
	Completely int `json:"completely"`
	Game       int `json:"game"`
	Hastly     int `json:"hastly"`
	Normally   int `json:"normally"`
}
