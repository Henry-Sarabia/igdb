package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct AgeRating -add-tags json -w

// AgeRating describes an age rating according to various organizations.
// For more information visit: https://api-docs.igdb.com/#age-rating
type AgeRating struct {
	ID                  int               `json:"id"`
	Category            AgeRatingCategory `json:"category"`
	ContentDescriptions []int             `json:"content_descriptions"`
	Rating              AgeRatingEnum     `json:"rating"`
	RatingCoverURL      string            `json:"rating_cover_url"`
	Synopsis            string            `json:"synopsis"`
}

// AgeRatingCategory specifies a regulatory organization.
type AgeRatingCategory int

//go:generate stringer -type=AgeRatingCategory,AgeRatingEnum

// Expected AgeRatingCategory enums from the IGDB.
const (
	AgeRatingESRB AgeRatingCategory = iota + 1
	AgeRatingPEGI
)

// AgeRatingEnum specifies a specific age rating.
type AgeRatingEnum int

// Expected AgeRatingEnum enums from the IGDB.
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

// AgeRatingService handles all the API calls for the IGDB AgeRating endpoint.
type AgeRatingService service

// Get returns a single AgeRating identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any AgeRatings, an error is returned.
func (as *AgeRatingService) Get(id int, opts ...Option) (*AgeRating, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var age []*AgeRating

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := as.client.post(as.end, &age, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AgeRating with ID %v", id)
	}

	return age[0], nil
}

// List returns a list of AgeRatings identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a AgeRating is ignored. If none of the IDs
// match a AgeRating, an error is returned.
func (as *AgeRatingService) List(ids []int, opts ...Option) ([]*AgeRating, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var age []*AgeRating

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := as.client.post(as.end, &age, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AgeRatings with IDs %v", ids)
	}

	return age, nil
}

// Index returns an index of AgeRatings based solely on the provided functional
// options used to sort, filter, and paginate the results. If no AgeRatings can
// be found using the provided options, an error is returned.
func (as *AgeRatingService) Index(opts ...Option) ([]*AgeRating, error) {
	var age []*AgeRating

	err := as.client.post(as.end, &age, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of AgeRatings")
	}

	return age, nil
}

// Count returns the number of AgeRatings available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which AgeRatings to count.
func (as *AgeRatingService) Count(opts ...Option) (int, error) {
	ct, err := as.client.getCount(as.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count AgeRatings")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB AgeRating object.
func (as *AgeRatingService) Fields() ([]string, error) {
	f, err := as.client.getFields(as.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AgeRating fields")
	}

	return f, nil
}
