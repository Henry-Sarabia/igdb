package igdb

import (
	"io/ioutil"
	"net/http"
	"strings"
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
				Body: ioutil.NopCloser(strings.NewReader(et.Body)),
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

func TestIsBracketPair(t *testing.T) {
	tests := []struct {
		name     string
		b        []byte
		wantBool bool
	}{
		{"Nil slice", nil, false},
		{"Empty byte slice", []byte{}, false},
		{"Single open bracket", []byte{91}, false},
		{"Single closed bracket", []byte{93}, false},
		{"Double open bracket", []byte{91, 91}, false},
		{"Double closed bracket", []byte{93, 93}, false},
		{"Triple open bracket", []byte{91, 91, 91}, false},
		{"Triple closed bracket", []byte{93, 93, 93}, false},
		{"Reversed bracket pair", []byte{93, 92}, false},
		{"Proper bracket pair with leading bracket", []byte{91, 91, 93}, false},
		{"Proper bracket pair with trailing bracket", []byte{91, 93, 93}, false},
		{"Double proper bracket pair", []byte{91, 93, 91, 93}, false},
		{"Proper bracket pair", []byte{91, 93}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isBracketPair(test.b)
			if got != test.wantBool {
				t.Errorf("got: <%v>, want: <%v>", got, test.wantBool)
			}
		})
	}
}
