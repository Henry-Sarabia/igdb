package igdb

import (
	"net/http"
	"testing"
)

const getCollectionResp = `
[{
	"id": 2404,
	"name": "Chocobo",
	"created_at": 1472012155368,
	"updated_at": 1472012155368,
	"slug": "chocobo",
	"url": "https://www.igdb.com/collections/chocobo",
	"games": [
		22896,
		22909,
		22905,
		22903,
		22907,
		22906,
		22898,
		22908,
		22901
	]
}]
`

const getCollectionsResp = `
[{
	"id": 338,
	"name": "Mega Man X",
	"created_at": 1352059968884,
	"updated_at": 1352059968884,
	"slug": "mega-man-x",
	"url": "https://www.igdb.com/collections/mega-man-x",
	"games": [
		1748,
		1751,
		1749,
		1743,
		1747,
		1745,
		1750,
		4001,
		1744,
		1741,
		1742,
		1746
	]
},
{
	"id": 1,
	"name": "Bioshock",
	"created_at": 1298843586000,
	"updated_at": 1323289207000,
	"slug": "bioshock",
	"url": "https://www.igdb.com/collections/bioshock",
	"games": [
		538,
		14543,
		19839,
		20,
		10047,
		21
	]
}]
`

const searchCollectionsResp = `
[{
	"id": 593,
	"name": "Paper Mario",
	"created_at": 1388708545370,
	"updated_at": 1388708545370,
	"slug": "paper-mario",
	"url": "https://www.igdb.com/collections/paper-mario",
	"games": [
		3340,
		2191,
		3350,
		18169,
		3349
	]
},
{
	"id": 597,
	"name": "Mario Tennis",
	"created_at": 1388973829025,
	"updated_at": 1388973829025,
	"slug": "mario-tennis",
	"url": "https://www.igdb.com/collections/mario-tennis",
	"games": [
		3406,
		20386,
		5712,
		3985,
		6504,
		11220
	]
},
{
	"id": 598,
	"name": "Mario Golf",
	"created_at": 1388973830641,
	"updated_at": 1388973830641,
	"slug": "mario-golf",
	"url": "https://www.igdb.com/collections/mario-golf",
	"games": [
		3402,
		3400,
		3404,
		3401,
		3405,
		3403
	]
}]
`

func TestGetCollection(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getCollectionResp)
	defer ts.Close()

	col, err := c.GetCollection(2404)
	if err != nil {
		t.Error(err)
	}

	en := "Chocobo"
	an := col.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/collections/chocobo")
	aURL := col.URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eID := []int{22896, 22909, 22905, 22903, 22907, 22906, 22898, 22908, 22901}
	aID := col.Games
	for i := range aID {
		if aID[i] != eID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
		}
	}
}

func TestGetCollections(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getCollectionsResp)
	defer ts.Close()

	ids := []int{338, 1}
	col, err := c.GetCollections(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(col)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Mega Man X"
	an := col[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1352059968884
	ac := col[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/collections/bioshock")
	aURL := col[1].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eID := []int{538, 14543, 19839, 20, 10047, 21}
	aID := col[1].Games
	for i := range aID {
		if aID[i] != eID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
		}
	}
}

func TestSearchCollections(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchCollectionsResp)
	defer ts.Close()

	col, err := c.SearchCollections("mario")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(col)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 593
	aID := col[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "Paper Mario"
	an := col[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1388973829025
	ac := col[1].CreatedAt
	if ac != ec {
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
	}

	es := "mario-tennis"
	as := col[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eURL := URL("https://www.igdb.com/collections/mario-golf")
	aURL := col[2].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eIDs := []int{3402, 3400, 3404, 3401, 3405, 3403}
	aIDs := col[2].Games
	for i := range aIDs {
		if aIDs[i] != eIDs[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eIDs[i], aIDs[i])
		}
	}
}
