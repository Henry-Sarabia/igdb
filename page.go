package igdb

//go:generate gomodifytags -file $GOFILE -struct Page -add-tags json -w

// Page represents an entry in the multipurpose page system
// currently used for youtubers and media organizations.
// For more information visit: https://api-docs.igdb.com/#page
type Page struct {
	Background       int             `json:"background"`
	Battlenet        string          `json:"battlenet"`
	Category         PageCategory    `json:"category"`
	Color            PageColor       `json:"color"`
	Company          int             `json:"company"`
	Country          int             `json:"country"`
	CreatedAt        int             `json:"created_at"`
	Description      string          `json:"description"`
	Feed             int             `json:"feed"`
	Game             int             `json:"game"`
	Name             string          `json:"name"`
	Origin           string          `json:"origin"`
	PageFollowsCount int             `json:"page_follows_count"`
	PageLogo         int             `json:"page_logo"`
	Slug             string          `json:"slug"`
	SubCategory      PageSubCategory `json:"sub_category"`
	UpdatedAt        int             `json:"updated_at"`
	Uplay            string          `json:"uplay"`
	URL              string          `json:"url"`
	User             int             `json:"user"`
	Websites         []int           `json:"websites"`
}

//go:generate stringer -type=PageCategory,PageSubCategory,PageColor

type PageCategory int

const (
	PagePersonality PageCategory = iota + 1
	PageMediaOrganization
	PageContentCreator
	PageClanTeam
)

type PageSubCategory int

const (
	PageUser PageSubCategory = iota + 1
	PageGame
	PageCompany
	PageConsumer
	PageIndustry
	PageESports
)

type PageColor int

const (
	PageGreen PageColor = iota
	PageBlue
	PageRed
	PageOrange
	PagePink
	PageYellow
)
