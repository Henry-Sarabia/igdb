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
	testArtworkGet  string = "test_data/artwork_get.json"
	testArtworkList string = "test_data/artwork_list.json"
)

func TestArtworkService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testArtworkGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Artwork, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name        string
		file        string
		id          int
		opts        []FuncOption
		wantArtwork *Artwork
		wantErr     error
	}{
		{"Valid response", testArtworkGet, 1336, []FuncOption{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 1336, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 1336, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			art, err := c.Artworks.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(art, test.wantArtwork) {
				t.Errorf("got: <%v>, \nwant: <%v>", art, test.wantArtwork)
			}
		})
	}
}

func TestArtworkService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testArtworkList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Artwork, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name         string
		file         string
		ids          []int
		opts         []FuncOption
		wantArtworks []*Artwork
		wantErr      error
	}{
		{"Valid response", testArtworkList, []int{1721, 2777}, []FuncOption{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1721, 2777}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1721, 2777}, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			art, err := c.Artworks.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(art, test.wantArtworks) {
				t.Errorf("got: <%v>, \nwant: <%v>", art, test.wantArtworks)
			}
		})
	}
}

func TestArtworkService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testArtworkList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Artwork, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name         string
		file         string
		opts         []FuncOption
		wantArtworks []*Artwork
		wantErr      error
	}{
		{"Valid response", testArtworkList, []FuncOption{SetLimit(5)}, init, nil},
		{"Empty response", testFileEmpty, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			art, err := c.Artworks.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(art, test.wantArtworks) {
				t.Errorf("got: <%v>, \nwant: <%v>", art, test.wantArtworks)
			}
		})
	}
}

func TestArtworkService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []FuncOption
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []FuncOption{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []FuncOption{SetLimit(100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.Artworks.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestArtworkService_Fields(t *testing.T) {
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

			fields, err := c.Artworks.Fields()
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
