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
	testCoverGet  string = "test_data/cover_get.json"
	testCoverList string = "test_data/cover_list.json"
)

func TestCoverService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testCoverGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Cover, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name      string
		file      string
		id        int
		opts      []Option
		wantCover *Cover
		wantErr   error
	}{
		{"Valid response", testCoverGet, 63541, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 63541, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 63541, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cov, err := c.Covers.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(cov, test.wantCover) {
				t.Errorf("got: <%v>, \nwant: <%v>", cov, test.wantCover)
			}
		})
	}
}

func TestCoverService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testCoverList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Cover, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name       string
		file       string
		ids        []int
		opts       []Option
		wantCovers []*Cover
		wantErr    error
	}{
		{"Valid response", testCoverList, []int{54614, 9206, 15242, 43854, 27257}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{54614, 9206, 15242, 43854, 27257}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{54614, 9206, 15242, 43854, 27257}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cov, err := c.Covers.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(cov, test.wantCovers) {
				t.Errorf("got: <%v>, \nwant: <%v>", cov, test.wantCovers)
			}
		})
	}
}

func TestCoverService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testCoverList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Cover, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name       string
		file       string
		opts       []Option
		wantCovers []*Cover
		wantErr    error
	}{
		{"Valid response", testCoverList, []Option{SetLimit(5)}, init, nil},
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

			cov, err := c.Covers.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(cov, test.wantCovers) {
				t.Errorf("got: <%v>, \nwant: <%v>", cov, test.wantCovers)
			}
		})
	}
}

func TestCoverService_Count(t *testing.T) {
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

			count, err := c.Covers.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestCoverService_Fields(t *testing.T) {
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

			fields, err := c.Covers.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
