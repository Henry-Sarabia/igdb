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
	testPulseGroupGet  string = "test_data/pulsegroup_get.json"
	testPulseGroupList string = "test_data/pulsegroup_list.json"
)

func TestPulseGroupService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPulseGroupGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PulseGroup, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name           string
		file           string
		id             int
		opts           []Option
		wantPulseGroup *PulseGroup
		wantErr        error
	}{
		{"Valid response", testPulseGroupGet, 61381, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 61381, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 61381, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.PulseGroups.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(pg, test.wantPulseGroup) {
				t.Errorf("got: <%v>, \nwant: <%v>", pg, test.wantPulseGroup)
			}
		})
	}
}

func TestPulseGroupService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPulseGroupList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PulseGroup, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name            string
		file            string
		ids             []int
		opts            []Option
		wantPulseGroups []*PulseGroup
		wantErr         error
	}{
		{"Valid response", testPulseGroupList, []int{65166, 41927, 65171, 71895}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{65166, 41927, 65171, 71895}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{65166, 41927, 65171, 71895}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.PulseGroups.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(pg, test.wantPulseGroups) {
				t.Errorf("got: <%v>, \nwant: <%v>", pg, test.wantPulseGroups)
			}
		})
	}
}

func TestPulseGroupService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPulseGroupList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PulseGroup, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name            string
		file            string
		opts            []Option
		wantPulseGroups []*PulseGroup
		wantErr         error
	}{
		{"Valid response", testPulseGroupList, []Option{SetLimit(5)}, init, nil},
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

			pg, err := c.PulseGroups.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(pg, test.wantPulseGroups) {
				t.Errorf("got: <%v>, \nwant: <%v>", pg, test.wantPulseGroups)
			}
		})
	}
}

func TestPulseGroupService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []Option{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(-100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.PulseGroups.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestPulseGroupService_Fields(t *testing.T) {
	var tests = []struct {
		name       string
		resp       string
		wantFields []string
		wantErr    error
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, nil},
		{"Dot operator", `["logo.url", "background.id"]`, []string{"background.id", "logo.url"}, nil},
		{"Asterisk", `["*"]`, []string{"*"}, nil},
		{"Empty response", "", nil, errInvalidJSON},
		{"No results", "[]", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			fields, err := c.PulseGroups.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
