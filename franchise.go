package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Franchise -add-tags json -w

// Franchise is a list of video game franchises such as Star Wars.
// For more information visit: https://api-docs.igdb.com/#franchise
type Franchise struct {
	ID        int    `json:"id"`
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	Url       string `json:"url"`
}

// FranchiseService handles all the API calls for the IGDB Franchise endpoint.
type FranchiseService service

// Get returns a single Franchise identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Franchises, an error is returned.
func (fs *FranchiseService) Get(id int, opts ...FuncOption) (*Franchise, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var fr []*Franchise

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := fs.client.get(fs.end, &fr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Franchise with ID %v", id)
	}

	return fr[0], nil
}

// List returns a list of Franchises identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Franchise is ignored. If none of the IDs
// match a Franchise, an error is returned.
func (fs *FranchiseService) List(ids []int, opts ...FuncOption) ([]*Franchise, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var fr []*Franchise

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := fs.client.get(fs.end, &fr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Franchises with IDs %v", ids)
	}

	return fr, nil
}

// Index returns an index of Franchises based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Franchises can
// be found using the provided options, an error is returned.
func (fs *FranchiseService) Index(opts ...FuncOption) ([]*Franchise, error) {
	var fr []*Franchise

	err := fs.client.get(fs.end, &fr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Franchises")
	}

	return fr, nil
}

// Count returns the number of Franchises available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Franchises to count.
func (fs *FranchiseService) Count(opts ...FuncOption) (int, error) {
	ct, err := fs.client.getCount(fs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Franchises")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Franchise object.
func (fs *FranchiseService) Fields() ([]string, error) {
	f, err := fs.client.getFields(fs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Franchise fields")
	}

	return f, nil
}
