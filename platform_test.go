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
	testPlatformGet  string = "test_data/platform_get.json"
	testPlatformList string = "test_data/platform_list.json"
)

func TestPlatformService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Platform, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name         string
		file         string
		id           int
		opts         []Option
		wantPlatform *Platform
		wantErr      error
	}{
		{"Valid response", testPlatformGet, 8, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 8, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 8, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			plat, err := c.Platforms.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(plat, test.wantPlatform) {
				t.Errorf("got: <%v>, \nwant: <%v>", plat, test.wantPlatform)
			}
		})
	}
}

func TestPlatformService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Platform, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name          string
		file          string
		ids           []int
		opts          []Option
		wantPlatforms []*Platform
		wantErr       error
	}{
		{"Valid response", testPlatformList, []int{96, 74, 133, 44, 19}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{96, 74, 133, 44, 19}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{96, 74, 133, 44, 19}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			plat, err := c.Platforms.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(plat, test.wantPlatforms) {
				t.Errorf("got: <%v>, \nwant: <%v>", plat, test.wantPlatforms)
			}
		})
	}
}

func TestPlatformService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Platform, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name          string
		file          string
		opts          []Option
		wantPlatforms []*Platform
		wantErr       error
	}{
		{"Valid response", testPlatformList, []Option{SetLimit(5)}, init, nil},
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

			plat, err := c.Platforms.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(plat, test.wantPlatforms) {
				t.Errorf("got: <%v>, \nwant: <%v>", plat, test.wantPlatforms)
			}
		})
	}
}

func TestPlatformService_Count(t *testing.T) {
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

			count, err := c.Platforms.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPlatformService_Fields(t *testing.T) {
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

			fields, err := c.Platforms.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
