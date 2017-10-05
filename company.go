package igdb

// CountryCode code ISO-3316-1
type CountryCode int

// Company is
type Company struct {
	ID                 int          `json:"ID"`
	Name               string       `json:"name"`
	Slug               string       `json:"slug"`
	URL                URL          `json:"url"`
	CreatedAt          int          `json:"created_at"`
	UpdatedAt          int          `json:"updated_at"`
	Logo               Image        `json:"logo"`
	Description        string       `json:"description"`
	Country            CountryCode  `json:"country"`
	Website            string       `json:"website"`
	StartDate          int          `json:"start_date"` //unix epoch
	StartDateCategory  DateCategory `json:"start_date_category"`
	ChangedID          int          `json:"changed_company_id"`
	ChangeDate         int          `json:"change_date"` //unix epoch
	ChangeDateCategory DateCategory `json:"change_date_category"`
	Twitter            string       `json:"twitter"`
	Published          []int        `json:"published"`
	Developed          []int        `json:"developed"`
	Parent             int          `json:"parent"`
	Facebook           string       `json:"facebook"`
}

// GetCompany gets IGDB information for a company identified by its unique IGDB ID.
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

// GetCompanies gets IGDB information for a list of companies identified by their
// unique IGDB IDs.
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

// SearchCompanies searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
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
