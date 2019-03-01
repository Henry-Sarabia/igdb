package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PulseSource -add-tags json -w

// PulseSource represents a news article source such as IGN.
// For more information visit: https://api-docs.igdb.com/#pulse-source
type PulseSource struct {
	ID   int    `json:"id"`
	Game int    `json:"game"`
	Name string `json:"name"`
	Page int    `json:"page"`
}

// PulseSourceService handles all the API
// calls for the IGDB PulseSource endpoint.
type PulseSourceService service

// Get returns a single PulseSource identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PulseSources, an error is returned.
func (ps *PulseSourceService) Get(id int, opts ...Option) (*PulseSource, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var src []*PulseSource

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &src, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PulseSource with ID %v", id)
	}

	return src[0], nil
}

// List returns a list of PulseSources identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PulseSource is ignored. If none of the IDs
// match a PulseSource, an error is returned.
func (ps *PulseSourceService) List(ids []int, opts ...Option) ([]*PulseSource, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var src []*PulseSource

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &src, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PulseSources with IDs %v", ids)
	}

	return src, nil
}

// Index returns an index of PulseSources based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PulseSources can
// be found using the provided options, an error is returned.
func (ps *PulseSourceService) Index(opts ...Option) ([]*PulseSource, error) {
	var src []*PulseSource

	err := ps.client.get(ps.end, &src, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PulseSources")
	}

	return src, nil
}

// Count returns the number of PulseSources available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PulseSources to count.
func (ps *PulseSourceService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PulseSources")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PulseSource object.
func (ps *PulseSourceService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PulseSource fields")
	}

	return f, nil
}
