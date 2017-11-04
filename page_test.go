package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPageTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	p := Page{}
	typ := reflect.ValueOf(p).Type()

	err := c.validateStruct(typ, PageEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPage(t *testing.T) {
	var pageTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_page.txt", 8, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 8, errEndOfJSON.Error()},
	}
	for _, tt := range pageTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.GetPage(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "IGN"
			an := pg.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ep := 2
			ap := pg.PageFollowCount
			if ap != ep {
				t.Errorf("Expected ID %d, got %d", ep, ap)
			}

			eyt := "https://www.youtube.com/ign"
			ayt := pg.Youtube
			if ayt != eyt {
				t.Errorf("Expected URL '%s', got '%s'", eyt, ayt)
			}

			ew := 1920
			aw := pg.Background.Width
			if aw != ew {
				t.Errorf("Expected width of %d, got %d", ew, aw)
			}
		})
	}
}

func TestGetPages(t *testing.T) {
	var pageTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_pages.txt", []int{36, 215}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-50}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{36, 215}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{36, 215}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range pageTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.GetPages(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(pg)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "TotalBiscuit, The Cynical Brit"
			an := pg[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eyt := "https://www.youtube.com/user/TotalHalibut"
			ayt := pg[0].Youtube
			if ayt != eyt {
				t.Errorf("Expected URL '%s', got '%s'", eyt, ayt)
			}

			eu := 1488287514804
			au := pg[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eh := 240
			ah := pg[1].Logo.Height
			if ah != eh {
				t.Errorf("Expected height of %d, got %d", eh, ah)
			}
		})
	}
}

func TestSearchPages(t *testing.T) {
	var pageTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_pages.txt", "PC", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_pages.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "PC", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "PC", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range pageTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.SearchPages(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(pg)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 133
			aID := pg[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ec := CountryCode(826)
			ac := pg[0].Country
			if ac != ec {
				t.Errorf("Expected country code %d, got %d", ec, ac)
			}

			efb := "https://www.facebook.com/Frankieonpcin1080p"
			afb := pg[0].Facebook
			if afb != efb {
				t.Errorf("Expected Facebook info '%s', got '%s'", efb, afb)
			}

			eCat := 2
			aCat := pg[1].Category
			if aCat != eCat {
				t.Errorf("Expected category %d, got %d", eCat, aCat)
			}

			ed := "The global authority on PC games. "
			ad := pg[1].Description
			if ad != ed {
				t.Errorf("Expected description '%s', got '%s'", ed, ad)
			}

			etw := "https://twitter.com/pcgamer"
			atw := pg[1].Twitter
			if atw != etw {
				t.Errorf("Expected Twitter info '%s', got '%s'", etw, atw)
			}
		})
	}
}
