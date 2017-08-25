package igdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Client type
type Client struct {
	http *http.Client
}

// NewClient returns a new client.
func NewClient() Client {
	return Client{http: http.DefaultClient}
}

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
