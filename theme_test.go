package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestThemeTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	r := Theme{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, ThemeEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestThemesGet(t *testing.T) {
	var themeTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/themes_get.txt", 17, ""},
		{"Invalid ID", "test_data/empty.txt", -15, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 17, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range themeTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			th, err := c.Themes.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Fantasy"
			an := th.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eID := 17
			aID := th.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eURL := URL("https://www.igdb.com/themes/fantasy")
			aURL := th.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{799, 651, 901, 929, 939, 800, 931, 876}
			agID := th.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestThemesList(t *testing.T) {
	var themeTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/themes_list.txt", []int{20, 23}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-50}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/themes_list.txt", nil, nil, ""},
		{"Empty response", "test_data/empty.txt", []int{20, 23}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{20, 23}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range themeTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			th, err := c.Themes.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(th)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Thriller"
			an := th[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			es := "thriller"
			as := th[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			eu := 1323289216000
			au := th[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			egID := []int{4, 820, 43, 500, 501, 433, 250, 377, 375}
			agID := th[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestThemesSearch(t *testing.T) {
	var themeTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/themes_search.txt", "horror", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []OptionFunc{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "horror", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "horror", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range themeTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			th, err := c.Themes.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 1
			al := len(th)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			ec := 1322605338000
			ac := th[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eURL := URL("https://www.igdb.com/themes/horror")
			aURL := th[0].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "horror"
			as := th[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			egl := 7
			agl := len(th[0].Games)
			if agl != egl {
				t.Errorf("Expected Games length %d, got %d", egl, agl)
			}
		})
	}
}

func TestThemesCount(t *testing.T) {
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

			count, err := c.Themes.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestThemesListFields(t *testing.T) {
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

			fields, err := c.Themes.ListFields()
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
