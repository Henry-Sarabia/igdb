package igdb

//go:generate gomodifytags -file $GOFILE -struct Rate -add-tags json -w

type Rate struct {
	ID     int     `json:"id"`
	Rating float64 `json:"rating"`
	User   int     `json:"user"`
}
