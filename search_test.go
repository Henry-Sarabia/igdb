package igdb

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const testSearch string = "test_data/search.json"

func TestClient_Search(t *testing.T) {
	f, err := ioutil.ReadFile(testSearch)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*SearchResult, 0)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		file    string
		qry     string
		opts    []Option
		wantRes []*SearchResult
		wantErr error
	}{
		{"Valid response", testSearch, "sonic", []Option{SetFields("*")}, init, nil},
		{"Empty query", testFileEmpty, "", []Option{SetLimit(50)}, nil, ErrEmptyQry},
		{"Empty response", testFileEmpty, "sonic", nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, "sonic", []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, "non-existent entry", nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			z, err := c.Search(test.qry, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(z, test.wantRes) {
				t.Errorf("got: <%v>, \nwant: <%v>", z, test.wantRes)
			}
		})
	}
}
