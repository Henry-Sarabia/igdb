package igdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	// errInvalidJSON occurs when encountering an unexpected end of JSON input.
	errInvalidJSON = errors.New("invalid JSON")
	ErrEmptyIDs    = errors.New("IDs argument empty")
)

// Errors returned when the IGDB responds with a problematic status code.
//
// For more information, visit: https://igdb.github.io/api/references/response-codes/
var (
	ErrAuthFailed    = errors.New("IGDB: authentication failed - need valid API key in user-key header")
	ErrBadRequest    = errors.New("IGDB: bad request - check query parameters")
	ErrInternalError = errors.New("IGDB: internal error - report bug")
)

// ServerError contains information on an
// error returned from an IGDB API call.
type ServerError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ServerError) Error() string {
	return "server error: status: " + strconv.Itoa(e.Status) + " message: " + e.Message
}

// checkResponse checks the provided HTTP response
// for errors returned by the IGDB.
func checkResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized, http.StatusForbidden:
		return ErrAuthFailed
	case http.StatusInternalServerError:
		return ErrInternalError
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var e ServerError

	err = json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	return e
}

// ErrNoResults occurs when the IGDB returns no results
var ErrNoResults = errors.New("results are empty")

// Byte representations of ASCII characters. Used for empty result checks.
const (
	// openBracketASCII represents the ASCII code for an open bracket.
	openBracketASCII = 91
	// closedBracketASCII represents the ASCII code for a closed bracket.
	closedBracketASCII = 93
)

// isBracketPair returns true if the provided slice of bytes is equivalent to
// an open and closed bracket pair in byte representation. Otherwise, false is returned.
// Used primarily to check if the results of an API call are an empty array.
func isBracketPair(b []byte) bool {
	if len(b) != 2 {
		return false
	}

	if b[0] == openBracketASCII && b[1] == closedBracketASCII {
		return true
	}

	return false
}
