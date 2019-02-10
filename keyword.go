package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Keyword -add-tags json -w

// Keyword represents a word or phrase that get tagged to a game
// such as "World War 2" or "Steampunk".
// For more information visit: https://api-docs.igdb.com/#keyword
type Keyword struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	Url       string `json:"url"`
}

// KeywordService handles all the API calls for the IGDB Keyword endpoint.
type KeywordService service

// Get returns a single Keyword identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Keywords, an error is returned.
func (ks *KeywordService) Get(id int, opts ...FuncOption) (*Keyword, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var key []*Keyword

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ks.client.get(ks.end, &key, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Keyword with ID %v", id)
	}

	return key[0], nil
}

// List returns a list of Keywords identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Keyword is ignored. If none of the IDs
// match a Keyword, an error is returned.
func (ks *KeywordService) List(ids []int, opts ...FuncOption) ([]*Keyword, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var key []*Keyword

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ks.client.get(ks.end, &key, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Keywords with IDs %v", ids)
	}

	return key, nil
}

// Index returns an index of Keywords based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Keywords can
// be found using the provided options, an error is returned.
func (ks *KeywordService) Index(opts ...FuncOption) ([]*Keyword, error) {
	var key []*Keyword

	err := ks.client.get(ks.end, &key, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Keywords")
	}

	return key, nil
}

// Count returns the number of Keywords available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Keywords to count.
func (ks *KeywordService) Count(opts ...FuncOption) (int, error) {
	ct, err := ks.client.getCount(ks.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Keywords")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Keyword object.
func (ks *KeywordService) Fields() ([]string, error) {
	f, err := ks.client.getFields(ks.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Keyword fields")
	}

	return f, nil
}
