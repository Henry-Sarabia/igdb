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
	testFeedFollowGet  string = "test_data/feedfollow_get.json"
	testFeedFollowList string = "test_data/feedfollow_list.json"
)

func TestFeedFollowService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testFeedFollowGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*FeedFollow, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name           string
		file           string
		id             int
		opts           []Option
		wantFeedFollow *FeedFollow
		wantErr        error
	}{
		{"Valid response", testFeedFollowGet, 1111, []Option{SetFields("name")}, init[0], nil},
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

			ff, err := c.FeedFollows.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ff, test.wantFeedFollow) {
				t.Errorf("got: <%v>, \nwant: <%v>", ff, test.wantFeedFollow)
			}
		})
	}
}

func TestFeedFollowService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testFeedFollowList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*FeedFollow, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name            string
		file            string
		ids             []int
		opts            []Option
		wantFeedFollows []*FeedFollow
		wantErr         error
	}{
		{"Valid response", testFeedFollowList, []int{1111, 2221}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1111, 2221}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1111, 2221}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ff, err := c.FeedFollows.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ff, test.wantFeedFollows) {
				t.Errorf("got: <%v>, \nwant: <%v>", ff, test.wantFeedFollows)
			}
		})
	}
}

func TestFeedFollowService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testFeedFollowList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*FeedFollow, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name            string
		file            string
		opts            []Option
		wantFeedFollows []*FeedFollow
		wantErr         error
	}{
		{"Valid response", testFeedFollowList, []Option{SetLimit(5)}, init, nil},
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

			ff, err := c.FeedFollows.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ff, test.wantFeedFollows) {
				t.Errorf("got: <%v>, \nwant: <%v>", ff, test.wantFeedFollows)
			}
		})
	}
}

func TestFeedFollowService_Count(t *testing.T) {
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

			count, err := c.FeedFollows.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestFeedFollowService_Fields(t *testing.T) {
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

			fields, err := c.FeedFollows.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
