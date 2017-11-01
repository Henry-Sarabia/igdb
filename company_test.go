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
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_company.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	com, err := c.GetCompany(58)
	if err != nil {
		t.Error(err)
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
}

func TestGetCompanies(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_companies.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{854, 7260}
	com, err := c.GetCompanies(ids)
	if err != nil {
		t.Error(err)
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
}

func TestSearchCompanies(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_companies.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	com, err := c.SearchCompanies("mario")
	if err != nil {
		t.Error(err)
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
}
