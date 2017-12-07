package igdb

import (
	"net/http"
	"testing"
)

func TestPerspectivesGet(t *testing.T) {
	var perspectiveTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/perspectives_get.txt", 7, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 7, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range perspectiveTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Perspectives.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Virtual Reality"
			an := p.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eu := 1462288484243
			au := p.UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eURL := URL("https://www.igdb.com/player_perspectives/virtual-reality")
			aURL := p.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{8654, 10724, 6415, 5639, 9254, 17244, 18157, 17986, 12302}
			agID := p.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestPerspectivesList(t *testing.T) {
	var perspectiveTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/perspectives_list.txt", []int{6, 3}, []FuncOption{SetLimit(5)}, ""},
		{"Zero IDs", "test_data/perspectives_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-10}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{6, 3}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{6, 3}, []FuncOption{SetOffset(99999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range perspectiveTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Perspectives.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Aural"
			an := p[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			es := "aural"
			as := p[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			eID := 3
			aID := p[1].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			egID := []int{6, 5, 35, 36, 13, 14, 17, 12, 76}
			agID := p[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestPerspectivesSearch(t *testing.T) {
	var perspectiveTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/perspectives_search.txt", "person", []FuncOption{SetLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{SetLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "person", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "person", []FuncOption{SetOffset(99999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existent entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range perspectiveTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Perspectives.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 1
			aID := p[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ec := 1298968658000
			ac := p[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eURL := URL("https://www.igdb.com/player_perspectives/third-person")
			aURL := p[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "third-person"
			as := p[1].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}
		})
	}
}

func TestPerspectivesCount(t *testing.T) {
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

			count, err := c.Perspectives.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestPerspectivesListFields(t *testing.T) {
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

			fields, err := c.Perspectives.ListFields()
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
