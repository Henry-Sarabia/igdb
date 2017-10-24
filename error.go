package igdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// errEndOfJSON occurs when encountering an unexpected end of JSON input.
var errEndOfJSON = errors.New("unexpected end of JSON input")

// Error contains information on an
// error returned from an IGDB API call.
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// checkError checks the provided HTTP Response for errors
// returned by the IGDB.
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

	msg := fmt.Sprintf("Status %d", e.Status)
	if e.Message != "" {
		msg += fmt.Sprintf(" - %v", e.Message)
	}
	return errors.New(msg)
}
