package igdb

import (
	"net/http"
	"testing"
)

func TestEnginesGet(t *testing.T) {
	var engineTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/engines_get.txt", 26, ""},
		{"Invalid ID", "test_data/empty.txt", -100, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 26, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range engineTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			eng, err := c.Engines.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "RAGE"
			an := eng.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ew := 476
			aw := eng.Logo.Width
			if aw != ew {
				t.Errorf("Expected width of %d, got %d", ew, aw)
			}

			egID := []int{731, 434, 7071, 1020, 960, 2541, 3174, 3265, 1969}
			agID := eng.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestEnginesList(t *testing.T) {
	var engineTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/engines_list.txt", []int{9, 22}, []FuncOption{SetLimit(5)}, ""},
		{"Zero IDs", "test_data/engines_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-999}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{9, 22}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{9, 22}, []FuncOption{SetOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range engineTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			eng, err := c.Engines.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(eng)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Anvil"
			an := eng[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eURL := URL("https://www.igdb.com/game_engines/anvil")
			aURL := eng[0].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			ecID := "pbscffthi6uhqxs3ubjk"
			acID := eng[1].Logo.ID
			if acID != ecID {
				t.Errorf("Expected Cloudinary ID '%s', got '%s'", ecID, acID)
			}

			egID := []int{7327, 981, 1968, 4756, 14533, 19726}
			agID := eng[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestEnginesSearch(t *testing.T) {
	var engineTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/engines_search.txt", "tool", []FuncOption{SetLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{SetLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "tool", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "tool", []FuncOption{SetOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range engineTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			eng, err := c.Engines.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(eng)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Telltale Tool"
			an := eng[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eu := 1492514717250
			au := eng[0].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eURL := URL("https://www.igdb.com/game_engines/crystal-tools")
			aURL := eng[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{389, 384, 2449}
			agID := eng[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestEnginesCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Resp     string
		Opts     []FuncOption
		ExpCount int
		ExpErr   string
	}{
		{"Happy path", `{"count": 100}`, []FuncOption{SetFilter("popularity", OpGreaterThan, "75")}, 100, ""},
		{"Empty response", "", nil, 0, errEndOfJSON.Error()},
		{"Invalid option", "", []FuncOption{SetLimit(100)}, 0, ErrOutOfRange.Error()},
		{"No results", "[]", nil, 0, ErrNoResults.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			count, err := c.Engines.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestEnginesListFields(t *testing.T) {
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

			fields, err := c.Engines.ListFields()
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
