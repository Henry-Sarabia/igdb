package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCreditTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	cr := Credit{}
	typ := reflect.ValueOf(cr).Type()

	err := c.validateStruct(typ, CreditEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCredit(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_credit.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	cr, err := c.GetCredit(1342182279)
	if err != nil {
		t.Error(err)
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
}

func TestGetCredits(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_credits.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{1342181334, 1342186852}
	cr, err := c.GetCredits(ids)
	if err != nil {
		t.Error(err)
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
}

func TestSearchCredits(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_credits.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	cr, err := c.SearchCredits("jim")
	if err != nil {
		t.Error(err)
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
}
