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
	testGameVersionGet  string = "test_data/gameversion_get.json"
	testGameVersionList string = "test_data/gameversion_list.json"
)

func TestGameVersionService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVersionGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVersion, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name            string
		file            string
		id              int
		opts            []Option
		wantGameVersion *GameVersion
		wantErr         error
	}{
		{"Valid response", testGameVersionGet, 59, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 59, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 59, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ver, err := c.GameVersions.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ver, test.wantGameVersion) {
				t.Errorf("got: <%v>, \nwant: <%v>", ver, test.wantGameVersion)
			}
		})
	}
}

func TestGameVersionService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVersionList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVersion, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name             string
		file             string
		ids              []int
		opts             []Option
		wantGameVersions []*GameVersion
		wantErr          error
	}{
		{"Valid response", testGameVersionList, []int{131, 95, 101, 109, 128}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{131, 95, 101, 109, 128}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{131, 95, 101, 109, 128}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ver, err := c.GameVersions.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ver, test.wantGameVersions) {
				t.Errorf("got: <%v>, \nwant: <%v>", ver, test.wantGameVersions)
			}
		})
	}
}

func TestGameVersionService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVersionList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVersion, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name             string
		file             string
		opts             []Option
		wantGameVersions []*GameVersion
		wantErr          error
	}{
		{"Valid response", testGameVersionList, []Option{SetLimit(5)}, init, nil},
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

			ver, err := c.GameVersions.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ver, test.wantGameVersions) {
				t.Errorf("got: <%v>, \nwant: <%v>", ver, test.wantGameVersions)
			}
		})
	}
}

func TestGameVersionService_Count(t *testing.T) {
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

			count, err := c.GameVersions.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestGameVersionService_Fields(t *testing.T) {
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

			fields, err := c.GameVersions.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
