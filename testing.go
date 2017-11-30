package igdb

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

// testKey mocks an IGDB API key.
const testKey = "notarealkey"

// Errors returned when performing type validation.
var (
	// ErrNotStruct occurs when a non-struct type is provided to a function expecting a struct.
	ErrNotStruct = errors.New("igdb: not a struct")
	// ErrNotSlice occurs when a non-slice type is provided to a function expecting a slice.
	ErrNotSlice = errors.New("igdb: not a slice")
)

// testHeader mocks a single HTTP header entry with a key and value field.
type testHeader struct {
	Key   string
	Value string
}

// startTestServer initializes and returns a test server that will respond with the provided status,
// response, and optional headers. startTestServer also returns a Client configured specifically for
// the initialized test server.
func startTestServer(status int, resp io.Reader, headers ...testHeader) (*httptest.Server, *Client) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		for _, h := range headers {
			w.Header().Add(h.Key, h.Value)
		}
		io.Copy(w, resp)
	}))

	c := NewClient(testKey, nil)
	c.http = ts.Client()
	c.rootURL = ts.URL + "/"

	return ts, c
}

// testServerString initializes and returns a test server that will respond with the provided
// status, response, and optional headers. testServerString also returns a Client configured
// specifically for the initialized test server.
func testServerString(status int, resp string, headers ...testHeader) (*httptest.Server, *Client) {
	return startTestServer(status, strings.NewReader(resp), headers...)
}

// testServerFile initializes and returns a test server that will respond with the provided status,
// response read from the given filename, and optional headers. testServerFile also returns a Client
// configured specifically for the initialized test server.
func testServerFile(status int, filename string, headers ...testHeader) (*httptest.Server, *Client, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	ts, c := startTestServer(status, f, headers...)
	return ts, c, nil
}

// assertError checks if the provided error and expected
// error string signal the same error. If they do not, the
// given test will fail. If they do, the test continues as normal.
func assertError(t *testing.T, err error, expErr string) {
	if err != nil {
		if err.Error() != expErr {
			t.Fatalf("Expected error '%v', got '%v'", expErr, err.Error())
		}
	} else {
		if expErr != "" {
			t.Fatalf("Expected error '%v', got nil error", expErr)
		}
	}
	return
}

// equalSlice returns true if two slices contain
// the same elements, otherwise it returns false.
//
// Adapted from github.com/emou/testify/assert/assertions.go
func equalSlice(x, y interface{}) (bool, error) {
	if x == nil || y == nil {
		return x == y, nil
	}

	if reflect.TypeOf(x).Kind() != reflect.Slice {
		return false, ErrNotSlice
	}

	if reflect.TypeOf(y).Kind() != reflect.Slice {
		return false, ErrNotSlice
	}

	vx := reflect.ValueOf(x)
	vy := reflect.ValueOf(y)

	if vx.Len() != vy.Len() {
		return false, nil
	}

	visited := make([]bool, vy.Len())

	for i := 0; i < vx.Len(); i++ {
		element := vx.Index(i).Interface()

		found := false
		for j := 0; j < vy.Len(); j++ {
			if visited[j] {
				continue
			}

			if reflect.DeepEqual(vy.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return false, nil
		}
	}

	return true, nil
}
