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
	testProductFamilyGet  string = "test_data/productfamily_get.json"
	testProductFamilyList string = "test_data/productfamily_list.json"
)

func TestProductFamilyService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testProductFamilyGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*ProductFamily, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name              string
		file              string
		id                int
		opts              []Option
		wantProductFamily *ProductFamily
		wantErr           error
	}{
		{"Valid response", testProductFamilyGet, 1, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 1, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 1, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.ProductFamilies.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(g, test.wantProductFamily) {
				t.Errorf("got: <%v>, \nwant: <%v>", g, test.wantProductFamily)
			}
		})
	}
}

func TestProductFamilyService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testProductFamilyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*ProductFamily, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                string
		file                string
		ids                 []int
		opts                []Option
		wantProductFamilies []*ProductFamily
		wantErr             error
	}{
		{"Valid response", testProductFamilyList, []int{3, 2, 4, 5}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{3, 2, 4, 5}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{3, 2, 4, 5}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.ProductFamilies.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(g, test.wantProductFamilies) {
				t.Errorf("got: <%v>, \nwant: <%v>", g, test.wantProductFamilies)
			}
		})
	}
}

func TestProductFamilyService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testProductFamilyList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*ProductFamily, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                string
		file                string
		opts                []Option
		wantProductFamilies []*ProductFamily
		wantErr             error
	}{
		{"Valid response", testProductFamilyList, []Option{SetLimit(5)}, init, nil},
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

			g, err := c.ProductFamilies.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(g, test.wantProductFamilies) {
				t.Errorf("got: <%v>, \nwant: <%v>", g, test.wantProductFamilies)
			}
		})
	}
}

func TestProductFamilyService_Count(t *testing.T) {
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

			count, err := c.ProductFamilies.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestProductFamilyService_Fields(t *testing.T) {
	var tests = []struct {
		name       string
		resp       string
		wantFields []string
		wantErr    error
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, nil},
		{"Dot operator", `["logo.url", "background.id"]`, []string{"background.id", "logo.url"}, nil},
		{"Asterisk", `["*"]`, []string{"*"}, nil},
		{"Empty response", "", nil, errInvalidJSON},
		{"No results", "[]", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			fields, err := c.ProductFamilies.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
