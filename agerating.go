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
	AgeRatingESRB AgeRatingCategory = iota + 1
	AgeRatingPEGI
)

type AgeRatingEnum int

const (
	AgeRatingThree AgeRatingEnum = iota + 1
	AgeRatingSeven
	AgeRatingTwelve
	AgeRatingSixteen
	AgeRatingEighteen
	AgeRatingRP
	AgeRatingEC
	AgeRatingE
	AgeRatingE10
	AgeRatingT
	AgeRatingM
	AgeRatingAO
)
