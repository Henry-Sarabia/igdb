package igdb

// Person type
type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	URL        URL    `json:"url"`
	CreatedAt  int    `json:"created_at"` //unix epoch
	UpdatedAt  int    `json:"updated_at"` //unix epoch
	Mugshot    Image  `json:"mug_shot"`
	Games      []int  `json:"games"`
	Characters []int  `json:"characters"`
	VoiceActed []int  `json:"voice_acted"`
}
