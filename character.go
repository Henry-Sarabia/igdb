package igdb

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
