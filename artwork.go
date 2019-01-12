package igdb

import (
	"encoding/json"
	"github.com/Henry-Sarabia/apicalypse"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Artwork struct {
	AlphaChannel bool   `json:"alpha_channel,omitempty"`
	Animated     bool   `json:"animated,omitempty"`
	Height       int    `json:"height,omitempty"`
	ImageID      string `json:"image_id,omitempty"`
	URL          string `json:"url,omitempty"`
	Width        int    `json:"width,omitempty"`
}

func (c *Client) GetArtwork(opts ...FuncOption) ([]Artwork, error) {
	req, err := c.request(GameEndpoint, opts...)
	if err != nil {
		return nil, err
	}

	var a []Artwork

	err = c.do(req, &a)
	if err != nil {
		return nil, errors.Wrap(err, "cannot make request")
	}

	return a, nil
}

func (c *Client) do(req *http.Request, result interface{}) error {
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
