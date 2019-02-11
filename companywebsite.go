package igdb

// CompanyWebsite represents a website for a specific company.
// For more information visit: https://api-docs.igdb.com/#company-website
type CompanyWebsite struct {
	ID       int             `json:"id"`
	Category WebsiteCategory `json:"category"`
	Trusted  bool            `json:"trusted"`
	URL      string          `json:"url"`
}
