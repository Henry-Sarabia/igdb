package igdb

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"strings"
)

const (
	// testClientID mocks an IGDB API key.
	testClientID = "notarealclientid"
	testToken    = "notarealtoken"
	// testFileEmpty is an empty file used for testing input.
	testFileEmpty string = "test_data/empty.json"
	// testFileEmptyArray is an empty array file used for testing input.
	testFileEmptyArray string = "test_data/empty_array.json"
	// testEndpoint mocks an endpoint.
	testEndpoint = "test/"
	// testResult is a mocked response from a Get request.
	testResult = `{"some_field": "some_value"}`
)

// testResultPlaceHolder mocks an IGDB object.
type testResultPlaceholder struct {
	SomeField string `json:"some_field"`
}

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

	c := NewClient(testClientID, testToken, ts.Client())
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

// equalSlice returns true if two slices contain
// the same elements, otherwise it returns false.
// The slices will be sorted.
func equalSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	sort.Strings(x)
	sort.Strings(y)

	return reflect.DeepEqual(x, y)
}
