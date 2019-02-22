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
	testInvolvedCompanyGet  string = "test_data/involvedcompany_get.json"
	testInvolvedCompanyList string = "test_data/involvedcompany_list.json"
)

func TestInvolvedCompanyService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testInvolvedCompanyGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*InvolvedCompany, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                string
		file                string
		id                  int
		opts                []Option
		wantInvolvedCompany *InvolvedCompany
		wantErr             error
	}{
		{"Valid response", testInvolvedCompanyGet, 36603, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 36603, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 36603, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.InvolvedCompanies.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(com, test.wantInvolvedCompany) {
				t.Errorf("got: <%v>, \nwant: <%v>", com, test.wantInvolvedCompany)
			}
		})
	}
}

func TestInvolvedCompanyService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testInvolvedCompanyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*InvolvedCompany, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                  string
		file                  string
		ids                   []int
		opts                  []Option
		wantInvolvedCompanies []*InvolvedCompany
		wantErr               error
	}{
		{"Valid response", testInvolvedCompanyList, []int{10268, 66143, 8, 65560, 67552}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{10268, 66143, 8, 65560, 67552}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{10268, 66143, 8, 65560, 67552}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.InvolvedCompanies.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(com, test.wantInvolvedCompanies) {
				t.Errorf("got: <%v>, \nwant: <%v>", com, test.wantInvolvedCompanies)
			}
		})
	}
}

func TestInvolvedCompanyService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testInvolvedCompanyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*InvolvedCompany, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                  string
		file                  string
		opts                  []Option
		wantInvolvedCompanies []*InvolvedCompany
		wantErr               error
	}{
		{"Valid response", testInvolvedCompanyList, []Option{SetLimit(5)}, init, nil},
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

			com, err := c.InvolvedCompanies.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(com, test.wantInvolvedCompanies) {
				t.Errorf("got: <%v>, \nwant: <%v>", com, test.wantInvolvedCompanies)
			}
		})
	}
}

func TestInvolvedCompanyService_Count(t *testing.T) {
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

			count, err := c.InvolvedCompanies.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestInvolvedCompanyService_Fields(t *testing.T) {
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

			fields, err := c.InvolvedCompanies.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
