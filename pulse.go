package igdb

// PulseVideo contains the ID and
// category for a video related to
// a pulse.
type PulseVideo struct {
	Category int    `json:"category"`
	ID       string `json:"video_id"`
}

// Pulse contains information on an IGDB
// entry for a single news article.
type Pulse struct {
	ID          int          `json:"id"`
	PulseSource int          `json:"pulse_source"`
	Category    int          `json:"category"`
	Title       string       `json:"title"`
	Summary     string       `json:"summary"`
	URL         URL          `json:"url"`
	UID         string       `json:"uid"`          //perhaps switch to ID
	CreatedAt   int          `json:"created_at"`   // Unix time in milliseconds
	UpdatedAt   int          `json:"updated_at"`   // Unix time in milliseconds
	PublishedAt int          `json:"published_at"` // Unix time in milliseconds
	ImageURL    URL          `json:"image"`
	PulseImage  Image        `json:"pulse_image"`
	Videos      []PulseVideo `json:"videos"`
	Author      string       `json:"author"`
	Tags        []Tag        `json:"tags"`
	Ignored     interface{}  `json:"ignored"`
}

// GetPulse gets IGDB information for a Pulse identified by its unique
// IGDB ID. GetPulse returns a single Pulse identified by the provided
// IGDB ID. Functional options may be provided but sorting and pagination
// will not have an effect due to GetPulse only returning a single Pulse
// object and not a list of Pulses.
func (c *Client) GetPulse(id int, opts ...OptionFunc) (*Pulse, error) {
	url, err := c.singleURL(PulseEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}
	var p []Pulse

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPulses returns a list of Pulses identified by the provided list of
// IGDB IDs. Provide functional options to filter, sort, and paginate the
// results. Providing an empty list of IDs will instead retrieve an index
// of Pulses based solely on the provided options.
func (c *Client) GetPulses(ids []int, opts ...OptionFunc) ([]*Pulse, error) {
	url, err := c.multiURL(PulseEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Pulse

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPulses returns a list of Pulses found by searching the IGDB using the
// provided query. Provide functional options to filter, sort, and paginate the results.
func (c *Client) SearchPulses(qry string, opts ...OptionFunc) ([]*Pulse, error) {
	url, err := c.searchURL(PulseEndpoint, qry, opts...)
	if err != nil {
		return nil, err
	}
	var p []*Pulse

	err = c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
