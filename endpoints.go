package igdb

type endpoint string

// Available IGDB API endpoints
const (
	EndpointAchievement                 endpoint = "achievements/"
	EndpointAchievementIcon             endpoint = "achievement_icons/"
	EndpointAgeRating                   endpoint = "age_ratings/"
	EndpointAgeRatingContentDescription endpoint = "age_rating_content_descriptions/"
	EndpointAlternativeName             endpoint = "alternative_names/"
	EndpointArtwork                     endpoint = "artworks/"
	EndpointCharacter                   endpoint = "characters/"
	EndpointCharacterMugShot            endpoint = "character_mug_shots/"
	EndpointCollection                  endpoint = "collections/"
	EndpointCompany                     endpoint = "companies/"
	EndpointCompanyLogo                 endpoint = "company_logos/"
	EndpointCompanyWebsite              endpoint = "company_websites/"
	EndpointCover                       endpoint = "covers/"
	EndpointExternalGame                endpoint = "external_games/"
	EndpointFeed                        endpoint = "feeds/"
	EndpointFranchise                   endpoint = "franchises/"
	EndpointGame                        endpoint = "games/"
	EndpointGameEngine                  endpoint = "game_engines/"
	EndpointGameEngineLogo              endpoint = "game_engine_logos/"
	EndpointGameMode                    endpoint = "game_modes/"
	EndpointGameVersion                 endpoint = "game_versions/"
	EndpointGameVersionFeature          endpoint = "game_version_features/"
	EndpointGameVersionFeatureValue     endpoint = "game_version_feature_values/"
	EndpointGameVideo                   endpoint = "game_videos/"
	EndpointGenre                       endpoint = "genres/"
	EndpointInvolvedCompany             endpoint = "involved_companies/"
	EndpointKeyword                     endpoint = "keywords/"
	EndpointMultiplayerMode             endpoint = "multiplayer_modes/"
	EndpointPage                        endpoint = "pages/"
	EndpointPackBackground              endpoint = "page_backgrounds/"
	EndpointPageLogo                    endpoint = "page_logos/"
	EndpointPageWebsite                 endpoint = "page_websites/"
	EndpointPlatform                    endpoint = "platforms/"
	EndpointPlatformLogo                endpoint = "platform_logos/"
	EndpointPlatformVersion             endpoint = "platform_versions/"
	EndpointPlatformVersionCompany      endpoint = "platform_version_companies/"
	EndpointPlatformVersionReleaseDate  endpoint = "platform_version_release_dates/"
	EndpointPlatformWebsite             endpoint = "platform_websites/"
	EndpointPlayerPerspective           endpoint = "player_perspectives/"
	EndpointProductFamily               endpoint = "product_families/"
	EndpointPulse                       endpoint = "pulses/"
	EndpointPulseGroup                  endpoint = "pulse_groups/"
	EndpointPulseSource                 endpoint = "pulse_sources/"
	EndpointPulseURL                    endpoint = "pulse_urls/"
	EndpointReleaseDate                 endpoint = "release_dates/"
	EndpointScreenshot                  endpoint = "screenshots/"
	EndpointSearch                      endpoint = "search/"
	EndpointTheme                       endpoint = "themes/"
	EndpointTimeToBeat                  endpoint = "time_to_beats/"
	EndpointTitle                       endpoint = "titles/"
	EndpointWebsite                     endpoint = "websites/"
)

// Count contains the number of objects
// of a certain type counted in the IGDB.
type Count struct {
	Count int `json:"count"`
}

// getEndpointFieldList returns a list of fields that represent the
// model of the data available at the given IGDB endpoint.
func (c *Client) getEndpointFieldList(end endpoint) ([]string, error) {
	req, err := c.request(end + "meta")
	if err != nil {
		return nil, err
	}

	var f []string

	if err = c.send(req, &f); err != nil && err != ErrNoResults {
		return nil, err
	}

	return f, nil
}

// getEndpointCount returns the count of entities available for the given IGDB endpoint.
func (c *Client) getEndpointCount(end endpoint, opts ...FuncOption) (int, error) {
	req, err := c.request(end+"count", opts...)
	if err != nil {
		return 0, err
	}

	var ct Count

	err = c.send(req, &ct)
	if err != nil {
		return 0, err
	}

	return ct.Count, nil
}
