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
	testPulseSourceGet  string = "test_data/pulsesource_get.json"
	testPulseSourceList string = "test_data/pulsesource_list.json"
)

func TestPulseSourceService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPulseSourceGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PulseSource, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name            string
		file            string
		id              int
		opts            []Option
		wantPulseSource *PulseSource
		wantErr         error
	}{
		{"Valid response", testPulseSourceGet, 54, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 54, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 54, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			src, err := c.PulseSources.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(src, test.wantPulseSource) {
				t.Errorf("got: <%v>, \nwant: <%v>", src, test.wantPulseSource)
			}
		})
	}
}

func TestPulseSourceService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPulseSourceList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PulseSource, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name             string
		file             string
		ids              []int
		opts             []Option
		wantPulseSources []*PulseSource
		wantErr          error
	}{
		{"Valid response", testPulseSourceList, []int{37, 35, 3, 22, 29}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{37, 35, 3, 22, 29}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{37, 35, 3, 22, 29}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			src, err := c.PulseSources.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(src, test.wantPulseSources) {
				t.Errorf("got: <%v>, \nwant: <%v>", src, test.wantPulseSources)
			}
		})
	}
}

func TestPulseSourceService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPulseSourceList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PulseSource, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name             string
		file             string
		opts             []Option
		wantPulseSources []*PulseSource
		wantErr          error
	}{
		{"Valid response", testPulseSourceList, []Option{SetLimit(5)}, init, nil},
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

			src, err := c.PulseSources.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(src, test.wantPulseSources) {
				t.Errorf("got: <%v>, \nwant: <%v>", src, test.wantPulseSources)
			}
		})
	}
}

func TestPulseSourceService_Count(t *testing.T) {
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

			count, err := c.PulseSources.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestPulseSourceService_Fields(t *testing.T) {
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

			fields, err := c.PulseSources.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
