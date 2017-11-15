package igdb

// CharacterService handles all the API calls
// for the IGDB Characters endpoint.
type CharacterService service

// Character contains information on an IGDB
// entry for a particular video game character.
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

// Get returns a single Character identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetCharacter only returning a single Character object and
// not a list of Characters.
func (cs *CharacterService) Get(id int, opts ...OptionFunc) (*Character, error) {
	url, err := cs.client.singleURL(CharacterEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var ch []Character

	err = cs.client.get(url, &ch)
	if err != nil {
		return nil, err
	}

	if len(ch) == 0 {
		return nil, nil
	}

	return &ch[0], nil
}

// MultiGet returns a list of Characters identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
func (cs *CharacterService) MultiGet(ids []int, opts ...OptionFunc) ([]*Character, error) {
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

// Search returns a list of Characters found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
// Providing an empty query will instead retrieve an index of Characters based solely on
// the provided options.
func (cs *CharacterService) Search(qry string, opts ...OptionFunc) ([]*Character, error) {
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
func (cs *CharacterService) Count() (int, error) {
	c, err := cs.client.GetEndpointCount(CharacterEndpoint)
	if err != nil {
		return 0, err
	}

	return c, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Character object.
func (cs *CharacterService) ListFields() ([]string, error) {
	fl, err := cs.client.GetEndpointFieldList(CharacterEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
