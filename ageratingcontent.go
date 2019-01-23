package igdb

//go:generate gomodifytags -file $GOFILE -struct AgeRatingContent -add-tags json -w

type AgeRatingContent struct {
	Category    AgeRatingContentCategory `json:"category"`
	Description string                   `json:"description"`
}

//go:generate stringer -type=AgeRatingContentCategory

type AgeRatingContentCategory int

const (
	AgeRatingContentPEGI AgeRatingContentCategory = iota + 1
	AgeRatingContentESRB
)
