package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PulseURL -add-tags json -w

// PulseURL represents a URL linking to an article.
// For more information visit: https://api-docs.igdb.com/#pulse-url
type PulseURL struct {
	ID      int    `json:"id,omitempty"`
	Trusted bool   `json:"trusted,omitempty"`
	URL     string `json:"url,omitempty"`
}

// PulseURLService handles all the API
// calls for the IGDB PulseURL endpoint.
type PulseURLService service

// Get returns a single PulseURL identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PulseURLs, an error is returned.
func (ps *PulseURLService) Get(id int, opts ...Option) (*PulseURL, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var URL []*PulseURL

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &URL, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PulseURL with ID %v", id)
	}

	return URL[0], nil
}

// List returns a list of PulseURLs identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PulseURL is ignored. If none of the IDs
// match a PulseURL, an error is returned.
func (ps *PulseURLService) List(ids []int, opts ...Option) ([]*PulseURL, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var URL []*PulseURL

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &URL, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PulseURLs with IDs %v", ids)
	}

	return URL, nil
}

// Index returns an index of PulseURLs based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PulseURLs can
// be found using the provided options, an error is returned.
func (ps *PulseURLService) Index(opts ...Option) ([]*PulseURL, error) {
	var URL []*PulseURL

	err := ps.client.get(ps.end, &URL, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PulseURLs")
	}

	return URL, nil
}

// Search returns a list of PulseURLs found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no PulseURLs are found using the provided query, an error is returned.
func (ps *PulseURLService) Search(qry string, opts ...Option) ([]*PulseURL, error) {
	var URL []*PulseURL

	opts = append(opts, setSearch(qry))
	err := ps.client.get(ps.end, &URL, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PulseURL with query %s", qry)
	}

	return URL, nil
}

// Count returns the number of PulseURLs available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PulseURLs to count.
func (ps *PulseURLService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PulseURLs")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PulseURL object.
func (ps *PulseURLService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PulseURL fields")
	}

	return f, nil
}
