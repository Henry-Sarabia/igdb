package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestReleaseDateTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	rd := ReleaseDate{}
	typ := reflect.ValueOf(rd).Type()

	err := c.validateStruct(typ, ReleaseDateEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetReleaseDate(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_releasedate.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	rd, err := c.GetReleaseDate(1073)
	if err != nil {
		t.Error(err)
	}

	eID := 1073
	aID := rd.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eg := 475
	ag := rd.Game
	if ag != eg {
		t.Errorf("Expected Game ID %d, got %d", eg, ag)
	}

	ec := 1303935024000
	ac := rd.CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}
}

func TestGetReleaseDates(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_releasedates.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{62566, 32350, 1077}
	rd, err := c.GetReleaseDates(ids)
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(rd)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	ec := DateCategory(2)
	ac := rd[0].Category
	if ac != ec {
		t.Errorf("Expected date category %d, got %d", ec, ac)
	}

	ep := 27
	ap := rd[0].Platform
	if ap != ep {
		t.Errorf("Expected platform %d, got %d", ep, ap)
	}

	ed := 978220800000
	ad := rd[1].Date
	if ad != ed {
		t.Errorf("Expected Unix time in milliseconds %d, got %d", ed, ad)
	}

	ey := 2000
	ay := rd[1].Year
	if ay != ey {
		t.Errorf("Expected year %d, got %d", ey, ay)
	}

	em := 10
	am := rd[2].Month
	if am != em {
		t.Errorf("Expected month %d, got %d", em, am)
	}

	eh := "2010-Oct-26"
	ah := rd[2].Human
	if ah != eh {
		t.Errorf("Expected date '%s', got '%s'", eh, ah)
	}

}
