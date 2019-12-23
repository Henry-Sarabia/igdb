package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Title -add-tags json -w

// Title represents a particular job title in the game industry.
// For more information visit: https://api-docs.igdb.com/#title
type Title struct {
	ID          int    `json:"id,omitempty"`
	CreatedAt   int    `json:"created_at,omitempty"`
	Description string `json:"description,omitempty"`
	Games       []int  `json:"games,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	UpdatedAt   int    `json:"updated_at,omitempty"`
	URL         string `json:"url,omitempty"`
}

// TitleService handles all the API calls for the IGDB Title endpoint.
type TitleService service

// Get returns a single Title identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Titles, an error is returned.
func (ts *TitleService) Get(id int, opts ...Option) (*Title, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var t []*Title

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ts.client.get(ts.end, &t, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Title with ID %v", id)
	}

	return t[0], nil
}

// List returns a list of Titles identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Title is ignored. If none of the IDs
// match a Title, an error is returned.
func (ts *TitleService) List(ids []int, opts ...Option) ([]*Title, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var t []*Title

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ts.client.get(ts.end, &t, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Titles with IDs %v", ids)
	}

	return t, nil
}

// Index returns an index of Titles based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Titles can
// be found using the provided options, an error is returned.
func (ts *TitleService) Index(opts ...Option) ([]*Title, error) {
	var t []*Title

	err := ts.client.get(ts.end, &t, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Titles")
	}

	return t, nil
}

// Count returns the number of Titles available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Titles to count.
func (ts *TitleService) Count(opts ...Option) (int, error) {
	ct, err := ts.client.getCount(ts.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Titles")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Title object.
func (ts *TitleService) Fields() ([]string, error) {
	f, err := ts.client.getFields(ts.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Title fields")
	}

	return f, nil
}
