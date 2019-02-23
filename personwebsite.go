package igdb

//go:generate gomodifytags -file $GOFILE -struct Zype -add-tags json -w

type PersonWebsite struct {
	ID       int             `json:"id"`
	Category WebsiteCategory `json:"category"`
	Trusted  bool            `json:"trusted"`
	URL      string          `json:"url"`
}
