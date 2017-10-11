package igdb

// CreditCategory corresponds to the IGDB
// enumerated credit type which describes
// the type of entry in an end credits list.
// CreditCategory implements the Stringer
// interface.
type CreditCategory int

// Credit is
type Credit struct {
	ID                    int            `json:"id"`
	Name                  string         `json:"name"`
	Slug                  string         `json:"slug"`
	URL                   URL            `json:"url"`
	CreatedAt             int            `json:"created_at"` //unix epoch
	UpdatedAt             int            `json:"updated_at"` //unix epoch
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

// GetCredit gets IGDB information for a credit identified by its unique IGDB ID.
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

// GetCredits gets IGDB information for a list of credits identified by their
// unique IGDB IDs.
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

// SearchCredits searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
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

// CreditCategory implements the Stringer interface
// by matching its code with the IGDBs enumerated type
// Credit Category and returns the category as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/credit-category/
func (c CreditCategory) String() string {
	switch c {
	case 1:
		return "voice_actor"
	case 2:
		return "language"
	case 3:
		return "company_credit"
	case 4:
		return "employee"
	case 5:
		return "misc"
	case 6:
		return "support_company"
	default:
		return "Undefined"
	}
}
