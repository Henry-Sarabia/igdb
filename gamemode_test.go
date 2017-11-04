package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGameModeTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	g := GameMode{}
	typ := reflect.ValueOf(g).Type()

	err := c.validateStruct(typ, GameModeEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetGameMode(t *testing.T) {
	var gameModeTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_gamemode.txt", 1, ""},
		{"Invalid ID", "test_data/empty.txt", -100, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 1, errEndOfJSON.Error()},
	}
	for _, tt := range gameModeTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.GetGameMode(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Single player"
			an := g.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eURL := URL("https://www.igdb.com/game_modes/single-player")
			aURL := g.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{37, 1, 2, 60, 54, 6, 38, 3, 5, 35, 36}
			agID := g.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestGetGameModes(t *testing.T) {
	var gameModeTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_gamemodes.txt", []int{3, 4}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-100}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{3, 4}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{3, 4}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range gameModeTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.GetGameModes(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(g)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Co-operative"
			an := g[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eURL := URL("https://www.igdb.com/game_modes/co-operative")
			aURL := g[0].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "split-screen"
			as := g[1].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			egID := []int{141, 143, 176, 211, 95, 545, 642, 847, 784}
			agID := g[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestSearchGameModes(t *testing.T) {
	var gameModeTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_gamemodes.txt", "multiplayer", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_gamemodes.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "multiplayer", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "multiplayer", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range gameModeTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.SearchGameModes(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(g)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 2
			aID := g[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Multiplayer"
			an := g[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := 1298968853000
			ac := g[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eURL := URL("https://www.igdb.com/game_modes/massively-multiplayer-online-mmo")
			aURL := g[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "massively-multiplayer-online-mmo"
			as := g[1].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			egID := []int{114, 123, 206, 145, 215, 228, 227, 229, 270, 282}
			agID := g[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}
