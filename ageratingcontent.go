package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct AgeRatingContent -add-tags json -w

// AgeRatingContent is the organization behind a specific rating.
type AgeRatingContent struct {
	ID          int                      `json:"id"`
	Category    AgeRatingContentCategory `json:"category"`
	Description string                   `json:"description"`
}

// AgeRatingContentCategory specifies a regulatory organization.
type AgeRatingContentCategory int

//go:generate stringer -type=AgeRatingContentCategory

// Expected AgeRatingContentCategory enums from the IGDB.
const (
	AgeRatingContentPEGI AgeRatingContentCategory = iota + 1
	AgeRatingContentESRB
)

// AgeRatingContentService handles all the API calls for the IGDB AgeRatingContent endpoint.
type AgeRatingContentService service

// Get returns a single AgeRatingContent identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any AgeRatingContents, an error is returned.
func (as *AgeRatingContentService) Get(id int, opts ...Option) (*AgeRatingContent, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var cont []*AgeRatingContent

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := as.client.post(as.end, &cont, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AgeRatingContent with ID %v", id)
	}

	return cont[0], nil
}

// List returns a list of AgeRatingContents identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a AgeRatingContent is ignored. If none of the IDs
// match a AgeRatingContent, an error is returned.
func (as *AgeRatingContentService) List(ids []int, opts ...Option) ([]*AgeRatingContent, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var cont []*AgeRatingContent

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := as.client.post(as.end, &cont, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AgeRatingContents with IDs %v", ids)
	}

	return cont, nil
}

// Index returns an index of AgeRatingContents based solely on the provided functional
// options used to sort, filter, and paginate the results. If no AgeRatingContents can
// be found using the provided options, an error is returned.
func (as *AgeRatingContentService) Index(opts ...Option) ([]*AgeRatingContent, error) {
	var cont []*AgeRatingContent

	err := as.client.post(as.end, &cont, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of AgeRatingContents")
	}

	return cont, nil
}

// Count returns the number of AgeRatingContents available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which AgeRatingContents to count.
func (as *AgeRatingContentService) Count(opts ...Option) (int, error) {
	ct, err := as.client.getCount(as.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count AgeRatingContents")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB AgeRatingContent object.
func (as *AgeRatingContentService) Fields() ([]string, error) {
	f, err := as.client.getFields(as.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AgeRatingContent fields")
	}

	return f, nil
}
