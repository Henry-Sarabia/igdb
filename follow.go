package igdb

//go:generate gomodifytags -file $GOFILE -struct Follow -add-tags json -w

type Follow struct {
	ID   int `json:"id"`
	Game int `json:"game"`
	User int `json:"user"`
}
