package igdb

//go:generate gomodifytags -file $GOFILE -struct GameVideo -add-tags json -w

// GameVideo represents a video associated with a particular game.
// For more information visit: https://api-docs.igdb.com/#game-video
type GameVideo struct {
	Game    int    `json:"game"`
	Name    string `json:"name"`
	VideoID string `json:"video_id"`
}
