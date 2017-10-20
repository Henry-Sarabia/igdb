package igdb

// Person contains information on an IGDB
// entry for a particular individual who
// works in the video game industry.
type Person struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	URL         URL         `json:"url"`
	CreatedAt   int         `json:"created_at"` // Unix time in milliseconds
	UpdatedAt   int         `json:"updated_at"` // Unix time in milliseconds
	DOB         int         `json:"dob"`
	Gender      GenderCode  `json:"gender"`
	Country     CountryCode `json:"country"`
	Mugshot     Image       `json:"mug_shot"`
	Bio         string      `json:"bio"`
	Description string      `json:"description"`
	Parent      int         `json:"parent"`
	Homepage    string      `json:"homepage"`
	Twitter     string      `json:"twitter"`
	LinkedIn    string      `json:"linkedin"`
	GooglePlus  string      `json:"google_plus"`
	Facebook    string      `json:"facebook"`
	Instagram   string      `json:"instagram"`
	Tumblr      string      `json:"tumblr"`
	Soundcloud  string      `json:"soundcloud"`
	Pinterest   string      `json:"pinterest"`
	Youtube     string      `json:"youtube"`
	Nicknames   []string    `json:"nicknames"`
	LovesCount  int         `json:"loves_count"`
	Games       []int       `json:"games"`
	Characters  []int       `json:"characters"`
	VoiceActed  []int       `json:"voice_acted"`
}

// GetPerson gets IGDB information for a Person identified by its unique
// IGDB ID. GetPerson returns a single Person identified by the provided
// IGDB ID. Functional options may be provided but sorting and pagination
// will not have an effect due to GetPerson only returning a single Person
// object and not a list of Persons.
func (c *Client) GetPerson(id int, opts ...OptionFunc) (*Person, error) {
	url, err := c.singleURL(PersonEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []Person

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPersons returns a list of Persons identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the
// results. Providing an empty list of IDs will instead retrieve an index of
// Persons based solely on the provided options.
func (c *Client) GetPersons(ids []int, opts ...OptionFunc) ([]*Person, error) {
	url, err := c.multiURL(PersonEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Person

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPersons returns a list of Persons found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchPersons(qry string, opts ...OptionFunc) ([]*Person, error) {
	url, err := c.searchURL(PersonEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Person

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
