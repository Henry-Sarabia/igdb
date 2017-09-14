package igdb

import (
	"net/http"
	"testing"
)

const getKeywordResp = `
[{
	"id": 2107,
	"name": "space adventure",
	"created_at": 1452513769816,
	"updated_at": 1452513769816,
	"slug": "space-adventure",
	"url": "https://www.igdb.com/categories/space-adventure",
	"games": [
		8506,
		26187,
		25919,
		23905,
		23908,
		27903,
		52205
	]
}]
`

const getKeywordsResp = `
[{
	"id": 2096,
	"name": "humor",
	"created_at": 1452454352817,
	"updated_at": 1452454352817,
	"slug": "humor",
	"url": "https://www.igdb.com/categories/humor",
	"games": [
		15857,
		3364,
		24415
	]
},
{
	"id": 1108,
	"name": "action sports",
	"created_at": 1403518560769,
	"updated_at": 1403518560769,
	"slug": "action-sports",
	"url": "https://www.igdb.com/categories/action-sports",
	"games": [
		6749,
		7591,
		10666,
		14952,
		36899
	]
}]
`

const searchKeywordsResp = `
[{
	"id": 3782,
	"name": "strategy",
	"created_at": 1495537045343,
	"updated_at": 1495537045343,
	"slug": "strategy",
	"url": "https://www.igdb.com/categories/strategy",
	"games": [
		30229,
		223,
		51997,
		46270
	]
},
{
	"id": 3908,
	"name": "historical strategy",
	"created_at": 1499532005267,
	"updated_at": 1499532005267,
	"slug": "historical-strategy",
	"url": "https://www.igdb.com/categories/historical-strategy",
	"games": [
		46076
	]
},
{
	"id": 2547,
	"name": "real time strategy",
	"created_at": 1469623570981,
	"updated_at": 1469623570981,
	"slug": "real-time-strategy--1",
	"url": "https://www.igdb.com/categories/real-time-strategy--1",
	"games": [
		21620,
		27254,
		21221,
		24273,
		27448,
		54723
	]
}]
`

func TestGetKeyword(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getKeywordResp)
	defer ts.Close()

	g, err := c.GetKeyword(2107)
	if err != nil {
		t.Error(err)
	}

	en := "space adventure"
	an := g.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 2107
	aID := g.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/categories/space-adventure")
	aURL := g.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{8506, 26187, 25919, 23905, 23908, 27903, 52205}
	agID := g.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestGetKeywords(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getKeywordsResp)
	defer ts.Close()

	ids := []int{2096, 1108}
	g, err := c.GetKeywords(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "humor"
	an := g[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/categories/humor")
	aURL := g[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eu := 1403518560769
	au := g[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}

	egID := []int{6749, 7591, 10666, 14952, 36899}
	agID := g[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestSearchKeywords(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchKeywordsResp)
	defer ts.Close()

	g, err := c.SearchKeywords("tool")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 3782
	aID := g[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "strategy"
	an := g[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1499532005267
	ac := g[1].CreatedAt
	if ac != ec {
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/categories/historical-strategy")
	aURL := g[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "real-time-strategy--1"
	as := g[2].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egID := []int{21620, 27254, 21221, 24273, 27448, 54723}
	agID := g[2].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}
