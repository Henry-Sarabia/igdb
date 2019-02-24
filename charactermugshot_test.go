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
	testCharacterMugshotGet  string = "test_data/charactermugshot_get.json"
	testCharacterMugshotList string = "test_data/charactermugshot_list.json"
)

func TestCharacterMugshotService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testCharacterMugshotGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*CharacterMugshot, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                 string
		file                 string
		id                   int
		opts                 []Option
		wantCharacterMugshot *CharacterMugshot
		wantErr              error
	}{
		{"Valid response", testCharacterMugshotGet, 3600, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 3600, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 3600, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			mug, err := c.CharacterMugshots.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(mug, test.wantCharacterMugshot) {
				t.Errorf("got: <%v>, \nwant: <%v>", mug, test.wantCharacterMugshot)
			}
		})
	}
}

func TestCharacterMugshotService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testCharacterMugshotList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*CharacterMugshot, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                  string
		file                  string
		ids                   []int
		opts                  []Option
		wantCharacterMugshots []*CharacterMugshot
		wantErr               error
	}{
		{"Valid response", testCharacterMugshotList, []int{3649, 3687, 3823, 3863, 3631}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{3649, 3687, 3823, 3863, 3631}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{3649, 3687, 3823, 3863, 3631}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			mug, err := c.CharacterMugshots.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(mug, test.wantCharacterMugshots) {
				t.Errorf("got: <%v>, \nwant: <%v>", mug, test.wantCharacterMugshots)
			}
		})
	}
}

func TestCharacterMugshotService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testCharacterMugshotList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*CharacterMugshot, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                  string
		file                  string
		opts                  []Option
		wantCharacterMugshots []*CharacterMugshot
		wantErr               error
	}{
		{"Valid response", testCharacterMugshotList, []Option{SetLimit(5)}, init, nil},
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

			mug, err := c.CharacterMugshots.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(mug, test.wantCharacterMugshots) {
				t.Errorf("got: <%v>, \nwant: <%v>", mug, test.wantCharacterMugshots)
			}
		})
	}
}

func TestCharacterMugshotService_Count(t *testing.T) {
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

			count, err := c.CharacterMugshots.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestCharacterMugshotService_Fields(t *testing.T) {
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

			fields, err := c.CharacterMugshots.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
