package igdb

//go:generate gomodifytags -file $GOFILE -struct Company -add-tags json -w

// Company represents a video game company.
// This includes both publishers and developers.
// For more information visit: https://api-docs.igdb.com/#company
type Company struct {
	ChangeDate         int                 `json:"change_date"`
	ChangeDateCategory CompanyDateCategory `json:"change_date_category"`
	ChangedCompanyID   int                 `json:"changed_company_id"`
	Country            int                 `json:"country"`
	CreatedAt          int                 `json:"created_at"`
	Description        string              `json:"description"`
	Developed          []int               `json:"developed"`
	Logo               int                 `json:"logo"`
	Name               string              `json:"name"`
	Parent             int                 `json:"parent"`
	Published          []int               `json:"published"`
	Slug               string              `json:"slug"`
	StartDate          int                 `json:"start_date"`
	StartDateCategory  CompanyDateCategory `json:"start_date_category"`
	UpdatedAt          int                 `json:"updated_at"`
	URL                string              `json:"url"`
	Websites           []int               `json:"websites"`
}

//go:generate stringer -type=CompanyDateCategory

type CompanyDateCategory int

const (
	DateYYYYMMMMDD CompanyDateCategory = iota
	DateYYYYMMMM
	DateYYYY
	DateYYYYQ1
	DateYYYYQ2
	DateYYYYQ3
	DateYYYYQ4
	DateTBD
)
