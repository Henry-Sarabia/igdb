package igdb

//go:generate gomodifytags -file $GOFILE -struct PlatformVersionCompany -add-tags json -w

// PlatformVersionCompany represents a platform developer.
// For more information visit: https://api-docs.igdb.com/#platform-version-company
type PlatformVersionCompany struct {
	Comment      string `json:"comment"`
	Company      int    `json:"company"`
	Developer    bool   `json:"developer"`
	Manufacturer bool   `json:"manufacturer"`
}
