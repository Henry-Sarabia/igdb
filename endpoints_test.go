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
		{"OK status with regular response", http.StatusOK, `["name", "slug", "url"]`, []string{"url", "slug", "name"}, ""},
		{"OK status with empty response", http.StatusOK, "", nil, errEndOfJSON.Error()},
		{"OK status with dot response", http.StatusOK, `["mugshot.width","name", "company.id"]`, []string{"company.id", "name", "mugshot.width"}, ""},
		{"OK status with asterisk response", http.StatusOK, `["*"]`, []string{"*"}, ""},
		{"Bad status with empty response", http.StatusBadRequest, "", nil, ErrBadRequest.Error()},
		{"Not found status with error response", http.StatusNotFound, testErrNotFound, nil, "Status 404 - status not found"},
	}

	for _, tt := range fieldsTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(tt.Status, tt.Resp)
			defer ts.Close()

			fields, err := c.getFields(testEndpoint)
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
		{"OK status with regular response", http.StatusOK, `{"count": 1234}`, 1234, ""},
		{"OK status with count of zero response", http.StatusOK, `{"count": 0}`, 0, ""},
		{"OK status with empty response", http.StatusOK, "", 0, errEndOfJSON.Error()},
		{"Bad status with empty response", http.StatusBadRequest, "", 0, ErrBadRequest.Error()},
		{"Not found status with error response", http.StatusNotFound, testErrNotFound, 0, "Status 404 - status not found"},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(tt.Status, tt.Resp)
			defer ts.Close()

			count, err := c.getCount(testEndpoint)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}
