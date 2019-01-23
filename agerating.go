package igdb

//go:generate gomodifytags -file $GOFILE -struct AgeRating -add-tags json -w

// AgeRating describes an age rating according to various organizations.
// For more information visit: https://api-docs.igdb.com/#age-rating
type AgeRating struct {
	Category            AgeRatingCategory `json:"category"`
	ContentDescriptions []int             `json:"content_descriptions"`
	Rating              AgeRatingEnum     `json:"rating"`
	RatingCoverURL      string            `json:"rating_cover_url"`
	Synopsis            string            `json:"synopsis"`
}

//go:generate stringer -type=AgeRatingCategory,AgeRatingEnum

type AgeRatingCategory int

const (
	ESRB AgeRatingCategory = iota + 1
	PEGI
)

type AgeRatingEnum int

const (
	Three AgeRatingEnum = iota + 1
	Seven
	Twelve
	Sixteen
	Eighteen
	RP
	EC
	E
	E10
	T
	M
	AO
)
