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
	testGameVideoGet  string = "test_data/gamevideo_get.json"
	testGameVideoList string = "test_data/gamevideo_list.json"
)

func TestGameVideoService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVideoGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVideo, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name          string
		file          string
		id            int
		opts          []Option
		wantGameVideo *GameVideo
		wantErr       error
	}{
		{"Valid response", testGameVideoGet, 24648, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 24648, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 24648, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			vid, err := c.GameVideos.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(vid, test.wantGameVideo) {
				t.Errorf("got: <%v>, \nwant: <%v>", vid, test.wantGameVideo)
			}
		})
	}
}

func TestGameVideoService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVideoList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVideo, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name           string
		file           string
		ids            []int
		opts           []Option
		wantGameVideos []*GameVideo
		wantErr        error
	}{
		{"Valid response", testGameVideoList, []int{24669, 24628, 24671, 24603, 24706}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{24669, 24628, 24671, 24603, 24706}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{24669, 24628, 24671, 24603, 24706}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			vid, err := c.GameVideos.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(vid, test.wantGameVideos) {
				t.Errorf("got: <%v>, \nwant: <%v>", vid, test.wantGameVideos)
			}
		})
	}
}

func TestGameVideoService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVideoList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVideo, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name           string
		file           string
		opts           []Option
		wantGameVideos []*GameVideo
		wantErr        error
	}{
		{"Valid response", testGameVideoList, []Option{SetLimit(5)}, init, nil},
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

			vid, err := c.GameVideos.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(vid, test.wantGameVideos) {
				t.Errorf("got: <%v>, \nwant: <%v>", vid, test.wantGameVideos)
			}
		})
	}
}

func TestGameVideoService_Count(t *testing.T) {
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

			count, err := c.GameVideos.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestGameVideoService_Fields(t *testing.T) {
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

			fields, err := c.GameVideos.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
