package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PlatformFamily -add-tags json -w

// PlatformFamily represents a collection of closely related platforms.
// For more information visit: https://api-docs.igdb.com/#platform-family
type PlatformFamily struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// PlatformFamilyService handles all the API
// calls for the IGDB PlatformFamily endpoint.
type PlatformFamilyService service

// Get returns a single PlatformFamily identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformFamilies, an error is returned.
func (ps *PlatformFamilyService) Get(id int, opts ...Option) (*PlatformFamily, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var fam []*PlatformFamily

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &fam, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformFamily with ID %v", id)
	}

	return fam[0], nil
}

// List returns a list of PlatformFamilies identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformFamily is ignored. If none of the IDs
// match a PlatformFamily, an error is returned.
func (ps *PlatformFamilyService) List(ids []int, opts ...Option) ([]*PlatformFamily, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var fam []*PlatformFamily

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &fam, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformFamilies with IDs %v", ids)
	}

	return fam, nil
}

// Index returns an index of PlatformFamilies based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformFamilies can
// be found using the provided options, an error is returned.
func (ps *PlatformFamilyService) Index(opts ...Option) ([]*PlatformFamily, error) {
	var fam []*PlatformFamily

	err := ps.client.get(ps.end, &fam, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformFamilies")
	}

	return fam, nil
}

// Count returns the number of PlatformFamilies available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformFamilies to count.
func (ps *PlatformFamilyService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformFamilies")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlatformFamily object.
func (ps *PlatformFamilyService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlatformFamily fields")
	}

	return f, nil
}
