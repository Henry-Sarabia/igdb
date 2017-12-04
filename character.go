package igdb

// CharacterService handles all the API calls
// for the IGDB Characters endpoint.
type CharacterService service

// Character contains information on an IGDB
// entry for a particular video game character.
//
// For more information, visit: https://igdb.github.io/api/endpoints/character/
type Character struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	URL         URL         `json:"url"`
	CreatedAt   int         `json:"created_at"` // Unix time in milliseconds
	UpdatedAt   int         `json:"updated_at"` // Unix time in milliseconds
	Mugshot     Image       `json:"mug_shot"`
	Gender      GenderCode  `json:"gender"`
	CountryName string      `json:"country_name"`
	AKAs        []string    `json:"akas"`
	Species     SpeciesCode `json:"species"`
	Games       []int       `json:"games"`
	People      []int       `json:"people"`
}

// Get returns a single Character identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to retrieve.
// If the ID does not match any Characters, an error is returned.
func (cs *CharacterService) Get(id int, opts ...FuncOption) (*Character, error) {
	url, err := cs.client.singleURL(CharacterEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var ch []Character

	err = cs.client.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return &ch[0], nil
}

// List returns a list of Characters identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of Characters based solely on the provided
// options. Any ID that does not match a Character is ignored. If none of the IDs
// match a Character, an error is returned.
func (cs *CharacterService) List(ids []int, opts ...FuncOption) ([]*Character, error) {
	url, err := cs.client.multiURL(CharacterEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var ch []*Character

	err = cs.client.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

// Search returns a list of Characters found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Characters are found using the provided query, an error is returned.
func (cs *CharacterService) Search(qry string, opts ...FuncOption) ([]*Character, error) {
	url, err := cs.client.searchURL(CharacterEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var ch []*Character

	err = cs.client.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

// Count returns the number of Characters available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Characters to count.
func (cs *CharacterService) Count(opts ...FuncOption) (int, error) {
	ct, err := cs.client.getEndpointCount(CharacterEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Character object.
func (cs *CharacterService) ListFields() ([]string, error) {
	fl, err := cs.client.getEndpointFieldList(CharacterEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
