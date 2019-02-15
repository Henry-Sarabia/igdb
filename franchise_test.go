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
	testFranchiseGet  string = "test_data/franchise_get.json"
	testFranchiseList string = "test_data/franchise_list.json"
)

func TestFranchiseService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testFranchiseGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Franchise, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name          string
		file          string
		id            int
		opts          []Option
		wantFranchise *Franchise
		wantErr       error
	}{
		{"Valid response", testFranchiseGet, 43, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 43, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 43, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			fr, err := c.Franchises.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(fr, test.wantFranchise) {
				t.Errorf("got: <%v>, \nwant: <%v>", fr, test.wantFranchise)
			}
		})
	}
}

func TestFranchiseService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testFranchiseList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Franchise, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name           string
		file           string
		ids            []int
		opts           []Option
		wantFranchises []*Franchise
		wantErr        error
	}{
		{"Valid response", testFranchiseList, []int{61, 133, 237, 9, 10}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{61, 133, 237, 9, 10}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{61, 133, 237, 9, 10}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			fr, err := c.Franchises.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(fr, test.wantFranchises) {
				t.Errorf("got: <%v>, \nwant: <%v>", fr, test.wantFranchises)
			}
		})
	}
}

func TestFranchiseService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testFranchiseList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Franchise, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name           string
		file           string
		opts           []Option
		wantFranchises []*Franchise
		wantErr        error
	}{
		{"Valid response", testFranchiseList, []Option{SetLimit(5)}, init, nil},
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

			fr, err := c.Franchises.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(fr, test.wantFranchises) {
				t.Errorf("got: <%v>, \nwant: <%v>", fr, test.wantFranchises)
			}
		})
	}
}

func TestFranchiseService_Count(t *testing.T) {
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

			count, err := c.Franchises.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestFranchiseService_Fields(t *testing.T) {
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

			fields, err := c.Franchises.Fields()
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
