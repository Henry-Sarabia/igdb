package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// ArtworkService handles all the API calls for the IGDB Artwork endpoint.
type ArtworkService service

// Artwork represents an official piece of artwork.
// Resolution and aspect ratio may vary.
// For more information visit: https://api-docs.igdb.com/#artwork
type Artwork struct {
	Image
	ID   int `json:"id"`
	Game int `json:"game"`
}

// Get returns a single Artwork identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Artworks, an error is returned.
func (as *ArtworkService) Get(id int, opts ...FuncOption) (*Artwork, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var art []*Artwork

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := as.client.get(as.end, &art, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Artwork with ID %v", id)
	}

	return art[0], nil
}

// List returns a list of Artworks identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Artwork is ignored. If none of the IDs
// match a Artwork, an error is returned.
func (as *ArtworkService) List(ids []int, opts ...FuncOption) ([]*Artwork, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var art []*Artwork

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := as.client.get(as.end, &art, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Artworks with IDs %v", ids)
	}

	return art, nil
}

// Index returns an index of Artworks based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Artworks can
// be found using the provided options, an error is returned.
func (as *ArtworkService) Index(opts ...FuncOption) ([]*Artwork, error) {
	var art []*Artwork

	err := as.client.get(as.end, &art, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Artworks")
	}

	return art, nil
}

// Count returns the number of Artworks available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Artworks to count.
func (as *ArtworkService) Count(opts ...FuncOption) (int, error) {
	ct, err := as.client.getCount(as.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Artworks")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Artwork object.
func (as *ArtworkService) Fields() ([]string, error) {
	f, err := as.client.getFields(as.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Artwork fields")
	}

	return f, nil
}
