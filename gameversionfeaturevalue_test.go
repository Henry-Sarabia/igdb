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
	testGameVersionFeatureValueGet  string = "test_data/gameversionfeaturevalue_get.json"
	testGameVersionFeatureValueList string = "test_data/gameversionfeaturevalue_list.json"
)

func TestGameVersionFeatureValueService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVersionFeatureValueGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVersionFeatureValue, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                        string
		file                        string
		id                          int
		opts                        []FuncOption
		wantGameVersionFeatureValue *GameVersionFeatureValue
		wantErr                     error
	}{
		{"Valid response", testGameVersionFeatureValueGet, 489, []FuncOption{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 489, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 489, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			val, err := c.GameVersionFeatureValues.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(val, test.wantGameVersionFeatureValue) {
				t.Errorf("got: <%v>, \nwant: <%v>", val, test.wantGameVersionFeatureValue)
			}
		})
	}
}

func TestGameVersionFeatureValueService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVersionFeatureValueList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVersionFeatureValue, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name                         string
		file                         string
		ids                          []int
		opts                         []FuncOption
		wantGameVersionFeatureValues []*GameVersionFeatureValue
		wantErr                      error
	}{
		{"Valid response", testGameVersionFeatureValueList, []int{1975, 1217, 1220, 1234, 1007}, []FuncOption{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{1975, 1217, 1220, 1234, 1007}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{1975, 1217, 1220, 1234, 1007}, []FuncOption{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			val, err := c.GameVersionFeatureValues.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(val, test.wantGameVersionFeatureValues) {
				t.Errorf("got: <%v>, \nwant: <%v>", val, test.wantGameVersionFeatureValues)
			}
		})
	}
}

func TestGameVersionFeatureValueService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testGameVersionFeatureValueList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*GameVersionFeatureValue, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name                         string
		file                         string
		opts                         []FuncOption
		wantGameVersionFeatureValues []*GameVersionFeatureValue
		wantErr                      error
	}{
		{"Valid response", testGameVersionFeatureValueList, []FuncOption{SetLimit(5)}, init, nil},
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

			val, err := c.GameVersionFeatureValues.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(val, test.wantGameVersionFeatureValues) {
				t.Errorf("got: <%v>, \nwant: <%v>", val, test.wantGameVersionFeatureValues)
			}
		})
	}
}

func TestGameVersionFeatureValueService_Count(t *testing.T) {
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

			count, err := c.GameVersionFeatureValues.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestGameVersionFeatureValueService_Fields(t *testing.T) {
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

			fields, err := c.GameVersionFeatureValues.Fields()
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
