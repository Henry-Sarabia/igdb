package igdb

//go:generate gomodifytags -file $GOFILE -struct Review -add-tags json -w

type Review struct {
	ID             int            `json:"id"`
	Category       ReviewCategory `json:"category"`
	Conclusion     string         `json:"conclusion"`
	Content        string         `json:"content"`
	CreatedAt      int            `json:"created_at"`
	Game           int            `json:"game"`
	Introduction   string         `json:"introduction"`
	Likes          int            `json:"likes"`
	NegativePoints string         `json:"negative_points"`
	Platform       int            `json:"platform"`
	PositivePoints string         `json:"positive_points"`
	Slug           string         `json:"slug"`
	Title          string         `json:"title"`
	UpdatedAt      int            `json:"updated_at"`
	URL            string         `json:"url"`
	User           int            `json:"user"`
	UserRating     int            `json:"user_rating"`
	Video          int            `json:"video"`
	Views          int            `json:"views"`
}

//go:generate stringer -type=ReviewCategory

type ReviewCategory int

const (
	ReviewText ReviewCategory = iota + 1
	ReviewVid
)
