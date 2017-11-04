package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCharacterTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	ch := Character{}
	typ := reflect.ValueOf(ch).Type()

	err := c.validateStruct(typ, CharacterEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCharacter(t *testing.T) {
	var characterTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_character.txt", 10617, ""},
		{"Invalid ID", "test_data/empty.txt", -500, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 10617, errEndOfJSON.Error()},
	}
	for _, tt := range characterTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ch, err := c.GetCharacter(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			eID := tt.ID
			aID := ch.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Princess Zelda"
			an := ch.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'\n", en, an)
			}

			eh := 479
			ah := ch.Mugshot.Height
			if ah != eh {
				t.Errorf("Expected height of %d, got %d\n", eh, ah)
			}
		})
	}
}

func TestGetCharacters(t *testing.T) {
	var characterTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_characters.txt", []int{3726, 9580}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-500}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{3726, 9580}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{3726, 9580}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range characterTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ch, err := c.GetCharacters(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(ch)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Mario"
			an := ch[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eAKA := "Super Mario"
			aAKA := ch[0].AKAs[0]
			if eAKA != aAKA {
				t.Errorf("Expected AKA '%s', got '%s'", eAKA, aAKA)
			}

			eID := 9580
			aID := ch[1].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eURL := URL("https://www.igdb.com/characters/samus-aran")
			aURL := ch[1].URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

		})
	}

}

func TestSearchCharacters(t *testing.T) {
	var characterTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_characters.txt", "snake", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_characters.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "snake", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "snake", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range characterTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ch, err := c.SearchCharacters(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(ch)
			if al != el {
				t.Errorf("Expected length of %d, got %d\n", el, al)
			}

			eURL := URL("https://www.igdb.com/characters/snake")
			aURL := ch[0].URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'\n", eURL, aURL)
			}

			en := "Solid Snake"
			an := ch[1].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'\n", en, an)
			}

			eID := 5378
			aID := ch[2].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d\n", eID, aID)
			}
		})
	}
}
