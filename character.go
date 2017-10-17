package igdb

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

// GetCharacter returns a single Character identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetCharacter only returning a single Character object and
// not a list of Characters.
func (c *Client) GetCharacter(id int, opts ...OptionFunc) (*Character, error) {
	url, err := c.singleURL(CharacterEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var ch []Character

	err = c.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return &ch[0], nil
}

// GetCharacters returns a list of Characters identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of Characters based
// solely on the provided options.
func (c *Client) GetCharacters(ids []int, opts ...OptionFunc) ([]*Character, error) {
	url, err := c.multiURL(CharacterEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var ch []*Character

	err = c.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

// SearchCharacters returns a list of Characters found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchCharacters(qry string, opts ...OptionFunc) ([]*Character, error) {
	url, err := c.searchURL(CharacterEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var ch []*Character

	err = c.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
