package igdb

import (
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// Mocked arguments for testing.
const (
	// testResult is a mocked response from a Get request.
	testResult = `{"some_field": "some_value"}`
)

// testResultPlaceHolder mocks an IGDB object.
type testResultPlaceholder struct {
	SomeField string `json:"some_field"`
}

func TestClient_Request(t *testing.T) {
	tests := []struct {
		name    string
		end     endpoint
		opts    []Option
		wantReq *http.Request
		wantErr error
	}{
		{"Zero options", testEndpoint, nil, httptest.NewRequest("GET", igdbURL+testEndpoint, nil), nil},
		{"Single option", testEndpoint, []Option{SetLimit(15)}, httptest.NewRequest("GET", igdbURL+testEndpoint, strings.NewReader("limit 15; ")), nil},
		{"Error option", testEndpoint, []Option{SetLimit(-99)}, httptest.NewRequest("GET", igdbURL+testEndpoint, nil), ErrOutOfRange},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewClient("somekey", nil)

			req, err := c.request(test.end, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			if req.Method != test.wantReq.Method {
				t.Errorf("got: <%v>, want: <%v>", req.Method, test.wantReq.Method)
			}

			if req.URL.String() != test.wantReq.URL.String() {
				t.Errorf("got: <%v>, want: <%v>", req.URL.String(), test.wantReq.URL.String())
			}

			if !reflect.DeepEqual(req.Body, test.wantReq.Body) {
				t.Errorf("got: <%v>, want: <%v>", req.Body, test.wantReq.Body)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	tests := []struct {
		name      string
		srvStatus int
		srvResp   string
		wantRes   testResultPlaceholder
		wantErr   error
	}{
		{"Status OK, populated response", http.StatusOK, testResult, testResultPlaceholder{SomeField: "some_value"}, nil},
		{"Status OK, empty array response", http.StatusOK, "[]", testResultPlaceholder{}, ErrNoResults},
		{"Status OK, empty response", http.StatusOK, "", testResultPlaceholder{}, errInvalidJSON},
		{"Status BadRequest, populated response", http.StatusBadRequest, testResult, testResultPlaceholder{}, ErrBadRequest},
		{"Status BadRequest, empty array response", http.StatusBadRequest, "[]", testResultPlaceholder{}, ErrBadRequest},
		{"Status BadRequest, empty response", http.StatusBadRequest, "", testResultPlaceholder{}, ErrBadRequest},
		{"Status Unauthorized, populated response", http.StatusUnauthorized, testResult, testResultPlaceholder{}, ErrAuthFailed},
		{"Status Unauthorized, empty array response", http.StatusUnauthorized, "[]", testResultPlaceholder{}, ErrAuthFailed},
		{"Status Unauthorized, empty response", http.StatusUnauthorized, "", testResultPlaceholder{}, ErrAuthFailed},
		{"Status Forbidden, populated response", http.StatusForbidden, testResult, testResultPlaceholder{}, ErrAuthFailed},
		{"Status Forbidden, empty array response", http.StatusForbidden, "[]", testResultPlaceholder{}, ErrAuthFailed},
		{"Status Forbidden, empty response", http.StatusForbidden, "", testResultPlaceholder{}, ErrAuthFailed},
		{"Status InternalServerError, populated response", http.StatusInternalServerError, testResult, testResultPlaceholder{}, ErrInternalError},
		{"Status InternalServerError, empty array response", http.StatusInternalServerError, "[]", testResultPlaceholder{}, ErrInternalError},
		{"Status InternalServerError, empty response", http.StatusInternalServerError, "", testResultPlaceholder{}, ErrInternalError},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(test.srvStatus, test.srvResp)
			defer ts.Close()

			req, err := http.NewRequest("GET", ts.URL, nil)
			if err != nil {
				t.Fatal(err)
			}

			res := testResultPlaceholder{}

			err = c.send(req, &res)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(res, test.wantRes) {
				t.Errorf("got: <%v>, want: <%v>", res, test.wantRes)
			}
		})
	}
}

//func TestGet(t *testing.T) {
//	var getTests = []struct {
//		Name    string
//		Status  int
//		Resp    string
//		URL     string
//		ExpResp string
//		ExpErr  string
//	}{
//		{"OK request with valid response", http.StatusOK, testResult, string(testEndpoint), "value", ""},
//		{"OK request with empty response", http.StatusOK, "", string(testEndpoint), "", errInvalidJSON.Error()},
//		{"Bad request with empty response", http.StatusNotFound, "", "badURL", "", errInvalidJSON.Error()},
//		{"Bad request with error response", http.StatusNotFound, testErrNotFound, "badURL", "", "Status 404 - status not found"},
//	}
//
//	for _, gt := range getTests {
//		testResp := testResultPlaceholder{}
//		t.Run(gt.Name, func(t *testing.T) {
//			ts, c := testServerString(gt.Status, gt.Resp)
//			defer ts.Close()
//
//			err := c.send(c.rootURL+gt.URL, &testResp)
//			//assertError(t, err, gt.ExpErr)
//
//			if testResp.Field != gt.ExpResp {
//				t.Fatalf("Expected response '%v', got '%v'", gt.ExpResp, testResp.Field)
//			}
//		})
//	}
//}
//
//func TestSingleURL(t *testing.T) {
//	var singleTests = []struct {
//		Name   string
//		ID     int
//		Opts   []Option
//		ExpURL string
//		ExpErr error
//	}{
//		{"Positive ID with no options", 1234, nil, igdbURL + string(testEndpoint) + "1234", nil},
//		{"Positive ID with limit and offset", 55, []Option{SetLimit(20), SetOffset(15)}, igdbURL + string(testEndpoint) + "55?limit=20&offset=15", nil},
//		{"Positive ID with fields and order", 100, []Option{SetFields("name", "rating"), SetOrder("rating", OrderDescending)}, igdbURL + string(testEndpoint) + "100?fields=name%2Crating&order=rating%3Adesc", nil},
//		{"Positive ID with filters", 55555, []Option{SetFilter("rating", OpGreaterThan, "80"), SetFilter("popularity", OpLessThan, "2")}, igdbURL + string(testEndpoint) + "55555?filter%5Bpopularity%5D%5Blt%5D=2&filter%5Brating%5D%5Bgt%5D=80", nil},
//		{"Negative ID with no options", -1234, nil, "", ErrNegativeID},
//		{"Negative ID with options", -55555, []Option{SetLimit(20), SetOffset(15)}, "", ErrNegativeID},
//		{"Positive ID with invalid option", 100, []Option{SetLimit(999)}, "", ErrOutOfRange},
//		{"Negative ID with invalid option", -100, []Option{SetLimit(999)}, "", ErrNegativeID},
//	}
//	for _, st := range singleTests {
//		t.Run(st.Name, func(t *testing.T) {
//			c := NewClient(testKey, nil)
//
//			url, err := c.singleURL(testEndpoint, st.ID, st.Opts...)
//			if !reflect.DeepEqual(err, st.ExpErr) {
//				t.Fatalf("Expected error '%v', got '%v'", st.ExpErr, err)
//			}
//
//			if url != st.ExpURL {
//				t.Fatalf("Expected URL '%s', got '%s'", st.ExpURL, url)
//			}
//		})
//	}
//}
//
//func TestMultiURL(t *testing.T) {
//	var multiTests = []struct {
//		Name   string
//		IDs    []int
//		Opts   []Option
//		ExpURL string
//		ExpErr error
//	}{
//		{"Positive ID with no options", []int{1234}, nil, igdbURL + string(testEndpoint) + "1234", nil},
//		{"Positive IDs with no options", []int{1, 2, 3, 4}, nil, igdbURL + string(testEndpoint) + "1,2,3,4", nil},
//		{"Positive IDs with limit and offset", []int{55, 110}, []Option{SetLimit(20), SetOffset(15)}, igdbURL + string(testEndpoint) + "55,110?limit=20&offset=15", nil},
//		{"Positive IDs with fields and order", []int{100}, []Option{SetFields("name", "rating"), SetOrder("rating", OrderDescending)}, igdbURL + string(testEndpoint) + "100?fields=name%2Crating&order=rating%3Adesc", nil},
//		{"Positive IDs with filters", []int{55555}, []Option{SetFilter("rating", OpGreaterThan, "80"), SetFilter("popularity", OpLessThan, "2")}, igdbURL + string(testEndpoint) + "55555?filter%5Bpopularity%5D%5Blt%5D=2&filter%5Brating%5D%5Bgt%5D=80", nil},
//		{"Negative ID with no options", []int{-1234}, nil, "", ErrNegativeID},
//		{"Negative IDs with no options", []int{-1, -2, -3, -4}, nil, "", ErrNegativeID},
//		{"Negative IDs with options", []int{-55555}, []Option{SetLimit(20), SetOffset(15)}, "", ErrNegativeID},
//		{"Mixed IDs with options", []int{100, -200, 300, -400}, []Option{SetLimit(5), SetOffset(20)}, "", ErrNegativeID},
//		{"Mixed IDs with no options", []int{-1, 2, -3, 4}, nil, "", ErrNegativeID},
//		{"No IDs with no options", nil, nil, igdbURL + string(testEndpoint), nil},
//		{"No IDs with options", nil, []Option{SetLimit(20), SetOffset(15)}, igdbURL + string(testEndpoint) + "?limit=20&offset=15", nil},
//		{"Positive IDs with invalid option", []int{100, 200}, []Option{SetLimit(999)}, "", ErrOutOfRange},
//		{"Negative IDs with invalid option", []int{-100, -200}, []Option{SetLimit(999)}, "", ErrNegativeID},
//		{"Mixed IDs with invalid option", []int{100, -200}, []Option{(SetLimit(999))}, "", ErrNegativeID},
//	}
//	for _, mt := range multiTests {
//		t.Run(mt.Name, func(t *testing.T) {
//			c := NewClient(testKey, nil)
//
//			url, err := c.multiURL(testEndpoint, mt.IDs, mt.Opts...)
//			if !reflect.DeepEqual(err, mt.ExpErr) {
//				t.Fatalf("Expected error '%v', got '%v'", mt.ExpErr, err)
//			}
//
//			if url != mt.ExpURL {
//				t.Fatalf("Expected URL '%s', got '%s'", mt.ExpURL, url)
//			}
//		})
//	}
//}
//
//func TestSearchURL(t *testing.T) {
//	var searchTests = []struct {
//		Name   string
//		Query  string
//		Opts   []Option
//		ExpURL string
//		ExpErr error
//	}{
//		{"Non-empty query with no options", "zelda", nil, igdbURL + string(testEndpoint) + "?search=zelda", nil},
//		{"Non-empty query with limit and offset", "zelda", []Option{SetLimit(20), SetOffset(15)}, igdbURL + string(testEndpoint) + "?limit=20&offset=15&search=zelda", nil},
//		{"Non-empty query with fields and order", "zelda", []Option{SetFields("name", "rating"), SetOrder("rating", OrderDescending)}, igdbURL + string(testEndpoint) + "?fields=name%2Crating&order=rating%3Adesc&search=zelda", nil},
//		{"Non-empty query with filters", "zelda", []Option{SetFilter("rating", OpGreaterThan, "80"), SetFilter("popularity", OpLessThan, "2")}, igdbURL + string(testEndpoint) + "?filter%5Bpopularity%5D%5Blt%5D=2&filter%5Brating%5D%5Bgt%5D=80&search=zelda", nil},
//		{"Empty query with no options", "", nil, "", ErrEmptyQuery},
//		{"Empty query with options", "", []Option{SetLimit(50), SetFilter("platforms", OpAny, "9")}, "", ErrEmptyQuery},
//		{"Space query with no options", "   ", nil, "", ErrEmptyQuery},
//		{"Space query with options", "   ", []Option{SetLimit(50), SetFilter("platforms", OpAny, "9")}, "", ErrEmptyQuery},
//		{"Non-empty query with invalid option", "zelda", []Option{SetOffset(-999)}, "", ErrOutOfRange},
//		{"Empty query with invalid option", "", []Option{SetOffset(-999)}, "", ErrEmptyQuery},
//		{"Space query with invalid option", "   ", []Option{SetOffset(-999)}, "", ErrEmptyQuery},
//	}
//
//	for _, st := range searchTests {
//		t.Run(st.Name, func(t *testing.T) {
//			c := NewClient(testKey, nil)
//
//			url, err := c.searchURL(testEndpoint, st.Query, st.Opts...)
//			if !reflect.DeepEqual(err, st.ExpErr) {
//				t.Fatalf("Expected error '%v', got '%v'", st.ExpErr, err)
//			}
//
//			if url != st.ExpURL {
//				t.Fatalf("Expected URL '%s', got '%s'", st.ExpURL, url)
//			}
//		})
//	}
//}
