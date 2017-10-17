package igdb

// Page is
type Page struct {
	ID              int         `json:"id"`
	Slug            string      `json:"slug"`
	URL             URL         `json:"url"`
	CreatedAt       int         `json:"created_at"` // Unix time in milliseconds
	UpdatedAt       int         `json:"updated_at"` // Unix time in milliseconds
	Name            string      `json:"name"`
	Content         string      `json:"content"`
	Category        int         `json:"category"`
	Subcategory     int         `json:"sub_category"`
	Country         CountryCode `json:"country"`
	Color           int         `json:"color"`
	Feed            int         `json:"feed"`
	User            int         `json:"user"`
	Game            int         `json:"game"`
	Company         int         `json:"company"`
	Description     string      `json:"description"`
	PageFollowCount int         `json:"page_follows_count"`
	Logo            Image       `json:"logo"`
	Background      Image       `json:"background"`
	Facebook        string      `json:"facebook"`
	Twitter         string      `json:"twitter"`
	Twitch          string      `json:"twitch"`
	Instagram       string      `json:"instagram"`
	Youtube         string      `json:"youtube"`
	Steam           string      `json:"steam"`
	Linkedin        string      `json:"linkedin"`
	Pinterest       string      `json:"pinterest"`
	Soundcloud      string      `json:"soundcloud"`
	GooglePlus      string      `json:"google_plus"`
	Reddit          string      `json:"reddit"`
	Battlenet       string      `json:"battlenet"`
	Origin          string      `json:"origin"`
	Uplay           string      `json:"uplay"`
	Discord         string      `json:"discord"`
}

// GetPage gets IGDB information for a page identified by its unique IGDB ID.
func (c *Client) GetPage(id int, opts ...optionFunc) (*Page, error) {
	url, err := c.singleURL(PageEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []Page

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPages gets IGDB information for a list of pages identified by their
// unique IGDB IDs.
func (c *Client) GetPages(ids []int, opts ...optionFunc) ([]*Page, error) {
	url, err := c.multiURL(PageEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Page

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPages searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchPages(qry string, opts ...optionFunc) ([]*Page, error) {
	url, err := c.searchURL(PageEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Page

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
