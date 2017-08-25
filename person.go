package igdb

// Person type
type Person struct {
	ID         ID     `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	URL        URL    `json:"url"`
	CreatedAt  int    `json:"created_at"` //unix epoch
	UpdatedAt  int    `json:"updated_at"` //unix epoch
	Mugshot    Image  `json:"mug_shot"`
	Games      []ID   `json:"games"`
	Characters []ID   `json:"characters"`
	VoiceActed []ID   `json:"voice_acted"`
}
