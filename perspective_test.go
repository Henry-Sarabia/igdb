package igdb

import (
	"net/http"
	"testing"
)

const getPerspectiveResp = `
[{
	"id": 7,
	"name": "Virtual Reality",
	"created_at": 1462288484243,
	"updated_at": 1462288484243,
	"slug": "virtual-reality",
	"url": "https://www.igdb.com/player_perspectives/virtual-reality",
	"games": [
		8654,
		10724,
		6415,
		5639,
		9254,
		17244,
		18157,
		17986,
		12302
	]
}]
`

const getPerspectivesResp = `
[{
	"id": 6,
	"name": "Aural",
	"created_at": 1413209511809,
	"updated_at": 1413209511809,
	"slug": "aural",
	"url": "https://www.igdb.com/player_perspectives/aural",
	"games": [
		9076,
		2629,
		7698,
		8597,
		8662,
		8612,
		8661,
		9520
	]
},
{
	"id": 3,
	"name": "Bird view",
	"created_at": 1298968714000,
	"updated_at": 1323289214000,
	"slug": "bird-view",
	"url": "https://www.igdb.com/player_perspectives/bird-view",
	"games": [
		6,
		5,
		35,
		36,
		13,
		14,
		17,
		12,
		76
	]
}]
`

const searchPerspectivesResp = `
[{
	"id": 1,
	"name": "First person",
	"created_at": 1298968658000,
	"updated_at": 1323289214000,
	"slug": "first-person",
	"url": "https://www.igdb.com/player_perspectives/first-person",
	"games": [
		1,
		2,
		54,
		3,
		15,
		16,
		11,
		21,
		41,
		42,
		43
	]
},
{
	"id": 2,
	"name": "Third person",
	"created_at": 1298968673000,
	"updated_at": 1323289214000,
	"slug": "third-person",
	"url": "https://www.igdb.com/player_perspectives/third-person",
	"games": [
		37,
		38,
		3,
		15,
		16,
		12,
		76,
		109,
		112,
		113
	]
}]
`

func TestGetPerspective(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPerspectiveResp)
	defer ts.Close()

	p, err := c.GetPerspective(7)
	if err != nil {
		t.Error(err)
	}

	en := "Virtual Reality"
	an := p.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eu := 1462288484243
	au := p.UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}

	eURL := URL("https://www.igdb.com/player_perspectives/virtual-reality")
	aURL := p.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{8654, 10724, 6415, 5639, 9254, 17244, 18157, 17986, 12302}
	agID := p.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestGetPerspectives(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPerspectivesResp)
	defer ts.Close()

	ids := []int{6, 3}
	p, err := c.GetPerspectives(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Aural"
	an := p[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	es := "aural"
	as := p[0].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eID := 3
	aID := p[1].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	egID := []int{6, 5, 35, 36, 13, 14, 17, 12, 76}
	agID := p[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestSearchPerspectives(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchPerspectivesResp)
	defer ts.Close()

	p, err := c.SearchPerspectives("person")
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 1
	aID := p[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	ec := 1298968658000
	ac := p[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/player_perspectives/third-person")
	aURL := p[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "third-person"
	as := p[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}
}
