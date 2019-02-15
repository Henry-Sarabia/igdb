package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct GameEngine -add-tags json -w

// GameEngine represents a video game engine such as Unreal Engine.
// For more information visit: https://api-docs.igdb.com/#game-engine
type GameEngine struct {
	ID          int    `json:"id"`
	Companies   []int  `json:"companies"`
	CreatedAt   int    `json:"created_at"`
	Description string `json:"description"`
	Logo        int    `json:"logo"`
	Name        string `json:"name"`
	Platforms   []int  `json:"platforms"`
	Slug        string `json:"slug"`
	UpdatedAt   int    `json:"updated_at"`
	URL         string `json:"url"`
}

// GameEngineService handles all the API calls for the IGDB GameEngine endpoint.
type GameEngineService service

// Get returns a single GameEngine identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameEngines, an error is returned.
func (gs *GameEngineService) Get(id int, opts ...Option) (*GameEngine, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var eng []*GameEngine

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &eng, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameEngine with ID %v", id)
	}

	return eng[0], nil
}

// List returns a list of GameEngines identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameEngine is ignored. If none of the IDs
// match a GameEngine, an error is returned.
func (gs *GameEngineService) List(ids []int, opts ...Option) ([]*GameEngine, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var eng []*GameEngine

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := gs.client.get(gs.end, &eng, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameEngines with IDs %v", ids)
	}

	return eng, nil
}

// Index returns an index of GameEngines based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameEngines can
// be found using the provided options, an error is returned.
func (gs *GameEngineService) Index(opts ...Option) ([]*GameEngine, error) {
	var eng []*GameEngine

	err := gs.client.get(gs.end, &eng, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameEngines")
	}

	return eng, nil
}

// Count returns the number of GameEngines available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameEngines to count.
func (gs *GameEngineService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameEngines")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameEngine object.
func (gs *GameEngineService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameEngine fields")
	}

	return f, nil
}
