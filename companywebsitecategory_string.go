// Code generated by "stringer -type=CompanyWebsiteCategory"; DO NOT EDIT.

package igdb

import "strconv"

const (
	_CompanyWebsiteCategory_name_0 = "WebsiteOfficialWebsiteWikiaWebsiteWikipediaWebsiteFacebookWebsiteTwitter"
	_CompanyWebsiteCategory_name_1 = "WebsiteTwitchWebsiteInstagramWebsiteYoutubeWebsiteIphoneWebsiteIpadWebsiteAndroidWebsiteSteamWebsiteRedditWebsiteDiscordWebsiteGooglePlusWebsiteTumblrWebsiteLinkedinWebsitePinterestWebsiteSoundcloud"
)

var (
	_CompanyWebsiteCategory_index_0 = [...]uint8{0, 15, 27, 43, 58, 72}
	_CompanyWebsiteCategory_index_1 = [...]uint8{0, 13, 29, 43, 56, 67, 81, 93, 106, 120, 137, 150, 165, 181, 198}
)

func (i CompanyWebsiteCategory) String() string {
	switch {
	case 1 <= i && i <= 5:
		i -= 1
		return _CompanyWebsiteCategory_name_0[_CompanyWebsiteCategory_index_0[i]:_CompanyWebsiteCategory_index_0[i+1]]
	case 7 <= i && i <= 20:
		i -= 7
		return _CompanyWebsiteCategory_name_1[_CompanyWebsiteCategory_index_1[i]:_CompanyWebsiteCategory_index_1[i+1]]
	default:
		return "CompanyWebsiteCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}