package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGenreTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	g := Genre{}
	typ := reflect.ValueOf(g).Type()

	err := c.validateStruct(typ, GenreEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetGenre(t *testing.T) {
	var genreTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_genre.txt", 8, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 8, errEndOfJSON.Error()},
	}
	for _, tt := range genreTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.GetGenre(tt.ID)
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

func TestGetGenres(t *testing.T) {
	var genreTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_genres.txt", []int{5, 10}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-5}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{5, 10}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{5, 10}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range genreTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.GetGenres(tt.IDs, tt.Opts...)
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

func TestSearchGenres(t *testing.T) {
	var genreTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_genres.txt", "strategy", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_genres.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "strategy", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "strategy", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range genreTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.SearchGenres(tt.Qry, tt.Opts...)
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
