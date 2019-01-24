package igdb

//go:generate gomodifytags -file $GOFILE -struct Character -add-tags json -w

// Character represents a video game character.
// For more information visit: https://api-docs.igdb.com/#character
type Character struct {
	AKAs        []string         `json:"ak_as"`
	CountryName string           `json:"country_name"`
	CreatedAt   int              `json:"created_at"`
	Description string           `json:"description"`
	Games       []int            `json:"games"`
	Gender      CharacterGender  `json:"gender"`
	MugShot     int              `json:"mug_shot"`
	Name        string           `json:"name"`
	People      []int            `json:"people"`
	Slug        string           `json:"slug"`
	Species     CharacterSpecies `json:"species"`
	UpdatedAt   int              `json:"updated_at"`
	URL         string           `json:"url"`
}

//go:generate stringer -type=CharacterGender,CharacterSpecies

type CharacterGender int

const (
	GenderMale CharacterGender = iota + 1
	GenderFemale
	GenderOther
)

type CharacterSpecies int

const (
	SpeciesHuman CharacterSpecies = iota + 1
	SpeciesAlien
	SpeciesAnimal
	SpeciesAndroid
	SpeciesUnknown
)
