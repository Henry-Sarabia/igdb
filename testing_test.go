package igdb

import (
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

func TestGetStructTags(t *testing.T) {
	noTags := testStructNoTags{}
	oneTag := testStructOneTag{}
	manyTags := testStructManyTags{}
	notStruct := "im not a struct"

	tagTests := []struct {
		Name    string
		Struct  reflect.Type
		ExpTags []string
		ExpErr  error
	}{
		{"Struct type with no tags", reflect.ValueOf(noTags).Type(), nil, nil},
		{"Struct type with one tag", reflect.ValueOf(oneTag).Type(), []string{"id"}, nil},
		{"Struct type with many tags", reflect.ValueOf(manyTags).Type(), []string{"id", "name", "rating"}, nil},
		{"Non-struct type", reflect.ValueOf(notStruct).Type(), nil, ErrNotStruct},
	}
	for _, tt := range tagTests {
		t.Run(tt.Name, func(t *testing.T) {
			tags, err := getStructTags(tt.Struct)
			if !reflect.DeepEqual(err, tt.ExpErr) {
				t.Fatalf("Expecter error '%v', got '%v'", tt.ExpErr, err)
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
