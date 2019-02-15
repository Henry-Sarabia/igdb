package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PlayerPerspective -add-tags json -w

// PlayerPerspective describes the view or perspective of the player in a video game.
// For more information visit: https://api-docs.igdb.com/#player-perspective
type PlayerPerspective struct {
	ID        int    `json:"id"`
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}

// PlayerPerspectiveService handles all the API calls for the IGDB PlayerPerspective endpoint.
type PlayerPerspectiveService service

// Get returns a single PlayerPerspective identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlayerPerspectives, an error is returned.
func (ps *PlayerPerspectiveService) Get(id int, opts ...Option) (*PlayerPerspective, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var pp []*PlayerPerspective

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &pp, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlayerPerspective with ID %v", id)
	}

	return pp[0], nil
}

// List returns a list of PlayerPerspectives identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlayerPerspective is ignored. If none of the IDs
// match a PlayerPerspective, an error is returned.
func (ps *PlayerPerspectiveService) List(ids []int, opts ...Option) ([]*PlayerPerspective, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var pp []*PlayerPerspective

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &pp, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlayerPerspectives with IDs %v", ids)
	}

	return pp, nil
}

// Index returns an index of PlayerPerspectives based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlayerPerspectives can
// be found using the provided options, an error is returned.
func (ps *PlayerPerspectiveService) Index(opts ...Option) ([]*PlayerPerspective, error) {
	var pp []*PlayerPerspective

	err := ps.client.get(ps.end, &pp, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlayerPerspectives")
	}

	return pp, nil
}

// Count returns the number of PlayerPerspectives available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlayerPerspectives to count.
func (ps *PlayerPerspectiveService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlayerPerspectives")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlayerPerspective object.
func (ps *PlayerPerspectiveService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlayerPerspective fields")
	}

	return f, nil
}
