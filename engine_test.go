package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestEngineTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	e := Engine{}
	typ := reflect.ValueOf(e).Type()

	err := c.validateStruct(typ, EngineEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetEngine(t *testing.T) {
	var engineTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_engine.txt", 26, ""},
		{"Invalid ID", "test_data/empty.txt", -100, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 26, errEndOfJSON.Error()},
	}
	for _, tt := range engineTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			eng, err := c.GetEngine(tt.ID)
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

func TestGetEngines(t *testing.T) {
	var engineTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_engines.txt", []int{9, 22}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-999}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{9, 22}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{9, 22}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range engineTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			eng, err := c.GetEngines(tt.IDs, tt.Opts...)
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

func TestSearchEngines(t *testing.T) {
	var engineTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_engines.txt", "tool", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_engines.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "tool", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "tool", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range engineTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			eng, err := c.SearchEngines(tt.Qry, tt.Opts...)
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
