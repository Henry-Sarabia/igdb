package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestTitleTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	r := Title{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, TitleEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTitle(t *testing.T) {
	var titleTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_title.txt", 16549, ""},
		{"Invalid ID", "test_data/empty.txt", -15000, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 16549, errEndOfJSON.Error()},
	}
	for _, tt := range titleTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ti, err := c.GetTitle(tt.ID)
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

func TestGetTitles(t *testing.T) {
	var titleTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_titles.txt", []int{7839, 25381}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-25000}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{7839, 25381}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{7839, 25381}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range titleTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ti, err := c.GetTitles(tt.IDs, tt.Opts...)
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

func TestSearchTitles(t *testing.T) {
	var titleTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_titles.txt", "graphic", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_titles.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "graphic", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "graphic", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range titleTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ti, err := c.SearchTitles(tt.Qry, tt.Opts...)
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
