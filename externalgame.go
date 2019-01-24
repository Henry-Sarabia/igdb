package igdb

//go:generate gomodifytags -file $GOFILE -struct ExternalGame -add-tags json -w

// ExternalGame contains the ID and other metadata for a game
// on a third party service.
// For more information visit: https://api-docs.igdb.com/#external-game
type ExternalGame struct {
	Category  ExternalGameCategory `json:"category"`
	CreatedAt int                  `json:"created_at"`
	Game      int                  `json:"game"`
	Name      string               `json:"name"`
	UID       string               `json:"uid"`
	UpdatedAt int                  `json:"updated_at"`
	Url       string               `json:"url"`
	Year      int                  `json:"year"`
}

//go:generate stringer -type=ExternalGameCategory

type ExternalGameCategory int

const (
	ExternalSteam = iota + 1
	_
	_
	_
	ExternalGOG
	_
	_
	_
	_
	ExternalYoutube
	ExternalMicrosoft
	_
	ExternalApple
	ExternalTwitch
	ExternalAndroid
)
