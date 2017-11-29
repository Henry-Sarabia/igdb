package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

// Mocked arguments for testing.
const (
	testEndpoint endpoint = "tests/"
	// testGetResp is a mocked response from a Get request.
	testGetResp = `{"field": "value"}`
)

// testResultPlaceHolder mocks an IGDB object.
type testResultPlaceholder struct {
	Field string `json:"field"`
}

func TestGet(t *testing.T) {
	var getTests = []struct {
		Name    string
		Status  int
		Resp    string
		URL     string
		ExpResp string
		ExpErr  string
	}{
		{"OK request with valid response", http.StatusOK, testGetResp, string(testEndpoint), "value", ""},
		{"OK request with empty response", http.StatusOK, "", string(testEndpoint), "", errEndOfJSON.Error()},
		{"Bad request with empty response", http.StatusNotFound, "", "badURL", "", errEndOfJSON.Error()},
		{"Bad request with error response", http.StatusNotFound, testErrNotFound, "badURL", "", "Status 404 - status not found"},
	}

	for _, gt := range getTests {
		testResp := testResultPlaceholder{}
		t.Run(gt.Name, func(t *testing.T) {
			ts, c := testServerString(gt.Status, gt.Resp)
			defer ts.Close()

			err := c.get(c.rootURL+gt.URL, &testResp)
			assertError(t, err, gt.ExpErr)

			if testResp.Field != gt.ExpResp {
				t.Fatalf("Expected response '%v', got '%v'", gt.ExpResp, testResp.Field)
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
			c := NewClient(testKey, nil)

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
			c := NewClient(testKey, nil)

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
	var searchTests = []struct {
		Name   string
		Query  string
		Opts   []OptionFunc
		ExpURL string
		ExpErr error
	}{
		{"Non-empty query with no options", "zelda", nil, igdbURL + string(testEndpoint) + "?search=zelda", nil},
		{"Non-empty query with limit and offset", "zelda", []OptionFunc{OptLimit(20), OptOffset(15)}, igdbURL + string(testEndpoint) + "?limit=20&offset=15&search=zelda", nil},
		{"Non-empty query with fields and order", "zelda", []OptionFunc{OptFields("name", "rating"), OptOrder("rating", OrderDescending)}, igdbURL + string(testEndpoint) + "?fields=name%2Crating&order=rating%3Adesc&search=zelda", nil},
		{"Non-empty query with filters", "zelda", []OptionFunc{OptFilter("rating", OpGreaterThan, "80"), OptFilter("popularity", OpLessThan, "2")}, igdbURL + string(testEndpoint) + "?filter%5Bpopularity%5D%5Blt%5D=2&filter%5Brating%5D%5Bgt%5D=80&search=zelda", nil},
		{"Empty query with no options", "", nil, "", ErrEmptyQuery},
		{"Empty query with options", "", []OptionFunc{OptLimit(50), OptFilter("platforms", OpAny, "9")}, "", ErrEmptyQuery},
		{"Space query with no options", "   ", nil, "", ErrEmptyQuery},
		{"Space query with options", "   ", []OptionFunc{OptLimit(50), OptFilter("platforms", OpAny, "9")}, "", ErrEmptyQuery},
		{"Non-empty query with invalid option", "zelda", []OptionFunc{OptOffset(-999)}, "", ErrOutOfRange},
		{"Empty query with invalid option", "", []OptionFunc{OptOffset(-999)}, "", ErrEmptyQuery},
		{"Space query with invalid option", "   ", []OptionFunc{OptOffset(-999)}, "", ErrEmptyQuery},
	}

	for _, st := range searchTests {
		t.Run(st.Name, func(t *testing.T) {
			c := NewClient(testKey, nil)

			url, err := c.searchURL(testEndpoint, st.Query, st.Opts...)
			if !reflect.DeepEqual(err, st.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", st.ExpErr, err)
			}

			if url != st.ExpURL {
				t.Fatalf("Expected URL '%s', got '%s'", st.ExpURL, url)
			}
		})
	}
}

func TestIntsToStrings(t *testing.T) {
	var tableTests = []struct {
		Name       string
		Ints       []int
		ExpStrings []string
	}{
		{"Empty slice", nil, nil},
		{"Zero int slice", []int{0}, []string{"0"}},
		{"Single positive int slice", []int{100}, []string{"100"}},
		{"Single negative int slice", []int{-100}, []string{"-100"}},
		{"Multiple positive ints slice", []int{100, 5, 999, 123456789}, []string{"100", "5", "999", "123456789"}},
		{"Multiple negative ints slice", []int{-100, -5, -999, -123456789}, []string{"-100", "-5", "-999", "-123456789"}},
		{"Mixed ints slice", []int{100, -200, 300, -400}, []string{"100", "-200", "300", "-400"}},
	}

	for _, tt := range tableTests {
		t.Run(tt.Name, func(t *testing.T) {
			s := intsToStrings(tt.Ints)

			ok, err := equalSlice(s, tt.ExpStrings)
			if err != nil {
				t.Error(err)
			}
			if !ok {
				t.Fatalf("Expected strings %v, got %v", tt.ExpStrings, s)
			}
		})
	}
}

func TestIntsToCommaString(t *testing.T) {
	var tableTests = []struct {
		Name      string
		Ints      []int
		Expstring string
	}{
		{"Empty slice", nil, ""},
		{"Zero int slice", []int{0}, "0"},
		{"Single positive int slice", []int{100}, "100"},
		{"Single negative int slice", []int{-100}, "-100"},
		{"Multiple positive ints slice", []int{100, 5, 999, 123456789}, "100,5,999,123456789"},
		{"Multiple negative ints slice", []int{-100, -5, -999, -123456789}, "-100,-5,-999,-123456789"},
		{"Mixed ints slice", []int{100, -200, 300, -400}, "100,-200,300,-400"},
	}

	for _, tt := range tableTests {
		t.Run(tt.Name, func(t *testing.T) {
			s := intsToCommaString(tt.Ints)

			if s != tt.Expstring {
				t.Fatalf("Expected string '%s', got '%s'", tt.Expstring, s)
			}
		})
	}
}
