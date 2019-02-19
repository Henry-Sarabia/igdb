// +build ignore

package igdb

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
)

const (
	testZypeGet  string = "test_data/zype_get.json"
	testZypeList string = "test_data/zype_list.json"
)

func TestZypeService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testZypeGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Zype, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name     string
		file     string
		id       int
		opts     []Option
		wantZype *Zype
		wantErr  error
	}{
		{"Valid response", testZypeGet, 777777, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 777777, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 777777, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			z, err := c.Zypes.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(z, test.wantZype) {
				t.Errorf("got: <%v>, \nwant: <%v>", z, test.wantZype)
			}
		})
	}
}

func TestZypeService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testZypeList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Zype, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name      string
		file      string
		ids       []int
		opts      []Option
		wantZypes []*Zype
		wantErr   error
	}{
		{"Valid response", testZypeList, []int{1111}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1111}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1111}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			z, err := c.Zypes.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(z, test.wantZypes) {
				t.Errorf("got: <%v>, \nwant: <%v>", z, test.wantZypes)
			}
		})
	}
}

func TestZypeService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testZypeList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Zype, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name      string
		file      string
		opts      []Option
		wantZypes []*Zype
		wantErr   error
	}{
		{"Valid response", testZypeList, []Option{SetLimit(5)}, init, nil},
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

			z, err := c.Zypes.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(z, test.wantZypes) {
				t.Errorf("got: <%v>, \nwant: <%v>", z, test.wantZypes)
			}
		})
	}
}

func TestZypeService_Count(t *testing.T) {
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

			count, err := c.Zypes.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestZypeService_Fields(t *testing.T) {
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

			fields, err := c.Zypes.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}

func TestZypeService_Search(t *testing.T) {
	f, err := ioutil.ReadFile(testZypeSearch)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Zype, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name      string
		file      string
		qry       string
		opts      []Option
		wantZypes []*Zype
		wantErr   error
	}{
		{"Valid response", testZypeSearch, "mario", []Option{SetLimit(50)}, init, nil},
		{"Empty query", testFileEmpty, "", []Option{SetLimit(50)}, nil, ErrEmptyQuery},
		{"Empty response", testFileEmpty, "mario", nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, "mario", []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, "non-existent entry", nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			z, err := c.Zypes.Search(test.qry, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(z, test.wantZypes) {
				t.Errorf("got: <%v>, \nwant: <%v>", z, test.wantZypes)
			}
		})
	}
}
