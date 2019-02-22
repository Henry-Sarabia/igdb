package igdb

//go:generate gomodifytags -file $GOFILE -struct ReviewVideo -add-tags json -w

type ReviewVideo struct {
	Trusted bool   `json:"trusted"`
	URL     string `json:"url"`
}
