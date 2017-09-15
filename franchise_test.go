package igdb

import (
	"net/http"
	"testing"
)

const getFranchiseResp = `
[{
	"id": 596,
	"name": "The Legend of Zelda",
	"created_at": 1439069965468,
	"updated_at": 1439069965468,
	"slug": "the-legend-of-zelda",
	"url": "https://www.igdb.com/franchises/the-legend-of-zelda",
	"games": [
		11607,
		1036,
		18017,
		18066,
		7346,
		25840,
		8534,
		41829,
		1628,
		9602
	]
}]
`

const getFranchisesResp = `
[{
	"id": 884,
	"name": "Red Dead",
	"created_at": 1476796671177,
	"updated_at": 1476796671177,
	"slug": "red-dead",
	"url": "https://www.igdb.com/franchises/red-dead",
	"games": [
		25076,
		434,
		1969
	]
},
{
	"id": 907,
	"name": "Dynasty Warriors",
	"created_at": 1479418914178,
	"updated_at": 1479418914178,
	"slug": "dynasty-warriors",
	"url": "https://www.igdb.com/franchises/dynasty-warriors",
	"games": [
		25639,
		26546,
		26180,
		28368,
		44157
	]
}]
`

const searchFranchisesResp = `
[{
	"id": 128,
	"name": "Super Man",
	"created_at": 1381669592350,
	"updated_at": 1381669592350,
	"slug": "super-man",
	"url": "https://www.igdb.com/franchises/super-man",
	"games": [
		3005,
		4190,
		6183,
		6182
	]
},
{
	"id": 860,
	"name": "Super Mario",
	"created_at": 1473419646160,
	"updated_at": 1473419646160,
	"slug": "super-mario",
	"url": "https://www.igdb.com/franchises/super-mario",
	"games": [
		41420
	]
},
{
	"id": 327,
	"name": "Marvel Super Hero Squad",
	"created_at": 1391704834814,
	"updated_at": 1391704834814,
	"slug": "marvel-super-hero-squad",
	"url": "https://www.igdb.com/franchises/marvel-super-hero-squad",
	"games": [
		4997,
		5188
	]
}]
`

func TestGetFranchise(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getFranchiseResp)
	defer ts.Close()

	f, err := c.GetFranchise(596)
	if err != nil {
		t.Error(err)
	}

	en := "The Legend of Zelda"
	an := f.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/franchises/the-legend-of-zelda")
	aURL := f.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{11607, 1036, 18017, 18066, 7346, 25840, 8534, 41829, 1628, 9602}
	agID := f.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestGetFranchises(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getFranchisesResp)
	defer ts.Close()

	ids := []int{9, 22}
	f, err := c.GetFranchises(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(f)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Red Dead"
	an := f[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/franchises/red-dead")
	aURL := f[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eu := 1479418914178
	au := f[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}

	egID := []int{25639, 26546, 26180, 28368, 44157}
	agID := f[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestSearchFranchises(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchFranchisesResp)
	defer ts.Close()

	f, err := c.SearchFranchises("super")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(f)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Super Man"
	an := f[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1381669592350
	ac := f[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eID := 860
	aID := f[1].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/franchises/super-mario")
	aURL := f[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "marvel-super-hero-squad"
	as := f[2].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egID := []int{4997, 5188}
	agID := f[2].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}
