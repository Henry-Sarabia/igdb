package igdb

import (
	"net/http"
	"testing"
)

func TestCompaniesGet(t *testing.T) {
	var companyTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/companies_get.txt", 58, ""},
		{"Invalid ID", "test_data/empty.txt", -999, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 58, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range companyTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.Companies.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Mojang AB"
			an := com.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eURL := URL("https://www.igdb.com/companies/mojang-ab")
			aURL := com.URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			eID := []int{2600, 1898, 121, 18977, 8339}
			aID := com.Developed
			for i := range aID {
				if aID[i] != eID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
				}
			}
		})
	}
}

func TestCompaniesList(t *testing.T) {
	var companyTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/companies_list.txt", []int{854, 7260}, []FuncOption{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/companies_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-400}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{854, 7260}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{854, 7260}, []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range companyTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.Companies.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(com)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Playdead"
			an := com[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eu := 1504811027097
			au := com[0].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eURL := URL("https://www.igdb.com/companies/night-school-studio")
			aURL := com[1].URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			eID := []int{14587, 22748, 36858}
			aID := com[1].Developed
			for i := range aID {
				if aID[i] != eID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
				}
			}
		})
	}
}

func TestCompaniesSearch(t *testing.T) {
	var companyTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/companies_search.txt", "toby fox", []FuncOption{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "toby fox", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "toby fox", []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range companyTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.Companies.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(com)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 6545
			aID := com[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eu := 1500415107616
			au := com[0].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			es := "terrible-toybox"
			as := com[1].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			ePub := 10232
			aPub := com[1].Published[0]
			if aPub != ePub {
				t.Errorf("Expected Game ID %d, got %d", ePub, aPub)
			}
		})
	}
}

func TestCompaniesCount(t *testing.T) {
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

			count, err := c.Companies.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestCompaniesListFields(t *testing.T) {
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

			fields, err := c.Companies.ListFields()
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
