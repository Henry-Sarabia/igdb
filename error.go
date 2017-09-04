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

func (c *Client) checkError(resp *http.Response) error {
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
