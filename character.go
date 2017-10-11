package igdb

// GenderCode corresponds to the IGDB
// enumerated gender type which describes
// a character's gender. GenderCode
// implements the Stringer interface.
type GenderCode int

// SpeciesCode corresponds to the IGDB
// enumerated species type which describes
// a character's species. SpeciesCode
// implements the Stringer interface.
type SpeciesCode int

// Character is
type Character struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	URL         URL         `json:"url"`
	CreatedAt   int         `json:"created_at"`
	UpdatedAt   int         `json:"updated_at"`
	Mugshot     Image       `json:"mug_shot"`
	Gender      GenderCode  `json:"gender"`
	CountryName string      `json:"country_name"`
	AKAs        []string    `json:"akas"`
	Species     SpeciesCode `json:"species"`
	Games       []int       `json:"games"`
	People      []int       `json:"people"`
}

// GetCharacter gets IGDB information for a character identified by its unique IGDB ID.
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

// GetCharacters gets IGDB information for a list of characters identified by their
// unique IGDB IDs.
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

// SearchCharacters searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
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

// GenderCode implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type Gender and returns the value as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/gender/
func (g GenderCode) String() string {
	switch g {
	case 0:
		return "Male"
	case 1:
		return "Female"
	case 2:
		return "Unknown"
	default:
		return "Undefined"
	}
}

// String will return the enumerated type as a string
// corresponding to its IGDB code. For more information,
// SpeciesCode implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type Species and returns the value as a string.
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/species/
func (s SpeciesCode) String() string {
	switch s {
	case 1:
		return "Human"
	case 2:
		return "Alien"
	case 3:
		return "Animal"
	case 4:
		return "Android"
	case 5:
		return "Unknown"
	default:
		return "Undefined"
	}
}
