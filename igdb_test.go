package igdb

import (
	"net/http"
	"strconv"
	"testing"
)

const testEndpoint endpoint = "tests/"
const testGetResp = `[{"example": "value"}]`

type testStruct struct {
	Example string `json:"example"`
}

func TestGet(t *testing.T) {
	testVar := testStruct{}
	var getTests = []struct {
		Name   string
		URL    string
		Result interface{}
		Code   int
		Resp   string
		ExpErr string
	}{
		{"bad request with bad response", "fakeurl", testVar, http.StatusBadRequest, "", "unexpected end of JSON input"},
		{"OK request with OK response", string(testEndpoint), testVar, http.StatusOK, testGetResp, ""},
		{"OK request with bad response", string(testEndpoint), testVar, http.StatusOK, "", "unexpected end of JSON input"},
	}

	for _, gt := range getTests {
		t.Run(gt.Name, func(t *testing.T) {
			ts, c := startTestServer(gt.Code, gt.Resp)
			defer ts.Close()

			err := c.get(c.rootURL+gt.URL, &gt.Result)
			if err == nil {
				if gt.ExpErr != "" {
					t.Fatalf("Expected error '%v', got nil error'", gt.ExpErr)
				}
				return
			}
			if err.Error() != gt.ExpErr {
				t.Fatalf("Expected error '%v', got error '%v'", gt.ExpErr, err.Error())
			}
		})
	}
}

func TestSingleURL(t *testing.T) {
	c := NewClient()

	eURL := "https://api-2445582011268.apicast.io/tests/1234?fields=id%2Cname%2Cpopularity&filter%5Bpopularity%5D%5Bgte%5D=50&limit=10&offset=5&order=popularity%3Adesc"
	aURL, err := c.singleURL(testEndpoint, 1234,
		OptFields("id", "name", "popularity"),
		OptFilter("popularity", GreaterThanEqual, strconv.Itoa(50)),
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
		OptFilter("popularity", GreaterThanEqual, strconv.Itoa(50)),
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
		OptFilter("popularity", GreaterThanEqual, strconv.Itoa(50)),
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
