package igdb

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
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
		wantErr error
	}{
		{"Status OK", http.StatusOK, "", nil},
		{"Status Bad Request", http.StatusBadRequest, "", ErrBadRequest},
		{"Status Unauthorized", http.StatusUnauthorized, "", ErrUnauthorized},
		{"Status Forbidden", http.StatusForbidden, "", ErrForbidden},
		{"Status Internal Server Error", http.StatusInternalServerError, "", ErrInternalError},
		{"Unexpected Status Not Found", http.StatusNotFound, testErrNotFound, ServerError{Status: 404, Msg: "status not found"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp := &http.Response{StatusCode: test.code,
				Body: ioutil.NopCloser(strings.NewReader(test.body)),
			}

			err := checkResponse(resp)

			if !reflect.DeepEqual(errors.Cause(err), test.wantErr) {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
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
