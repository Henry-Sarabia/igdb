package igdb

//go:generate gomodifytags -file $GOFILE -struct Person -add-tags json -w

type Person struct {
	ID            int             `json:"id"`
	Bio           string          `json:"bio"`
	Characters    []int           `json:"characters"`
	Country       int             `json:"country"`
	CreatedAt     int             `json:"created_at"`
	CreditedGames []int           `json:"credited_games"`
	Description   string          `json:"description"`
	DOB           int             `json:"dob"`
	Gender        CharacterGender `json:"gender"`
	LovesCount    int             `json:"loves_count"`
	MugShot       int             `json:"mug_shot"`
	Name          string          `json:"name"`
	Nicknames     []string        `json:"nicknames"`
	Parent        int             `json:"parent"`
	Slug          string          `json:"slug"`
	UpdatedAt     int             `json:"updated_at"`
	URL           string          `json:"url"`
	VoiceActed    []int           `json:"voice_acted"`
	Websites      []int           `json:"websites"`
}
