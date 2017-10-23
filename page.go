package igdb

// Page contains information on an IGDB entry
// for a multipurpose page used for Youtubers
// and media corporations.
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

// GetPage returns a single Page identified by the provided IGDB ID.
// Functional options may be provided but sorting and pagination will
// not have an effect due to GetPage only returning a single Page
// object and not a list of Pages.
func (c *Client) GetPage(id int, opts ...OptionFunc) (*Page, error) {
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

// GetPages returns a list of Pages identified by the provided list of IGDB
// IDs. Provide functional options to filter, sort, and paginate the results.
func (c *Client) GetPages(ids []int, opts ...OptionFunc) ([]*Page, error) {
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

// SearchPages returns a list of Pages found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate
// the results. Providing an empty query will instead retrieve an index of 
// Pages based solely on the provided options.
func (c *Client) SearchPages(qry string, opts ...OptionFunc) ([]*Page, error) {
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
