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
	testPlatformLogoGet  string = "test_data/platformlogo_get.json"
	testPlatformLogoList string = "test_data/platformlogo_list.json"
)

func TestPlatformLogoService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformLogoGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformLogo, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name             string
		file             string
		id               int
		opts             []Option
		wantPlatformLogo *PlatformLogo
		wantErr          error
	}{
		{"Valid response", testPlatformLogoGet, 61, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 61, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 61, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			logo, err := c.PlatformLogos.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(logo, test.wantPlatformLogo) {
				t.Errorf("got: <%v>, \nwant: <%v>", logo, test.wantPlatformLogo)
			}
		})
	}
}

func TestPlatformLogoService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformLogoList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformLogo, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name              string
		file              string
		ids               []int
		opts              []Option
		wantPlatformLogos []*PlatformLogo
		wantErr           error
	}{
		{"Valid response", testPlatformLogoList, []int{32, 23, 41, 49, 34}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{32, 23, 41, 49, 34}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{32, 23, 41, 49, 34}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			logo, err := c.PlatformLogos.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(logo, test.wantPlatformLogos) {
				t.Errorf("got: <%v>, \nwant: <%v>", logo, test.wantPlatformLogos)
			}
		})
	}
}

func TestPlatformLogoService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testPlatformLogoList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*PlatformLogo, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name              string
		file              string
		opts              []Option
		wantPlatformLogos []*PlatformLogo
		wantErr           error
	}{
		{"Valid response", testPlatformLogoList, []Option{SetLimit(5)}, init, nil},
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

			logo, err := c.PlatformLogos.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(logo, test.wantPlatformLogos) {
				t.Errorf("got: <%v>, \nwant: <%v>", logo, test.wantPlatformLogos)
			}
		})
	}
}

func TestPlatformLogoService_Count(t *testing.T) {
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

			count, err := c.PlatformLogos.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)

			}
		})
	}
}

func TestPlatformLogoService_Fields(t *testing.T) {
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

			fields, err := c.PlatformLogos.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}
