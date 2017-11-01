package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPulseGroupTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	pg := PulseGroup{}
	typ := reflect.ValueOf(pg).Type()

	err := c.validateStruct(typ, PulseGroupEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPulseGroup(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_pulsegroup.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	pg, err := c.GetPulseGroup(4943)
	if err != nil {
		t.Error(err)
	}

	en := "DOTA 2"
	an := pg.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 4943
	aID := pg.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	etl := 14
	atl := len(pg.Tags)
	if atl != etl {
		t.Errorf("Expected Tags length %d, got %d", etl, atl)
	}

	ep := 224467
	ap := pg.Pulses[0]
	if ap != ep {
		t.Errorf("Expected Pulse ID %d, got %d", ep, ap)
	}
}

func TestGetPulseGroups(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_pulsegroups.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{2096, 1108}
	pg, err := c.GetPulseGroups(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(pg)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Battleborn"
	an := pg[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	epID := []int{268536, 268550, 268533}
	apID := pg[0].Pulses
	for i := range apID {
		if apID[i] != epID[i] {
			t.Errorf("Expected Pulse ID %d, got %d", epID[i], apID[i])
		}
	}

	ec := 1
	ac := pg[1].Category
	if ac != ec {
		t.Errorf("Expected category %d, got %d", ec, ac)
	}

	eu := 1501184656575
	au := pg[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
	}
}

func TestSearchPulseGroups(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_pulsegroups.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	pg, err := c.SearchPulseGroups("LeagueofLegends")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(pg)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 3907
	aID := pg[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "League of Legends"
	an := pg[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1500792572855
	ac := pg[1].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	eu := 1500792572855
	au := pg[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, %d", eu, au)
	}

	eg := 115
	ag := pg[2].Game
	if ag != eg {
		t.Errorf("Expected Game ID %d, got %d", eg, ag)
	}

	etl := 9
	atl := len(pg[2].Tags)
	if atl != etl {
		t.Errorf("Expected Tags length %d, got %d", etl, atl)
	}
}
