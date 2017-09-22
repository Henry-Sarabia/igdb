package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getPulseGroupResp = `
[{
	"id": 4943,
	"name": "DOTA 2",
	"category": 1,
	"created_at": 1500399501694,
	"updated_at": 1500399501694,
	"published_at": 1500390000000,
	"tags": [
		1,
		17,
		39,
		268435467,
		268435468,
		268435471,
		268435480,
		536871004,
		536871060,
		536871717,
		536872567,
		805309331,
		1073741827,
		1073741826
	],
	"game": 2963,
	"pulses": [
		224467
	]
}]
`

const getPulseGroupsResp = `
[{
	"id": 13385,
	"name": "Battleborn",
	"category": 1,
	"created_at": 1505670431652,
	"updated_at": 1505672182912,
	"published_at": 1505664360000,
	"tags": [
		1,
		17,
		18,
		27,
		268435461,
		268435468,
		536871378,
		536871458,
		536872244,
		536873085,
		536873286,
		536873287,
		805314055,
		1073741825
	],
	"game": 7687,
	"pulses": [
		268536,
		268550,
		268533
	]
},
{
	"id": 6126,
	"name": "Heroes of the Storm",
	"category": 1,
	"created_at": 1501184656575,
	"updated_at": 1501184656575,
	"published_at": 1501162200000,
	"tags": [
		1,
		268435467,
		536871060,
		536871439,
		536872159,
		536872160,
		536872161,
		536872162,
		536872163,
		536872567,
		805313681,
		1073741827
	],
	"game": 7313,
	"pulses": [
		252265
	]
}]
`

const searchPulseGroupsResp = `
[{
	"id": 3907,
	"name": "League of Legends",
	"category": 1,
	"created_at": 1499759917719,
	"updated_at": 1499759917719,
	"published_at": 1499734800000,
	"tags": [
		1,
		17,
		268435467,
		268435471,
		536871004,
		536871060,
		536872567,
		805306483,
		1073741827
	],
	"game": 115,
	"pulses": [
		153381
	]
},
{
	"id": 5501,
	"name": "League of Legends",
	"category": 1,
	"created_at": 1500792572855,
	"updated_at": 1500792572855,
	"published_at": 1500771647000,
	"tags": [
		1,
		17,
		268435467,
		268435471,
		536871004,
		536871060,
		536872567,
		805306483,
		1073741827
	],
	"game": 115,
	"pulses": [
		250686
	]
},
{
	"id": 6356,
	"name": "League of Legends",
	"category": 1,
	"created_at": 1501359594279,
	"updated_at": 1501359594279,
	"published_at": 1501358634000,
	"tags": [
		1,
		17,
		268435467,
		268435471,
		536871004,
		536871060,
		536872567,
		805306483,
		1073741827
	],
	"game": 115,
	"pulses": [
		252828
	]
}]
`

func TestPulseGroupTypeIntegrity(t *testing.T) {
	c := NewClient()

	pg := PulseGroup{}
	typ := reflect.ValueOf(pg).Type()

	err := c.validateStruct(typ, PulseGroupEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPulseGroup(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPulseGroupResp)
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
	ts, c := startTestServer(http.StatusOK, getPulseGroupsResp)
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
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}
}

func TestSearchPulseGroups(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchPulseGroupsResp)
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
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eu := 1500792572855
	au := pg[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, %d", eu, au)
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
