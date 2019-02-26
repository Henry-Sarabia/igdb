package igdb

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const testStatus string = "test_data/status.json"

func TestClient_Status(t *testing.T) {
	f, err := ioutil.ReadFile(testStatus)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Status, 1)
	json.Unmarshal(f, &init)

	tests := []struct {
		name     string
		file     string
		wantStat *Status
		wantErr  error
	}{
		{"Valid response", testStatus, init[0], nil},
		{"Empty response", testFileEmpty, nil, errInvalidJSON},
		{"No results", testFileEmptyArray, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			stat, err := c.Status()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(stat, test.wantStat) {
				t.Errorf("got: <%v>, \nwant: <%v>", stat, test.wantStat)
			}
		})
	}
}
