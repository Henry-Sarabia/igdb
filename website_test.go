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
	testWebsiteGet  string = "test_data/website_get.json"
	testWebsiteList string = "test_data/website_list.json"
)

func TestWebsiteService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testWebsiteGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Website, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name        string
		file        string
		id          int
		opts        []Option
		wantWebsite *Website
		wantErr     error
	}{
		{"Valid response", testWebsiteGet, 52133, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 52133, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 52133, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.Websites.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantWebsite) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantWebsite)
			}
		})
	}
}

func TestWebsiteService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Website, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name         string
		file         string
		ids          []int
		opts         []Option
		wantWebsites []*Website
		wantErr      error
	}{
		{"Valid response", testWebsiteList, []int{95440, 94413, 90071, 20460, 83935}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{95440, 94413, 90071, 20460, 83935}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{95440, 94413, 90071, 20460, 83935}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.Websites.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantWebsites)
			}
		})
	}
}

func TestWebsiteService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Website, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name         string
		file         string
		opts         []Option
		wantWebsites []*Website
		wantErr      error
	}{
		{"Valid response", testWebsiteList, []Option{SetLimit(5)}, init, nil},
		{"Empty response", testFileEmpty, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.Websites.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantWebsites)
			}
		})
	}
}

func TestWebsiteService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []Option{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(-99999)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.Websites.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestWebsiteService_Fields(t *testing.T) {
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

			fields, err := c.Websites.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
