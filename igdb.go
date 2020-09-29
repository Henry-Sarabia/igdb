package igdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Henry-Sarabia/apicalypse"
	"github.com/pkg/errors"
)

// igdbURL is the base URL for the IGDB API.
const igdbURL string = "https://api.igdb.com/v4/"

// service is the underlying struct that handles
// all API calls for different IGDB endpoints.
type service struct {
	client *Client
	end    endpoint
}

// Client wraps an HTTP Client used to communicate with the IGDB,
// the root URL of the IGDB, and the user's IGDB API key.
// Client also initializes all the separate services to communicate
// with each individual IGDB API endpoint.
type Client struct {
	http     *http.Client
	rootURL  string
	clientID string
	token    string

	// Services
	AgeRatings                  *AgeRatingService
	AgeRatingContents           *AgeRatingContentService
	AlternativeNames            *AlternativeNameService
	Artworks                    *ArtworkService
	Characters                  *CharacterService
	CharacterMugshots           *CharacterMugshotService
	Collections                 *CollectionService
	Companies                   *CompanyService
	CompanyLogos                *CompanyLogoService
	CompanyWebsites             *CompanyWebsiteService
	Covers                      *CoverService
	ExternalGames               *ExternalGameService
	Franchises                  *FranchiseService
	Games                       *GameService
	GameEngines                 *GameEngineService
	GameEngineLogos             *GameEngineLogoService
	GameModes                   *GameModeService
	GameVersions                *GameVersionService
	GameVersionFeatures         *GameVersionFeatureService
	GameVersionFeatureValues    *GameVersionFeatureValueService
	GameVideos                  *GameVideoService
	Genres                      *GenreService
	InvolvedCompanies           *InvolvedCompanyService
	Keywords                    *KeywordService
	MultiplayerModes            *MultiplayerModeService
	Platforms                   *PlatformService
	PlatformLogos               *PlatformLogoService
	PlatformVersions            *PlatformVersionService
	PlatformVersionCompanies    *PlatformVersionCompanyService
	PlatformVersionReleaseDates *PlatformVersionReleaseDateService
	PlatformWebsites            *PlatformWebsiteService
	PlayerPerspectives          *PlayerPerspectiveService
	PlatformFamilies            *PlatformFamilyService
	ReleaseDates                *ReleaseDateService
	Screenshots                 *ScreenshotService
	Themes                      *ThemeService
	Websites                    *WebsiteService

	// Private Services
	TestDummies *TestDummyService
}

