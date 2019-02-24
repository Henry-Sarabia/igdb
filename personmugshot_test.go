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
	testPersonMugshotGet  string = "test_data/personmugshot_get.json"
	testPersonMugshotList string = "test_data/personmugshot_list.json"
)

func TestPersonMugshotService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPersonMugshotGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PersonMugshot, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name              string
		file              string
		id                int
		opts              []Option
		wantPersonMugshot *PersonMugshot
		wantErr           error
	}{
		{"Valid response", testPersonMugshotGet, 1111, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 1111, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 1111, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			mug, err := c.PersonMugshots.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(mug, test.wantPersonMugshot) {
				t.Errorf("got: <%v>, \nwant: <%v>", mug, test.wantPersonMugshot)
			}
		})
	}
}

func TestPersonMugshotService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPersonMugshotList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PersonMugshot, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name               string
		file               string
		ids                []int
		opts               []Option
		wantPersonMugshots []*PersonMugshot
		wantErr            error
	}{
		{"Valid response", testPersonMugshotList, []int{1111, 2222}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1111, 2222}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1111, 2222}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			mug, err := c.PersonMugshots.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(mug, test.wantPersonMugshots) {
				t.Errorf("got: <%v>, \nwant: <%v>", mug, test.wantPersonMugshots)
			}
		})
	}
}

func TestPersonMugshotService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPersonMugshotList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PersonMugshot, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name               string
		file               string
		opts               []Option
		wantPersonMugshots []*PersonMugshot
		wantErr            error
	}{
		{"Valid response", testPersonMugshotList, []Option{SetLimit(5)}, init, nil},
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

			mug, err := c.PersonMugshots.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(mug, test.wantPersonMugshots) {
				t.Errorf("got: <%v>, \nwant: <%v>", mug, test.wantPersonMugshots)
			}
		})
	}
}

func TestPersonMugshotService_Count(t *testing.T) {
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

			count, err := c.PersonMugshots.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPersonMugshotService_Fields(t *testing.T) {
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

			fields, err := c.PersonMugshots.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
