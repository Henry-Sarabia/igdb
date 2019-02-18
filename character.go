package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Character -add-tags json -w

// Character represents a video game character.
// For more information visit: https://api-docs.igdb.com/#character
type Character struct {
	ID          int              `json:"ID"`
	AKAS        []string         `json:"akas"`
	CountryName string           `json:"country_name"`
	CreatedAt   int              `json:"created_at"`
	Description string           `json:"description"`
	Games       []int            `json:"games"`
	Gender      CharacterGender  `json:"gender"`
	MugShot     int              `json:"mug_shot"`
	Name        string           `json:"name"`
	People      []int            `json:"people"`
	Slug        string           `json:"slug"`
	Species     CharacterSpecies `json:"species"`
	UpdatedAt   int              `json:"updated_at"`
	URL         string           `json:"url"`
}

type CharacterGender int

//go:generate stringer -type=CharacterGender,CharacterSpecies

const (
	GenderMale CharacterGender = iota + 1
	GenderFemale
	GenderOther
)

type CharacterSpecies int

const (
	SpeciesHuman CharacterSpecies = iota + 1
	SpeciesAlien
	SpeciesAnimal
	SpeciesAndroid
	SpeciesUnknown
)

// CharacterService handles all the API calls for the IGDB Character endpoint.
type CharacterService service

// Get returns a single Character identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Characters, an error is returned.
func (cs *CharacterService) Get(id int, opts ...Option) (*Character, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ch []*Character

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.get(cs.end, &ch, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Character with ID %v", id)
	}

	return ch[0], nil
}

// List returns a list of Characters identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Character is ignored. If none of the IDs
// match a Character, an error is returned.
func (cs *CharacterService) List(ids []int, opts ...Option) ([]*Character, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ch []*Character

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.get(cs.end, &ch, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Characters with IDs %v", ids)
	}

	return ch, nil
}

// Index returns an index of Characters based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Characters can
// be found using the provided options, an error is returned.
func (cs *CharacterService) Index(opts ...Option) ([]*Character, error) {
	var ch []*Character

	err := cs.client.get(cs.end, &ch, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Characters")
	}

	return ch, nil
}

// Search returns a list of Characters found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Characters are found using the provided query, an error is returned.
func (cs *CharacterService) Search(qry string, opts ...Option) ([]*Character, error) {
	var ch []*Character

	opts = append(opts, setSearch(qry))
	err := cs.client.get(cs.end, &ch, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Character with query %s", qry)
	}

	return ch, nil
}

// Count returns the number of Characters available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Characters to count.
func (cs *CharacterService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Characters")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Character object.
func (cs *CharacterService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Character fields")
	}

	return f, nil
}
