package igdb

// Pulse type
type Pulse struct {
	ID          int    `json:"id"`
	PulseSource int    `json:"pulse_source"` //not uint
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	URL         URL    `json:"url"`
	UID         string `json:"uid"`          //perhaps switch to ID
	CreatedAt   int    `json:"created_at"`   //unix epoch
	UpdatedAt   int    `json:"updated_at"`   //unix epoch
	PublishedAt int    `json:"published_at"` //unix epoch
	ImageURL    URL    `json:"image"`
	Author      string `json:"author"`
	Tags        []Tag  `json:"tags"`
}

// PulseGroup type
type PulseGroup struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Tags      []Tag  `json:"tags"`
	Pulses    []int  `json:"pulses"`
	Game      []int  `json:"game"`
}

// PulseSource type
type PulseSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Game int    `json:"game"`
	Page int    `json:"page"`
}
