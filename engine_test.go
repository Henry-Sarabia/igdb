package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getEngineResp = `
[{
	"id": 26,
	"name": "RAGE",
	"created_at": 1402868699497,
	"updated_at": 1499128991855,
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/ssdjdq31lxtqqgpfkycg.jpg",
		"cloudinary_id": "ssdjdq31lxtqqgpfkycg",
		"width": 476,
		"height": 201
	},
	"slug": "rage",
	"url": "https://www.igdb.com/game_engines/rage",
	"games": [
		731,
		434,
		7071,
		1020,
		960,
		2541,
		3174,
		3265,
		1969
	],
	"platforms": [
		5,
		9,
		12,
		48,
		49,
		6
	],
	"companies": [
		364
	]
}]
`

const getEnginesResp = `
[{
	"id": 9,
	"name": "Anvil",
	"created_at": 1399589546212,
	"updated_at": 1420976957180,
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/ugq9i3qhp1m1bvitr4h3.jpg",
		"cloudinary_id": "ugq9i3qhp1m1bvitr4h3",
		"width": 840,
		"height": 315
	},
	"slug": "anvil",
	"url": "https://www.igdb.com/game_engines/anvil",
	"games": [
		127,
		113,
		537,
		1855,
		8216,
		8217,
		2468
	],
	"platforms": [
		6,
		9,
		48,
		46,
		41,
		12,
		49
	],
	"companies": [
		38
	]
},
{
	"id": 22,
	"name": "UbiArt Framework",
	"created_at": 1402435973694,
	"updated_at": 1466690712748,
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/pbscffthi6uhqxs3ubjk.jpg",
		"cloudinary_id": "pbscffthi6uhqxs3ubjk",
		"width": 1070,
		"height": 1377
	},
	"slug": "ubiart-framework",
	"url": "https://www.igdb.com/game_engines/ubiart-framework",
	"games": [
		7327,
		981,
		1968,
		4756,
		14533,
		19726
	]
}]
`

const searchEnginesResp = `
[{
	"id": 12,
	"name": "Telltale Tool",
	"created_at": 1399826402127,
	"updated_at": 1492514717250,
	"slug": "telltale-tool",
	"url": "https://www.igdb.com/game_engines/telltale-tool",
	"games": [
		6707,
		1871,
		3097,
		6778,
		2993,
		7610,
		8339,
		11307,
		9463,
		19268,
		4781,
		7025,
		3232
	]
},
{
	"id": 114,
	"name": "Crystal Tools",
	"created_at": 1415283742080,
	"updated_at": 1499989864468,
	"slug": "crystal-tools",
	"url": "https://www.igdb.com/game_engines/crystal-tools",
	"games": [
		389,
		384,
		2449
	]
}]
`

func TestEngineTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	e := Engine{}
	typ := reflect.ValueOf(e).Type()

	err := c.validateStruct(typ, EngineEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetEngine(t *testing.T) {
	ts, c := testServerString(http.StatusOK, getEngineResp)
	defer ts.Close()

	eng, err := c.GetEngine(26)
	if err != nil {
		t.Error(err)
	}

	en := "RAGE"
	an := eng.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ew := 476
	aw := eng.Logo.Width
	if aw != ew {
		t.Errorf("Expected width of %d, got %d", ew, aw)
	}

	egID := []int{731, 434, 7071, 1020, 960, 2541, 3174, 3265, 1969}
	agID := eng.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestGetEngines(t *testing.T) {
	ts, c := testServerString(http.StatusOK, getEnginesResp)
	defer ts.Close()

	ids := []int{9, 22}
	eng, err := c.GetEngines(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(eng)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Anvil"
	an := eng[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/game_engines/anvil")
	aURL := eng[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	ecID := "pbscffthi6uhqxs3ubjk"
	acID := eng[1].Logo.ID
	if acID != ecID {
		t.Errorf("Expected Cloudinary ID '%s', got '%s'", ecID, acID)
	}

	egID := []int{7327, 981, 1968, 4756, 14533, 19726}
	agID := eng[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestSearchEngines(t *testing.T) {
	ts, c := testServerString(http.StatusOK, searchEnginesResp)
	defer ts.Close()

	eng, err := c.SearchEngines("tool")
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(eng)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Telltale Tool"
	an := eng[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eu := 1492514717250
	au := eng[0].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
	}

	eURL := URL("https://www.igdb.com/game_engines/crystal-tools")
	aURL := eng[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{389, 384, 2449}
	agID := eng[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}
