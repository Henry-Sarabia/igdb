package igdb

import (
	"net/http"
	"testing"
)

func TestPagesGet(t *testing.T) {
	var pageTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/pages_get.txt", 8, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 8, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range pageTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.Pages.Get(tt.ID)
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

func TestPagesList(t *testing.T) {
	var pageTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/pages_list.txt", []int{36, 215}, []FuncOption{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/pages_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-50}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{36, 215}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{36, 215}, []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range pageTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.Pages.List(tt.IDs, tt.Opts...)
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

func TestPagesSearch(t *testing.T) {
	var pageTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/pages_search.txt", "PC", []FuncOption{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "PC", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "PC", []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range pageTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.Pages.Search(tt.Qry, tt.Opts...)
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

func TestPagesCount(t *testing.T) {
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

			count, err := c.Pages.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestPagesListFields(t *testing.T) {
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

			fields, err := c.Pages.ListFields()
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
