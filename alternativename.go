package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct AlternativeName -add-tags json -w

type AlternativeName struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
	Game    int    `json:"game"`
	Name    string `json:"name"`
}

// AlternativeNameService handles all the API calls for the IGDB AlternativeName endpoint.
type AlternativeNameService service

// Get returns a single AlternativeName identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any AlternativeNames, an error is returned.
func (as *AlternativeNameService) Get(id int, opts ...Option) (*AlternativeName, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var alt []*AlternativeName

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := as.client.get(as.end, &alt, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AlternativeName with ID %v", id)
	}

	return alt[0], nil
}

// List returns a list of AlternativeNames identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a AlternativeName is ignored. If none of the IDs
// match a AlternativeName, an error is returned.
func (as *AlternativeNameService) List(ids []int, opts ...Option) ([]*AlternativeName, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var alt []*AlternativeName

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := as.client.get(as.end, &alt, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AlternativeNames with IDs %v", ids)
	}

	return alt, nil
}

// Index returns an index of AlternativeNames based solely on the provided functional
// options used to sort, filter, and paginate the results. If no AlternativeNames can
// be found using the provided options, an error is returned.
func (as *AlternativeNameService) Index(opts ...Option) ([]*AlternativeName, error) {
	var alt []*AlternativeName

	err := as.client.get(as.end, &alt, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of AlternativeNames")
	}

	return alt, nil
}

// Count returns the number of AlternativeNames available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which AlternativeNames to count.
func (as *AlternativeNameService) Count(opts ...Option) (int, error) {
	ct, err := as.client.getCount(as.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count AlternativeNames")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB AlternativeName object.
func (as *AlternativeNameService) Fields() ([]string, error) {
	f, err := as.client.getFields(as.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AlternativeName fields")
	}

	return f, nil
}
