package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PulseGroup -add-tags json -w

// PulseGroup represents a combined array of news articles about a specific
// game that were published around the same time period.
// For more information visit: https://api-docs.igdb.com/#pulse-group
type PulseGroup struct {
	ID          int    `json:"id"`
	CreatedAt   int    `json:"created_at"`
	Game        int    `json:"game"`
	Name        string `json:"name"`
	PublishedAt int    `json:"published_at"`
	Pulses      []int  `json:"pulses"`
	Tags        []Tag  `json:"tags"`
	UpdatedAt   int    `json:"updated_at"`
}

// PulseGroupService handles all the API
// calls for the IGDB PulseGroup endpoint.
type PulseGroupService service

// Get returns a single PulseGroup identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PulseGroups, an error is returned.
func (ps *PulseGroupService) Get(id int, opts ...Option) (*PulseGroup, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var pg []*PulseGroup

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &pg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PulseGroup with ID %v", id)
	}

	return pg[0], nil
}

// List returns a list of PulseGroups identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PulseGroup is ignored. If none of the IDs
// match a PulseGroup, an error is returned.
func (ps *PulseGroupService) List(ids []int, opts ...Option) ([]*PulseGroup, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var pg []*PulseGroup

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &pg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PulseGroups with IDs %v", ids)
	}

	return pg, nil
}

// Index returns an index of PulseGroups based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PulseGroups can
// be found using the provided options, an error is returned.
func (ps *PulseGroupService) Index(opts ...Option) ([]*PulseGroup, error) {
	var pg []*PulseGroup

	err := ps.client.get(ps.end, &pg, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PulseGroups")
	}

	return pg, nil
}

// Count returns the number of PulseGroups available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PulseGroups to count.
func (ps *PulseGroupService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PulseGroups")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PulseGroup object.
func (ps *PulseGroupService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PulseGroup fields")
	}

	return f, nil
}
