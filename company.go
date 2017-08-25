package igdb

// CountryCode code ISO-3316-1
type CountryCode int

// Company is
type Company struct {
	ID                 ID           `json:"ID"`
	Name               string       `json:"name"`
	Slug               string       `json:"slug"`
	URL                URL          `json:"url"`
	CreatedAt          int          `json:"created_at"`
	UpdatedAt          int          `json:"updated_at"`
	Logo               Image        `json:"logo"`
	Description        string       `json:"description"`
	Country            CountryCode  `json:"country"`
	Website            string       `json:"website"`
	StartDate          int          `json:"start_date"` //unix epoch
	StartDateCategory  DateCategory `json:"start_date_category"`
	ChangedID          ID           `json:"changed_company_id"`
	ChangeDate         int          `json:"changed_date"` //unix epoch
	ChangeDateCategory DateCategory `json:"change_date_category"`
	Twitter            string       `json:"twitter"`
	Published          []ID         `json:"published"`
	Developed          []ID         `json:"developed"`
}
