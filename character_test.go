package igdb

import (
	"net/http"
	"testing"
)

func TestCharactersGet(t *testing.T) {
	var characterTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/characters_get.txt", 10617, ""},
		{"Invalid ID", "test_data/empty.txt", -500, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 10617, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range characterTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()
			ch, err := c.Characters.Get(tt.ID)
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

func TestCharactersList(t *testing.T) {
	var characterTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/characters_list.txt", []int{3726, 9580}, []FuncOption{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/characters_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-500}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{3726, 9580}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{3726, 9580}, []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{3726, 9580}, nil, ErrNoResults.Error()},
	}
	for _, tt := range characterTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ch, err := c.Characters.List(tt.IDs, tt.Opts...)
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

func TestCharactersSearch(t *testing.T) {
	var characterTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/characters_search.txt", "snake", []FuncOption{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "snake", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "snake", []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "snake", nil, ErrNoResults.Error()},
	}
	for _, tt := range characterTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ch, err := c.Characters.Search(tt.Qry, tt.Opts...)
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

func TestCharactersCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Resp     string
		Opts     []FuncOption
		ExpCount int
		ExpErr   string
	}{
		{"Happy path", `{"count": 100}`, []FuncOption{OptFilter("popularity", OpGreaterThan, "75")}, 100, ""},
		{"Empty response", "", nil, 0, errEndOfJSON.Error()},
		{"Invalid option", "", []FuncOption{OptLimit(100)}, 0, ErrOutOfRange.Error()},
		{"No results", "[]", nil, 0, ErrNoResults.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			count, err := c.Characters.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestCharactersListFields(t *testing.T) {
	var fieldTests = []struct {
		Name      string
		Resp      string
		ExpFields []string
		ExpErr    string
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, ""},
		{"Dot operator", `["logo.url", "background.id"]`, []string{"background.id", "logo.url"}, ""},
		{"Asterisk", `["*"]`, []string{"*"}, ""},
		{"Empty response", "", nil, errEndOfJSON.Error()},
		{"No results", "[]", nil, ""},
	}

	for _, tt := range fieldTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			fields, err := c.Characters.ListFields()
			assertError(t, err, tt.ExpErr)

			ok, err := equalSlice(fields, tt.ExpFields)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Fatalf("Expected fields '%v', got '%v'", tt.ExpFields, fields)
			}
		})
	}
}
