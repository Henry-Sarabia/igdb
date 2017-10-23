package igdb

import (
	"net/http"
	"reflect"
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
	var singleTests = []struct {
		Name   string
		ID     int
		Opts   []OptionFunc
		ExpURL string
		ExpErr error
	}{
		{"Positive ID with no options", 1234, nil, igdbURL + string(testEndpoint) + "1234", nil},
		{"Positive ID with limit and offset", 55, []OptionFunc{OptLimit(20), OptOffset(15)}, igdbURL + string(testEndpoint) + "55?limit=20&offset=15", nil},
		{"Positive ID with fields and order", 100, []OptionFunc{OptFields("name", "rating"), OptOrder("rating", OrderDescending)}, igdbURL + string(testEndpoint) + "100?fields=name%2Crating&order=rating%3Adesc", nil},
		{"Positive ID with filters", 55555, []OptionFunc{OptFilter("rating", OpGreaterThan, "80"), OptFilter("popularity", OpLessThan, "2")}, igdbURL + string(testEndpoint) + "55555?filter%5Bpopularity%5D%5Blt%5D=2&filter%5Brating%5D%5Bgt%5D=80", nil},
		{"Negative ID with no options", -1234, nil, "", ErrNegativeID},
		{"Negative ID with options", -55555, []OptionFunc{OptLimit(20), OptOffset(15)}, "", ErrNegativeID},
		{"Positive ID with invalid option", 100, []OptionFunc{OptLimit(999)}, "", ErrOutOfRange},
		{"Negative ID with invalid option", -100, []OptionFunc{OptLimit(999)}, "", ErrNegativeID},
	}
	for _, st := range singleTests {
		t.Run(st.Name, func(t *testing.T) {
			c := NewClient()

			url, err := c.singleURL(testEndpoint, st.ID, st.Opts...)
			if !reflect.DeepEqual(err, st.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", st.ExpErr, err)
			}

			if url != st.ExpURL {
				t.Fatalf("Expected URL '%s', got '%s'", st.ExpURL, url)
			}
		})
	}
}

func TestMultiURL(t *testing.T) {
	var multiTests = []struct {
		Name   string
		IDs    []int
		Opts   []OptionFunc
		ExpURL string
		ExpErr error
	}{
		{"Positive ID with no options", []int{1234}, nil, igdbURL + string(testEndpoint) + "1234", nil},
		{"Positive IDs with no options", []int{1, 2, 3, 4}, nil, igdbURL + string(testEndpoint) + "1,2,3,4", nil},
		{"Positive IDs with limit and offset", []int{55, 110}, []OptionFunc{OptLimit(20), OptOffset(15)}, igdbURL + string(testEndpoint) + "55,110?limit=20&offset=15", nil},
		{"Positive IDs with fields and order", []int{100}, []OptionFunc{OptFields("name", "rating"), OptOrder("rating", OrderDescending)}, igdbURL + string(testEndpoint) + "100?fields=name%2Crating&order=rating%3Adesc", nil},
		{"Positive IDs with filters", []int{55555}, []OptionFunc{OptFilter("rating", OpGreaterThan, "80"), OptFilter("popularity", OpLessThan, "2")}, igdbURL + string(testEndpoint) + "55555?filter%5Bpopularity%5D%5Blt%5D=2&filter%5Brating%5D%5Bgt%5D=80", nil},
		{"Negative ID with no options", []int{-1234}, nil, "", ErrNegativeID},
		{"Negative IDs with no options", []int{-1, -2, -3, -4}, nil, "", ErrNegativeID},
		{"Negative IDs with options", []int{-55555}, []OptionFunc{OptLimit(20), OptOffset(15)}, "", ErrNegativeID},
		{"Mixed IDs with options", []int{100, -200, 300, -400}, []OptionFunc{OptLimit(5), OptOffset(20)}, "", ErrNegativeID},
		{"Mixed IDs with no options", []int{-1, 2, -3, 4}, nil, "", ErrNegativeID},
		{"No IDs with no options", nil, nil, igdbURL + string(testEndpoint), nil},
		{"No IDs with options", nil, []OptionFunc{OptLimit(20), OptOffset(15)}, igdbURL + string(testEndpoint) + "?limit=20&offset=15", nil},
		{"Positive IDs with invalid option", []int{100, 200}, []OptionFunc{OptLimit(999)}, "", ErrOutOfRange},
		{"Negative IDs with invalid option", []int{-100, -200}, []OptionFunc{OptLimit(999)}, "", ErrNegativeID},
		{"Mixed IDs with invalid option", []int{100, -200}, []OptionFunc{(OptLimit(999))}, "", ErrNegativeID},
	}
	for _, mt := range multiTests {
		t.Run(mt.Name, func(t *testing.T) {
			c := NewClient()

			url, err := c.multiURL(testEndpoint, mt.IDs, mt.Opts...)
			if !reflect.DeepEqual(err, mt.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", mt.ExpErr, err)
			}

			if url != mt.ExpURL {
				t.Fatalf("Expected URL '%s', got '%s'", mt.ExpURL, url)
			}
		})
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
