package igdb

type endpoint string

// Endpoints to their respective IGDB endpoints
const (
	CharacterEndpoint   endpoint = "characters/"
	CollectionEndpoint  endpoint = "collections/"
	CompanyEndpoint     endpoint = "companies/"
	CreditEndpoint      endpoint = "credits/"
	EngineEndpoint      endpoint = "game_engines/"
	FeedEndpoint        endpoint = "feeds/"
	FranchiseEndpoint   endpoint = "franchises/"
	GameEndpoint        endpoint = "games/"
	GameModeEndpoint    endpoint = "game_modes/"
	GenreEndpoint       endpoint = "genres/"
	KeywordEndpoint     endpoint = "keywords/"
	PageEndpoint        endpoint = "pages/"
	PersonEndpoint      endpoint = "people/"
	PlatformEndpoint    endpoint = "platforms/"
	PerspectiveEndpoint endpoint = "player_perspectives/"
	PulseEndpoint       endpoint = "pulses/"
	PulseGroupEndpoint  endpoint = "pulse_groups/"
	PulseSourceEndpoint endpoint = "pulse_sources/"
	ReleaseDateEndpoint endpoint = "release_dates/"
	ReviewEndpoint      endpoint = "reviews/"
	ThemeEndpoint       endpoint = "themes/"
	TitleEndpoint       endpoint = "titles/"
)

// GetEndpointFieldList returns a list of fields that represent the
// model of the data available at the given IGDB endpoint.
func (c *Client) GetEndpointFieldList(end endpoint) ([]string, error) {
	url := c.rootURL + string(end) + "meta"

	var f []string

	err := c.get(url, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetEndpointCount returns the count of entities available for
// the given IGDB endpoint.
func (c *Client) GetEndpointCount(end endpoint) (int, error) {
	url := c.rootURL + string(end) + "count"

	var ct struct {
		Count int `json:"count"`
	}

	err := c.get(url, &ct)
	if err != nil {
		return 0, err
	}

	return ct.Count, nil
}
