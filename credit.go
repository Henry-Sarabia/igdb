package igdb

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

// GetCredit returns a single Credit identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetCredit only returning a single Credit object and
// not a list of Credits.
func (c *Client) GetCredit(id int, opts ...OptionFunc) (*Credit, error) {
	url, err := c.singleURL(CreditEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var cr []Credit

	err = c.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return &cr[0], nil
}

// GetCredits returns a list of Credits identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
// Providing an empty list of IDs will instead retrieve an index of Credits based
// solely on the provided options.
func (c *Client) GetCredits(ids []int, opts ...OptionFunc) ([]*Credit, error) {
	url, err := c.multiURL(CreditEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var cr []*Credit

	err = c.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}

// SearchCredits returns a list of Credits found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchCredits(qry string, opts ...OptionFunc) ([]*Credit, error) {
	url, err := c.searchURL(CreditEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var cr []*Credit

	err = c.get(url, &cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}
