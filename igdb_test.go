package igdb

import (
	"net/http"
	"strconv"
	"testing"
)

const testEndpoint endpoint = "tests/"
const testGetResp = `{"example": "value"}`
const testNextHeader = "/games/scroll/DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAECqQWdzRScEJ3YkVUVldrb0pycUtGR2l4QQ==/?fields=name,id,rating,popularity"

type testResultPlaceholder struct {
	Example string `json:"example"`
}

func TestGet(t *testing.T) {
	var getTests = []struct {
		Name    string
		URL     string
		Code    int
		ExpResp string
		ExpErr  string
	}{
		{"Bad request with bad response", "fakeurl", http.StatusBadRequest, "", errEndOfJSON.Error()},
		{"OK request with OK response", string(testEndpoint), http.StatusOK, testGetResp, ""},
		{"OK request with bad response", string(testEndpoint), http.StatusOK, "", errEndOfJSON.Error()},
	}

	for _, gt := range getTests {
		testResp := testResultPlaceholder{}
		t.Run(gt.Name, func(t *testing.T) {
			ts, c := startTestServer(gt.Code, gt.ExpResp)
			defer ts.Close()

			err := c.get(c.rootURL+gt.URL, &testResp)
			if err == nil {
				if gt.ExpErr != "" {
					t.Fatalf("Expected error '%v', got nil error'", gt.ExpErr)
				}
				return
			} else if err.Error() != gt.ExpErr {
				t.Fatalf("Expected error '%v', got error '%v'", gt.ExpErr, err.Error())
			}

			if testResp.Example != gt.ExpResp {
				t.Fatalf("Expected response '%v', got '%v'", gt.ExpResp, testResp.Example)
			}
		})
	}
}

func TestSetScrollHeaders(t *testing.T) {
	var headerTests = []struct {
		Name     string
		Headers  []testHeader
		ExpNext  string
		ExpCount int
		ExpErr   string
	}{
		{"Non-empty Next and non-empty Count", []testHeader{{"X-Next-Page", testNextHeader}, {"X-Count", "105"}}, testNextHeader, 105, ""},
		{"Empty Next and non-empty Count", []testHeader{{"X-Count", "31"}}, "", 31, ""},
		{"Non-empty Next and empty Count", []testHeader{{"X-Next-Page", testNextHeader}}, testNextHeader, 0, ""},
		{"Empty Next and empty Count", []testHeader{}, "", 0, ""},
		{"Non-empty Next and invalid Count", []testHeader{{"X-Next-Page", testNextHeader}, {"X-Count", "abcd"}}, testNextHeader, 0, `strconv.Atoi: parsing "abcd": invalid syntax`},
	}

	for _, ht := range headerTests {
		t.Run(ht.Name, func(t *testing.T) {
			hdr := http.Header{}
			for _, h := range ht.Headers {
				hdr.Set(h.Key, h.Value)
			}

			c := NewClient()
			err := c.setScrollHeaders(hdr)
			if err != nil {
				if err.Error() != ht.ExpErr {
					t.Fatalf("Expected error '%v', got '%v'", ht.ExpErr, err.Error())
				}
			} else {
				if ht.ExpErr != "" {
					t.Fatalf("Expected error '%v', got nil error", ht.ExpErr)
				}
			}

			if c.ScrollNext != ht.ExpNext {
				t.Errorf("Expected ScrollNext of '%s', got '%s'", ht.ExpNext, c.ScrollNext)
			}

			if c.ScrollCount != ht.ExpCount {
				t.Errorf("Expected ScrollCount of %d, got %d", ht.ExpCount, c.ScrollCount)
			}
		})
	}
}

func TestSingleURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/1234?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc"
	aURL, err := c.singleURL(testEndpoint, 1234,
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", OpGreaterThanEqual, strconv.Itoa(50)),
		OptLimit(10),
		OptOffset(5),
		OptOrder("popularity", OrderDescending))
	if err != nil {
		t.Error(err)
	}
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}

func TestMultiURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/1234,5678,9?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc"
	aURL, err := c.multiURL(testEndpoint, []int{1234, 5678, 9},
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", OpGreaterThanEqual, strconv.Itoa(50)),
		OptLimit(10),
		OptOffset(5),
		OptOrder("popularity", OrderDescending))
	if err != nil {
		t.Error(err)
	}
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}

func TestSearchURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc&search=mario+party"
	aURL, err := c.searchURL(testEndpoint, "mario party",
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", OpGreaterThanEqual, strconv.Itoa(50)),
		OptLimit(10),
		OptOffset(5),
		OptOrder("popularity", OrderDescending))
	if err != nil {
		t.Error(err)
	}
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}
