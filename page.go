package igdb

// PageService handles all the API
// calls for the IGDB Page endpoint.
type PageService service

// Page contains information on an IGDB entry for a multipurpose
// page used for Youtubers and media corporations.
//
// For more information, visit: https://igdb.github.io/api/endpoints/page/
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

// Get returns a single Page identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Pages, an error is returned.
func (ps *PageService) Get(id int, opts ...FuncOption) (*Page, error) {
	url, err := ps.client.singleURL(PageEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var pg []Page

	err = ps.client.get(url, &pg)
	if err != nil {
		return nil, err
	}

	return &pg[0], nil
}

// List returns a list of Pages identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of Pages based solely on the provided
// options. Any ID that does not match a Page is ignored. If none of the IDs
// match a Page, an error is returned.
func (ps *PageService) List(ids []int, opts ...FuncOption) ([]*Page, error) {
	url, err := ps.client.multiURL(PageEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var pg []*Page

	err = ps.client.get(url, &pg)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

// Search returns a list of Pages found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Pages are found using the provided query, an error is returned.
func (ps *PageService) Search(qry string, opts ...FuncOption) ([]*Page, error) {
	url, err := ps.client.searchURL(PageEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}

	var pg []*Page

	err = ps.client.get(url, &pg)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

// Count returns the number of Pages available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Pages to count.
func (ps *PageService) Count(opts ...FuncOption) (int, error) {
	ct, err := ps.client.getEndpointCount(PageEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Page object.
func (ps *PageService) ListFields() ([]string, error) {
	fl, err := ps.client.getEndpointFieldList(PageEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
