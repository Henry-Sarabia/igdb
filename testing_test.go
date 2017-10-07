package igdb

import (
	"testing"
)

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
