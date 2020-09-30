package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

// Screenshot represents a screenshot of a particular game.
// For more information visit: https://api-docs.igdb.com/#screenshot
type Screenshot struct {
	Image
	ID   int `json:"id"`
	Game int `json:"game"`
}

// ScreenshotService handles all the API calls for the IGDB Screenshot endpoint.
type ScreenshotService service

// Get returns a single Screenshot identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Screenshots, an error is returned.
func (ss *ScreenshotService) Get(id int, opts ...Option) (*Screenshot, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var shot []*Screenshot

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ss.client.post(ss.end, &shot, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Screenshot with ID %v", id)
	}

	return shot[0], nil
}

// List returns a list of Screenshots identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Screenshot is ignored. If none of the IDs
// match a Screenshot, an error is returned.
func (ss *ScreenshotService) List(ids []int, opts ...Option) ([]*Screenshot, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var shot []*Screenshot

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ss.client.post(ss.end, &shot, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Screenshots with IDs %v", ids)
	}

	return shot, nil
}

// Index returns an index of Screenshots based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Screenshots can
// be found using the provided options, an error is returned.
func (ss *ScreenshotService) Index(opts ...Option) ([]*Screenshot, error) {
	var shot []*Screenshot

	err := ss.client.post(ss.end, &shot, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Screenshots")
	}

	return shot, nil
}

// Count returns the number of Screenshots available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Screenshots to count.
func (ss *ScreenshotService) Count(opts ...Option) (int, error) {
	ct, err := ss.client.getCount(ss.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Screenshots")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Screenshot object.
func (ss *ScreenshotService) Fields() ([]string, error) {
	f, err := ss.client.getFields(ss.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Screenshot fields")
	}

	return f, nil
}
