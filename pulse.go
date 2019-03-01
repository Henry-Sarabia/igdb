package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Pulse -add-tags json -w

// Pulse represents a single news article.
// For more information visit: https://api-docs.igdb.com/#pulse
type Pulse struct {
	ID          int      `json:"id"`
	Author      string   `json:"author"`
	CreatedAt   int      `json:"created_at"`
	Image       string   `json:"image"`
	PublishedAt int      `json:"published_at"`
	PulseSource int      `json:"pulse_source"`
	Summary     string   `json:"summary"`
	Tags        []Tag    `json:"tags"`
	Title       string   `json:"title"`
	UID         string   `json:"uid"`
	UpdatedAt   int      `json:"updated_at"`
	Videos      []string `json:"videos"`
	Website     int      `json:"website"`
}

// PulseService handles all the API
// calls for the IGDB Pulse endpoint.
type PulseService service

// Get returns a single Pulse identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Pulses, an error is returned.
func (ps *PulseService) Get(id int, opts ...Option) (*Pulse, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var p []*Pulse

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &p, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Pulse with ID %v", id)
	}

	return p[0], nil
}

// List returns a list of Pulses identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Pulse is ignored. If none of the IDs
// match a Pulse, an error is returned.
func (ps *PulseService) List(ids []int, opts ...Option) ([]*Pulse, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var p []*Pulse

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &p, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Pulses with IDs %v", ids)
	}

	return p, nil
}

// Index returns an index of Pulses based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Pulses can
// be found using the provided options, an error is returned.
func (ps *PulseService) Index(opts ...Option) ([]*Pulse, error) {
	var p []*Pulse

	err := ps.client.get(ps.end, &p, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Pulses")
	}

	return p, nil
}

// Count returns the number of Pulses available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Pulses to count.
func (ps *PulseService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Pulses")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Pulse object.
func (ps *PulseService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Pulse fields")
	}

	return f, nil
}
