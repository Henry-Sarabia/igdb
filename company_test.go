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
	testCompanyGet  string = "test_data/company_get.json"
	testCompanyList string = "test_data/company_list.json"
)

func TestCompanyService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testCompanyGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Company, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name        string
		file        string
		id          int
		opts        []Option
		wantCompany *Company
		wantErr     error
	}{
		{"Valid response", testCompanyGet, 13710, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 13710, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 13710, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			comp, err := c.Companies.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(comp, test.wantCompany) {
				t.Errorf("got: <%v>, \nwant: <%v>", comp, test.wantCompany)
			}
		})
	}
}

func TestCompanyService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testCompanyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Company, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name          string
		file          string
		ids           []int
		opts          []Option
		wantCompanies []*Company
		wantErr       error
	}{
		{"Valid response", testCompanyList, []int{10815, 16954, 8199, 14672, 13535}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{10815, 16954, 8199, 14672, 13535}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{10815, 16954, 8199, 14672, 13535}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			comp, err := c.Companies.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(comp, test.wantCompanies) {
				t.Errorf("got: <%v>, \nwant: <%v>", comp, test.wantCompanies)
			}
		})
	}
}

func TestCompanyService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testCompanyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Company, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name          string
		file          string
		opts          []Option
		wantCompanies []*Company
		wantErr       error
	}{
		{"Valid response", testCompanyList, []Option{SetLimit(5)}, init, nil},
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

			comp, err := c.Companies.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(comp, test.wantCompanies) {
				t.Errorf("got: <%v>, \nwant: <%v>", comp, test.wantCompanies)
			}
		})
	}
}

func TestCompanyService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []Option{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(-100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.Companies.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestCompanyService_Fields(t *testing.T) {
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

			fields, err := c.Companies.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
