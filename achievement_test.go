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
	testAchievementGet  string = "test_data/achievement_get.json"
	testAchievementList string = "test_data/achievement_list.json"
)

func TestAchievementService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testAchievementGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Achievement, 1)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name            string
		file            string
		id              int
		opts            []Option
		wantAchievement *Achievement
		wantErr         error
	}{
		{"Valid response", testAchievementGet, 123, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 123, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 123, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ach, err := c.Achievements.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ach, test.wantAchievement) {
				t.Errorf("got: <%v>, \nwant: <%v>", ach, test.wantAchievement)
			}
		})
	}
}

func TestAchievementService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testAchievementList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Achievement, 0)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name             string
		file             string
		ids              []int
		opts             []Option
		wantAchievements []*Achievement
		wantErr          error
	}{
		{"Valid response", testAchievementList, []int{123, 456}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{123, 456}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{123, 456}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ach, err := c.Achievements.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ach, test.wantAchievements) {
				t.Errorf("got: <%v>, \nwant: <%v>", ach, test.wantAchievements)
			}
		})
	}
}

func TestAchievementService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testAchievementList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Achievement, 0)
	err = json.Unmarshal(f, &init)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name             string
		file             string
		opts             []Option
		wantAchievements []*Achievement
		wantErr          error
	}{
		{"Valid response", testAchievementList, nil, init, nil},
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

			a, err := c.Achievements.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(a, test.wantAchievements) {
				t.Errorf("got: <%v>, \nwant: <%v>", a, test.wantAchievements)
			}
		})
	}
}

func TestAchievementService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Valid response", `{"count": 100}`, []Option{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(-100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.Achievements.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestAchievementService_Fields(t *testing.T) {
	var tests = []struct {
		name       string
		resp       string
		wantFields []string
		wantErr    error
	}{
		{"Valid response", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, nil},
		{"Dot operator", `["logo.url", "background.id"]`, []string{"background.id", "logo.url"}, nil},
		{"Asterisk", `["*"]`, []string{"*"}, nil},
		{"Empty response", "", nil, errInvalidJSON},
		{"No results", "[]", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			fields, err := c.Achievements.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
