package igdb

// PulseVideo contains information for
// a video specifically for a pulse.
type PulseVideo struct {
	Category int    `json:"category"`
	ID       string `json:"video_id"`
}

// Pulse type
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

// GetPulse gets IGDB information for a pulse identified by its unique IGDB ID.
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

// GetPulses gets IGDB information for a list of pulses identified by their
// unique IGDB IDs.
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

// SearchPulses searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
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
