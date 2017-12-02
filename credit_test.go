package igdb

import (
	"net/http"
	"testing"
)

func TestCreditsGet(t *testing.T) {
	var creditTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/credits_get.txt", 1342182279, ""},
		{"Invalid ID", "test_data/empty.txt", -321, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 1342182279, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range creditTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cr, err := c.Credits.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			eID := 1342182279
			aID := cr.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Michael"
			an := cr.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := CreditCategory(5)
			ac := cr.Category
			if ac != ec {
				t.Errorf("Expected category %d, got %d", ec, ac)
			}

			ep := 45
			ap := cr.Position
			if ap != ep {
				t.Errorf("Expected position %d, got %d", ep, ap)
			}
		})
	}
}

func TestCreditsList(t *testing.T) {
	var creditTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/credits_list.txt", []int{1342181334, 1342186852}, []FuncOption{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/credits_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-100}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{1342181334, 1342186852}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{1342181334, 1342186852}, []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range creditTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cr, err := c.Credits.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(cr)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 1342180316
			aID := cr[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Scott"
			an := cr[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := CreditCategory(5)
			ac := cr[1].Category
			if ac != ec {
				t.Errorf("Expected category %d, got %d", ec, ac)
			}

			ep := 140
			ap := cr[1].Position
			if ap != ep {
				t.Errorf("Expected position %d, got %d", ep, ap)
			}
		})
	}
}

func TestCreditsSearch(t *testing.T) {
	var creditTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/credits_search.txt", "jim", []FuncOption{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "jim", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "jim", []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range creditTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cr, err := c.Credits.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(cr)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 1342181334
			aID := cr[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Justin - Mom Cody Mark Josh Jim Kerri"
			an := cr[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := 1463521290038
			ac := cr[1].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eu := 1463521290038
			au := cr[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eCat := CreditCategory(5)
			aCat := cr[2].Category
			if aCat != eCat {
				t.Errorf("Expected category %d, got %d", eCat, aCat)
			}

			ep := 365
			ap := cr[2].Position
			if ap != ep {
				t.Errorf("Expected position %d, got %d", ep, ap)
			}
		})
	}
}

func TestCreditsCount(t *testing.T) {
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

			count, err := c.Credits.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestCreditsListFields(t *testing.T) {
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

			fields, err := c.Credits.ListFields()
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
