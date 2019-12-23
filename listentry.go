package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct ListEntry -add-tags json -w

// ListEntry represents an entry in a user-created list of games.
// For more information visit: https://api-docs.igdb.com/#list-entry
type ListEntry struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Game        int    `json:"game,omitempty"`
	List        int    `json:"list,omitempty"`
	Platform    int    `json:"platform,omitempty"`
	Position    int    `json:"position,omitempty"`
	Private     bool   `json:"private,omitempty"`
	User        int    `json:"user,omitempty"`
}

// ListEntryService handles all the API calls for the IGDB ListEntry endpoint.
type ListEntryService service

// Get returns a single ListEntry identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any ListEntrys, an error is returned.
func (ls *ListEntryService) Get(id int, opts ...Option) (*ListEntry, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var le []*ListEntry

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ls.client.get(ls.end, &le, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ListEntry with ID %v", id)
	}

	return le[0], nil
}

// List returns a list of ListEntrys identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a ListEntry is ignored. If none of the IDs
// match a ListEntry, an error is returned.
func (ls *ListEntryService) List(ids []int, opts ...Option) ([]*ListEntry, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var le []*ListEntry

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ls.client.get(ls.end, &le, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ListEntrys with IDs %v", ids)
	}

	return le, nil
}

// Index returns an index of ListEntrys based solely on the provided functional
// options used to sort, filter, and paginate the results. If no ListEntrys can
// be found using the provided options, an error is returned.
func (ls *ListEntryService) Index(opts ...Option) ([]*ListEntry, error) {
	var le []*ListEntry

	err := ls.client.get(ls.end, &le, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of ListEntrys")
	}

	return le, nil
}

// Count returns the number of ListEntrys available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ListEntrys to count.
func (ls *ListEntryService) Count(opts ...Option) (int, error) {
	ct, err := ls.client.getCount(ls.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count ListEntrys")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB ListEntry object.
func (ls *ListEntryService) Fields() ([]string, error) {
	f, err := ls.client.getFields(ls.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ListEntry fields")
	}

	return f, nil
}
