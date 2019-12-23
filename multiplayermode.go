package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct MultiplayerMode -add-tags json -w

// MultiplayerMode contains data about the supported multiplayer types.
// For more information visit: https://api-docs.igdb.com/#multiplayer-mode
type MultiplayerMode struct {
	Campaigncoop      bool `json:"campaigncoop,omitempty"`
	Dropin            bool `json:"dropin,omitempty"`
	Lancoop           bool `json:"lancoop,omitempty"`
	Offlinecoop       bool `json:"offlinecoop,omitempty"`
	Offlinecoopmax    int  `json:"offlinecoopmax,omitempty"`
	Offlinemax        int  `json:"offlinemax,omitempty"`
	Onlinecoop        bool `json:"onlinecoop,omitempty"`
	Onlinecoopmax     int  `json:"onlinecoopmax,omitempty"`
	Onlinemax         int  `json:"onlinemax,omitempty"`
	Platform          int  `json:"platform,omitempty"`
	Splitscreen       bool `json:"splitscreen,omitempty"`
	Splitscreenonline bool `json:"splitscreenonline,omitempty"`
}

// MultiplayerModeService handles all the API calls for the IGDB MultiplayerMode endpoint.
type MultiplayerModeService service

// Get returns a single MultiplayerMode identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any MultiplayerModes, an error is returned.
func (ms *MultiplayerModeService) Get(id int, opts ...Option) (*MultiplayerMode, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var mode []*MultiplayerMode

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ms.client.get(ms.end, &mode, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get MultiplayerMode with ID %v", id)
	}

	return mode[0], nil
}

// List returns a list of MultiplayerModes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a MultiplayerMode is ignored. If none of the IDs
// match a MultiplayerMode, an error is returned.
func (ms *MultiplayerModeService) List(ids []int, opts ...Option) ([]*MultiplayerMode, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var mode []*MultiplayerMode

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ms.client.get(ms.end, &mode, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get MultiplayerModes with IDs %v", ids)
	}

	return mode, nil
}

// Index returns an index of MultiplayerModes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no MultiplayerModes can
// be found using the provided options, an error is returned.
func (ms *MultiplayerModeService) Index(opts ...Option) ([]*MultiplayerMode, error) {
	var mode []*MultiplayerMode

	err := ms.client.get(ms.end, &mode, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of MultiplayerModes")
	}

	return mode, nil
}

// Count returns the number of MultiplayerModes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which MultiplayerModes to count.
func (ms *MultiplayerModeService) Count(opts ...Option) (int, error) {
	ct, err := ms.client.getCount(ms.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count MultiplayerModes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB MultiplayerMode object.
func (ms *MultiplayerModeService) Fields() ([]string, error) {
	f, err := ms.client.getFields(ms.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get MultiplayerMode fields")
	}

	return f, nil
}
