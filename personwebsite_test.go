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
	testPersonWebsiteGet  string = "test_data/personwebsite_get.json"
	testPersonWebsiteList string = "test_data/personwebsite_list.json"
)

func TestPersonWebsiteService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPersonWebsiteGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PersonWebsite, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name              string
		file              string
		id                int
		opts              []Option
		wantPersonWebsite *PersonWebsite
		wantErr           error
	}{
		{"Valid response", testPersonWebsiteGet, 1111, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 1111, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 1111, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PersonWebsites.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPersonWebsite) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPersonWebsite)
			}
		})
	}
}

func TestPersonWebsiteService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPersonWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PersonWebsite, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name               string
		file               string
		ids                []int
		opts               []Option
		wantPersonWebsites []*PersonWebsite
		wantErr            error
	}{
		{"Valid response", testPersonWebsiteList, []int{1111, 2222}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1111, 2222}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1111, 2222}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PersonWebsites.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPersonWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPersonWebsites)
			}
		})
	}
}

func TestPersonWebsiteService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPersonWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PersonWebsite, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name               string
		file               string
		opts               []Option
		wantPersonWebsites []*PersonWebsite
		wantErr            error
	}{
		{"Valid response", testPersonWebsiteList, []Option{SetLimit(5)}, init, nil},
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

			web, err := c.PersonWebsites.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPersonWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPersonWebsites)
			}
		})
	}
}

func TestPersonWebsiteService_Count(t *testing.T) {
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

			count, err := c.PersonWebsites.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPersonWebsiteService_Fields(t *testing.T) {
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

			fields, err := c.PersonWebsites.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
