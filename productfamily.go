package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct ProductFamily -add-tags json -w

// ProductFamily represents a collection of closely related platforms.
// For more information visit: https://api-docs.igdb.com/#product-family
type ProductFamily struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ProductFamilyService handles all the API
// calls for the IGDB ProductFamily endpoint.
type ProductFamilyService service

// Get returns a single ProductFamily identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any ProductFamilies, an error is returned.
func (ps *ProductFamilyService) Get(id int, opts ...Option) (*ProductFamily, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var fam []*ProductFamily

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &fam, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ProductFamily with ID %v", id)
	}

	return fam[0], nil
}

// List returns a list of ProductFamilies identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a ProductFamily is ignored. If none of the IDs
// match a ProductFamily, an error is returned.
func (ps *ProductFamilyService) List(ids []int, opts ...Option) ([]*ProductFamily, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var fam []*ProductFamily

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &fam, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ProductFamilies with IDs %v", ids)
	}

	return fam, nil
}

// Index returns an index of ProductFamilies based solely on the provided functional
// options used to sort, filter, and paginate the results. If no ProductFamilies can
// be found using the provided options, an error is returned.
func (ps *ProductFamilyService) Index(opts ...Option) ([]*ProductFamily, error) {
	var fam []*ProductFamily

	err := ps.client.get(ps.end, &fam, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of ProductFamilies")
	}

	return fam, nil
}

// Count returns the number of ProductFamilies available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ProductFamilies to count.
func (ps *ProductFamilyService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count ProductFamilies")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB ProductFamily object.
func (ps *ProductFamilyService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ProductFamily fields")
	}

	return f, nil
}
