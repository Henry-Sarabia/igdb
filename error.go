package igdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Error contains information from the IGDB
// when an API call receives an error in response.
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// checkError checks an http.Response for an error response
// from the IGDB servers.
func (c *Client) checkError(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var e Error

	err = json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Status %d - %s", e.Status, e.Message)
	return errors.New(msg)
}
