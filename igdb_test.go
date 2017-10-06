package igdb

import (
	"net/http"
	"strconv"
	"testing"
)

const testEndpoint endpoint = "tests/"
const testGetResp = `{"example": "value"}`

type testStruct struct {
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
		{"Bad request with bad response", "fakeurl", http.StatusBadRequest, "", "unexpected end of JSON input"},
		{"OK request with OK response", string(testEndpoint), http.StatusOK, testGetResp, ""},
		{"OK request with bad response", string(testEndpoint), http.StatusOK, "", "unexpected end of JSON input"},
	}

	for _, gt := range getTests {
		testResp := testStruct{}
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

func TestSingleURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/1234?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc"
	aURL, err := c.singleURL(testEndpoint, 1234,
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", OpGreaterThanEqual, strconv.Itoa(50)),
		OptLimit(10),
		OptOffset(5),
		OptOrder("popularity", DescendingOrder))
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
		OptOrder("popularity", DescendingOrder))
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
		OptOrder("popularity", DescendingOrder))
	if err != nil {
		t.Error(err)
	}
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}
