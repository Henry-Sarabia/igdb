package igdb

import (
	"net/http"
	"testing"
)

func TestGetEndpointFieldList(t *testing.T) {
	var fieldsTests = []struct {
		Name      string
		Status    int
		Resp      string
		ExpFields []string
		ExpErr    string
	}{
		{"OK request with non-empty response", http.StatusOK, `["name", "slug", "url"]`, []string{"url", "slug", "name"}, ""},
		{"OK request with empty response", http.StatusOK, "", nil, errEndOfJSON.Error()},
		{"Bad request with empty response", http.StatusBadRequest, "", nil, errEndOfJSON.Error()},
	}

	for _, tt := range fieldsTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(tt.Status, tt.Resp)
			defer ts.Close()

			fields, err := c.GetEndpointFieldList(testEndpoint)
			assertError(t, err, tt.ExpErr)

			ok, err := equalSlice(fields, tt.ExpFields)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Fatalf("Expected fields '%v', got '%v'", tt.ExpFields, fields)
			}
		})
	}
}

func TestGetEndpointCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Status   int
		Resp     string
		ExpCount int
		ExpErr   string
	}{
		{"OK request with non-empty response", http.StatusOK, `{"count": 1234}`, 1234, ""},
		{"OK request with empty response", http.StatusOK, "", 0, errEndOfJSON.Error()},
		{"Bad request with empty response", http.StatusBadRequest, "", 0, errEndOfJSON.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(tt.Status, tt.Resp)
			defer ts.Close()

			count, err := c.GetEndpointCount(testEndpoint)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}
