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
	testCompanyWebsiteGet  string = "test_data/companywebsite_get.json"
	testCompanyWebsiteList string = "test_data/companywebsite_list.json"
)

func TestCompanyWebsiteService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testCompanyWebsiteGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*CompanyWebsite, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name               string
		file               string
		id                 int
		opts               []Option
		wantCompanyWebsite *CompanyWebsite
		wantErr            error
	}{
		{"Valid response", testCompanyWebsiteGet, 1707, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 1707, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 1707, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.CompanyWebsites.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantCompanyWebsite) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantCompanyWebsite)
			}
		})
	}
}

func TestCompanyWebsiteService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testCompanyWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*CompanyWebsite, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                string
		file                string
		ids                 []int
		opts                []Option
		wantCompanyWebsites []*CompanyWebsite
		wantErr             error
	}{
		{"Valid response", testCompanyWebsiteList, []int{1709, 453, 1710, 1322, 133}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1709, 453, 1710, 1322, 133}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1709, 453, 1710, 1322, 133}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.CompanyWebsites.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantCompanyWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantCompanyWebsites)
			}
		})
	}
}

func TestCompanyWebsiteService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testCompanyWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*CompanyWebsite, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                string
		file                string
		opts                []Option
		wantCompanyWebsites []*CompanyWebsite
		wantErr             error
	}{
		{"Valid response", testCompanyWebsiteList, []Option{SetLimit(5)}, init, nil},
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

			web, err := c.CompanyWebsites.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantCompanyWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantCompanyWebsites)
			}
		})
	}
}

func TestCompanyWebsiteService_Count(t *testing.T) {
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

			count, err := c.CompanyWebsites.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestCompanyWebsiteService_Fields(t *testing.T) {
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

			fields, err := c.CompanyWebsites.Fields()
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
