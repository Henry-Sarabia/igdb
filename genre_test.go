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
	testGenreGet  string = "test_data/genre_get.json"
	testGenreList string = "test_data/genre_list.json"
)

func TestGenreService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testGenreGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Genre, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name      string
		file      string
		id        int
		opts      []Option
		wantGenre *Genre
		wantErr   error
	}{
		{"Valid response", testGenreGet, 13, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 13, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 13, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			gen, err := c.Genres.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(gen, test.wantGenre) {
				t.Errorf("got: <%v>, \nwant: <%v>", gen, test.wantGenre)
			}
		})
	}
}

func TestGenreService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testGenreList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Genre, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name       string
		file       string
		ids        []int
		opts       []Option
		wantGenres []*Genre
		wantErr    error
	}{
		{"Valid response", testGenreList, []int{24, 26, 4, 15, 31}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{26, 4, 15, 31}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{26, 4, 15, 31}, []Option{SetOffset(-99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			gen, err := c.Genres.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(gen, test.wantGenres) {
				t.Errorf("got: <%v>, \nwant: <%v>", gen, test.wantGenres)
			}
		})
	}
}

func TestGenreService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testGenreList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Genre, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name       string
		file       string
		opts       []Option
		wantGenres []*Genre
		wantErr    error
	}{
		{"Valid response", testGenreList, []Option{SetLimit(5)}, init, nil},
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

			gen, err := c.Genres.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(gen, test.wantGenres) {
				t.Errorf("got: <%v>, \nwant: <%v>", gen, test.wantGenres)
			}
		})
	}
}

func TestGenreService_Count(t *testing.T) {
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

			count, err := c.Genres.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestGenreService_Fields(t *testing.T) {
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

			fields, err := c.Genres.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
