package igdb

//go:generate gomodifytags -file $GOFILE -struct AlternativeName -add-tags json -w

type AlternativeName struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
}
