package igdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Client wraps a typical http.Client.
type Client struct {
	http    *http.Client
	rootURL string
}

// NewClient returns a new client.
func NewClient() Client {
	return Client{http: http.DefaultClient, rootURL: igdbURL}
}

// get sends a GET request to the url and stores the response
// in the result interface{} if no errors are encountered.
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

	if resp.StatusCode != http.StatusOK {
		err = c.checkError(resp)
		if err != nil {
			return err
		}
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

// getRaw sends a GET request to the url and returns a raw
// encoded JSON value.
func (c *Client) getRaw(url string) (json.RawMessage, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", APIkey)
	req.Header.Add("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = c.checkError(resp)
		if err != nil {
			return nil, err
		}
	}

	var raw json.RawMessage

	raw, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// singleURL creates a URL configured to request a single IGDB entity
// identified by its unique IGDB ID using the given endpoint.
func (c *Client) singleURL(end endpoint, id int, opts ...OptionFunc) string {
	opt := newOpt(opts...)

	url := c.rootURL + string(end) + strconv.Itoa(id)
	url = encodeURL(opt.Values, url)

	return url
}

// multiURL creates a URL configured to request multiple IGDB entities identified
// by their unique IGDB IDs using the given endpoint.
func (c *Client) multiURL(end endpoint, ids []int, opts ...OptionFunc) string {
	opt := newOpt(opts...)

	url := c.rootURL + string(end) + intsToCommaString(ids)
	url = encodeURL(opt.Values, url)

	return url
}

// singleGet returns a raw encoded JSON value describing the IGDB information for a single
// entity identified by its unique IGDB ID.
func (c *Client) singleGet(id int, end endpoint, opts ...OptionFunc) (json.RawMessage, error) {
	opt := newOpt(opts...)

	url := c.rootURL + string(end) + strconv.Itoa(id)
	url = encodeURL(opt.Values, url)

	raw, err := c.getRaw(url)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// multiGet returns a raw encoded JSON value describing the IGDB information for a list of
// entities identified by their unique IGDB IDs.
func (c *Client) multiGet(ids []int, end endpoint, opts ...OptionFunc) (json.RawMessage, error) {
	opt := newOpt(opts...)

	url := c.rootURL + string(end) + intsToCommaString(ids)
	url = encodeURL(opt.Values, url)

	raw, err := c.getRaw(url)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// search searches the IGDB using the given query and returns a raw encoded JSON value describing
// the results of the search. Use functional options for pagination and result sorting by parameter.
func (c *Client) search(qry string, end endpoint, opts ...OptionFunc) (json.RawMessage, error) {
	opts = append(opts, optSearch(qry))
	opt := newOpt(opts...)

	url := c.rootURL + string(end)
	url = encodeURL(opt.Values, url)

	raw, err := c.getRaw(url)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// Encoder is implemented by any values that has
// an encode method, which returns the "encoded"
// format for that value. The Encode method is
// used to print a case-sensitive key value map
// used for query parameters or form values as a
// string.
type Encoder interface {
	Encode() string
}

// encodeURL encodes the url with the query
// parameters provided by the encoder.
func encodeURL(enc Encoder, url string) string {
	url = strings.Replace(url, " ", "", -1)

	if values := enc.Encode(); values != "" {
		url += "?" + values
	}

	return url
}
