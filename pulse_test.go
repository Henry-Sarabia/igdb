package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPulseTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	p := Pulse{}
	typ := reflect.ValueOf(p).Type()

	err := c.validateStruct(typ, PulseEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPulse(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_pulse.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	p, err := c.GetPulse(145346)
	if err != nil {
		t.Error(err)
	}

	et := "Nintendo announces new Mario, Zelda amiibo"
	at := p.Title
	if at != et {
		t.Errorf("Expected title '%s', got '%s'", et, at)
	}

	ep := 2
	ap := p.PulseSource
	if ap != ep {
		t.Errorf("Expected Pulse Source %d, got %d", ep, ap)
	}

	eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/i1fti435exzyu1ydftu4.jpg")
	aURL := p.PulseImage.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	etID := []Tag{1, 17, 38, 268435468, 268435487, 536871422, 536872221}
	atID := p.Tags
	for i := range atID {
		if atID[i] != etID[i] {
			t.Errorf("Expected Tag ID %d, got %d\n", etID[i], atID[i])
		}
	}
}

func TestGetPulses(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_pulses.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{132354, 257394, 109415}
	p, err := c.GetPulses(ids)
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	et := "Battleborn: a great game with a fatal flaw"
	at := p[0].Title
	if at != et {
		t.Errorf("Expected title '%s', got '%s'", et, at)
	}

	ea := "Darren Nakamura"
	aa := p[0].Author
	if aa != ea {
		t.Errorf("Expected slug '%s', got '%s'", ea, aa)
	}

	eUID := "5fc0c5269a2aa7887d1a2c13a27c5bd2"
	aUID := p[1].UID
	if aUID != eUID {
		t.Errorf("Expected ID '%s', got '%s'", eUID, aUID)
	}

	eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/ibmrifg0uxp8w3y6hfzo.jpg")
	aURL := p[1].PulseImage.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	ep := 8
	ap := p[2].PulseSource
	if ap != ep {
		t.Errorf("Expected Pulse Source %d, got %d", ep, ap)
	}

	etID := []Tag{1, 18, 27, 268435461, 268435468, 268435471, 536871198}
	atID := p[2].Tags
	for i := range atID {
		if atID[i] != etID[i] {
			t.Errorf("Expected Tag ID %d, got %d\n", etID[i], atID[i])
		}
	}
}

func TestSearchPulses(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_pulses.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	p, err := c.SearchPulses("megaman")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eCat := 10
	aCat := p[0].Category
	if aCat != eCat {
		t.Errorf("Expected category %d, got %d", eCat, aCat)
	}

	ec := 1502176226691
	ac := p[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	et := "Retroid talks Mega Man tonight"
	at := p[1].Title
	if at != et {
		t.Errorf("Expected title '%s', got '%s'", et, at)
	}

	eURL := URL("http://feedproxy.google.com/~r/Destructoid/~3/ChzHjztL10M/retroid-talks-mega-man-tonight-456282.phtml")
	aURL := p[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	ep := 1479339000000
	ap := p[2].PublishedAt
	if ap != ep {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ep, ap)
	}

	eID := "y4wwcqqbkuyeteoq4l2n"
	aID := p[2].PulseImage.ID
	if aID != eID {
		t.Errorf("Expected ID '%s', got '%s'", eID, aID)
	}
}
