package igdb

//go:generate gomodifytags -file $GOFILE -struct ListEntry -add-tags json -w

type ListEntry struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Game        int    `json:"game"`
	List        int    `json:"list"`
	Platform    int    `json:"platform"`
	Position    int    `json:"position"`
	Private     bool   `json:"private"`
	User        int    `json:"user"`
}
