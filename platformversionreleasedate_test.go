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
	testPlatformVersionReleaseDateGet  string = "test_data/platformversionreleasedate_get.json"
	testPlatformVersionReleaseDateList string = "test_data/platformversionreleasedate_list.json"
)

func TestPlatformVersionReleaseDateService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformVersionReleaseDateGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformVersionReleaseDate, 1)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name                           string
		file                           string
		id                             int
		opts                           []Option
		wantPlatformVersionReleaseDate *PlatformVersionReleaseDate
		wantErr                        error
	}{
		{"Valid response", testPlatformVersionReleaseDateGet, 6, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 6, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 6, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			date, err := c.PlatformVersionReleaseDates.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(date, test.wantPlatformVersionReleaseDate) {
				t.Errorf("got: <%v>, \nwant: <%v>", date, test.wantPlatformVersionReleaseDate)
			}
		})
	}
}

func TestPlatformVersionReleaseDateService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformVersionReleaseDateList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformVersionReleaseDate, 0)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name                            string
		file                            string
		ids                             []int
		opts                            []Option
		wantPlatformVersionReleaseDates []*PlatformVersionReleaseDate
		wantErr                         error
	}{
		{"Valid response", testPlatformVersionReleaseDateList, []int{29, 37, 40, 48, 81}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{29, 37, 40, 48, 81}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{29, 37, 40, 48, 81}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			date, err := c.PlatformVersionReleaseDates.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(date, test.wantPlatformVersionReleaseDates) {
				t.Errorf("got: <%v>, \nwant: <%v>", date, test.wantPlatformVersionReleaseDates)
			}
		})
	}
}

func TestPlatformVersionReleaseDateService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformVersionReleaseDateList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformVersionReleaseDate, 0)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                            string
		file                            string
		opts                            []Option
		wantPlatformVersionReleaseDates []*PlatformVersionReleaseDate
		wantErr                         error
	}{
		{"Valid response", testPlatformVersionReleaseDateList, []Option{SetLimit(5)}, init, nil},
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

			date, err := c.PlatformVersionReleaseDates.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(date, test.wantPlatformVersionReleaseDates) {
				t.Errorf("got: <%v>, \nwant: <%v>", date, test.wantPlatformVersionReleaseDates)
			}
		})
	}
}

func TestPlatformVersionReleaseDateService_Count(t *testing.T) {
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

			count, err := c.PlatformVersionReleaseDates.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPlatformVersionReleaseDateService_Fields(t *testing.T) {
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

			fields, err := c.PlatformVersionReleaseDates.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
