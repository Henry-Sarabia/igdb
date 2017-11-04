package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCompanyTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	com := Company{}
	typ := reflect.ValueOf(com).Type()

	err := c.validateStruct(typ, CompanyEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCompany(t *testing.T) {
	var companyTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_company.txt", 58, ""},
		{"Invalid ID", "test_data/empty.txt", -999, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 58, errEndOfJSON.Error()},
	}
	for _, tt := range companyTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.GetCompany(tt.ID)
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

func TestGetCompanies(t *testing.T) {
	var companyTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_companies.txt", []int{854, 7260}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-400}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{854, 7260}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{854, 7260}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range companyTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.GetCompanies(tt.IDs, tt.Opts...)
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

func TestSearchCompanies(t *testing.T) {
	var companyTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_companies.txt", "toby fox", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_companies.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "toby fox", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "toby fox", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range companyTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			com, err := c.SearchCompanies(tt.Qry, tt.Opts...)
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
