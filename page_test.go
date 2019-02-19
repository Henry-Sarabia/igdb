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
	testPageGet  string = "test_data/page_get.json"
	testPageList string = "test_data/page_list.json"
)

func TestPageService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPageGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Page, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name     string
		file     string
		id       int
		opts     []Option
		wantPage *Page
		wantErr  error
	}{
		{"Valid response", testPageGet, 34, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 34, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 34, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.Pages.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(pg, test.wantPage) {
				t.Errorf("got: <%v>, \nwant: <%v>", pg, test.wantPage)
			}
		})
	}
}

func TestPageService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPageList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Page, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name      string
		file      string
		ids       []int
		opts      []Option
		wantPages []*Page
		wantErr   error
	}{
		{"Valid response", testPageList, []int{7, 90, 23, 154, 386}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{7, 90, 23, 154, 386}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{7, 90, 23, 154, 386}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.Pages.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(pg, test.wantPages) {
				t.Errorf("got: <%v>, \nwant: <%v>", pg, test.wantPages)
			}
		})
	}
}

func TestPageService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPageList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Page, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name      string
		file      string
		opts      []Option
		wantPages []*Page
		wantErr   error
	}{
		{"Valid response", testPageList, []Option{SetLimit(5)}, init, nil},
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

			pg, err := c.Pages.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(pg, test.wantPages) {
				t.Errorf("got: <%v>, \nwant: <%v>", pg, test.wantPages)
			}
		})
	}
}

func TestPageService_Count(t *testing.T) {
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

			count, err := c.Pages.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPageService_Fields(t *testing.T) {
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

			fields, err := c.Pages.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
