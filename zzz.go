// +build ignore

package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Zype -add-tags json -w

// ZypeService handles all the API calls for the IGDB Zype endpoint.
type ZypeService service

// Get returns a single Zype identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Zypes, an error is returned.
func (zs *ZypeService) Get(id int, opts ...Option) (*Zype, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var z []*Zype

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := zs.client.get(zs.end, &z, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Zype with ID %v", id)
	}

	return z[0], nil
}

// List returns a list of Zypes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Zype is ignored. If none of the IDs
// match a Zype, an error is returned.
func (zs *ZypeService) List(ids []int, opts ...Option) ([]*Zype, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var z []*Zype

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := zs.client.get(zs.end, &z, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Zypes with IDs %v", ids)
	}

	return z, nil
}

// Index returns an index of Zypes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Zypes can
// be found using the provided options, an error is returned.
func (zs *ZypeService) Index(opts ...Option) ([]*Zype, error) {
	var z []*Zype

	err := zs.client.get(zs.end, &z, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Zypes")
	}

	return z, nil
}

// Count returns the number of Zypes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Zypes to count.
func (zs *ZypeService) Count(opts ...Option) (int, error) {
	ct, err := zs.client.getCount(zs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Zypes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Zype object.
func (zs *ZypeService) Fields() ([]string, error) {
	f, err := zs.client.getFields(zs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Zype fields")
	}

	return f, nil
}

// Search returns a list of Zypes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Zypes are found using the provided query, an error is returned.
func (zs *ZypeService) Search(qry string, opts ...Option) ([]*Zype, error) {
	var z []*Zype

	opts = append(opts, setSearch(qry))
	err := zs.client.get(zs.end, &z, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Zype with query %s", qry)
	}

	return z, nil
}
