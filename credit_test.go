package igdb

import (
	"net/http"
	"testing"
)

const getCreditResp = `
[{
	"id": 1342182279,
	"name": "Michael",
	"created_at": 1424124855090,
	"updated_at": 1424124855090,
	"category": 5,
	"position": 45
}]
`

const getCreditsResp = `
[{
	"id": 1342180316,
	"name": "Scott",
	"created_at": 1414000040656,
	"updated_at": 1414000040656,
	"category": 5,
	"position": 518
},
{
	"id": 1342186852,
	"name": "Thanks  for the inspiration (Scott W.):",
	"created_at": 1463521214105,
	"updated_at": 1463521214105,
	"category": 5,
	"position": 140
}]
`

const searchCreditsResp = `
[{
	"id": 1342181334,
	"name": "Justin - Mom Cody Mark Josh Jim Kerri",
	"created_at": 1417362858828,
	"updated_at": 1417362858828,
	"category": 5,
	"position": 267
},
{
	"id": 1342186871,
	"name": "Jim Gardner for being the MAN!",
	"created_at": 1463521290038,
	"updated_at": 1463521290038,
	"category": 5,
	"position": 288
},
{
	"id": 1342178993,
	"name": "Thanks to Becki Sanders, Sofia Sanders, Sara Sanders, and the USMC. - Jim Sanders",
	"created_at": 1410439937050,
	"updated_at": 1410439937050,
	"category": 5,
	"position": 365
}]
`

func TestGetCredit(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getCreditResp)
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
	ts, c := startTestServer(http.StatusOK, getCreditsResp)
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
	ts, c := startTestServer(http.StatusOK, searchCreditsResp)
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
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eu := 1463521290038
	au := cr[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
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
