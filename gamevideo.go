package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct GameVideo -add-tags json -w

// GameVideo represents a video associated with a particular game.
// For more information visit: https://api-docs.igdb.com/#game-video
type GameVideo struct {
	Game    int    `json:"game"`
	Name    string `json:"name"`
	VideoID string `json:"video_id"`
}

// GameVideoService handles all the API calls for the IGDB GameVideo endpoint.
type GameVideoService service

// Get returns a single GameVideo identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any GameVideos, an error is returned.
func (gs *GameVideoService) Get(id int, opts ...Option) (*GameVideo, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var vid []*GameVideo

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &vid, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVideo with ID %v", id)
	}

	return vid[0], nil
}

// List returns a list of GameVideos identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a GameVideo is ignored. If none of the IDs
// match a GameVideo, an error is returned.
func (gs *GameVideoService) List(ids []int, opts ...Option) ([]*GameVideo, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var vid []*GameVideo

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := gs.client.get(gs.end, &vid, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get GameVideos with IDs %v", ids)
	}

	return vid, nil
}

// Index returns an index of GameVideos based solely on the provided functional
// options used to sort, filter, and paginate the results. If no GameVideos can
// be found using the provided options, an error is returned.
func (gs *GameVideoService) Index(opts ...Option) ([]*GameVideo, error) {
	var vid []*GameVideo

	err := gs.client.get(gs.end, &vid, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of GameVideos")
	}

	return vid, nil
}

// Count returns the number of GameVideos available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which GameVideos to count.
func (gs *GameVideoService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count GameVideos")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB GameVideo object.
func (gs *GameVideoService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get GameVideo fields")
	}

	return f, nil
}
