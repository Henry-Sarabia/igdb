package igdb

import (
	"net/http"
	"testing"
)

func TestTitlesGet(t *testing.T) {
	var titleTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/titles_get.txt", 16549, ""},
		{"Invalid ID", "test_data/empty.txt", -15000, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 16549, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range titleTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ti, err := c.Titles.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Dev. Support Engineer"
			an := ti.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eID := 16549
			aID := ti.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eURL := URL("https://www.igdb.com/titles/dev-support-engineer")
			aURL := ti.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{1000, 1001}
			agID := ti.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestTitlesList(t *testing.T) {
	var titleTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/titles_list.txt", []int{7839, 25381}, []FuncOption{SetLimit(5)}, ""},
		{"Zero IDs", "test_data/titles_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-25000}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{7839, 25381}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{7839, 25381}, []FuncOption{SetOffset(99999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range titleTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ti, err := c.Titles.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(ti)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Web Dev"
			an := ti[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			es := "web-dev"
			as := ti[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			eu := 1472328870944
			au := ti[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			egID := []int{556, 15894, 105, 11582}
			agID := ti[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestTitlesSearch(t *testing.T) {
	var titleTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/titles_search.txt", "graphic", []FuncOption{SetLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{SetLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "graphic", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "graphic", []FuncOption{SetOffset(99999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range titleTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ti, err := c.Titles.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(ti)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			ec := 1437825411943
			ac := ti[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eURL := URL("https://www.igdb.com/titles/senior-graphic-and-interface-designer")
			aURL := ti[0].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "graphic-coordinator"
			as := ti[1].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			egl := 2
			agl := len(ti[1].Games)
			if agl != egl {
				t.Errorf("Expected Games lengti %d, got %d", egl, agl)
			}

			en := "Character Graphic"
			an := ti[2].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eu := 1433023454932
			au := ti[2].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}
		})
	}
}

func TestTitlesCount(t *testing.T) {
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

			count, err := c.Titles.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestTitlesListFields(t *testing.T) {
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

			fields, err := c.Titles.ListFields()
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
