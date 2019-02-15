package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct GameMode -add-tags json -w

// GameMode represents a video game mode such as single or multi player.
// For more information visit: https://api-docs.igdb.com/#game-mode
type GameMode struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}

// GameModeService handles all the API calls for the IGDB GameMode endpoint.
type GameModeService service

// Get returns a single GameMode identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameModes, an error is returned.
func (gs *GameModeService) Get(id int, opts ...Option) (*GameMode, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var mode []*GameMode

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &mode, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameMode with ID %v", id)
	}

	return mode[0], nil
}

// List returns a list of GameModes identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameMode is ignored. If none of the IDs
// match a GameMode, an error is returned.
func (gs *GameModeService) List(ids []int, opts ...Option) ([]*GameMode, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var mode []*GameMode

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := gs.client.get(gs.end, &mode, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameModes with IDs %v", ids)
	}

	return mode, nil
}

// Index returns an index of GameModes based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameModes can
// be found using the provided options, an error is returned.
func (gs *GameModeService) Index(opts ...Option) ([]*GameMode, error) {
	var mode []*GameMode

	err := gs.client.get(gs.end, &mode, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameModes")
	}

	return mode, nil
}

// Count returns the number of GameModes available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameModes to count.
func (gs *GameModeService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameModes")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameMode object.
func (gs *GameModeService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameMode fields")
	}

	return f, nil
}
