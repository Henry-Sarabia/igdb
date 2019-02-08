package igdb

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const (
	testCollectionGet    string = "test_data/collection_get.json"
	testCollectionList   string = "test_data/collection_list.json"
	testCollectionSearch string = "test_data/collection_search.json"
)

func TestCollectionService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testCollectionGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Collection, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name           string
		file           string
		id             int
		opts           []FuncOption
		wantCollection *Collection
		wantErr        error
	}{
		{"Valid response", testCollectionGet, 286, []FuncOption{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 286, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 286, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			col, err := c.Collections.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(col, test.wantCollection) {
				t.Errorf("got: <%v>, \nwant: <%v>", col, test.wantCollection)
			}
		})
	}
}

func TestCollectionService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testCollectionList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Collection, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name            string
		file            string
		ids             []int
		opts            []FuncOption
		wantCollections []*Collection
		wantErr         error
	}{
		{"Valid response", testCollectionList, []int{301, 4010, 364, 457, 719}, []FuncOption{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{301, 4010, 364, 457, 719}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{301, 4010, 364, 457, 719}, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			col, err := c.Collections.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(col, test.wantCollections) {
				t.Errorf("got: <%v>, \nwant: <%v>", col, test.wantCollections)
			}
		})
	}
}

func TestCollectionService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testCollectionList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Collection, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name            string
		file            string
		opts            []FuncOption
		wantCollections []*Collection
		wantErr         error
	}{
		{"Valid response", testCollectionList, []FuncOption{SetLimit(5)}, init, nil},
		{"Empty response", testFileEmpty, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			col, err := c.Collections.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(col, test.wantCollections) {
				t.Errorf("got: <%v>, \nwant: <%v>", col, test.wantCollections)
			}
		})
	}
}

func TestCollectionService_Search(t *testing.T) {
	f, err := ioutil.ReadFile(testCollectionSearch)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Collection, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name            string
		file            string
		qry             string
		opts            []FuncOption
		wantCollections []*Collection
		wantErr         error
	}{
		{"Valid response", testCollectionSearch, "super", []FuncOption{SetLimit(50)}, init, nil},
		{"Empty query", testFileEmpty, "", []FuncOption{SetLimit(50)}, nil, ErrEmptyQuery},
		{"Empty response", testFileEmpty, "super", nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, "super", []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, "non-existent entry", nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			col, err := c.Collections.Search(test.qry, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(col, test.wantCollections) {
				t.Errorf("got: <%v>, \nwant: <%v>", col, test.wantCollections)
			}
		})
	}
}

func TestCollectionService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []FuncOption
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []FuncOption{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []FuncOption{SetLimit(100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.Collections.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestCollectionService_Fields(t *testing.T) {
	var tests = []struct {
		name       string
		resp       string
		wantFields []string
		wantErr    error
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, nil},
		{"Asterisk", `["*"]`, []string{"*"}, nil},
		{"Empty response", "", nil, errInvalidJSON},
		{"No results", "[]", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			fields, err := c.Collections.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			ok, err := equalSlice(fields, test.wantFields)
			if err != nil {
				t.Fatal(err)
			}

			if !ok {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
