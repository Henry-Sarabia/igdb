package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// GameEngineLogo represents the logo of a particular game engine.
// For more information visit: https://api-docs.igdb.com/#game-engine-logo
type GameEngineLogo struct {
	Image
	ID int `json:"id"`
}

// GameEngineLogoService handles all the API calls for the IGDB GameEngineLogo endpoint.
type GameEngineLogoService service

// Get returns a single GameEngineLogo identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameEngineLogos, an error is returned.
func (gs *GameEngineLogoService) Get(id int, opts ...Option) (*GameEngineLogo, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var logo []*GameEngineLogo

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameEngineLogo with ID %v", id)
	}

	return logo[0], nil
}

// List returns a list of GameEngineLogos identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameEngineLogo is ignored. If none of the IDs
// match a GameEngineLogo, an error is returned.
func (gs *GameEngineLogoService) List(ids []int, opts ...Option) ([]*GameEngineLogo, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var logo []*GameEngineLogo

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := gs.client.get(gs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameEngineLogos with IDs %v", ids)
	}

	return logo, nil
}

// Index returns an index of GameEngineLogos based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameEngineLogos can
// be found using the provided options, an error is returned.
func (gs *GameEngineLogoService) Index(opts ...Option) ([]*GameEngineLogo, error) {
	var logo []*GameEngineLogo

	err := gs.client.get(gs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameEngineLogos")
	}

	return logo, nil
}

// Count returns the number of GameEngineLogos available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameEngineLogos to count.
func (gs *GameEngineLogoService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameEngineLogos")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameEngineLogo object.
func (gs *GameEngineLogoService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameEngineLogo fields")
	}

	return f, nil
}
