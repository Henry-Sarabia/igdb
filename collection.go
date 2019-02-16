package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Collection -add-tags json -w

// Collection represents a video game series.
// For more information visit: https://api-docs.igdb.com/#collection
type Collection struct {
	ID        int    `json:"id"`
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}

// CollectionService handles all the API calls for the IGDB Collection endpoint.
type CollectionService service

// Get returns a single Collection identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Collections, an error is returned.
func (cs *CollectionService) Get(id int, opts ...Option) (*Collection, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var col []*Collection

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.get(cs.end, &col, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Collection with ID %v", id)
	}

	return col[0], nil
}

// List returns a list of Collections identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Collection is ignored. If none of the IDs
// match a Collection, an error is returned.
func (cs *CollectionService) List(ids []int, opts ...Option) ([]*Collection, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var col []*Collection

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.get(cs.end, &col, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Collections with IDs %v", ids)
	}

	return col, nil
}

// Index returns an index of Collections based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Collections can
// be found using the provided options, an error is returned.
func (cs *CollectionService) Index(opts ...Option) ([]*Collection, error) {
	var col []*Collection

	err := cs.client.get(cs.end, &col, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Collections")
	}

	return col, nil
}

// Search returns a list of Collections found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Collections are found using the provided query, an error is returned.
func (cs *CollectionService) Search(qry string, opts ...Option) ([]*Collection, error) {
	var col []*Collection

	opts = append(opts, setSearch(qry))
	err := cs.client.get(cs.end, &col, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Collection with query %s", qry)
	}

	return col, nil
}

// Count returns the number of Collections available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Collections to count.
func (cs *CollectionService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Collections")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Collection object.
func (cs *CollectionService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Collection fields")
	}

	return f, nil
}
