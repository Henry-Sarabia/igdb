package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getCharacterResp = `
[{
	"id": 10617,
	"name": "Princess Zelda",
	"created_at": 1492752468377,
	"updated_at": 1492754718998,
	"slug": "princess-zelda",
	"url": "https://www.igdb.com/characters/princess-zelda",
	"mug_shot": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/q64bgyomnbk5spyikkex.jpg",
		"cloudinary_id": "q64bgyomnbk5spyikkex",
		"width": 204,
		"height": 479
	},
	"gender": 1,
	"species": 1
}]
`

const getCharactersResp = `
[{
	"id": 3726,
	"name": "Mario",
	"created_at": 1428832054608,
	"updated_at": 1468601612724,
	"slug": "mario",
	"url": "https://www.igdb.com/characters/mario",
	"mug_shot": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/iurhmhenrrsdnsc4zbva.jpg",
		"cloudinary_id": "iurhmhenrrsdnsc4zbva",
		"width": 1840,
		"height": 3784
	},
	"gender": 0,
	"akas": [
		"Super Mario"
	],
	"species": 1,
	"people": [
		28769,
		132118
	],
	"games": [
		1074
	]
},
{
	"id": 9580,
	"name": "Samus Aran",
	"created_at": 1472328217237,
	"updated_at": 1492754716428,
	"slug": "samus-aran",
	"url": "https://www.igdb.com/characters/samus-aran",
	"mug_shot": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/ilgqsndahl8sjk5navaw.jpg",
		"cloudinary_id": "ilgqsndahl8sjk5navaw",
		"width": 352,
		"height": 480
	},
	"gender": 1,
	"species": 1,
	"people": [
		25893
	],
	"games": [
		1113
	]
}]
`

const searchCharacterResp = `
[{
	"id": 2168,
	"name": "Snake",
	"created_at": 1413724358001,
	"updated_at": 1472328003683,
	"slug": "snake",
	"url": "https://www.igdb.com/characters/snake",
	"gender": 0,
	"species": 1,
	"people": [
		34569,
		85790
	],
	"games": [
		5328,
		376,
		379,
		1985,
		382
	]
},
{
	"id": 5352,
	"name": "Solid Snake",
	"created_at": 1438931494248,
	"updated_at": 1438931494561,
	"slug": "solid-snake",
	"url": "https://www.igdb.com/characters/solid-snake",
	"people": [
		85790
	],
	"games": [
		375
	]
},
{
	"id": 5378,
	"name": "Old Snake",
	"created_at": 1440480929228,
	"updated_at": 1440480929498,
	"slug": "old-snake",
	"url": "https://www.igdb.com/characters/old-snake",
	"people": [
		85790
	],
	"games": [
		380
	]
}]
`

func TestCharacterTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	ch := Character{}
	typ := reflect.ValueOf(ch).Type()

	err := c.validateStruct(typ, CharacterEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCharacter(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getCharacterResp)
	defer ts.Close()

	ch, err := c.GetCharacter(10617)
	if err != nil {
		t.Error(err)
	}

	eID := 10617
	aID := ch.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "Princess Zelda"
	an := ch.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'\n", en, an)
	}

	eh := 479
	ah := ch.Mugshot.Height
	if ah != eh {
		t.Errorf("Expected height of %d, got %d\n", eh, ah)
	}
}

func TestGetCharacters(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getCharactersResp)
	defer ts.Close()

	ids := []int{3726, 9580}
	ch, err := c.GetCharacters(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(ch)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Mario"
	an := ch[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eAKA := "Super Mario"
	aAKA := ch[0].AKAs[0]
	if eAKA != aAKA {
		t.Errorf("Expected AKA '%s', got '%s'", eAKA, aAKA)
	}

	eID := 9580
	aID := ch[1].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/characters/samus-aran")
	aURL := ch[1].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

}

func TestSearchCharacters(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchCharacterResp)
	defer ts.Close()

	ch, err := c.SearchCharacters("snake")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(ch)
	if al != el {
		t.Errorf("Expected length of %d, got %d\n", el, al)
	}

	eURL := URL("https://www.igdb.com/characters/snake")
	aURL := ch[0].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'\n", eURL, aURL)
	}

	en := "Solid Snake"
	an := ch[1].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'\n", en, an)
	}

	eID := 5378
	aID := ch[2].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d\n", eID, aID)
	}
}
