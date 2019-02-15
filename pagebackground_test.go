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
	testPageBackgroundGet  string = "test_data/pagebackground_get.json"
	testPageBackgroundList string = "test_data/pagebackground_list.json"
)

func TestPageBackgroundService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPageBackgroundGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PageBackground, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name               string
		file               string
		id                 int
		opts               []Option
		wantPageBackground *PageBackground
		wantErr            error
	}{
		{"Valid response", testPageBackgroundGet, 1, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 1, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 1, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			bg, err := c.PageBackgrounds.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(bg, test.wantPageBackground) {
				t.Errorf("got: <%v>, \nwant: <%v>", bg, test.wantPageBackground)
			}
		})
	}
}

func TestPageBackgroundService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPageBackgroundList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PageBackground, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                string
		file                string
		ids                 []int
		opts                []Option
		wantPageBackgrounds []*PageBackground
		wantErr             error
	}{
		{"Valid response", testPageBackgroundList, []int{2, 3, 4, 5, 6}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{2, 3, 4, 5, 6}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{2, 3, 4, 5, 6}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			bg, err := c.PageBackgrounds.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(bg, test.wantPageBackgrounds) {
				t.Errorf("got: <%v>, \nwant: <%v>", bg, test.wantPageBackgrounds)
			}
		})
	}
}

func TestPageBackgroundService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPageBackgroundList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PageBackground, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                string
		file                string
		opts                []Option
		wantPageBackgrounds []*PageBackground
		wantErr             error
	}{
		{"Valid response", testPageBackgroundList, []Option{SetLimit(5)}, init, nil},
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

			bg, err := c.PageBackgrounds.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(bg, test.wantPageBackgrounds) {
				t.Errorf("got: <%v>, \nwant: <%v>", bg, test.wantPageBackgrounds)
			}
		})
	}
}

func TestPageBackgroundService_Count(t *testing.T) {
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

			count, err := c.PageBackgrounds.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPageBackgroundService_Fields(t *testing.T) {
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

			fields, err := c.PageBackgrounds.Fields()
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
