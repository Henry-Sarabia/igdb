package igdb

// Page is
type Page struct {
	ID              ID          `json:"id"`
	Slug            string      `json:"slug"`
	CreatedAt       int         `json:"created_at"` //unix epoch
	UpdatedAt       int         `json:"updated_at"` //unix epoch
	Name            string      `json:"name"`
	Content         string      `json:"content"`
	Category        int         `json:"category"`
	Subcategory     int         `json:"sub_category"`
	Country         CountryCode `json:"country"`
	Color           int         `json:"color"`
	User            int         `json:"user"`
	Game            ID          `json:"game"`
	Company         ID          `json:"company"`
	Description     string      `json:"description"`
	PageFollowCount int         `json:"page_follows_count"`
	Logo            []Image     `json:"logo"`       // Might not be a slice
	Background      []Image     `json:"background"` // Might not be a slice
	Facebook        string      `json:"facebook"`
	Twitter         string      `json:"twitter"`
	Twitch          string      `json:"twitch"`
	Instagram       string      `json:"instagram"`
	Youtube         string      `json:"youtube"`
	Steam           string      `json:"steam"`
	Linkedin        string      `json:"linkedin"`
	Pinterest       string      `json:"pinterest"`
	Soundcloud      string      `json:"soundcloud"`
	GooglePlus      string      `json:"google_plus"`
	Reddit          string      `json:"reddit"`
	Battlenet       string      `json:"battlenet"`
	Origin          string      `json:"origin"`
	Uplay           string      `json:"uplay"`
	Discord         string      `json:"discord"`
}
