package igdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
