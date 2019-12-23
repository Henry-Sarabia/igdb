package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Theme -add-tags json -w

// Theme represents a particular video game theme.
// For more information visit: https://api-docs.igdb.com/#theme
type Theme struct {
	ID        int    `json:"id,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	Name      string `json:"name,omitempty"`
	Slug      string `json:"slug,omitempty"`
	UpdatedAt int    `json:"updated_at,omitempty"`
	URL       string `json:"url,omitempty"`
}

// ThemeService handles all the API calls for the IGDB Theme endpoint.
type ThemeService service

// Get returns a single Theme identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Themes, an error is returned.
func (ts *ThemeService) Get(id int, opts ...Option) (*Theme, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var th []*Theme

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ts.client.get(ts.end, &th, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Theme with ID %v", id)
	}

	return th[0], nil
}

// List returns a list of Themes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Theme is ignored. If none of the IDs
// match a Theme, an error is returned.
func (ts *ThemeService) List(ids []int, opts ...Option) ([]*Theme, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var th []*Theme

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ts.client.get(ts.end, &th, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Themes with IDs %v", ids)
	}

	return th, nil
}

// Index returns an index of Themes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Themes can
// be found using the provided options, an error is returned.
func (ts *ThemeService) Index(opts ...Option) ([]*Theme, error) {
	var th []*Theme

	err := ts.client.get(ts.end, &th, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Themes")
	}

	return th, nil
}

// Search returns a list of Themes found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Themes are found using the provided query, an error is returned.
func (ts *ThemeService) Search(qry string, opts ...Option) ([]*Theme, error) {
	var th []*Theme

	opts = append(opts, setSearch(qry))
	err := ts.client.get(ts.end, &th, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Theme with query %s", qry)
	}

	return th, nil
}

// Count returns the number of Themes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Themes to count.
func (ts *ThemeService) Count(opts ...Option) (int, error) {
	ct, err := ts.client.getCount(ts.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Themes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Theme object.
func (ts *ThemeService) Fields() ([]string, error) {
	f, err := ts.client.getFields(ts.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Theme fields")
	}

	return f, nil
}
