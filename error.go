package igdb

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	// ErrNegativeID occurs when a negative ID is used as an argument in an API call.
	ErrNegativeID = errors.New("ID cannot be negative")
	// ErrEmptyIDs occurs when a List function is called without a populated int slice.
	ErrEmptyIDs = errors.New("IDs argument empty")
	// ErrNoResults occurs when the IGDB returns an empty array, void of results.
	ErrNoResults = errors.New("results are empty")
	// errInvalidJSON occurs when encountering an unexpected end of JSON input.
	errInvalidJSON = errors.New("invalid JSON")
)

// Errors returned when encountering error status codes.
var (
	// ErrBadRequest occurs when a request is malformed.
	ErrBadRequest = ServerError{
		Status: http.StatusBadRequest,
		Msg:    "bad request: check query parameters",
	}
	// ErrUnauthorized occurs when a request is made without authorization.
	ErrUnauthorized = ServerError{
		Status: http.StatusUnauthorized,
		Msg:    "authentication failed: check for valid API key in user-key header",
	}
	// ErrForbidden occurs when a request is made without authorization.
	ErrForbidden = ServerError{
		Status: http.StatusForbidden,
		Msg:    "authentication failed: check for valid API key in user-key header",
	}
	// ErrInternalError occurs when an unexpected IGDB server error occurs and should be reported.
	ErrInternalError = ServerError{
		Status: http.StatusInternalServerError,
		Msg:    "internal error: report bug",
	}
)

// ServerError contains information on an
// error returned from an IGDB API call.
type ServerError struct {
	Status int    `json:"status,omitempty"`
	Msg    string `json:"message,omitempty"`
}

// Error formats the ServerError and fulfills the error interface.
func (e ServerError) Error() string {
	return "igdb server error: status: " + strconv.Itoa(e.Status) + " message: " + e.Msg
}

// checkResponse checks the provided HTTP response
// for errors returned by the IGDB.
func checkResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusInternalServerError:
		return ErrInternalError
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(b))

	var e ServerError

	err = json.Unmarshal(b, &e)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal server error message")
	}

	return e
}

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
