package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getGenreResp = `
[{
	"id": 8,
	"name": "Platform",
	"created_at": 1297639288000,
	"updated_at": 1323289215000,
	"slug": "platform",
	"url": "https://www.igdb.com/genres/platform",
	"games": [
		358,
		360,
		452,
		337,
		454,
		185,
		190,
		71,
		72,
		217
	]
}]
`

const getGenresResp = `
[{
	"id": 5,
	"name": "Shooter",
	"created_at": 1297639288000,
	"updated_at": 1323289215000,
	"slug": "shooter",
	"url": "https://www.igdb.com/genres/shooter",
	"games": [
		1,
		2,
		3,
		15,
		16,
		20,
		21,
		41,
		42
	]
},
{
	"id": 10,
	"name": "Racing",
	"created_at": 1297639288000,
	"updated_at": 1323289215000,
	"slug": "racing",
	"url": "https://www.igdb.com/genres/racing",
	"games": [
		143,
		154,
		177,
		390,
		422,
		90,
		91,
		92,
		99
	]
}]
`

const searchGenresResp = `
[{
	"id": 15,
	"name": "Strategy",
	"created_at": 1297639288000,
	"updated_at": 1323289215000,
	"slug": "strategy",
	"url": "https://www.igdb.com/genres/strategy",
	"games": [
		67,
		6,
		205,
		175,
		82,
		432,
		435,
		334
	]
},
{
	"id": 16,
	"name": "Turn-based strategy (TBS)",
	"created_at": 1297678340000,
	"updated_at": 1323289215000,
	"slug": "turn-based-strategy-tbs",
	"url": "https://www.igdb.com/genres/turn-based-strategy-tbs",
	"games": [
		8,
		9,
		13,
		14,
		17,
		25,
		26,
		67
	]
},
{
	"id": 11,
	"name": "Real Time Strategy (RTS)",
	"created_at": 1297639288000,
	"updated_at": 1323289215000,
	"slug": "real-time-strategy-rts",
	"url": "https://www.igdb.com/genres/real-time-strategy-rts",
	"games": [
		34,
		35,
		36,
		133,
		151,
		159,
		289
	]
}]
`

func TestGenreTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	g := Genre{}
	typ := reflect.ValueOf(g).Type()

	err := c.validateStruct(typ, GenreEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetGenre(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getGenreResp)
	defer ts.Close()

	g, err := c.GetGenre(8)
	if err != nil {
		t.Error(err)
	}

	en := "Platform"
	an := g.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 8
	aID := g.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/genres/platform")
	aURL := g.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{358, 360, 452, 337, 454, 185, 190, 71, 72, 217}
	agID := g.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestGetGenres(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getGenresResp)
	defer ts.Close()

	ids := []int{5, 10}
	g, err := c.GetGenres(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Shooter"
	an := g[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/genres/shooter")
	aURL := g[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eu := 1323289215000
	au := g[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}

	egID := []int{143, 154, 177, 390, 422, 90, 91, 92, 99}
	agID := g[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestSearchGenres(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchGenresResp)
	defer ts.Close()

	g, err := c.SearchGenres("strategy")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 15
	aID := g[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "Strategy"
	an := g[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1297678340000
	ac := g[1].CreatedAt
	if ac != ec {
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/genres/turn-based-strategy-tbs")
	aURL := g[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "real-time-strategy-rts"
	as := g[2].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egID := []int{34, 35, 36, 133, 151, 159, 289}
	agID := g[2].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}
