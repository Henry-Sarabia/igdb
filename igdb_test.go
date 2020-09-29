package igdb

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestClient_Request(t *testing.T) {
	tests := []struct {
		name    string
		end     endpoint
		opts    []Option
		wantReq *http.Request
		wantErr error
	}{
		{"Zero options", testEndpoint, nil, httptest.NewRequest("POST", igdbURL+testEndpoint, nil), nil},
		{"Single option", testEndpoint, []Option{SetLimit(15)}, httptest.NewRequest("POST", igdbURL+testEndpoint, strings.NewReader("limit 15; ")), nil},
		{"Error option", testEndpoint, []Option{SetLimit(-99)}, httptest.NewRequest("POST", igdbURL+testEndpoint, nil), ErrOutOfRange},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

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
		{"Status Unauthorized, populated response", http.StatusUnauthorized, testResult, testResultPlaceholder{}, ErrUnauthorized},
		{"Status Unauthorized, empty array response", http.StatusUnauthorized, "[]", testResultPlaceholder{}, ErrUnauthorized},
		{"Status Unauthorized, empty response", http.StatusUnauthorized, "", testResultPlaceholder{}, ErrUnauthorized},
		{"Status Forbidden, populated response", http.StatusForbidden, testResult, testResultPlaceholder{}, ErrForbidden},
		{"Status Forbidden, empty array response", http.StatusForbidden, "[]", testResultPlaceholder{}, ErrForbidden},
		{"Status Forbidden, empty response", http.StatusForbidden, "", testResultPlaceholder{}, ErrForbidden},
		{"Status InternalServerError, populated response", http.StatusInternalServerError, testResult, testResultPlaceholder{}, ErrInternalError},
		{"Status InternalServerError, empty array response", http.StatusInternalServerError, "[]", testResultPlaceholder{}, ErrInternalError},
		{"Status InternalServerError, empty response", http.StatusInternalServerError, "", testResultPlaceholder{}, ErrInternalError},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(test.srvStatus, test.srvResp)
			defer ts.Close()

			req, err := http.NewRequest("POST", ts.URL, nil)
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

func TestClient_Get(t *testing.T) {
	tests := []struct {
		name      string
		srvStatus int
		srvResp   string
		opts      []Option
		wantRes   testResultPlaceholder
		wantErr   error
	}{
		{
			"Status OK, populated response, no options",
			http.StatusOK,
			testResult,
			[]Option{},
			testResultPlaceholder{SomeField: "some_value"},
			nil,
		},
		{
			"Status OK, populated response, single valid option",
			http.StatusOK,
			testResult,
			[]Option{SetLimit(15)},
			testResultPlaceholder{SomeField: "some_value"},
			nil,
		},
		{
			"Status OK, populated response, multiple valid options",
			http.StatusOK,
			testResult,
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{SomeField: "some_value"},
			nil,
		},
		{
			"Status OK, populated response, single invalid option",
			http.StatusOK,
			testResult,
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status OK, empty array response, no options",
			http.StatusOK,
			"[]",
			[]Option{},
			testResultPlaceholder{},
			ErrNoResults,
		},
		{
			"Status OK, empty array response, single valid option",
			http.StatusOK,
			"[]",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrNoResults,
		},
		{
			"Status OK, empty array response, multiple valid options",
			http.StatusOK,
			"[]",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrNoResults,
		},
		{
			"Status OK, empty array response, single invalid option",
			http.StatusOK,
			"[]",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status OK, empty response, no options",
			http.StatusOK,
			"",
			[]Option{},
			testResultPlaceholder{},
			errInvalidJSON,
		},
		{
			"Status OK, empty response, single valid option",
			http.StatusOK,
			"",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			errInvalidJSON,
		},
		{
			"Status OK, empty response, multiple valid options",
			http.StatusOK,
			"",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			errInvalidJSON,
		},
		{
			"Status OK, empty response, single invalid option",
			http.StatusOK,
			"",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status BadRequest, populated response, no options",
			http.StatusBadRequest,
			testResult,
			[]Option{},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, populated response, single valid option",
			http.StatusBadRequest,
			testResult,
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, populated response, multiple valid options",
			http.StatusBadRequest,
			testResult,
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, populated response, single invalid option",
			http.StatusBadRequest,
			testResult,
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status BadRequest, empty array response, no options",
			http.StatusBadRequest,
			"[]",
			[]Option{},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, empty array response, single valid option",
			http.StatusBadRequest,
			"[]",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, empty array response, multiple valid options",
			http.StatusBadRequest,
			"[]",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, empty array response, single invalid option",
			http.StatusBadRequest,
			"[]",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status BadRequest, empty response, no options",
			http.StatusBadRequest,
			"",
			[]Option{},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, empty response, single valid option",
			http.StatusBadRequest,
			"",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, empty response, multiple valid options",
			http.StatusBadRequest,
			"",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrBadRequest,
		},
		{
			"Status BadRequest, empty response, single invalid option",
			http.StatusBadRequest,
			"",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status Unauthorized, populated response, no options",
			http.StatusUnauthorized,
			testResult,
			[]Option{},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, populated response, single valid option",
			http.StatusUnauthorized,
			testResult,
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, populated response, multiple valid options",
			http.StatusUnauthorized,
			testResult,
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, populated response, single invalid option",
			http.StatusUnauthorized,
			testResult,
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status Unauthorized, empty array response, no options",
			http.StatusUnauthorized,
			"[]",
			[]Option{},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, empty array response, single valid option",
			http.StatusUnauthorized,
			"[]",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, empty array response, multiple valid options",
			http.StatusUnauthorized,
			"[]",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, empty array response, single invalid option",
			http.StatusUnauthorized,
			"[]",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status Unauthorized, empty response, no options",
			http.StatusUnauthorized,
			"",
			[]Option{},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, empty response, single valid option",
			http.StatusUnauthorized,
			"",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, empty response, multiple valid options",
			http.StatusUnauthorized,
			"",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrUnauthorized,
		},
		{
			"Status Unauthorized, empty response, single invalid option",
			http.StatusUnauthorized,
			"",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status Forbidden, populated response, no options",
			http.StatusForbidden,
			testResult,
			[]Option{},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, populated response, single valid option",
			http.StatusForbidden,
			testResult,
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, populated response, multiple valid options",
			http.StatusForbidden,
			testResult,
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, populated response, single invalid option",
			http.StatusForbidden,
			testResult,
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status Forbidden, empty array response, no options",
			http.StatusForbidden,
			"[]",
			[]Option{},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, empty array response, single valid option",
			http.StatusForbidden,
			"[]",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, empty array response, multiple valid options",
			http.StatusForbidden,
			"[]",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, empty array response, single invalid option",
			http.StatusForbidden,
			"[]",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status Forbidden, empty response, no options",
			http.StatusForbidden,
			"",
			[]Option{},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, empty response, single valid option",
			http.StatusForbidden,
			"",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, empty response, multiple valid options",
			http.StatusForbidden,
			"",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrForbidden,
		},
		{
			"Status Forbidden, empty response, single invalid option",
			http.StatusForbidden,
			"",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status InternalServerError, populated response, no options",
			http.StatusInternalServerError,
			testResult,
			[]Option{},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, populated response, single valid option",
			http.StatusInternalServerError,
			testResult,
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, populated response, multiple valid options",
			http.StatusInternalServerError,
			testResult,
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, populated response, single invalid option",
			http.StatusInternalServerError,
			testResult,
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status InternalServerError, empty array response, no options",
			http.StatusInternalServerError,
			"[]",
			[]Option{},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, empty array response, single valid option",
			http.StatusInternalServerError,
			"[]",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, empty array response, multiple valid options",
			http.StatusInternalServerError,
			"[]",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, empty array response, single invalid option",
			http.StatusInternalServerError,
			"[]",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
		{
			"Status InternalServerError, empty response, no options",
			http.StatusInternalServerError,
			"",
			[]Option{},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, empty response, single valid option",
			http.StatusInternalServerError,
			"",
			[]Option{SetLimit(15)},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, empty response, multiple valid options",
			http.StatusInternalServerError,
			"",
			[]Option{SetLimit(15), SetOffset(20)},
			testResultPlaceholder{},
			ErrInternalError,
		},
		{
			"Status InternalServerError, empty response, single invalid option",
			http.StatusInternalServerError,
			"",
			[]Option{SetLimit(-99)},
			testResultPlaceholder{},
			ErrOutOfRange,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(test.srvStatus, test.srvResp)
			defer ts.Close()

			res := testResultPlaceholder{}

			err := c.get(testEndpoint, &res, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(res, test.wantRes) {
				t.Errorf("got: <%v>, want: <%v>", res, test.wantRes)
			}
		})
	}
}