// NewClient returns a new Client configured to communicate with the IGDB.
// The provided clientID and appAccessToken will be used to make requests on your behalf.
// The provided HTTP Client will be the client making requests to the IGDB. If no
// HTTP Client is provided, a default HTTP client is used instead.
//
// If you need an IGDB/Twitch API keys, please visit: https://api-docs.igdb.com/#account-creation
func NewClient(clientID string, appAccessToken string, custom *http.Client) *Client {
	if custom == nil {
		custom = http.DefaultClient
	}

	c := &Client{
		http:     custom,
		rootURL:  igdbURL,
		clientID: clientID,
		token:    appAccessToken,
	}

	c.AgeRatings = &AgeRatingService{client: c, end: EndpointAgeRating}
	c.AgeRatingContents = &AgeRatingContentService{client: c, end: EndpointAgeRatingContent}
	c.AlternativeNames = &AlternativeNameService{client: c, end: EndpointAlternativeName}
	c.Artworks = &ArtworkService{client: c, end: EndpointArtwork}
	c.Characters = &CharacterService{client: c, end: EndpointCharacter}
	c.CharacterMugshots = &CharacterMugshotService{client: c, end: EndpointCharacterMugshot}
	c.Collections = &CollectionService{client: c, end: EndpointCollection}
	c.Companies = &CompanyService{client: c, end: EndpointCompany}
	c.CompanyLogos = &CompanyLogoService{client: c, end: EndpointCompanyLogo}
	c.CompanyWebsites = &CompanyWebsiteService{client: c, end: EndpointCompanyWebsite}
	c.Covers = &CoverService{client: c, end: EndpointCover}
	c.ExternalGames = &ExternalGameService{client: c, end: EndpointExternalGame}
	c.Franchises = &FranchiseService{client: c, end: EndpointFranchise}
	c.Games = &GameService{client: c, end: EndpointGame}
	c.GameEngines = &GameEngineService{client: c, end: EndpointGameEngine}
	c.GameEngineLogos = &GameEngineLogoService{client: c, end: EndpointGameEngineLogo}
	c.GameModes = &GameModeService{client: c, end: EndpointGameMode}
	c.GameVersions = &GameVersionService{client: c, end: EndpointGameVersion}
	c.GameVersionFeatures = &GameVersionFeatureService{client: c, end: EndpointGameVersionFeature}
	c.GameVersionFeatureValues = &GameVersionFeatureValueService{client: c, end: EndpointGameVersionFeatureValue}
	c.GameVideos = &GameVideoService{client: c, end: EndpointGameVideo}
	c.Genres = &GenreService{client: c, end: EndpointGenre}
	c.InvolvedCompanies = &InvolvedCompanyService{client: c, end: EndpointInvolvedCompany}
	c.Keywords = &KeywordService{client: c, end: EndpointKeyword}
	c.MultiplayerModes = &MultiplayerModeService{client: c, end: EndpointMultiplayerMode}
	c.Platforms = &PlatformService{client: c, end: EndpointPlatform}
	c.PlatformLogos = &PlatformLogoService{client: c, end: EndpointPlatformLogo}
	c.PlatformVersions = &PlatformVersionService{client: c, end: EndpointPlatformVersion}
	c.PlatformVersionCompanies = &PlatformVersionCompanyService{client: c, end: EndpointPlatformVersionCompany}
	c.PlatformVersionReleaseDates = &PlatformVersionReleaseDateService{client: c, end: EndpointPlatformVersionReleaseDate}
	c.PlatformWebsites = &PlatformWebsiteService{client: c, end: EndpointPlatformWebsite}
	c.PlayerPerspectives = &PlayerPerspectiveService{client: c, end: EndpointPlayerPerspective}
	c.PlatformFamilies = &PlatformFamilyService{client: c, end: EndpointPlatformFamily}
	c.ReleaseDates = &ReleaseDateService{client: c, end: EndpointReleaseDate}
	c.Screenshots = &ScreenshotService{client: c, end: EndpointScreenshot}
	c.Themes = &ThemeService{client: c, end: EndpointTheme}
	c.Websites = &WebsiteService{client: c, end: EndpointWebsite}

	c.TestDummies = &TestDummyService{client: c, end: EndpointTestDummy}
	return c
}

// Request configures a new request for the provided URL and
// adds the necessary headers to communicate with the IGDB.
func (c *Client) request(end endpoint, opts ...Option) (*http.Request, error) {
	unwrapped, err := unwrapOptions(opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create request with invalid options")
	}

	req, err := apicalypse.NewRequest("POST", c.rootURL+string(end), unwrapped...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot make request for '%s' endpoint", end)
	}

	req.Header.Add("client-id", c.clientID)
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("x-user-agent", "HenrySarabia/igdb")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// Send sends the provided request and stores the response in the value pointed to by result.
// The response will be checked and return any errors.
func (c *Client) send(req *http.Request, result interface{}) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client cannot send request")
	}
	defer resp.Body.Close()

	if err = checkResponse(resp); err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "cannot read response body")
	}

	if isBracketPair(b) {
		return ErrNoResults
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return errors.Wrap(errInvalidJSON, err.Error())
	}

	return nil
}

// Get sends a GET request to the provided endpoint with the provided options and
// stores the results in the value pointed to by result.
func (c *Client) get(end endpoint, result interface{}, opts ...Option) error {
	req, err := c.request(end, opts...)
	if err != nil {
		return err
	}

	err = c.send(req, result)
	if err != nil {
		return errors.Wrap(err, "cannot make GET request")
	}

	return nil
}
