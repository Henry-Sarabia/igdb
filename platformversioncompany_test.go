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
	testPlatformVersionCompanyGet  string = "test_data/platformversioncompany_get.json"
	testPlatformVersionCompanyList string = "test_data/platformversioncompany_list.json"
)

func TestPlatformVersionCompanyService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformVersionCompanyGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformVersionCompany, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                       string
		file                       string
		id                         int
		opts                       []Option
		wantPlatformVersionCompany *PlatformVersionCompany
		wantErr                    error
	}{
		{"Valid response", testPlatformVersionCompanyGet, 151, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 151, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 151, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.PlatformVersionCompanies.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(com, test.wantPlatformVersionCompany) {
				t.Errorf("got: <%v>, \nwant: <%v>", com, test.wantPlatformVersionCompany)
			}
		})
	}
}

func TestPlatformVersionCompanyService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformVersionCompanyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformVersionCompany, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                         string
		file                         string
		ids                          []int
		opts                         []Option
		wantPlatformVersionCompanies []*PlatformVersionCompany
		wantErr                      error
	}{
		{"Valid response", testPlatformVersionCompanyList, []int{152, 159, 117, 162, 87}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{152, 159, 117, 162, 87}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{152, 159, 117, 162, 87}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.PlatformVersionCompanies.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(com, test.wantPlatformVersionCompanies) {
				t.Errorf("got: <%v>, \nwant: <%v>", com, test.wantPlatformVersionCompanies)
			}
		})
	}
}

func TestPlatformVersionCompanyService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformVersionCompanyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformVersionCompany, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                         string
		file                         string
		opts                         []Option
		wantPlatformVersionCompanies []*PlatformVersionCompany
		wantErr                      error
	}{
		{"Valid response", testPlatformVersionCompanyList, []Option{SetLimit(5)}, init, nil},
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

			com, err := c.PlatformVersionCompanies.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(com, test.wantPlatformVersionCompanies) {
				t.Errorf("got: <%v>, \nwant: <%v>", com, test.wantPlatformVersionCompanies)
			}
		})
	}
}

func TestPlatformVersionCompanyService_Count(t *testing.T) {
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

			count, err := c.PlatformVersionCompanies.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPlatformVersionCompanyService_Fields(t *testing.T) {
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

			fields, err := c.PlatformVersionCompanies.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
