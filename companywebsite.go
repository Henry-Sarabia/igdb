package igdb

//go:generate gomodifytags -file $GOFILE -struct CompanyWebsite -add-tags json -w

// CompanyWebsite represents a website for a specific company.
// For more information visit: https://api-docs.igdb.com/#company-website
type CompanyWebsite struct {
	Category CompanyWebsiteCategory `json:"category"`
	Trusted  bool                   `json:"trusted"`
	URL      string                 `json:"url"`
}

//go:generate stringer -type=CompanyWebsiteCategory

type CompanyWebsiteCategory int

const (
	WebsiteOfficial CompanyWebsiteCategory = iota + 1
	WebsiteWikia
	WebsiteWikipedia
	WebsiteFacebook
	WebsiteTwitter
	_
	WebsiteTwitch
	WebsiteInstagram
	WebsiteYoutube
	WebsiteIphone
	WebsiteIpad
	WebsiteAndroid
	WebsiteSteam
	WebsiteReddit
	WebsiteDiscord
	WebsiteGooglePlus
	WebsiteTumblr
	WebsiteLinkedin
	WebsitePinterest
	WebsiteSoundcloud
)
