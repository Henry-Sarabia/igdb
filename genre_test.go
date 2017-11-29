package igdb

import (
	"net/http"
	"testing"
)

func TestGenresGet(t *testing.T) {
	var genreTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/genres_get.txt", 8, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 8, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range genreTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Genres.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Platform"
			an := g.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eID := 8
			aID := g.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eURL := URL("https://www.igdb.com/genres/platform")
			aURL := g.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{358, 360, 452, 337, 454, 185, 190, 71, 72, 217}
			agID := g.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestGenresList(t *testing.T) {
	var genreTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/genres_list.txt", []int{5, 10}, []OptionFunc{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/genres_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-5}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{5, 10}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{5, 10}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range genreTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Genres.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(g)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Shooter"
			an := g[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eURL := URL("https://www.igdb.com/genres/shooter")
			aURL := g[0].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			eu := 1323289215000
			au := g[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			egID := []int{143, 154, 177, 390, 422, 90, 91, 92, 99}
			agID := g[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestGenresSearch(t *testing.T) {
	var genreTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/genres_search.txt", "strategy", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []OptionFunc{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "strategy", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "strategy", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range genreTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Genres.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(g)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 15
			aID := g[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Strategy"
			an := g[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := 1297678340000
			ac := g[1].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eURL := URL("https://www.igdb.com/genres/turn-based-strategy-tbs")
			aURL := g[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "real-time-strategy-rts"
			as := g[2].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			egID := []int{34, 35, 36, 133, 151, 159, 289}
			agID := g[2].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestGenresCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Resp     string
		Opts     []OptionFunc
		ExpCount int
		ExpErr   string
	}{
		{"Happy path", `{"count": 100}`, []OptionFunc{OptFilter("popularity", OpGreaterThan, "75")}, 100, ""},
		{"Empty response", "", nil, 0, errEndOfJSON.Error()},
		{"Invalid option", "", []OptionFunc{OptLimit(100)}, 0, ErrOutOfRange.Error()},
		{"No results", "[]", nil, 0, ErrNoResults.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			count, err := c.Genres.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestGenresListFields(t *testing.T) {
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

			fields, err := c.Genres.ListFields()
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
