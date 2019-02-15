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
	testKeywordGet  string = "test_data/keyword_get.json"
	testKeywordList string = "test_data/keyword_list.json"
)

func TestKeywordService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testKeywordGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Keyword, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name        string
		file        string
		id          int
		opts        []Option
		wantKeyword *Keyword
		wantErr     error
	}{
		{"Valid response", testKeywordGet, 19226, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 19226, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 19226, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			key, err := c.Keywords.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(key, test.wantKeyword) {
				t.Errorf("got: <%v>, \nwant: <%v>", key, test.wantKeyword)
			}
		})
	}
}

func TestKeywordService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testKeywordList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Keyword, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name         string
		file         string
		ids          []int
		opts         []Option
		wantKeywords []*Keyword
		wantErr      error
	}{
		{"Valid response", testKeywordList, []int{31, 18534, 12071, 6939, 7281}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{31, 18534, 12071, 6939, 7281}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{31, 18534, 12071, 6939, 7281}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			key, err := c.Keywords.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(key, test.wantKeywords) {
				t.Errorf("got: <%v>, \nwant: <%v>", key, test.wantKeywords)
			}
		})
	}
}

func TestKeywordService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testKeywordList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Keyword, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name         string
		file         string
		opts         []Option
		wantKeywords []*Keyword
		wantErr      error
	}{
		{"Valid response", testKeywordList, []Option{SetLimit(5)}, init, nil},
		{"Empty response", testFileEmpty, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			key, err := c.Keywords.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(key, test.wantKeywords) {
				t.Errorf("got: <%v>, \nwant: <%v>", key, test.wantKeywords)
			}
		})
	}
}

func TestKeywordService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []Option{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.Keywords.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestKeywordService_Fields(t *testing.T) {
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

			fields, err := c.Keywords.Fields()
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
