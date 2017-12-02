package igdb

import (
	"net/http"
	"testing"
)

func TestKeywordsGet(t *testing.T) {
	var keywordTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/keywords_get.txt", 2107, ""},
		{"Invalid ID", "test_data/empty.txt", -1000, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 2107, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range keywordTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			kw, err := c.Keywords.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "space adventure"
			an := kw.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eID := 2107
			aID := kw.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eURL := URL("https://www.igdb.com/categories/space-adventure")
			aURL := kw.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{8506, 26187, 25919, 23905, 23908, 27903, 52205}
			agID := kw.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestKeywordsList(t *testing.T) {
	var keywordTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/keywords_list.txt", []int{2096, 1108}, []FuncOption{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/keywords_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-2000}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{2096, 1108}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{2096, 1108}, []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range keywordTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			kw, err := c.Keywords.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(kw)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "humor"
			an := kw[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eURL := URL("https://www.igdb.com/categories/humor")
			aURL := kw[0].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			eu := 1403518560769
			au := kw[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			egID := []int{6749, 7591, 10666, 14952, 36899}
			agID := kw[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestKeywordsSearch(t *testing.T) {
	var keywordTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/keywords_search.txt", "strategy", []FuncOption{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "strategy", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "strategy", []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range keywordTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			kw, err := c.Keywords.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(kw)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 3782
			aID := kw[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "strategy"
			an := kw[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := 1499532005267
			ac := kw[1].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eURL := URL("https://www.igdb.com/categories/historical-strategy")
			aURL := kw[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "real-time-strategy--1"
			as := kw[2].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			egID := []int{21620, 27254, 21221, 24273, 27448, 54723}
			agID := kw[2].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestKeywordsCount(t *testing.T) {
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

			count, err := c.Keywords.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestKeywordsListFields(t *testing.T) {
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

			fields, err := c.Keywords.ListFields()
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
