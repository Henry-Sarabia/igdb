package igdb

// CreditService handles all the API
// calls for the IGDB Credit endpoint.
type CreditService service

// Credit contains information on an IGDB entry
// for an employee responsible for working on
// a particular video game.
type Credit struct {
	ID                    int            `json:"id"`
	Name                  string         `json:"name"`
	Slug                  string         `json:"slug"`
	URL                   URL            `json:"url"`
	CreatedAt             int            `json:"created_at"` // Unix time in milliseconds
	UpdatedAt             int            `json:"updated_at"` // Unix time in milliseconds
	Game                  int            `json:"game"`
	Category              CreditCategory `json:"category"`
	Company               int            `json:"company"`
	Position              int            `json:"position"`
	Person                int            `json:"person"`
	Character             int            `json:"character"`
	Title                 int            `json:"title"`
	Country               CountryCode    `json:"country"`
	CreditedName          string         `json:"credited_name"`
	CharacterCreditedName string         `json:"character_credited_name"`
	PersonTitle           interface{}    `json:"person_title"`
}

// Get returns a single Credit identified by the provided IGDB ID. Provide
// the OptFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Credits, an error is returned.
func (cs *CreditService) Get(id int, opts ...OptionFunc) (*Credit, error) {
	url, err := cs.client.singleURL(CreditEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var cr []Credit

	err = cs.client.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return &cr[0], nil
}

// List returns a list of Credits identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of Credits based solely on the provided
// options. Any ID that does not match a Credit is ignored. If none of the IDs
// match a Credit, an error is returned.
func (cs *CreditService) List(ids []int, opts ...OptionFunc) ([]*Credit, error) {
	url, err := cs.client.multiURL(CreditEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var cr []*Credit

	err = cs.client.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}

// Search returns a list of Credits found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate  the results. If
// no Credits are found using the provided query, an error is returned.
func (cs *CreditService) Search(qry string, opts ...OptionFunc) ([]*Credit, error) {
	url, err := cs.client.searchURL(CreditEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var cr []*Credit

	err = cs.client.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}

// Count returns the number of Credits available in the IGDB.
// Provide the OptFilter functional option if you need to filter
// which Credits to count.
func (cs *CreditService) Count(opts ...OptionFunc) (int, error) {
	ct, err := cs.client.GetEndpointCount(CreditEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Credit object.
func (cs *CreditService) ListFields() ([]string, error) {
	fl, err := cs.client.GetEndpointFieldList(CreditEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
