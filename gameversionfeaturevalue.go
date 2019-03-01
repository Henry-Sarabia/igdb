package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct GameVersionFeatureValue -add-tags json -w

// GameVersionFeatureValue represents the bool/text value of a particular feature.
// For more information visit: https://api-docs.igdb.com/#game-version-feature-value
type GameVersionFeatureValue struct {
	ID              int                     `json:"id"`
	Game            int                     `json:"game"`
	GameFeature     int                     `json:"game_feature"`
	IncludedFeature VersionFeatureInclusion `json:"included_feature"`
	Note            string                  `json:"note"`
}

//go:generate stringer -type=VersionFeatureInclusion

// VersionFeatureInclusion specifies whether a feature is included or not.
type VersionFeatureInclusion int

// Expected VersionFeatureInclusion enums from the IGDB.
const (
	VersionFeatureNotIncluded VersionFeatureInclusion = iota
	VersionFeatureIncluded
	VersionFeaturePreOrderOnly
)

// GameVersionFeatureValueService handles all the API calls for the IGDB GameVersionFeatureValue endpoint.
type GameVersionFeatureValueService service

// Get returns a single GameVersionFeatureValue identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameVersionFeatureValues, an error is returned.
func (gs *GameVersionFeatureValueService) Get(id int, opts ...Option) (*GameVersionFeatureValue, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var val []*GameVersionFeatureValue

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &val, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVersionFeatureValue with ID %v", id)
	}

	return val[0], nil
}

// List returns a list of GameVersionFeatureValues identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameVersionFeatureValue is ignored. If none of the IDs
// match a GameVersionFeatureValue, an error is returned.
func (gs *GameVersionFeatureValueService) List(ids []int, opts ...Option) ([]*GameVersionFeatureValue, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var val []*GameVersionFeatureValue

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := gs.client.get(gs.end, &val, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVersionFeatureValues with IDs %v", ids)
	}

	return val, nil
}

// Index returns an index of GameVersionFeatureValues based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameVersionFeatureValues can
// be found using the provided options, an error is returned.
func (gs *GameVersionFeatureValueService) Index(opts ...Option) ([]*GameVersionFeatureValue, error) {
	var val []*GameVersionFeatureValue

	err := gs.client.get(gs.end, &val, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameVersionFeatureValues")
	}

	return val, nil
}

// Count returns the number of GameVersionFeatureValues available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameVersionFeatureValues to count.
func (gs *GameVersionFeatureValueService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameVersionFeatureValues")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameVersionFeatureValue object.
func (gs *GameVersionFeatureValueService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameVersionFeatureValue fields")
	}

	return f, nil
}
