package igdb

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
)

// Errors returned when performing struct type validation.
var (
	// ErrNotStruct occurs when a non-struct type is provided to a function expecting a struct.
	ErrNotStruct = errors.New("igdb: not a struct")
	// ErrNotSlice occurs when a non-slice type is provided to a function expecting a slice.
	ErrNotSlice = errors.New("igdb: not a slice")
)

// testHeader mocks a single HTTP header
// entry with a key and value field.
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

	c := NewClient()
	c.http = ts.Client()
	c.rootURL = ts.URL + "/"

	return ts, &c
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

// validateStruct checks if the given struct contains all of the fields
// listed by the field manifest provided by the given IGDB endpoint.
func (c *Client) validateStruct(str reflect.Type, end endpoint) error {
	f, err := c.GetEndpointFieldManifest(end)
	if err != nil {
		return err
	}

	f = removeSubfields(f)

	err = validateStructTags(str, f)
	if err != nil {
		return err
	}
	return nil
}

// validateStructTags checks if the given struct contains all of
// the struct tags in the provided field manifest.
func validateStructTags(str reflect.Type, fm []string) error {
	old, err := getStructTags(str)
	if err != nil {
		return err
	}

	found := make(map[string]bool)
	for _, v := range fm {
		found[v] = false
	}

	for _, v := range old {
		if _, ok := found[v]; ok {
			found[v] = true
		}
	}

	var missing []string
	for k, v := range found {
		if !v {
			missing = append(missing, k)
		}
	}

	if missing != nil {
		return errors.New("missing struct tags: " + strings.Join(missing, ", "))
	}

	return nil
}

// getStructTags collects the struct tags of every
// available field in the given struct.
func getStructTags(str reflect.Type) ([]string, error) {
	if str.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	var f []string
	for i := 0; i < str.NumField(); i++ {
		tag := str.Field(i).Tag.Get("json")
		if tag != "" {
			f = append(f, tag)
		}
	}
	return f, nil
}

// removeSubfields returns a slice of strings
// containing all fields from f except for any
// subfields, identified by their inclusion of
// the period character. Additionally, excess
// space characters and empty strings are removed.
func removeSubfields(f []string) []string {
	var out []string
	for _, val := range f {
		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}
		if !strings.Contains(val, ".") {
			out = append(out, val)
		}
	}
	return out
}

// equalSlice returns true if two slices contain
// the same elements, otherwise it returns false.
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
