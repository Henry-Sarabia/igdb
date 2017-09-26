package igdb

import "testing"

const testEndpoint endpoint = "tests/"

func TestSingleURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/1234?fields=id%2Cname%2Cpopularity&limit=5"
	aURL := c.singleURL(testEndpoint, 1234, OptLimit(5), OptFields("id", "name", "popularity"))
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}
