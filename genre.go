package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Genre -add-tags json -w

// Genre represents the genre of a particular video game.
// For more information visit: https://api-docs.igdb.com/#genre
type Genre struct {
	CreatedAt int    `json:"created_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	UpdatedAt int    `json:"updated_at"`
	URL       string `json:"url"`
}

// GenreService handles all the API calls for the IGDB Genre endpoint.
type GenreService service

// Get returns a single Genre identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Genres, an error is returned.
func (gs *GenreService) Get(id int, opts ...Option) (*Genre, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var gen []*Genre

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := gs.client.get(gs.end, &gen, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Genre with ID %v", id)
	}

	return gen[0], nil
}

// List returns a list of Genres identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Genre is ignored. If none of the IDs
// match a Genre, an error is returned.
func (gs *GenreService) List(ids []int, opts ...Option) ([]*Genre, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var gen []*Genre

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := gs.client.get(gs.end, &gen, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Genres with IDs %v", ids)
	}

	return gen, nil
}

// Index returns an index of Genres based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Genres can
// be found using the provided options, an error is returned.
func (gs *GenreService) Index(opts ...Option) ([]*Genre, error) {
	var gen []*Genre

	err := gs.client.get(gs.end, &gen, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Genres")
	}

	return gen, nil
}

// Count returns the number of Genres available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Genres to count.
func (gs *GenreService) Count(opts ...Option) (int, error) {
	ct, err := gs.client.getCount(gs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Genres")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Genre object.
func (gs *GenreService) Fields() ([]string, error) {
	f, err := gs.client.getFields(gs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Genre fields")
	}

	return f, nil
}
