package igdb

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// igdbURL is the base URL for the IGDB API.
const igdbURL string = "https://api-v3.igdb.com/"

// Errors returned when creating URLs for API calls.
var (
	// ErrNegativeID occurs when a negative ID is used as an argument in an API call.
	ErrNegativeID = errors.New("igdb.Client: negative ID")
	// ErrNoResults occurs when the IGDB returns no results
	ErrNoResults = errors.New("igdb.Client: no results")
)

// URL represents a URL as a string.
type URL string

// service is the underlying struct that handles
// all API calls for different IGDB endpoints.
type service struct {
	client *Client
}

// Client wraps an HTTP Client used to communicate with the IGDB,
// the root URL of the IGDB, and the user's IGDB API key.
// Client also initializes all the separate services to communicate
// with each individual IGDB API endpoint.
type Client struct {
	http    *http.Client
	rootURL string
	key     string

	common service

	// Services
	Games *GameService
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

	c.common.client = c
	c.Games = (*GameService)(&c.common)

	return c
}

// get sends a GET request to the provided url and stores
// the response in the provided result empty interface.
// The response will be checked and return any errors.
func (c *Client) get(req *http.Request, result interface{}) error {
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

	return err
}

// newRequest configures a new request for the provided URL and
// adds the necesarry headers to communicate with the IGDB.
//func (c *Client) newRequest(url string) (*http.Request, error) {
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Add("user-key", c.key)
//	req.Header.Add("Accept", "application/json")
//
//	return req, nil
//}

// searchURL creates a URL configured to search the IGDB based on the given query using
// the provided endpoint and options.
//func (c *Client) searchURL(end endpoint, qry string, opts ...FuncOption) (string, error) {
//	if strings.TrimSpace(qry) == "" {
//		return "", ErrEmptyQuery
//	}
//
//	opts = append(opts, setSearch(qry))
//	opt, err := newOpt(opts...)
//	if err != nil {
//		return "", err
//	}
//
//	url := c.rootURL + string(end)
//	url = encodeURL(&opt.Values, url)
//
//	return url, nil
//}

// countURL creates a URL configured to retrieve the count of IGDB objects
// using the provided endpoint and options.
//func (c *Client) countURL(end endpoint, opts ...FuncOption) (string, error) {
//	opt, err := newOpt(opts...)
//	if err != nil {
//		return "", err
//	}
//
//	url := c.rootURL + string(end) + "count"
//	url = encodeURL(&opt.Values, url)
//
//	return url, nil
//}

// Byte representations of ASCII characters. Used for empty result checks.
const (
	// openBracketASCII represents the ASCII code for an open bracket.
	openBracketASCII = 91
	// closedBracketASCII represents the ASCII code for a closed bracket.
	closedBracketASCII = 93
)

// checkResults checks if the results of an API call are an empty array.
// If they are, an error is returned. Otherwise, nil is returned.
func checkResults(r []byte) error {
	if len(r) != 2 {
		return nil
	}

	if r[0] == openBracketASCII && r[1] == closedBracketASCII {
		return ErrNoResults
	}

	return nil
}

// IntsToStrings is a helper function that converts a slice of ints to a
// slice of strings. Useful for functions that require a variadic number
// of strings instead of ints such as SetFilter.
func IntsToStrings(ints []int) []string {
	var str []string
	for _, i := range ints {
		str = append(str, strconv.Itoa(i))
	}
	return str
}

// intsToCommaString is a helper function that returns a comma separated
// list of ints as a single string.
func intsToCommaString(ints []int) string {
	s := IntsToStrings(ints)
	return strings.Join(s, ",")
}
