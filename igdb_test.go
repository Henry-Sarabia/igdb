package igdb

import (
	"strconv"
	"testing"
)

const testEndpoint endpoint = "tests/"

func TestSingleURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/1234?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc"
	aURL := c.singleURL(testEndpoint, 1234,
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", GTE, strconv.Itoa(50)),
		OptLimit(10),
		OptOffset(5),
		OptOrder("popularity", Desc))
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}

func TestMultiURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/1234,5678,9?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc"
	aURL := c.multiURL(testEndpoint, []int{1234, 5678, 9},
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", GTE, strconv.Itoa(50)),
		OptLimit(10),
		OptOffset(5),
		OptOrder("popularity", Desc))
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}

func TestSearchURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc&search=mario+party"
	aURL := c.searchURL(testEndpoint, "mario party",
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", GTE, strconv.Itoa(50)),
		OptLimit(10),
		OptOffset(5),
		OptOrder("popularity", Desc))
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}
