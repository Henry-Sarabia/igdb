package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct GameVersion -add-tags json -w

// GameVersion provides details about game editions and versions.
// For more information visit: https://api-docs.igdb.com/#game-version
type GameVersion struct {
	CreatedAt int    `json:"created_at"`
	Features  []int  `json:"features"`
	Game      int    `json:"game"`
	Games     []int  `json:"games"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}

// GameVersionService handles all the API calls for the IGDB GameVersion endpoint.
type GameVersionService service

// Get returns a single GameVersion identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameVersions, an error is returned.
func (gs *GameVersionService) Get(id int, opts ...FuncOption) (*GameVersion, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ver []*GameVersion

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVersion with ID %v", id)
	}

	return ver[0], nil
}

// List returns a list of GameVersions identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameVersion is ignored. If none of the IDs
// match a GameVersion, an error is returned.
func (gs *GameVersionService) List(ids []int, opts ...FuncOption) ([]*GameVersion, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ver []*GameVersion

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := gs.client.get(gs.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVersions with IDs %v", ids)
	}

	return ver, nil
}

// Index returns an index of GameVersions based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameVersions can
// be found using the provided options, an error is returned.
func (gs *GameVersionService) Index(opts ...FuncOption) ([]*GameVersion, error) {
	var ver []*GameVersion

	err := gs.client.get(gs.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameVersions")
	}

	return ver, nil
}

// Count returns the number of GameVersions available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameVersions to count.
func (gs *GameVersionService) Count(opts ...FuncOption) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameVersions")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameVersion object.
func (gs *GameVersionService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameVersion fields")
	}

	return f, nil
}
