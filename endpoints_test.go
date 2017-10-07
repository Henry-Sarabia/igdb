package igdb

import (
	"net/http"
	"testing"
)

func TestGetEndpointFields(t *testing.T) {
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

	for _, ft := range fieldsTests {
		t.Run(ft.Name, func(t *testing.T) {
			ts, c := startTestServer(ft.Status, ft.Resp)
			defer ts.Close()

			fields, err := c.GetEndpointFields(testEndpoint)
			actErr := ""
			if err != nil {
				actErr = err.Error()
			}

			if actErr != ft.ExpErr {
				t.Fatalf("Expected error '%s', got '%s'", ft.ExpErr, actErr)
			}

			ok, err := equalSlice(fields, ft.ExpFields)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Fatalf("Expected fields '%v', got '%v'", ft.ExpFields, fields)
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

	for _, ct := range countTests {
		t.Run(ct.Name, func(t *testing.T) {
			ts, c := startTestServer(ct.Status, ct.Resp)
			defer ts.Close()

			count, err := c.GetEndpointCount(testEndpoint)
			actErr := ""
			if err != nil {
				actErr = err.Error()
			}

			if actErr != ct.ExpErr {
				t.Fatalf("Expecter error '%v', got '%v'", ct.ExpErr, actErr)
			}

			if count != ct.ExpCount {
				t.Fatalf("Expected count %d, got %d", ct.ExpCount, count)
			}
		})
	}
}
