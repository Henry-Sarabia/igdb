package igdb

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

const testErrNotFound = `
{
	"status": 404,
	"message": "status not found"
}
`

func TestCheckResponse(t *testing.T) {
	var errTests = []struct {
		Name string
		Code int
		Body string
		Exp  string
	}{
		{"Status OK", http.StatusOK, "", ""},
		{"Status Bad Request", http.StatusBadRequest, "", ErrBadRequest.Error()},
		{"Status Unauthorized", http.StatusUnauthorized, "", ErrAuthFailed.Error()},
		{"Status Forbidden", http.StatusForbidden, "", ErrAuthFailed.Error()},
		{"Status Internal Server Error", http.StatusInternalServerError, "", ErrInternalError.Error()},
		{"Unexpected Status Not Found", http.StatusNotFound, testErrNotFound, "Status 404 - status not found"},
	}

	for _, et := range errTests {
		t.Run(et.Name, func(t *testing.T) {
			resp := &http.Response{StatusCode: et.Code,
				Body: ioutil.NopCloser(bytes.NewBufferString(et.Body)),
			}

			err := checkResponse(resp)
			if resp.StatusCode == http.StatusOK {
				if err != nil {
					t.Fatalf("Expected nil err, got '%v'", err)
				}
				return
			}
			if err.Error() != et.Exp {
				t.Fatalf("Expected '%v', got '%v'", et.Exp, err.Error())
			}
		})
	}
}
