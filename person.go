package igdb

// PersonService handles all the API
// calls for the IGDB Person endpoint.
type PersonService service

// Person contains information on an IGDB entry for a particular
// individual who works in the video game industry.
//
// For more information, visit: https://igdb.github.io/api/endpoints/person/
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

// Get returns a single Person identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any People, an error is returned.
func (ps *PersonService) Get(id int, opts ...FuncOption) (*Person, error) {
	url, err := ps.client.singleURL(PersonEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var p []Person

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// List returns a list of People identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of People based solely on the provided
// options. Any ID that does not match a Person is ignored. If none of the IDs
// match a Person, an error is returned.
func (ps *PersonService) List(ids []int, opts ...FuncOption) ([]*Person, error) {
	url, err := ps.client.multiURL(PersonEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var p []*Person

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Search returns a list of People found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no People are found using the provided query, an error is returned.
func (ps *PersonService) Search(qry string, opts ...FuncOption) ([]*Person, error) {
	url, err := ps.client.searchURL(PersonEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var p []*Person

	err = ps.client.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Count returns the number of People available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which People to count.
func (ps *PersonService) Count(opts ...FuncOption) (int, error) {
	ct, err := ps.client.getEndpointCount(PersonEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Person object.
func (ps *PersonService) ListFields() ([]string, error) {
	fl, err := ps.client.getEndpointFieldList(PersonEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
