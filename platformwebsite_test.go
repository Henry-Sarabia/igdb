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
	testPlatformWebsiteGet  string = "test_data/platformwebsite_get.json"
	testPlatformWebsiteList string = "test_data/platformwebsite_list.json"
)

func TestPlatformWebsiteService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformWebsiteGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformWebsite, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                string
		file                string
		id                  int
		opts                []FuncOption
		wantPlatformWebsite *PlatformWebsite
		wantErr             error
	}{
		{"Valid response", testPlatformWebsiteGet, 16, []FuncOption{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 16, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 16, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PlatformWebsites.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPlatformWebsite) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPlatformWebsite)
			}
		})
	}
}

func TestPlatformWebsiteService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformWebsite, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                 string
		file                 string
		ids                  []int
		opts                 []FuncOption
		wantPlatformWebsites []*PlatformWebsite
		wantErr              error
	}{
		{"Valid response", testPlatformWebsiteList, []int{1, 18, 32, 6, 29}, []FuncOption{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1, 18, 32, 6, 29}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1, 18, 32, 6, 29}, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PlatformWebsites.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPlatformWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPlatformWebsites)
			}
		})
	}
}

func TestPlatformWebsiteService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformWebsiteList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformWebsite, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                 string
		file                 string
		opts                 []FuncOption
		wantPlatformWebsites []*PlatformWebsite
		wantErr              error
	}{
		{"Valid response", testPlatformWebsiteList, []FuncOption{SetLimit(5)}, init, nil},
		{"Empty response", testFileEmpty, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			web, err := c.PlatformWebsites.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(web, test.wantPlatformWebsites) {
				t.Errorf("got: <%v>, \nwant: <%v>", web, test.wantPlatformWebsites)
			}
		})
	}
}

func TestPlatformWebsiteService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []FuncOption
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []FuncOption{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []FuncOption{SetLimit(100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.PlatformWebsites.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPlatformWebsiteService_Fields(t *testing.T) {
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

			fields, err := c.PlatformWebsites.Fields()
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
