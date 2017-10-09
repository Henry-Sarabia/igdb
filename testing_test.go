package igdb

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

type testStructNoTags struct {
	ID int
}

type testStructOneTag struct {
	ID int `json:"id"`
}

type testStructManyTags struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

func TestValidateStruct(t *testing.T) {
	noTags := reflect.ValueOf(testStructNoTags{}).Type()
	oneTag := reflect.ValueOf(testStructOneTag{}).Type()
	manyTags := reflect.ValueOf(testStructManyTags{}).Type()
	notStruct := reflect.ValueOf("im not a struct").Type()

	validateTests := []struct {
		Name   string
		Struct reflect.Type
		Resp   string
		ExpErr error
	}{
		{"No tag struct with corresponding tag response", noTags, `[]`, nil},
		{"No tag struct without corresponding tag response", noTags, `["id"]`, errors.New("missing struct tags: id")},
		{"One tag struct with corresponding tag strings", oneTag, `["id"]`, nil},
		{"One tag struct without corresponding tag strings", oneTag, `["id", "url"]`, errors.New("missing struct tags: url")},
		{"Many tag struct with corresponding tag strings", manyTags, `["id", "name", "rating"]`, nil},
		{"Many tag struct without corresponding tag strings", manyTags, `["id", "name", "rating", "url"]`, errors.New("missing struct tags: url")},
		{"Non-struct type", notStruct, `[]`, ErrNotStruct},
	}

	for _, vt := range validateTests {
		t.Run(vt.Name, func(t *testing.T) {
			ts, c := startTestServer(http.StatusOK, vt.Resp)
			defer ts.Close()

			err := c.validateStruct(vt.Struct, testEndpoint)
			if !reflect.DeepEqual(err, vt.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", vt.ExpErr, err)
			}
		})
	}

	networkTests := []struct {
		Name   string
		Status int
		Resp   string
		ExpErr string
	}{
		{"Bad request with empty response", http.StatusBadRequest, "", errEndOfJSON.Error()},
		{"Bad request with Error type response", http.StatusBadRequest, `{"status": 400,"message": "status bad request"}`, "Status 400 - status bad request"},
		{"OK request with empty response", http.StatusOK, "", errEndOfJSON.Error()},
	}

	for _, nt := range networkTests {
		t.Run(nt.Name, func(t *testing.T) {
			ts, c := startTestServer(nt.Status, nt.Resp)
			defer ts.Close()

			err := c.validateStruct(manyTags, testEndpoint)
			if err == nil {
				if nt.ExpErr != "" {
					t.Fatalf("Expected error '%v', got nil error", nt.ExpErr)
				}
				return
			} else if err.Error() != nt.ExpErr {
				t.Fatalf("Expected error '%v', got error '%v'", nt.ExpErr, err.Error())
			}
		})
	}
}

func TestValidateStructTags(t *testing.T) {
	noTags := reflect.ValueOf(testStructNoTags{}).Type()
	oneTag := reflect.ValueOf(testStructOneTag{}).Type()
	manyTags := reflect.ValueOf(testStructManyTags{}).Type()
	notStruct := reflect.ValueOf("im not a struct").Type()

	validateTests := []struct {
		Name   string
		Struct reflect.Type
		Tags   []string
		ExpErr error
	}{
		{"No tag struct with corresponding tag strings", noTags, nil, nil},
		{"No tag struct without corresponding tag strings", noTags, []string{"id"}, errors.New("missing struct tags: id")},
		{"One tag struct with corresponding tag strings", oneTag, []string{"id"}, nil},
		{"One tag struct without corresponding tag strings", oneTag, []string{"id", "url"}, errors.New("missing struct tags: url")},
		{"Many tag struct with corresponding tag strings", manyTags, []string{"id", "name", "rating"}, nil},
		{"Many tag struct without corresponding tag strings", manyTags, []string{"id", "name", "rating", "url"}, errors.New("missing struct tags: url")},
		{"Non-struct type", notStruct, nil, ErrNotStruct},
	}

	for _, vt := range validateTests {
		t.Run(vt.Name, func(t *testing.T) {
			err := validateStructTags(vt.Struct, vt.Tags)
			if !reflect.DeepEqual(err, vt.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", vt.ExpErr, err)
			}
		})
	}
}

func TestGetStructTags(t *testing.T) {
	noTags := reflect.ValueOf(testStructNoTags{}).Type()
	oneTag := reflect.ValueOf(testStructOneTag{}).Type()
	manyTags := reflect.ValueOf(testStructManyTags{}).Type()
	notStruct := reflect.ValueOf("im not a struct").Type()

	tagTests := []struct {
		Name    string
		Struct  reflect.Type
		ExpTags []string
		ExpErr  error
	}{
		{"Struct type with no tags", noTags, nil, nil},
		{"Struct type with one tag", oneTag, []string{"id"}, nil},
		{"Struct type with many tags", manyTags, []string{"id", "name", "rating"}, nil},
		{"Non-struct type", notStruct, nil, ErrNotStruct},
	}
	for _, tt := range tagTests {
		t.Run(tt.Name, func(t *testing.T) {
			tags, err := getStructTags(tt.Struct)
			if !reflect.DeepEqual(err, tt.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", tt.ExpErr, err)
			}

			ok, err := equalSlice(tags, tt.ExpTags)
			if err != nil {
				t.Fatal(err)
			}

			if !ok {
				t.Fatalf("Expected tags '%v', got '%v'", tt.ExpTags, tags)
			}
		})
	}
}

func TestRemoveSubFields(t *testing.T) {
	removeTests := []struct {
		Name      string
		Fields    []string
		ExpFields []string
	}{
		{"Empty fields slice", nil, nil},
		{"Single empty field", []string{""}, nil},
		{"Multiple empty fields", []string{"", "", ""}, nil},
		{"Single non-empty field", []string{"name"}, []string{"name"}},
		{"Multiple non-empty fields", []string{"name", "rating"}, []string{"name", "rating"}},
		{"Mixed empty and non-empty fields", []string{"game", "", "person"}, []string{"game", "person"}},
		{"Out of order mixed empty and non-empty fields", []string{"game", "", "person"}, []string{"person", "game"}},
		{"Single subfield", []string{"game.id"}, nil},
		{"Multiple subfields", []string{"game.id", "game.rating"}, nil},
		{"Mixed non-empty fields and subfields", []string{"game.name", "game", "franchise", "franchise.developer"}, []string{"game", "franchise"}},
		{"Mixed empty fields, non-empty fields, and subfields", []string{"", "genre", "genre.games"}, []string{"genre"}},
	}
	for _, rt := range removeTests {
		t.Run(rt.Name, func(t *testing.T) {
			actFields := removeSubfields(rt.Fields)
			ok, err := equalSlice(actFields, rt.ExpFields)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Fatalf("Expected fields '%v', got '%v'", rt.ExpFields, actFields)
			}
		})
	}
}
