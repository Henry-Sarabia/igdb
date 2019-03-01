package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct List -add-tags json -w

// List represents a user-created list of games.
// For more information visit: https://api-docs.igdb.com/#list
type List struct {
	ID           int    `json:"id"`
	CreatedAt    int    `json:"created_at"`
	Description  string `json:"description"`
	EntriesCount int    `json:"entries_count"`
	ListEntries  []int  `json:"list_entries"`
	ListTags     []int  `json:"list_tags"`
	ListedGames  []int  `json:"listed_games"`
	Name         string `json:"name"`
	Numbering    bool   `json:"numbering"`
	Private      bool   `json:"private"`
	SimilarLists []int  `json:"similar_lists"`
	Slug         string `json:"slug"`
	UpdatedAt    int    `json:"updated_at"`
	URL          string `json:"url"`
	User         int    `json:"user"`
}

// ListService handles all the API calls for the IGDB List endpoint.
type ListService service

// Get returns a single List identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Lists, an error is returned.
func (ls *ListService) Get(id int, opts ...Option) (*List, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var l []*List

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ls.client.get(ls.end, &l, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get List with ID %v", id)
	}

	return l[0], nil
}

// List returns a list of Lists identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a List is ignored. If none of the IDs
// match a List, an error is returned.
func (ls *ListService) List(ids []int, opts ...Option) ([]*List, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var l []*List

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ls.client.get(ls.end, &l, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Lists with IDs %v", ids)
	}

	return l, nil
}

// Index returns an index of Lists based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Lists can
// be found using the provided options, an error is returned.
func (ls *ListService) Index(opts ...Option) ([]*List, error) {
	var l []*List

	err := ls.client.get(ls.end, &l, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Lists")
	}

	return l, nil
}

// Count returns the number of Lists available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Lists to count.
func (ls *ListService) Count(opts ...Option) (int, error) {
	ct, err := ls.client.getCount(ls.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Lists")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB List object.
func (ls *ListService) Fields() ([]string, error) {
	f, err := ls.client.getFields(ls.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get List fields")
	}

	return f, nil
}
