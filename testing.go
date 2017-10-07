package igdb

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
)

// startTestServer initializes a test server that will respond with the
// given status and response. A Client configured especially for this
// test server is also returned.
func startTestServer(status int, resp string) (*httptest.Server, Client) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprint(w, resp)
	}))
	c := Client{
		http:    ts.Client(),
		rootURL: ts.URL + "/",
	}

	return ts, c
}

// validateStruct checks if the given struct contains all of the fields
// it should according to the appropriate IGDB endpoint.
func (c *Client) validateStruct(str reflect.Type, end endpoint) error {
	f, err := c.GetEndpointFields(end)
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
// the struct tags according to the slice of strings representing
// the appropriate struct tags.
func validateStructTags(str reflect.Type, new []string) error {
	old, err := getStructTags(str)
	if err != nil {
		return err
	}

	found := make(map[string]bool)
	for _, v := range new {
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

// getStructTags collects the struct tags of every available
// field in the given struct.
func getStructTags(str reflect.Type) ([]string, error) {
	if str.Kind() != reflect.Struct {
		return nil, errors.New("input type's kind not a struct")
	}

	var f []string
	for i := 0; i < str.NumField(); i++ {
		f = append(f, str.Field(i).Tag.Get("json"))
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
// adapted from github.com/emou/testify/assert/assertions.go
func equalSlice(x, y interface{}) (bool, error) {
	if x == nil || y == nil {
		return x == y, nil
	}

	if reflect.TypeOf(x).Kind() != reflect.Slice {
		return false, errors.New("not a slice")
	}

	if reflect.TypeOf(y).Kind() != reflect.Slice {
		return false, errors.New("not a slice")
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
