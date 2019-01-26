package igdb

//go:generate gomodifytags -file $GOFILE -struct MultiplayerMode -add-tags json -w

// MultiplayerMode contains data about the supported multiplayer types.
// For more information visit: https://api-docs.igdb.com/#multiplayer-mode
type MultiplayerMode struct {
	Campaigncoop      bool `json:"campaigncoop"`
	Dropin            bool `json:"dropin"`
	Lancoop           bool `json:"lancoop"`
	Offlinecoop       bool `json:"offlinecoop"`
	Offlinecoopmax    int  `json:"offlinecoopmax"`
	Offlinemax        int  `json:"offlinemax"`
	Onlinecoop        bool `json:"onlinecoop"`
	Onlinecoopmax     int  `json:"onlinecoopmax"`
	Onlinemax         int  `json:"onlinemax"`
	Platform          int  `json:"platform"`
	Splitscreen       bool `json:"splitscreen"`
	Splitscreenonline bool `json:"splitscreenonline"`
}
