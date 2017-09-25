package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getGameModeResp = `
[{
	"id": 1,
	"name": "Single player",
	"created_at": 1298968834000,
	"updated_at": 1323289216000,
	"slug": "single-player",
	"url": "https://www.igdb.com/game_modes/single-player",
	"games": [
		37,
		1,
		2,
		60,
		54,
		6,
		38,
		3,
		5,
		35,
		36
	]
}]
`

const getGameModesResp = `
[{
	"id": 3,
	"name": "Co-operative",
	"created_at": 1298968887000,
	"updated_at": 1323289216000,
	"slug": "co-operative",
	"url": "https://www.igdb.com/game_modes/co-operative",
	"games": [
		84,
		115,
		120,
		122,
		124,
		125,
		126,
		141,
		83,
		498,
		470
	]
},
{
	"id": 4,
	"name": "Split screen",
	"created_at": 1298968900000,
	"updated_at": 1323289216000,
	"slug": "split-screen",
	"url": "https://www.igdb.com/game_modes/split-screen",
	"games": [
		141,
		143,
		176,
		211,
		95,
		545,
		642,
		847,
		784
	]
}]
`

const searchGameModesResp = `
[{
	"id": 2,
	"name": "Multiplayer",
	"created_at": 1298968853000,
	"updated_at": 1323289216000,
	"slug": "multiplayer",
	"url": "https://www.igdb.com/game_modes/multiplayer",
	"games": [
		35,
		36,
		21,
		84,
		113,
		115,
		119,
		123,
		124
	]
},
{
	"id": 5,
	"name": "Massively Multiplayer Online (MMO)",
	"created_at": 1298969041000,
	"updated_at": 1323289216000,
	"slug": "massively-multiplayer-online-mmo",
	"url": "https://www.igdb.com/game_modes/massively-multiplayer-online-mmo",
	"games": [
		114,
		123,
		206,
		145,
		215,
		228,
		227,
		229,
		270,
		282
	]
}]
`

func TestGameModeTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	g := GameMode{}
	typ := reflect.ValueOf(g).Type()

	err := c.validateStruct(typ, GameModeEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetGameMode(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getGameModeResp)
	defer ts.Close()

	g, err := c.GetGameMode(1)
	if err != nil {
		t.Error(err)
	}

	en := "Single player"
	an := g.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/game_modes/single-player")
	aURL := g.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{37, 1, 2, 60, 54, 6, 38, 3, 5, 35, 36}
	agID := g.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestGetGameModes(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getGameModesResp)
	defer ts.Close()

	ids := []int{3, 4}
	g, err := c.GetGameModes(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Co-operative"
	an := g[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/game_modes/co-operative")
	aURL := g[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "split-screen"
	as := g[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egID := []int{141, 143, 176, 211, 95, 545, 642, 847, 784}
	agID := g[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestSearchGameModes(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchGameModesResp)
	defer ts.Close()

	g, err := c.SearchGameModes("multiplayer")
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 2
	aID := g[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "Multiplayer"
	an := g[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1298968853000
	ac := g[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/game_modes/massively-multiplayer-online-mmo")
	aURL := g[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "massively-multiplayer-online-mmo"
	as := g[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egID := []int{114, 123, 206, 145, 215, 228, 227, 229, 270, 282}
	agID := g[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}
