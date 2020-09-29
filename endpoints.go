package igdb

type endpoint string

// Public IGDB API endpoints
const (
	EndpointAgeRating                  endpoint = "age_ratings/"
	EndpointAgeRatingContent           endpoint = "age_rating_content_descriptions/"
	EndpointAlternativeName            endpoint = "alternative_names/"
	EndpointArtwork                    endpoint = "artworks/"
	EndpointCharacter                  endpoint = "characters/"
	EndpointCharacterMugshot           endpoint = "character_mug_shots/"
	EndpointCollection                 endpoint = "collections/"
	EndpointCompany                    endpoint = "companies/"
	EndpointCompanyLogo                endpoint = "company_logos/"
	EndpointCompanyWebsite             endpoint = "company_websites/"
	EndpointCover                      endpoint = "covers/"
	EndpointExternalGame               endpoint = "external_games/"
	EndpointFranchise                  endpoint = "franchises/"
	EndpointGame                       endpoint = "games/"
	EndpointGameEngine                 endpoint = "game_engines/"
	EndpointGameEngineLogo             endpoint = "game_engine_logos/"
	EndpointGameMode                   endpoint = "game_modes/"
	EndpointGameVersion                endpoint = "game_versions/"
	EndpointGameVersionFeature         endpoint = "game_version_features/"
	EndpointGameVersionFeatureValue    endpoint = "game_version_feature_values/"
	EndpointGameVideo                  endpoint = "game_videos/"
	EndpointGenre                      endpoint = "genres/"
	EndpointInvolvedCompany            endpoint = "involved_companies/"
	EndpointKeyword                    endpoint = "keywords/"
	EndpointMultiplayerMode            endpoint = "multiplayer_modes/"
	EndpointPlatform                   endpoint = "platforms/"
	EndpointPlatformLogo               endpoint = "platform_logos/"
	EndpointPlatformVersion            endpoint = "platform_versions/"
	EndpointPlatformVersionCompany     endpoint = "platform_version_companies/"
	EndpointPlatformVersionReleaseDate endpoint = "platform_version_release_dates/"
	EndpointPlatformWebsite            endpoint = "platform_websites/"
	EndpointPlayerPerspective          endpoint = "player_perspectives/"
	EndpointPlatformFamily             endpoint = "product_families/"
	EndpointPulse                      endpoint = "pulses/"
	EndpointReleaseDate                endpoint = "release_dates/"
	EndpointScreenshot                 endpoint = "screenshots/"
	EndpointSearch                     endpoint = "search/"
	EndpointTheme                      endpoint = "themes/"
	EndpointTitle                      endpoint = "titles/"
	EndpointWebsite                    endpoint = "websites/"
)

// Count contains the number of objects
// of a certain type counted in the IGDB.
type Count struct {
	Count int `json:"count"`
}

// getFields returns a list of fields that represent the
// model of the data available at the given IGDB endpoint.
func (c *Client) getFields(end endpoint) ([]string, error) {
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

// getCount returns the count of entities available for the given IGDB endpoint.
func (c *Client) getCount(end endpoint, opts ...Option) (int, error) {
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
