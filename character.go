package igdb

import "strconv"

// GenderCode codes
type GenderCode int

// SpeciesCode codes
type SpeciesCode int

// Character is
type Character struct {
	ID        ID          `json:"id"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
	URL       URL         `json:"url"`
	CreatedAt int         `json:"created_at"`
	UpdatedAt int         `json:"updated_at"`
	MugShot   Image       `json:"mug_shot"`
	Gender    GenderCode  `json:"gender"`
	AKAs      []string    `json:"akas"`
	Species   SpeciesCode `json:"species"`
	Games     []ID        `json:"games"`
	People    []ID        `json:"people"`
}

// GetCharacter gets IGDB information for a character identified by its unique IGDB ID.
func (c *Client) GetCharacter(id int, opts ...OptionFunc) (*Character, error) {
	opt := newOpt()

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := rootURL + "characters/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var ch []Character

	err := c.get(url, &ch)
	if err != nil {
		return nil, err
	}

	return &ch[0], nil
}
