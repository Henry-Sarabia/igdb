package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

// CharacterMugshotService handles all the API calls for the IGDB CharacterMugshot endpoint.
type CharacterMugshotService service

// CharacterMugshot represents an image depicting a game character.
// For more information visit: https://api-docs.igdb.com/#character-mug-shot
type CharacterMugshot struct {
	Image
	ID int `json:"id"`
}

// Get returns a single CharacterMugshot identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any CharacterMugshots, an error is returned.
func (cs *CharacterMugshotService) Get(id int, opts ...Option) (*CharacterMugshot, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var mug []*CharacterMugshot

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.get(cs.end, &mug, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CharacterMugshot with ID %v", id)
	}

	return mug[0], nil
}

// List returns a list of CharacterMugshots identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a CharacterMugshot is ignored. If none of the IDs
// match a CharacterMugshot, an error is returned.
func (cs *CharacterMugshotService) List(ids []int, opts ...Option) ([]*CharacterMugshot, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var mug []*CharacterMugshot

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.get(cs.end, &mug, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CharacterMugshots with IDs %v", ids)
	}

	return mug, nil
}

// Index returns an index of CharacterMugshots based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CharacterMugshots can
// be found using the provided options, an error is returned.
func (cs *CharacterMugshotService) Index(opts ...Option) ([]*CharacterMugshot, error) {
	var mug []*CharacterMugshot

	err := cs.client.get(cs.end, &mug, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CharacterMugshots")
	}

	return mug, nil
}

// Count returns the number of CharacterMugshots available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CharacterMugshots to count.
func (cs *CharacterMugshotService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CharacterMugshots")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CharacterMugshot object.
func (cs *CharacterMugshotService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CharacterMugshot fields")
	}

	return f, nil
}
