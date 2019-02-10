package igdb

import (
	"encoding/json"
	"github.com/Henry-Sarabia/apicalypse"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// igdbURL is the base URL for the IGDB API.
const igdbURL string = "https://api-v3.igdb.com/"

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
	http    *http.Client
	rootURL string
	key     string

	// Services
	Achievements             *AchievementService
	AchievementIcons         *AchievementIconService
	AgeRatings               *AgeRatingService
	AgeRatingContents        *AgeRatingContentService
	AlternativeNames         *AlternativeNameService
	Artworks                 *ArtworkService
	Characters               *CharacterService
	CharacterMugshots        *CharacterMugshotService
	Collections              *CollectionService
	Companies                *CompanyService
	CompanyLogos             *CompanyLogoService
	Covers                   *CoverService
	ExternalGames            *ExternalGameService
	Feeds                    *FeedService
	Franchises               *FranchiseService
	Games                    *GameService
	GameEngines              *GameEngineService
	GameEngineLogos          *GameEngineLogoService
	GameModes                *GameModeService
	GameVersions             *GameVersionService
	GameVersionFeatures      *GameVersionFeatureService
	GameVersionFeatureValues *GameVersionFeatureValueService
	GameVideos               *GameVideoService
	Genres                   *GenreService
	InvolvedCompanies        *InvolvedCompanyService
	Keywords                 *KeywordService
	MultiplayerModes         *MultiplayerModeService
	Pages                    *PageService
	PageBackgrounds          *PageBackgroundService
	PageLogos                *PageLogoService
	PageWebsites             *PageWebsiteService
}

// NewClient returns a new Client configured to communicate with the IGDB.
// The provided apiKey will be used to make requests on your behalf. The
// provided HTTP Client will be the client making requests to the IGDB.
// If no HTTP Client is provided, a default HTTP client is used instead.
//
// If you need an IGDB API key, please visit: https://api.igdb.com/signup
func NewClient(apiKey string, custom *http.Client) *Client {
	if custom == nil {
		custom = http.DefaultClient
	}
	c := &Client{}
	c.http = custom
	c.key = apiKey
	c.rootURL = igdbURL

	c.Achievements = &AchievementService{client: c, end: EndpointAchievement}
	c.AchievementIcons = &AchievementIconService{client: c, end: EndpointAchievementIcon}
	c.AgeRatings = &AgeRatingService{client: c, end: EndpointAgeRating}
	c.AgeRatingContents = &AgeRatingContentService{client: c, end: EndpointAgeRatingContent}
	c.AlternativeNames = &AlternativeNameService{client: c, end: EndpointAlternativeName}
	c.Artworks = &ArtworkService{client: c, end: EndpointArtwork}
	c.Characters = &CharacterService{client: c, end: EndpointCharacter}
	c.CharacterMugshots = &CharacterMugshotService{client: c, end: EndpointCharacterMugshot}
	c.Collections = &CollectionService{client: c, end: EndpointCollection}
	c.Companies = &CompanyService{client: c, end: EndpointCompany}
	c.CompanyLogos = &CompanyLogoService{client: c, end: EndpointCompanyLogo}
	c.Covers = &CoverService{client: c, end: EndpointCover}
	c.ExternalGames = &ExternalGameService{client: c, end: EndpointExternalGame}
	c.Feeds = &FeedService{client: c, end: EndpointFeed}
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
	c.Pages = &PageService{client: c, end: EndpointPage}
	c.PageBackgrounds = &PageBackgroundService{client: c, end: EndpointPageBackground}
	c.PageLogos = &PageLogoService{client: c, end: EndpointPageLogo}
	c.PageWebsites = &PageWebsiteService{client: c, end: EndpointPageWebsite}

	return c
}

// Request configures a new request for the provided URL and
// adds the necessary headers to communicate with the IGDB.
func (c *Client) request(end endpoint, opts ...FuncOption) (*http.Request, error) {
	unwrapped, err := unwrapOptions(opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create request with invalid options")
	}

	req, err := apicalypse.NewRequest("GET", c.rootURL+string(end), unwrapped...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot make request for '%s' endpoint", end)
	}

	req.Header.Add("user-key", c.key)
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

	if err = checkResults(b); err != nil {
		return err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return errors.Wrap(errInvalidJSON, err.Error())
	}

	return nil
}

// Get sends a GET request to the provided endpoint with the provided options and
// stores the results in the value pointed to by result.
func (c *Client) get(end endpoint, result interface{}, opts ...FuncOption) error {
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

// intsToStrings is a helper function that converts a slice of ints to a
// slice of strings.
func intsToStrings(ints []int) []string {
	var str []string
	for _, i := range ints {
		str = append(str, strconv.Itoa(i))
	}
	return str
}
