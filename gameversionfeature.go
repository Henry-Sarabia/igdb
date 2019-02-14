package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct GameVersionFeature -add-tags json -w

// GameVersionFeature represents features and descriptions of what makes
// each version/edition different from their main game.
// For more information visit: https://api-docs.igdb.com/#game-version-feature
type GameVersionFeature struct {
	ID          int                    `json:"id"`
	Category    VersionFeatureCategory `json:"category"`
	Description string                 `json:"description"`
	Position    int                    `json:"position"`
	Title       string                 `json:"title"`
	Values      []int                  `json:"values"`
}

//go:generate stringer -type=VersionFeatureCategory

type VersionFeatureCategory int

const (
	VersionFeatureBoolean VersionFeatureCategory = iota
	VersionFeatureDescription
)

// GameVersionFeatureService handles all the API calls for the IGDB GameVersionFeature endpoint.
type GameVersionFeatureService service

// Get returns a single GameVersionFeature identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameVersionFeatures, an error is returned.
func (gs *GameVersionFeatureService) Get(id int, opts ...FuncOption) (*GameVersionFeature, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ft []*GameVersionFeature

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &ft, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVersionFeature with ID %v", id)
	}

	return ft[0], nil
}

// List returns a list of GameVersionFeatures identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameVersionFeature is ignored. If none of the IDs
// match a GameVersionFeature, an error is returned.
func (gs *GameVersionFeatureService) List(ids []int, opts ...FuncOption) ([]*GameVersionFeature, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ft []*GameVersionFeature

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := gs.client.get(gs.end, &ft, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVersionFeatures with IDs %v", ids)
	}

	return ft, nil
}

// Index returns an index of GameVersionFeatures based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameVersionFeatures can
// be found using the provided options, an error is returned.
func (gs *GameVersionFeatureService) Index(opts ...FuncOption) ([]*GameVersionFeature, error) {
	var ft []*GameVersionFeature

	err := gs.client.get(gs.end, &ft, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameVersionFeatures")
	}

	return ft, nil
}

// Count returns the number of GameVersionFeatures available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameVersionFeatures to count.
func (gs *GameVersionFeatureService) Count(opts ...FuncOption) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameVersionFeatures")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameVersionFeature object.
func (gs *GameVersionFeatureService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameVersionFeature fields")
	}

	return f, nil
}
