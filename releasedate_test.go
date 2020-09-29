package igdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

const (
	testReleaseDateGet  string = "test_data/releasedate_get.json"
	testReleaseDateList string = "test_data/releasedate_list.json"
)

func TestReleaseDateService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testReleaseDateGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*ReleaseDate, 1)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name            string
		file            string
		id              int
		opts            []Option
		wantReleaseDate *ReleaseDate
		wantErr         error
	}{
		{"Valid response", testReleaseDateGet, 26259, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 26259, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 26259, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			date, err := c.ReleaseDates.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(date, test.wantReleaseDate) {
				t.Errorf("got: <%v>, \nwant: <%v>", date, test.wantReleaseDate)
			}
		})
	}
}

func TestReleaseDateService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testReleaseDateList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*ReleaseDate, 0)
	err = json.Unmarshal(f, &init)

	var tests = []struct {
		name             string
		file             string
		ids              []int
		opts             []Option
		wantReleaseDates []*ReleaseDate
		wantErr          error
	}{
		{"Valid response", testReleaseDateList, []int{16309, 52698, 16321, 106291, 16905}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{16309, 52698, 16321, 106291, 16905}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{16309, 52698, 16321, 106291, 16905}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			date, err := c.ReleaseDates.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(date, test.wantReleaseDates) {
				t.Errorf("got: <%v>, \nwant: <%v>", date, test.wantReleaseDates)
			}
		})
	}
}

func TestReleaseDateService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testReleaseDateList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*ReleaseDate, 0)
	err = json.Unmarshal(f, &init)

	tests := []struct {
		name             string
		file             string
		opts             []Option
		wantReleaseDates []*ReleaseDate
		wantErr          error
	}{
		{"Valid response", testReleaseDateList, []Option{SetLimit(5)}, init, nil},
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

			date, err := c.ReleaseDates.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(date, test.wantReleaseDates) {
				t.Errorf("got: <%v>, \nwant: <%v>", date, test.wantReleaseDates)
			}
		})
	}
}

func TestReleaseDateService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []Option{SetFilter("hypes", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(-99999)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.ReleaseDates.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestReleaseDateService_Fields(t *testing.T) {
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

			fields, err := c.ReleaseDates.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
