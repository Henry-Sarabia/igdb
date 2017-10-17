package igdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// igdbURL is the base URL for the IGDB API.
const igdbURL string = "https://api-2445582011268.apicast.io/"

// ErrNegativeID is returned by a client when a negative
// ID is used as an argument in an API call.
var ErrNegativeID = errors.New("igdb.Client: negative ID")

// URL represents a URL as a string.
type URL string

// Client wraps a typical http.Client.
// Client is used for all IGDB API calls.
type Client struct {
	http    *http.Client
	rootURL string
}

// NewClient returns a new Client set with a default HTTP
// client and the default IGDB root URL.
func NewClient() Client {
	return Client{http: http.DefaultClient, rootURL: igdbURL}
}

// get sends a GET request to the provided url and stores
// the response in the provided result empty interface.
func (c *Client) get(url string, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("user-key", APIkey)
	req.Header.Add("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = c.checkError(resp)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return err
	}
	return nil
}

// singleURL creates a URL configured to request a single IGDB object
// identified by its unique IGDB ID using the provided endpoint.
func (c *Client) singleURL(end endpoint, id int, opts ...optionFunc) (string, error) {
	if id < 0 {
		return "", ErrNegativeID
	}
	opt, err := newOpt(opts...)
	if err != nil {
		return "", err
	}

	url := c.rootURL + string(end) + strconv.Itoa(id)
	url = encodeURL(&opt.Values, url)

	return url, nil
}

// multiURL creates a URL configured to request multiple IGDB object identified
// by their unique IGDB IDs using the provided endpoint.
func (c *Client) multiURL(end endpoint, ids []int, opts ...optionFunc) (string, error) {
	for _, id := range ids {
		if id < 0 {
			return "", ErrNegativeID
		}
	}
	opt, err := newOpt(opts...)
	if err != nil {
		return "", err
	}

	url := c.rootURL + string(end) + intsToCommaString(ids)
	url = encodeURL(&opt.Values, url)

	return url, nil
}

// searchURL creates a URL configured to search the IGDB based on the given query
// using the provided endpoint.
func (c *Client) searchURL(end endpoint, qry string, opts ...optionFunc) (string, error) {
	opts = append(opts, optSearch(qry))
	opt, err := newOpt(opts...)
	if err != nil {
		return "", err
	}

	url := c.rootURL + string(end)
	url = encodeURL(&opt.Values, url)

	return url, nil
}

// intsToStrings converts a slice of ints
// to a slice of strings.
func intsToStrings(ints []int) []string {
	var str []string
	for _, i := range ints {
		str = append(str, strconv.Itoa(i))
	}
	return str
}

// intsToCommaString returns a comma separated
// list of ints as a single string.
func intsToCommaString(ints []int) string {
	s := intsToStrings(ints)
	return strings.Join(s, ",")
}
