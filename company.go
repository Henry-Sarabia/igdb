package igdb

// Company contains information on an IGDB entry
// for a particular video game company, including
// both publishers and developers.
type Company struct {
	ID                 int          `json:"ID"`
	Name               string       `json:"name"`
	Slug               string       `json:"slug"`
	URL                URL          `json:"url"`
	CreatedAt          int          `json:"created_at"` // Unix time in milliseconds
	UpdatedAt          int          `json:"updated_at"` // Unix time in milliseconds
	Logo               Image        `json:"logo"`
	Description        string       `json:"description"`
	Country            CountryCode  `json:"country"`
	Website            string       `json:"website"`
	StartDate          int          `json:"start_date"` // Unix time in milliseconds
	StartDateCategory  DateCategory `json:"start_date_category"`
	ChangedID          int          `json:"changed_company_id"`
	ChangeDate         int          `json:"change_date"` // Unix time in milliseconds
	ChangeDateCategory DateCategory `json:"change_date_category"`
	Twitter            string       `json:"twitter"`
	Published          []int        `json:"published"`
	Developed          []int        `json:"developed"`
	Parent             int          `json:"parent"`
	Facebook           string       `json:"facebook"`
}

// GetCompany returns a single Company identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will not have
// an effect due to GetCompany only returning a single Company object and
// not a list of Companies.
func (c *Client) GetCompany(id int, opts ...OptionFunc) (*Company, error) {
	url, err := c.singleURL(CompanyEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var com []Company

	err = c.get(url, &com)
	if err != nil {
		return nil, err
	}

	return &com[0], nil
}

// GetCompanies returns a list of Companies identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the results.
func (c *Client) GetCompanies(ids []int, opts ...OptionFunc) ([]*Company, error) {
	url, err := c.multiURL(CompanyEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var com []*Company

	err = c.get(url, &com)
	if err != nil {
		return nil, err
	}

	return com, nil
}

// SearchCompanies returns a list of Companies found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
// Providing an empty query will instead retrieve an index of Companies based solely
// on the provided options.
func (c *Client) SearchCompanies(qry string, opts ...OptionFunc) ([]*Company, error) {
	url, err := c.searchURL(CompanyEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var com []*Company

	err = c.get(url, &com)
	if err != nil {
		return nil, err
	}

	return com, nil
}
