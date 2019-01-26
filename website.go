package igdb

// Website represents a website and its URL; usually associated with a game.
// For more information visit: https://api-docs.igdb.com/#website
type Website struct {
	Category WebsiteCategory
	Trusted  bool
	URL      string
}

// WebsiteCategory represents the IGDB enumerated type Website Category which
// simply describes the category in which a website or URL falls under.
// Use the Stringer interface to access the corresponding category values
// as strings.
type WebsiteCategory int

const (
	WebsiteOfficial WebsiteCategory = iota + 1
	WebsiteWikia
	WebsiteWikipedia
	WebsiteFacebook
	WebsiteTwitter
	WebsiteTwitch
	_
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
