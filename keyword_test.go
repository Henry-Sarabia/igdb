package igdb

import (
	"net/http"
	"reflect"
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

func TestKeywordTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	k := Keyword{}
	typ := reflect.ValueOf(k).Type()

	err := c.validateStruct(typ, KeywordEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetKeyword(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getKeywordResp)
	defer ts.Close()

	kw, err := c.GetKeyword(2107)
	if err != nil {
		t.Error(err)
	}

	en := "space adventure"
	an := kw.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 2107
	aID := kw.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/categories/space-adventure")
	aURL := kw.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{8506, 26187, 25919, 23905, 23908, 27903, 52205}
	agID := kw.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestGetKeywords(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getKeywordsResp)
	defer ts.Close()

	ids := []int{2096, 1108}
	kw, err := c.GetKeywords(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(kw)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "humor"
	an := kw[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/categories/humor")
	aURL := kw[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eu := 1403518560769
	au := kw[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
	}

	egID := []int{6749, 7591, 10666, 14952, 36899}
	agID := kw[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestSearchKeywords(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchKeywordsResp)
	defer ts.Close()

	kw, err := c.SearchKeywords("strategy")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(kw)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 3782
	aID := kw[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "strategy"
	an := kw[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1499532005267
	ac := kw[1].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/categories/historical-strategy")
	aURL := kw[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "real-time-strategy--1"
	as := kw[2].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egID := []int{21620, 27254, 21221, 24273, 27448, 54723}
	agID := kw[2].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}
