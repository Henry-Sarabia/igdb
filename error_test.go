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
	var tests = []struct {
		name    string
		code    int
		body    string
		wantErr string
	}{
		{"Status OK", http.StatusOK, "", ""},
		{"Status Bad Request", http.StatusBadRequest, "", ErrBadRequest.Error()},
		{"Status Unauthorized", http.StatusUnauthorized, "", ErrAuthFailed.Error()},
		{"Status Forbidden", http.StatusForbidden, "", ErrAuthFailed.Error()},
		{"Status Internal Server Error", http.StatusInternalServerError, "", ErrInternalError.Error()},
		{"Unexpected Status Not Found", http.StatusNotFound, testErrNotFound, "Status 404 - status not found"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp := &http.Response{StatusCode: test.code,
				Body: ioutil.NopCloser(strings.NewReader(test.body)),
			}

			err := checkResponse(resp)
			if resp.StatusCode == http.StatusOK {
				if err != nil {
					t.Errorf("got: <%v>, want: <%v>", err, nil)
				}
				return
			}

			if err.Error() != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", err.Error(), test.wantErr)
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
