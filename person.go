package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Person -add-tags json -w

// Person represents a person in the video game industry.
// For more information visit: https://api-docs.igdb.com/#person
type Person struct {
	ID            int             `json:"id,omitempty"`
	Bio           string          `json:"bio,omitempty"`
	Characters    []int           `json:"characters,omitempty"`
	Country       int             `json:"country,omitempty"`
	CreatedAt     int             `json:"created_at,omitempty"`
	CreditedGames []int           `json:"credited_games,omitempty"`
	Description   string          `json:"description,omitempty"`
	DOB           int             `json:"dob,omitempty"`
	Gender        CharacterGender `json:"gender,omitempty"`
	LovesCount    int             `json:"loves_count,omitempty"`
	MugShot       int             `json:"mug_shot,omitempty"`
	Name          string          `json:"name,omitempty"`
	Nicknames     []string        `json:"nicknames,omitempty"`
	Parent        int             `json:"parent,omitempty"`
	Slug          string          `json:"slug,omitempty"`
	UpdatedAt     int             `json:"updated_at,omitempty"`
	URL           string          `json:"url,omitempty"`
	VoiceActed    []int           `json:"voice_acted,omitempty"`
	Websites      []int           `json:"websites,omitempty"`
}

// PersonService handles all the API calls for the IGDB Person endpoint.
type PersonService service

// Get returns a single Person identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any People, an error is returned.
func (ps *PersonService) Get(id int, opts ...Option) (*Person, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var p []*Person

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &p, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Person with ID %v", id)
	}

	return p[0], nil
}

// List returns a list of People identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Person is ignored. If none of the IDs
// match a Person, an error is returned.
func (ps *PersonService) List(ids []int, opts ...Option) ([]*Person, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var p []*Person

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &p, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get People with IDs %v", ids)
	}

	return p, nil
}

// Index returns an index of People based solely on the provided functional
// options used to sort, filter, and paginate the results. If no People can
// be found using the provided options, an error is returned.
func (ps *PersonService) Index(opts ...Option) ([]*Person, error) {
	var p []*Person

	err := ps.client.get(ps.end, &p, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of People")
	}

	return p, nil
}

// Search returns a list of People found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no People are found using the provided query, an error is returned.
func (ps *PersonService) Search(qry string, opts ...Option) ([]*Person, error) {
	var p []*Person

	opts = append(opts, setSearch(qry))
	err := ps.client.get(ps.end, &p, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Person with query %s", qry)
	}

	return p, nil
}

// Count returns the number of People available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which People to count.
func (ps *PersonService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count People")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Person object.
func (ps *PersonService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Person fields")
	}

	return f, nil
}
