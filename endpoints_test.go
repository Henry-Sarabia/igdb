package igdb

import (
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"testing"
)

const testEndpoint = "test/"

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
		{"Not found status with error response", http.StatusNotFound, testErrNotFound, nil, ServerError{Status: 404, Message: "status not found"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(test.status, test.resp)
			defer ts.Close()

			f, err := c.getFields(testEndpoint)
			if !reflect.DeepEqual(errors.Cause(err), test.wantErr) {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			ok, err := equalSlice(f, test.wantFields)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Errorf("got: <%v>, want: <%v>", test.wantFields, f)
			}
		})
	}
}

//func TestClient_GetCount(t *testing.T) {
//	var countTests = []struct {
//		Name     string
//		Status   int
//		Resp     string
//		ExpCount int
//		ExpErr   string
//	}{
//		{"OK status with regular response", http.StatusOK, `{"count": 1234}`, 1234, ""},
//		{"OK status with count of zero response", http.StatusOK, `{"count": 0}`, 0, ""},
//		{"OK status with empty response", http.StatusOK, "", 0, errInvalidJSON.Error()},
//		{"Bad status with empty response", http.StatusBadRequest, "", 0, ErrBadRequest.Error()},
//		{"Not found status with error response", http.StatusNotFound, testErrNotFound, 0, "Status 404 - status not found"},
//	}
//
//	for _, tt := range countTests {
//		t.Run(tt.Name, func(t *testing.T) {
//			ts, c := testServerString(tt.Status, tt.Resp)
//			defer ts.Close()
//
//			count, err := c.getCount(testEndpoint)
//			//assertError(t, err, tt.ExpErr)
//
//			if count != tt.ExpCount {
//				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
//			}
//		})
//	}
//}
