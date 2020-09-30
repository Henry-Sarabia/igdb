package igdb

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
)

func TestClient_GetFields(t *testing.T) {
	var tests = []struct {
		name       string
		status     int
		resp       string
		wantFields []string
		wantErr    error
	}{
		{"OK status with regular response", http.StatusOK, `["name", "slug", "url"]`, []string{"url", "slug", "name"}, nil},
		{"OK status with empty response", http.StatusOK, "", nil, errInvalidJSON},
		{"OK status with dot response", http.StatusOK, `["mugshot.width","name", "company.id"]`, []string{"company.id", "name", "mugshot.width"}, nil},
		{"OK status with asterisk response", http.StatusOK, `["*"]`, []string{"*"}, nil},
		{"Bad status with empty response", http.StatusBadRequest, "", nil, ErrBadRequest},
		{"Not found status with error response", http.StatusNotFound, testErrNotFound, nil, ServerError{Status: 404, Msg: "status not found"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(test.status, test.resp)
			defer ts.Close()

			f, err := c.getFields(testEndpoint)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(f, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, f)
			}
		})
	}
}

func TestClient_GetCount(t *testing.T) {
	var tests = []struct {
		name      string
		status    int
		resp      string
		wantCount int
		wantErr   error
	}{
		{"OK status with regular response", http.StatusOK, `{"count": 1234}`, 1234, nil},
		{"OK status with count of zero response", http.StatusOK, `{"count": 0}`, 0, nil},
		{"OK status with empty response", http.StatusOK, "", 0, errInvalidJSON},
		{"Bad status with empty response", http.StatusBadRequest, "", 0, ErrBadRequest},
		{"Not found status with error response", http.StatusNotFound, testErrNotFound, 0, ServerError{Status: 404, Msg: "status not found"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(test.status, test.resp)
			defer ts.Close()

			count, err := c.getCount(testEndpoint)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("Expected count %d, got %d", test.wantCount, count)
			}
		})
	}
}
